package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/danthegoodman1/PermissionPanther/logger"
	"github.com/danthegoodman1/PermissionPanther/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var (
	GRPCServer *grpc.Server
)

type server struct {
	pb.UnimplementedPermissionPantherServer
}

func StartGRPCServer(port string) {
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%s", port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	var opts []grpc.ServerOption

	GRPCServer = grpc.NewServer(opts...)
	pb.RegisterPermissionPantherServer(GRPCServer, &server{})
	logger.Info("Starting BILLABULL gRPC Server on port ", port)
	err = GRPCServer.Serve(lis)
	if err != nil {
		logger.Error("Error closing grpc server:")
		logger.Error(err.Error())
	}
}

func (server) CheckDirectPermission(ctx context.Context, in *pb.CheckDirectReq) (out *pb.CheckDirectRes, err error) {
	// TODO: Intercepted auth
	// TEMP AUTH
	if in.Key != "thisisasupersecretkeythatyouwillneverguesshahahahahagoodluckidiothackers" {
		err = status.Error(codes.PermissionDenied, "Permission denied")
	}

	foundAt := CheckPermissions("nspc", in.Object, in.Permission, in.Entity, 0, 5)
	if foundAt >= 0 {
		out.Valid = true
		out.Recursion = int32(foundAt)
	} else {
		out.Valid = false
		out.Recursion = -1
	}
	return
}
