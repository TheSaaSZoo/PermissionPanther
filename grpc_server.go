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
	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	var opts []grpc.ServerOption

	GRPCServer = grpc.NewServer(opts...)
	pb.RegisterPermissionPantherServer(GRPCServer, &server{})
	logger.Info("Starting Permission Panther gRPC Server on port %s", port)
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
	var found bool
	if in.DenyPermission != "" {
		found, err = CheckPermissionDirect("nspc", in.Object, in.DenyPermission, in.Entity)
		if err != nil {
			logger.Error("Error checking permission direct for deny check")
			logger.Error(err.Error())
			err = status.Errorf(codes.Internal, "Internal error")
			return
		}
		if found {
			out.Valid = false
			out.Recursion = 0
			return
		}
	}

	var foundAt int
	if in.Recursive {
		foundAt, err = CheckPermissions("nspc", in.Object, in.Permission, in.Entity, 0, MAX_RECURSIONS)
		if err != nil {
			logger.Error("Error checking permissions for recursive")
			logger.Error(err.Error())
			err = status.Errorf(codes.Internal, "Internal error")
			return
		}
		if foundAt >= 0 {
			out.Valid = true
			out.Recursion = int32(foundAt)
		} else {
			out.Valid = false
			out.Recursion = -1
		}
	} else {
		found, err = CheckPermissionDirect("nspc", in.Object, in.Permission, in.Entity)
		if err != nil {
			logger.Error("Error checking permissions direct for non recursive")
			logger.Error(err.Error())
			err = status.Errorf(codes.Internal, "Internal error")
			return
		}
		out.Valid = found
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

	out.Relations, err = ListEntityPermissions("nspc", in.Entity, in.Permission)
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

	out.Relations, err = ListObjectPermissions("nspc", in.Object, in.Permission)
	if err != nil {
		logger.Error("Error listing object permissions")
		logger.Error(err.Error())
		err = status.Errorf(codes.Internal, "Internal error")
	}

	return
}

func (server) SetPermission(ctx context.Context, in *pb.RelationReq) (out *pb.NoContent, err error) {
	out = &pb.NoContent{}

	if in.Key != "thisisasupersecretkeythatyouwillneverguesshahahahahagoodluckidiothackers" {
		err = status.Error(codes.PermissionDenied, "Permission denied")
		return
	}

	err = UpsertRelation("nspc", in.Object, in.Permission, in.Entity)
	if err != nil {
		logger.Error("Error upserting relation")
		logger.Error(err.Error())
		err = status.Errorf(codes.Internal, "Internal error")
	}

	return
}

func (server) RemovePermission(ctx context.Context, in *pb.RelationReq) (out *pb.NoContent, err error) {
	out = &pb.NoContent{}

	if in.Key != "thisisasupersecretkeythatyouwillneverguesshahahahahagoodluckidiothackers" {
		err = status.Error(codes.PermissionDenied, "Permission denied")
		return
	}

	err = DeleteRelation("nspc", in.Object, in.Permission, in.Entity)
	if err != nil {
		logger.Error("Error deleting relation")
		logger.Error(err.Error())
		err = status.Errorf(codes.Internal, "Internal error")
	}

	return
}
