/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_NOTESREPLICATIONENTRY_CLASS.html */
package domigo

import (
	"github.com/monstermichl/domigo/internal/helpers"

	ole "github.com/go-ole/go-ole"
)

type NotesReplicationEntry struct {
	NotesStruct
}

func NewNotesReplicationEntry(dispatchPtr *ole.IDispatch) NotesReplicationEntry {
	return NotesReplicationEntry{NewNotesStruct(dispatchPtr)}
}

/* --------------------------------- Properties --------------------------------- */
/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_DESTINATION_PROPERTY_RE.html */
func (r NotesReplicationEntry) Destination() (String, error) {
	val, err := getComProperty(r, "Destination")
	return helpers.CastValue[String](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_FORMULA_PROPERTY_RE.html */
func (r NotesReplicationEntry) Formula() (String, error) {
	val, err := getComProperty(r, "Formula")
	return helpers.CastValue[String](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_FORMULA_PROPERTY_RE.html */
func (r NotesReplicationEntry) SetFormula(v String) error {
	return putComProperty(r, "Formula", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISINCLUDEACL_PROPERTY_RE.html */
func (r NotesReplicationEntry) IsIncludeACL() (Boolean, error) {
	val, err := getComProperty(r, "IsIncludeACL")
	return helpers.CastValue[Boolean](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISINCLUDEACL_PROPERTY_RE.html */
func (r NotesReplicationEntry) SetIsIncludeACL(v Boolean) error {
	return putComProperty(r, "IsIncludeACL", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISINCLUDEAGENTS_PROPERTY_RE.html */
func (r NotesReplicationEntry) IsIncludeAgents() (Boolean, error) {
	val, err := getComProperty(r, "IsIncludeAgents")
	return helpers.CastValue[Boolean](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISINCLUDEAGENTS_PROPERTY_RE.html */
func (r NotesReplicationEntry) SetIsIncludeAgents(v Boolean) error {
	return putComProperty(r, "IsIncludeAgents", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISINCLUDEDOCUMENTS_PROPERTY_RE.html */
func (r NotesReplicationEntry) IsIncludeDocuments() (Boolean, error) {
	val, err := getComProperty(r, "IsIncludeDocuments")
	return helpers.CastValue[Boolean](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISINCLUDEDOCUMENTS_PROPERTY_RE.html */
func (r NotesReplicationEntry) SetIsIncludeDocuments(v Boolean) error {
	return putComProperty(r, "IsIncludeDocuments", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISINCLUDEFORMS_PROPERTY_RE.html */
func (r NotesReplicationEntry) IsIncludeForms() (Boolean, error) {
	val, err := getComProperty(r, "IsIncludeForms")
	return helpers.CastValue[Boolean](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISINCLUDEFORMS_PROPERTY_RE.html */
func (r NotesReplicationEntry) SetIsIncludeForms(v Boolean) error {
	return putComProperty(r, "IsIncludeForms", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISINCLUDEFORMULAS_PROPERTY_RE.html */
func (r NotesReplicationEntry) IsIncludeFormulas() (Boolean, error) {
	val, err := getComProperty(r, "IsIncludeFormulas")
	return helpers.CastValue[Boolean](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISINCLUDEFORMULAS_PROPERTY_RE.html */
func (r NotesReplicationEntry) SetIsIncludeFormulas(v Boolean) error {
	return putComProperty(r, "IsIncludeFormulas", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SOURCE_PROPERTY_RE.html */
func (r NotesReplicationEntry) Source() (String, error) {
	val, err := getComProperty(r, "Source")
	return helpers.CastValue[String](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_VIEWS_PROPERTY_RE.html */
func (r NotesReplicationEntry) Views() (String, error) {
	val, err := getComProperty(r, "Views")
	return helpers.CastValue[String](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_VIEWS_PROPERTY_RE.html */
func (r NotesReplicationEntry) SetViews(v String) error {
	return putComProperty(r, "Views", v)
}

/* --------------------------------- Methods ------------------------------------ */
/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_REMOVE_METHOD_RE.html */
func (r NotesReplicationEntry) Remove() error {
	_, err := callComMethod(r, "Remove")
	return err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SAVE_METHOD_RE.html */
func (r NotesReplicationEntry) Save() error {
	_, err := callComMethod(r, "Save")
	return err
}
