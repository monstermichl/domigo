/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_NOTESVIEWNAVIGATOR_CLASS.html */
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
type createViewNavParams struct {
	cacheSize *domino.Long
}

type createViewNavParam func(*createViewNavParams)

func WithCreateViewNavCacheSize(cacheSize domino.Long) createViewNavParam {
	return func(c *createViewNavParams) {
		c.cacheSize = &cacheSize
	}
}

func NotesViewCreateViewNav(v notesview.NotesView, params ...createViewNavParam) (NotesViewNavigator, error) {
	paramsStruct := &createViewNavParams{}
	paramsOrdered := []interface{}{}

	for _, p := range params {
		p(paramsStruct)
	}

	if paramsStruct.cacheSize != nil {
		paramsOrdered = append(paramsOrdered, *paramsStruct.cacheSize)
	}
	dispatchPtr, err := v.Com().CallObjectMethod("CreateViewNav", paramsOrdered...)
	return New(dispatchPtr), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_CREATEVIEWNAVFROM_METHOD_5742.html */
/* Moved from NotesView. */
type createViewNavFromParams struct {
	cacheSize *domino.Long
}

type createViewNavFromParam func(*createViewNavFromParams)

func WithCreateViewNavFromCacheSize(cacheSize domino.Long) createViewNavFromParam {
	return func(c *createViewNavFromParams) {
		c.cacheSize = &cacheSize
	}
}

func NotesViewCreateViewNavFrom(v notesview.NotesView, navigatorObject domino.NotesConnector, params ...createViewNavFromParam) (NotesViewNavigator, error) {
	paramsStruct := &createViewNavFromParams{}
	paramsOrdered := []interface{}{navigatorObject.Com().Dispatch()}

	for _, p := range params {
		p(paramsStruct)
	}

	if paramsStruct.cacheSize != nil {
		paramsOrdered = append(paramsOrdered, *paramsStruct.cacheSize)
	}
	dispatchPtr, err := v.Com().CallObjectMethod("CreateViewNavFrom", paramsOrdered...)
	return New(dispatchPtr), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_CREATEVIEWNAVFROMALLUNREAD.html */
/* Moved from NotesView. */
type createViewNavFromAllUnreadParams struct {
	username *domino.String
}

type createViewNavFromAllUnreadParam func(*createViewNavFromAllUnreadParams)

func WithCreateViewNavFromAllUnreadUsername(username domino.String) createViewNavFromAllUnreadParam {
	return func(c *createViewNavFromAllUnreadParams) {
		c.username = &username
	}
}

func NotesViewCreateViewNavFromAllUnread(v notesview.NotesView, params ...createViewNavFromAllUnreadParam) (NotesViewNavigator, error) {
	paramsStruct := &createViewNavFromAllUnreadParams{}
	paramsOrdered := []interface{}{}

	for _, p := range params {
		p(paramsStruct)
	}

	if paramsStruct.username != nil {
		paramsOrdered = append(paramsOrdered, *paramsStruct.username)
	}
	dispatchPtr, err := v.Com().CallObjectMethod("CreateViewNavFromAllUnread", paramsOrdered...)
	return New(dispatchPtr), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_CREATEVIEWNAVFROMCATEGORY_METHOD_3595.html */
/* Moved from NotesView. */
type createViewNavFromCategoryParams struct {
	cacheSize *domino.Long
}

type createViewNavFromCategoryParam func(*createViewNavFromCategoryParams)

func WithCreateViewNavFromCategoryCacheSize(cacheSize domino.Long) createViewNavFromCategoryParam {
	return func(c *createViewNavFromCategoryParams) {
		c.cacheSize = &cacheSize
	}
}

func NotesViewCreateViewNavFromCategory(v notesview.NotesView, category domino.String, params ...createViewNavFromCategoryParam) (NotesViewNavigator, error) {
	paramsStruct := &createViewNavFromCategoryParams{}
	paramsOrdered := []interface{}{category}

	for _, p := range params {
		p(paramsStruct)
	}

	if paramsStruct.cacheSize != nil {
		paramsOrdered = append(paramsOrdered, *paramsStruct.cacheSize)
	}
	dispatchPtr, err := v.Com().CallObjectMethod("CreateViewNavFromCategory", paramsOrdered...)
	return New(dispatchPtr), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_CREATEVIEWNAVFROMCHILDREN_METHOD_9100.html */
/* Moved from NotesView. */
type createViewNavFromChildrenParams struct {
	cacheSize *domino.Long
}

type createViewNavFromChildrenParam func(*createViewNavFromChildrenParams)

func WithCreateViewNavFromChildrenCacheSize(cacheSize domino.Long) createViewNavFromChildrenParam {
	return func(c *createViewNavFromChildrenParams) {
		c.cacheSize = &cacheSize
	}
}

func NotesViewCreateViewNavFromChildren(v notesview.NotesView, navigatorObject domino.NotesConnector, params ...createViewNavFromChildrenParam) (NotesViewNavigator, error) {
	paramsStruct := &createViewNavFromChildrenParams{}
	paramsOrdered := []interface{}{navigatorObject.Com().Dispatch()}

	for _, p := range params {
		p(paramsStruct)
	}

	if paramsStruct.cacheSize != nil {
		paramsOrdered = append(paramsOrdered, *paramsStruct.cacheSize)
	}
	dispatchPtr, err := v.Com().CallObjectMethod("CreateViewNavFromChildren", paramsOrdered...)
	return New(dispatchPtr), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_CREATEVIEWNAVFROMDESCENDANTS_METHOD_2893.html */
/* Moved from NotesView. */
type createViewNavFromDescendantsParams struct {
	cacheSize *domino.Long
}

type createViewNavFromDescendantsParam func(*createViewNavFromDescendantsParams)

func WithCreateViewNavFromDescendantsCacheSize(cacheSize domino.Long) createViewNavFromDescendantsParam {
	return func(c *createViewNavFromDescendantsParams) {
		c.cacheSize = &cacheSize
	}
}

func NotesViewCreateViewNavFromDescendants(v notesview.NotesView, navigatorObject domino.NotesConnector, params ...createViewNavFromDescendantsParam) (NotesViewNavigator, error) {
	paramsStruct := &createViewNavFromDescendantsParams{}
	paramsOrdered := []interface{}{navigatorObject.Com().Dispatch()}

	for _, p := range params {
		p(paramsStruct)
	}

	if paramsStruct.cacheSize != nil {
		paramsOrdered = append(paramsOrdered, *paramsStruct.cacheSize)
	}
	dispatchPtr, err := v.Com().CallObjectMethod("CreateViewNavFromDescendants", paramsOrdered...)
	return New(dispatchPtr), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_CREATEVIEWNAVMAXLEVEL_METHOD_NOTESVIEW_CLASS.html */
/* Moved from NotesView. */
type createViewNavMaxLevelParams struct {
	cacheSize *domino.Long
}

type createViewNavMaxLevelParam func(*createViewNavMaxLevelParams)

func WithCreateViewNavMaxLevelCacheSize(cacheSize domino.Long) createViewNavMaxLevelParam {
	return func(c *createViewNavMaxLevelParams) {
		c.cacheSize = &cacheSize
	}
}

func NotesViewCreateViewNavMaxLevel(v notesview.NotesView, level domino.Long, params ...createViewNavMaxLevelParam) (NotesViewNavigator, error) {
	paramsStruct := &createViewNavMaxLevelParams{}
	paramsOrdered := []interface{}{level}

	for _, p := range params {
		p(paramsStruct)
	}

	if paramsStruct.cacheSize != nil {
		paramsOrdered = append(paramsOrdered, *paramsStruct.cacheSize)
	}
	dispatchPtr, err := v.Com().CallObjectMethod("CreateViewNavMaxLevel", paramsOrdered...)
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
func (v NotesViewNavigator) GetEntry(entry domino.NotesConnector) (notesviewentry.NotesViewEntry, error) {
	dispatchPtr, err := v.Com().CallObjectMethod("GetEntry", entry.Com().Dispatch())
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
func (v NotesViewNavigator) GotoEntry(objUnknown domino.NotesConnector) error {
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
type markAllReadParams struct {
	username *domino.String
}

type markAllReadParam func(*markAllReadParams)

func WithMarkAllReadUsername(username domino.String) markAllReadParam {
	return func(c *markAllReadParams) {
		c.username = &username
	}
}

func (v NotesViewNavigator) MarkAllRead(params ...markAllReadParam) error {
	paramsStruct := &markAllReadParams{}
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
type markAllUnreadParams struct {
	username *domino.String
}

type markAllUnreadParam func(*markAllUnreadParams)

func WithMarkAllUnreadUsername(username domino.String) markAllUnreadParam {
	return func(c *markAllUnreadParams) {
		c.username = &username
	}
}

func (v NotesViewNavigator) MarkAllUnread(params ...markAllUnreadParam) error {
	paramsStruct := &markAllUnreadParams{}
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
