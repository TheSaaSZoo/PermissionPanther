// package: 
// file: pb/main.proto

/* tslint:disable */
/* eslint-disable */

import * as jspb from "google-protobuf";
import * as pb_permissions_pb from "../pb/permissions_pb";

export class NoContent extends jspb.Message { 

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): NoContent.AsObject;
    static toObject(includeInstance: boolean, msg: NoContent): NoContent.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: NoContent, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): NoContent;
    static deserializeBinaryFromReader(message: NoContent, reader: jspb.BinaryReader): NoContent;
}

export namespace NoContent {
    export type AsObject = {
    }
}
