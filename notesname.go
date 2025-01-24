/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_NOTESNAME_CLASS.html */
package domigo

import (
	ole "github.com/go-ole/go-ole"
)

type NotesName struct {
	NotesStruct
}

func newNotesName(dispatchPtr *ole.IDispatch) NotesName {
	return NotesName{newNotesStruct(dispatchPtr)}
}

/* --------------------------------- Properties --------------------------------- */
/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ABBREVIATED_PROPERTY.html */
func (n NotesName) Abbreviated() (String, error) {
	return getComProperty[String](n, "Abbreviated")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ADDR821_PROPERTY_6733.html */
func (n NotesName) Addr821() (String, error) {
	return getComProperty[String](n, "Addr821")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ADDR822COMMENT1_PROPERTY_7368.html */
func (n NotesName) Addr822Comment1() (String, error) {
	return getComProperty[String](n, "Addr822Comment1")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ADDR822COMMENT2_PROPERTY_8387.html */
func (n NotesName) Addr822Comment2() (String, error) {
	return getComProperty[String](n, "Addr822Comment2")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ADDR822COMMENT3_PROPERTY_5250.html */
func (n NotesName) Addr822Comment3() (String, error) {
	return getComProperty[String](n, "Addr822Comment3")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ADDR822LOCALPART_PROPERTY_2191.html */
func (n NotesName) Addr822LocalPart() (String, error) {
	return getComProperty[String](n, "Addr822LocalPart")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ADDR822PHRASE_PROPERTY_7159.html */
func (n NotesName) Addr822Phrase() (String, error) {
	return getComProperty[String](n, "Addr822Phrase")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ADMD_PROPERTY.html */
func (n NotesName) ADMD() (String, error) {
	return getComProperty[String](n, "ADMD")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_CANONICAL_PROPERTY.html */
func (n NotesName) Canonical() (String, error) {
	return getComProperty[String](n, "Canonical")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_COMMON_PROPERTY.html */
func (n NotesName) Common() (String, error) {
	return getComProperty[String](n, "Common")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_COUNTRY_PROPERTY.html */
func (n NotesName) Country() (String, error) {
	return getComProperty[String](n, "Country")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GENERATION_PROPERTY.html */
func (n NotesName) Generation() (String, error) {
	return getComProperty[String](n, "Generation")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GIVEN_PROPERTY.html */
func (n NotesName) Given() (String, error) {
	return getComProperty[String](n, "Given")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_INITIALS_PROPERTY.html */
func (n NotesName) Initials() (String, error) {
	return getComProperty[String](n, "Initials")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISHIERARCHICAL_PROPERTY.html */
func (n NotesName) IsHierarchical() (Boolean, error) {
	return getComProperty[Boolean](n, "IsHierarchical")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_KEYWORD_PROPERTY.html */
func (n NotesName) Keyword() (String, error) {
	return getComProperty[String](n, "Keyword")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_LANGUAGE_PROPERTY_123.html */
func (n NotesName) Language() (String, error) {
	return getComProperty[String](n, "Language")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ORGANIZATION_PROPERTY.html */
func (n NotesName) Organization() (String, error) {
	return getComProperty[String](n, "Organization")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ORGUNIT1_PROPERTY.html */
func (n NotesName) OrgUnit1() (String, error) {
	return getComProperty[String](n, "OrgUnit1")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ORGUNIT2_PROPERTY.html */
func (n NotesName) OrgUnit2() (String, error) {
	return getComProperty[String](n, "OrgUnit2")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ORGUNIT3_PROPERTY.html */
func (n NotesName) OrgUnit3() (String, error) {
	return getComProperty[String](n, "OrgUnit3")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ORGUNIT4_PROPERTY.html */
func (n NotesName) OrgUnit4() (String, error) {
	return getComProperty[String](n, "OrgUnit4")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_PARENT_PROPERTY_NAME_COM.html */
func (n NotesName) Parent() (NotesSession, error) {
	return getComObjectProperty(n, newNotesSession, "Parent")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_PRMD_PROPERTY.html */
func (n NotesName) PRMD() (String, error) {
	return getComProperty[String](n, "PRMD")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SURNAME_PROPERTY.html */
func (n NotesName) Surname() (String, error) {
	return getComProperty[String](n, "Surname")
}

/* --------------------------------- Methods ------------------------------------ */
