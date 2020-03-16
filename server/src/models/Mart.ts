"use strict"
import {arrayProp, prop, Ref, Typegoose} from "typegoose"
import IModelBase from "./modelBase"
import DataSet from "./DataSet"

class Mart extends Typegoose implements IModelBase<Mart> {

    @arrayProp({ itemsRef: DataSet, required: false } )
    public dfs: Ref<DataSet>[]

    @prop({ default: "", required: true })
    public name: string

    @prop({ default: "", required: true })
    public url: string

    @prop({ default: "", required: true })
    public dataType: string

    @prop({ default: "", required: false })
    public description: string

    public getModel() {
        return this.getModelForClass(Mart)
    }
}

export default Mart
