/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_NOTESOUTLINE_CLASS.html */
package domigo

import (
	"github.com/monstermichl/domigo/helpers"

	ole "github.com/go-ole/go-ole"
)

type NotesOutline struct {
	NotesStruct
}

func NewNotesOutline(dispatchPtr *ole.IDispatch) NotesOutline {
	return NotesOutline{NewNotesStruct(dispatchPtr)}
}

/* --------------------------------- Properties --------------------------------- */
/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ALIAS_PROPERTY_OUTLINE.html */
func (o NotesOutline) Alias() (String, error) {
	val, err := o.Com().GetProperty("Alias")
	return helpers.CastValue[String](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ALIAS_PROPERTY_OUTLINE.html */
func (o NotesOutline) SetAlias(v String) error {
	return o.Com().PutProperty("Alias", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_COMMENT_PROPERTY_OUTLINE.html */
func (o NotesOutline) Comment() (String, error) {
	val, err := o.Com().GetProperty("Comment")
	return helpers.CastValue[String](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_COMMENT_PROPERTY_OUTLINE.html */
func (o NotesOutline) SetComment(v String) error {
	return o.Com().PutProperty("Comment", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_NAME_PROPERTY_OUTLINE.html */
func (o NotesOutline) Name() (String, error) {
	val, err := o.Com().GetProperty("Name")
	return helpers.CastValue[String](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_NAME_PROPERTY_OUTLINE.html */
func (o NotesOutline) SetName(v String) error {
	return o.Com().PutProperty("Name", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_PARENTDATABASE_PROPERTY_OUTLINE_COM.html */
func (o NotesOutline) ParentDatabase() (NotesDatabase, error) {
	dispatchPtr, err := o.Com().GetObjectProperty("ParentDatabase")
	return NewNotesDatabase(dispatchPtr), err
}

/* --------------------------------- Methods ------------------------------------ */
/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ADDENTRY_METHOD_OUTLINE.html */
type notesOutlineAddEntryParams struct {
	refEntry  *NotesOutlineEntry
	moveAfter *Boolean
	AsChild   *Boolean
}

type notesOutlineAddEntryParam func(*notesOutlineAddEntryParams)

func WithNotesOutlineAddEntryRefEntry(refEntry NotesOutlineEntry) notesOutlineAddEntryParam {
	return func(c *notesOutlineAddEntryParams) {
		c.refEntry = &refEntry
	}
}

func WithNotesOutlineAddEntryMoveAfter(moveAfter Boolean) notesOutlineAddEntryParam {
	return func(c *notesOutlineAddEntryParams) {
		c.moveAfter = &moveAfter
	}
}

func WithNotesOutlineAddEntryAsChild(AsChild Boolean) notesOutlineAddEntryParam {
	return func(c *notesOutlineAddEntryParams) {
		c.AsChild = &AsChild
	}
}

func (o NotesOutline) AddEntry(newEntry NotesOutlineEntry, params ...notesOutlineAddEntryParam) error {
	paramsStruct := &notesOutlineAddEntryParams{}
	paramsOrdered := []interface{}{newEntry.Com().Dispatch()}

	for _, p := range params {
		p(paramsStruct)
	}

	if paramsStruct.refEntry != nil {
		paramsOrdered = append(paramsOrdered, *paramsStruct.refEntry.Com().Dispatch())
		if paramsStruct.moveAfter != nil {
			paramsOrdered = append(paramsOrdered, *paramsStruct.moveAfter)
			if paramsStruct.AsChild != nil {
				paramsOrdered = append(paramsOrdered, *paramsStruct.AsChild)
			}
		}
	}
	_, err := o.Com().CallMethod("AddEntry", paramsOrdered...)
	return err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_CREATEENTRY_METHOD_OUTLINE.html */
type notesOutlineCreateEntryParams struct {
	refentry *NotesOutlineEntry
	addAfter *Boolean
	asChild  *Boolean
}

type notesOutlineCreateEntryParam func(*notesOutlineCreateEntryParams)

func WithNotesOutlineCreateEntryRefentry(refentry NotesOutlineEntry) notesOutlineCreateEntryParam {
	return func(c *notesOutlineCreateEntryParams) {
		c.refentry = &refentry
	}
}

func WithNotesOutlineCreateEntryAddAfter(addAfter Boolean) notesOutlineCreateEntryParam {
	return func(c *notesOutlineCreateEntryParams) {
		c.addAfter = &addAfter
	}
}

func WithNotesOutlineCreateEntryAsChild(asChild Boolean) notesOutlineCreateEntryParam {
	return func(c *notesOutlineCreateEntryParams) {
		c.asChild = &asChild
	}
}

func (o NotesOutline) CreateEntry(name String, params ...notesOutlineCreateEntryParam) (NotesOutlineEntry, error) {
	paramsStruct := &notesOutlineCreateEntryParams{}
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
	return NewNotesOutlineEntry(dispatchPtr), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_CREATEENTRYFROM_METHOD_MEMDEF_NOTESOUTLINE.html */
type notesOutlineCreateEntryFromParams struct {
	refentry *NotesOutlineEntry
	addAfter *Boolean
	asChild  *Boolean
}

type notesOutlineCreateEntryFromParam func(*notesOutlineCreateEntryFromParams)

func WithNotesOutlineCreateEntryFromRefentry(refentry NotesOutlineEntry) notesOutlineCreateEntryFromParam {
	return func(c *notesOutlineCreateEntryFromParams) {
		c.refentry = &refentry
	}
}

func WithNotesOutlineCreateEntryFromAddAfter(addAfter Boolean) notesOutlineCreateEntryFromParam {
	return func(c *notesOutlineCreateEntryFromParams) {
		c.addAfter = &addAfter
	}
}

func WithNotesOutlineCreateEntryFromAsChild(asChild Boolean) notesOutlineCreateEntryFromParam {
	return func(c *notesOutlineCreateEntryFromParams) {
		c.asChild = &asChild
	}
}

func (o NotesOutline) CreateEntryFrom(entry NotesOutlineEntry, params ...notesOutlineCreateEntryFromParam) (NotesOutlineEntry, error) {
	paramsStruct := &notesOutlineCreateEntryFromParams{}
	paramsOrdered := []interface{}{entry.Com().Dispatch()}

	for _, p := range params {
		p(paramsStruct)
	}

	if paramsStruct.refentry != nil {
		paramsOrdered = append(paramsOrdered, *paramsStruct.refentry.Com().Dispatch())
		if paramsStruct.addAfter != nil {
			paramsOrdered = append(paramsOrdered, *paramsStruct.addAfter)
			if paramsStruct.asChild != nil {
				paramsOrdered = append(paramsOrdered, *paramsStruct.asChild)
			}
		}
	}
	dispatchPtr, err := o.Com().CallObjectMethod("CreateEntryFrom", paramsOrdered...)
	return NewNotesOutlineEntry(dispatchPtr), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GETCHILD_METHOD_NOTESOUTLINE.html */
func (o NotesOutline) GetChild(entry NotesOutlineEntry) (NotesOutlineEntry, error) {
	dispatchPtr, err := o.Com().CallObjectMethod("GetChild", entry.Com().Dispatch())
	return NewNotesOutlineEntry(dispatchPtr), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GETFIRST_METHOD_OUTLINE.html */
func (o NotesOutline) GetFirst() (NotesOutlineEntry, error) {
	dispatchPtr, err := o.Com().CallObjectMethod("GetFirst")
	return NewNotesOutlineEntry(dispatchPtr), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GETLAST_METHOD_OUTLINE.html */
func (o NotesOutline) GetLast() (NotesOutlineEntry, error) {
	dispatchPtr, err := o.Com().CallObjectMethod("GetLast")
	return NewNotesOutlineEntry(dispatchPtr), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GETNEXT_METHOD_OUTLINE.html */
func (o NotesOutline) GetNext(entry NotesOutlineEntry) (NotesOutlineEntry, error) {
	dispatchPtr, err := o.Com().CallObjectMethod("GetNext", entry.Com().Dispatch())
	return NewNotesOutlineEntry(dispatchPtr), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GETNEXTSIBLING_METHOD_OUTLINE.html */
func (o NotesOutline) GetNextSibling(entry NotesOutlineEntry) (NotesOutlineEntry, error) {
	dispatchPtr, err := o.Com().CallObjectMethod("GetNextSibling", entry.Com().Dispatch())
	return NewNotesOutlineEntry(dispatchPtr), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GETPARENT_METHOD_OUTLINE.html */
func (o NotesOutline) GetParent(entry NotesOutlineEntry) (NotesOutlineEntry, error) {
	dispatchPtr, err := o.Com().CallObjectMethod("GetParent", entry.Com().Dispatch())
	return NewNotesOutlineEntry(dispatchPtr), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GETPREV_METHOD_OUTLINE.html */
func (o NotesOutline) GetPrev(entry NotesOutlineEntry) (NotesOutlineEntry, error) {
	dispatchPtr, err := o.Com().CallObjectMethod("GetPrev", entry.Com().Dispatch())
	return NewNotesOutlineEntry(dispatchPtr), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GETPREVSIBLING_METHOD_OUTLINE.html */
func (o NotesOutline) GetPrevSibling(entry NotesOutlineEntry) (NotesOutlineEntry, error) {
	dispatchPtr, err := o.Com().CallObjectMethod("GetPrevSibling", entry.Com().Dispatch())
	return NewNotesOutlineEntry(dispatchPtr), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_MOVEENTRY_METHOD_OUTLINE.html */
type notesOutlineMoveEntryParams struct {
	moveAfter *Boolean
	AsChild   *Boolean
}

type notesOutlineMoveEntryParam func(*notesOutlineMoveEntryParams)

func WithNotesOutlineMoveEntryMoveAfter(moveAfter Boolean) notesOutlineMoveEntryParam {
	return func(c *notesOutlineMoveEntryParams) {
		c.moveAfter = &moveAfter
	}
}

func WithNotesOutlineMoveEntryAsChild(AsChild Boolean) notesOutlineMoveEntryParam {
	return func(c *notesOutlineMoveEntryParams) {
		c.AsChild = &AsChild
	}
}

func (o NotesOutline) MoveEntry(currentEntry NotesOutlineEntry, refEntry NotesOutlineEntry, params ...notesOutlineMoveEntryParam) error {
	paramsStruct := &notesOutlineMoveEntryParams{}
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
func (o NotesOutline) RemoveEntry(currentEntry NotesOutlineEntry) error {
	_, err := o.Com().CallMethod("RemoveEntry", currentEntry.Com().Dispatch())
	return err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SAVE_METHOD_OUTLINE.html */
func (o NotesOutline) Save() error {
	_, err := o.Com().CallMethod("Save")
	return err
}
