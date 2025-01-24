/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_NOTESREPLICATIONENTRY_CLASS.html */
package domigo

import (
	ole "github.com/go-ole/go-ole"
)

type NotesReplicationEntry struct {
	NotesStruct
}

func newNotesReplicationEntry(dispatchPtr *ole.IDispatch) NotesReplicationEntry {
	return NotesReplicationEntry{newNotesStruct(dispatchPtr)}
}

/* --------------------------------- Properties --------------------------------- */
/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_DESTINATION_PROPERTY_RE.html */
func (r NotesReplicationEntry) Destination() (String, error) {
	return getComProperty[String](r, "Destination")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_FORMULA_PROPERTY_RE.html */
func (r NotesReplicationEntry) Formula() (String, error) {
	return getComProperty[String](r, "Formula")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_FORMULA_PROPERTY_RE.html */
func (r NotesReplicationEntry) SetFormula(v String) error {
	return putComProperty(r, "Formula", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISINCLUDEACL_PROPERTY_RE.html */
func (r NotesReplicationEntry) IsIncludeACL() (Boolean, error) {
	return getComProperty[Boolean](r, "IsIncludeACL")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISINCLUDEACL_PROPERTY_RE.html */
func (r NotesReplicationEntry) SetIsIncludeACL(v Boolean) error {
	return putComProperty(r, "IsIncludeACL", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISINCLUDEAGENTS_PROPERTY_RE.html */
func (r NotesReplicationEntry) IsIncludeAgents() (Boolean, error) {
	return getComProperty[Boolean](r, "IsIncludeAgents")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISINCLUDEAGENTS_PROPERTY_RE.html */
func (r NotesReplicationEntry) SetIsIncludeAgents(v Boolean) error {
	return putComProperty(r, "IsIncludeAgents", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISINCLUDEDOCUMENTS_PROPERTY_RE.html */
func (r NotesReplicationEntry) IsIncludeDocuments() (Boolean, error) {
	return getComProperty[Boolean](r, "IsIncludeDocuments")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISINCLUDEDOCUMENTS_PROPERTY_RE.html */
func (r NotesReplicationEntry) SetIsIncludeDocuments(v Boolean) error {
	return putComProperty(r, "IsIncludeDocuments", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISINCLUDEFORMS_PROPERTY_RE.html */
func (r NotesReplicationEntry) IsIncludeForms() (Boolean, error) {
	return getComProperty[Boolean](r, "IsIncludeForms")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISINCLUDEFORMS_PROPERTY_RE.html */
func (r NotesReplicationEntry) SetIsIncludeForms(v Boolean) error {
	return putComProperty(r, "IsIncludeForms", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISINCLUDEFORMULAS_PROPERTY_RE.html */
func (r NotesReplicationEntry) IsIncludeFormulas() (Boolean, error) {
	return getComProperty[Boolean](r, "IsIncludeFormulas")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISINCLUDEFORMULAS_PROPERTY_RE.html */
func (r NotesReplicationEntry) SetIsIncludeFormulas(v Boolean) error {
	return putComProperty(r, "IsIncludeFormulas", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SOURCE_PROPERTY_RE.html */
func (r NotesReplicationEntry) Source() (String, error) {
	return getComProperty[String](r, "Source")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_VIEWS_PROPERTY_RE.html */
func (r NotesReplicationEntry) Views() (String, error) {
	return getComProperty[String](r, "Views")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_VIEWS_PROPERTY_RE.html */
func (r NotesReplicationEntry) SetViews(v String) error {
	return putComProperty(r, "Views", v)
}

/* --------------------------------- Methods ------------------------------------ */
/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_REMOVE_METHOD_RE.html */
func (r NotesReplicationEntry) Remove() error {
	return callComVoidMethod(r, "Remove")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SAVE_METHOD_RE.html */
func (r NotesReplicationEntry) Save() error {
	return callComVoidMethod(r, "Save")
}
