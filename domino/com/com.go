package com

/* Got started with this example: https://github.com/go-ole/go-ole/blob/master/_example/excel/excel.go.*/

import (
	"domigo/helpers"
	"errors"
	"fmt"

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
	dispatchPtrs, err := c.GetObjectArrayProperty(name, params...)

	if err != nil {
		return []T{}, err
	}
	objects := make([]T, len(dispatchPtrs))

	for i, ptr := range dispatchPtrs {
		objects[i] = modifyFn(ptr)
	}
	return objects, err
}

func (c Com) Release() {
	err := c.checkPreconditions()

	if err == nil {
		c.dispatchPtr.Release()
	}
}

func (c Com) CallMethod(name string, params ...interface{}) (interface{}, error) {
	variantPtr, err := c.CallMethodInternal(name, params...)

	if err != nil {
		return nil, err
	}
	return variantPtr.Value(), err
}

func (c Com) CallObjectMethod(name string, params ...interface{}) (*ole.IDispatch, error) {
	variantPtr, err := c.CallMethodInternal(name, params...)

	if err != nil {
		return nil, err
	}
	return variantPtr.ToIDispatch(), err
}

func (c Com) GetProperty(name string, params ...interface{}) (interface{}, error) {
	err := c.checkPreconditions()

	if err != nil {
		return nil, err
	}
	variantPtr, err := oleutil.GetProperty(c.dispatchPtr, name, params...)

	if err != nil {
		return nil, fmt.Errorf("property '%s' could not be retrieved", name)
	}
	return variantPtr.Value(), err
}

func (c Com) GetObjectProperty(name string, params ...interface{}) (*ole.IDispatch, error) {
	variantPtr, err := c.getPropertyInternal(name, params...)

	if err != nil {
		return nil, err
	}
	return variantPtr.ToIDispatch(), err
}

func (c Com) GetArrayProperty(name string, params ...interface{}) ([]interface{}, error) {
	variantPtr, err := c.getPropertyInternal(name, params...)

	if err != nil {
		return []any{}, err
	}

	if variantPtr.ToIDispatch() != nil {
		return []any{}, fmt.Errorf("property '%s' is not an array property", name)
	}
	return variantPtr.ToArray().ToValueArray(), err
}

func (c Com) GetObjectArrayProperty(name string, params ...interface{}) ([]*ole.IDispatch, error) {
	array, err := c.GetArrayProperty(name, params...)
	return helpers.CastSlice[*ole.IDispatch](array), err
}

func (c Com) PutProperty(name string, params ...interface{}) error {
	err := c.checkPreconditions()

	if err != nil {
		return err
	}
	_, err = c.dispatchPtr.PutProperty(name, params...)

	if err != nil {
		err = fmt.Errorf("property '%s' could not be set", name)
	}
	return err
}

func (c Com) CallMethodInternal(name string, params ...interface{}) (*ole.VARIANT, error) {
	err := c.checkPreconditions()

	if err != nil {
		return nil, err
	}
	variantPtr, err := oleutil.CallMethod(c.dispatchPtr, name, params...)

	if err != nil {
		return nil, fmt.Errorf("method '%s' could not be called", name)
	}
	return variantPtr, err
}

func (c Com) getPropertyInternal(name string, params ...interface{}) (*ole.VARIANT, error) {
	err := c.checkPreconditions()

	if err != nil {
		return nil, err
	}
	variantPtr, err := oleutil.GetProperty(c.dispatchPtr, name, params...)

	if err != nil {
		return nil, fmt.Errorf("property '%s' could not be retrieved", name)
	}
	return variantPtr, err
}

func (c Com) checkPreconditions() error {
	var err error = nil

	if c.dispatchPtr == nil {
		err = errors.New("invalid COM dispatch")
	}
	return err
}
