package domigo

import (
	"errors"
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

type NotesStruct struct {
	c com.Com
}

func NewNotesStruct(dispatchPtr *ole.IDispatch) NotesStruct {
	return NotesStruct{com.New(dispatchPtr)}
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

func getComObjectArrayProperty[T notesStruct](s notesStruct, modifyFn com.ModifyFunc[T], name string, params ...any) ([]T, error) {
	params = convertParams(params...)
	return com.GetObjectArrayProperty(s.com(), modifyFn, name, params...)
}

func callComObjectArrayMethod[T notesStruct](s notesStruct, modifyFn com.ModifyFunc[T], name string, params ...any) ([]T, error) {
	params = convertParams(params...)
	return com.CallObjectArrayMethod(s.com(), modifyFn, name, params...)
}

func callComMethod[T primitiveType](s notesStruct, name string, params ...any) (T, error) {
	params = convertParams(params...)
	val, err := s.com().CallMethod(name, params...)
	return helpers.CastValue[T](val), err
}

func callComAnyMethod(s notesStruct, name string, params ...any) (any, error) {
	params = convertParams(params...)
	return s.com().CallMethod(name, params...)
}

func callComVoidMethod(s notesStruct, name string, params ...any) error {
	params = convertParams(params...)
	_, err := s.com().CallMethod(name, params...)
	return err
}

func callComObjectMethod[T notesStruct](s notesStruct, createFn com.ModifyFunc[T], name string, params ...any) (T, error) {
	params = convertParams(params...)
	return com.CallObjectMethod[T](s.com(), createFn, name, params...)
}

func callComArrayMethod[T primitiveType](s notesStruct, name string, params ...any) ([]T, error) {
	params = convertParams(params...)
	vals, err := s.com().CallArrayMethod(name, params...)
	return helpers.CastSlice[T](vals), err
}

func callComAnyArrayMethod(s notesStruct, name string, params ...any) (any, error) {
	params = convertParams(params...)
	return s.com().CallArrayMethod(name, params...)
}

func getComProperty[T primitiveType](s notesStruct, name string, params ...any) (T, error) {
	params = convertParams(params...)
	val, err := s.com().GetProperty(name, params...)
	return helpers.CastValue[T](val), err
}

func getComAnyProperty(s notesStruct, name string, params ...any) (any, error) {
	params = convertParams(params...)
	return s.com().GetProperty(name, params...)
}

func getComObjectProperty[T notesStruct](s notesStruct, createFn com.ModifyFunc[T], name string, params ...any) (T, error) {
	params = convertParams(params...)
	dispatchPtr, err := s.com().GetObjectProperty(name, params...)
	return createFn(dispatchPtr), err
}

func getComArrayProperty[T primitiveType](s notesStruct, name string, params ...any) ([]T, error) {
	params = convertParams(params...)
	vals, err := s.com().GetArrayProperty(name, params...)
	return helpers.CastSlice[T](vals), err
}

func getComAnyArrayProperty(s notesStruct, name string, params ...any) ([]any, error) {
	params = convertParams(params...)
	return s.com().GetArrayProperty(name, params...)
}

func putComProperty(s notesStruct, name string, params ...any) error {
	params = convertParams(params...)
	return s.com().PutProperty(name, params...)
}

func getNotesStruct(val any) (NotesStruct, error) {
	var err error
	var s NotesStruct

	if reflect.TypeOf(val).Kind() == reflect.Struct {
		v := reflect.ValueOf(val)
		fieldName := reflect.TypeOf(NotesStruct{}).Name()
		field := v.FieldByName(fieldName)

		if !field.IsValid() {
			err = fmt.Errorf("no %s field found", fieldName)
		} else {
			s = field.Interface().(NotesStruct)
		}
	} else {
		err = errors.New("value is not a struct")
	}
	return s, err
}

func convertParams(params ...any) []any {
	for i, param := range params {
		if s, err := getNotesStruct(param); err == nil {
			params[i] = s.com().Dispatch()
		} else if reflect.TypeOf(param).Kind() == reflect.Slice {
			params[i] = convertParams(param.([]any)...)
		}
	}
	return params
}

/* https://dusted.codes/using-go-generics-to-pass-struct-slices-for-interface-slices */
func dispatchSlice[T notesStruct](conns []T) []*ole.IDispatch {
	s := make([]*ole.IDispatch, len(conns))

	for i, c := range conns {
		s[i] = c.com().Dispatch()
	}
	return s
}
