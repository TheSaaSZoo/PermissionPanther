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
        this.key = config.key;
        this.target = config.endpoint;
        this.client = new main_grpc_pb_1.PermissionPantherClient(this.target, grpc.credentials.createInsecure());
    }
    /**
     * Check permission
     */
    async CheckPermission(input) {
        return new Promise((resolve, reject) => {
            const req = new permissions_pb_1.CheckDirectReq();
            req.setKey(this.key);
            req.setEntity(input.entity);
            req.setPermission(input.permission);
            req.setObject(input.object);
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
}
exports.default = PermissionPanther;
