"use strict"
import {arrayProp, prop, Ref, Typegoose} from "typegoose"
import IModelBase from "./modelBase"

class DataSet extends Typegoose implements IModelBase<DataSet> {

    @arrayProp({ items: String, default: [], required: true })
    public schema: string[]

    @prop({ default: [], required: true })
    public length: number

    @prop({ default: "", required: true })
    public url: string

    @prop({ ref: DataSet, required: false})
    public parent?: Ref<DataSet>

    @prop({ default: "", required: false})
    public jobId?: string

    public getModel() {
        return this.getModelForClass(DataSet)
    }
}

export default DataSet
