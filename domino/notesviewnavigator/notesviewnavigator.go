package notesviewnavigator

import (
	"domigo/domino"
	"domigo/domino/notesview"
	"domigo/domino/notesviewentry"

	"domigo/helpers"

	ole "github.com/go-ole/go-ole"
)

type NotesViewNavigator struct {
	domino.NotesStruct
}

func New(dispatchPtr *ole.IDispatch) NotesViewNavigator {
	return NotesViewNavigator{domino.NewNotesStruct(dispatchPtr)}
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_CREATEVIEWNAV_METHOD_1631.html */
/* Moved from NotesView. */
func NotesViewCreateViewNav(v notesview.NotesView, cacheSize domino.Long) (NotesViewNavigator, error) {
	dispatchPtr, err := v.Com().CallObjectMethod("CreateViewNav", cacheSize)
	return New(dispatchPtr), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_CREATEVIEWNAVFROM_METHOD_5742.html */
/* Moved from NotesView. */
/* TODO: Handle Variant types. */
func NotesViewCreateViewNavFrom(v notesview.NotesView, navigatorObject domino.Variant, cacheSize domino.Long) (NotesViewNavigator, error) {
	dispatchPtr, err := v.Com().CallObjectMethod("CreateViewNavFrom", navigatorObject, cacheSize)
	return New(dispatchPtr), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_CREATEVIEWNAVFROMALLUNREAD.html */
/* Moved from NotesView. */
func NotesViewCreateViewNavFromAllUnread(v notesview.NotesView, username domino.String) (NotesViewNavigator, error) {
	dispatchPtr, err := v.Com().CallObjectMethod("CreateViewNavFromAllUnread", username)
	return New(dispatchPtr), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_CREATEVIEWNAVFROMCATEGORY_METHOD_3595.html */
/* Moved from NotesView. */
func NotesViewCreateViewNavFromCategory(v notesview.NotesView, category domino.String, cacheSize domino.Long) (NotesViewNavigator, error) {
	dispatchPtr, err := v.Com().CallObjectMethod("CreateViewNavFromCategory", category, cacheSize)
	return New(dispatchPtr), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_CREATEVIEWNAVFROMCHILDREN_METHOD_9100.html */
/* Moved from NotesView. */
/* TODO: Handle Variant types. */
func NotesViewCreateViewNavFromChildren(v notesview.NotesView, navigatorObject domino.Variant, cacheSize domino.Long) (NotesViewNavigator, error) {
	dispatchPtr, err := v.Com().CallObjectMethod("CreateViewNavFromChildren", navigatorObject, cacheSize)
	return New(dispatchPtr), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_CREATEVIEWNAVFROMDESCENDANTS_METHOD_2893.html */
/* Moved from NotesView. */
/* TODO: Handle Variant types. */
func NotesViewCreateViewNavFromDescendants(v notesview.NotesView, navigatorObject domino.Variant, cacheSize domino.Long) (NotesViewNavigator, error) {
	dispatchPtr, err := v.Com().CallObjectMethod("CreateViewNavFromDescendants", navigatorObject, cacheSize)
	return New(dispatchPtr), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_CREATEVIEWNAVMAXLEVEL_METHOD_NOTESVIEW_CLASS.html */
/* Moved from NotesView. */
func NotesViewCreateViewNavMaxLevel(v notesview.NotesView, level domino.Long, cacheSize domino.Long) (NotesViewNavigator, error) {
	dispatchPtr, err := v.Com().CallObjectMethod("CreateViewNavMaxLevel", level, cacheSize)
	return New(dispatchPtr), err
}

/* --------------------------------- Properties --------------------------------- */
/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_CACHESIZE_PROPERTY_NOTESVIEWNAVIGATOR.html */
func (v NotesViewNavigator) CacheSize() (domino.Long, error) {
	val, err := v.Com().GetProperty("CacheSize")
	return helpers.CastValue[domino.Long](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_CACHESIZE_PROPERTY_NOTESVIEWNAVIGATOR.html */
func (v NotesViewNavigator) SetCacheSize(val domino.Long) error {
	return v.Com().PutProperty("CacheSize", val)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_COUNT_PROPERTY_VIEWNAV.html */
func (v NotesViewNavigator) Count() (domino.Long, error) {
	val, err := v.Com().GetProperty("Count")
	return helpers.CastValue[domino.Long](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_MAXLEVEL_PROPERTY_VIEWNAV.html */
func (v NotesViewNavigator) MaxLevel() (domino.Long, error) {
	val, err := v.Com().GetProperty("MaxLevel")
	return helpers.CastValue[domino.Long](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_MAXLEVEL_PROPERTY_VIEWNAV.html */
func (v NotesViewNavigator) SetMaxLevel(val domino.Long) error {
	return v.Com().PutProperty("MaxLevel", val)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_PARENTVIEW_PROPERTY_VIEWNAV.html */
func (v NotesViewNavigator) ParentView() (notesview.NotesView, error) {
	dispatchPtr, err := v.Com().GetObjectProperty("ParentView")
	return notesview.New(dispatchPtr), err
}

/* --------------------------------- Methods ------------------------------------ */
/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GETCHILD_METHOD_VIEWNAV.html */
func (v NotesViewNavigator) GetChild(entry notesviewentry.NotesViewEntry) (notesviewentry.NotesViewEntry, error) {
	dispatchPtr, err := v.Com().CallObjectMethod("GetChild", entry.Com().Dispatch())
	return notesviewentry.New(dispatchPtr), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GETCURRENT_METHOD_VIEWNAV.html */
func (v NotesViewNavigator) GetCurrent() (notesviewentry.NotesViewEntry, error) {
	dispatchPtr, err := v.Com().CallObjectMethod("GetCurrent")
	return notesviewentry.New(dispatchPtr), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GETENTRY_METHOD_VIEWNAV.html */
func (v NotesViewNavigator) GetEntry(entry domino.Variant) (notesviewentry.NotesViewEntry, error) {
	dispatchPtr, err := v.Com().CallObjectMethod("GetEntry", entry)
	return notesviewentry.New(dispatchPtr), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GETFIRST_METHOD_VIEWNAV.html */
func (v NotesViewNavigator) GetFirst() (notesviewentry.NotesViewEntry, error) {
	dispatchPtr, err := v.Com().CallObjectMethod("GetFirst")
	return notesviewentry.New(dispatchPtr), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GETFIRSTDOCUMENT_METHOD_VIEWNAV.html */
func (v NotesViewNavigator) GetFirstDocument() (notesviewentry.NotesViewEntry, error) {
	dispatchPtr, err := v.Com().CallObjectMethod("GetFirstDocument")
	return notesviewentry.New(dispatchPtr), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GETLAST_METHOD_VIEWNAV.html */
func (v NotesViewNavigator) GetLast() (notesviewentry.NotesViewEntry, error) {
	dispatchPtr, err := v.Com().CallObjectMethod("GetLast")
	return notesviewentry.New(dispatchPtr), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GETLASTDOCUMENT_METHOD_VIEWNAV.html */
func (v NotesViewNavigator) GetLastDocument() (notesviewentry.NotesViewEntry, error) {
	dispatchPtr, err := v.Com().CallObjectMethod("GetLastDocument")
	return notesviewentry.New(dispatchPtr), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GETNEXT_METHOD_VIEWNAV.html */
func (v NotesViewNavigator) GetNext(entry notesviewentry.NotesViewEntry) (notesviewentry.NotesViewEntry, error) {
	dispatchPtr, err := v.Com().CallObjectMethod("GetNext", entry.Com().Dispatch())
	return notesviewentry.New(dispatchPtr), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GETNEXTCATEGORY_METHOD_VIEWNAV.html */
func (v NotesViewNavigator) GetNextCategory(entry notesviewentry.NotesViewEntry) (notesviewentry.NotesViewEntry, error) {
	dispatchPtr, err := v.Com().CallObjectMethod("GetNextCategory", entry.Com().Dispatch())
	return notesviewentry.New(dispatchPtr), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GETNEXTDOCUMENT_METHOD_VIEWNAV.html */
func (v NotesViewNavigator) GetNextDocument(entry notesviewentry.NotesViewEntry) (notesviewentry.NotesViewEntry, error) {
	dispatchPtr, err := v.Com().CallObjectMethod("GetNextDocument", entry.Com().Dispatch())
	return notesviewentry.New(dispatchPtr), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GETNEXTSIBLING_METHOD_VIEWNAV.html */
func (v NotesViewNavigator) GetNextSibling(entry notesviewentry.NotesViewEntry) (notesviewentry.NotesViewEntry, error) {
	dispatchPtr, err := v.Com().CallObjectMethod("GetNextSibling", entry.Com().Dispatch())
	return notesviewentry.New(dispatchPtr), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GETNTH_METHOD_VIEWNAV.html */
func (v NotesViewNavigator) GetNth(index domino.Long) (notesviewentry.NotesViewEntry, error) {
	dispatchPtr, err := v.Com().CallObjectMethod("GetNth", index)
	return notesviewentry.New(dispatchPtr), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GETPARENT_METHOD_VIEWNAV.html */
func (v NotesViewNavigator) GetParent(entry notesviewentry.NotesViewEntry) (notesviewentry.NotesViewEntry, error) {
	dispatchPtr, err := v.Com().CallObjectMethod("GetParent", entry.Com().Dispatch())
	return notesviewentry.New(dispatchPtr), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GETPOS_METHOD_VIEWNAV.html */
func (v NotesViewNavigator) GetPos(position domino.String, separator domino.String) (notesviewentry.NotesViewEntry, error) {
	dispatchPtr, err := v.Com().CallObjectMethod("GetPos", position, separator)
	return notesviewentry.New(dispatchPtr), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GETPREV_METHOD_VIEWNAV.html */
func (v NotesViewNavigator) GetPrev(entry notesviewentry.NotesViewEntry) (notesviewentry.NotesViewEntry, error) {
	dispatchPtr, err := v.Com().CallObjectMethod("GetPrev", entry.Com().Dispatch())
	return notesviewentry.New(dispatchPtr), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GETPREVCATEGORY_METHOD_VIEWNAV.html */
func (v NotesViewNavigator) GetPrevCategory(entry notesviewentry.NotesViewEntry) (notesviewentry.NotesViewEntry, error) {
	dispatchPtr, err := v.Com().CallObjectMethod("GetPrevCategory", entry.Com().Dispatch())
	return notesviewentry.New(dispatchPtr), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GETPREVDOCUMENT_METHOD_VIEWNAV.html */
func (v NotesViewNavigator) GetPrevDocument(entry notesviewentry.NotesViewEntry) (notesviewentry.NotesViewEntry, error) {
	dispatchPtr, err := v.Com().CallObjectMethod("GetPrevDocument", entry.Com().Dispatch())
	return notesviewentry.New(dispatchPtr), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GETPREVSIBLING_METHOD_VIEWNAV.html */
func (v NotesViewNavigator) GetPrevSibling(entry notesviewentry.NotesViewEntry) (notesviewentry.NotesViewEntry, error) {
	dispatchPtr, err := v.Com().CallObjectMethod("GetPrevSibling", entry.Com().Dispatch())
	return notesviewentry.New(dispatchPtr), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GOTOCHILD_METHOD_NOTESVIEWNAV.html */
func (v NotesViewNavigator) GotoChild(entry notesviewentry.NotesViewEntry) error {
	_, err := v.Com().CallMethod("GotoChild", entry.Com().Dispatch())
	return err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GOTOENTRY_METHOD_VIEWNAV.html */
/* TODO: Find out what to pass to the function. */
// func (v NotesViewNavigator) GotoEntry(objUnknown domino.Object) error {
// 	_, err := v.Com().CallMethod("GotoEntry", objUnknown)
// 	return err
// }

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
func (v NotesViewNavigator) GotoNext(entry notesviewentry.NotesViewEntry) error {
	_, err := v.Com().CallMethod("GotoNext", entry.Com().Dispatch())
	return err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GOTONEXTCATEGORY_METHOD_VIEWNAV.html */
func (v NotesViewNavigator) GotoNextCategory(entry notesviewentry.NotesViewEntry) error {
	_, err := v.Com().CallMethod("GotoNextCategory", entry.Com().Dispatch())
	return err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GOTONEXTDOCUMENT_METHOD_VIEWNAV.html */
func (v NotesViewNavigator) GotoNextDocument(entry notesviewentry.NotesViewEntry) error {
	_, err := v.Com().CallMethod("GotoNextDocument", entry.Com().Dispatch())
	return err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GOTONEXTSIBLING_METHOD_VIEWNAV.html */
func (v NotesViewNavigator) GotoNextSibling(entry notesviewentry.NotesViewEntry) error {
	_, err := v.Com().CallMethod("GotoNextSibling", entry.Com().Dispatch())
	return err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GOTOPARENT_METHOD_VIEWNAV.html */
func (v NotesViewNavigator) GotoParent(entry notesviewentry.NotesViewEntry) error {
	_, err := v.Com().CallMethod("GotoParent", entry.Com().Dispatch())
	return err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GOTOPOS_METHOD_VIEWNAV.html */
func (v NotesViewNavigator) GotoPos(pos domino.String, separator domino.String) error {
	_, err := v.Com().CallMethod("GotoPos", pos, separator)
	return err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GOTOPREV_METHOD_VIEWNAV.html */
func (v NotesViewNavigator) GotoPrev(entry notesviewentry.NotesViewEntry) error {
	_, err := v.Com().CallMethod("GotoPrev", entry.Com().Dispatch())
	return err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GOTOPREVCATEGORY_METHOD_VIEWNAV.html */
func (v NotesViewNavigator) GotoPrevCategory(entry notesviewentry.NotesViewEntry) error {
	_, err := v.Com().CallMethod("GotoPrevCategory", entry.Com().Dispatch())
	return err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GOTOPREVDOCUMENT_METHOD_VIEWNAV.html */
func (v NotesViewNavigator) GotoPrevDocument(entry notesviewentry.NotesViewEntry) error {
	_, err := v.Com().CallMethod("GotoPrevDocument", entry.Com().Dispatch())
	return err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GOTOPREVSIBLING_METHOD_VIEWNAV.html */
func (v NotesViewNavigator) GotoPrevSibling(entry notesviewentry.NotesViewEntry) error {
	_, err := v.Com().CallMethod("GotoPrevSibling", entry.Com().Dispatch())
	return err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_MARKALLREAD_VIEWNAV.html */
/* TODO: Some parameters are optional. Make sure to handle them correctly. */
func (v NotesViewNavigator) MarkAllRead(username domino.String) error {
	_, err := v.Com().CallMethod("MarkAllRead", username)
	return err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_MARKALLUNREAD_VIEWNAV.html */
/* TODO: Some parameters are optional. Make sure to handle them correctly. */
func (v NotesViewNavigator) MarkAllUnread(username domino.String) error {
	_, err := v.Com().CallMethod("MarkAllUnread", username)
	return err
}
