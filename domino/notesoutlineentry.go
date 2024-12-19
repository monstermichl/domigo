/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_NOTESOUTLINEENTRY_CLASS.html */
package domigo

import (
	"github.com/monstermichl/domigo/helpers"

	ole "github.com/go-ole/go-ole"
)

type NotesOutlineEntry struct {
	NotesStruct
}

func NewNotesOutlineEntry(dispatchPtr *ole.IDispatch) NotesOutlineEntry {
	return NotesOutlineEntry{NewNotesStruct(dispatchPtr)}
}

/* --------------------------------- Properties --------------------------------- */
/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ALIAS_PROPERTY_OUTLINEENTRY.html */
func (o NotesOutlineEntry) Alias() (String, error) {
	val, err := o.Com().GetProperty("Alias")
	return helpers.CastValue[String](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ALIAS_PROPERTY_OUTLINEENTRY.html */
func (o NotesOutlineEntry) SetAlias(v String) error {
	return o.Com().PutProperty("Alias", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_DATABASE_PROPERTY_OUTLINEENTRY.html */
func (o NotesOutlineEntry) Database() (NotesDatabase, error) {
	dispatchPtr, err := o.Com().GetObjectProperty("Database")
	return NewNotesDatabase(dispatchPtr), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_DOCUMENT_PROPERTY_OUTLINEENTRY.html */
func (o NotesOutlineEntry) Document() (NotesDocument, error) {
	dispatchPtr, err := o.Com().GetObjectProperty("Document")
	return NewNotesDocument(dispatchPtr), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ENTRYCLASS_PROPERTY_OUTLINEENTRY.html */
func (o NotesOutlineEntry) EntryClass() (Long, error) {
	val, err := o.Com().GetProperty("EntryClass")
	return helpers.CastValue[Long](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_FORMULA_PROPERTY_OUTLINEENTRY.html */
func (o NotesOutlineEntry) Formula() (String, error) {
	val, err := o.Com().GetProperty("Formula")
	return helpers.CastValue[String](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_FRAMETEXT_PROPERTY_OUTLINEENTRY.html */
func (o NotesOutlineEntry) FrameText() (String, error) {
	val, err := o.Com().GetProperty("FrameText")
	return helpers.CastValue[String](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_FRAMETEXT_PROPERTY_OUTLINEENTRY.html */
func (o NotesOutlineEntry) SetFrameText(v String) error {
	return o.Com().PutProperty("FrameText", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_HASCHILDREN_PROPERTY_OUTLINEENTRY.html */
func (o NotesOutlineEntry) HasChildren() (Boolean, error) {
	val, err := o.Com().GetProperty("HasChildren")
	return helpers.CastValue[Boolean](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_HIDEFORMULA_PROPERTY_OUTLINEENTRY.html */
func (o NotesOutlineEntry) HideFormula() (String, error) {
	val, err := o.Com().GetProperty("HideFormula")
	return helpers.CastValue[String](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_HIDEFORMULA_PROPERTY_OUTLINEENTRY.html */
func (o NotesOutlineEntry) SetHideFormula(v String) error {
	return o.Com().PutProperty("HideFormula", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_IMAGESTEXT_PROPERTY_OUTLINEENTRY.html */
func (o NotesOutlineEntry) ImagesText() (String, error) {
	val, err := o.Com().GetProperty("ImagesText")
	return helpers.CastValue[String](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_IMAGESTEXT_PROPERTY_OUTLINEENTRY.html */
func (o NotesOutlineEntry) SetImagesText(v String) error {
	return o.Com().PutProperty("ImagesText", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISHIDDEN_PROPERTY_OUTLINEENTRY.html */
func (o NotesOutlineEntry) IsHidden() (Boolean, error) {
	val, err := o.Com().GetProperty("IsHidden")
	return helpers.CastValue[Boolean](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISHIDDEN_PROPERTY_OUTLINEENTRY.html */
func (o NotesOutlineEntry) SetIsHidden(v Boolean) error {
	return o.Com().PutProperty("IsHidden", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISHIDDENFROMNOTES_PROPERTY_OUTLINEENTRY.html */
func (o NotesOutlineEntry) IsHiddenFromNotes() (Boolean, error) {
	val, err := o.Com().GetProperty("IsHiddenFromNotes")
	return helpers.CastValue[Boolean](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISHIDDENFROMNOTES_PROPERTY_OUTLINEENTRY.html */
func (o NotesOutlineEntry) SetIsHiddenFromNotes(v Boolean) error {
	return o.Com().PutProperty("IsHiddenFromNotes", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISHIDDENFROMWEB_PROPERTY_OUTLINEENTRY.html */
func (o NotesOutlineEntry) IsHiddenFromWeb() (Boolean, error) {
	val, err := o.Com().GetProperty("IsHiddenFromWeb")
	return helpers.CastValue[Boolean](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISHIDDENFROMWEB_PROPERTY_OUTLINEENTRY.html */
func (o NotesOutlineEntry) SetIsHiddenFromWeb(v Boolean) error {
	return o.Com().PutProperty("IsHiddenFromWeb", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISINTHISDB_PROPERTY_OUTLINEENTRY.html */
func (o NotesOutlineEntry) IsInThisDB() (Boolean, error) {
	val, err := o.Com().GetProperty("IsInThisDB")
	return helpers.CastValue[Boolean](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISPRIVATE_PROPERTY_OUTLINEENTRY.html */
func (o NotesOutlineEntry) IsPrivate() (Boolean, error) {
	val, err := o.Com().GetProperty("IsPrivate")
	return helpers.CastValue[Boolean](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_KEEPSELECTIONFOCUS_PROPERTY_OUTLINEENTRY.html */
func (o NotesOutlineEntry) KeepSelectionFocus() (Boolean, error) {
	val, err := o.Com().GetProperty("KeepSelectionFocus")
	return helpers.CastValue[Boolean](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_KEEPSELECTIONFOCUS_PROPERTY_OUTLINEENTRY.html */
func (o NotesOutlineEntry) SetKeepSelectionFocus(v Boolean) error {
	return o.Com().PutProperty("KeepSelectionFocus", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_LABEL_PROPERTY_OUTLINEENTRY.html */
func (o NotesOutlineEntry) Label() (String, error) {
	val, err := o.Com().GetProperty("Label")
	return helpers.CastValue[String](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_LABEL_PROPERTY_OUTLINEENTRY.html */
func (o NotesOutlineEntry) SetLabel(v String) error {
	return o.Com().PutProperty("Label", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_LEVEL_PROPERTY_OUTLINEENTRY.html */
func (o NotesOutlineEntry) Level() (Long, error) {
	val, err := o.Com().GetProperty("Level")
	return helpers.CastValue[Long](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_NAMEDELEMENT_PROPERTY_OUTLINEENTRY.html */
func (o NotesOutlineEntry) NamedElement() (String, error) {
	val, err := o.Com().GetProperty("NamedElement")
	return helpers.CastValue[String](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_PARENT_PROPERTY_OUTLINEENTRY_COM.html */
func (o NotesOutlineEntry) Parent() (NotesOutline, error) {
	dispatchPtr, err := o.Com().GetObjectProperty("Parent")
	return NewNotesOutline(dispatchPtr), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_TYPE_PROPERTY_OUTLINEENTRY.html */
func (o NotesOutlineEntry) Type() (Long, error) {
	val, err := o.Com().GetProperty("Type")
	return helpers.CastValue[Long](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_URL_PROPERTY_OUTLINEENTRY.html */
func (o NotesOutlineEntry) URL() (String, error) {
	val, err := o.Com().GetProperty("URL")
	return helpers.CastValue[String](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_USEHIDEFORMULA_PROPERTY_OUTLINEENTRY.html */
func (o NotesOutlineEntry) UseHideFormula() (Boolean, error) {
	val, err := o.Com().GetProperty("UseHideFormula")
	return helpers.CastValue[Boolean](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_USEHIDEFORMULA_PROPERTY_OUTLINEENTRY.html */
func (o NotesOutlineEntry) SetUseHideFormula(v Boolean) error {
	return o.Com().PutProperty("UseHideFormula", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_VIEW_PROPERTY_OUTLINEENTRY.html */
func (o NotesOutlineEntry) View() (NotesView, error) {
	dispatchPtr, err := o.Com().GetObjectProperty("View")
	return NewNotesView(dispatchPtr), err
}

/* --------------------------------- Methods ------------------------------------ */
/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SETACTION_METHOD_OUTLINEENTRY.html */
func (o NotesOutlineEntry) SetAction(formula String) (Boolean, error) {
	val, err := o.Com().CallMethod("SetAction", formula)
	return helpers.CastValue[Boolean](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SETNAMEDELEMENT_METHOD_OUTLINEENTRY.html */
func (o NotesOutlineEntry) SetNamedElement(notesDatabase NotesDatabase, elementName String, entryclass Long) (Boolean, error) {
	val, err := o.Com().CallMethod("SetNamedElement", notesDatabase.Com().Dispatch(), elementName, entryclass)
	return helpers.CastValue[Boolean](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SETNOTELINK_METHOD_OUTLINEENTRY.html */
func (o NotesOutlineEntry) SetNoteLink(notesDatabase NotesDatabase, notesView NotesView, notesDocument NotesDocument, Obj NotesConnector) (Boolean, error) {
	val, err := o.Com().CallMethod("SetNoteLink", notesDatabase.Com().Dispatch(), notesView.Com().Dispatch(), notesDocument.Com().Dispatch(), Obj.Com().Dispatch())
	return helpers.CastValue[Boolean](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SETURL_METHOD_OUTLINEENTRY.html */
func (o NotesOutlineEntry) SetURL(url String) (Boolean, error) {
	val, err := o.Com().CallMethod("SetURL", url)
	return helpers.CastValue[Boolean](val), err
}
