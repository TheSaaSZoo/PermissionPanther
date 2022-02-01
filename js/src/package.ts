import { PantherConfig } from "./types";

export default class PermissionPanther {
  key: string
  constructor(config: PantherConfig) {
    this.key = config.key
  }
}
