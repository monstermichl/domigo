/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_NOTESEMBEDDEDOBJECT_CLASS.html */
package domigo

import (
	"github.com/monstermichl/domigo/internal/helpers"

	ole "github.com/go-ole/go-ole"
)

type NotesEmbeddedObjectEmbedType Long

const (
	NOTESEMBEDDEDOBJECT_EMBED_ATTACHMENT NotesEmbeddedObjectEmbedType = 1454
	NOTESEMBEDDEDOBJECT_EMBED_OBJECT     NotesEmbeddedObjectEmbedType = 1453
	NOTESEMBEDDEDOBJECT_EMBED_OBJECTLINK NotesEmbeddedObjectEmbedType = 1452
)

type NotesEmbeddedObject struct {
	NotesStruct
}

func NewNotesEmbeddedObject(dispatchPtr *ole.IDispatch) NotesEmbeddedObject {
	return NotesEmbeddedObject{NewNotesStruct(dispatchPtr)}
}

/* --------------------------------- Properties --------------------------------- */
/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_CLASS_PROPERTY.html */
func (e NotesEmbeddedObject) Class() (String, error) {
	val, err := e.com().GetProperty("Class")
	return helpers.CastValue[String](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_FILECREATED_PROPERTY.html */
func (e NotesEmbeddedObject) FileCreated() (String, error) {
	val, err := e.com().GetProperty("FileCreated")
	return helpers.CastValue[String](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_FILEENCODING_PROPERTY.html */
func (e NotesEmbeddedObject) FileEncoding() (Integer, error) {
	val, err := e.com().GetProperty("FileEncoding")
	return helpers.CastValue[Integer](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_FILEMODIFIED_PROPERTY.html */
func (e NotesEmbeddedObject) FileModified() (String, error) {
	val, err := e.com().GetProperty("FileModified")
	return helpers.CastValue[String](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_FILESIZE_PROPERTY.html */
func (e NotesEmbeddedObject) FileSize() (Long, error) {
	val, err := e.com().GetProperty("FileSize")
	return helpers.CastValue[Long](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_FITBELOWFIELDS_PROPERTY_EMBEDDEDOBJ.html */
func (e NotesEmbeddedObject) FitBelowFields() (Boolean, error) {
	val, err := e.com().GetProperty("FitBelowFields")
	return helpers.CastValue[Boolean](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_FITBELOWFIELDS_PROPERTY_EMBEDDEDOBJ.html */
func (e NotesEmbeddedObject) SetFitBelowFields(v Boolean) error {
	return e.com().PutProperty("FitBelowFields", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_FITTOWINDOW_PROPERTY_EMBEDDEDOBJ.html */
func (e NotesEmbeddedObject) FitToWindow() (Boolean, error) {
	val, err := e.com().GetProperty("FitToWindow")
	return helpers.CastValue[Boolean](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_FITTOWINDOW_PROPERTY_EMBEDDEDOBJ.html */
func (e NotesEmbeddedObject) SetFitToWindow(v Boolean) error {
	return e.com().PutProperty("FitToWindow", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_NAME_PROPERTY_OBJECT.html */
func (e NotesEmbeddedObject) Name() (String, error) {
	val, err := e.com().GetProperty("Name")
	return helpers.CastValue[String](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_OBJECT_PROPERTY.html */
func (e NotesEmbeddedObject) Object() (Variant, error) {
	val, err := e.com().GetProperty("Object")
	return helpers.CastValue[Variant](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_PARENT_PROPERTY_OBJECT.html */
func (e NotesEmbeddedObject) Parent() (NotesRichTextItem, error) {
	dispatchPtr, err := e.com().GetObjectProperty("Parent")
	return NewNotesRichTextItem(dispatchPtr), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_RUNREADONLY_PROPERTY_EMBEDDEDOBJ.html */
func (e NotesEmbeddedObject) RunReadOnly() (Boolean, error) {
	val, err := e.com().GetProperty("RunReadOnly")
	return helpers.CastValue[Boolean](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_RUNREADONLY_PROPERTY_EMBEDDEDOBJ.html */
func (e NotesEmbeddedObject) SetRunReadOnly(v Boolean) error {
	return e.com().PutProperty("RunReadOnly", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SOURCE_PROPERTY.html */
func (e NotesEmbeddedObject) Source() (String, error) {
	val, err := e.com().GetProperty("Source")
	return helpers.CastValue[String](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_TYPE_PROPERTY_OBJECT.html */
func (e NotesEmbeddedObject) Type() (NotesEmbeddedObjectEmbedType, error) {
	val, err := e.com().GetProperty("Type")
	return helpers.CastValue[NotesEmbeddedObjectEmbedType](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_VERBS_PROPERTY.html */
func (e NotesEmbeddedObject) Verbs() ([]String, error) {
	vals, err := e.com().GetArrayProperty("Verbs")
	return helpers.CastSlice[String](vals), err
}

/* --------------------------------- Methods ------------------------------------ */
/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ACTIVATE_METHOD.html */
func (e NotesEmbeddedObject) Activate(show Boolean) (*ole.IDispatch, error) {
	return e.com().CallObjectMethod("Activate", show)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_DOVERB_METHOD.html */
func (e NotesEmbeddedObject) DoVerb(verb String) (Variant, error) {
	val, err := e.com().CallMethod("DoVerb", verb)
	return helpers.CastValue[Variant](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_EXTRACTFILE_METHOD.html */
func (e NotesEmbeddedObject) ExtractFile(path String) error {
	_, err := e.com().CallMethod("ExtractFile", path)
	return err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_REMOVE_METHOD_OBJECT.html */
func (e NotesEmbeddedObject) Remove() error {
	_, err := e.com().CallMethod("Remove")
	return err
}
