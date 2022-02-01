export class PermissionDenied extends Error {
	constructor() {
		super()

		Error.captureStackTrace(this, this.constructor)
		this.name = this.constructor.name
	}
}
