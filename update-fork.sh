find . -name "opts.proto" -exec sed -i 's/(criu)/(criu_cedana)/g' {} +
find . -name "*.proto" -exec sed -i 's/(criu)/(criu_cedana)/g' {} +
make -C scripts/proto-gen

old_name="github.com/checkpoint-restore/go-criu/v7"
new_name="github.com/cedana/go-criu/v7"

go mod edit -module "${new_name}"
find . -type f -name '*.go' -exec sed -i -e "s,\"${old_name}/,\"${new_name}/,g" {} \;
find . -type f -name '*.go' -exec sed -i -e "s,\"${old_name}\",\"${new_name}\",g" {} \;
