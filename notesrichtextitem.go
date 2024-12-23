/* https://help.hcl-software.com/dom_designer/10.0.1/basic/H_NOTESRICHTEXTITEM_CLASS.html */
package domigo

import (
	"github.com/monstermichl/domigo/internal/com"

	ole "github.com/go-ole/go-ole"
)

type NotesRichTextItem struct {
	NotesStruct
}

func NewNotesRichTextItem(dispatchPtr *ole.IDispatch) NotesRichTextItem {
	return NotesRichTextItem{NewNotesStruct(dispatchPtr)}
}

/* --------------------------------- Properties --------------------------------- */
/* https://help.hcl-software.com/dom_designer/10.0.1/basic/H_EMBEDDEDOBJECTS_PROPERTY_RTITEM.html */
func (r NotesRichTextItem) EmbeddedObjects() ([]NotesEmbeddedObject, error) {
	return com.GetObjectArrayProperty(r.com(), NewNotesEmbeddedObject, "EmbeddedObjects")
}

/* --------------------------------- Methods ------------------------------------ */
/* https://help.hcl-software.com/dom_designer/10.0.1/basic/H_ADDNEWLINE_METHOD.html */
type notesRichTextItemAddNewLineParams struct {
	n              *Integer
	forceParagraph *Boolean
}

type notesRichTextItemAddNewLineParam func(*notesRichTextItemAddNewLineParams)

func WithNotesRichTextItemAddNewLineN(n Integer) notesRichTextItemAddNewLineParam {
	return func(c *notesRichTextItemAddNewLineParams) {
		c.n = &n
	}
}

func WithNotesRichTextItemAddNewLineForceParagraph(forceParagraph Boolean) notesRichTextItemAddNewLineParam {
	return func(c *notesRichTextItemAddNewLineParams) {
		c.forceParagraph = &forceParagraph
	}
}

func (r NotesRichTextItem) AddNewLine(params ...notesRichTextItemAddNewLineParam) error {
	paramsStruct := &notesRichTextItemAddNewLineParams{}
	paramsOrdered := []interface{}{}

	for _, p := range params {
		p(paramsStruct)
	}

	if paramsStruct.n != nil {
		paramsOrdered = append(paramsOrdered, *paramsStruct.n)
		if paramsStruct.forceParagraph != nil {
			paramsOrdered = append(paramsOrdered, *paramsStruct.forceParagraph)
		}
	}
	return callComVoidMethod(r, "AddNewLine", paramsOrdered...)
}

/* https://help.hcl-software.com/dom_designer/10.0.1/basic/H_ADDPAGEBREAK_METHOD_290_ABOUT.html */
type notesRichTextItemAddPageBreakParams struct {
	notesRichTextParagraphStyle *NotesRichTextParagraphStyle
}

type notesRichTextItemAddPageBreakParam func(*notesRichTextItemAddPageBreakParams)

func WithNotesRichTextItemAddPageBreakNotesRichTextParagraphStyle(notesRichTextParagraphStyle NotesRichTextParagraphStyle) notesRichTextItemAddPageBreakParam {
	return func(c *notesRichTextItemAddPageBreakParams) {
		c.notesRichTextParagraphStyle = &notesRichTextParagraphStyle
	}
}

func (r NotesRichTextItem) AddPageBreak(params ...notesRichTextItemAddPageBreakParam) error {
	paramsStruct := &notesRichTextItemAddPageBreakParams{}
	paramsOrdered := []interface{}{}

	for _, p := range params {
		p(paramsStruct)
	}

	if paramsStruct.notesRichTextParagraphStyle != nil {
		paramsOrdered = append(paramsOrdered, *paramsStruct.notesRichTextParagraphStyle)
	}
	return callComVoidMethod(r, "AddPageBreak", paramsOrdered...)
}

/* https://help.hcl-software.com/dom_designer/10.0.1/basic/H_ADDTAB_METHOD.html */
type notesRichTextItemAddTabParams struct {
	n *Integer
}

type notesRichTextItemAddTabParam func(*notesRichTextItemAddTabParams)

func WithNotesRichTextItemAddTabN(n Integer) notesRichTextItemAddTabParam {
	return func(c *notesRichTextItemAddTabParams) {
		c.n = &n
	}
}

func (r NotesRichTextItem) AddTab(params ...notesRichTextItemAddTabParam) error {
	paramsStruct := &notesRichTextItemAddTabParams{}
	paramsOrdered := []interface{}{}

	for _, p := range params {
		p(paramsStruct)
	}

	if paramsStruct.n != nil {
		paramsOrdered = append(paramsOrdered, *paramsStruct.n)
	}
	return callComVoidMethod(r, "AddTab", paramsOrdered...)
}

/* https://help.hcl-software.com/dom_designer/10.0.1/basic/H_APPENDDOCLINK_METHOD.html */
type notesRichTextItemAppendDocLinkParams struct {
	HotSpotText *String
}

type notesRichTextItemAppendDocLinkParam func(*notesRichTextItemAppendDocLinkParams)

func WithNotesRichTextItemAppendDocLinkHotSpotText(HotSpotText String) notesRichTextItemAppendDocLinkParam {
	return func(c *notesRichTextItemAppendDocLinkParams) {
		c.HotSpotText = &HotSpotText
	}
}

func (r NotesRichTextItem) AppendDocLink(linkTo notesStruct, comment String, params ...notesRichTextItemAppendDocLinkParam) error {
	paramsStruct := &notesRichTextItemAppendDocLinkParams{}
	paramsOrdered := []interface{}{linkTo, comment}

	for _, p := range params {
		p(paramsStruct)
	}

	if paramsStruct.HotSpotText != nil {
		paramsOrdered = append(paramsOrdered, *paramsStruct.HotSpotText)
	}
	return callComVoidMethod(r, "AppendDocLink", paramsOrdered...)
}

/* https://help.hcl-software.com/dom_designer/10.0.1/basic/H_APPENDPARAGRAPHSTYLE_METHOD_1636_ABOUT.html */
func (r NotesRichTextItem) AppendParagraphStyle(notesRichTextParagraphStyle NotesRichTextParagraphStyle) error {
	return callComVoidMethod(r, "AppendParagraphStyle", notesRichTextParagraphStyle)
}

/* https://help.hcl-software.com/dom_designer/10.0.1/basic/H_APPENDRTITEM_METHOD.html */
func (r NotesRichTextItem) AppendRTItem(notesRichTextItem NotesRichTextItem) error {
	return callComVoidMethod(r, "AppendRTItem", notesRichTextItem)
}

/* https://help.hcl-software.com/dom_designer/10.0.1/basic/H_APPENDSTYLE_METHOD.html */
func (r NotesRichTextItem) AppendStyle(notesRichTextStyle NotesRichTextStyle) error {
	return callComVoidMethod(r, "AppendStyle", notesRichTextStyle)
}

/* https://help.hcl-software.com/dom_designer/10.0.1/basic/H_APPENDTABLE_METHOD_RTITEM.html */
type notesRichTextItemAppendTableParams struct {
	labels         *[]String
	leftMargin     *Long
	rtpsStyleArray *[]NotesRichTextParagraphStyle
}

type notesRichTextItemAppendTableParam func(*notesRichTextItemAppendTableParams)

func WithNotesRichTextItemAppendTableLabels(labels []String) notesRichTextItemAppendTableParam {
	return func(c *notesRichTextItemAppendTableParams) {
		c.labels = &labels
	}
}

func WithNotesRichTextItemAppendTableLeftMargin(leftMargin Long) notesRichTextItemAppendTableParam {
	return func(c *notesRichTextItemAppendTableParams) {
		c.leftMargin = &leftMargin
	}
}

func WithNotesRichTextItemAppendTableRtpsStyleArray(rtpsStyleArray []NotesRichTextParagraphStyle) notesRichTextItemAppendTableParam {
	return func(c *notesRichTextItemAppendTableParams) {
		c.rtpsStyleArray = &rtpsStyleArray
	}
}

func (r NotesRichTextItem) AppendTable(rows Integer, columns Integer, params ...notesRichTextItemAppendTableParam) error {
	paramsStruct := &notesRichTextItemAppendTableParams{}
	paramsOrdered := []interface{}{rows, columns}

	for _, p := range params {
		p(paramsStruct)
	}

	if paramsStruct.labels != nil {
		paramsOrdered = append(paramsOrdered, *paramsStruct.labels)
		if paramsStruct.leftMargin != nil {
			paramsOrdered = append(paramsOrdered, *paramsStruct.leftMargin)
			if paramsStruct.rtpsStyleArray != nil {
				paramsOrdered = append(paramsOrdered, dispatchSlice(*paramsStruct.rtpsStyleArray))
			}
		}
	}
	return callComVoidMethod(r, "AppendTable", paramsOrdered...)
}

/* https://help.hcl-software.com/dom_designer/10.0.1/basic/H_APPENDTEXT_METHOD.html */
func (r NotesRichTextItem) AppendText(text String) error {
	return callComVoidMethod(r, "AppendText", text)
}

/* https://help.hcl-software.com/dom_designer/10.0.1/basic/H_BEGININSERT_METHOD_RTITEM.html */
type notesRichTextItemBeginInsertParams struct {
	after *Boolean
}

type notesRichTextItemBeginInsertParam func(*notesRichTextItemBeginInsertParams)

func WithNotesRichTextItemBeginInsertAfter(after Boolean) notesRichTextItemBeginInsertParam {
	return func(c *notesRichTextItemBeginInsertParams) {
		c.after = &after
	}
}

func (r NotesRichTextItem) BeginInsert(element notesStruct, params ...notesRichTextItemBeginInsertParam) error {
	paramsStruct := &notesRichTextItemBeginInsertParams{}
	paramsOrdered := []interface{}{element}

	for _, p := range params {
		p(paramsStruct)
	}

	if paramsStruct.after != nil {
		paramsOrdered = append(paramsOrdered, *paramsStruct.after)
	}
	return callComVoidMethod(r, "BeginInsert", paramsOrdered...)
}

/* https://help.hcl-software.com/dom_designer/10.0.1/basic/H_BEGINSECTION_METHOD_RTITEM.html */
type notesRichTextItemBeginSectionParams struct {
	titleStyle *NotesRichTextStyle
	barColor   *NotesColorObject
	expand     *Boolean
}

type notesRichTextItemBeginSectionParam func(*notesRichTextItemBeginSectionParams)

func WithNotesRichTextItemBeginSectionTitleStyle(titleStyle NotesRichTextStyle) notesRichTextItemBeginSectionParam {
	return func(c *notesRichTextItemBeginSectionParams) {
		c.titleStyle = &titleStyle
	}
}

func WithNotesRichTextItemBeginSectionBarColor(barColor NotesColorObject) notesRichTextItemBeginSectionParam {
	return func(c *notesRichTextItemBeginSectionParams) {
		c.barColor = &barColor
	}
}

func WithNotesRichTextItemBeginSectionExpand(expand Boolean) notesRichTextItemBeginSectionParam {
	return func(c *notesRichTextItemBeginSectionParams) {
		c.expand = &expand
	}
}

func (r NotesRichTextItem) BeginSection(title String, params ...notesRichTextItemBeginSectionParam) error {
	paramsStruct := &notesRichTextItemBeginSectionParams{}
	paramsOrdered := []interface{}{title}

	for _, p := range params {
		p(paramsStruct)
	}

	if paramsStruct.titleStyle != nil {
		paramsOrdered = append(paramsOrdered, *paramsStruct.titleStyle)
		if paramsStruct.barColor != nil {
			paramsOrdered = append(paramsOrdered, *paramsStruct.barColor)
			if paramsStruct.expand != nil {
				paramsOrdered = append(paramsOrdered, *paramsStruct.expand)
			}
		}
	}
	return callComVoidMethod(r, "BeginSection", paramsOrdered...)
}

/* https://help.hcl-software.com/dom_designer/10.0.1/basic/H_COMPACT_METHOD_RTITEM.html */
func (r NotesRichTextItem) Compact() error {
	return callComVoidMethod(r, "Compact")
}

/* https://help.hcl-software.com/dom_designer/10.0.1/basic/H_CONVERTOHTML_METHOD_NOTESRICHTEXTITEM.html */
func (r NotesRichTextItem) Converttohtml() error {
	return callComVoidMethod(r, "Converttohtml")
}

/* https://help.hcl-software.com/dom_designer/10.0.1/basic/H_CREATENAVIGATOR_METHOD_RTITEM.html */
func (r NotesRichTextItem) CreateNavigator() (NotesRichTextNavigator, error) {
	return callComObjectMethod(r, NewNotesRichTextNavigator, "CreateNavigator")
}

/* https://help.hcl-software.com/dom_designer/10.0.1/basic/H_CREATERANGE_METHOD_RTITEM.html */
func (r NotesRichTextItem) CreateRange() (NotesRichTextRange, error) {
	return callComObjectMethod(r, NewNotesRichTextRange, "CreateRange")
}

/* https://help.hcl-software.com/dom_designer/10.0.1/basic/H_EMBEDOBJECT_METHOD.html */
type notesRichTextItemEmbedObjectParams struct {
	name *String
}

type notesRichTextItemEmbedObjectParam func(*notesRichTextItemEmbedObjectParams)

func WithNotesRichTextItemEmbedObjectName(name String) notesRichTextItemEmbedObjectParam {
	return func(c *notesRichTextItemEmbedObjectParams) {
		c.name = &name
	}
}

func (r NotesRichTextItem) EmbedObject(embedType NotesEmbeddedObjectEmbedType, class String, source String, params ...notesRichTextItemEmbedObjectParam) (NotesEmbeddedObject, error) {
	paramsStruct := &notesRichTextItemEmbedObjectParams{}
	paramsOrdered := []interface{}{embedType, class, source}

	for _, p := range params {
		p(paramsStruct)
	}

	if paramsStruct.name != nil {
		paramsOrdered = append(paramsOrdered, *paramsStruct.name)
	}
	return callComObjectMethod(r, NewNotesEmbeddedObject, "EmbedObject", paramsOrdered...)
}

/* https://help.hcl-software.com/dom_designer/10.0.1/basic/H_ENDINSERT_METHOD_RTITEM.html */
func (r NotesRichTextItem) EndInsert() error {
	return callComVoidMethod(r, "EndInsert")
}

/* https://help.hcl-software.com/dom_designer/10.0.1/basic/H_ENDSECTION_METHOD_RTITEM.html */
func (r NotesRichTextItem) EndSection() error {
	return callComVoidMethod(r, "EndSection")
}

/* https://help.hcl-software.com/dom_designer/10.0.1/basic/H_GETEMBEDDEDOBJECT_METHOD.html */
func (r NotesRichTextItem) GetEmbeddedObject(name String) (NotesEmbeddedObject, error) {
	return callComObjectMethod(r, NewNotesEmbeddedObject, "GetEmbeddedObject", name)
}

/* https://help.hcl-software.com/dom_designer/10.0.1/basic/H_GETFORMATTEDTEXT_METHOD.html */
func (r NotesRichTextItem) GetFormattedText(tabstrip Boolean, lineLength Integer) (String, error) {
	return callComMethod[String](r, "GetFormattedText", tabstrip, lineLength)
}

/* https://help.hcl-software.com/dom_designer/10.0.1/basic/H_GETHTMLREFERENCES_METHOD_NOTESRICHTEXTITEM.html */
func (r NotesRichTextItem) GetHTMLReferences() (Variant, error) {
	return callComMethod[Variant](r, "GetHTMLReferences")
}

/* https://help.hcl-software.com/dom_designer/10.0.1/basic/H_GETNOTESFONT_METHOD_RTITEM.html */
type notesRichTextItemGetNotesFontParams struct {
	addOnFail *Boolean
}

type notesRichTextItemGetNotesFontParam func(*notesRichTextItemGetNotesFontParams)

func WithGetNotesFontAddOnFail(addOnFail Boolean) notesRichTextItemGetNotesFontParam {
	return func(c *notesRichTextItemGetNotesFontParams) {
		c.addOnFail = &addOnFail
	}
}

func (r NotesRichTextItem) GetNotesFont(faceName String, params ...notesRichTextItemGetNotesFontParam) (Long, error) {
	paramsStruct := &notesRichTextItemGetNotesFontParams{}
	paramsOrdered := []interface{}{faceName}

	for _, p := range params {
		p(paramsStruct)
	}

	if paramsStruct.addOnFail != nil {
		paramsOrdered = append(paramsOrdered, *paramsStruct.addOnFail)
	}
	return callComMethod[Long](r, "GetNotesFont", paramsOrdered...)
}

/* https://help.hcl-software.com/dom_designer/10.0.1/basic/H_GETUNFORMATTEDTEXT_METHOD.html */
func (r NotesRichTextItem) GetUnformattedText() (String, error) {
	return callComMethod[String](r, "GetUnformattedText")
}

/* https://help.hcl-software.com/dom_designer/10.0.1/basic/H_UPDATE_METHOD_RTITEM.html */
func (r NotesRichTextItem) Update() error {
	return callComVoidMethod(r, "Update")
}
