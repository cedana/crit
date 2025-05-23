package phaul

import (
	"errors"
	"fmt"

	"github.com/cedana/go-criu/v7"
	"github.com/cedana/go-criu/v7/crit"
	"github.com/cedana/go-criu/v7/crit/images/stats"
	"github.com/cedana/go-criu/v7/rpc"
	"google.golang.org/protobuf/proto"
)

const (
	minPagesWritten uint64 = 64
	maxIters        int    = 8
	maxGrowDelta    int64  = 32
)

// Client struct
type Client struct {
	local  Local
	remote Remote
	cfg    Config
}

// MakePhaulClient function
// Main entry point. Caller should create the client object by
// passing here local, remote and comm. See comment in corresponding
// interfaces/structs for explanation.
//
// Then call client.Migrate() and enjoy :)
func MakePhaulClient(l Local, r Remote, c Config) (*Client, error) {
	return &Client{local: l, remote: r, cfg: c}, nil
}

func isLastIter(iter int, stats *stats.DumpStatsEntry, prevStats *stats.DumpStatsEntry) bool {
	if iter >= maxIters {
		fmt.Printf("`- max iters reached\n")
		return true
	}

	pagesWritten := stats.GetPagesWritten()
	if pagesWritten < minPagesWritten {
		fmt.Printf("`- tiny pre-dump (%d) reached\n", int(pagesWritten))
		return true
	}

	pagesDelta := int64(pagesWritten) - int64(prevStats.GetPagesWritten())
	if pagesDelta >= maxGrowDelta {
		fmt.Printf("`- grow iter (%d) reached\n", int(pagesDelta))
		return true
	}

	return false
}

// Migrate function
func (pc *Client) Migrate() (retErr error) {
	criu := criu.MakeCriu()
	psi := rpc.CriuPageServerInfo{
		Fd: proto.Int32(int32(pc.cfg.Memfd)),
	}
	opts := &rpc.CriuOpts{
		Pid:      proto.Int32(int32(pc.cfg.Pid)),
		LogLevel: proto.Int32(4),
		LogFile:  proto.String("pre-dump.log"),
		Ps:       &psi,
	}

	err := criu.Prepare()
	if err != nil {
		return err
	}

	defer func() {
		// append any cleanup errors to the returned error
		err := criu.Cleanup()
		if err != nil {
			retErr = errors.Join(retErr, err)
		}
	}()

	imgs, err := preparePhaulImages(pc.cfg.Wdir)
	if err != nil {
		return err
	}
	prevStats := &stats.DumpStatsEntry{}
	iter := 0

	for {
		err = pc.remote.StartIter()
		if err != nil {
			return err
		}

		prevP := imgs.lastImagesDir()
		imgDir, err := imgs.openNextDir()
		if err != nil {
			return err
		}

		opts.ImagesDirFd = proto.Int32(int32(imgDir.Fd()))
		if prevP != "" {
			opts.ParentImg = proto.String(prevP)
		}

		err = criu.PreDump(opts, nil)
		imgDir.Close()
		if err != nil {
			return err
		}

		iter++

		err = pc.remote.StopIter()
		if err != nil {
			return err
		}

		// Get dump statistics with crit
		stats, err := crit.GetDumpStats(imgDir.Name())
		if err != nil {
			return err
		}

		if isLastIter(iter, stats, prevStats) {
			break
		}

		prevStats = stats
	}

	err = pc.remote.StartIter()
	if err == nil {
		prevP := imgs.lastImagesDir()
		err = pc.local.DumpCopyRestore(criu, pc.cfg, prevP)
		err2 := pc.remote.StopIter()
		if err == nil {
			err = err2
		}
	}

	return err
}
