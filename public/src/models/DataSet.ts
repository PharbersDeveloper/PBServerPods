"use strict"
import {arrayProp, prop, Ref, Typegoose} from "typegoose"
import IModelBase from "./modelBase"

class DataSet extends Typegoose implements IModelBase<DataSet> {

    @arrayProp({ items: String, default: [], required: false } )
    public colNames?: string[]

    @prop({ default: 0, required: false} )
    public length?: number

    @prop({ default: "", required: false } )
    public url?: string

    @prop({ default: "", required: false } )
    public description?: string

    @prop({ ref: DataSet, required: false } )
    public parent?: Ref<DataSet>

    @prop({ default: "", required: true } )
    public jobId: string

    public getModel() {
        return this.getModelForClass(DataSet)
    }
}

export default DataSet
