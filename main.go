package main

import (
	"github.com/danthegoodman1/PermissionPanther/logger"
	"github.com/danthegoodman1/PermissionPanther/scylla"
)

func main() {
	logger.ConfigureLogger()
	scylla.DBConfig()
	scylla.DBConnectWithKeyspace()
}
