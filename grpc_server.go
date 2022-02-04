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
	logger.Info("Starting BILLABULL gRPC Server on port %s", port)
	err = GRPCServer.Serve(lis)
	if err != nil {
		logger.Error("Error closing grpc server:")
		logger.Error(err.Error())
	}
}

func (server) CheckDirectPermission(ctx context.Context, in *pb.CheckDirectReq) (out *pb.CheckDirectRes, err error) {
	out = &pb.CheckDirectRes{}
	// TODO: Intercepted auth
	// TEMP AUTH
	if in.Key != "thisisasupersecretkeythatyouwillneverguesshahahahahagoodluckidiothackers" {
		err = status.Error(codes.PermissionDenied, "Permission denied")
		return
	}

	// First check explicit deny if provided
	if in.DenyPermission != "" {
		directChan := make(chan bool)
		go CheckPermissionDirect(directChan, "nspc", in.Object, in.DenyPermission, in.Entity)
		if <-directChan {
			out.Valid = false
			out.Recursion = 0
			return
		}
	}

	if in.Recursive {
		foundAt := CheckPermissions("nspc", in.Object, in.Permission, in.Entity, 0, MAX_RECURSIONS)
		if foundAt >= 0 {
			out.Valid = true
			out.Recursion = int32(foundAt)
		} else {
			out.Valid = false
			out.Recursion = -1
		}
	} else {
		directChan := make(chan bool)
		go CheckPermissionDirect(directChan, "nspc", in.Object, in.Permission, in.Entity)
		isValid := <-directChan
		out.Valid = isValid
		out.Recursion = 0
	}
	return
}

func (server) ListEntityRelations(ctx context.Context, in *pb.ListEntityRelationsReq) (out *pb.RelationsResponse, err error) {
	out = &pb.RelationsResponse{}

	if in.Key != "thisisasupersecretkeythatyouwillneverguesshahahahahagoodluckidiothackers" {
		err = status.Error(codes.PermissionDenied, "Permission denied")
		return
	}

	var permissionFilter string
	if in.Permission != "" {
		permissionFilter = in.Permission
	}
	out.Relations, err = ListEntityPermissions("nspc", in.Entity, &permissionFilter)
	if err != nil {
		logger.Error("Error listing entity permissions")
		logger.Error(err.Error())
		err = status.Errorf(codes.Internal, "Internal error")
	}

	return
}

func (server) ListObjectRelations(ctx context.Context, in *pb.ListObjectRelationsReq) (out *pb.RelationsResponse, err error) {
	out = &pb.RelationsResponse{}

	if in.Key != "thisisasupersecretkeythatyouwillneverguesshahahahahagoodluckidiothackers" {
		err = status.Error(codes.PermissionDenied, "Permission denied")
		return
	}

	var permissionFilter string
	if in.Permission != "" {
		permissionFilter = in.Permission
	}
	out.Relations, err = ListEntityPermissions("nspc", in.Object, &permissionFilter)
	if err != nil {
		logger.Error("Error listing object permissions")
		logger.Error(err.Error())
		err = status.Errorf(codes.Internal, "Internal error")
	}

	return
}
