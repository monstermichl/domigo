/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_NOTESVIEWNAVIGATOR_CLASS.html */
package domigo

import (
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
	return getComProperty[Long](v, "CacheSize")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_CACHESIZE_PROPERTY_NOTESVIEWNAVIGATOR.html */
func (v NotesViewNavigator) SetCacheSize(val Long) error {
	return putComProperty(v, "CacheSize", val)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_COUNT_PROPERTY_VIEWNAV.html */
func (v NotesViewNavigator) Count() (Long, error) {
	return getComProperty[Long](v, "Count")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_MAXLEVEL_PROPERTY_VIEWNAV.html */
func (v NotesViewNavigator) MaxLevel() (Long, error) {
	return getComProperty[Long](v, "MaxLevel")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_MAXLEVEL_PROPERTY_VIEWNAV.html */
func (v NotesViewNavigator) SetMaxLevel(val Long) error {
	return putComProperty(v, "MaxLevel", val)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_PARENTVIEW_PROPERTY_VIEWNAV.html */
func (v NotesViewNavigator) ParentView() (NotesView, error) {
	return getComObjectProperty(v, NewNotesView, "ParentView")
}

/* --------------------------------- Methods ------------------------------------ */
/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GETCHILD_METHOD_VIEWNAV.html */
func (v NotesViewNavigator) GetChild(entry NotesViewEntry) (NotesViewEntry, error) {
	return callComObjectMethod(v, NewNotesViewEntry, "GetChild", entry)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GETCURRENT_METHOD_VIEWNAV.html */
func (v NotesViewNavigator) GetCurrent() (NotesViewEntry, error) {
	return callComObjectMethod(v, NewNotesViewEntry, "GetCurrent")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GETENTRY_METHOD_VIEWNAV.html */
func (v NotesViewNavigator) GetEntry(entry notesStruct) (NotesViewEntry, error) {
	return callComObjectMethod(v, NewNotesViewEntry, "GetEntry", entry)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GETFIRST_METHOD_VIEWNAV.html */
func (v NotesViewNavigator) GetFirst() (NotesViewEntry, error) {
	return callComObjectMethod(v, NewNotesViewEntry, "GetFirst")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GETFIRSTDOCUMENT_METHOD_VIEWNAV.html */
func (v NotesViewNavigator) GetFirstDocument() (NotesViewEntry, error) {
	return callComObjectMethod(v, NewNotesViewEntry, "GetFirstDocument")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GETLAST_METHOD_VIEWNAV.html */
func (v NotesViewNavigator) GetLast() (NotesViewEntry, error) {
	return callComObjectMethod(v, NewNotesViewEntry, "GetLast")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GETLASTDOCUMENT_METHOD_VIEWNAV.html */
func (v NotesViewNavigator) GetLastDocument() (NotesViewEntry, error) {
	return callComObjectMethod(v, NewNotesViewEntry, "GetLastDocument")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GETNEXT_METHOD_VIEWNAV.html */
func (v NotesViewNavigator) GetNext(entry NotesViewEntry) (NotesViewEntry, error) {
	return callComObjectMethod(v, NewNotesViewEntry, "GetNext", entry)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GETNEXTCATEGORY_METHOD_VIEWNAV.html */
func (v NotesViewNavigator) GetNextCategory(entry NotesViewEntry) (NotesViewEntry, error) {
	return callComObjectMethod(v, NewNotesViewEntry, "GetNextCategory", entry)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GETNEXTDOCUMENT_METHOD_VIEWNAV.html */
func (v NotesViewNavigator) GetNextDocument(entry NotesViewEntry) (NotesViewEntry, error) {
	return callComObjectMethod(v, NewNotesViewEntry, "GetNextDocument", entry)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GETNEXTSIBLING_METHOD_VIEWNAV.html */
func (v NotesViewNavigator) GetNextSibling(entry NotesViewEntry) (NotesViewEntry, error) {
	return callComObjectMethod(v, NewNotesViewEntry, "GetNextSibling", entry)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GETNTH_METHOD_VIEWNAV.html */
func (v NotesViewNavigator) GetNth(index Long) (NotesViewEntry, error) {
	return callComObjectMethod(v, NewNotesViewEntry, "GetNth", index)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GETPARENT_METHOD_VIEWNAV.html */
func (v NotesViewNavigator) GetParent(entry NotesViewEntry) (NotesViewEntry, error) {
	return callComObjectMethod(v, NewNotesViewEntry, "GetParent", entry)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GETPOS_METHOD_VIEWNAV.html */
func (v NotesViewNavigator) GetPos(position String, separator String) (NotesViewEntry, error) {
	return callComObjectMethod(v, NewNotesViewEntry, "GetPos", position, separator)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GETPREV_METHOD_VIEWNAV.html */
func (v NotesViewNavigator) GetPrev(entry NotesViewEntry) (NotesViewEntry, error) {
	return callComObjectMethod(v, NewNotesViewEntry, "GetPrev", entry)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GETPREVCATEGORY_METHOD_VIEWNAV.html */
func (v NotesViewNavigator) GetPrevCategory(entry NotesViewEntry) (NotesViewEntry, error) {
	return callComObjectMethod(v, NewNotesViewEntry, "GetPrevCategory", entry)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GETPREVDOCUMENT_METHOD_VIEWNAV.html */
func (v NotesViewNavigator) GetPrevDocument(entry NotesViewEntry) (NotesViewEntry, error) {
	return callComObjectMethod(v, NewNotesViewEntry, "GetPrevDocument", entry)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GETPREVSIBLING_METHOD_VIEWNAV.html */
func (v NotesViewNavigator) GetPrevSibling(entry NotesViewEntry) (NotesViewEntry, error) {
	return callComObjectMethod(v, NewNotesViewEntry, "GetPrevSibling", entry)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GOTOCHILD_METHOD_NOTESVIEWNAV.html */
func (v NotesViewNavigator) GotoChild(entry NotesViewEntry) error {
	return callComVoidMethod(v, "GotoChild", entry)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GOTOENTRY_METHOD_VIEWNAV.html */
func (v NotesViewNavigator) GotoEntry(objUnknown notesStruct) error {
	return callComVoidMethod(v, "GotoEntry", objUnknown)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GOTOFIRST_METHOD_VIEWNAV.html */
func (v NotesViewNavigator) GotoFirst() error {
	return callComVoidMethod(v, "GotoFirst")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GOTOFIRSTDOCUMENT_METHOD_VIEWNAV.html */
func (v NotesViewNavigator) GotoFirstDocument() error {
	return callComVoidMethod(v, "GotoFirstDocument")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GOTOLAST_METHOD_VIEWNAV.html */
func (v NotesViewNavigator) GotoLast() error {
	return callComVoidMethod(v, "GotoLast")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GOTOLASTDOCUMENT_METHOD_VIEWNAV.html */
func (v NotesViewNavigator) GotoLastDocument() error {
	return callComVoidMethod(v, "GotoLastDocument")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GOTONEXT_METHOD_VIEWNAV.html */
func (v NotesViewNavigator) GotoNext(entry NotesViewEntry) error {
	return callComVoidMethod(v, "GotoNext", entry)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GOTONEXTCATEGORY_METHOD_VIEWNAV.html */
func (v NotesViewNavigator) GotoNextCategory(entry NotesViewEntry) error {
	return callComVoidMethod(v, "GotoNextCategory", entry)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GOTONEXTDOCUMENT_METHOD_VIEWNAV.html */
func (v NotesViewNavigator) GotoNextDocument(entry NotesViewEntry) error {
	return callComVoidMethod(v, "GotoNextDocument", entry)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GOTONEXTSIBLING_METHOD_VIEWNAV.html */
func (v NotesViewNavigator) GotoNextSibling(entry NotesViewEntry) error {
	return callComVoidMethod(v, "GotoNextSibling", entry)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GOTOPARENT_METHOD_VIEWNAV.html */
func (v NotesViewNavigator) GotoParent(entry NotesViewEntry) error {
	return callComVoidMethod(v, "GotoParent", entry)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GOTOPOS_METHOD_VIEWNAV.html */
func (v NotesViewNavigator) GotoPos(pos String, separator String) error {
	return callComVoidMethod(v, "GotoPos", pos, separator)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GOTOPREV_METHOD_VIEWNAV.html */
func (v NotesViewNavigator) GotoPrev(entry NotesViewEntry) error {
	return callComVoidMethod(v, "GotoPrev", entry)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GOTOPREVCATEGORY_METHOD_VIEWNAV.html */
func (v NotesViewNavigator) GotoPrevCategory(entry NotesViewEntry) error {
	return callComVoidMethod(v, "GotoPrevCategory", entry)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GOTOPREVDOCUMENT_METHOD_VIEWNAV.html */
func (v NotesViewNavigator) GotoPrevDocument(entry NotesViewEntry) error {
	return callComVoidMethod(v, "GotoPrevDocument", entry)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GOTOPREVSIBLING_METHOD_VIEWNAV.html */
func (v NotesViewNavigator) GotoPrevSibling(entry NotesViewEntry) error {
	return callComVoidMethod(v, "GotoPrevSibling", entry)
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
	return callComVoidMethod(v, "MarkAllRead", paramsOrdered...)
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
	return callComVoidMethod(v, "MarkAllUnread", paramsOrdered...)
}
