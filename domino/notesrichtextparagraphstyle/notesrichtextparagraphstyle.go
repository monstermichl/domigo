/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_NOTESRICHTEXTPARAGRAPHSTYLE_CLASS_3756.html */
package notesrichtextparagraphstyle

import (
	"github.com/monstermichl/domigo/domino"
	"github.com/monstermichl/domigo/domino/com"
	"github.com/monstermichl/domigo/domino/notesrichtexttab"
	"github.com/monstermichl/domigo/helpers"

	ole "github.com/go-ole/go-ole"
)

type RulerPosition = domino.Long

const (
	RULER_ONE_INCH       RulerPosition = 1440
	RULER_ONE_CENTIMETER RulerPosition = 567
)

type NotesRichTextParagraphStyle struct {
	domino.NotesStruct
}

func New(dispatchPtr *ole.IDispatch) NotesRichTextParagraphStyle {
	return NotesRichTextParagraphStyle{domino.NewNotesStruct(dispatchPtr)}
}

/* --------------------------------- Properties --------------------------------- */
/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ALIGNMENT_PROPERTY_6587.html */
func (r NotesRichTextParagraphStyle) Alignment() (domino.Long, error) {
	val, err := r.Com().GetProperty("Alignment")
	return helpers.CastValue[domino.Long](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ALIGNMENT_PROPERTY_6587.html */
func (r NotesRichTextParagraphStyle) SetAlignment(v domino.Long) error {
	return r.Com().PutProperty("Alignment", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_FIRSTLINELEFTMARGIN_PROPERTY_4412.html */
func (r NotesRichTextParagraphStyle) FirstLineLeftMargin() (domino.Long, error) {
	val, err := r.Com().GetProperty("FirstLineLeftMargin")
	return helpers.CastValue[domino.Long](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_FIRSTLINELEFTMARGIN_PROPERTY_4412.html */
func (r NotesRichTextParagraphStyle) SetFirstLineLeftMargin(v domino.Long) error {
	return r.Com().PutProperty("FirstLineLeftMargin", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_INTERLINESPACING_PROPERTY_9542.html */
func (r NotesRichTextParagraphStyle) InterLineSpacing() (domino.Long, error) {
	val, err := r.Com().GetProperty("InterLineSpacing")
	return helpers.CastValue[domino.Long](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_INTERLINESPACING_PROPERTY_9542.html */
func (r NotesRichTextParagraphStyle) SetInterLineSpacing(v domino.Long) error {
	return r.Com().PutProperty("InterLineSpacing", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_LEFTMARGIN_PROPERTY_4793.html */
func (r NotesRichTextParagraphStyle) LeftMargin() (domino.Long, error) {
	val, err := r.Com().GetProperty("LeftMargin")
	return helpers.CastValue[domino.Long](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_LEFTMARGIN_PROPERTY_4793.html */
func (r NotesRichTextParagraphStyle) SetLeftMargin(v domino.Long) error {
	return r.Com().PutProperty("LeftMargin", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_PAGINATION_PROPERTY_RTPARAGRAPHSTYLE.html */
func (r NotesRichTextParagraphStyle) Pagination() (domino.Long, error) {
	val, err := r.Com().GetProperty("Pagination")
	return helpers.CastValue[domino.Long](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_PAGINATION_PROPERTY_RTPARAGRAPHSTYLE.html */
func (r NotesRichTextParagraphStyle) SetPagination(v domino.Long) error {
	return r.Com().PutProperty("Pagination", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_RIGHTMARGIN_PROPERTY_480.html */
func (r NotesRichTextParagraphStyle) RightMargin() (domino.Long, error) {
	val, err := r.Com().GetProperty("RightMargin")
	return helpers.CastValue[domino.Long](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_RIGHTMARGIN_PROPERTY_480.html */
func (r NotesRichTextParagraphStyle) SetRightMargin(v domino.Long) error {
	return r.Com().PutProperty("RightMargin", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SPACINGABOVE_PROPERTY_7663.html */
func (r NotesRichTextParagraphStyle) SpacingAbove() (domino.Long, error) {
	val, err := r.Com().GetProperty("SpacingAbove")
	return helpers.CastValue[domino.Long](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SPACINGABOVE_PROPERTY_7663.html */
func (r NotesRichTextParagraphStyle) SetSpacingAbove(v domino.Long) error {
	return r.Com().PutProperty("SpacingAbove", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SPACINGBELOW_PROPERTY_8858.html */
func (r NotesRichTextParagraphStyle) SpacingBelow() (domino.Long, error) {
	val, err := r.Com().GetProperty("SpacingBelow")
	return helpers.CastValue[domino.Long](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SPACINGBELOW_PROPERTY_8858.html */
func (r NotesRichTextParagraphStyle) SetSpacingBelow(v domino.Long) error {
	return r.Com().PutProperty("SpacingBelow", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_TABS_PROPERTY_5218.html */
func (r NotesRichTextParagraphStyle) Tabs() ([]notesrichtexttab.NotesRichTextTab, error) {
	return com.GetObjectArrayProperty(r.Com(), notesrichtexttab.New, "Tabs")
}

/* --------------------------------- Methods ------------------------------------ */
/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_CLEARALLTABS_METHOD_3625.html */
func (r NotesRichTextParagraphStyle) ClearAllTabs() error {
	_, err := r.Com().CallMethod("ClearAllTabs")
	return err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SETTAB_METHOD_2382.html */
type setTabParams struct {
	tabType *notesrichtexttab.TabType
}

type setTabParam func(*setTabParams)

func WithSetTabType(tabType notesrichtexttab.TabType) setTabParam {
	return func(c *setTabParams) {
		c.tabType = &tabType
	}
}

func (r NotesRichTextParagraphStyle) SetTab(position domino.Long, params ...setTabParam) error {
	paramsStruct := &setTabParams{}
	paramsOrdered := []interface{}{position}

	for _, p := range params {
		p(paramsStruct)
	}

	if paramsStruct.tabType != nil {
		paramsOrdered = append(paramsOrdered, *paramsStruct.tabType)
	}
	_, err := r.Com().CallMethod("SetTab", paramsOrdered...)
	return err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SETTABS_METHOD_181.html */
type setTabsParams struct {
	interval *domino.Integer
	tabType  *notesrichtexttab.TabType
}

type setTabsParam func(*setTabsParams)

func WithSetTabsInterval(interval domino.Integer) setTabsParam {
	return func(c *setTabsParams) {
		c.interval = &interval
	}
}

func WithSetTabsType(tabType notesrichtexttab.TabType) setTabsParam {
	return func(c *setTabsParams) {
		c.tabType = &tabType
	}
}

func (r NotesRichTextParagraphStyle) SetTabs(number domino.Integer, startposition domino.Long, params ...setTabsParam) error {
	paramsStruct := &setTabsParams{}
	paramsOrdered := []interface{}{number, startposition}

	for _, p := range params {
		p(paramsStruct)
	}

	if paramsStruct.interval != nil {
		paramsOrdered = append(paramsOrdered, *paramsStruct.interval)
		if paramsStruct.tabType != nil {
			paramsOrdered = append(paramsOrdered, *paramsStruct.tabType)
		}
	}
	_, err := r.Com().CallMethod("SetTabs", paramsOrdered...)
	return err
}
