"use strict"
import {arrayProp, prop, Ref, Typegoose} from "typegoose"
import FileVersion from "./FileVersion"
import IModelBase from "./modelBase"

class FileMetaDatum extends Typegoose implements IModelBase<FileMetaDatum> {

    @prop({ default: "", required: true })
    public name: string

    @prop({ default: "", required: true })
    public extension: string

    @prop({ default: 0, required: true })
    public created: number

    @prop({ default: "", required: true })
    public Kind: string

    @prop({ default: "", required: true })
    public ownerID: string

    @prop({ default: "", required: true })
    public ownerName: string

    @prop({ default: "", required: true })
    public groupID: string

    @prop({ default: "", required: true })
    public mod: string

    @prop({ default: 0.0, required: true })
    public size: number

    @prop({ default: "", required: true })
    public where: string

    @prop({ default: "", required: true })
    public kind: string

    @prop({ default: "", required: true })
    public tag: string

    @arrayProp( { itemsRef: FileVersion, required: true } )
    public fileVersions: Ref<FileVersion>[]

    public getModel() {
        return this.getModelForClass(FileMetaDatum)
    }
}

export default FileMetaDatum
