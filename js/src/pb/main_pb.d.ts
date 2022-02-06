// package: 
// file: pb/main.proto

/* tslint:disable */
/* eslint-disable */

import * as jspb from "google-protobuf";
import * as pb_permissions_pb from "../pb/permissions_pb";

export class Applied extends jspb.Message { 
    getApplied(): boolean;
    setApplied(value: boolean): Applied;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): Applied.AsObject;
    static toObject(includeInstance: boolean, msg: Applied): Applied.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: Applied, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): Applied;
    static deserializeBinaryFromReader(message: Applied, reader: jspb.BinaryReader): Applied;
}

export namespace Applied {
    export type AsObject = {
        applied: boolean,
    }
}
