"use strict";
var __createBinding = (this && this.__createBinding) || (Object.create ? (function(o, m, k, k2) {
    if (k2 === undefined) k2 = k;
    Object.defineProperty(o, k2, { enumerable: true, get: function() { return m[k]; } });
}) : (function(o, m, k, k2) {
    if (k2 === undefined) k2 = k;
    o[k2] = m[k];
}));
var __setModuleDefault = (this && this.__setModuleDefault) || (Object.create ? (function(o, v) {
    Object.defineProperty(o, "default", { enumerable: true, value: v });
}) : function(o, v) {
    o["default"] = v;
});
var __importStar = (this && this.__importStar) || function (mod) {
    if (mod && mod.__esModule) return mod;
    var result = {};
    if (mod != null) for (var k in mod) if (k !== "default" && Object.prototype.hasOwnProperty.call(mod, k)) __createBinding(result, mod, k);
    __setModuleDefault(result, mod);
    return result;
};
Object.defineProperty(exports, "__esModule", { value: true });
const grpc = __importStar(require("@grpc/grpc-js"));
const main_grpc_pb_1 = require("./pb/main_grpc_pb");
const permissions_pb_1 = require("./pb/permissions_pb");
const errors_1 = require("./errors");
class PermissionPanther {
    constructor(config) {
        this.keyID = config.keyID;
        this.keySecret = config.keySecret;
        this.target = config.endpoint;
        if (config.insecure === true) {
            this.client = new main_grpc_pb_1.PermissionPantherClient(this.target, grpc.credentials.createInsecure());
        }
        else {
            this.client = new main_grpc_pb_1.PermissionPantherClient(this.target, grpc.credentials.createSsl());
        }
    }
    /**
     * Checks whether an entity has a permission on an object. Optionally specify explicity deny permission, and group inheritance checks.
     */
    async CheckPermission(input) {
        return new Promise((resolve, reject) => {
            const req = new permissions_pb_1.CheckDirectReq();
            req.setKeyid(this.keyID);
            req.setKeysecret(this.keySecret);
            req.setEntity(input.entity);
            req.setPermission(input.permission);
            req.setObject(input.object);
            if (input.denyPermission) {
                req.setDenypermission(input.denyPermission);
            }
            if (input.inheritance === false) {
                req.setRecursive(false);
            }
            else {
                req.setRecursive(true);
            }
            this.client.checkDirectPermission(req, (err, res) => {
                if (err) {
                    switch (err.code) {
                        case grpc.status.PERMISSION_DENIED:
                            reject(new errors_1.PermissionDenied());
                        default:
                            reject(err);
                    }
                    return;
                }
                resolve({
                    recursion: res.getRecursion(),
                    valid: res.getValid()
                });
            });
        });
    }
    /**
     * Lists an entity's relations to find what objects they have permission on. Optionally specify a `permission` to look for objects that the entity has a specific permission on.
     */
    async ListEntityRelations(input) {
        return new Promise((resolve, reject) => {
            const req = new permissions_pb_1.ListEntityRelationsReq();
            req.setKeyid(this.keyID);
            req.setKeysecret(this.keySecret);
            req.setEntity(input.entity);
            if (input.permission) {
                req.setPermission(input.permission);
            }
            this.client.listEntityRelations(req, (err, res) => {
                if (err) {
                    switch (err.code) {
                        case grpc.status.PERMISSION_DENIED:
                            reject(new errors_1.PermissionDenied());
                        default:
                            reject(err);
                    }
                }
                const rel = [];
                for (const r of res.getRelationsList()) {
                    rel.push({
                        entity: r.getEntity(),
                        object: r.getObject(),
                        permission: r.getPermission()
                    });
                }
                resolve({
                    offset: "",
                    relations: rel
                });
            });
        });
    }
    /**
     * Lists an object's relations to find what entities have permission on it. Optionally specify a `permission` to look for entities who have a specific permission on the object.
     */
    async ListObjectRelations(input) {
        return new Promise((resolve, reject) => {
            const req = new permissions_pb_1.ListObjectRelationsReq();
            req.setKeyid(this.keyID);
            req.setKeysecret(this.keySecret);
            req.setObject(input.object);
            if (input.permission) {
                req.setPermission(input.permission);
            }
            this.client.listObjectRelations(req, (err, res) => {
                if (err) {
                    switch (err.code) {
                        case grpc.status.PERMISSION_DENIED:
                            reject(new errors_1.PermissionDenied());
                        default:
                            reject(err);
                    }
                }
                const rel = [];
                for (const r of res.getRelationsList()) {
                    rel.push({
                        entity: r.getEntity(),
                        object: r.getObject(),
                        permission: r.getPermission()
                    });
                }
                resolve({
                    offset: "",
                    relations: rel
                });
            });
        });
    }
    /**
     * Sets a permission.
     * Returns whether the relation was created (did not exist).
     */
    async SetPermission(input) {
        return new Promise((resolve, reject) => {
            const req = new permissions_pb_1.RelationReq();
            req.setEntity(input.entity);
            req.setKeyid(this.keyID);
            req.setKeysecret(this.keySecret);
            req.setObject(input.object);
            req.setPermission(input.permission);
            this.client.setPermission(req, (err, res) => {
                if (err) {
                    switch (err.code) {
                        case grpc.status.PERMISSION_DENIED:
                            reject(new errors_1.PermissionDenied());
                        default:
                            reject(err);
                    }
                }
                resolve(res.getApplied());
            });
        });
    }
    /**
     * Removes a permission.
     * Returns whether the relation was deleted (existed).
     */
    async RemovePermission(input) {
        return new Promise((resolve, reject) => {
            const req = new permissions_pb_1.RelationReq();
            req.setEntity(input.entity);
            req.setKeyid(this.keyID);
            req.setKeysecret(this.keySecret);
            req.setObject(input.object);
            req.setPermission(input.permission);
            this.client.removePermission(req, (err, res) => {
                if (err) {
                    switch (err.code) {
                        case grpc.status.PERMISSION_DENIED:
                            reject(new errors_1.PermissionDenied());
                        default:
                            reject(err);
                    }
                }
                resolve(res.getApplied());
            });
        });
    }
}
exports.default = PermissionPanther;
