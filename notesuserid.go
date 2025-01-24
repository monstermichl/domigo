/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_NOTESUSERID_CLASS.html */
package domigo

import (
	ole "github.com/go-ole/go-ole"
)

type NotesUserID struct {
	NotesStruct
}

func newNotesUserID(dispatchPtr *ole.IDispatch) NotesUserID {
	return NotesUserID{newNotesStruct(dispatchPtr)}
}

/* --------------------------------- Properties --------------------------------- */

/* --------------------------------- Methods ------------------------------------ */
/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GETENCRYPTIONKEYS_USERID.html */
func (u NotesUserID) GetEncryptionKeys() ([]String, error) {
	return callComArrayMethod[String](u, "GetEncryptionKeys")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GETUSERNAME_USERID.html */
func (u NotesUserID) GetUserName() (String, error) {
	return callComMethod[String](u, "GetUserName")
}
