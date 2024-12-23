package domigo

import (
	"fmt"
	"reflect"

	"github.com/monstermichl/domigo/internal/com"
	"github.com/monstermichl/domigo/internal/helpers"

	ole "github.com/go-ole/go-ole"
)

type notesStruct interface {
	Release()
	IsReady() bool
	com() com.Com
}

type primitiveType interface {
	bool | string | int | int8 | int16 | int32 | int64 | uint | uint8 | uint16 | uint32 | uint64 | uintptr | float32 | float64
}

type NotesStruct struct {
	c com.Com
}

func NewNotesStruct(dispatchPtr *ole.IDispatch) NotesStruct {
	return NotesStruct{com.New(dispatchPtr)}
}

func GetNotesStruct(val any) (NotesStruct, error) {
	var err error
	var s NotesStruct

	v := reflect.ValueOf(val)
	fieldName := reflect.TypeOf(NotesStruct{}).Name()
	field := v.FieldByName(fieldName)

	if !field.IsValid() {
		err = fmt.Errorf("no %s field found", fieldName)
	} else {
		s = field.Interface().(NotesStruct)
	}
	return s, err
}

func (s NotesStruct) Release() {
	s.c.Release()
}

func (s NotesStruct) IsReady() bool {
	return s.com().IsReady()
}

func (s NotesStruct) com() com.Com {
	return s.c
}

func getComObjectArrayProperty[T notesStruct](s notesStruct, modifyFn com.ModifyFunc[T], name string, params ...interface{}) ([]T, error) {
	return com.GetObjectArrayProperty(s.com(), modifyFn, name, params...)
}

func callComObjectArrayMethod[T notesStruct](s notesStruct, modifyFn com.ModifyFunc[T], name string, params ...interface{}) ([]T, error) {
	return com.CallObjectArrayMethod(s.com(), modifyFn, name, params...)
}

func callComMethod[T primitiveType](s notesStruct, name string, params ...interface{}) (T, error) {
	val, err := s.com().CallMethod(name, params...)
	return helpers.CastValue[T](val), err
}

func callComObjectMethod[T notesStruct](s notesStruct, createFn com.ModifyFunc[T], name string, params ...interface{}) (T, error) {
	return com.CallObjectMethod[T](s.com(), createFn, name, params...)
}

func callComArrayMethod[T primitiveType](s notesStruct, name string, params ...interface{}) ([]T, error) {
	vals, err := s.com().CallArrayMethod(name, params...)
	return helpers.CastSlice[T](vals), err
}

func getComProperty[T primitiveType](s notesStruct, name string, params ...interface{}) (T, error) {
	val, err := s.com().GetProperty(name, params...)
	return helpers.CastValue[T](val), err
}

func getComObjectProperty[T notesStruct](s notesStruct, createFn com.ModifyFunc[T], name string, params ...interface{}) (T, error) {
	dispatchPtr, err := s.com().GetObjectProperty(name, params...)
	return createFn(dispatchPtr), err
}

func getComArrayProperty[T primitiveType](s notesStruct, name string, params ...interface{}) ([]T, error) {
	vals, err := s.com().GetArrayProperty(name, params...)
	return helpers.CastSlice[T](vals), err
}

func putComProperty(s notesStruct, name string, params ...interface{}) error {
	return s.com().PutProperty(name, params...)
}

/* https://dusted.codes/using-go-generics-to-pass-struct-slices-for-interface-slices */
func dispatchSlice[T notesStruct](conns []T) []*ole.IDispatch {
	s := make([]*ole.IDispatch, len(conns))

	for i, c := range conns {
		s[i] = c.com().Dispatch()
	}
	return s
}
