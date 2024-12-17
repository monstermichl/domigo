/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_NOTESOUTLINEENTRY_CLASS.html */
package notesoutlineentry

import (
	"domigo/domino"
	"domigo/domino/notesdocument"
	"domigo/domino/notesview"
	"domigo/helpers"

	ole "github.com/go-ole/go-ole"
)

type NotesOutlineEntry struct {
	domino.NotesStruct
}

func New(dispatchPtr *ole.IDispatch) NotesOutlineEntry {
	return NotesOutlineEntry{domino.NewNotesStruct(dispatchPtr)}
}

/* --------------------------------- Properties --------------------------------- */
/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ALIAS_PROPERTY_OUTLINEENTRY.html */
func (o NotesOutlineEntry) Alias() (domino.String, error) {
	val, err := o.Com().GetProperty("Alias")
	return helpers.CastValue[domino.String](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ALIAS_PROPERTY_OUTLINEENTRY.html */
func (o NotesOutlineEntry) SetAlias(v domino.String) error {
	return o.Com().PutProperty("Alias", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_DOCUMENT_PROPERTY_OUTLINEENTRY.html */
func (o NotesOutlineEntry) Document() (notesdocument.NotesDocument, error) {
	dispatchPtr, err := o.Com().GetObjectProperty("Document")
	return notesdocument.New(dispatchPtr), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ENTRYCLASS_PROPERTY_OUTLINEENTRY.html */
func (o NotesOutlineEntry) EntryClass() (domino.Long, error) {
	val, err := o.Com().GetProperty("EntryClass")
	return helpers.CastValue[domino.Long](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_FORMULA_PROPERTY_OUTLINEENTRY.html */
func (o NotesOutlineEntry) Formula() (domino.String, error) {
	val, err := o.Com().GetProperty("Formula")
	return helpers.CastValue[domino.String](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_FRAMETEXT_PROPERTY_OUTLINEENTRY.html */
func (o NotesOutlineEntry) FrameText() (domino.String, error) {
	val, err := o.Com().GetProperty("FrameText")
	return helpers.CastValue[domino.String](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_FRAMETEXT_PROPERTY_OUTLINEENTRY.html */
func (o NotesOutlineEntry) SetFrameText(v domino.String) error {
	return o.Com().PutProperty("FrameText", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_HASCHILDREN_PROPERTY_OUTLINEENTRY.html */
func (o NotesOutlineEntry) HasChildren() (domino.Boolean, error) {
	val, err := o.Com().GetProperty("HasChildren")
	return helpers.CastValue[domino.Boolean](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_HIDEFORMULA_PROPERTY_OUTLINEENTRY.html */
func (o NotesOutlineEntry) HideFormula() (domino.String, error) {
	val, err := o.Com().GetProperty("HideFormula")
	return helpers.CastValue[domino.String](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_HIDEFORMULA_PROPERTY_OUTLINEENTRY.html */
func (o NotesOutlineEntry) SetHideFormula(v domino.String) error {
	return o.Com().PutProperty("HideFormula", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_IMAGESTEXT_PROPERTY_OUTLINEENTRY.html */
func (o NotesOutlineEntry) ImagesText() (domino.String, error) {
	val, err := o.Com().GetProperty("ImagesText")
	return helpers.CastValue[domino.String](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_IMAGESTEXT_PROPERTY_OUTLINEENTRY.html */
func (o NotesOutlineEntry) SetImagesText(v domino.String) error {
	return o.Com().PutProperty("ImagesText", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISHIDDEN_PROPERTY_OUTLINEENTRY.html */
func (o NotesOutlineEntry) IsHidden() (domino.Boolean, error) {
	val, err := o.Com().GetProperty("IsHidden")
	return helpers.CastValue[domino.Boolean](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISHIDDEN_PROPERTY_OUTLINEENTRY.html */
func (o NotesOutlineEntry) SetIsHidden(v domino.Boolean) error {
	return o.Com().PutProperty("IsHidden", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISHIDDENFROMNOTES_PROPERTY_OUTLINEENTRY.html */
func (o NotesOutlineEntry) IsHiddenFromNotes() (domino.Boolean, error) {
	val, err := o.Com().GetProperty("IsHiddenFromNotes")
	return helpers.CastValue[domino.Boolean](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISHIDDENFROMNOTES_PROPERTY_OUTLINEENTRY.html */
func (o NotesOutlineEntry) SetIsHiddenFromNotes(v domino.Boolean) error {
	return o.Com().PutProperty("IsHiddenFromNotes", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISHIDDENFROMWEB_PROPERTY_OUTLINEENTRY.html */
func (o NotesOutlineEntry) IsHiddenFromWeb() (domino.Boolean, error) {
	val, err := o.Com().GetProperty("IsHiddenFromWeb")
	return helpers.CastValue[domino.Boolean](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISHIDDENFROMWEB_PROPERTY_OUTLINEENTRY.html */
func (o NotesOutlineEntry) SetIsHiddenFromWeb(v domino.Boolean) error {
	return o.Com().PutProperty("IsHiddenFromWeb", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISINTHISDB_PROPERTY_OUTLINEENTRY.html */
func (o NotesOutlineEntry) IsInThisDB() (domino.Boolean, error) {
	val, err := o.Com().GetProperty("IsInThisDB")
	return helpers.CastValue[domino.Boolean](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISPRIVATE_PROPERTY_OUTLINEENTRY.html */
func (o NotesOutlineEntry) IsPrivate() (domino.Boolean, error) {
	val, err := o.Com().GetProperty("IsPrivate")
	return helpers.CastValue[domino.Boolean](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_KEEPSELECTIONFOCUS_PROPERTY_OUTLINEENTRY.html */
func (o NotesOutlineEntry) KeepSelectionFocus() (domino.Boolean, error) {
	val, err := o.Com().GetProperty("KeepSelectionFocus")
	return helpers.CastValue[domino.Boolean](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_KEEPSELECTIONFOCUS_PROPERTY_OUTLINEENTRY.html */
func (o NotesOutlineEntry) SetKeepSelectionFocus(v domino.Boolean) error {
	return o.Com().PutProperty("KeepSelectionFocus", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_LABEL_PROPERTY_OUTLINEENTRY.html */
func (o NotesOutlineEntry) Label() (domino.String, error) {
	val, err := o.Com().GetProperty("Label")
	return helpers.CastValue[domino.String](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_LABEL_PROPERTY_OUTLINEENTRY.html */
func (o NotesOutlineEntry) SetLabel(v domino.String) error {
	return o.Com().PutProperty("Label", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_LEVEL_PROPERTY_OUTLINEENTRY.html */
func (o NotesOutlineEntry) Level() (domino.Long, error) {
	val, err := o.Com().GetProperty("Level")
	return helpers.CastValue[domino.Long](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_NAMEDELEMENT_PROPERTY_OUTLINEENTRY.html */
func (o NotesOutlineEntry) NamedElement() (domino.String, error) {
	val, err := o.Com().GetProperty("NamedElement")
	return helpers.CastValue[domino.String](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_TYPE_PROPERTY_OUTLINEENTRY.html */
func (o NotesOutlineEntry) Type() (domino.Long, error) {
	val, err := o.Com().GetProperty("Type")
	return helpers.CastValue[domino.Long](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_URL_PROPERTY_OUTLINEENTRY.html */
func (o NotesOutlineEntry) URL() (domino.String, error) {
	val, err := o.Com().GetProperty("URL")
	return helpers.CastValue[domino.String](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_USEHIDEFORMULA_PROPERTY_OUTLINEENTRY.html */
func (o NotesOutlineEntry) UseHideFormula() (domino.Boolean, error) {
	val, err := o.Com().GetProperty("UseHideFormula")
	return helpers.CastValue[domino.Boolean](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_USEHIDEFORMULA_PROPERTY_OUTLINEENTRY.html */
func (o NotesOutlineEntry) SetUseHideFormula(v domino.Boolean) error {
	return o.Com().PutProperty("UseHideFormula", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_VIEW_PROPERTY_OUTLINEENTRY.html */
func (o NotesOutlineEntry) View() (notesview.NotesView, error) {
	dispatchPtr, err := o.Com().GetObjectProperty("View")
	return notesview.New(dispatchPtr), err
}

/* --------------------------------- Methods ------------------------------------ */
/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SETACTION_METHOD_OUTLINEENTRY.html */
func (o NotesOutlineEntry) SetAction(formula domino.String) (domino.Boolean, error) {
	val, err := o.Com().CallMethod("SetAction", formula)
	return helpers.CastValue[domino.Boolean](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SETURL_METHOD_OUTLINEENTRY.html */
func (o NotesOutlineEntry) SetURL(url domino.String) (domino.Boolean, error) {
	val, err := o.Com().CallMethod("SetURL", url)
	return helpers.CastValue[domino.Boolean](val), err
}
