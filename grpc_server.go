package main

import (
	"context"
	"net"

	"github.com/danthegoodman1/PermissionPanther/logger"
	"github.com/danthegoodman1/PermissionPanther/pb"
	"github.com/jackc/pgx/v4"
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

func StartGRPCServer(lis net.Listener) {
	var opts []grpc.ServerOption

	GRPCServer = grpc.NewServer(opts...)
	pb.RegisterPermissionPantherServer(GRPCServer, &server{})
	logger.Info("Starting Permission Panther gRPC Server")
	err := GRPCServer.Serve(lis)
	if err != nil {
		logger.Error("Error closing grpc server:")
		logger.Error(err.Error())
	}
}

func (server) CheckDirectPermission(ctx context.Context, in *pb.CheckDirectReq) (out *pb.CheckDirectRes, err error) {
	out = &pb.CheckDirectRes{}

	apiKey, err := CheckAPIKey(in.KeyID, in.KeySecret)
	if err != nil {
		switch err {
		case pgx.ErrNoRows:
			err = status.Error(codes.PermissionDenied, codes.PermissionDenied.String())
		case ErrInvalidHash:
			err = status.Error(codes.PermissionDenied, codes.PermissionDenied.String())
		default:
			logger.Error("CheckDirectPermission(): Check api key error")
			logger.Error(err.Error())
			err = status.Error(codes.Internal, codes.Internal.String())
		}
		return
	}

	// First check explicit deny if provided
	var found bool
	if in.DenyPermission != "" {
		found, err = CheckPermissionDirect(apiKey.Namespace, in.Object, in.DenyPermission, in.Entity)
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
		foundAt, err = CheckPermissions(apiKey.Namespace, in.Object, in.Permission, in.Entity, 0, apiKey.MaxRecursions)
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
		found, err = CheckPermissionDirect(apiKey.Namespace, in.Object, in.Permission, in.Entity)
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

	apiKey, err := CheckAPIKey(in.KeyID, in.KeySecret)
	if err != nil {
		switch err {
		case pgx.ErrNoRows:
			err = status.Error(codes.PermissionDenied, codes.PermissionDenied.String())
		case ErrInvalidHash:
			err = status.Error(codes.PermissionDenied, codes.PermissionDenied.String())
		default:
			logger.Error("CheckDirectPermission(): Check api key error")
			logger.Error(err.Error())
			err = status.Error(codes.Internal, codes.Internal.String())
		}
		return
	}

	out.Relations, err = ListEntityPermissions(apiKey.Namespace, in.Entity, in.Permission)
	if err != nil {
		logger.Error("Error listing entity permissions")
		logger.Error(err.Error())
		err = status.Errorf(codes.Internal, "Internal error")
	}

	return
}

func (server) ListObjectRelations(ctx context.Context, in *pb.ListObjectRelationsReq) (out *pb.RelationsResponse, err error) {
	out = &pb.RelationsResponse{}

	apiKey, err := CheckAPIKey(in.KeyID, in.KeySecret)
	if err != nil {
		switch err {
		case pgx.ErrNoRows:
			err = status.Error(codes.PermissionDenied, codes.PermissionDenied.String())
		case ErrInvalidHash:
			err = status.Error(codes.PermissionDenied, codes.PermissionDenied.String())
		default:
			logger.Error("CheckDirectPermission(): Check api key error")
			logger.Error(err.Error())
			err = status.Error(codes.Internal, codes.Internal.String())
		}
		return
	}

	out.Relations, err = ListObjectPermissions(apiKey.Namespace, in.Object, in.Permission)
	if err != nil {
		logger.Error("Error listing object permissions")
		logger.Error(err.Error())
		err = status.Errorf(codes.Internal, "Internal error")
	}

	return
}

func (server) SetPermission(ctx context.Context, in *pb.RelationReq) (out *pb.Applied, err error) {
	out = &pb.Applied{}

	apiKey, err := CheckAPIKey(in.KeyID, in.KeySecret)
	if err != nil {
		switch err {
		case pgx.ErrNoRows:
			err = status.Error(codes.PermissionDenied, codes.PermissionDenied.String())
		case ErrInvalidHash:
			err = status.Error(codes.PermissionDenied, codes.PermissionDenied.String())
		default:
			logger.Error("CheckDirectPermission(): Check api key error")
			logger.Error(err.Error())
			err = status.Error(codes.Internal, codes.Internal.String())
		}
		return
	}

	// TODO: Check if this is a permission or a group

	out.Applied, err = UpsertRelation(apiKey.Namespace, in.Object, in.Permission, in.Entity)
	if err != nil {
		logger.Error("Error upserting relation")
		logger.Error(err.Error())
		err = status.Errorf(codes.Internal, "Internal error")
	}

	return
}

func (server) RemovePermission(ctx context.Context, in *pb.RelationReq) (out *pb.Applied, err error) {
	out = &pb.Applied{}

	apiKey, err := CheckAPIKey(in.KeyID, in.KeySecret)
	if err != nil {
		switch err {
		case pgx.ErrNoRows:
			err = status.Error(codes.PermissionDenied, codes.PermissionDenied.String())
		case ErrInvalidHash:
			err = status.Error(codes.PermissionDenied, codes.PermissionDenied.String())
		default:
			logger.Error("CheckDirectPermission(): Check api key error")
			logger.Error(err.Error())
			err = status.Error(codes.Internal, codes.Internal.String())
		}
		return
	}

	// TODO: Check if this is a permission or a group

	out.Applied, err = DeleteRelation(apiKey.Namespace, in.Object, in.Permission, in.Entity)
	if err != nil {
		logger.Error("Error deleting relation")
		logger.Error(err.Error())
		err = status.Errorf(codes.Internal, "Internal error")
	}

	return
}

// func (server) CreatePermissionGroup(ctx context.Context, in *pb.PermissionGroupReq) (out *pb.Applied, err error) {
// 	out = &pb.Applied{}
// }

// func (server) DeletePermissionGroup(ctx context.Context, in *pb.PermissionGroupReq) (out *pb.Applied, err error) {
// 	out = &pb.Applied{}
// }

// func (server) AddPermissionToGroup(ctx context.Context, in *pb.ModifyPermissionGroupReq) (out *pb.Applied, err error) {
// 	out = &pb.Applied{}
// }

// func (server) RemovePermissionFromGroup(ctx context.Context, in *pb.ModifyPermissionGroupReq) (out *pb.Applied, err error) {
// 	out = &pb.Applied{}
// }
