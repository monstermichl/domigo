/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_NOTESACL_CLASS.html */
package notesacl

import (
	"domigo/domino"
	"domigo/domino/notesaclentry"
	"domigo/helpers"

	ole "github.com/go-ole/go-ole"
)

type NotesACL struct {
	domino.NotesStruct
}

func New(dispatchPtr *ole.IDispatch) NotesACL {
	return NotesACL{domino.NewNotesStruct(dispatchPtr)}
}

func (a NotesACL) getEntry(funName string, params ...interface{}) (notesaclentry.NotesACLEntry, error) {
	dispatchPtr, err := a.Com.CallObjectMethod(funName, params...)

	if err != nil {
		return notesaclentry.NotesACLEntry{}, err
	}
	return notesaclentry.New(dispatchPtr), err
}

/* --------------------------------- Properties --------------------------------- */
func (a NotesACL) AdministrationServer() (string, error) {
	val, err := a.Com.GetProperty("AdministrationServer")
	return helpers.CastValue[string](val), err
}

func (a NotesACL) SetAdministrationServer(s string) error {
	return a.Com.PutProperty("AdministrationServer", s)
}

func (a NotesACL) InternetLevel() (int32, error) {
	val, err := a.Com.GetProperty("InternetLevel")
	return helpers.CastValue[int32](val), err
}

func (a NotesACL) SetInternetLevel(l int32) error {
	return a.Com.PutProperty("InternetLevel", l)
}

func (a NotesACL) IsAdminNames() (bool, error) {
	val, err := a.Com.GetProperty("IsAdminNames")
	return helpers.CastValue[bool](val), err
}

func (a NotesACL) SetIsAdminNames(i bool) error {
	return a.Com.PutProperty("IsAdminNames", i)
}

func (a NotesACL) IsAdminReaderAuthor() (bool, error) {
	val, err := a.Com.GetProperty("IsAdminReaderAuthor")
	return helpers.CastValue[bool](val), err
}

func (a NotesACL) SetIsAdminReaderAuthor(i bool) error {
	return a.Com.PutProperty("IsAdminReaderAuthor", i)
}

func (a NotesACL) IsExtendedAccess() (bool, error) {
	val, err := a.Com.GetProperty("IsExtendedAccess")
	return helpers.CastValue[bool](val), err
}

func (a NotesACL) SetIsExtendedAccess(i bool) error {
	return a.Com.PutProperty("IsExtendedAccess", i)
}

func (a NotesACL) Roles() ([]string, error) {
	vals, err := a.Com.GetArrayProperty("Roles")
	return helpers.CastSlice[string](vals), err
}

func (a NotesACL) UniformAccess() (bool, error) {
	val, err := a.Com.GetProperty("UniformAccess")
	return helpers.CastValue[bool](val), err
}

func (a NotesACL) SetUniformAccess(u bool) error {
	return a.Com.PutProperty("UniformAccess", u)
}

/* --------------------------------- Methods ------------------------------------ */
func (a NotesACL) AddRole(name string) error {
	_, err := a.Com.CallMethod("AddRole", name)
	return err
}

func (a NotesACL) CreateACLEntry(name string, level int32) error {
	_, err := a.Com.CallMethod("CreateACLEntry", name, level)
	return err
}

func (a NotesACL) DeleteRole(name string) error {
	_, err := a.Com.CallMethod("DeleteRole", name)
	return err
}

func (a NotesACL) GetEntry(name string) (notesaclentry.NotesACLEntry, error) {
	return a.getEntry("GetEntry", name)
}

func (a NotesACL) GetFirstEntry() (notesaclentry.NotesACLEntry, error) {
	return a.getEntry("GetFirstEntry")
}

/* TODO: Find out how to implement GetNextEntry, because entry needs to be passed
   https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GETNEXTENTRY_METHOD.html. */
// func (a NotesACL) GetNextEntry() (notesaclentry.NotesACLEntry, error) {
// 	return a.getEntry("GetNextEntry")
// }

func (a NotesACL) RemoveACLEntry(name string) error {
	_, err := a.Com.CallMethod("RemoveACLEntry", name)
	return err
}

func (a NotesACL) RenameRole(old string, new string) error {
	_, err := a.Com.CallMethod("RenameRole", old, new)
	return err
}

func (a NotesACL) Save() error {
	_, err := a.Com.CallMethod("Save")
	return err
}
