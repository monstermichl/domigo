/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_NOTESNAME_CLASS.html */
package notesname

import (
	"github.com/monstermichl/domigo/domino"
	"github.com/monstermichl/domigo/helpers"

	ole "github.com/go-ole/go-ole"
)

type NotesName struct {
	domino.NotesStruct
}

func New(dispatchPtr *ole.IDispatch) NotesName {
	return NotesName{domino.NewNotesStruct(dispatchPtr)}
}

/* --------------------------------- Properties --------------------------------- */
func (n NotesName) Abbreviated() (domino.String, error) {
	val, err := n.Com().GetProperty("Abbreviated")
	return helpers.CastValue[domino.String](val), err
}

func (n NotesName) Addr821() (domino.String, error) {
	val, err := n.Com().GetProperty("Addr821")
	return helpers.CastValue[domino.String](val), err
}

func (n NotesName) Addr822Comment1() (domino.String, error) {
	val, err := n.Com().GetProperty("Addr822Comment1")
	return helpers.CastValue[domino.String](val), err
}

func (n NotesName) Addr822Comment2() (domino.String, error) {
	val, err := n.Com().GetProperty("Addr822Comment2")
	return helpers.CastValue[domino.String](val), err
}

func (n NotesName) Addr822Comment3() (domino.String, error) {
	val, err := n.Com().GetProperty("Addr822Comment3")
	return helpers.CastValue[domino.String](val), err
}

func (n NotesName) Addr822LocalPart() (domino.String, error) {
	val, err := n.Com().GetProperty("Addr822LocalPart")
	return helpers.CastValue[domino.String](val), err
}

func (n NotesName) Addr822Phrase() (domino.String, error) {
	val, err := n.Com().GetProperty("Addr822Phrase")
	return helpers.CastValue[domino.String](val), err
}

func (n NotesName) ADMD() (domino.String, error) {
	val, err := n.Com().GetProperty("ADMD")
	return helpers.CastValue[domino.String](val), err
}

func (n NotesName) Canonical() (domino.String, error) {
	val, err := n.Com().GetProperty("Canonical")
	return helpers.CastValue[domino.String](val), err
}

func (n NotesName) Common() (domino.String, error) {
	val, err := n.Com().GetProperty("Common")
	return helpers.CastValue[domino.String](val), err
}

func (n NotesName) Country() (domino.String, error) {
	val, err := n.Com().GetProperty("Country")
	return helpers.CastValue[domino.String](val), err
}

func (n NotesName) Generation() (domino.String, error) {
	val, err := n.Com().GetProperty("Generation")
	return helpers.CastValue[domino.String](val), err
}

func (n NotesName) Given() (domino.String, error) {
	val, err := n.Com().GetProperty("Given")
	return helpers.CastValue[domino.String](val), err
}

func (n NotesName) Initials() (domino.String, error) {
	val, err := n.Com().GetProperty("Initials")
	return helpers.CastValue[domino.String](val), err
}

func (n NotesName) IsHierarchical() (domino.Boolean, error) {
	val, err := n.Com().GetProperty("IsHierarchical")
	return helpers.CastValue[domino.Boolean](val), err
}

func (n NotesName) Keyword() (domino.String, error) {
	val, err := n.Com().GetProperty("Keyword")
	return helpers.CastValue[domino.String](val), err
}

func (n NotesName) Language() (domino.String, error) {
	val, err := n.Com().GetProperty("Language")
	return helpers.CastValue[domino.String](val), err
}

func (n NotesName) Organization() (domino.String, error) {
	val, err := n.Com().GetProperty("Organization")
	return helpers.CastValue[domino.String](val), err
}

func (n NotesName) OrgUnit1() (domino.String, error) {
	val, err := n.Com().GetProperty("OrgUnit1")
	return helpers.CastValue[domino.String](val), err
}

func (n NotesName) OrgUnit2() (domino.String, error) {
	val, err := n.Com().GetProperty("OrgUnit2")
	return helpers.CastValue[domino.String](val), err
}

func (n NotesName) OrgUnit3() (domino.String, error) {
	val, err := n.Com().GetProperty("OrgUnit3")
	return helpers.CastValue[domino.String](val), err
}

func (n NotesName) OrgUnit4() (domino.String, error) {
	val, err := n.Com().GetProperty("OrgUnit4")
	return helpers.CastValue[domino.String](val), err
}

func (n NotesName) PRMD() (domino.String, error) {
	val, err := n.Com().GetProperty("PRMD")
	return helpers.CastValue[domino.String](val), err
}

func (n NotesName) Surname() (domino.String, error) {
	val, err := n.Com().GetProperty("Surname")
	return helpers.CastValue[domino.String](val), err
}

/* --------------------------------- Methods ------------------------------------ */
