/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_NOTESNEWSLETTER_CLASS.html */
package notesnewsletter

import (
	"github.com/monstermichl/domigo/domino"
	"github.com/monstermichl/domigo/domino/notesdatabase"
	"github.com/monstermichl/domigo/domino/notesdocument"
	"github.com/monstermichl/domigo/helpers"

	ole "github.com/go-ole/go-ole"
)

type NotesNewsletter struct {
	domino.NotesStruct
}

func New(dispatchPtr *ole.IDispatch) NotesNewsletter {
	return NotesNewsletter{domino.NewNotesStruct(dispatchPtr)}
}

/* --------------------------------- Properties --------------------------------- */
/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_DOSCORE_PROPERTY.html */
func (n NotesNewsletter) DoScore() (domino.Boolean, error) {
	val, err := n.Com().GetProperty("DoScore")
	return helpers.CastValue[domino.Boolean](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_DOSCORE_PROPERTY.html */
func (n NotesNewsletter) SetDoScore(v domino.Boolean) error {
	return n.Com().PutProperty("DoScore", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_DOSUBJECT_PROPERTY.html */
func (n NotesNewsletter) DoSubject() (domino.Boolean, error) {
	val, err := n.Com().GetProperty("DoSubject")
	return helpers.CastValue[domino.Boolean](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_DOSUBJECT_PROPERTY.html */
func (n NotesNewsletter) SetDoSubject(v domino.Boolean) error {
	return n.Com().PutProperty("DoSubject", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SUBJECTITEMNAME_PROPERTY.html */
func (n NotesNewsletter) SubjectItemName() (domino.String, error) {
	val, err := n.Com().GetProperty("SubjectItemName")
	return helpers.CastValue[domino.String](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SUBJECTITEMNAME_PROPERTY.html */
func (n NotesNewsletter) SetSubjectItemName(v domino.String) error {
	return n.Com().PutProperty("SubjectItemName", v)
}

/* --------------------------------- Methods ------------------------------------ */
/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_FORMATDOCUMENT_METHOD.html */
func (n NotesNewsletter) FormatDocument(notesDatabase notesdatabase.NotesDatabase, documentNumber domino.Integer) (notesdocument.NotesDocument, error) {
	dispatchPtr, err := n.Com().CallObjectMethod("FormatDocument", notesDatabase.Com().Dispatch(), documentNumber)
	return notesdocument.New(dispatchPtr), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_FORMATMSGWITHDOCLINKS_METHOD.html */
func (n NotesNewsletter) FormatMsgWithDoclinks(notesDatabase notesdatabase.NotesDatabase) (notesdocument.NotesDocument, error) {
	dispatchPtr, err := n.Com().CallObjectMethod("FormatMsgWithDoclinks", notesDatabase.Com().Dispatch())
	return notesdocument.New(dispatchPtr), err
}
