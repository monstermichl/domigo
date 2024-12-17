package com

/* Got started with this example: https://github.com/go-ole/go-ole/blob/master/_example/excel/excel.go.*/

import (
	"errors"
	"fmt"

	"github.com/monstermichl/domigo/helpers"

	ole "github.com/go-ole/go-ole"
	"github.com/go-ole/go-ole/oleutil"
)

type Com struct {
	dispatchPtr *ole.IDispatch
}

type ModifyFunc[T any] func(ptr *ole.IDispatch) T

var initialized = false

func New(dispatchPtr *ole.IDispatch) Com {
	return Com{dispatchPtr}
}

func CreateObject(id string) (Com, error) {
	var err error = nil
	var com Com

	/* Make sure COM interface got initialized. */
	if !initialized {
		err = ole.CoInitialize(0)
	}

	if err != nil {
		return com, errors.New("COM interface could not be initalized")
	}
	initialized = true
	unknown, err := oleutil.CreateObject(id)

	/* HINT: The CreateObject call failed, if the program wasn't compiled for 32 bit. */
	if err != nil {
		return com, errors.New("COM object could not be created")
	}
	dispatchPtr, err := unknown.QueryInterface(ole.IID_IDispatch)

	if err != nil {
		return com, errors.New("interface could not be retrieved")
	}
	return New(dispatchPtr), err
}

/* TODO: Find out if there's a better way... */
func GetObjectArrayProperty[T any](c Com, modifyFn ModifyFunc[T], name string, params ...interface{}) ([]T, error) {
	return objectArrayActionInternal(c.GetObjectArrayProperty, modifyFn, name, params...)
}

func CallObjectArrayMethod[T any](c Com, modifyFn ModifyFunc[T], name string, params ...interface{}) ([]T, error) {
	return objectArrayActionInternal(c.CallObjectArrayMethod, modifyFn, name, params...)
}

func (c Com) Dispatch() *ole.IDispatch {
	return c.dispatchPtr
}

func (c Com) Release() {
	err := c.checkPreconditions()

	if err == nil {
		c.dispatchPtr.Release()
	}
}

func (c Com) CallMethod(name string, params ...interface{}) (interface{}, error) {
	return c.valueActionInternal(c.callMethodInternal, name, params...)
}

func (c Com) CallObjectMethod(name string, params ...interface{}) (*ole.IDispatch, error) {
	return c.objectActionInternal(c.callMethodInternal, name, params...)
}

func (c Com) CallArrayMethod(name string, params ...interface{}) ([]interface{}, error) {
	return c.arrayActionInternal(c.callMethodInternal, name, params...)
}

func (c Com) CallObjectArrayMethod(name string, params ...interface{}) ([]*ole.IDispatch, error) {
	return c.objectArrayActionInternal(c.CallArrayMethod, name, params...)
}

func (c Com) GetProperty(name string, params ...interface{}) (interface{}, error) {
	return c.valueActionInternal(c.getPropertyInternal, name, params...)
}

func (c Com) GetObjectProperty(name string, params ...interface{}) (*ole.IDispatch, error) {
	return c.objectActionInternal(c.getPropertyInternal, name, params...)
}

func (c Com) GetArrayProperty(name string, params ...interface{}) ([]interface{}, error) {
	return c.arrayActionInternal(c.getPropertyInternal, name, params...)
}

func (c Com) GetObjectArrayProperty(name string, params ...interface{}) ([]*ole.IDispatch, error) {
	return c.objectArrayActionInternal(c.GetArrayProperty, name, params...)
}

func (c Com) PutProperty(name string, params ...interface{}) error {
	_, err := c.actionInternal(oleutil.PutProperty, name, params...)
	return err
}

func objectArrayActionInternal[T any](fn func(string, ...interface{}) ([]*ole.IDispatch, error), modifyFn ModifyFunc[T], name string, params ...interface{}) ([]T, error) {
	dispatchPtrs, err := fn(name, params...)

	if err != nil {
		return []T{}, err
	}
	objects := make([]T, len(dispatchPtrs))

	for i, ptr := range dispatchPtrs {
		objects[i] = modifyFn(ptr)
	}
	return objects, err
}

func (c Com) actionInternal(fn func(*ole.IDispatch, string, ...interface{}) (*ole.VARIANT, error), name string, params ...interface{}) (*ole.VARIANT, error) {
	err := c.checkPreconditions()

	if err != nil {
		return nil, err
	}
	variantPtr, err := fn(c.dispatchPtr, name, params...)

	if err != nil {
		return nil, fmt.Errorf("COM command execution failed: %s", err)
	}
	return variantPtr, err
}

func (c Com) valueActionInternal(fn func(string, ...interface{}) (*ole.VARIANT, error), name string, params ...interface{}) (interface{}, error) {
	variantPtr, err := fn(name, params...)

	if err != nil {
		return nil, err
	}
	return variantPtr.Value(), err
}

func (c Com) objectActionInternal(fn func(string, ...interface{}) (*ole.VARIANT, error), name string, params ...interface{}) (*ole.IDispatch, error) {
	variantPtr, err := fn(name, params...)

	if err != nil {
		return nil, err
	}
	return variantPtr.ToIDispatch(), err
}

func (c Com) arrayActionInternal(fn func(name string, params ...interface{}) (*ole.VARIANT, error), name string, params ...interface{}) ([]interface{}, error) {
	variantPtr, err := fn(name, params...)

	if err != nil {
		return []any{}, err
	}
	dispatchPtr := variantPtr.ToIDispatch()

	if dispatchPtr != nil {
		dispatchPtr.Release()
		return []any{}, fmt.Errorf("returned value for '%s' is not an array", name)
	}
	safeArray := variantPtr.ToArray()

	if safeArray == nil {
		return []any{}, err
	}
	return safeArray.ToValueArray(), err
}

func (c Com) objectArrayActionInternal(fn func(name string, params ...interface{}) ([]interface{}, error), name string, params ...interface{}) ([]*ole.IDispatch, error) {
	array, err := fn(name, params...)
	return helpers.CastSlice[*ole.IDispatch](array), err
}

func (c Com) getPropertyInternal(name string, params ...interface{}) (*ole.VARIANT, error) {
	return c.actionInternal(oleutil.GetProperty, name, params...)
}

func (c Com) callMethodInternal(name string, params ...interface{}) (*ole.VARIANT, error) {
	return c.actionInternal(oleutil.CallMethod, name, params...)
}

func (c Com) checkPreconditions() error {
	var err error = nil

	if c.dispatchPtr == nil {
		err = errors.New("invalid COM dispatch")
	}
	return err
}
