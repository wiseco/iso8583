package iso8583

import (
	"fmt"
	"log"
	"reflect"

	"github.com/moov-io/iso8583/field"
)

type Message struct {
	fields    map[int]field.Field
	spec      *MessageSpec
	data      interface{}
	fieldsMap map[int]struct{}
	bitmap    *field.Bitmap
}

func NewMessage(spec *MessageSpec) *Message {
	fields := spec.CreateMessageFields()

	return &Message{
		fields:    fields,
		spec:      spec,
		fieldsMap: map[int]struct{}{},
	}
}

func (m *Message) Data() interface{} {
	return m.data
}

func (m *Message) SetData(data interface{}) error {
	m.data = data

	if m.data == nil {
		return nil
	}

	// get the struct
	str := reflect.ValueOf(data).Elem()

	if reflect.TypeOf(str).Kind() != reflect.Struct {
		return fmt.Errorf("failed to set data as struct is expected, got: %s", reflect.TypeOf(str).Kind())
	}

	for i, fl := range m.fields {
		fieldName := fmt.Sprintf("F%d", i)

		// get the struct field
		dataField := str.FieldByName(fieldName)

		if dataField == (reflect.Value{}) || dataField.IsNil() {
			continue
		}

		if dataField.Type() != reflect.TypeOf(fl) {
			return fmt.Errorf("failed to set data: type of the field %d: %v does not match the type in the spec: %v", i, dataField.Type(), reflect.TypeOf(fl))
		}

		// set data field spec for the message spec field
		specField := m.fields[i]
		df := dataField.Interface().(field.Field)
		df.SetSpec(specField.Spec())

		// use data field as a message field
		m.fields[i] = df
		m.fieldsMap[i] = struct{}{}
	}

	return nil
}

func (m *Message) Bitmap() *field.Bitmap {
	if m.bitmap != nil {
		return m.bitmap
	}

	m.bitmap = m.fields[1].(*field.Bitmap)
	m.fieldsMap[1] = struct{}{}

	return m.bitmap
}

func (m *Message) MTI(val string) {
	m.fieldsMap[0] = struct{}{}
	m.fields[0].SetBytes([]byte(val))
}

func (m *Message) Field(id int, val string) {
	m.fieldsMap[id] = struct{}{}
	m.fields[id].SetBytes([]byte(val))
}

func (m *Message) BinaryField(id int, val []byte) {
	m.fieldsMap[id] = struct{}{}
	m.fields[id].SetBytes(val)
}

func (m *Message) GetMTI() string {
	// check index
	return m.fields[0].String()
}

func (m *Message) GetString(id int) string {
	if _, ok := m.fieldsMap[id]; ok {
		return m.fields[id].String()
	}

	return ""
}

func (m *Message) GetBytes(id int) []byte {
	if _, ok := m.fieldsMap[id]; ok {
		return m.fields[id].Bytes()
	}

	return nil
}

func (m *Message) Pack() ([]byte, error) {
	packed := []byte{}
	m.Bitmap().Reset()

	// build the bitmap
	maxId := 0

	for id := range m.fieldsMap {
		if id > maxId {
			maxId = id
		}

		// indexes 0 and 1 are for mti and bitmap
		// regular field number startd from index 2
		if id < 2 {
			continue
		}

		m.Bitmap().Set(id)
	}

	// pack fields
	for i := 0; i <= maxId; i++ {
		if _, ok := m.fieldsMap[i]; ok {
			field, ok := m.fields[i]
			if !ok {
				return nil, fmt.Errorf("failed to pack field %d: no specification found", i)
			}

			packedField, err := field.Pack()
			if err != nil {
				return nil, fmt.Errorf("failed to pack field %d (%s): %v", i, field.Spec().Description, err)
			}
			packed = append(packed, packedField...)
		}
	}

	return packed, nil
}

func (m *Message) Unpack(src []byte) error {
	var off int
	log.Println("Unpacking started")

	m.fieldsMap = map[int]struct{}{}
	m.Bitmap().Reset()

	// unpack MTI
	read, err := m.fields[0].Unpack(src)
	if err != nil {
		return err
	}

	// Set fields map
	m.fieldsMap[0] = struct{}{}

	off = read

	// unpack Bitmap
	read, err = m.fields[1].Unpack(src[off:])
	if err != nil {
		return err
	}

	// Set in fields map
	m.fieldsMap[1] = struct{}{}

	off += read

	for i := 2; i <= m.Bitmap().Len(); i++ {
		log.Println("Unpacking of each bitmap", i, off, m.Bitmap().IsSet(i))

		if m.Bitmap().IsSet(i) {
			field, ok := m.fields[i]
			if !ok {
				return fmt.Errorf("failed to unpack field %d: no specification found", i)
			}

			m.fieldsMap[i] = struct{}{}
			read, err = field.Unpack(src[off:])
			if err != nil {
				return fmt.Errorf("failed to unpack field %d (%s): %v", i, field.Spec().Description, err)
			}

			err = m.linkDataFieldWithMessageField(i, field)
			if err != nil {
				return fmt.Errorf("failed to unpack field %d: %v", i, err)
			}
			off += read

			log.Println("End of Unpacking of each bitmap", off, read)
		}
	}

	return nil
}

func (m *Message) linkDataFieldWithMessageField(i int, fl field.Field) error {
	if m.data == nil {
		return nil
	}

	// get the struct
	str := reflect.ValueOf(m.data).Elem()

	fieldName := fmt.Sprintf("F%d", i)

	// get the struct field
	dataField := str.FieldByName(fieldName)
	if dataField == (reflect.Value{}) {
		return nil
	}

	if dataField.Type() != reflect.TypeOf(fl) {
		return fmt.Errorf("field type: %v does not match the type in the spec: %v", dataField.Type(), reflect.TypeOf(fl))
	}

	dataField.Addr().Elem().Set(reflect.ValueOf(fl))

	return nil
}
