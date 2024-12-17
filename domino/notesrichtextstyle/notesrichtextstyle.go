/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_NOTESRICHTEXTSTYLE_CLASS.html */
package notesrichtextstyle

import (
	"github.com/monstermichl/domigo/domino"
	"github.com/monstermichl/domigo/helpers"

	ole "github.com/go-ole/go-ole"
)

type NotesRichTextStyle struct {
	domino.NotesStruct
}

func New(dispatchPtr *ole.IDispatch) NotesRichTextStyle {
	return NotesRichTextStyle{domino.NewNotesStruct(dispatchPtr)}
}

/* --------------------------------- Properties --------------------------------- */
/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_BOLD_PROPERTY.html */
func (r NotesRichTextStyle) Bold() (domino.Long, error) {
	val, err := r.Com().GetProperty("Bold")
	return helpers.CastValue[domino.Long](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_BOLD_PROPERTY.html */
func (r NotesRichTextStyle) SetBold(v domino.Long) error {
	return r.Com().PutProperty("Bold", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_EFFECTS_PROPERTY.html */
func (r NotesRichTextStyle) Effects() (domino.Long, error) {
	val, err := r.Com().GetProperty("Effects")
	return helpers.CastValue[domino.Long](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_EFFECTS_PROPERTY.html */
func (r NotesRichTextStyle) SetEffects(v domino.Long) error {
	return r.Com().PutProperty("Effects", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_FONTSIZE_PROPERTY.html */
func (r NotesRichTextStyle) FontSize() (domino.Long, error) {
	val, err := r.Com().GetProperty("FontSize")
	return helpers.CastValue[domino.Long](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_FONTSIZE_PROPERTY.html */
func (r NotesRichTextStyle) SetFontSize(v domino.Long) error {
	return r.Com().PutProperty("FontSize", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISDEFAULT_PROPERTY_RTSTYLE.html */
func (r NotesRichTextStyle) IsDefault() (domino.Boolean, error) {
	val, err := r.Com().GetProperty("IsDefault")
	return helpers.CastValue[domino.Boolean](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ITALIC_PROPERTY.html */
func (r NotesRichTextStyle) Italic() (domino.Long, error) {
	val, err := r.Com().GetProperty("Italic")
	return helpers.CastValue[domino.Long](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ITALIC_PROPERTY.html */
func (r NotesRichTextStyle) SetItalic(v domino.Long) error {
	return r.Com().PutProperty("Italic", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_NOTESCOLOR_PROPERTY.html */
func (r NotesRichTextStyle) NotesColor() (domino.Long, error) {
	val, err := r.Com().GetProperty("NotesColor")
	return helpers.CastValue[domino.Long](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_NOTESCOLOR_PROPERTY.html */
func (r NotesRichTextStyle) SetNotesColor(v domino.Long) error {
	return r.Com().PutProperty("NotesColor", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_NOTESFONT_PROPERTY.html */
func (r NotesRichTextStyle) NotesFont() (domino.Long, error) {
	val, err := r.Com().GetProperty("NotesFont")
	return helpers.CastValue[domino.Long](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_NOTESFONT_PROPERTY.html */
func (r NotesRichTextStyle) SetNotesFont(v domino.Long) error {
	return r.Com().PutProperty("NotesFont", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_PASSTHRUHTML_PROPERTY_514_ABOUT.html */
func (r NotesRichTextStyle) PassThruHTML() (domino.Long, error) {
	val, err := r.Com().GetProperty("PassThruHTML")
	return helpers.CastValue[domino.Long](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_PASSTHRUHTML_PROPERTY_514_ABOUT.html */
func (r NotesRichTextStyle) SetPassThruHTML(v domino.Long) error {
	return r.Com().PutProperty("PassThruHTML", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_STRIKETHROUGH_PROPERTY.html */
func (r NotesRichTextStyle) StrikeThrough() (domino.Long, error) {
	val, err := r.Com().GetProperty("StrikeThrough")
	return helpers.CastValue[domino.Long](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_STRIKETHROUGH_PROPERTY.html */
func (r NotesRichTextStyle) SetStrikeThrough(v domino.Long) error {
	return r.Com().PutProperty("StrikeThrough", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_UNDERLINE_PROPERTY.html */
func (r NotesRichTextStyle) Underline() (domino.Long, error) {
	val, err := r.Com().GetProperty("Underline")
	return helpers.CastValue[domino.Long](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_UNDERLINE_PROPERTY.html */
func (r NotesRichTextStyle) SetUnderline(v domino.Long) error {
	return r.Com().PutProperty("Underline", v)
}

/* --------------------------------- Methods ------------------------------------ */
