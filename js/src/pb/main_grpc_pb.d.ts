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
    createPermissionGroup: IPermissionPantherService_ICreatePermissionGroup;
    deletePermissionGroup: IPermissionPantherService_IDeletePermissionGroup;
    addPermissionToGroup: IPermissionPantherService_IAddPermissionToGroup;
    removePermissionFromGroup: IPermissionPantherService_IRemovePermissionFromGroup;
    listEntitiesInGroup: IPermissionPantherService_IListEntitiesInGroup;
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
interface IPermissionPantherService_ICreatePermissionGroup extends grpc.MethodDefinition<pb_permissions_pb.CreatePermissionGroupReq, pb_main_pb.Applied> {
    path: "/PermissionPanther/CreatePermissionGroup";
    requestStream: false;
    responseStream: false;
    requestSerialize: grpc.serialize<pb_permissions_pb.CreatePermissionGroupReq>;
    requestDeserialize: grpc.deserialize<pb_permissions_pb.CreatePermissionGroupReq>;
    responseSerialize: grpc.serialize<pb_main_pb.Applied>;
    responseDeserialize: grpc.deserialize<pb_main_pb.Applied>;
}
interface IPermissionPantherService_IDeletePermissionGroup extends grpc.MethodDefinition<pb_permissions_pb.DeletePermissionGroupReq, pb_main_pb.Applied> {
    path: "/PermissionPanther/DeletePermissionGroup";
    requestStream: false;
    responseStream: false;
    requestSerialize: grpc.serialize<pb_permissions_pb.DeletePermissionGroupReq>;
    requestDeserialize: grpc.deserialize<pb_permissions_pb.DeletePermissionGroupReq>;
    responseSerialize: grpc.serialize<pb_main_pb.Applied>;
    responseDeserialize: grpc.deserialize<pb_main_pb.Applied>;
}
interface IPermissionPantherService_IAddPermissionToGroup extends grpc.MethodDefinition<pb_permissions_pb.ModifyPermissionGroupReq, pb_main_pb.Applied> {
    path: "/PermissionPanther/AddPermissionToGroup";
    requestStream: false;
    responseStream: false;
    requestSerialize: grpc.serialize<pb_permissions_pb.ModifyPermissionGroupReq>;
    requestDeserialize: grpc.deserialize<pb_permissions_pb.ModifyPermissionGroupReq>;
    responseSerialize: grpc.serialize<pb_main_pb.Applied>;
    responseDeserialize: grpc.deserialize<pb_main_pb.Applied>;
}
interface IPermissionPantherService_IRemovePermissionFromGroup extends grpc.MethodDefinition<pb_permissions_pb.ModifyPermissionGroupReq, pb_main_pb.Applied> {
    path: "/PermissionPanther/RemovePermissionFromGroup";
    requestStream: false;
    responseStream: false;
    requestSerialize: grpc.serialize<pb_permissions_pb.ModifyPermissionGroupReq>;
    requestDeserialize: grpc.deserialize<pb_permissions_pb.ModifyPermissionGroupReq>;
    responseSerialize: grpc.serialize<pb_main_pb.Applied>;
    responseDeserialize: grpc.deserialize<pb_main_pb.Applied>;
}
interface IPermissionPantherService_IListEntitiesInGroup extends grpc.MethodDefinition<pb_permissions_pb.ListPermissionGroupReq, pb_permissions_pb.ListPermissionGroupRes> {
    path: "/PermissionPanther/ListEntitiesInGroup";
    requestStream: false;
    responseStream: false;
    requestSerialize: grpc.serialize<pb_permissions_pb.ListPermissionGroupReq>;
    requestDeserialize: grpc.deserialize<pb_permissions_pb.ListPermissionGroupReq>;
    responseSerialize: grpc.serialize<pb_permissions_pb.ListPermissionGroupRes>;
    responseDeserialize: grpc.deserialize<pb_permissions_pb.ListPermissionGroupRes>;
}

export const PermissionPantherService: IPermissionPantherService;

export interface IPermissionPantherServer {
    checkDirectPermission: grpc.handleUnaryCall<pb_permissions_pb.CheckDirectReq, pb_permissions_pb.CheckDirectRes>;
    listEntityRelations: grpc.handleUnaryCall<pb_permissions_pb.ListEntityRelationsReq, pb_permissions_pb.RelationsResponse>;
    listObjectRelations: grpc.handleUnaryCall<pb_permissions_pb.ListObjectRelationsReq, pb_permissions_pb.RelationsResponse>;
    setPermission: grpc.handleUnaryCall<pb_permissions_pb.RelationReq, pb_main_pb.Applied>;
    removePermission: grpc.handleUnaryCall<pb_permissions_pb.RelationReq, pb_main_pb.Applied>;
    createPermissionGroup: grpc.handleUnaryCall<pb_permissions_pb.CreatePermissionGroupReq, pb_main_pb.Applied>;
    deletePermissionGroup: grpc.handleUnaryCall<pb_permissions_pb.DeletePermissionGroupReq, pb_main_pb.Applied>;
    addPermissionToGroup: grpc.handleUnaryCall<pb_permissions_pb.ModifyPermissionGroupReq, pb_main_pb.Applied>;
    removePermissionFromGroup: grpc.handleUnaryCall<pb_permissions_pb.ModifyPermissionGroupReq, pb_main_pb.Applied>;
    listEntitiesInGroup: grpc.handleUnaryCall<pb_permissions_pb.ListPermissionGroupReq, pb_permissions_pb.ListPermissionGroupRes>;
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
    createPermissionGroup(request: pb_permissions_pb.CreatePermissionGroupReq, callback: (error: grpc.ServiceError | null, response: pb_main_pb.Applied) => void): grpc.ClientUnaryCall;
    createPermissionGroup(request: pb_permissions_pb.CreatePermissionGroupReq, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: pb_main_pb.Applied) => void): grpc.ClientUnaryCall;
    createPermissionGroup(request: pb_permissions_pb.CreatePermissionGroupReq, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: pb_main_pb.Applied) => void): grpc.ClientUnaryCall;
    deletePermissionGroup(request: pb_permissions_pb.DeletePermissionGroupReq, callback: (error: grpc.ServiceError | null, response: pb_main_pb.Applied) => void): grpc.ClientUnaryCall;
    deletePermissionGroup(request: pb_permissions_pb.DeletePermissionGroupReq, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: pb_main_pb.Applied) => void): grpc.ClientUnaryCall;
    deletePermissionGroup(request: pb_permissions_pb.DeletePermissionGroupReq, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: pb_main_pb.Applied) => void): grpc.ClientUnaryCall;
    addPermissionToGroup(request: pb_permissions_pb.ModifyPermissionGroupReq, callback: (error: grpc.ServiceError | null, response: pb_main_pb.Applied) => void): grpc.ClientUnaryCall;
    addPermissionToGroup(request: pb_permissions_pb.ModifyPermissionGroupReq, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: pb_main_pb.Applied) => void): grpc.ClientUnaryCall;
    addPermissionToGroup(request: pb_permissions_pb.ModifyPermissionGroupReq, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: pb_main_pb.Applied) => void): grpc.ClientUnaryCall;
    removePermissionFromGroup(request: pb_permissions_pb.ModifyPermissionGroupReq, callback: (error: grpc.ServiceError | null, response: pb_main_pb.Applied) => void): grpc.ClientUnaryCall;
    removePermissionFromGroup(request: pb_permissions_pb.ModifyPermissionGroupReq, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: pb_main_pb.Applied) => void): grpc.ClientUnaryCall;
    removePermissionFromGroup(request: pb_permissions_pb.ModifyPermissionGroupReq, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: pb_main_pb.Applied) => void): grpc.ClientUnaryCall;
    listEntitiesInGroup(request: pb_permissions_pb.ListPermissionGroupReq, callback: (error: grpc.ServiceError | null, response: pb_permissions_pb.ListPermissionGroupRes) => void): grpc.ClientUnaryCall;
    listEntitiesInGroup(request: pb_permissions_pb.ListPermissionGroupReq, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: pb_permissions_pb.ListPermissionGroupRes) => void): grpc.ClientUnaryCall;
    listEntitiesInGroup(request: pb_permissions_pb.ListPermissionGroupReq, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: pb_permissions_pb.ListPermissionGroupRes) => void): grpc.ClientUnaryCall;
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
    public createPermissionGroup(request: pb_permissions_pb.CreatePermissionGroupReq, callback: (error: grpc.ServiceError | null, response: pb_main_pb.Applied) => void): grpc.ClientUnaryCall;
    public createPermissionGroup(request: pb_permissions_pb.CreatePermissionGroupReq, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: pb_main_pb.Applied) => void): grpc.ClientUnaryCall;
    public createPermissionGroup(request: pb_permissions_pb.CreatePermissionGroupReq, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: pb_main_pb.Applied) => void): grpc.ClientUnaryCall;
    public deletePermissionGroup(request: pb_permissions_pb.DeletePermissionGroupReq, callback: (error: grpc.ServiceError | null, response: pb_main_pb.Applied) => void): grpc.ClientUnaryCall;
    public deletePermissionGroup(request: pb_permissions_pb.DeletePermissionGroupReq, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: pb_main_pb.Applied) => void): grpc.ClientUnaryCall;
    public deletePermissionGroup(request: pb_permissions_pb.DeletePermissionGroupReq, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: pb_main_pb.Applied) => void): grpc.ClientUnaryCall;
    public addPermissionToGroup(request: pb_permissions_pb.ModifyPermissionGroupReq, callback: (error: grpc.ServiceError | null, response: pb_main_pb.Applied) => void): grpc.ClientUnaryCall;
    public addPermissionToGroup(request: pb_permissions_pb.ModifyPermissionGroupReq, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: pb_main_pb.Applied) => void): grpc.ClientUnaryCall;
    public addPermissionToGroup(request: pb_permissions_pb.ModifyPermissionGroupReq, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: pb_main_pb.Applied) => void): grpc.ClientUnaryCall;
    public removePermissionFromGroup(request: pb_permissions_pb.ModifyPermissionGroupReq, callback: (error: grpc.ServiceError | null, response: pb_main_pb.Applied) => void): grpc.ClientUnaryCall;
    public removePermissionFromGroup(request: pb_permissions_pb.ModifyPermissionGroupReq, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: pb_main_pb.Applied) => void): grpc.ClientUnaryCall;
    public removePermissionFromGroup(request: pb_permissions_pb.ModifyPermissionGroupReq, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: pb_main_pb.Applied) => void): grpc.ClientUnaryCall;
    public listEntitiesInGroup(request: pb_permissions_pb.ListPermissionGroupReq, callback: (error: grpc.ServiceError | null, response: pb_permissions_pb.ListPermissionGroupRes) => void): grpc.ClientUnaryCall;
    public listEntitiesInGroup(request: pb_permissions_pb.ListPermissionGroupReq, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: pb_permissions_pb.ListPermissionGroupRes) => void): grpc.ClientUnaryCall;
    public listEntitiesInGroup(request: pb_permissions_pb.ListPermissionGroupReq, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: pb_permissions_pb.ListPermissionGroupRes) => void): grpc.ClientUnaryCall;
}
