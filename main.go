package main

import "github.com/danthegoodman1/PermissionPanther/scylla"

func main() {
	scylla.DBConfig()
	scylla.DBConnectWithKeyspace()
}
