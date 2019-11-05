"use strict"
import {arrayProp, prop, Ref, Typegoose} from "typegoose"
import DataSet from "./DataSet"
import DbSource from "./DbSource"
import IModelBase from "./modelBase"
import File from "./File"

class Assent extends Typegoose implements IModelBase<Assent> {

    @prop({ default: "", required: true })
    public traceId: string

    @prop({default: "", required: true})
    public assent: string

    @prop({default: 0, required: true})
    public version: number

    @prop({ default: "file", required: true })
    public dataType: string // candidate: database, file, stream, application

    @prop({ ref: File, required: false} )
    public file?: Ref<File>

    @prop({ ref: DbSource, required: false} )
    public dbs?: Ref<DbSource>

    @arrayProp({ itemsRef: DataSet, required: false} )
    public dfs?: Ref<DataSet>[]

    @arrayProp({ items: String, default: [], required: true} )
    public providers: string[]

    @arrayProp({ items: String, default: [], required: true} )
    public markets: string[]

    @arrayProp({ items: String, default: [], required: true} )
    public molecules: string[]

    @arrayProp({ items: Number, default: [], required: true} )
    public dataCover: string[]

    @arrayProp({ items: String, default: [], required: true} )
    public geoCover: string[]

    @arrayProp({ items: String, default: [], required: true} )
    public Labels: string[]

    public getModel() {
        return this.getModelForClass(Assent)
    }
}

export default Assent
