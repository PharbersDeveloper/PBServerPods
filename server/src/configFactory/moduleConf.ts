"use strict"
import {JsonObject, JsonProperty} from "json2typescript"

@JsonObject("ModuleConf")
export class ModuleConf {
    @JsonProperty("protocol", String)
    public protocol: string = "http"

    @JsonProperty("host", String)
    public host: string = undefined

    @JsonProperty("port", Number)
    public port: number = undefined

    @JsonProperty("routers", [String])
    public routers: string[] = undefined
}
