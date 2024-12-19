/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_NOTESVIEWENTRY_CLASS_2925.html */
package domigo

import (
	"github.com/monstermichl/domigo/helpers"

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
	val, err := v.Com().GetProperty("ChildCount")
	return helpers.CastValue[Long](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_COLUMNINDENTLEVEL_PROPERTY_VIEWENTRY.html */
func (v NotesViewEntry) ColumnIndentLevel() (Long, error) {
	val, err := v.Com().GetProperty("ColumnIndentLevel")
	return helpers.CastValue[Long](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_COLUMNVALUES_PROPERTY_5887.html */
func (v NotesViewEntry) ColumnValues() ([]any, error) {
	return v.Com().GetArrayProperty("ColumnValues")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_DESCENDANTCOUNT_PROPERTY_2777.html */
func (v NotesViewEntry) DescendantCount() (Long, error) {
	val, err := v.Com().GetProperty("DescendantCount")
	return helpers.CastValue[Long](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_DOCUMENT_PROPERTY_7936.html */
func (v NotesViewEntry) Document() (NotesDocument, error) {
	dispatchPtr, err := v.Com().GetObjectProperty("Document")
	return NewNotesDocument(dispatchPtr), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_FTSEARCHSCORE_PROPERTY_2767.html */
func (v NotesViewEntry) FTSearchScore() (Long, error) {
	val, err := v.Com().GetProperty("FTSearchScore")
	return helpers.CastValue[Long](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_INDENTLEVEL_PROPERTY_8244.html */
func (v NotesViewEntry) IndentLevel() (Long, error) {
	val, err := v.Com().GetProperty("IndentLevel")
	return helpers.CastValue[Long](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISCATEGORY_PROPERTY_9511.html */
func (v NotesViewEntry) IsCategory() (Boolean, error) {
	val, err := v.Com().GetProperty("IsCategory")
	return helpers.CastValue[Boolean](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISCONFLICT_PROPERTY_1204.html */
func (v NotesViewEntry) IsConflict() (Boolean, error) {
	val, err := v.Com().GetProperty("IsConflict")
	return helpers.CastValue[Boolean](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISDOCUMENT_PROPERTY_6058.html */
func (v NotesViewEntry) IsDocument() (Boolean, error) {
	val, err := v.Com().GetProperty("IsDocument")
	return helpers.CastValue[Boolean](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISTOTAL_PROPERTY_9664.html */
func (v NotesViewEntry) IsTotal() (Boolean, error) {
	val, err := v.Com().GetProperty("IsTotal")
	return helpers.CastValue[Boolean](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISVALID_PROPERTY_2176.html */
func (v NotesViewEntry) IsValid() (Boolean, error) {
	val, err := v.Com().GetProperty("IsValid")
	return helpers.CastValue[Boolean](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_NOTEID_PROPERTY_9904.html */
func (v NotesViewEntry) NoteID() (String, error) {
	val, err := v.Com().GetProperty("NoteID")
	return helpers.CastValue[String](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_PARENT_PROPERTY_3044.html */
func (v NotesViewEntry) Parent() (NotesView, error) {
	dispatchPtr, err := v.Com().GetObjectProperty("Parent")
	return NewNotesView(dispatchPtr), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SIBLINGCOUNT_PROPERTY_5642.html */
func (v NotesViewEntry) SiblingCount() (Long, error) {
	val, err := v.Com().GetProperty("SiblingCount")
	return helpers.CastValue[Long](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_UNIVERSALID_PROPERTY_8372.html */
func (v NotesViewEntry) UniversalID() (String, error) {
	val, err := v.Com().GetProperty("UniversalID")
	return helpers.CastValue[String](val), err
}

/* --------------------------------- Methods ------------------------------------ */
/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GETPOSITION_METHOD_8121.html */
func (v NotesViewEntry) GetPosition(separator String) (String, error) {
	val, err := v.Com().CallMethod("GetPosition", separator)
	return helpers.CastValue[String](val), err
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
	val, err := v.Com().CallMethod("GetRead", paramsOrdered...)
	return helpers.CastValue[Boolean](val), err
}
