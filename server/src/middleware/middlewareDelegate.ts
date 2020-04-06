
import TokenCheckMiddleWare from "./tokenCheckMiddleWare"
import bodyParser from "body-parser"

export default class MiddlewareDelegate {
    public exec(app: any, router: any) {
        app.use(bodyParser.json())
        app.use( bodyParser.urlencoded( {
            extended: true
        }))

        new TokenCheckMiddleWare().exec(router)

        app.use("/", router)
    }
}