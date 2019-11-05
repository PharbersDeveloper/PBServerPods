"use strict"
import {arrayProp, prop, Ref, Typegoose} from "typegoose"
import IModelBase from "./modelBase"

class DbSource extends Typegoose implements IModelBase<DbSource> {

    @prop({ default: "", required: true })
    public dbType: string

    @prop({ default: "", required: true })
    public url: string

    @prop({ default: "", required: true })
    public userName: number

    @prop({ default: "", required: true })
    public pwd: number

    @prop({ default: "", required: true })
    public dbName: number

    public getModel() {
        return this.getModelForClass(DbSource)
    }
}

export default DbSource
