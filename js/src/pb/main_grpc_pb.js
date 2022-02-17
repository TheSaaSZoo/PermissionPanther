// GENERATED CODE -- DO NOT EDIT!

'use strict';
var grpc = require('@grpc/grpc-js');
var pb_main_pb = require('../pb/main_pb.js');
var pb_permissions_pb = require('../pb/permissions_pb.js');

function serialize_Applied(arg) {
  if (!(arg instanceof pb_main_pb.Applied)) {
    throw new Error('Expected argument of type Applied');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_Applied(buffer_arg) {
  return pb_main_pb.Applied.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_CheckDirectReq(arg) {
  if (!(arg instanceof pb_permissions_pb.CheckDirectReq)) {
    throw new Error('Expected argument of type CheckDirectReq');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_CheckDirectReq(buffer_arg) {
  return pb_permissions_pb.CheckDirectReq.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_CheckDirectRes(arg) {
  if (!(arg instanceof pb_permissions_pb.CheckDirectRes)) {
    throw new Error('Expected argument of type CheckDirectRes');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_CheckDirectRes(buffer_arg) {
  return pb_permissions_pb.CheckDirectRes.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_CreatePermissionGroupReq(arg) {
  if (!(arg instanceof pb_permissions_pb.CreatePermissionGroupReq)) {
    throw new Error('Expected argument of type CreatePermissionGroupReq');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_CreatePermissionGroupReq(buffer_arg) {
  return pb_permissions_pb.CreatePermissionGroupReq.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_DeletePermissionGroupReq(arg) {
  if (!(arg instanceof pb_permissions_pb.DeletePermissionGroupReq)) {
    throw new Error('Expected argument of type DeletePermissionGroupReq');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_DeletePermissionGroupReq(buffer_arg) {
  return pb_permissions_pb.DeletePermissionGroupReq.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_ListEntityRelationsReq(arg) {
  if (!(arg instanceof pb_permissions_pb.ListEntityRelationsReq)) {
    throw new Error('Expected argument of type ListEntityRelationsReq');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_ListEntityRelationsReq(buffer_arg) {
  return pb_permissions_pb.ListEntityRelationsReq.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_ListObjectRelationsReq(arg) {
  if (!(arg instanceof pb_permissions_pb.ListObjectRelationsReq)) {
    throw new Error('Expected argument of type ListObjectRelationsReq');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_ListObjectRelationsReq(buffer_arg) {
  return pb_permissions_pb.ListObjectRelationsReq.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_ListPermissionGroupReq(arg) {
  if (!(arg instanceof pb_permissions_pb.ListPermissionGroupReq)) {
    throw new Error('Expected argument of type ListPermissionGroupReq');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_ListPermissionGroupReq(buffer_arg) {
  return pb_permissions_pb.ListPermissionGroupReq.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_ListPermissionGroupRes(arg) {
  if (!(arg instanceof pb_permissions_pb.ListPermissionGroupRes)) {
    throw new Error('Expected argument of type ListPermissionGroupRes');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_ListPermissionGroupRes(buffer_arg) {
  return pb_permissions_pb.ListPermissionGroupRes.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_ModifyPermissionGroupReq(arg) {
  if (!(arg instanceof pb_permissions_pb.ModifyPermissionGroupReq)) {
    throw new Error('Expected argument of type ModifyPermissionGroupReq');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_ModifyPermissionGroupReq(buffer_arg) {
  return pb_permissions_pb.ModifyPermissionGroupReq.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_RelationReq(arg) {
  if (!(arg instanceof pb_permissions_pb.RelationReq)) {
    throw new Error('Expected argument of type RelationReq');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_RelationReq(buffer_arg) {
  return pb_permissions_pb.RelationReq.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_RelationsResponse(arg) {
  if (!(arg instanceof pb_permissions_pb.RelationsResponse)) {
    throw new Error('Expected argument of type RelationsResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_RelationsResponse(buffer_arg) {
  return pb_permissions_pb.RelationsResponse.deserializeBinary(new Uint8Array(buffer_arg));
}


var PermissionPantherService = exports.PermissionPantherService = {
  // Checks whether a permission exists, and at what recursion level. If there is an explicit `deny` permission then group checking will be aborted.
checkDirectPermission: {
    path: '/PermissionPanther/CheckDirectPermission',
    requestStream: false,
    responseStream: false,
    requestType: pb_permissions_pb.CheckDirectReq,
    responseType: pb_permissions_pb.CheckDirectRes,
    requestSerialize: serialize_CheckDirectReq,
    requestDeserialize: deserialize_CheckDirectReq,
    responseSerialize: serialize_CheckDirectRes,
    responseDeserialize: deserialize_CheckDirectRes,
  },
  // Lists all the permissions an entity has, optionally specify permissions to filter on
listEntityRelations: {
    path: '/PermissionPanther/ListEntityRelations',
    requestStream: false,
    responseStream: false,
    requestType: pb_permissions_pb.ListEntityRelationsReq,
    responseType: pb_permissions_pb.RelationsResponse,
    requestSerialize: serialize_ListEntityRelationsReq,
    requestDeserialize: deserialize_ListEntityRelationsReq,
    responseSerialize: serialize_RelationsResponse,
    responseDeserialize: deserialize_RelationsResponse,
  },
  // List all relations for an object, optoinally specify permissions to filter on
listObjectRelations: {
    path: '/PermissionPanther/ListObjectRelations',
    requestStream: false,
    responseStream: false,
    requestType: pb_permissions_pb.ListObjectRelationsReq,
    responseType: pb_permissions_pb.RelationsResponse,
    requestSerialize: serialize_ListObjectRelationsReq,
    requestDeserialize: deserialize_ListObjectRelationsReq,
    responseSerialize: serialize_RelationsResponse,
    responseDeserialize: deserialize_RelationsResponse,
  },
  // Will set a permission for an entity on an object. If the permission already exists it is a no-op.
setPermission: {
    path: '/PermissionPanther/SetPermission',
    requestStream: false,
    responseStream: false,
    requestType: pb_permissions_pb.RelationReq,
    responseType: pb_main_pb.Applied,
    requestSerialize: serialize_RelationReq,
    requestDeserialize: deserialize_RelationReq,
    responseSerialize: serialize_Applied,
    responseDeserialize: deserialize_Applied,
  },
  // Will remove a permission for an entity on an object. If the permission does not exist it is a no-op.
removePermission: {
    path: '/PermissionPanther/RemovePermission',
    requestStream: false,
    responseStream: false,
    requestType: pb_permissions_pb.RelationReq,
    responseType: pb_main_pb.Applied,
    requestSerialize: serialize_RelationReq,
    requestDeserialize: deserialize_RelationReq,
    responseSerialize: serialize_Applied,
    responseDeserialize: deserialize_Applied,
  },
  // Creates a permission group if it does not exist
createPermissionGroup: {
    path: '/PermissionPanther/CreatePermissionGroup',
    requestStream: false,
    responseStream: false,
    requestType: pb_permissions_pb.CreatePermissionGroupReq,
    responseType: pb_main_pb.Applied,
    requestSerialize: serialize_CreatePermissionGroupReq,
    requestDeserialize: deserialize_CreatePermissionGroupReq,
    responseSerialize: serialize_Applied,
    responseDeserialize: deserialize_Applied,
  },
  // Deletes a permission group if it exists
deletePermissionGroup: {
    path: '/PermissionPanther/DeletePermissionGroup',
    requestStream: false,
    responseStream: false,
    requestType: pb_permissions_pb.DeletePermissionGroupReq,
    responseType: pb_main_pb.Applied,
    requestSerialize: serialize_DeletePermissionGroupReq,
    requestDeserialize: deserialize_DeletePermissionGroupReq,
    responseSerialize: serialize_Applied,
    responseDeserialize: deserialize_Applied,
  },
  // Adds one or more permissions a to group if it exists, and the permissions are not already in the group
addPermissionToGroup: {
    path: '/PermissionPanther/AddPermissionToGroup',
    requestStream: false,
    responseStream: false,
    requestType: pb_permissions_pb.ModifyPermissionGroupReq,
    responseType: pb_main_pb.Applied,
    requestSerialize: serialize_ModifyPermissionGroupReq,
    requestDeserialize: deserialize_ModifyPermissionGroupReq,
    responseSerialize: serialize_Applied,
    responseDeserialize: deserialize_Applied,
  },
  // Removes one or more permissions to a group if it exists, and the permissions are in the group
removePermissionFromGroup: {
    path: '/PermissionPanther/RemovePermissionFromGroup',
    requestStream: false,
    responseStream: false,
    requestType: pb_permissions_pb.ModifyPermissionGroupReq,
    responseType: pb_main_pb.Applied,
    requestSerialize: serialize_ModifyPermissionGroupReq,
    requestDeserialize: deserialize_ModifyPermissionGroupReq,
    responseSerialize: serialize_Applied,
    responseDeserialize: deserialize_Applied,
  },
  listEntitiesInGroup: {
    path: '/PermissionPanther/ListEntitiesInGroup',
    requestStream: false,
    responseStream: false,
    requestType: pb_permissions_pb.ListPermissionGroupReq,
    responseType: pb_permissions_pb.ListPermissionGroupRes,
    requestSerialize: serialize_ListPermissionGroupReq,
    requestDeserialize: deserialize_ListPermissionGroupReq,
    responseSerialize: serialize_ListPermissionGroupRes,
    responseDeserialize: deserialize_ListPermissionGroupRes,
  },
};

exports.PermissionPantherClient = grpc.makeGenericClientConstructor(PermissionPantherService);
