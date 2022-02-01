proto:
	protoc --go_out=. --go_opt=paths=source_relative \
	--go-grpc_out=. --go-grpc_opt=paths=source_relative \
	--js_out=import_style=commonjs,binary:./js \
	-I=./ pb/*.proto
