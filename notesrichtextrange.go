/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_NOTESRICHTEXTRANGE_CLASS.html */
package domigo

import (
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
	return getComObjectProperty(r, NewNotesRichTextNavigator, "Navigator")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_STYLE_PROPERTY_RTRANGE.html */
func (r NotesRichTextRange) Style() (NotesRichTextStyle, error) {
	return getComObjectProperty(r, NewNotesRichTextStyle, "Style")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_TEXTPARAGRAPH_PROPERTY_RTRANGE.html */
func (r NotesRichTextRange) TextParagraph() (String, error) {
	return getComProperty[String](r, "TextParagraph")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_TEXTRUN_PROPERTY_RTRANGE.html */
func (r NotesRichTextRange) TextRun() (String, error) {
	return getComProperty[String](r, "TextRun")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_TYPE_PROPERTY_RTRANGE.html */
func (r NotesRichTextRange) Type() (Long, error) {
	return getComProperty[Long](r, "Type")
}

/* --------------------------------- Methods ------------------------------------ */
/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_CLONE_METHOD_RTRANGE.html */
func (r NotesRichTextRange) Clone() (NotesRichTextRange, error) {
	return callComObjectMethod(r, NewNotesRichTextRange, "Clone")
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
	return callComMethod[Long](r, "FindAndReplace", paramsOrdered...)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_REMOVE_METHOD_RTRANGE.html */
func (r NotesRichTextRange) Remove() error {
	return callComVoidMethod(r, "Remove")
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
	return callComVoidMethod(r, "Reset", paramsOrdered...)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SETBEGIN_METHOD_RTRANGE.html */
func (r NotesRichTextRange) SetBegin(element notesStruct) error {
	return callComVoidMethod(r, "SetBegin", element.com().Dispatch())
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SETEND_METHOD_RTRANGE.html */
func (r NotesRichTextRange) SetEnd(element notesStruct) error {
	return callComVoidMethod(r, "SetEnd", element.com().Dispatch())
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SETSTYLE_METHOD_RTRANGE.html */
func (r NotesRichTextRange) SetStyle(style NotesRichTextStyle) error {
	return callComVoidMethod(r, "SetStyle", style.com().Dispatch())
}
