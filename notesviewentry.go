/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_NOTESVIEWENTRY_CLASS_2925.html */
package domigo

import (
	ole "github.com/go-ole/go-ole"
)

type NotesViewEntry struct {
	NotesStruct
}

func NewNotesViewEntry(dispatchPtr *ole.IDispatch) NotesViewEntry {
	return NotesViewEntry{NewNotesStruct(dispatchPtr)}
}

/* --------------------------------- Properties --------------------------------- */
/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_CHILDCOUNT_PROPERTY_8963.html */
func (v NotesViewEntry) ChildCount() (Long, error) {
	return getComProperty[Long](v, "ChildCount")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_COLUMNINDENTLEVEL_PROPERTY_VIEWENTRY.html */
func (v NotesViewEntry) ColumnIndentLevel() (Long, error) {
	return getComProperty[Long](v, "ColumnIndentLevel")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_COLUMNVALUES_PROPERTY_5887.html */
func (v NotesViewEntry) ColumnValues() ([]any, error) {
	return getComArrayProperty(v, "ColumnValues")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_DESCENDANTCOUNT_PROPERTY_2777.html */
func (v NotesViewEntry) DescendantCount() (Long, error) {
	return getComProperty[Long](v, "DescendantCount")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_DOCUMENT_PROPERTY_7936.html */
func (v NotesViewEntry) Document() (NotesDocument, error) {
	return getComObjectProperty(v, NewNotesDocument, "Document")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_FTSEARCHSCORE_PROPERTY_2767.html */
func (v NotesViewEntry) FTSearchScore() (Long, error) {
	return getComProperty[Long](v, "FTSearchScore")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_INDENTLEVEL_PROPERTY_8244.html */
func (v NotesViewEntry) IndentLevel() (Long, error) {
	return getComProperty[Long](v, "IndentLevel")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISCATEGORY_PROPERTY_9511.html */
func (v NotesViewEntry) IsCategory() (Boolean, error) {
	return getComProperty[Boolean](v, "IsCategory")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISCONFLICT_PROPERTY_1204.html */
func (v NotesViewEntry) IsConflict() (Boolean, error) {
	return getComProperty[Boolean](v, "IsConflict")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISDOCUMENT_PROPERTY_6058.html */
func (v NotesViewEntry) IsDocument() (Boolean, error) {
	return getComProperty[Boolean](v, "IsDocument")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISTOTAL_PROPERTY_9664.html */
func (v NotesViewEntry) IsTotal() (Boolean, error) {
	return getComProperty[Boolean](v, "IsTotal")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISVALID_PROPERTY_2176.html */
func (v NotesViewEntry) IsValid() (Boolean, error) {
	return getComProperty[Boolean](v, "IsValid")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_NOTEID_PROPERTY_9904.html */
func (v NotesViewEntry) NoteID() (String, error) {
	return getComProperty[String](v, "NoteID")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_PARENT_PROPERTY_3044.html */
func (v NotesViewEntry) Parent() (NotesView, error) {
	return getComObjectProperty(v, NewNotesView, "Parent")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SIBLINGCOUNT_PROPERTY_5642.html */
func (v NotesViewEntry) SiblingCount() (Long, error) {
	return getComProperty[Long](v, "SiblingCount")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_UNIVERSALID_PROPERTY_8372.html */
func (v NotesViewEntry) UniversalID() (String, error) {
	return getComProperty[String](v, "UniversalID")
}

/* --------------------------------- Methods ------------------------------------ */
/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GETPOSITION_METHOD_8121.html */
func (v NotesViewEntry) GetPosition(separator String) (String, error) {
	return callComMethod[String](v, "GetPosition", separator)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GETREAD_METHOD_ENTRY.html */
type notesViewEntryGetReadParams struct {
	username *String
}

type notesViewEntryGetReadParam func(*notesViewEntryGetReadParams)

func WithNotesViewEntryGetReadUsername(username String) notesViewEntryGetReadParam {
	return func(c *notesViewEntryGetReadParams) {
		c.username = &username
	}
}

func (v NotesViewEntry) GetRead(params ...notesViewEntryGetReadParam) (Boolean, error) {
	paramsStruct := &notesViewEntryGetReadParams{}
	paramsOrdered := []interface{}{}

	for _, p := range params {
		p(paramsStruct)
	}

	if paramsStruct.username != nil {
		paramsOrdered = append(paramsOrdered, *paramsStruct.username)
	}
	return callComMethod[Boolean](v, "GetRead", paramsOrdered...)
}
