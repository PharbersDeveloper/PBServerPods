// Code generated by github.com/actgardner/gogen-avro. DO NOT EDIT.
/*
 * SOURCE:
 *     UploadEnd.avsc
 */

package PhUploadEnd

import (
	"github.com/actgardner/gogen-avro/compiler"
	"github.com/actgardner/gogen-avro/container"
	"github.com/actgardner/gogen-avro/vm"
	"github.com/actgardner/gogen-avro/vm/types"
	"io"
)

type UploadEnd struct {
	DataSetId string
	TraceId   string
}

func NewUploadEndWriter(writer io.Writer, codec container.Codec, recordsPerBlock int64) (*container.Writer, error) {
	str := &UploadEnd{}
	return container.NewWriter(writer, codec, recordsPerBlock, str.Schema())
}

func DeserializeUploadEnd(r io.Reader) (*UploadEnd, error) {
	t := NewUploadEnd()
	err := deserializeField(r, t.Schema(), t.Schema(), t)
	return t, err
}

func DeserializeUploadEndFromSchema(r io.Reader, schema string) (*UploadEnd, error) {
	t := NewUploadEnd()
	err := deserializeField(r, schema, t.Schema(), t)
	return t, err
}

func NewUploadEnd() *UploadEnd {
	return &UploadEnd{}
}

func (r *UploadEnd) Schema() string {
	return "{\"fields\":[{\"name\":\"dataSetId\",\"type\":\"string\"},{\"name\":\"traceId\",\"type\":\"string\"}],\"name\":\"UploadEnd\",\"namespace\":\"com.pharbers.kafka.schema\",\"type\":\"record\"}"
}

func (r *UploadEnd) SchemaName() string {
	return "com.pharbers.kafka.schema.UploadEnd"
}

func (r *UploadEnd) Serialize(w io.Writer) error {
	return writeUploadEnd(r, w)
}

func (_ *UploadEnd) SetBoolean(v bool)    { panic("Unsupported operation") }
func (_ *UploadEnd) SetInt(v int32)       { panic("Unsupported operation") }
func (_ *UploadEnd) SetLong(v int64)      { panic("Unsupported operation") }
func (_ *UploadEnd) SetFloat(v float32)   { panic("Unsupported operation") }
func (_ *UploadEnd) SetDouble(v float64)  { panic("Unsupported operation") }
func (_ *UploadEnd) SetBytes(v []byte)    { panic("Unsupported operation") }
func (_ *UploadEnd) SetString(v string)   { panic("Unsupported operation") }
func (_ *UploadEnd) SetUnionElem(v int64) { panic("Unsupported operation") }
func (r *UploadEnd) Get(i int) types.Field {
	switch i {
	case 0:
		return (*types.String)(&r.DataSetId)
	case 1:
		return (*types.String)(&r.TraceId)

	}
	panic("Unknown field index")
}
func (r *UploadEnd) SetDefault(i int) {
	switch i {

	}
	panic("Unknown field index")
}
func (_ *UploadEnd) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ *UploadEnd) AppendArray() types.Field         { panic("Unsupported operation") }
func (_ *UploadEnd) Finalize()                        {}

type UploadEndReader struct {
	r io.Reader
	p *vm.Program
}

func NewUploadEndReader(r io.Reader) (*UploadEndReader, error) {
	containerReader, err := container.NewReader(r)
	if err != nil {
		return nil, err
	}

	t := NewUploadEnd()
	deser, err := compiler.CompileSchemaBytes([]byte(containerReader.AvroContainerSchema()), []byte(t.Schema()))
	if err != nil {
		return nil, err
	}

	return &UploadEndReader{
		r: containerReader,
		p: deser,
	}, nil
}

func (r *UploadEndReader) Read() (*UploadEnd, error) {
	t := NewUploadEnd()
	err := vm.Eval(r.r, r.p, t)
	return t, err
}
