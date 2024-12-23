/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_NOTESNEWSLETTER_CLASS.html */
package domigo

import (
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
	return getComProperty[Boolean](n, "DoScore")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_DOSCORE_PROPERTY.html */
func (n NotesNewsletter) SetDoScore(v Boolean) error {
	return putComProperty(n, "DoScore", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_DOSUBJECT_PROPERTY.html */
func (n NotesNewsletter) DoSubject() (Boolean, error) {
	return getComProperty[Boolean](n, "DoSubject")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_DOSUBJECT_PROPERTY.html */
func (n NotesNewsletter) SetDoSubject(v Boolean) error {
	return putComProperty(n, "DoSubject", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_PARENT_PROPERTY_NEWSLETTER_COM.html */
func (n NotesNewsletter) Parent() (NotesSession, error) {
	return getComObjectProperty(n, NewNotesSession, "Parent")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SUBJECTITEMNAME_PROPERTY.html */
func (n NotesNewsletter) SubjectItemName() (String, error) {
	return getComProperty[String](n, "SubjectItemName")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SUBJECTITEMNAME_PROPERTY.html */
func (n NotesNewsletter) SetSubjectItemName(v String) error {
	return putComProperty(n, "SubjectItemName", v)
}

/* --------------------------------- Methods ------------------------------------ */
/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_FORMATDOCUMENT_METHOD.html */
func (n NotesNewsletter) FormatDocument(notesDatabase NotesDatabase, documentNumber Integer) (NotesDocument, error) {
	return callComObjectMethod(n, NewNotesDocument, "FormatDocument", notesDatabase.com().Dispatch(), documentNumber)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_FORMATMSGWITHDOCLINKS_METHOD.html */
func (n NotesNewsletter) FormatMsgWithDoclinks(notesDatabase NotesDatabase) (NotesDocument, error) {
	return callComObjectMethod(n, NewNotesDocument, "FormatMsgWithDoclinks", notesDatabase.com().Dispatch())
}
