package domigo

import (
	"github.com/monstermichl/domigo/internal/com"

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

/* https://dusted.codes/using-go-generics-to-pass-struct-slices-for-interface-slices */
func dispatchSlice[T notesStruct](conns []T) []*ole.IDispatch {
	s := make([]*ole.IDispatch, len(conns))

	for i, c := range conns {
		s[i] = c.com().Dispatch()
	}
	return s
}
