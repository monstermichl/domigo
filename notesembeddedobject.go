/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_NOTESEMBEDDEDOBJECT_CLASS.html */
package domigo

import (
	"github.com/monstermichl/domigo/internal/helpers"

	ole "github.com/go-ole/go-ole"
)

type NotesEmbeddedObjectEmbedType = Long

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
	val, err := getComProperty(e, "Class")
	return helpers.CastValue[String](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_FILECREATED_PROPERTY.html */
func (e NotesEmbeddedObject) FileCreated() (String, error) {
	val, err := getComProperty(e, "FileCreated")
	return helpers.CastValue[String](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_FILEENCODING_PROPERTY.html */
func (e NotesEmbeddedObject) FileEncoding() (Integer, error) {
	val, err := getComProperty(e, "FileEncoding")
	return helpers.CastValue[Integer](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_FILEMODIFIED_PROPERTY.html */
func (e NotesEmbeddedObject) FileModified() (String, error) {
	val, err := getComProperty(e, "FileModified")
	return helpers.CastValue[String](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_FILESIZE_PROPERTY.html */
func (e NotesEmbeddedObject) FileSize() (Long, error) {
	val, err := getComProperty(e, "FileSize")
	return helpers.CastValue[Long](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_FITBELOWFIELDS_PROPERTY_EMBEDDEDOBJ.html */
func (e NotesEmbeddedObject) FitBelowFields() (Boolean, error) {
	val, err := getComProperty(e, "FitBelowFields")
	return helpers.CastValue[Boolean](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_FITBELOWFIELDS_PROPERTY_EMBEDDEDOBJ.html */
func (e NotesEmbeddedObject) SetFitBelowFields(v Boolean) error {
	return putComProperty(e, "FitBelowFields", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_FITTOWINDOW_PROPERTY_EMBEDDEDOBJ.html */
func (e NotesEmbeddedObject) FitToWindow() (Boolean, error) {
	val, err := getComProperty(e, "FitToWindow")
	return helpers.CastValue[Boolean](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_FITTOWINDOW_PROPERTY_EMBEDDEDOBJ.html */
func (e NotesEmbeddedObject) SetFitToWindow(v Boolean) error {
	return putComProperty(e, "FitToWindow", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_NAME_PROPERTY_OBJECT.html */
func (e NotesEmbeddedObject) Name() (String, error) {
	val, err := getComProperty(e, "Name")
	return helpers.CastValue[String](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_OBJECT_PROPERTY.html */
func (e NotesEmbeddedObject) Object() (Variant, error) {
	val, err := getComProperty(e, "Object")
	return helpers.CastValue[Variant](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_PARENT_PROPERTY_OBJECT.html */
func (e NotesEmbeddedObject) Parent() (NotesRichTextItem, error) {
	dispatchPtr, err := getComObjectProperty(e, "Parent")
	return NewNotesRichTextItem(dispatchPtr), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_RUNREADONLY_PROPERTY_EMBEDDEDOBJ.html */
func (e NotesEmbeddedObject) RunReadOnly() (Boolean, error) {
	val, err := getComProperty(e, "RunReadOnly")
	return helpers.CastValue[Boolean](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_RUNREADONLY_PROPERTY_EMBEDDEDOBJ.html */
func (e NotesEmbeddedObject) SetRunReadOnly(v Boolean) error {
	return putComProperty(e, "RunReadOnly", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SOURCE_PROPERTY.html */
func (e NotesEmbeddedObject) Source() (String, error) {
	val, err := getComProperty(e, "Source")
	return helpers.CastValue[String](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_TYPE_PROPERTY_OBJECT.html */
func (e NotesEmbeddedObject) Type() (NotesEmbeddedObjectEmbedType, error) {
	val, err := getComProperty(e, "Type")
	return helpers.CastValue[NotesEmbeddedObjectEmbedType](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_VERBS_PROPERTY.html */
func (e NotesEmbeddedObject) Verbs() ([]String, error) {
	vals, err := getComArrayProperty(e, "Verbs")
	return helpers.CastSlice[String](vals), err
}

/* --------------------------------- Methods ------------------------------------ */
/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ACTIVATE_METHOD.html */
func (e NotesEmbeddedObject) Activate(show Boolean) (*ole.IDispatch, error) {
	return callComObjectMethod(e, "Activate", show)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_DOVERB_METHOD.html */
func (e NotesEmbeddedObject) DoVerb(verb String) (Variant, error) {
	val, err := callComMethod(e, "DoVerb", verb)
	return helpers.CastValue[Variant](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_EXTRACTFILE_METHOD.html */
func (e NotesEmbeddedObject) ExtractFile(path String) error {
	_, err := callComMethod(e, "ExtractFile", path)
	return err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_REMOVE_METHOD_OBJECT.html */
func (e NotesEmbeddedObject) Remove() error {
	_, err := callComMethod(e, "Remove")
	return err
}
