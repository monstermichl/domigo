package notesuserid

import (
	"github.com/monstermichl/domigo/domino"
	"github.com/monstermichl/domigo/helpers"

	ole "github.com/go-ole/go-ole"
)

type NotesUserID struct {
	domino.NotesStruct
}

func New(dispatchPtr *ole.IDispatch) NotesUserID {
	return NotesUserID{domino.NewNotesStruct(dispatchPtr)}
}

/* --------------------------------- Properties --------------------------------- */

/* --------------------------------- Methods ------------------------------------ */
/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GETENCRYPTIONKEYS_USERID.html */
func (u NotesUserID) GetEncryptionKeys() ([]domino.String, error) {
	vals, err := u.Com().CallArrayMethod("GetEncryptionKeys")
	return helpers.CastSlice[string](vals), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GETUSERNAME_USERID.html */
func (u NotesUserID) GetUserName() (domino.String, error) {
	val, err := u.Com().CallMethod("GetUserName")
	return helpers.CastValue[domino.String](val), err
}
