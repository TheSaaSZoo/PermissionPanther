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
        const responseType_3: typeof pb_main_pb.Applied;
        export { responseType_3 as responseType };
        export { serialize_RelationReq as requestSerialize };
        export { deserialize_RelationReq as requestDeserialize };
        export { serialize_Applied as responseSerialize };
        export { deserialize_Applied as responseDeserialize };
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
        const responseType_4: typeof pb_main_pb.Applied;
        export { responseType_4 as responseType };
        export { serialize_RelationReq as requestSerialize };
        export { deserialize_RelationReq as requestDeserialize };
        export { serialize_Applied as responseSerialize };
        export { deserialize_Applied as responseDeserialize };
    }
    namespace createPermissionGroup {
        const path_5: string;
        export { path_5 as path };
        const requestStream_5: boolean;
        export { requestStream_5 as requestStream };
        const responseStream_5: boolean;
        export { responseStream_5 as responseStream };
        const requestType_5: typeof pb_permissions_pb.CreatePermissionGroupReq;
        export { requestType_5 as requestType };
        const responseType_5: typeof pb_main_pb.Applied;
        export { responseType_5 as responseType };
        export { serialize_CreatePermissionGroupReq as requestSerialize };
        export { deserialize_CreatePermissionGroupReq as requestDeserialize };
        export { serialize_Applied as responseSerialize };
        export { deserialize_Applied as responseDeserialize };
    }
    namespace deletePermissionGroup {
        const path_6: string;
        export { path_6 as path };
        const requestStream_6: boolean;
        export { requestStream_6 as requestStream };
        const responseStream_6: boolean;
        export { responseStream_6 as responseStream };
        const requestType_6: typeof pb_permissions_pb.DeletePermissionGroupReq;
        export { requestType_6 as requestType };
        const responseType_6: typeof pb_main_pb.Applied;
        export { responseType_6 as responseType };
        export { serialize_DeletePermissionGroupReq as requestSerialize };
        export { deserialize_DeletePermissionGroupReq as requestDeserialize };
        export { serialize_Applied as responseSerialize };
        export { deserialize_Applied as responseDeserialize };
    }
    namespace addPermissionToGroup {
        const path_7: string;
        export { path_7 as path };
        const requestStream_7: boolean;
        export { requestStream_7 as requestStream };
        const responseStream_7: boolean;
        export { responseStream_7 as responseStream };
        const requestType_7: typeof pb_permissions_pb.ModifyPermissionGroupReq;
        export { requestType_7 as requestType };
        const responseType_7: typeof pb_main_pb.Applied;
        export { responseType_7 as responseType };
        export { serialize_ModifyPermissionGroupReq as requestSerialize };
        export { deserialize_ModifyPermissionGroupReq as requestDeserialize };
        export { serialize_Applied as responseSerialize };
        export { deserialize_Applied as responseDeserialize };
    }
    namespace removePermissionFromGroup {
        const path_8: string;
        export { path_8 as path };
        const requestStream_8: boolean;
        export { requestStream_8 as requestStream };
        const responseStream_8: boolean;
        export { responseStream_8 as responseStream };
        const requestType_8: typeof pb_permissions_pb.ModifyPermissionGroupReq;
        export { requestType_8 as requestType };
        const responseType_8: typeof pb_main_pb.Applied;
        export { responseType_8 as responseType };
        export { serialize_ModifyPermissionGroupReq as requestSerialize };
        export { deserialize_ModifyPermissionGroupReq as requestDeserialize };
        export { serialize_Applied as responseSerialize };
        export { deserialize_Applied as responseDeserialize };
    }
    namespace listEntitiesInGroup {
        const path_9: string;
        export { path_9 as path };
        const requestStream_9: boolean;
        export { requestStream_9 as requestStream };
        const responseStream_9: boolean;
        export { responseStream_9 as responseStream };
        const requestType_9: typeof pb_permissions_pb.ListPermissionGroupReq;
        export { requestType_9 as requestType };
        const responseType_9: typeof pb_permissions_pb.ListPermissionGroupRes;
        export { responseType_9 as responseType };
        export { serialize_ListPermissionGroupReq as requestSerialize };
        export { deserialize_ListPermissionGroupReq as requestDeserialize };
        export { serialize_ListPermissionGroupRes as responseSerialize };
        export { deserialize_ListPermissionGroupRes as responseDeserialize };
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
declare function serialize_Applied(arg: any): Buffer;
declare function deserialize_Applied(buffer_arg: any): pb_main_pb.Applied;
declare function serialize_CreatePermissionGroupReq(arg: any): Buffer;
declare function deserialize_CreatePermissionGroupReq(buffer_arg: any): pb_permissions_pb.CreatePermissionGroupReq;
declare function serialize_DeletePermissionGroupReq(arg: any): Buffer;
declare function deserialize_DeletePermissionGroupReq(buffer_arg: any): pb_permissions_pb.DeletePermissionGroupReq;
declare function serialize_ModifyPermissionGroupReq(arg: any): Buffer;
declare function deserialize_ModifyPermissionGroupReq(buffer_arg: any): pb_permissions_pb.ModifyPermissionGroupReq;
declare function serialize_ListPermissionGroupReq(arg: any): Buffer;
declare function deserialize_ListPermissionGroupReq(buffer_arg: any): pb_permissions_pb.ListPermissionGroupReq;
declare function serialize_ListPermissionGroupRes(arg: any): Buffer;
declare function deserialize_ListPermissionGroupRes(buffer_arg: any): pb_permissions_pb.ListPermissionGroupRes;
import grpc = require("@grpc/grpc-js");
export {};
