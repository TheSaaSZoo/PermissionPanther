"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.PermissionDenied = void 0;
class PermissionDenied extends Error {
    constructor() {
        super();
        Error.captureStackTrace(this, this.constructor);
        this.name = this.constructor.name;
    }
}
exports.PermissionDenied = PermissionDenied;
