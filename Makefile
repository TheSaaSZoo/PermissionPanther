proto:
	protoc --go_out=. --go_opt=paths=source_relative \
	--go-grpc_out=. --go-grpc_opt=paths=source_relative \
	--js_out=import_style=commonjs,binary:./js \
	-I=./ pb/*.proto

# https://medium.com/blokur/how-to-implement-a-grpc-client-and-server-in-typescript-fa3ac807855e
jsclient:
	grpc_tools_node_protoc --js_out=import_style=commonjs,binary:./js/src \
	--plugin=protoc-gen-ts=./node_modules/.bin/protoc-gen-ts \
	--grpc_out=grpc_js:./js/src \
	--ts_out=./js/src \
	pb/*.proto
