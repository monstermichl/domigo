/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_NOTESNAME_CLASS.html */
package domigo

import (
	"github.com/monstermichl/domigo/internal/helpers"

	ole "github.com/go-ole/go-ole"
)

type NotesName struct {
	NotesStruct
}

func NewNotesName(dispatchPtr *ole.IDispatch) NotesName {
	return NotesName{NewNotesStruct(dispatchPtr)}
}

/* --------------------------------- Properties --------------------------------- */
/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ABBREVIATED_PROPERTY.html */
func (n NotesName) Abbreviated() (String, error) {
	val, err := n.com().GetProperty("Abbreviated")
	return helpers.CastValue[String](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ADDR821_PROPERTY_6733.html */
func (n NotesName) Addr821() (String, error) {
	val, err := n.com().GetProperty("Addr821")
	return helpers.CastValue[String](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ADDR822COMMENT1_PROPERTY_7368.html */
func (n NotesName) Addr822Comment1() (String, error) {
	val, err := n.com().GetProperty("Addr822Comment1")
	return helpers.CastValue[String](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ADDR822COMMENT2_PROPERTY_8387.html */
func (n NotesName) Addr822Comment2() (String, error) {
	val, err := n.com().GetProperty("Addr822Comment2")
	return helpers.CastValue[String](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ADDR822COMMENT3_PROPERTY_5250.html */
func (n NotesName) Addr822Comment3() (String, error) {
	val, err := n.com().GetProperty("Addr822Comment3")
	return helpers.CastValue[String](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ADDR822LOCALPART_PROPERTY_2191.html */
func (n NotesName) Addr822LocalPart() (String, error) {
	val, err := n.com().GetProperty("Addr822LocalPart")
	return helpers.CastValue[String](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ADDR822PHRASE_PROPERTY_7159.html */
func (n NotesName) Addr822Phrase() (String, error) {
	val, err := n.com().GetProperty("Addr822Phrase")
	return helpers.CastValue[String](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ADMD_PROPERTY.html */
func (n NotesName) ADMD() (String, error) {
	val, err := n.com().GetProperty("ADMD")
	return helpers.CastValue[String](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_CANONICAL_PROPERTY.html */
func (n NotesName) Canonical() (String, error) {
	val, err := n.com().GetProperty("Canonical")
	return helpers.CastValue[String](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_COMMON_PROPERTY.html */
func (n NotesName) Common() (String, error) {
	val, err := n.com().GetProperty("Common")
	return helpers.CastValue[String](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_COUNTRY_PROPERTY.html */
func (n NotesName) Country() (String, error) {
	val, err := n.com().GetProperty("Country")
	return helpers.CastValue[String](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GENERATION_PROPERTY.html */
func (n NotesName) Generation() (String, error) {
	val, err := n.com().GetProperty("Generation")
	return helpers.CastValue[String](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GIVEN_PROPERTY.html */
func (n NotesName) Given() (String, error) {
	val, err := n.com().GetProperty("Given")
	return helpers.CastValue[String](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_INITIALS_PROPERTY.html */
func (n NotesName) Initials() (String, error) {
	val, err := n.com().GetProperty("Initials")
	return helpers.CastValue[String](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISHIERARCHICAL_PROPERTY.html */
func (n NotesName) IsHierarchical() (Boolean, error) {
	val, err := n.com().GetProperty("IsHierarchical")
	return helpers.CastValue[Boolean](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_KEYWORD_PROPERTY.html */
func (n NotesName) Keyword() (String, error) {
	val, err := n.com().GetProperty("Keyword")
	return helpers.CastValue[String](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_LANGUAGE_PROPERTY_123.html */
func (n NotesName) Language() (String, error) {
	val, err := n.com().GetProperty("Language")
	return helpers.CastValue[String](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ORGANIZATION_PROPERTY.html */
func (n NotesName) Organization() (String, error) {
	val, err := n.com().GetProperty("Organization")
	return helpers.CastValue[String](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ORGUNIT1_PROPERTY.html */
func (n NotesName) OrgUnit1() (String, error) {
	val, err := n.com().GetProperty("OrgUnit1")
	return helpers.CastValue[String](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ORGUNIT2_PROPERTY.html */
func (n NotesName) OrgUnit2() (String, error) {
	val, err := n.com().GetProperty("OrgUnit2")
	return helpers.CastValue[String](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ORGUNIT3_PROPERTY.html */
func (n NotesName) OrgUnit3() (String, error) {
	val, err := n.com().GetProperty("OrgUnit3")
	return helpers.CastValue[String](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ORGUNIT4_PROPERTY.html */
func (n NotesName) OrgUnit4() (String, error) {
	val, err := n.com().GetProperty("OrgUnit4")
	return helpers.CastValue[String](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_PARENT_PROPERTY_NAME_COM.html */
func (n NotesName) Parent() (NotesSession, error) {
	dispatchPtr, err := n.com().GetObjectProperty("Parent")
	return NewNotesSession(dispatchPtr), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_PRMD_PROPERTY.html */
func (n NotesName) PRMD() (String, error) {
	val, err := n.com().GetProperty("PRMD")
	return helpers.CastValue[String](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SURNAME_PROPERTY.html */
func (n NotesName) Surname() (String, error) {
	val, err := n.com().GetProperty("Surname")
	return helpers.CastValue[String](val), err
}

/* --------------------------------- Methods ------------------------------------ */
