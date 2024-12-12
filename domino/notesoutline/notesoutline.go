/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_NOTESOUTLINE_CLASS.html */
package notesoutline

import (
	"domigo/domino"
	"domigo/domino/notesoutlineentry"
	"domigo/helpers"

	ole "github.com/go-ole/go-ole"
)

type NotesOutline struct {
	domino.NotesStruct
}

func New(dispatchPtr *ole.IDispatch) NotesOutline {
	return NotesOutline{domino.NewNotesStruct(dispatchPtr)}
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_PARENT_PROPERTY_OUTLINEENTRY_COM.html */
/* Moved from NotesOutlineEntry. */
func NotesOutlineEntryParent(o notesoutlineentry.NotesOutlineEntry) (NotesOutline, error) {
	dispatchPtr, err := o.Com().GetObjectProperty("Parent")
	return New(dispatchPtr), err
}

/* --------------------------------- Properties --------------------------------- */
/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ALIAS_PROPERTY_OUTLINE.html */
func (o NotesOutline) Alias() (domino.String, error) {
	val, err := o.Com().GetProperty("Alias")
	return helpers.CastValue[domino.String](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ALIAS_PROPERTY_OUTLINE.html */
func (o NotesOutline) SetAlias(v domino.String) error {
	return o.Com().PutProperty("Alias", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_COMMENT_PROPERTY_OUTLINE.html */
func (o NotesOutline) Comment() (domino.String, error) {
	val, err := o.Com().GetProperty("Comment")
	return helpers.CastValue[domino.String](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_COMMENT_PROPERTY_OUTLINE.html */
func (o NotesOutline) SetComment(v domino.String) error {
	return o.Com().PutProperty("Comment", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_NAME_PROPERTY_OUTLINE.html */
func (o NotesOutline) Name() (domino.String, error) {
	val, err := o.Com().GetProperty("Name")
	return helpers.CastValue[domino.String](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_NAME_PROPERTY_OUTLINE.html */
func (o NotesOutline) SetName(v domino.String) error {
	return o.Com().PutProperty("Name", v)
}

/* --------------------------------- Methods ------------------------------------ */

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_CREATEENTRY_METHOD_OUTLINE.html */
type createEntryParams struct {
	refentry *notesoutlineentry.NotesOutlineEntry
	addAfter *domino.Boolean
	asChild  *domino.Boolean
}

type createEntryParam func(*createEntryParams)

func WithCreateEntryRefentry(refentry notesoutlineentry.NotesOutlineEntry) createEntryParam {
	return func(c *createEntryParams) {
		c.refentry = &refentry
	}
}

func WithCreateEntryAddAfter(addAfter domino.Boolean) createEntryParam {
	return func(c *createEntryParams) {
		c.addAfter = &addAfter
	}
}

func WithCreateEntryAsChild(asChild domino.Boolean) createEntryParam {
	return func(c *createEntryParams) {
		c.asChild = &asChild
	}
}

func (o NotesOutline) CreateEntry(name domino.String, params ...createEntryParam) (notesoutlineentry.NotesOutlineEntry, error) {
	paramsStruct := &createEntryParams{}
	paramsOrdered := []interface{}{name}

	for _, p := range params {
		p(paramsStruct)
	}

	/* Note: This parameter is required in COM. Specify "Nothing"
	   to get the default (check documentation for details). */
	if paramsStruct.refentry == nil {
		paramsOrdered = append(paramsOrdered, nil)
	} else {
		paramsOrdered = append(paramsOrdered, paramsStruct.refentry.Com().Dispatch())
	}

	if paramsStruct.addAfter != nil {
		paramsOrdered = append(paramsOrdered, *paramsStruct.addAfter)
		if paramsStruct.asChild != nil {
			paramsOrdered = append(paramsOrdered, *paramsStruct.asChild)
		}
	}
	dispatchPtr, err := o.Com().CallObjectMethod("CreateEntry", paramsOrdered...)
	return notesoutlineentry.New(dispatchPtr), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_CREATEENTRYFROM_METHOD_MEMDEF_NOTESOUTLINE.html */
type createEntryFromParams struct {
	refentry *notesoutlineentry.NotesOutlineEntry
	addAfter *domino.Boolean
	asChild  *domino.Boolean
}

type createEntryFromParam func(*createEntryFromParams)

func WithCreateEntryFromRefentry(refentry notesoutlineentry.NotesOutlineEntry) createEntryFromParam {
	return func(c *createEntryFromParams) {
		c.refentry = &refentry
	}
}

func WithCreateEntryFromAddAfter(addAfter domino.Boolean) createEntryFromParam {
	return func(c *createEntryFromParams) {
		c.addAfter = &addAfter
	}
}

func WithCreateEntryFromAsChild(asChild domino.Boolean) createEntryFromParam {
	return func(c *createEntryFromParams) {
		c.asChild = &asChild
	}
}

func (o NotesOutline) CreateEntryFrom(entry notesoutlineentry.NotesOutlineEntry, params ...createEntryFromParam) (notesoutlineentry.NotesOutlineEntry, error) {
	paramsStruct := &createEntryFromParams{}
	paramsOrdered := []interface{}{entry.Com().Dispatch()}

	for _, p := range params {
		p(paramsStruct)
	}

	if paramsStruct.refentry != nil {
		paramsOrdered = append(paramsOrdered, *paramsStruct.refentry)
		if paramsStruct.addAfter != nil {
			paramsOrdered = append(paramsOrdered, *paramsStruct.addAfter)
			if paramsStruct.asChild != nil {
				paramsOrdered = append(paramsOrdered, *paramsStruct.asChild)
			}
		}
	}
	dispatchPtr, err := o.Com().CallObjectMethod("CreateEntryFrom", paramsOrdered...)
	return notesoutlineentry.New(dispatchPtr), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GETCHILD_METHOD_NOTESOUTLINE.html */
func (o NotesOutline) GetChild(entry notesoutlineentry.NotesOutlineEntry) (notesoutlineentry.NotesOutlineEntry, error) {
	dispatchPtr, err := o.Com().CallObjectMethod("GetChild", entry.Com().Dispatch())
	return notesoutlineentry.New(dispatchPtr), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GETFIRST_METHOD_OUTLINE.html */
func (o NotesOutline) GetFirst() (notesoutlineentry.NotesOutlineEntry, error) {
	dispatchPtr, err := o.Com().CallObjectMethod("GetFirst")
	return notesoutlineentry.New(dispatchPtr), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GETLAST_METHOD_OUTLINE.html */
func (o NotesOutline) GetLast() (notesoutlineentry.NotesOutlineEntry, error) {
	dispatchPtr, err := o.Com().CallObjectMethod("GetLast")
	return notesoutlineentry.New(dispatchPtr), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GETNEXT_METHOD_OUTLINE.html */
func (o NotesOutline) GetNext(entry notesoutlineentry.NotesOutlineEntry) (notesoutlineentry.NotesOutlineEntry, error) {
	dispatchPtr, err := o.Com().CallObjectMethod("GetNext", entry.Com().Dispatch())
	return notesoutlineentry.New(dispatchPtr), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GETNEXTSIBLING_METHOD_OUTLINE.html */
func (o NotesOutline) GetNextSibling(entry notesoutlineentry.NotesOutlineEntry) (notesoutlineentry.NotesOutlineEntry, error) {
	dispatchPtr, err := o.Com().CallObjectMethod("GetNextSibling", entry.Com().Dispatch())
	return notesoutlineentry.New(dispatchPtr), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GETPARENT_METHOD_OUTLINE.html */
func (o NotesOutline) GetParent(entry notesoutlineentry.NotesOutlineEntry) (notesoutlineentry.NotesOutlineEntry, error) {
	dispatchPtr, err := o.Com().CallObjectMethod("GetParent", entry.Com().Dispatch())
	return notesoutlineentry.New(dispatchPtr), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GETPREV_METHOD_OUTLINE.html */
func (o NotesOutline) GetPrev(entry notesoutlineentry.NotesOutlineEntry) (notesoutlineentry.NotesOutlineEntry, error) {
	dispatchPtr, err := o.Com().CallObjectMethod("GetPrev", entry.Com().Dispatch())
	return notesoutlineentry.New(dispatchPtr), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GETPREVSIBLING_METHOD_OUTLINE.html */
func (o NotesOutline) GetPrevSibling(entry notesoutlineentry.NotesOutlineEntry) (notesoutlineentry.NotesOutlineEntry, error) {
	dispatchPtr, err := o.Com().CallObjectMethod("GetPrevSibling", entry.Com().Dispatch())
	return notesoutlineentry.New(dispatchPtr), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_MOVEENTRY_METHOD_OUTLINE.html */
type moveEntryParams struct {
	moveAfter *domino.Boolean
	AsChild   *domino.Boolean
}

type moveEntryParam func(*moveEntryParams)

func WithMoveEntryMoveAfter(moveAfter domino.Boolean) moveEntryParam {
	return func(c *moveEntryParams) {
		c.moveAfter = &moveAfter
	}
}

func WithMoveEntryAsChild(AsChild domino.Boolean) moveEntryParam {
	return func(c *moveEntryParams) {
		c.AsChild = &AsChild
	}
}

func (o NotesOutline) MoveEntry(currentEntry notesoutlineentry.NotesOutlineEntry, refEntry notesoutlineentry.NotesOutlineEntry, params ...moveEntryParam) error {
	paramsStruct := &moveEntryParams{}
	paramsOrdered := []interface{}{currentEntry.Com().Dispatch(), refEntry.Com().Dispatch()}

	for _, p := range params {
		p(paramsStruct)
	}

	if paramsStruct.moveAfter != nil {
		paramsOrdered = append(paramsOrdered, *paramsStruct.moveAfter)
		if paramsStruct.AsChild != nil {
			paramsOrdered = append(paramsOrdered, *paramsStruct.AsChild)
		}
	}
	_, err := o.Com().CallMethod("MoveEntry", paramsOrdered...)
	return err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_REMOVEENTRY_METHOD_OUTLINE.html */
func (o NotesOutline) RemoveEntry(currentEntry notesoutlineentry.NotesOutlineEntry) error {
	_, err := o.Com().CallMethod("RemoveEntry", currentEntry.Com().Dispatch())
	return err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SAVE_METHOD_OUTLINE.html */
func (o NotesOutline) Save() error {
	_, err := o.Com().CallMethod("Save")
	return err
}
