#!/bin/sh
mkdir -p bin;

echo "GOOS and GOARCH setting"
GOOS="$(go env GOOS)"
GOARCH="$(go env GOARCH)"

echo "building teleport..."
cd ./teleport;
GOOS=$GOOS GOARCH=$GOARCH go build -o ../bin/tp_server benchmark.pb.go tp_server.go;
GOOS=$GOOS GOARCH=$GOARCH go build -o ../bin/tp_client benchmark.pb.go tp_client.go;
GOOS=$GOOS GOARCH=$GOARCH go build -o ../bin/tp_mclient  benchmark.pb.go tp_mclient.go;
echo "bin/teleport OK"

echo "building grpc..."
cd ../grpc;
GOOS=$GOOS GOARCH=$GOARCH go build -o ../bin/grpc_server  grpc_benchmark.pb.go grpc_server.go;
GOOS=$GOOS GOARCH=$GOARCH go build -o ../bin/grpc_client  grpc_benchmark.pb.go grpc_client.go;
GOOS=$GOOS GOARCH=$GOARCH go build -o ../bin/grpc_mclient  grpc_benchmark.pb.go grpc_mclient.go;
echo "bin/grpc OK"

echo "building rpcx..."
cd ../rpcx;
GOOS=$GOOS GOARCH=$GOARCH go build -o ../bin/rpcx_server  benchmark.pb.go rpcx_server.go;
GOOS=$GOOS GOARCH=$GOARCH go build -o ../bin/rpcx_client  benchmark.pb.go rpcx_client.go;
GOOS=$GOOS GOARCH=$GOARCH go build -o ../bin/rpcx_mclient  benchmark.pb.go rpcx_mclient.go;
echo "bin/rpcx OK"

cd ../;
