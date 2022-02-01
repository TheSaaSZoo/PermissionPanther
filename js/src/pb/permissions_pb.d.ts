// package: 
// file: pb/permissions.proto

/* tslint:disable */
/* eslint-disable */

import * as jspb from "google-protobuf";

export class CheckDirectReq extends jspb.Message { 
    getKey(): string;
    setKey(value: string): CheckDirectReq;
    getEntity(): string;
    setEntity(value: string): CheckDirectReq;
    getPermission(): string;
    setPermission(value: string): CheckDirectReq;
    getObject(): string;
    setObject(value: string): CheckDirectReq;

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
        key: string,
        entity: string,
        permission: string,
        object: string,
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
    getKey(): string;
    setKey(value: string): ListEntityRelationsReq;
    getEntity(): string;
    setEntity(value: string): ListEntityRelationsReq;
    getPermission(): string;
    setPermission(value: string): ListEntityRelationsReq;

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
        key: string,
        entity: string,
        permission: string,
    }
}

export class ListObjectRelationsReq extends jspb.Message { 
    getKey(): string;
    setKey(value: string): ListObjectRelationsReq;
    getObject(): string;
    setObject(value: string): ListObjectRelationsReq;
    getPermission(): string;
    setPermission(value: string): ListObjectRelationsReq;

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
        key: string,
        object: string,
        permission: string,
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
