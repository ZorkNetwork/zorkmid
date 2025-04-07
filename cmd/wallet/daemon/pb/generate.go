//go:generate echo "Generating wallet messages..."
//go:generate protoc --go_out=. --go-grpc_out=. --go_opt=paths=source_relative --go-grpc_opt=paths=source_relative walletd.proto

package pb
