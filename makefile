PROTO_DIR=proto
PROTO_FILES=$(wildcard $(PROTO_DIR)/*.proto)
OUT_DIR=.

GENERATE_PROTO=protoc \
    -I $(PROTO_DIR) \
    --go_out=$(OUT_DIR) \
    --go-grpc_out=$(OUT_DIR)

.PHONY: all proto clean

all: proto

proto:
	$(GENERATE_PROTO) $(PROTO_FILES)

clean:
	rm -f internal/proto/*.pb.go
