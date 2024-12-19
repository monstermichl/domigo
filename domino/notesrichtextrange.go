/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_NOTESRICHTEXTRANGE_CLASS.html */
package domigo

import (
	"github.com/monstermichl/domigo/helpers"

	ole "github.com/go-ole/go-ole"
)

type NotesRichTextRange struct {
	NotesStruct
}

func NewNotesRichTextRange(dispatchPtr *ole.IDispatch) NotesRichTextRange {
	return NotesRichTextRange{NewNotesStruct(dispatchPtr)}
}

/* --------------------------------- Properties --------------------------------- */
/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_NAVIGATOR_PROPERTY_RTRANGE.html */
func (r NotesRichTextRange) Navigator() (NotesRichTextNavigator, error) {
	dispatchPtr, err := r.com().GetObjectProperty("Navigator")
	return NewNotesRichTextNavigator(dispatchPtr), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_STYLE_PROPERTY_RTRANGE.html */
func (r NotesRichTextRange) Style() (NotesRichTextStyle, error) {
	dispatchPtr, err := r.com().GetObjectProperty("Style")
	return NewNotesRichTextStyle(dispatchPtr), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_TEXTPARAGRAPH_PROPERTY_RTRANGE.html */
func (r NotesRichTextRange) TextParagraph() (String, error) {
	val, err := r.com().GetProperty("TextParagraph")
	return helpers.CastValue[String](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_TEXTRUN_PROPERTY_RTRANGE.html */
func (r NotesRichTextRange) TextRun() (String, error) {
	val, err := r.com().GetProperty("TextRun")
	return helpers.CastValue[String](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_TYPE_PROPERTY_RTRANGE.html */
func (r NotesRichTextRange) Type() (Long, error) {
	val, err := r.com().GetProperty("Type")
	return helpers.CastValue[Long](val), err
}

/* --------------------------------- Methods ------------------------------------ */
/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_CLONE_METHOD_RTRANGE.html */
func (r NotesRichTextRange) Clone() (NotesRichTextRange, error) {
	dispatchPtr, err := r.com().CallObjectMethod("Clone")
	return NewNotesRichTextRange(dispatchPtr), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_FINDANDREPLACE_METHOD_RTRANGE.html */
type notesRichTextRangeFindAndReplaceParams struct {
	options *Long
}

type notesRichTextRangeFindAndReplaceParam func(*notesRichTextRangeFindAndReplaceParams)

func WithNotesRichTextRangeFindAndReplaceOptions(options Long) notesRichTextRangeFindAndReplaceParam {
	return func(c *notesRichTextRangeFindAndReplaceParams) {
		c.options = &options
	}
}

func (r NotesRichTextRange) FindAndReplace(target String, replacement String, params ...notesRichTextRangeFindAndReplaceParam) (Long, error) {
	paramsStruct := &notesRichTextRangeFindAndReplaceParams{}
	paramsOrdered := []interface{}{target, replacement}

	for _, p := range params {
		p(paramsStruct)
	}

	if paramsStruct.options != nil {
		paramsOrdered = append(paramsOrdered, *paramsStruct.options)
	}
	val, err := r.com().CallMethod("FindAndReplace", paramsOrdered...)
	return helpers.CastValue[Long](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_REMOVE_METHOD_RTRANGE.html */
func (r NotesRichTextRange) Remove() error {
	_, err := r.com().CallMethod("Remove")
	return err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_RESET_METHOD_RTRANGE.html */
type notesRichTextRangeResetParams struct {
	begin *Boolean
	end   *Boolean
}

type notesRichTextRangeResetParam func(*notesRichTextRangeResetParams)

func WithNotesRichTextRangeResetBegin(begin Boolean) notesRichTextRangeResetParam {
	return func(c *notesRichTextRangeResetParams) {
		c.begin = &begin
	}
}

func WithNotesRichTextRangeResetEnd(end Boolean) notesRichTextRangeResetParam {
	return func(c *notesRichTextRangeResetParams) {
		c.end = &end
	}
}

func (r NotesRichTextRange) Reset(params ...notesRichTextRangeResetParam) error {
	paramsStruct := &notesRichTextRangeResetParams{}
	paramsOrdered := []interface{}{}

	for _, p := range params {
		p(paramsStruct)
	}

	if paramsStruct.begin != nil {
		paramsOrdered = append(paramsOrdered, *paramsStruct.begin)
		if paramsStruct.end != nil {
			paramsOrdered = append(paramsOrdered, *paramsStruct.end)
		}
	}
	_, err := r.com().CallMethod("Reset", paramsOrdered...)
	return err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SETBEGIN_METHOD_RTRANGE.html */
func (r NotesRichTextRange) SetBegin(element notesStruct) error {
	_, err := r.com().CallMethod("SetBegin", element.com().Dispatch())
	return err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SETEND_METHOD_RTRANGE.html */
func (r NotesRichTextRange) SetEnd(element notesStruct) error {
	_, err := r.com().CallMethod("SetEnd", element.com().Dispatch())
	return err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SETSTYLE_METHOD_RTRANGE.html */
func (r NotesRichTextRange) SetStyle(style NotesRichTextStyle) error {
	_, err := r.com().CallMethod("SetStyle", style.com().Dispatch())
	return err
}
