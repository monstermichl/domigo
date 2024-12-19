/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_NOTESRICHTEXTSTYLE_CLASS.html */
package domigo

import (
	"github.com/monstermichl/domigo/internal/helpers"

	ole "github.com/go-ole/go-ole"
)

type NotesRichTextStyle struct {
	NotesStruct
}

func NewNotesRichTextStyle(dispatchPtr *ole.IDispatch) NotesRichTextStyle {
	return NotesRichTextStyle{NewNotesStruct(dispatchPtr)}
}

/* --------------------------------- Properties --------------------------------- */
/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_BOLD_PROPERTY.html */
func (r NotesRichTextStyle) Bold() (Long, error) {
	val, err := r.com().GetProperty("Bold")
	return helpers.CastValue[Long](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_BOLD_PROPERTY.html */
func (r NotesRichTextStyle) SetBold(v Long) error {
	return r.com().PutProperty("Bold", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_EFFECTS_PROPERTY.html */
func (r NotesRichTextStyle) Effects() (Long, error) {
	val, err := r.com().GetProperty("Effects")
	return helpers.CastValue[Long](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_EFFECTS_PROPERTY.html */
func (r NotesRichTextStyle) SetEffects(v Long) error {
	return r.com().PutProperty("Effects", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_FONTSIZE_PROPERTY.html */
func (r NotesRichTextStyle) FontSize() (Long, error) {
	val, err := r.com().GetProperty("FontSize")
	return helpers.CastValue[Long](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_FONTSIZE_PROPERTY.html */
func (r NotesRichTextStyle) SetFontSize(v Long) error {
	return r.com().PutProperty("FontSize", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISDEFAULT_PROPERTY_RTSTYLE.html */
func (r NotesRichTextStyle) IsDefault() (Boolean, error) {
	val, err := r.com().GetProperty("IsDefault")
	return helpers.CastValue[Boolean](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ITALIC_PROPERTY.html */
func (r NotesRichTextStyle) Italic() (Long, error) {
	val, err := r.com().GetProperty("Italic")
	return helpers.CastValue[Long](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ITALIC_PROPERTY.html */
func (r NotesRichTextStyle) SetItalic(v Long) error {
	return r.com().PutProperty("Italic", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_NOTESCOLOR_PROPERTY.html */
func (r NotesRichTextStyle) NotesColor() (Long, error) {
	val, err := r.com().GetProperty("NotesColor")
	return helpers.CastValue[Long](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_NOTESCOLOR_PROPERTY.html */
func (r NotesRichTextStyle) SetNotesColor(v Long) error {
	return r.com().PutProperty("NotesColor", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_NOTESFONT_PROPERTY.html */
func (r NotesRichTextStyle) NotesFont() (Long, error) {
	val, err := r.com().GetProperty("NotesFont")
	return helpers.CastValue[Long](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_NOTESFONT_PROPERTY.html */
func (r NotesRichTextStyle) SetNotesFont(v Long) error {
	return r.com().PutProperty("NotesFont", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_PARENT_PROPERTY_RICHTEXTSTYLE_COM.html */
func (r NotesRichTextStyle) Parent() (NotesSession, error) {
	dispatchPtr, err := r.com().GetObjectProperty("Parent")
	return NewNotesSession(dispatchPtr), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_PASSTHRUHTML_PROPERTY_514_ABOUT.html */
func (r NotesRichTextStyle) PassThruHTML() (Long, error) {
	val, err := r.com().GetProperty("PassThruHTML")
	return helpers.CastValue[Long](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_PASSTHRUHTML_PROPERTY_514_ABOUT.html */
func (r NotesRichTextStyle) SetPassThruHTML(v Long) error {
	return r.com().PutProperty("PassThruHTML", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_STRIKETHROUGH_PROPERTY.html */
func (r NotesRichTextStyle) StrikeThrough() (Long, error) {
	val, err := r.com().GetProperty("StrikeThrough")
	return helpers.CastValue[Long](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_STRIKETHROUGH_PROPERTY.html */
func (r NotesRichTextStyle) SetStrikeThrough(v Long) error {
	return r.com().PutProperty("StrikeThrough", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_UNDERLINE_PROPERTY.html */
func (r NotesRichTextStyle) Underline() (Long, error) {
	val, err := r.com().GetProperty("Underline")
	return helpers.CastValue[Long](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_UNDERLINE_PROPERTY.html */
func (r NotesRichTextStyle) SetUnderline(v Long) error {
	return r.com().PutProperty("Underline", v)
}

/* --------------------------------- Methods ------------------------------------ */
