// package: 
// file: pb/main.proto

/* tslint:disable */
/* eslint-disable */

import * as grpc from "grpc";
import * as pb_main_pb from "../pb/main_pb";
import * as pb_permissions_pb from "../pb/permissions_pb";

interface IPermissionPantherService extends grpc.ServiceDefinition<grpc.UntypedServiceImplementation> {
    checkDirectPermission: IPermissionPantherService_ICheckDirectPermission;
    listEntityRelations: IPermissionPantherService_IListEntityRelations;
    listObjectRelations: IPermissionPantherService_IListObjectRelations;
    setPermission: IPermissionPantherService_ISetPermission;
    removePermission: IPermissionPantherService_IRemovePermission;
}

interface IPermissionPantherService_ICheckDirectPermission extends grpc.MethodDefinition<pb_permissions_pb.CheckDirectReq, pb_permissions_pb.CheckDirectRes> {
    path: "/PermissionPanther/CheckDirectPermission";
    requestStream: false;
    responseStream: false;
    requestSerialize: grpc.serialize<pb_permissions_pb.CheckDirectReq>;
    requestDeserialize: grpc.deserialize<pb_permissions_pb.CheckDirectReq>;
    responseSerialize: grpc.serialize<pb_permissions_pb.CheckDirectRes>;
    responseDeserialize: grpc.deserialize<pb_permissions_pb.CheckDirectRes>;
}
interface IPermissionPantherService_IListEntityRelations extends grpc.MethodDefinition<pb_permissions_pb.ListEntityRelationsReq, pb_permissions_pb.RelationsResponse> {
    path: "/PermissionPanther/ListEntityRelations";
    requestStream: false;
    responseStream: false;
    requestSerialize: grpc.serialize<pb_permissions_pb.ListEntityRelationsReq>;
    requestDeserialize: grpc.deserialize<pb_permissions_pb.ListEntityRelationsReq>;
    responseSerialize: grpc.serialize<pb_permissions_pb.RelationsResponse>;
    responseDeserialize: grpc.deserialize<pb_permissions_pb.RelationsResponse>;
}
interface IPermissionPantherService_IListObjectRelations extends grpc.MethodDefinition<pb_permissions_pb.ListObjectRelationsReq, pb_permissions_pb.RelationsResponse> {
    path: "/PermissionPanther/ListObjectRelations";
    requestStream: false;
    responseStream: false;
    requestSerialize: grpc.serialize<pb_permissions_pb.ListObjectRelationsReq>;
    requestDeserialize: grpc.deserialize<pb_permissions_pb.ListObjectRelationsReq>;
    responseSerialize: grpc.serialize<pb_permissions_pb.RelationsResponse>;
    responseDeserialize: grpc.deserialize<pb_permissions_pb.RelationsResponse>;
}
interface IPermissionPantherService_ISetPermission extends grpc.MethodDefinition<pb_permissions_pb.RelationReq, pb_main_pb.Applied> {
    path: "/PermissionPanther/SetPermission";
    requestStream: false;
    responseStream: false;
    requestSerialize: grpc.serialize<pb_permissions_pb.RelationReq>;
    requestDeserialize: grpc.deserialize<pb_permissions_pb.RelationReq>;
    responseSerialize: grpc.serialize<pb_main_pb.Applied>;
    responseDeserialize: grpc.deserialize<pb_main_pb.Applied>;
}
interface IPermissionPantherService_IRemovePermission extends grpc.MethodDefinition<pb_permissions_pb.RelationReq, pb_main_pb.Applied> {
    path: "/PermissionPanther/RemovePermission";
    requestStream: false;
    responseStream: false;
    requestSerialize: grpc.serialize<pb_permissions_pb.RelationReq>;
    requestDeserialize: grpc.deserialize<pb_permissions_pb.RelationReq>;
    responseSerialize: grpc.serialize<pb_main_pb.Applied>;
    responseDeserialize: grpc.deserialize<pb_main_pb.Applied>;
}

export const PermissionPantherService: IPermissionPantherService;

export interface IPermissionPantherServer {
    checkDirectPermission: grpc.handleUnaryCall<pb_permissions_pb.CheckDirectReq, pb_permissions_pb.CheckDirectRes>;
    listEntityRelations: grpc.handleUnaryCall<pb_permissions_pb.ListEntityRelationsReq, pb_permissions_pb.RelationsResponse>;
    listObjectRelations: grpc.handleUnaryCall<pb_permissions_pb.ListObjectRelationsReq, pb_permissions_pb.RelationsResponse>;
    setPermission: grpc.handleUnaryCall<pb_permissions_pb.RelationReq, pb_main_pb.Applied>;
    removePermission: grpc.handleUnaryCall<pb_permissions_pb.RelationReq, pb_main_pb.Applied>;
}

export interface IPermissionPantherClient {
    checkDirectPermission(request: pb_permissions_pb.CheckDirectReq, callback: (error: grpc.ServiceError | null, response: pb_permissions_pb.CheckDirectRes) => void): grpc.ClientUnaryCall;
    checkDirectPermission(request: pb_permissions_pb.CheckDirectReq, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: pb_permissions_pb.CheckDirectRes) => void): grpc.ClientUnaryCall;
    checkDirectPermission(request: pb_permissions_pb.CheckDirectReq, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: pb_permissions_pb.CheckDirectRes) => void): grpc.ClientUnaryCall;
    listEntityRelations(request: pb_permissions_pb.ListEntityRelationsReq, callback: (error: grpc.ServiceError | null, response: pb_permissions_pb.RelationsResponse) => void): grpc.ClientUnaryCall;
    listEntityRelations(request: pb_permissions_pb.ListEntityRelationsReq, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: pb_permissions_pb.RelationsResponse) => void): grpc.ClientUnaryCall;
    listEntityRelations(request: pb_permissions_pb.ListEntityRelationsReq, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: pb_permissions_pb.RelationsResponse) => void): grpc.ClientUnaryCall;
    listObjectRelations(request: pb_permissions_pb.ListObjectRelationsReq, callback: (error: grpc.ServiceError | null, response: pb_permissions_pb.RelationsResponse) => void): grpc.ClientUnaryCall;
    listObjectRelations(request: pb_permissions_pb.ListObjectRelationsReq, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: pb_permissions_pb.RelationsResponse) => void): grpc.ClientUnaryCall;
    listObjectRelations(request: pb_permissions_pb.ListObjectRelationsReq, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: pb_permissions_pb.RelationsResponse) => void): grpc.ClientUnaryCall;
    setPermission(request: pb_permissions_pb.RelationReq, callback: (error: grpc.ServiceError | null, response: pb_main_pb.Applied) => void): grpc.ClientUnaryCall;
    setPermission(request: pb_permissions_pb.RelationReq, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: pb_main_pb.Applied) => void): grpc.ClientUnaryCall;
    setPermission(request: pb_permissions_pb.RelationReq, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: pb_main_pb.Applied) => void): grpc.ClientUnaryCall;
    removePermission(request: pb_permissions_pb.RelationReq, callback: (error: grpc.ServiceError | null, response: pb_main_pb.Applied) => void): grpc.ClientUnaryCall;
    removePermission(request: pb_permissions_pb.RelationReq, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: pb_main_pb.Applied) => void): grpc.ClientUnaryCall;
    removePermission(request: pb_permissions_pb.RelationReq, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: pb_main_pb.Applied) => void): grpc.ClientUnaryCall;
}

export class PermissionPantherClient extends grpc.Client implements IPermissionPantherClient {
    constructor(address: string, credentials: grpc.ChannelCredentials, options?: object);
    public checkDirectPermission(request: pb_permissions_pb.CheckDirectReq, callback: (error: grpc.ServiceError | null, response: pb_permissions_pb.CheckDirectRes) => void): grpc.ClientUnaryCall;
    public checkDirectPermission(request: pb_permissions_pb.CheckDirectReq, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: pb_permissions_pb.CheckDirectRes) => void): grpc.ClientUnaryCall;
    public checkDirectPermission(request: pb_permissions_pb.CheckDirectReq, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: pb_permissions_pb.CheckDirectRes) => void): grpc.ClientUnaryCall;
    public listEntityRelations(request: pb_permissions_pb.ListEntityRelationsReq, callback: (error: grpc.ServiceError | null, response: pb_permissions_pb.RelationsResponse) => void): grpc.ClientUnaryCall;
    public listEntityRelations(request: pb_permissions_pb.ListEntityRelationsReq, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: pb_permissions_pb.RelationsResponse) => void): grpc.ClientUnaryCall;
    public listEntityRelations(request: pb_permissions_pb.ListEntityRelationsReq, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: pb_permissions_pb.RelationsResponse) => void): grpc.ClientUnaryCall;
    public listObjectRelations(request: pb_permissions_pb.ListObjectRelationsReq, callback: (error: grpc.ServiceError | null, response: pb_permissions_pb.RelationsResponse) => void): grpc.ClientUnaryCall;
    public listObjectRelations(request: pb_permissions_pb.ListObjectRelationsReq, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: pb_permissions_pb.RelationsResponse) => void): grpc.ClientUnaryCall;
    public listObjectRelations(request: pb_permissions_pb.ListObjectRelationsReq, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: pb_permissions_pb.RelationsResponse) => void): grpc.ClientUnaryCall;
    public setPermission(request: pb_permissions_pb.RelationReq, callback: (error: grpc.ServiceError | null, response: pb_main_pb.Applied) => void): grpc.ClientUnaryCall;
    public setPermission(request: pb_permissions_pb.RelationReq, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: pb_main_pb.Applied) => void): grpc.ClientUnaryCall;
    public setPermission(request: pb_permissions_pb.RelationReq, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: pb_main_pb.Applied) => void): grpc.ClientUnaryCall;
    public removePermission(request: pb_permissions_pb.RelationReq, callback: (error: grpc.ServiceError | null, response: pb_main_pb.Applied) => void): grpc.ClientUnaryCall;
    public removePermission(request: pb_permissions_pb.RelationReq, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: pb_main_pb.Applied) => void): grpc.ClientUnaryCall;
    public removePermission(request: pb_permissions_pb.RelationReq, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: pb_main_pb.Applied) => void): grpc.ClientUnaryCall;
}
