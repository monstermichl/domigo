/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_NOTESEMBEDDEDOBJECT_CLASS.html */
package notesembeddedobject

import (
	"github.com/monstermichl/domigo/domino"
	"github.com/monstermichl/domigo/helpers"

	ole "github.com/go-ole/go-ole"
)

type EmbedType = domino.Long

const (
	EMBED_ATTACHMENT EmbedType = 1454
	EMBED_OBJECT     EmbedType = 1453
	EMBED_OBJECTLINK EmbedType = 1452
)

type NotesEmbeddedObject struct {
	domino.NotesStruct
}

func New(dispatchPtr *ole.IDispatch) NotesEmbeddedObject {
	return NotesEmbeddedObject{domino.NewNotesStruct(dispatchPtr)}
}

/* --------------------------------- Properties --------------------------------- */
/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_CLASS_PROPERTY.html */
func (e NotesEmbeddedObject) Class() (domino.String, error) {
	val, err := e.Com().GetProperty("Class")
	return helpers.CastValue[domino.String](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_FILECREATED_PROPERTY.html */
func (e NotesEmbeddedObject) FileCreated() (domino.String, error) {
	val, err := e.Com().GetProperty("FileCreated")
	return helpers.CastValue[domino.String](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_FILEENCODING_PROPERTY.html */
func (e NotesEmbeddedObject) FileEncoding() (domino.Integer, error) {
	val, err := e.Com().GetProperty("FileEncoding")
	return helpers.CastValue[domino.Integer](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_FILEMODIFIED_PROPERTY.html */
func (e NotesEmbeddedObject) FileModified() (domino.String, error) {
	val, err := e.Com().GetProperty("FileModified")
	return helpers.CastValue[domino.String](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_FILESIZE_PROPERTY.html */
func (e NotesEmbeddedObject) FileSize() (domino.Long, error) {
	val, err := e.Com().GetProperty("FileSize")
	return helpers.CastValue[domino.Long](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_FITBELOWFIELDS_PROPERTY_EMBEDDEDOBJ.html */
func (e NotesEmbeddedObject) FitBelowFields() (domino.Boolean, error) {
	val, err := e.Com().GetProperty("FitBelowFields")
	return helpers.CastValue[domino.Boolean](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_FITBELOWFIELDS_PROPERTY_EMBEDDEDOBJ.html */
func (e NotesEmbeddedObject) SetFitBelowFields(v domino.Boolean) error {
	return e.Com().PutProperty("FitBelowFields", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_FITTOWINDOW_PROPERTY_EMBEDDEDOBJ.html */
func (e NotesEmbeddedObject) FitToWindow() (domino.Boolean, error) {
	val, err := e.Com().GetProperty("FitToWindow")
	return helpers.CastValue[domino.Boolean](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_FITTOWINDOW_PROPERTY_EMBEDDEDOBJ.html */
func (e NotesEmbeddedObject) SetFitToWindow(v domino.Boolean) error {
	return e.Com().PutProperty("FitToWindow", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_NAME_PROPERTY_OBJECT.html */
func (e NotesEmbeddedObject) Name() (domino.String, error) {
	val, err := e.Com().GetProperty("Name")
	return helpers.CastValue[domino.String](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_OBJECT_PROPERTY.html */
func (e NotesEmbeddedObject) Object() (domino.Variant, error) {
	val, err := e.Com().GetProperty("Object")
	return helpers.CastValue[domino.Variant](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_RUNREADONLY_PROPERTY_EMBEDDEDOBJ.html */
func (e NotesEmbeddedObject) RunReadOnly() (domino.Boolean, error) {
	val, err := e.Com().GetProperty("RunReadOnly")
	return helpers.CastValue[domino.Boolean](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_RUNREADONLY_PROPERTY_EMBEDDEDOBJ.html */
func (e NotesEmbeddedObject) SetRunReadOnly(v domino.Boolean) error {
	return e.Com().PutProperty("RunReadOnly", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SOURCE_PROPERTY.html */
func (e NotesEmbeddedObject) Source() (domino.String, error) {
	val, err := e.Com().GetProperty("Source")
	return helpers.CastValue[domino.String](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_TYPE_PROPERTY_OBJECT.html */
func (e NotesEmbeddedObject) Type() (EmbedType, error) {
	val, err := e.Com().GetProperty("Type")
	return helpers.CastValue[EmbedType](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_VERBS_PROPERTY.html */
func (e NotesEmbeddedObject) Verbs() ([]domino.String, error) {
	vals, err := e.Com().GetArrayProperty("Verbs")
	return helpers.CastSlice[domino.String](vals), err
}

/* --------------------------------- Methods ------------------------------------ */
/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ACTIVATE_METHOD.html */
func (e NotesEmbeddedObject) Activate(show domino.Boolean) (*ole.IDispatch, error) {
	return e.Com().CallObjectMethod("Activate", show)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_DOVERB_METHOD.html */
func (e NotesEmbeddedObject) DoVerb(verb domino.String) (domino.Variant, error) {
	val, err := e.Com().CallMethod("DoVerb", verb)
	return helpers.CastValue[domino.Variant](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_EXTRACTFILE_METHOD.html */
func (e NotesEmbeddedObject) ExtractFile(path domino.String) error {
	_, err := e.Com().CallMethod("ExtractFile", path)
	return err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_REMOVE_METHOD_OBJECT.html */
func (e NotesEmbeddedObject) Remove() error {
	_, err := e.Com().CallMethod("Remove")
	return err
}
