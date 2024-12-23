/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_NOTESOUTLINE_CLASS.html */
package domigo

import (
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
	return getComProperty[String](o, "Alias")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ALIAS_PROPERTY_OUTLINE.html */
func (o NotesOutline) SetAlias(v String) error {
	return putComProperty(o, "Alias", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_COMMENT_PROPERTY_OUTLINE.html */
func (o NotesOutline) Comment() (String, error) {
	return getComProperty[String](o, "Comment")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_COMMENT_PROPERTY_OUTLINE.html */
func (o NotesOutline) SetComment(v String) error {
	return putComProperty(o, "Comment", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_NAME_PROPERTY_OUTLINE.html */
func (o NotesOutline) Name() (String, error) {
	return getComProperty[String](o, "Name")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_NAME_PROPERTY_OUTLINE.html */
func (o NotesOutline) SetName(v String) error {
	return putComProperty(o, "Name", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_PARENTDATABASE_PROPERTY_OUTLINE_COM.html */
func (o NotesOutline) ParentDatabase() (NotesDatabase, error) {
	return getComObjectProperty(o, NewNotesDatabase, "ParentDatabase")
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
	paramsOrdered := []interface{}{newEntry.com().Dispatch()}

	for _, p := range params {
		p(paramsStruct)
	}

	if paramsStruct.refEntry != nil {
		paramsOrdered = append(paramsOrdered, *paramsStruct.refEntry.com().Dispatch())
		if paramsStruct.moveAfter != nil {
			paramsOrdered = append(paramsOrdered, *paramsStruct.moveAfter)
			if paramsStruct.AsChild != nil {
				paramsOrdered = append(paramsOrdered, *paramsStruct.AsChild)
			}
		}
	}
	return callComVoidMethod(o, "AddEntry", paramsOrdered...)
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
		paramsOrdered = append(paramsOrdered, paramsStruct.refentry.com().Dispatch())
	}

	if paramsStruct.addAfter != nil {
		paramsOrdered = append(paramsOrdered, *paramsStruct.addAfter)
		if paramsStruct.asChild != nil {
			paramsOrdered = append(paramsOrdered, *paramsStruct.asChild)
		}
	}
	return callComObjectMethod(o, NewNotesOutlineEntry, "CreateEntry", paramsOrdered...)
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
	paramsOrdered := []interface{}{entry.com().Dispatch()}

	for _, p := range params {
		p(paramsStruct)
	}

	if paramsStruct.refentry != nil {
		paramsOrdered = append(paramsOrdered, *paramsStruct.refentry.com().Dispatch())
		if paramsStruct.addAfter != nil {
			paramsOrdered = append(paramsOrdered, *paramsStruct.addAfter)
			if paramsStruct.asChild != nil {
				paramsOrdered = append(paramsOrdered, *paramsStruct.asChild)
			}
		}
	}
	return callComObjectMethod(o, NewNotesOutlineEntry, "CreateEntryFrom", paramsOrdered...)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GETCHILD_METHOD_NOTESOUTLINE.html */
func (o NotesOutline) GetChild(entry NotesOutlineEntry) (NotesOutlineEntry, error) {
	return callComObjectMethod(o, NewNotesOutlineEntry, "GetChild", entry.com().Dispatch())
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GETFIRST_METHOD_OUTLINE.html */
func (o NotesOutline) GetFirst() (NotesOutlineEntry, error) {
	return callComObjectMethod(o, NewNotesOutlineEntry, "GetFirst")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GETLAST_METHOD_OUTLINE.html */
func (o NotesOutline) GetLast() (NotesOutlineEntry, error) {
	return callComObjectMethod(o, NewNotesOutlineEntry, "GetLast")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GETNEXT_METHOD_OUTLINE.html */
func (o NotesOutline) GetNext(entry NotesOutlineEntry) (NotesOutlineEntry, error) {
	return callComObjectMethod(o, NewNotesOutlineEntry, "GetNext", entry.com().Dispatch())
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GETNEXTSIBLING_METHOD_OUTLINE.html */
func (o NotesOutline) GetNextSibling(entry NotesOutlineEntry) (NotesOutlineEntry, error) {
	return callComObjectMethod(o, NewNotesOutlineEntry, "GetNextSibling", entry.com().Dispatch())
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GETPARENT_METHOD_OUTLINE.html */
func (o NotesOutline) GetParent(entry NotesOutlineEntry) (NotesOutlineEntry, error) {
	return callComObjectMethod(o, NewNotesOutlineEntry, "GetParent", entry.com().Dispatch())
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GETPREV_METHOD_OUTLINE.html */
func (o NotesOutline) GetPrev(entry NotesOutlineEntry) (NotesOutlineEntry, error) {
	return callComObjectMethod(o, NewNotesOutlineEntry, "GetPrev", entry.com().Dispatch())
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GETPREVSIBLING_METHOD_OUTLINE.html */
func (o NotesOutline) GetPrevSibling(entry NotesOutlineEntry) (NotesOutlineEntry, error) {
	return callComObjectMethod(o, NewNotesOutlineEntry, "GetPrevSibling", entry.com().Dispatch())
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
	paramsOrdered := []interface{}{currentEntry.com().Dispatch(), refEntry.com().Dispatch()}

	for _, p := range params {
		p(paramsStruct)
	}

	if paramsStruct.moveAfter != nil {
		paramsOrdered = append(paramsOrdered, *paramsStruct.moveAfter)
		if paramsStruct.AsChild != nil {
			paramsOrdered = append(paramsOrdered, *paramsStruct.AsChild)
		}
	}
	return callComVoidMethod(o, "MoveEntry", paramsOrdered...)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_REMOVEENTRY_METHOD_OUTLINE.html */
func (o NotesOutline) RemoveEntry(currentEntry NotesOutlineEntry) error {
	return callComVoidMethod(o, "RemoveEntry", currentEntry.com().Dispatch())
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SAVE_METHOD_OUTLINE.html */
func (o NotesOutline) Save() error {
	return callComVoidMethod(o, "Save")
}
