import {ServerConf} from "../configFactory/serverConf"
import {SingletonInitConf} from "../configFactory/singletonConf"

export const CONFIG: ServerConf = new SingletonInitConf().getConf()

