/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_NOTESACLENTRY_CLASS.html */
package notesaclentry

import (
	"domigo/domino"
	"domigo/helpers"

	ole "github.com/go-ole/go-ole"
)

type NotesACLEntry struct {
	domino.NotesStruct
}

func New(ptr *ole.IDispatch) NotesACLEntry {
	return NotesACLEntry{domino.NewNotesStruct(ptr)}
}

/* --------------------------------- Properties --------------------------------- */
func (e NotesACLEntry) CanCreateDocuments() (bool, error) {
	val, err := e.Com.GetProperty("CanCreateDocuments")
	return helpers.CastValue[bool](val), err
}

func (e NotesACLEntry) SetCanCreateDocuments(i bool) error {
	return e.Com.PutProperty("CanCreateDocuments", i)
}

/* --------------------------------- Methods ------------------------------------ */
