/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_NOTESRICHTEXTRANGE_CLASS.html */
package notesrichtextrange

import (
	"domigo/domino"
	"domigo/domino/notesrichtextnavigator"
	"domigo/domino/notesrichtextstyle"
	"domigo/helpers"

	ole "github.com/go-ole/go-ole"
)

type NotesRichTextRange struct {
	domino.NotesStruct
}

func New(dispatchPtr *ole.IDispatch) NotesRichTextRange {
	return NotesRichTextRange{domino.NewNotesStruct(dispatchPtr)}
}

/* --------------------------------- Properties --------------------------------- */
/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_NAVIGATOR_PROPERTY_RTRANGE.html */
func (r NotesRichTextRange) Navigator() (notesrichtextnavigator.NotesRichTextNavigator, error) {
	dispatchPtr, err := r.Com().GetObjectProperty("Navigator")
	return notesrichtextnavigator.New(dispatchPtr), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_STYLE_PROPERTY_RTRANGE.html */
func (r NotesRichTextRange) Style() (notesrichtextstyle.NotesRichTextStyle, error) {
	dispatchPtr, err := r.Com().GetObjectProperty("Style")
	return notesrichtextstyle.New(dispatchPtr), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_TEXTPARAGRAPH_PROPERTY_RTRANGE.html */
func (r NotesRichTextRange) TextParagraph() (domino.String, error) {
	val, err := r.Com().GetProperty("TextParagraph")
	return helpers.CastValue[domino.String](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_TEXTRUN_PROPERTY_RTRANGE.html */
func (r NotesRichTextRange) TextRun() (domino.String, error) {
	val, err := r.Com().GetProperty("TextRun")
	return helpers.CastValue[domino.String](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_TYPE_PROPERTY_RTRANGE.html */
func (r NotesRichTextRange) Type() (domino.Long, error) {
	val, err := r.Com().GetProperty("Type")
	return helpers.CastValue[domino.Long](val), err
}

/* --------------------------------- Methods ------------------------------------ */
/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_CLONE_METHOD_RTRANGE.html */
func (r NotesRichTextRange) Clone() (NotesRichTextRange, error) {
	dispatchPtr, err := r.Com().CallObjectMethod("Clone")
	return New(dispatchPtr), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_FINDANDREPLACE_METHOD_RTRANGE.html */
type findAndReplaceParams struct {
	options *domino.Long
}

type findAndReplaceParam func(*findAndReplaceParams)

func WithFindAndReplaceOptions(options domino.Long) findAndReplaceParam {
	return func(c *findAndReplaceParams) {
		c.options = &options
	}
}

func (r NotesRichTextRange) FindAndReplace(target domino.String, replacement domino.String, params ...findAndReplaceParam) (domino.Long, error) {
	paramsStruct := &findAndReplaceParams{}
	paramsOrdered := []interface{}{target, replacement}

	for _, p := range params {
		p(paramsStruct)
	}

	if paramsStruct.options != nil {
		paramsOrdered = append(paramsOrdered, *paramsStruct.options)
	}
	val, err := r.Com().CallMethod("FindAndReplace", paramsOrdered...)
	return helpers.CastValue[domino.Long](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_REMOVE_METHOD_RTRANGE.html */
func (r NotesRichTextRange) Remove() error {
	_, err := r.Com().CallMethod("Remove")
	return err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_RESET_METHOD_RTRANGE.html */
type resetParams struct {
	begin *domino.Boolean
	end   *domino.Boolean
}

type resetParam func(*resetParams)

func WithResetBegin(begin domino.Boolean) resetParam {
	return func(c *resetParams) {
		c.begin = &begin
	}
}

func WithResetEnd(end domino.Boolean) resetParam {
	return func(c *resetParams) {
		c.end = &end
	}
}

func (r NotesRichTextRange) Reset(params ...resetParam) error {
	paramsStruct := &resetParams{}
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
	_, err := r.Com().CallMethod("Reset", paramsOrdered...)
	return err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SETBEGIN_METHOD_RTRANGE.html */
func (r NotesRichTextRange) SetBegin(element domino.NotesConnector) error {
	_, err := r.Com().CallMethod("SetBegin", element.Com().Dispatch())
	return err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SETEND_METHOD_RTRANGE.html */
func (r NotesRichTextRange) SetEnd(element domino.NotesConnector) error {
	_, err := r.Com().CallMethod("SetEnd", element.Com().Dispatch())
	return err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SETSTYLE_METHOD_RTRANGE.html */
func (r NotesRichTextRange) SetStyle(style notesrichtextstyle.NotesRichTextStyle) error {
	_, err := r.Com().CallMethod("SetStyle", style.Com().Dispatch())
	return err
}
