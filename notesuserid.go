/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_NOTESUSERID_CLASS.html */
package domigo

import (
	"github.com/monstermichl/domigo/internal/helpers"

	ole "github.com/go-ole/go-ole"
)

type NotesUserID struct {
	NotesStruct
}

func NewNotesUserID(dispatchPtr *ole.IDispatch) NotesUserID {
	return NotesUserID{NewNotesStruct(dispatchPtr)}
}

/* --------------------------------- Properties --------------------------------- */

/* --------------------------------- Methods ------------------------------------ */
/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GETENCRYPTIONKEYS_USERID.html */
func (u NotesUserID) GetEncryptionKeys() ([]String, error) {
	vals, err := callComArrayMethod(u, "GetEncryptionKeys")
	return helpers.CastSlice[String](vals), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GETUSERNAME_USERID.html */
func (u NotesUserID) GetUserName() (String, error) {
	val, err := callComMethod(u, "GetUserName")
	return helpers.CastValue[String](val), err
}
