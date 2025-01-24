/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_NOTESRICHTEXTSTYLE_CLASS.html */
package domigo

import (
	ole "github.com/go-ole/go-ole"
)

type NotesRichTextStyle struct {
	NotesStruct
}

func newNotesRichTextStyle(dispatchPtr *ole.IDispatch) NotesRichTextStyle {
	return NotesRichTextStyle{newNotesStruct(dispatchPtr)}
}

/* --------------------------------- Properties --------------------------------- */
/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_BOLD_PROPERTY.html */
func (r NotesRichTextStyle) Bold() (Long, error) {
	return getComProperty[Long](r, "Bold")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_BOLD_PROPERTY.html */
func (r NotesRichTextStyle) SetBold(v Long) error {
	return putComProperty(r, "Bold", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_EFFECTS_PROPERTY.html */
func (r NotesRichTextStyle) Effects() (Long, error) {
	return getComProperty[Long](r, "Effects")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_EFFECTS_PROPERTY.html */
func (r NotesRichTextStyle) SetEffects(v Long) error {
	return putComProperty(r, "Effects", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_FONTSIZE_PROPERTY.html */
func (r NotesRichTextStyle) FontSize() (Long, error) {
	return getComProperty[Long](r, "FontSize")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_FONTSIZE_PROPERTY.html */
func (r NotesRichTextStyle) SetFontSize(v Long) error {
	return putComProperty(r, "FontSize", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISDEFAULT_PROPERTY_RTSTYLE.html */
func (r NotesRichTextStyle) IsDefault() (Boolean, error) {
	return getComProperty[Boolean](r, "IsDefault")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ITALIC_PROPERTY.html */
func (r NotesRichTextStyle) Italic() (Long, error) {
	return getComProperty[Long](r, "Italic")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ITALIC_PROPERTY.html */
func (r NotesRichTextStyle) SetItalic(v Long) error {
	return putComProperty(r, "Italic", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_NOTESCOLOR_PROPERTY.html */
func (r NotesRichTextStyle) NotesColor() (Long, error) {
	return getComProperty[Long](r, "NotesColor")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_NOTESCOLOR_PROPERTY.html */
func (r NotesRichTextStyle) SetNotesColor(v Long) error {
	return putComProperty(r, "NotesColor", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_NOTESFONT_PROPERTY.html */
func (r NotesRichTextStyle) NotesFont() (Long, error) {
	return getComProperty[Long](r, "NotesFont")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_NOTESFONT_PROPERTY.html */
func (r NotesRichTextStyle) SetNotesFont(v Long) error {
	return putComProperty(r, "NotesFont", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_PARENT_PROPERTY_RICHTEXTSTYLE_COM.html */
func (r NotesRichTextStyle) Parent() (NotesSession, error) {
	return getComObjectProperty(r, newNotesSession, "Parent")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_PASSTHRUHTML_PROPERTY_514_ABOUT.html */
func (r NotesRichTextStyle) PassThruHTML() (Long, error) {
	return getComProperty[Long](r, "PassThruHTML")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_PASSTHRUHTML_PROPERTY_514_ABOUT.html */
func (r NotesRichTextStyle) SetPassThruHTML(v Long) error {
	return putComProperty(r, "PassThruHTML", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_STRIKETHROUGH_PROPERTY.html */
func (r NotesRichTextStyle) StrikeThrough() (Long, error) {
	return getComProperty[Long](r, "StrikeThrough")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_STRIKETHROUGH_PROPERTY.html */
func (r NotesRichTextStyle) SetStrikeThrough(v Long) error {
	return putComProperty(r, "StrikeThrough", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_UNDERLINE_PROPERTY.html */
func (r NotesRichTextStyle) Underline() (Long, error) {
	return getComProperty[Long](r, "Underline")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_UNDERLINE_PROPERTY.html */
func (r NotesRichTextStyle) SetUnderline(v Long) error {
	return putComProperty(r, "Underline", v)
}

/* --------------------------------- Methods ------------------------------------ */
