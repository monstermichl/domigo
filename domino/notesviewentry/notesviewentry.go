package notesviewentry

import (
	"domigo/domino"
	"domigo/domino/notesdocument"

	"domigo/helpers"

	ole "github.com/go-ole/go-ole"
)

type NotesViewEntry struct {
	domino.NotesStruct
}

func New(dispatchPtr *ole.IDispatch) NotesViewEntry {
	return NotesViewEntry{domino.NewNotesStruct(dispatchPtr)}
}

/* --------------------------------- Properties --------------------------------- */
/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_CHILDCOUNT_PROPERTY_8963.html */
func (v NotesViewEntry) ChildCount() (domino.Integer, error) {
	val, err := v.Com().GetProperty("ChildCount")
	return helpers.CastValue[domino.Integer](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_COLUMNINDENTLEVEL_PROPERTY_VIEWENTRY.html */
func (v NotesViewEntry) ColumnIndentLevel() (domino.Integer, error) {
	val, err := v.Com().GetProperty("ColumnIndentLevel")
	return helpers.CastValue[domino.Integer](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_COLUMNVALUES_PROPERTY_5887.html */
func (v NotesViewEntry) ColumnValues() ([]domino.Variant, error) {
	vals, err := v.Com().GetArrayProperty("ColumnValues")
	return helpers.CastSlice[domino.Variant](vals), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_DESCENDANTCOUNT_PROPERTY_2777.html */
func (v NotesViewEntry) DescendantCount() (domino.Integer, error) {
	val, err := v.Com().GetProperty("DescendantCount")
	return helpers.CastValue[domino.Integer](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_DOCUMENT_PROPERTY_7936.html */
func (v NotesViewEntry) Document() (notesdocument.NotesDocument, error) {
	dispatchPtr, err := v.Com().GetObjectProperty("Document")
	return notesdocument.New(dispatchPtr), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_FTSEARCHSCORE_PROPERTY_2767.html */
func (v NotesViewEntry) FTSearchScore() (domino.Integer, error) {
	val, err := v.Com().GetProperty("FTSearchScore")
	return helpers.CastValue[domino.Integer](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_INDENTLEVEL_PROPERTY_8244.html */
func (v NotesViewEntry) IndentLevel() (domino.Integer, error) {
	val, err := v.Com().GetProperty("IndentLevel")
	return helpers.CastValue[domino.Integer](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISCATEGORY_PROPERTY_9511.html */
func (v NotesViewEntry) IsCategory() (domino.Boolean, error) {
	val, err := v.Com().GetProperty("IsCategory")
	return helpers.CastValue[domino.Boolean](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISCONFLICT_PROPERTY_1204.html */
func (v NotesViewEntry) IsConflict() (domino.Boolean, error) {
	val, err := v.Com().GetProperty("IsConflict")
	return helpers.CastValue[domino.Boolean](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISDOCUMENT_PROPERTY_6058.html */
func (v NotesViewEntry) IsDocument() (domino.Boolean, error) {
	val, err := v.Com().GetProperty("IsDocument")
	return helpers.CastValue[domino.Boolean](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISTOTAL_PROPERTY_9664.html */
func (v NotesViewEntry) IsTotal() (domino.Boolean, error) {
	val, err := v.Com().GetProperty("IsTotal")
	return helpers.CastValue[domino.Boolean](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISVALID_PROPERTY_2176.html */
func (v NotesViewEntry) IsValid() (domino.Boolean, error) {
	val, err := v.Com().GetProperty("IsValid")
	return helpers.CastValue[domino.Boolean](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_NOTEID_PROPERTY_9904.html */
func (v NotesViewEntry) NoteID() (domino.String, error) {
	val, err := v.Com().GetProperty("NoteID")
	return helpers.CastValue[domino.String](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SIBLINGCOUNT_PROPERTY_5642.html */
func (v NotesViewEntry) SiblingCount() (domino.Integer, error) {
	val, err := v.Com().GetProperty("SiblingCount")
	return helpers.CastValue[domino.Integer](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_UNIVERSALID_PROPERTY_8372.html */
func (v NotesViewEntry) UniversalID() (domino.String, error) {
	val, err := v.Com().GetProperty("UniversalID")
	return helpers.CastValue[domino.String](val), err
}

/* --------------------------------- Methods ------------------------------------ */
/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GETPOSITION_METHOD_8121.html */
func (v NotesViewEntry) GetPosition(separator domino.String) (domino.String, error) {
	val, err := v.Com().CallMethod("GetPosition", separator)
	return helpers.CastValue[domino.String](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GETREAD_METHOD_ENTRY.html */
/* TODO: Some parameters are optional. Make sure to handle them correctly. */
func (v NotesViewEntry) GetRead(username domino.String) (domino.Boolean, error) {
	val, err := v.Com().CallMethod("GetRead", username)
	return helpers.CastValue[domino.Boolean](val), err
}
