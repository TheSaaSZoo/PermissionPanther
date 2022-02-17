// package: 
// file: pb/permissions.proto

/* tslint:disable */
/* eslint-disable */

import * as jspb from "google-protobuf";

export class CheckDirectReq extends jspb.Message { 
    getKeyid(): string;
    setKeyid(value: string): CheckDirectReq;
    getKeysecret(): string;
    setKeysecret(value: string): CheckDirectReq;
    getEntity(): string;
    setEntity(value: string): CheckDirectReq;
    getPermission(): string;
    setPermission(value: string): CheckDirectReq;
    getObject(): string;
    setObject(value: string): CheckDirectReq;
    getRecursive(): boolean;
    setRecursive(value: boolean): CheckDirectReq;
    getDenypermission(): string;
    setDenypermission(value: string): CheckDirectReq;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): CheckDirectReq.AsObject;
    static toObject(includeInstance: boolean, msg: CheckDirectReq): CheckDirectReq.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: CheckDirectReq, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): CheckDirectReq;
    static deserializeBinaryFromReader(message: CheckDirectReq, reader: jspb.BinaryReader): CheckDirectReq;
}

export namespace CheckDirectReq {
    export type AsObject = {
        keyid: string,
        keysecret: string,
        entity: string,
        permission: string,
        object: string,
        recursive: boolean,
        denypermission: string,
    }
}

export class CheckDirectRes extends jspb.Message { 
    getValid(): boolean;
    setValid(value: boolean): CheckDirectRes;
    getRecursion(): number;
    setRecursion(value: number): CheckDirectRes;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): CheckDirectRes.AsObject;
    static toObject(includeInstance: boolean, msg: CheckDirectRes): CheckDirectRes.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: CheckDirectRes, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): CheckDirectRes;
    static deserializeBinaryFromReader(message: CheckDirectRes, reader: jspb.BinaryReader): CheckDirectRes;
}

export namespace CheckDirectRes {
    export type AsObject = {
        valid: boolean,
        recursion: number,
    }
}

export class ListEntityRelationsReq extends jspb.Message { 
    getKeyid(): string;
    setKeyid(value: string): ListEntityRelationsReq;
    getKeysecret(): string;
    setKeysecret(value: string): ListEntityRelationsReq;
    getEntity(): string;
    setEntity(value: string): ListEntityRelationsReq;
    getPermission(): string;
    setPermission(value: string): ListEntityRelationsReq;
    getOffset(): number;
    setOffset(value: number): ListEntityRelationsReq;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): ListEntityRelationsReq.AsObject;
    static toObject(includeInstance: boolean, msg: ListEntityRelationsReq): ListEntityRelationsReq.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: ListEntityRelationsReq, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): ListEntityRelationsReq;
    static deserializeBinaryFromReader(message: ListEntityRelationsReq, reader: jspb.BinaryReader): ListEntityRelationsReq;
}

export namespace ListEntityRelationsReq {
    export type AsObject = {
        keyid: string,
        keysecret: string,
        entity: string,
        permission: string,
        offset: number,
    }
}

export class ListObjectRelationsReq extends jspb.Message { 
    getKeyid(): string;
    setKeyid(value: string): ListObjectRelationsReq;
    getKeysecret(): string;
    setKeysecret(value: string): ListObjectRelationsReq;
    getObject(): string;
    setObject(value: string): ListObjectRelationsReq;
    getPermission(): string;
    setPermission(value: string): ListObjectRelationsReq;
    getOffset(): number;
    setOffset(value: number): ListObjectRelationsReq;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): ListObjectRelationsReq.AsObject;
    static toObject(includeInstance: boolean, msg: ListObjectRelationsReq): ListObjectRelationsReq.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: ListObjectRelationsReq, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): ListObjectRelationsReq;
    static deserializeBinaryFromReader(message: ListObjectRelationsReq, reader: jspb.BinaryReader): ListObjectRelationsReq;
}

export namespace ListObjectRelationsReq {
    export type AsObject = {
        keyid: string,
        keysecret: string,
        object: string,
        permission: string,
        offset: number,
    }
}

export class RelationsResponse extends jspb.Message { 
    clearRelationsList(): void;
    getRelationsList(): Array<Relation>;
    setRelationsList(value: Array<Relation>): RelationsResponse;
    addRelations(value?: Relation, index?: number): Relation;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): RelationsResponse.AsObject;
    static toObject(includeInstance: boolean, msg: RelationsResponse): RelationsResponse.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: RelationsResponse, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): RelationsResponse;
    static deserializeBinaryFromReader(message: RelationsResponse, reader: jspb.BinaryReader): RelationsResponse;
}

export namespace RelationsResponse {
    export type AsObject = {
        relationsList: Array<Relation.AsObject>,
    }
}

export class Relation extends jspb.Message { 
    getEntity(): string;
    setEntity(value: string): Relation;
    getPermission(): string;
    setPermission(value: string): Relation;
    getObject(): string;
    setObject(value: string): Relation;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): Relation.AsObject;
    static toObject(includeInstance: boolean, msg: Relation): Relation.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: Relation, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): Relation;
    static deserializeBinaryFromReader(message: Relation, reader: jspb.BinaryReader): Relation;
}

export namespace Relation {
    export type AsObject = {
        entity: string,
        permission: string,
        object: string,
    }
}

export class RelationReq extends jspb.Message { 
    getKeyid(): string;
    setKeyid(value: string): RelationReq;
    getKeysecret(): string;
    setKeysecret(value: string): RelationReq;
    getEntity(): string;
    setEntity(value: string): RelationReq;
    getPermission(): string;
    setPermission(value: string): RelationReq;
    getObject(): string;
    setObject(value: string): RelationReq;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): RelationReq.AsObject;
    static toObject(includeInstance: boolean, msg: RelationReq): RelationReq.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: RelationReq, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): RelationReq;
    static deserializeBinaryFromReader(message: RelationReq, reader: jspb.BinaryReader): RelationReq;
}

export namespace RelationReq {
    export type AsObject = {
        keyid: string,
        keysecret: string,
        entity: string,
        permission: string,
        object: string,
    }
}

export class CreatePermissionGroupReq extends jspb.Message { 
    getKeyid(): string;
    setKeyid(value: string): CreatePermissionGroupReq;
    getKeysecret(): string;
    setKeysecret(value: string): CreatePermissionGroupReq;
    getGroupname(): string;
    setGroupname(value: string): CreatePermissionGroupReq;
    clearPermissionsList(): void;
    getPermissionsList(): Array<string>;
    setPermissionsList(value: Array<string>): CreatePermissionGroupReq;
    addPermissions(value: string, index?: number): string;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): CreatePermissionGroupReq.AsObject;
    static toObject(includeInstance: boolean, msg: CreatePermissionGroupReq): CreatePermissionGroupReq.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: CreatePermissionGroupReq, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): CreatePermissionGroupReq;
    static deserializeBinaryFromReader(message: CreatePermissionGroupReq, reader: jspb.BinaryReader): CreatePermissionGroupReq;
}

export namespace CreatePermissionGroupReq {
    export type AsObject = {
        keyid: string,
        keysecret: string,
        groupname: string,
        permissionsList: Array<string>,
    }
}

export class DeletePermissionGroupReq extends jspb.Message { 
    getKeyid(): string;
    setKeyid(value: string): DeletePermissionGroupReq;
    getKeysecret(): string;
    setKeysecret(value: string): DeletePermissionGroupReq;
    getGroupname(): string;
    setGroupname(value: string): DeletePermissionGroupReq;
    getPropagate(): boolean;
    setPropagate(value: boolean): DeletePermissionGroupReq;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): DeletePermissionGroupReq.AsObject;
    static toObject(includeInstance: boolean, msg: DeletePermissionGroupReq): DeletePermissionGroupReq.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: DeletePermissionGroupReq, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): DeletePermissionGroupReq;
    static deserializeBinaryFromReader(message: DeletePermissionGroupReq, reader: jspb.BinaryReader): DeletePermissionGroupReq;
}

export namespace DeletePermissionGroupReq {
    export type AsObject = {
        keyid: string,
        keysecret: string,
        groupname: string,
        propagate: boolean,
    }
}

export class ModifyPermissionGroupReq extends jspb.Message { 
    getKeyid(): string;
    setKeyid(value: string): ModifyPermissionGroupReq;
    getKeysecret(): string;
    setKeysecret(value: string): ModifyPermissionGroupReq;
    getGroupname(): string;
    setGroupname(value: string): ModifyPermissionGroupReq;
    getPermission(): string;
    setPermission(value: string): ModifyPermissionGroupReq;
    getPropagate(): boolean;
    setPropagate(value: boolean): ModifyPermissionGroupReq;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): ModifyPermissionGroupReq.AsObject;
    static toObject(includeInstance: boolean, msg: ModifyPermissionGroupReq): ModifyPermissionGroupReq.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: ModifyPermissionGroupReq, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): ModifyPermissionGroupReq;
    static deserializeBinaryFromReader(message: ModifyPermissionGroupReq, reader: jspb.BinaryReader): ModifyPermissionGroupReq;
}

export namespace ModifyPermissionGroupReq {
    export type AsObject = {
        keyid: string,
        keysecret: string,
        groupname: string,
        permission: string,
        propagate: boolean,
    }
}

export class ListPermissionGroupReq extends jspb.Message { 
    getKeyid(): string;
    setKeyid(value: string): ListPermissionGroupReq;
    getKeysecret(): string;
    setKeysecret(value: string): ListPermissionGroupReq;
    getGroupname(): string;
    setGroupname(value: string): ListPermissionGroupReq;
    getOffset(): string;
    setOffset(value: string): ListPermissionGroupReq;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): ListPermissionGroupReq.AsObject;
    static toObject(includeInstance: boolean, msg: ListPermissionGroupReq): ListPermissionGroupReq.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: ListPermissionGroupReq, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): ListPermissionGroupReq;
    static deserializeBinaryFromReader(message: ListPermissionGroupReq, reader: jspb.BinaryReader): ListPermissionGroupReq;
}

export namespace ListPermissionGroupReq {
    export type AsObject = {
        keyid: string,
        keysecret: string,
        groupname: string,
        offset: string,
    }
}

export class ListPermissionGroupRes extends jspb.Message { 
    clearMembersList(): void;
    getMembersList(): Array<PermissionGroupMembership>;
    setMembersList(value: Array<PermissionGroupMembership>): ListPermissionGroupRes;
    addMembers(value?: PermissionGroupMembership, index?: number): PermissionGroupMembership;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): ListPermissionGroupRes.AsObject;
    static toObject(includeInstance: boolean, msg: ListPermissionGroupRes): ListPermissionGroupRes.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: ListPermissionGroupRes, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): ListPermissionGroupRes;
    static deserializeBinaryFromReader(message: ListPermissionGroupRes, reader: jspb.BinaryReader): ListPermissionGroupRes;
}

export namespace ListPermissionGroupRes {
    export type AsObject = {
        membersList: Array<PermissionGroupMembership.AsObject>,
    }
}

export class PermissionGroupMembership extends jspb.Message { 
    getGroupname(): string;
    setGroupname(value: string): PermissionGroupMembership;
    getEntity(): string;
    setEntity(value: string): PermissionGroupMembership;
    getObject(): string;
    setObject(value: string): PermissionGroupMembership;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): PermissionGroupMembership.AsObject;
    static toObject(includeInstance: boolean, msg: PermissionGroupMembership): PermissionGroupMembership.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: PermissionGroupMembership, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): PermissionGroupMembership;
    static deserializeBinaryFromReader(message: PermissionGroupMembership, reader: jspb.BinaryReader): PermissionGroupMembership;
}

export namespace PermissionGroupMembership {
    export type AsObject = {
        groupname: string,
        entity: string,
        object: string,
    }
}
