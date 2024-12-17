package notesreplicationentry

import (
	"github.com/monstermichl/domigo/domino"
	"github.com/monstermichl/domigo/helpers"

	ole "github.com/go-ole/go-ole"
)

type NotesReplicationEntry struct {
	domino.NotesStruct
}

func New(dispatchPtr *ole.IDispatch) NotesReplicationEntry {
	return NotesReplicationEntry{domino.NewNotesStruct(dispatchPtr)}
}

/* --------------------------------- Properties --------------------------------- */
/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_DESTINATION_PROPERTY_RE.html */
func (r NotesReplicationEntry) Destination() (domino.String, error) {
	val, err := r.Com().GetProperty("Destination")
	return helpers.CastValue[domino.String](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_FORMULA_PROPERTY_RE.html */
func (r NotesReplicationEntry) Formula() (domino.String, error) {
	val, err := r.Com().GetProperty("Formula")
	return helpers.CastValue[domino.String](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_FORMULA_PROPERTY_RE.html */
func (r NotesReplicationEntry) SetFormula(v domino.String) error {
	return r.Com().PutProperty("Formula", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISINCLUDEACL_PROPERTY_RE.html */
func (r NotesReplicationEntry) IsIncludeACL() (domino.Boolean, error) {
	val, err := r.Com().GetProperty("IsIncludeACL")
	return helpers.CastValue[domino.Boolean](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISINCLUDEACL_PROPERTY_RE.html */
func (r NotesReplicationEntry) SetIsIncludeACL(v domino.Boolean) error {
	return r.Com().PutProperty("IsIncludeACL", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISINCLUDEAGENTS_PROPERTY_RE.html */
func (r NotesReplicationEntry) IsIncludeAgents() (domino.Boolean, error) {
	val, err := r.Com().GetProperty("IsIncludeAgents")
	return helpers.CastValue[domino.Boolean](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISINCLUDEAGENTS_PROPERTY_RE.html */
func (r NotesReplicationEntry) SetIsIncludeAgents(v domino.Boolean) error {
	return r.Com().PutProperty("IsIncludeAgents", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISINCLUDEDOCUMENTS_PROPERTY_RE.html */
func (r NotesReplicationEntry) IsIncludeDocuments() (domino.Boolean, error) {
	val, err := r.Com().GetProperty("IsIncludeDocuments")
	return helpers.CastValue[domino.Boolean](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISINCLUDEDOCUMENTS_PROPERTY_RE.html */
func (r NotesReplicationEntry) SetIsIncludeDocuments(v domino.Boolean) error {
	return r.Com().PutProperty("IsIncludeDocuments", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISINCLUDEFORMS_PROPERTY_RE.html */
func (r NotesReplicationEntry) IsIncludeForms() (domino.Boolean, error) {
	val, err := r.Com().GetProperty("IsIncludeForms")
	return helpers.CastValue[domino.Boolean](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISINCLUDEFORMS_PROPERTY_RE.html */
func (r NotesReplicationEntry) SetIsIncludeForms(v domino.Boolean) error {
	return r.Com().PutProperty("IsIncludeForms", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISINCLUDEFORMULAS_PROPERTY_RE.html */
func (r NotesReplicationEntry) IsIncludeFormulas() (domino.Boolean, error) {
	val, err := r.Com().GetProperty("IsIncludeFormulas")
	return helpers.CastValue[domino.Boolean](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISINCLUDEFORMULAS_PROPERTY_RE.html */
func (r NotesReplicationEntry) SetIsIncludeFormulas(v domino.Boolean) error {
	return r.Com().PutProperty("IsIncludeFormulas", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SOURCE_PROPERTY_RE.html */
func (r NotesReplicationEntry) Source() (domino.String, error) {
	val, err := r.Com().GetProperty("Source")
	return helpers.CastValue[domino.String](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_VIEWS_PROPERTY_RE.html */
func (r NotesReplicationEntry) Views() (domino.String, error) {
	val, err := r.Com().GetProperty("Views")
	return helpers.CastValue[domino.String](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_VIEWS_PROPERTY_RE.html */
func (r NotesReplicationEntry) SetViews(v domino.String) error {
	return r.Com().PutProperty("Views", v)
}

/* --------------------------------- Methods ------------------------------------ */
/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_REMOVE_METHOD_RE.html */
func (r NotesReplicationEntry) Remove() error {
	_, err := r.Com().CallMethod("Remove")
	return err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SAVE_METHOD_RE.html */
func (r NotesReplicationEntry) Save() error {
	_, err := r.Com().CallMethod("Save")
	return err
}
