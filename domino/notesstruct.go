package domino

import (
	"domigo/domino/com"

	ole "github.com/go-ole/go-ole"
)

type NotesStruct struct {
	Com com.Com
}

func NewNotesStruct(dispatchPtr *ole.IDispatch) NotesStruct {
	return NotesStruct{
		Com: com.New(dispatchPtr),
	}
}

/* --------------------------------- Properties --------------------------------- */
func (s NotesStruct) Parent() (interface{}, error) {
	val, err := s.Com.GetProperty("Parent")
	return val, err
}
