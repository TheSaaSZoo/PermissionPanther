export namespace PermissionPantherService {
    namespace checkDirectPermission {
        export const path: string;
        export const requestStream: boolean;
        export const responseStream: boolean;
        export const requestType: typeof pb_permissions_pb.CheckDirectReq;
        export const responseType: typeof pb_permissions_pb.CheckDirectRes;
        export { serialize_CheckDirectReq as requestSerialize };
        export { deserialize_CheckDirectReq as requestDeserialize };
        export { serialize_CheckDirectRes as responseSerialize };
        export { deserialize_CheckDirectRes as responseDeserialize };
    }
    namespace listEntityRelations {
        const path_1: string;
        export { path_1 as path };
        const requestStream_1: boolean;
        export { requestStream_1 as requestStream };
        const responseStream_1: boolean;
        export { responseStream_1 as responseStream };
        const requestType_1: typeof pb_permissions_pb.ListEntityRelationsReq;
        export { requestType_1 as requestType };
        const responseType_1: typeof pb_permissions_pb.RelationsResponse;
        export { responseType_1 as responseType };
        export { serialize_ListEntityRelationsReq as requestSerialize };
        export { deserialize_ListEntityRelationsReq as requestDeserialize };
        export { serialize_RelationsResponse as responseSerialize };
        export { deserialize_RelationsResponse as responseDeserialize };
    }
    namespace listObjectRelations {
        const path_2: string;
        export { path_2 as path };
        const requestStream_2: boolean;
        export { requestStream_2 as requestStream };
        const responseStream_2: boolean;
        export { responseStream_2 as responseStream };
        const requestType_2: typeof pb_permissions_pb.ListObjectRelationsReq;
        export { requestType_2 as requestType };
        const responseType_2: typeof pb_permissions_pb.RelationsResponse;
        export { responseType_2 as responseType };
        export { serialize_ListObjectRelationsReq as requestSerialize };
        export { deserialize_ListObjectRelationsReq as requestDeserialize };
        export { serialize_RelationsResponse as responseSerialize };
        export { deserialize_RelationsResponse as responseDeserialize };
    }
    namespace setPermission {
        const path_3: string;
        export { path_3 as path };
        const requestStream_3: boolean;
        export { requestStream_3 as requestStream };
        const responseStream_3: boolean;
        export { responseStream_3 as responseStream };
        const requestType_3: typeof pb_permissions_pb.RelationReq;
        export { requestType_3 as requestType };
        const responseType_3: typeof pb_main_pb.NoContent;
        export { responseType_3 as responseType };
        export { serialize_RelationReq as requestSerialize };
        export { deserialize_RelationReq as requestDeserialize };
        export { serialize_NoContent as responseSerialize };
        export { deserialize_NoContent as responseDeserialize };
    }
    namespace removePermission {
        const path_4: string;
        export { path_4 as path };
        const requestStream_4: boolean;
        export { requestStream_4 as requestStream };
        const responseStream_4: boolean;
        export { responseStream_4 as responseStream };
        const requestType_4: typeof pb_permissions_pb.RelationReq;
        export { requestType_4 as requestType };
        const responseType_4: typeof pb_main_pb.NoContent;
        export { responseType_4 as responseType };
        export { serialize_RelationReq as requestSerialize };
        export { deserialize_RelationReq as requestDeserialize };
        export { serialize_NoContent as responseSerialize };
        export { deserialize_NoContent as responseDeserialize };
    }
}
export var PermissionPantherClient: grpc.ServiceClientConstructor;
import pb_permissions_pb = require("../pb/permissions_pb.js");
declare function serialize_CheckDirectReq(arg: any): Buffer;
declare function deserialize_CheckDirectReq(buffer_arg: any): pb_permissions_pb.CheckDirectReq;
declare function serialize_CheckDirectRes(arg: any): Buffer;
declare function deserialize_CheckDirectRes(buffer_arg: any): pb_permissions_pb.CheckDirectRes;
declare function serialize_ListEntityRelationsReq(arg: any): Buffer;
declare function deserialize_ListEntityRelationsReq(buffer_arg: any): pb_permissions_pb.ListEntityRelationsReq;
declare function serialize_RelationsResponse(arg: any): Buffer;
declare function deserialize_RelationsResponse(buffer_arg: any): pb_permissions_pb.RelationsResponse;
declare function serialize_ListObjectRelationsReq(arg: any): Buffer;
declare function deserialize_ListObjectRelationsReq(buffer_arg: any): pb_permissions_pb.ListObjectRelationsReq;
import pb_main_pb = require("../pb/main_pb.js");
declare function serialize_RelationReq(arg: any): Buffer;
declare function deserialize_RelationReq(buffer_arg: any): pb_permissions_pb.RelationReq;
declare function serialize_NoContent(arg: any): Buffer;
declare function deserialize_NoContent(buffer_arg: any): pb_main_pb.NoContent;
import grpc = require("@grpc/grpc-js");
export {};
