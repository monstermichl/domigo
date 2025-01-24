/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_NOTESOUTLINEENTRY_CLASS.html */
package domigo

import (
	ole "github.com/go-ole/go-ole"
)

type NotesOutlineEntry struct {
	NotesStruct
}

func newNotesOutlineEntry(dispatchPtr *ole.IDispatch) NotesOutlineEntry {
	return NotesOutlineEntry{newNotesStruct(dispatchPtr)}
}

/* --------------------------------- Properties --------------------------------- */
/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ALIAS_PROPERTY_OUTLINEENTRY.html */
func (o NotesOutlineEntry) Alias() (String, error) {
	return getComProperty[String](o, "Alias")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ALIAS_PROPERTY_OUTLINEENTRY.html */
func (o NotesOutlineEntry) SetAlias(v String) error {
	return putComProperty(o, "Alias", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_DATABASE_PROPERTY_OUTLINEENTRY.html */
func (o NotesOutlineEntry) Database() (NotesDatabase, error) {
	return getComObjectProperty(o, newNotesDatabase, "Database")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_DOCUMENT_PROPERTY_OUTLINEENTRY.html */
func (o NotesOutlineEntry) Document() (NotesDocument, error) {
	return getComObjectProperty(o, newNotesDocument, "Document")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ENTRYCLASS_PROPERTY_OUTLINEENTRY.html */
func (o NotesOutlineEntry) EntryClass() (Long, error) {
	return getComProperty[Long](o, "EntryClass")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_FORMULA_PROPERTY_OUTLINEENTRY.html */
func (o NotesOutlineEntry) Formula() (String, error) {
	return getComProperty[String](o, "Formula")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_FRAMETEXT_PROPERTY_OUTLINEENTRY.html */
func (o NotesOutlineEntry) FrameText() (String, error) {
	return getComProperty[String](o, "FrameText")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_FRAMETEXT_PROPERTY_OUTLINEENTRY.html */
func (o NotesOutlineEntry) SetFrameText(v String) error {
	return putComProperty(o, "FrameText", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_HASCHILDREN_PROPERTY_OUTLINEENTRY.html */
func (o NotesOutlineEntry) HasChildren() (Boolean, error) {
	return getComProperty[Boolean](o, "HasChildren")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_HIDEFORMULA_PROPERTY_OUTLINEENTRY.html */
func (o NotesOutlineEntry) HideFormula() (String, error) {
	return getComProperty[String](o, "HideFormula")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_HIDEFORMULA_PROPERTY_OUTLINEENTRY.html */
func (o NotesOutlineEntry) SetHideFormula(v String) error {
	return putComProperty(o, "HideFormula", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_IMAGESTEXT_PROPERTY_OUTLINEENTRY.html */
func (o NotesOutlineEntry) ImagesText() (String, error) {
	return getComProperty[String](o, "ImagesText")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_IMAGESTEXT_PROPERTY_OUTLINEENTRY.html */
func (o NotesOutlineEntry) SetImagesText(v String) error {
	return putComProperty(o, "ImagesText", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISHIDDEN_PROPERTY_OUTLINEENTRY.html */
func (o NotesOutlineEntry) IsHidden() (Boolean, error) {
	return getComProperty[Boolean](o, "IsHidden")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISHIDDEN_PROPERTY_OUTLINEENTRY.html */
func (o NotesOutlineEntry) SetIsHidden(v Boolean) error {
	return putComProperty(o, "IsHidden", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISHIDDENFROMNOTES_PROPERTY_OUTLINEENTRY.html */
func (o NotesOutlineEntry) IsHiddenFromNotes() (Boolean, error) {
	return getComProperty[Boolean](o, "IsHiddenFromNotes")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISHIDDENFROMNOTES_PROPERTY_OUTLINEENTRY.html */
func (o NotesOutlineEntry) SetIsHiddenFromNotes(v Boolean) error {
	return putComProperty(o, "IsHiddenFromNotes", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISHIDDENFROMWEB_PROPERTY_OUTLINEENTRY.html */
func (o NotesOutlineEntry) IsHiddenFromWeb() (Boolean, error) {
	return getComProperty[Boolean](o, "IsHiddenFromWeb")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISHIDDENFROMWEB_PROPERTY_OUTLINEENTRY.html */
func (o NotesOutlineEntry) SetIsHiddenFromWeb(v Boolean) error {
	return putComProperty(o, "IsHiddenFromWeb", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISINTHISDB_PROPERTY_OUTLINEENTRY.html */
func (o NotesOutlineEntry) IsInThisDB() (Boolean, error) {
	return getComProperty[Boolean](o, "IsInThisDB")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISPRIVATE_PROPERTY_OUTLINEENTRY.html */
func (o NotesOutlineEntry) IsPrivate() (Boolean, error) {
	return getComProperty[Boolean](o, "IsPrivate")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_KEEPSELECTIONFOCUS_PROPERTY_OUTLINEENTRY.html */
func (o NotesOutlineEntry) KeepSelectionFocus() (Boolean, error) {
	return getComProperty[Boolean](o, "KeepSelectionFocus")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_KEEPSELECTIONFOCUS_PROPERTY_OUTLINEENTRY.html */
func (o NotesOutlineEntry) SetKeepSelectionFocus(v Boolean) error {
	return putComProperty(o, "KeepSelectionFocus", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_LABEL_PROPERTY_OUTLINEENTRY.html */
func (o NotesOutlineEntry) Label() (String, error) {
	return getComProperty[String](o, "Label")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_LABEL_PROPERTY_OUTLINEENTRY.html */
func (o NotesOutlineEntry) SetLabel(v String) error {
	return putComProperty(o, "Label", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_LEVEL_PROPERTY_OUTLINEENTRY.html */
func (o NotesOutlineEntry) Level() (Long, error) {
	return getComProperty[Long](o, "Level")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_NAMEDELEMENT_PROPERTY_OUTLINEENTRY.html */
func (o NotesOutlineEntry) NamedElement() (String, error) {
	return getComProperty[String](o, "NamedElement")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_PARENT_PROPERTY_OUTLINEENTRY_COM.html */
func (o NotesOutlineEntry) Parent() (NotesOutline, error) {
	return getComObjectProperty(o, newNotesOutline, "Parent")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_TYPE_PROPERTY_OUTLINEENTRY.html */
func (o NotesOutlineEntry) Type() (Long, error) {
	return getComProperty[Long](o, "Type")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_URL_PROPERTY_OUTLINEENTRY.html */
func (o NotesOutlineEntry) URL() (String, error) {
	return getComProperty[String](o, "URL")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_USEHIDEFORMULA_PROPERTY_OUTLINEENTRY.html */
func (o NotesOutlineEntry) UseHideFormula() (Boolean, error) {
	return getComProperty[Boolean](o, "UseHideFormula")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_USEHIDEFORMULA_PROPERTY_OUTLINEENTRY.html */
func (o NotesOutlineEntry) SetUseHideFormula(v Boolean) error {
	return putComProperty(o, "UseHideFormula", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_VIEW_PROPERTY_OUTLINEENTRY.html */
func (o NotesOutlineEntry) View() (NotesView, error) {
	return getComObjectProperty(o, newNotesView, "View")
}

/* --------------------------------- Methods ------------------------------------ */
/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SETACTION_METHOD_OUTLINEENTRY.html */
func (o NotesOutlineEntry) SetAction(formula String) (Boolean, error) {
	return callComMethod[Boolean](o, "SetAction", formula)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SETNAMEDELEMENT_METHOD_OUTLINEENTRY.html */
func (o NotesOutlineEntry) SetNamedElement(notesDatabase NotesDatabase, elementName String, entryclass Long) (Boolean, error) {
	return callComMethod[Boolean](o, "SetNamedElement", notesDatabase, elementName, entryclass)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SETNOTELINK_METHOD_OUTLINEENTRY.html */
func (o NotesOutlineEntry) SetNoteLink(notesDatabase NotesDatabase, notesView NotesView, notesDocument NotesDocument, Obj notesStruct) (Boolean, error) {
	return callComMethod[Boolean](o, "SetNoteLink", notesDatabase, notesView, notesDocument, Obj)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SETURL_METHOD_OUTLINEENTRY.html */
func (o NotesOutlineEntry) SetURL(url String) (Boolean, error) {
	return callComMethod[Boolean](o, "SetURL", url)
}
