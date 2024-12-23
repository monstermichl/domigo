/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_NOTESVIEWNAVIGATOR_CLASS.html */
package domigo

import (
	"github.com/monstermichl/domigo/internal/helpers"

	ole "github.com/go-ole/go-ole"
)

type NotesViewNavigator struct {
	NotesStruct
}

func NewNotesViewNavigator(dispatchPtr *ole.IDispatch) NotesViewNavigator {
	return NotesViewNavigator{NewNotesStruct(dispatchPtr)}
}

/* --------------------------------- Properties --------------------------------- */
/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_CACHESIZE_PROPERTY_NOTESVIEWNAVIGATOR.html */
func (v NotesViewNavigator) CacheSize() (Long, error) {
	val, err := getComProperty(v, "CacheSize")
	return helpers.CastValue[Long](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_CACHESIZE_PROPERTY_NOTESVIEWNAVIGATOR.html */
func (v NotesViewNavigator) SetCacheSize(val Long) error {
	return putComProperty(v, "CacheSize", val)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_COUNT_PROPERTY_VIEWNAV.html */
func (v NotesViewNavigator) Count() (Long, error) {
	val, err := getComProperty(v, "Count")
	return helpers.CastValue[Long](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_MAXLEVEL_PROPERTY_VIEWNAV.html */
func (v NotesViewNavigator) MaxLevel() (Long, error) {
	val, err := getComProperty(v, "MaxLevel")
	return helpers.CastValue[Long](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_MAXLEVEL_PROPERTY_VIEWNAV.html */
func (v NotesViewNavigator) SetMaxLevel(val Long) error {
	return putComProperty(v, "MaxLevel", val)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_PARENTVIEW_PROPERTY_VIEWNAV.html */
func (v NotesViewNavigator) ParentView() (NotesView, error) {
	dispatchPtr, err := getComObjectProperty(v, "ParentView")
	return NewNotesView(dispatchPtr), err
}

/* --------------------------------- Methods ------------------------------------ */
/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GETCHILD_METHOD_VIEWNAV.html */
func (v NotesViewNavigator) GetChild(entry NotesViewEntry) (NotesViewEntry, error) {
	dispatchPtr, err := callComObjectMethod(v, "GetChild", entry.com().Dispatch())
	return NewNotesViewEntry(dispatchPtr), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GETCURRENT_METHOD_VIEWNAV.html */
func (v NotesViewNavigator) GetCurrent() (NotesViewEntry, error) {
	dispatchPtr, err := callComObjectMethod(v, "GetCurrent")
	return NewNotesViewEntry(dispatchPtr), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GETENTRY_METHOD_VIEWNAV.html */
func (v NotesViewNavigator) GetEntry(entry notesStruct) (NotesViewEntry, error) {
	dispatchPtr, err := callComObjectMethod(v, "GetEntry", entry.com().Dispatch())
	return NewNotesViewEntry(dispatchPtr), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GETFIRST_METHOD_VIEWNAV.html */
func (v NotesViewNavigator) GetFirst() (NotesViewEntry, error) {
	dispatchPtr, err := callComObjectMethod(v, "GetFirst")
	return NewNotesViewEntry(dispatchPtr), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GETFIRSTDOCUMENT_METHOD_VIEWNAV.html */
func (v NotesViewNavigator) GetFirstDocument() (NotesViewEntry, error) {
	dispatchPtr, err := callComObjectMethod(v, "GetFirstDocument")
	return NewNotesViewEntry(dispatchPtr), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GETLAST_METHOD_VIEWNAV.html */
func (v NotesViewNavigator) GetLast() (NotesViewEntry, error) {
	dispatchPtr, err := callComObjectMethod(v, "GetLast")
	return NewNotesViewEntry(dispatchPtr), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GETLASTDOCUMENT_METHOD_VIEWNAV.html */
func (v NotesViewNavigator) GetLastDocument() (NotesViewEntry, error) {
	dispatchPtr, err := callComObjectMethod(v, "GetLastDocument")
	return NewNotesViewEntry(dispatchPtr), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GETNEXT_METHOD_VIEWNAV.html */
func (v NotesViewNavigator) GetNext(entry NotesViewEntry) (NotesViewEntry, error) {
	dispatchPtr, err := callComObjectMethod(v, "GetNext", entry.com().Dispatch())
	return NewNotesViewEntry(dispatchPtr), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GETNEXTCATEGORY_METHOD_VIEWNAV.html */
func (v NotesViewNavigator) GetNextCategory(entry NotesViewEntry) (NotesViewEntry, error) {
	dispatchPtr, err := callComObjectMethod(v, "GetNextCategory", entry.com().Dispatch())
	return NewNotesViewEntry(dispatchPtr), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GETNEXTDOCUMENT_METHOD_VIEWNAV.html */
func (v NotesViewNavigator) GetNextDocument(entry NotesViewEntry) (NotesViewEntry, error) {
	dispatchPtr, err := callComObjectMethod(v, "GetNextDocument", entry.com().Dispatch())
	return NewNotesViewEntry(dispatchPtr), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GETNEXTSIBLING_METHOD_VIEWNAV.html */
func (v NotesViewNavigator) GetNextSibling(entry NotesViewEntry) (NotesViewEntry, error) {
	dispatchPtr, err := callComObjectMethod(v, "GetNextSibling", entry.com().Dispatch())
	return NewNotesViewEntry(dispatchPtr), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GETNTH_METHOD_VIEWNAV.html */
func (v NotesViewNavigator) GetNth(index Long) (NotesViewEntry, error) {
	dispatchPtr, err := callComObjectMethod(v, "GetNth", index)
	return NewNotesViewEntry(dispatchPtr), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GETPARENT_METHOD_VIEWNAV.html */
func (v NotesViewNavigator) GetParent(entry NotesViewEntry) (NotesViewEntry, error) {
	dispatchPtr, err := callComObjectMethod(v, "GetParent", entry.com().Dispatch())
	return NewNotesViewEntry(dispatchPtr), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GETPOS_METHOD_VIEWNAV.html */
func (v NotesViewNavigator) GetPos(position String, separator String) (NotesViewEntry, error) {
	dispatchPtr, err := callComObjectMethod(v, "GetPos", position, separator)
	return NewNotesViewEntry(dispatchPtr), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GETPREV_METHOD_VIEWNAV.html */
func (v NotesViewNavigator) GetPrev(entry NotesViewEntry) (NotesViewEntry, error) {
	dispatchPtr, err := callComObjectMethod(v, "GetPrev", entry.com().Dispatch())
	return NewNotesViewEntry(dispatchPtr), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GETPREVCATEGORY_METHOD_VIEWNAV.html */
func (v NotesViewNavigator) GetPrevCategory(entry NotesViewEntry) (NotesViewEntry, error) {
	dispatchPtr, err := callComObjectMethod(v, "GetPrevCategory", entry.com().Dispatch())
	return NewNotesViewEntry(dispatchPtr), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GETPREVDOCUMENT_METHOD_VIEWNAV.html */
func (v NotesViewNavigator) GetPrevDocument(entry NotesViewEntry) (NotesViewEntry, error) {
	dispatchPtr, err := callComObjectMethod(v, "GetPrevDocument", entry.com().Dispatch())
	return NewNotesViewEntry(dispatchPtr), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GETPREVSIBLING_METHOD_VIEWNAV.html */
func (v NotesViewNavigator) GetPrevSibling(entry NotesViewEntry) (NotesViewEntry, error) {
	dispatchPtr, err := callComObjectMethod(v, "GetPrevSibling", entry.com().Dispatch())
	return NewNotesViewEntry(dispatchPtr), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GOTOCHILD_METHOD_NOTESVIEWNAV.html */
func (v NotesViewNavigator) GotoChild(entry NotesViewEntry) error {
	_, err := callComMethod(v, "GotoChild", entry.com().Dispatch())
	return err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GOTOENTRY_METHOD_VIEWNAV.html */
func (v NotesViewNavigator) GotoEntry(objUnknown notesStruct) error {
	_, err := callComMethod(v, "GotoEntry", objUnknown.com().Dispatch())
	return err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GOTOFIRST_METHOD_VIEWNAV.html */
func (v NotesViewNavigator) GotoFirst() error {
	_, err := callComMethod(v, "GotoFirst")
	return err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GOTOFIRSTDOCUMENT_METHOD_VIEWNAV.html */
func (v NotesViewNavigator) GotoFirstDocument() error {
	_, err := callComMethod(v, "GotoFirstDocument")
	return err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GOTOLAST_METHOD_VIEWNAV.html */
func (v NotesViewNavigator) GotoLast() error {
	_, err := callComMethod(v, "GotoLast")
	return err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GOTOLASTDOCUMENT_METHOD_VIEWNAV.html */
func (v NotesViewNavigator) GotoLastDocument() error {
	_, err := callComMethod(v, "GotoLastDocument")
	return err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GOTONEXT_METHOD_VIEWNAV.html */
func (v NotesViewNavigator) GotoNext(entry NotesViewEntry) error {
	_, err := callComMethod(v, "GotoNext", entry.com().Dispatch())
	return err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GOTONEXTCATEGORY_METHOD_VIEWNAV.html */
func (v NotesViewNavigator) GotoNextCategory(entry NotesViewEntry) error {
	_, err := callComMethod(v, "GotoNextCategory", entry.com().Dispatch())
	return err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GOTONEXTDOCUMENT_METHOD_VIEWNAV.html */
func (v NotesViewNavigator) GotoNextDocument(entry NotesViewEntry) error {
	_, err := callComMethod(v, "GotoNextDocument", entry.com().Dispatch())
	return err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GOTONEXTSIBLING_METHOD_VIEWNAV.html */
func (v NotesViewNavigator) GotoNextSibling(entry NotesViewEntry) error {
	_, err := callComMethod(v, "GotoNextSibling", entry.com().Dispatch())
	return err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GOTOPARENT_METHOD_VIEWNAV.html */
func (v NotesViewNavigator) GotoParent(entry NotesViewEntry) error {
	_, err := callComMethod(v, "GotoParent", entry.com().Dispatch())
	return err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GOTOPOS_METHOD_VIEWNAV.html */
func (v NotesViewNavigator) GotoPos(pos String, separator String) error {
	_, err := callComMethod(v, "GotoPos", pos, separator)
	return err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GOTOPREV_METHOD_VIEWNAV.html */
func (v NotesViewNavigator) GotoPrev(entry NotesViewEntry) error {
	_, err := callComMethod(v, "GotoPrev", entry.com().Dispatch())
	return err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GOTOPREVCATEGORY_METHOD_VIEWNAV.html */
func (v NotesViewNavigator) GotoPrevCategory(entry NotesViewEntry) error {
	_, err := callComMethod(v, "GotoPrevCategory", entry.com().Dispatch())
	return err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GOTOPREVDOCUMENT_METHOD_VIEWNAV.html */
func (v NotesViewNavigator) GotoPrevDocument(entry NotesViewEntry) error {
	_, err := callComMethod(v, "GotoPrevDocument", entry.com().Dispatch())
	return err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GOTOPREVSIBLING_METHOD_VIEWNAV.html */
func (v NotesViewNavigator) GotoPrevSibling(entry NotesViewEntry) error {
	_, err := callComMethod(v, "GotoPrevSibling", entry.com().Dispatch())
	return err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_MARKALLREAD_VIEWNAV.html */
type notesViewNavigatorMarkAllReadParams struct {
	username *String
}

type notesViewNavigatorMarkAllReadParam func(*notesViewNavigatorMarkAllReadParams)

func WithNotesViewNavigatorMarkAllReadUsername(username String) notesViewNavigatorMarkAllReadParam {
	return func(c *notesViewNavigatorMarkAllReadParams) {
		c.username = &username
	}
}

func (v NotesViewNavigator) MarkAllRead(params ...notesViewNavigatorMarkAllReadParam) error {
	paramsStruct := &notesViewNavigatorMarkAllReadParams{}
	paramsOrdered := []interface{}{}

	for _, p := range params {
		p(paramsStruct)
	}

	if paramsStruct.username != nil {
		paramsOrdered = append(paramsOrdered, *paramsStruct.username)
	}
	_, err := callComMethod(v, "MarkAllRead", paramsOrdered...)
	return err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_MARKALLUNREAD_VIEWNAV.html */
type notesViewNavigatorMarkAllUnreadParams struct {
	username *String
}

type notesViewNavigatorMarkAllUnreadParam func(*notesViewNavigatorMarkAllUnreadParams)

func WithNotesViewNavigatorMarkAllUnreadUsername(username String) notesViewNavigatorMarkAllUnreadParam {
	return func(c *notesViewNavigatorMarkAllUnreadParams) {
		c.username = &username
	}
}

func (v NotesViewNavigator) MarkAllUnread(params ...notesViewNavigatorMarkAllUnreadParam) error {
	paramsStruct := &notesViewNavigatorMarkAllUnreadParams{}
	paramsOrdered := []interface{}{}

	for _, p := range params {
		p(paramsStruct)
	}

	if paramsStruct.username != nil {
		paramsOrdered = append(paramsOrdered, *paramsStruct.username)
	}
	_, err := callComMethod(v, "MarkAllUnread", paramsOrdered...)
	return err
}
