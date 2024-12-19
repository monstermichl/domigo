/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_NOTESNEWSLETTER_CLASS.html */
package domigo

import (
	"github.com/monstermichl/domigo/internal/helpers"

	ole "github.com/go-ole/go-ole"
)

type NotesNewsletter struct {
	NotesStruct
}

func NewNotesNewsletter(dispatchPtr *ole.IDispatch) NotesNewsletter {
	return NotesNewsletter{NewNotesStruct(dispatchPtr)}
}

/* --------------------------------- Properties --------------------------------- */
/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_DOSCORE_PROPERTY.html */
func (n NotesNewsletter) DoScore() (Boolean, error) {
	val, err := n.com().GetProperty("DoScore")
	return helpers.CastValue[Boolean](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_DOSCORE_PROPERTY.html */
func (n NotesNewsletter) SetDoScore(v Boolean) error {
	return n.com().PutProperty("DoScore", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_DOSUBJECT_PROPERTY.html */
func (n NotesNewsletter) DoSubject() (Boolean, error) {
	val, err := n.com().GetProperty("DoSubject")
	return helpers.CastValue[Boolean](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_DOSUBJECT_PROPERTY.html */
func (n NotesNewsletter) SetDoSubject(v Boolean) error {
	return n.com().PutProperty("DoSubject", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_PARENT_PROPERTY_NEWSLETTER_COM.html */
func (n NotesNewsletter) Parent() (NotesSession, error) {
	dispatchPtr, err := n.com().GetObjectProperty("Parent")
	return NewNotesSession(dispatchPtr), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SUBJECTITEMNAME_PROPERTY.html */
func (n NotesNewsletter) SubjectItemName() (String, error) {
	val, err := n.com().GetProperty("SubjectItemName")
	return helpers.CastValue[String](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SUBJECTITEMNAME_PROPERTY.html */
func (n NotesNewsletter) SetSubjectItemName(v String) error {
	return n.com().PutProperty("SubjectItemName", v)
}

/* --------------------------------- Methods ------------------------------------ */
/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_FORMATDOCUMENT_METHOD.html */
func (n NotesNewsletter) FormatDocument(notesDatabase NotesDatabase, documentNumber Integer) (NotesDocument, error) {
	dispatchPtr, err := n.com().CallObjectMethod("FormatDocument", notesDatabase.com().Dispatch(), documentNumber)
	return NewNotesDocument(dispatchPtr), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_FORMATMSGWITHDOCLINKS_METHOD.html */
func (n NotesNewsletter) FormatMsgWithDoclinks(notesDatabase NotesDatabase) (NotesDocument, error) {
	dispatchPtr, err := n.com().CallObjectMethod("FormatMsgWithDoclinks", notesDatabase.com().Dispatch())
	return NewNotesDocument(dispatchPtr), err
}
