"use strict"
import {arrayProp, prop, Ref, Typegoose} from "typegoose"
import FileMetaDatum from "./FileMetaDatum"
import IModelBase from "./modelBase"

class SandboxIndex extends Typegoose implements IModelBase<SandboxIndex> {

    @prop({ default: 0, required: true })
    public salesTarget?: number

    @prop({ default: "", required: true })
    public AccountID: string

    @arrayProp( { itemsRef: FileMetaDatum, required: true } )
    public fileVersions: Array<Ref<FileMetaDatum>>

    public getModel() {
        return this.getModelForClass(SandboxIndex)
    }
}

export default SandboxIndex
