.PHONY: proto
proto:
	protoc \
		--go_out=common/api/generated \
        --go-grpc_out=common/api/generated \
		--go_opt=module=github.com/justclimber/fda/common \
        --go-grpc_opt=module=github.com/justclimber/fda/common \
	  	--go-grpc_opt=require_unimplemented_servers=false \
        common/api/proto/api.proto