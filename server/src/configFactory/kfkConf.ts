"use strict"

import {JsonObject, JsonProperty} from "json2typescript"

@JsonObject("KfkConf")
export class KfkConf {

    @JsonProperty("kafkaBrokerList", String)
    public brokerLst: string = undefined

    @JsonProperty("kafkaTopic", String)
    public kafkaTopic: string = undefined
}
