/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_NOTESVIEWNAVIGATOR_CLASS.html */
package domigo

import (
	"github.com/monstermichl/domigo/helpers"

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
	val, err := v.Com().GetProperty("CacheSize")
	return helpers.CastValue[Long](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_CACHESIZE_PROPERTY_NOTESVIEWNAVIGATOR.html */
func (v NotesViewNavigator) SetCacheSize(val Long) error {
	return v.Com().PutProperty("CacheSize", val)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_COUNT_PROPERTY_VIEWNAV.html */
func (v NotesViewNavigator) Count() (Long, error) {
	val, err := v.Com().GetProperty("Count")
	return helpers.CastValue[Long](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_MAXLEVEL_PROPERTY_VIEWNAV.html */
func (v NotesViewNavigator) MaxLevel() (Long, error) {
	val, err := v.Com().GetProperty("MaxLevel")
	return helpers.CastValue[Long](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_MAXLEVEL_PROPERTY_VIEWNAV.html */
func (v NotesViewNavigator) SetMaxLevel(val Long) error {
	return v.Com().PutProperty("MaxLevel", val)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_PARENTVIEW_PROPERTY_VIEWNAV.html */
func (v NotesViewNavigator) ParentView() (NotesView, error) {
	dispatchPtr, err := v.Com().GetObjectProperty("ParentView")
	return NewNotesView(dispatchPtr), err
}

/* --------------------------------- Methods ------------------------------------ */
/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GETCHILD_METHOD_VIEWNAV.html */
func (v NotesViewNavigator) GetChild(entry NotesViewEntry) (NotesViewEntry, error) {
	dispatchPtr, err := v.Com().CallObjectMethod("GetChild", entry.Com().Dispatch())
	return NewNotesViewEntry(dispatchPtr), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GETCURRENT_METHOD_VIEWNAV.html */
func (v NotesViewNavigator) GetCurrent() (NotesViewEntry, error) {
	dispatchPtr, err := v.Com().CallObjectMethod("GetCurrent")
	return NewNotesViewEntry(dispatchPtr), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GETENTRY_METHOD_VIEWNAV.html */
func (v NotesViewNavigator) GetEntry(entry NotesConnector) (NotesViewEntry, error) {
	dispatchPtr, err := v.Com().CallObjectMethod("GetEntry", entry.Com().Dispatch())
	return NewNotesViewEntry(dispatchPtr), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GETFIRST_METHOD_VIEWNAV.html */
func (v NotesViewNavigator) GetFirst() (NotesViewEntry, error) {
	dispatchPtr, err := v.Com().CallObjectMethod("GetFirst")
	return NewNotesViewEntry(dispatchPtr), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GETFIRSTDOCUMENT_METHOD_VIEWNAV.html */
func (v NotesViewNavigator) GetFirstDocument() (NotesViewEntry, error) {
	dispatchPtr, err := v.Com().CallObjectMethod("GetFirstDocument")
	return NewNotesViewEntry(dispatchPtr), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GETLAST_METHOD_VIEWNAV.html */
func (v NotesViewNavigator) GetLast() (NotesViewEntry, error) {
	dispatchPtr, err := v.Com().CallObjectMethod("GetLast")
	return NewNotesViewEntry(dispatchPtr), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GETLASTDOCUMENT_METHOD_VIEWNAV.html */
func (v NotesViewNavigator) GetLastDocument() (NotesViewEntry, error) {
	dispatchPtr, err := v.Com().CallObjectMethod("GetLastDocument")
	return NewNotesViewEntry(dispatchPtr), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GETNEXT_METHOD_VIEWNAV.html */
func (v NotesViewNavigator) GetNext(entry NotesViewEntry) (NotesViewEntry, error) {
	dispatchPtr, err := v.Com().CallObjectMethod("GetNext", entry.Com().Dispatch())
	return NewNotesViewEntry(dispatchPtr), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GETNEXTCATEGORY_METHOD_VIEWNAV.html */
func (v NotesViewNavigator) GetNextCategory(entry NotesViewEntry) (NotesViewEntry, error) {
	dispatchPtr, err := v.Com().CallObjectMethod("GetNextCategory", entry.Com().Dispatch())
	return NewNotesViewEntry(dispatchPtr), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GETNEXTDOCUMENT_METHOD_VIEWNAV.html */
func (v NotesViewNavigator) GetNextDocument(entry NotesViewEntry) (NotesViewEntry, error) {
	dispatchPtr, err := v.Com().CallObjectMethod("GetNextDocument", entry.Com().Dispatch())
	return NewNotesViewEntry(dispatchPtr), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GETNEXTSIBLING_METHOD_VIEWNAV.html */
func (v NotesViewNavigator) GetNextSibling(entry NotesViewEntry) (NotesViewEntry, error) {
	dispatchPtr, err := v.Com().CallObjectMethod("GetNextSibling", entry.Com().Dispatch())
	return NewNotesViewEntry(dispatchPtr), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GETNTH_METHOD_VIEWNAV.html */
func (v NotesViewNavigator) GetNth(index Long) (NotesViewEntry, error) {
	dispatchPtr, err := v.Com().CallObjectMethod("GetNth", index)
	return NewNotesViewEntry(dispatchPtr), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GETPARENT_METHOD_VIEWNAV.html */
func (v NotesViewNavigator) GetParent(entry NotesViewEntry) (NotesViewEntry, error) {
	dispatchPtr, err := v.Com().CallObjectMethod("GetParent", entry.Com().Dispatch())
	return NewNotesViewEntry(dispatchPtr), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GETPOS_METHOD_VIEWNAV.html */
func (v NotesViewNavigator) GetPos(position String, separator String) (NotesViewEntry, error) {
	dispatchPtr, err := v.Com().CallObjectMethod("GetPos", position, separator)
	return NewNotesViewEntry(dispatchPtr), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GETPREV_METHOD_VIEWNAV.html */
func (v NotesViewNavigator) GetPrev(entry NotesViewEntry) (NotesViewEntry, error) {
	dispatchPtr, err := v.Com().CallObjectMethod("GetPrev", entry.Com().Dispatch())
	return NewNotesViewEntry(dispatchPtr), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GETPREVCATEGORY_METHOD_VIEWNAV.html */
func (v NotesViewNavigator) GetPrevCategory(entry NotesViewEntry) (NotesViewEntry, error) {
	dispatchPtr, err := v.Com().CallObjectMethod("GetPrevCategory", entry.Com().Dispatch())
	return NewNotesViewEntry(dispatchPtr), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GETPREVDOCUMENT_METHOD_VIEWNAV.html */
func (v NotesViewNavigator) GetPrevDocument(entry NotesViewEntry) (NotesViewEntry, error) {
	dispatchPtr, err := v.Com().CallObjectMethod("GetPrevDocument", entry.Com().Dispatch())
	return NewNotesViewEntry(dispatchPtr), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GETPREVSIBLING_METHOD_VIEWNAV.html */
func (v NotesViewNavigator) GetPrevSibling(entry NotesViewEntry) (NotesViewEntry, error) {
	dispatchPtr, err := v.Com().CallObjectMethod("GetPrevSibling", entry.Com().Dispatch())
	return NewNotesViewEntry(dispatchPtr), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GOTOCHILD_METHOD_NOTESVIEWNAV.html */
func (v NotesViewNavigator) GotoChild(entry NotesViewEntry) error {
	_, err := v.Com().CallMethod("GotoChild", entry.Com().Dispatch())
	return err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GOTOENTRY_METHOD_VIEWNAV.html */
func (v NotesViewNavigator) GotoEntry(objUnknown NotesConnector) error {
	_, err := v.Com().CallMethod("GotoEntry", objUnknown.Com().Dispatch())
	return err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GOTOFIRST_METHOD_VIEWNAV.html */
func (v NotesViewNavigator) GotoFirst() error {
	_, err := v.Com().CallMethod("GotoFirst")
	return err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GOTOFIRSTDOCUMENT_METHOD_VIEWNAV.html */
func (v NotesViewNavigator) GotoFirstDocument() error {
	_, err := v.Com().CallMethod("GotoFirstDocument")
	return err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GOTOLAST_METHOD_VIEWNAV.html */
func (v NotesViewNavigator) GotoLast() error {
	_, err := v.Com().CallMethod("GotoLast")
	return err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GOTOLASTDOCUMENT_METHOD_VIEWNAV.html */
func (v NotesViewNavigator) GotoLastDocument() error {
	_, err := v.Com().CallMethod("GotoLastDocument")
	return err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GOTONEXT_METHOD_VIEWNAV.html */
func (v NotesViewNavigator) GotoNext(entry NotesViewEntry) error {
	_, err := v.Com().CallMethod("GotoNext", entry.Com().Dispatch())
	return err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GOTONEXTCATEGORY_METHOD_VIEWNAV.html */
func (v NotesViewNavigator) GotoNextCategory(entry NotesViewEntry) error {
	_, err := v.Com().CallMethod("GotoNextCategory", entry.Com().Dispatch())
	return err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GOTONEXTDOCUMENT_METHOD_VIEWNAV.html */
func (v NotesViewNavigator) GotoNextDocument(entry NotesViewEntry) error {
	_, err := v.Com().CallMethod("GotoNextDocument", entry.Com().Dispatch())
	return err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GOTONEXTSIBLING_METHOD_VIEWNAV.html */
func (v NotesViewNavigator) GotoNextSibling(entry NotesViewEntry) error {
	_, err := v.Com().CallMethod("GotoNextSibling", entry.Com().Dispatch())
	return err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GOTOPARENT_METHOD_VIEWNAV.html */
func (v NotesViewNavigator) GotoParent(entry NotesViewEntry) error {
	_, err := v.Com().CallMethod("GotoParent", entry.Com().Dispatch())
	return err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GOTOPOS_METHOD_VIEWNAV.html */
func (v NotesViewNavigator) GotoPos(pos String, separator String) error {
	_, err := v.Com().CallMethod("GotoPos", pos, separator)
	return err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GOTOPREV_METHOD_VIEWNAV.html */
func (v NotesViewNavigator) GotoPrev(entry NotesViewEntry) error {
	_, err := v.Com().CallMethod("GotoPrev", entry.Com().Dispatch())
	return err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GOTOPREVCATEGORY_METHOD_VIEWNAV.html */
func (v NotesViewNavigator) GotoPrevCategory(entry NotesViewEntry) error {
	_, err := v.Com().CallMethod("GotoPrevCategory", entry.Com().Dispatch())
	return err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GOTOPREVDOCUMENT_METHOD_VIEWNAV.html */
func (v NotesViewNavigator) GotoPrevDocument(entry NotesViewEntry) error {
	_, err := v.Com().CallMethod("GotoPrevDocument", entry.Com().Dispatch())
	return err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GOTOPREVSIBLING_METHOD_VIEWNAV.html */
func (v NotesViewNavigator) GotoPrevSibling(entry NotesViewEntry) error {
	_, err := v.Com().CallMethod("GotoPrevSibling", entry.Com().Dispatch())
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
	_, err := v.Com().CallMethod("MarkAllRead", paramsOrdered...)
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
	_, err := v.Com().CallMethod("MarkAllUnread", paramsOrdered...)
	return err
}
