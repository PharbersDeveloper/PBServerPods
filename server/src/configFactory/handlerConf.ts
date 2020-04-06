"use strict"
import {JsonObject, JsonProperty} from "json2typescript"

@JsonObject("HandlerConf")
export class HandlerConf {
    @JsonProperty("file", String)
    public file: string = undefined

    @JsonProperty("entrance", String)
    public entrance: string = undefined
}
