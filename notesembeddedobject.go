/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_NOTESEMBEDDEDOBJECT_CLASS.html */
package domigo

import (
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
	return getComProperty[String](e, "Class")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_FILECREATED_PROPERTY.html */
func (e NotesEmbeddedObject) FileCreated() (String, error) {
	return getComProperty[String](e, "FileCreated")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_FILEENCODING_PROPERTY.html */
func (e NotesEmbeddedObject) FileEncoding() (Integer, error) {
	return getComProperty[Integer](e, "FileEncoding")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_FILEMODIFIED_PROPERTY.html */
func (e NotesEmbeddedObject) FileModified() (String, error) {
	return getComProperty[String](e, "FileModified")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_FILESIZE_PROPERTY.html */
func (e NotesEmbeddedObject) FileSize() (Long, error) {
	return getComProperty[Long](e, "FileSize")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_FITBELOWFIELDS_PROPERTY_EMBEDDEDOBJ.html */
func (e NotesEmbeddedObject) FitBelowFields() (Boolean, error) {
	return getComProperty[Boolean](e, "FitBelowFields")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_FITBELOWFIELDS_PROPERTY_EMBEDDEDOBJ.html */
func (e NotesEmbeddedObject) SetFitBelowFields(v Boolean) error {
	return putComProperty(e, "FitBelowFields", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_FITTOWINDOW_PROPERTY_EMBEDDEDOBJ.html */
func (e NotesEmbeddedObject) FitToWindow() (Boolean, error) {
	return getComProperty[Boolean](e, "FitToWindow")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_FITTOWINDOW_PROPERTY_EMBEDDEDOBJ.html */
func (e NotesEmbeddedObject) SetFitToWindow(v Boolean) error {
	return putComProperty(e, "FitToWindow", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_NAME_PROPERTY_OBJECT.html */
func (e NotesEmbeddedObject) Name() (String, error) {
	return getComProperty[String](e, "Name")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_OBJECT_PROPERTY.html */
func (e NotesEmbeddedObject) Object() (Variant, error) {
	return getComProperty[Variant](e, "Object")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_PARENT_PROPERTY_OBJECT.html */
func (e NotesEmbeddedObject) Parent() (NotesRichTextItem, error) {
	return getComObjectProperty(e, NewNotesRichTextItem, "Parent")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_RUNREADONLY_PROPERTY_EMBEDDEDOBJ.html */
func (e NotesEmbeddedObject) RunReadOnly() (Boolean, error) {
	return getComProperty[Boolean](e, "RunReadOnly")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_RUNREADONLY_PROPERTY_EMBEDDEDOBJ.html */
func (e NotesEmbeddedObject) SetRunReadOnly(v Boolean) error {
	return putComProperty(e, "RunReadOnly", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SOURCE_PROPERTY.html */
func (e NotesEmbeddedObject) Source() (String, error) {
	return getComProperty[String](e, "Source")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_TYPE_PROPERTY_OBJECT.html */
func (e NotesEmbeddedObject) Type() (NotesEmbeddedObjectEmbedType, error) {
	return getComProperty[NotesEmbeddedObjectEmbedType](e, "Type")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_VERBS_PROPERTY.html */
func (e NotesEmbeddedObject) Verbs() ([]String, error) {
	return getComArrayProperty[String](e, "Verbs")
}

/* --------------------------------- Methods ------------------------------------ */
/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ACTIVATE_METHOD.html */
func (e NotesEmbeddedObject) Activate(show Boolean) (*ole.IDispatch, error) {
	return e.com().CallObjectMethod("Activate", show)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_DOVERB_METHOD.html */
func (e NotesEmbeddedObject) DoVerb(verb String) (Variant, error) {
	return callComMethod[Variant](e, "DoVerb", verb)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_EXTRACTFILE_METHOD.html */
func (e NotesEmbeddedObject) ExtractFile(path String) error {
	return callComVoidMethod(e, "ExtractFile", path)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_REMOVE_METHOD_OBJECT.html */
func (e NotesEmbeddedObject) Remove() error {
	return callComVoidMethod(e, "Remove")
}
