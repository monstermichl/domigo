/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_NOTESRICHTEXTPARAGRAPHSTYLE_CLASS_3756.html */
package domigo

import (
	"github.com/monstermichl/domigo/internal/com"

	ole "github.com/go-ole/go-ole"
)

type NotesRichTextParagraphStyleRulerPosition = Long

const (
	NOTESRICHTEXTPARAGRAPHSTYLE_RULER_ONE_INCH       NotesRichTextParagraphStyleRulerPosition = 1440
	NOTESRICHTEXTPARAGRAPHSTYLE_RULER_ONE_CENTIMETER NotesRichTextParagraphStyleRulerPosition = 567
)

type NotesRichTextParagraphStyle struct {
	NotesStruct
}

func NewNotesRichTextParagraphStyle(dispatchPtr *ole.IDispatch) NotesRichTextParagraphStyle {
	return NotesRichTextParagraphStyle{NewNotesStruct(dispatchPtr)}
}

/* --------------------------------- Properties --------------------------------- */
/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ALIGNMENT_PROPERTY_6587.html */
func (r NotesRichTextParagraphStyle) Alignment() (Long, error) {
	return getComProperty[Long](r, "Alignment")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ALIGNMENT_PROPERTY_6587.html */
func (r NotesRichTextParagraphStyle) SetAlignment(v Long) error {
	return putComProperty(r, "Alignment", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_FIRSTLINELEFTMARGIN_PROPERTY_4412.html */
func (r NotesRichTextParagraphStyle) FirstLineLeftMargin() (Long, error) {
	return getComProperty[Long](r, "FirstLineLeftMargin")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_FIRSTLINELEFTMARGIN_PROPERTY_4412.html */
func (r NotesRichTextParagraphStyle) SetFirstLineLeftMargin(v Long) error {
	return putComProperty(r, "FirstLineLeftMargin", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_INTERLINESPACING_PROPERTY_9542.html */
func (r NotesRichTextParagraphStyle) InterLineSpacing() (Long, error) {
	return getComProperty[Long](r, "InterLineSpacing")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_INTERLINESPACING_PROPERTY_9542.html */
func (r NotesRichTextParagraphStyle) SetInterLineSpacing(v Long) error {
	return putComProperty(r, "InterLineSpacing", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_LEFTMARGIN_PROPERTY_4793.html */
func (r NotesRichTextParagraphStyle) LeftMargin() (Long, error) {
	return getComProperty[Long](r, "LeftMargin")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_LEFTMARGIN_PROPERTY_4793.html */
func (r NotesRichTextParagraphStyle) SetLeftMargin(v Long) error {
	return putComProperty(r, "LeftMargin", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_PAGINATION_PROPERTY_RTPARAGRAPHSTYLE.html */
func (r NotesRichTextParagraphStyle) Pagination() (Long, error) {
	return getComProperty[Long](r, "Pagination")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_PAGINATION_PROPERTY_RTPARAGRAPHSTYLE.html */
func (r NotesRichTextParagraphStyle) SetPagination(v Long) error {
	return putComProperty(r, "Pagination", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_RIGHTMARGIN_PROPERTY_480.html */
func (r NotesRichTextParagraphStyle) RightMargin() (Long, error) {
	return getComProperty[Long](r, "RightMargin")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_RIGHTMARGIN_PROPERTY_480.html */
func (r NotesRichTextParagraphStyle) SetRightMargin(v Long) error {
	return putComProperty(r, "RightMargin", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SPACINGABOVE_PROPERTY_7663.html */
func (r NotesRichTextParagraphStyle) SpacingAbove() (Long, error) {
	return getComProperty[Long](r, "SpacingAbove")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SPACINGABOVE_PROPERTY_7663.html */
func (r NotesRichTextParagraphStyle) SetSpacingAbove(v Long) error {
	return putComProperty(r, "SpacingAbove", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SPACINGBELOW_PROPERTY_8858.html */
func (r NotesRichTextParagraphStyle) SpacingBelow() (Long, error) {
	return getComProperty[Long](r, "SpacingBelow")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SPACINGBELOW_PROPERTY_8858.html */
func (r NotesRichTextParagraphStyle) SetSpacingBelow(v Long) error {
	return putComProperty(r, "SpacingBelow", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_TABS_PROPERTY_5218.html */
func (r NotesRichTextParagraphStyle) Tabs() ([]NotesRichTextTab, error) {
	return com.GetObjectArrayProperty(r.com(), NewNotesRichTextTab, "Tabs")
}

/* --------------------------------- Methods ------------------------------------ */
/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_CLEARALLTABS_METHOD_3625.html */
func (r NotesRichTextParagraphStyle) ClearAllTabs() error {
	return callComVoidMethod(r, "ClearAllTabs")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SETTAB_METHOD_2382.html */
type notesRichTextParagraphStyleSetTabParams struct {
	tabType *NotesRichTextTabType
}

type notesRichTextParagraphStyleSetTabParam func(*notesRichTextParagraphStyleSetTabParams)

func WithNotesRichTextParagraphStyleSetTabType(tabType NotesRichTextTabType) notesRichTextParagraphStyleSetTabParam {
	return func(c *notesRichTextParagraphStyleSetTabParams) {
		c.tabType = &tabType
	}
}

func (r NotesRichTextParagraphStyle) SetTab(position Long, params ...notesRichTextParagraphStyleSetTabParam) error {
	paramsStruct := &notesRichTextParagraphStyleSetTabParams{}
	paramsOrdered := []interface{}{position}

	for _, p := range params {
		p(paramsStruct)
	}

	if paramsStruct.tabType != nil {
		paramsOrdered = append(paramsOrdered, *paramsStruct.tabType)
	}
	return callComVoidMethod(r, "SetTab", paramsOrdered...)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SETTABS_METHOD_181.html */
type notesRichTextParagraphStyleSetTabsParams struct {
	interval *Integer
	tabType  *NotesRichTextTabType
}

type notesRichTextParagraphStyleSetTabsParam func(*notesRichTextParagraphStyleSetTabsParams)

func WithNotesRichTextParagraphStyleSetTabsInterval(interval Integer) notesRichTextParagraphStyleSetTabsParam {
	return func(c *notesRichTextParagraphStyleSetTabsParams) {
		c.interval = &interval
	}
}

func WithNotesRichTextParagraphStyleSetTabsType(tabType NotesRichTextTabType) notesRichTextParagraphStyleSetTabsParam {
	return func(c *notesRichTextParagraphStyleSetTabsParams) {
		c.tabType = &tabType
	}
}

func (r NotesRichTextParagraphStyle) SetTabs(number Integer, startposition Long, params ...notesRichTextParagraphStyleSetTabsParam) error {
	paramsStruct := &notesRichTextParagraphStyleSetTabsParams{}
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
	return callComVoidMethod(r, "SetTabs", paramsOrdered...)
}
