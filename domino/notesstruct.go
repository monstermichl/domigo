package domino

import (
	"domigo/domino/com"

	ole "github.com/go-ole/go-ole"
)

type NotesStruct struct {
	com com.Com
}

func NewNotesStruct(dispatchPtr *ole.IDispatch) NotesStruct {
	return NotesStruct{com.New(dispatchPtr)}
}

func (s NotesStruct) Com() com.Com {
	return s.com
}

func (s NotesStruct) Release() {
	s.Com().Release()
}
