/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_NOTESDATABASE_CLASS.html */
package notesdatabase

import (
	"domigo/domino"
	"domigo/domino/com"
	"domigo/domino/notesacl"
	"domigo/domino/notesform"
	"domigo/helpers"

	ole "github.com/go-ole/go-ole"
)

type NotesDatabase struct {
	domino.NotesStruct
}

func New(ptr *ole.IDispatch) NotesDatabase {
	return NotesDatabase{domino.NewNotesStruct(ptr)}
}

/* --------------------------------- Properties --------------------------------- */
func (d NotesDatabase) ACL() (notesacl.NotesACL, error) {
	dispatchPtr, err := d.Com.GetObjectProperty("ACL")

	if err != nil {
		return notesacl.NotesACL{}, err
	}
	return notesacl.New(dispatchPtr), err
}

func (d NotesDatabase) Forms() ([]notesform.NotesForm, error) {
	vals, err := com.GetObjectArrayProperty(d.Com, func(ptr *ole.IDispatch) notesform.NotesForm {
		return notesform.New(ptr)
	}, "Forms")
	return vals, err
}

func (d NotesDatabase) IsDesignLockingEnabled() (bool, error) {
	val, err := d.Com.GetProperty("IsDesignLockingEnabled")
	return helpers.CastValue[bool](val), err
}

func (d NotesDatabase) SetIsDesignLockingEnabled(e bool) error {
	return d.Com.PutProperty("IsDesignLockingEnabled", e)
}

func (d NotesDatabase) Title() (string, error) {
	val, err := d.Com.GetProperty("Title")
	return helpers.CastValue[string](val), err
}

/* --------------------------------- Methods ------------------------------------ */
