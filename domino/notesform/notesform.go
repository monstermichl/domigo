package notesform

import (
	"domigo/domino"
	"domigo/helpers"

	ole "github.com/go-ole/go-ole"
)

type NotesForm struct {
	domino.NotesStruct
}

func New(dispatchPtr *ole.IDispatch) NotesForm {
	return NotesForm{domino.NewNotesStruct(dispatchPtr)}
}

/* --------------------------------- Properties --------------------------------- */
func (n NotesForm) Aliases() ([]string, error) {
	vals, err := n.Com.GetArrayProperty("Aliases")
	return helpers.CastSlice[string](vals), err
}

func (n NotesForm) Fields() ([]string, error) {
	vals, err := n.Com.GetArrayProperty("Fields")
	return helpers.CastSlice[string](vals), err
}

func (n NotesForm) FormUsers() ([]string, error) {
	vals, err := n.Com.GetArrayProperty("FormUsers")
	return helpers.CastSlice[string](vals), err
}

func (n NotesForm) SetFormUsers(users []string) error {
	return n.Com.PutProperty("FormUsers", users)
}

func (n NotesForm) HttpURL() (string, error) {
	val, err := n.Com.GetProperty("HttpURL")
	return helpers.CastValue[string](val), err
}

func (n NotesForm) IsSubForm() (bool, error) {
	val, err := n.Com.GetProperty("IsSubForm")
	return helpers.CastValue[bool](val), err
}

func (n NotesForm) LockHolders() ([]string, error) {
	vals, err := n.Com.GetArrayProperty("LockHolders")
	return helpers.CastSlice[string](vals), err
}

func (n NotesForm) Name() (string, error) {
	val, err := n.Com.GetProperty("Name")
	return helpers.CastValue[string](val), err
}

func (n NotesForm) NotesURL() (string, error) {
	val, err := n.Com.GetProperty("NotesURL")
	return helpers.CastValue[string](val), err
}

func (n NotesForm) ProtectReaders() (bool, error) {
	val, err := n.Com.GetProperty("ProtectReaders")
	return helpers.CastValue[bool](val), err
}

func (n NotesForm) SetProtectReaders(protect bool) error {
	return n.Com.PutProperty("ProtectReaders", protect)
}

func (n NotesForm) ProtectUsers() (bool, error) {
	val, err := n.Com.GetProperty("ProtectUsers")
	return helpers.CastValue[bool](val), err
}

func (n NotesForm) SetProtectUsers(protect bool) error {
	return n.Com.PutProperty("ProtectUsers", protect)
}

func (n NotesForm) Readers() ([]string, error) {
	vals, err := n.Com.GetArrayProperty("Readers")
	return helpers.CastSlice[string](vals), err
}

func (n NotesForm) SetReaders(readers []string) error {
	return n.Com.PutProperty("Readers", readers)
}

/* --------------------------------- Methods ------------------------------------ */
func (n NotesForm) GetFieldType(name string) (int, error) {
	val, err := n.Com.CallMethod("GetFieldType", name)
	return helpers.CastValue[int](val), err
}

func (n NotesForm) Lock(names []string, provisional bool) (bool, error) {
	val, err := n.Com.CallMethod("Lock", names, provisional)
	return helpers.CastValue[bool](val), err
}

func (n NotesForm) LockProvisional(names []string) (bool, error) {
	return n.Lock(names, true)
}

func (n NotesForm) Remove() error {
	_, err := n.Com.CallMethod("Remove")
	return err
}

func (n NotesForm) UnLock() error {
	_, err := n.Com.CallMethod("UnLock")
	return err
}
