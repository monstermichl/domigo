/* https://help.hcl-software.com/dom_designer/10.0.1/basic/H_NOTESRICHTEXTITEM_CLASS.html */
package notesrichtextitem

import (
	"github.com/monstermichl/domigo/domino"
	"github.com/monstermichl/domigo/domino/com"
	"github.com/monstermichl/domigo/domino/notescolorobject"
	"github.com/monstermichl/domigo/domino/notesembeddedobject"
	"github.com/monstermichl/domigo/domino/notesitem"
	"github.com/monstermichl/domigo/domino/notesrichtextnavigator"
	"github.com/monstermichl/domigo/domino/notesrichtextparagraphstyle"
	"github.com/monstermichl/domigo/domino/notesrichtextrange"
	"github.com/monstermichl/domigo/domino/notesrichtextstyle"
	"github.com/monstermichl/domigo/helpers"

	ole "github.com/go-ole/go-ole"
)

type NotesRichTextItem struct {
	notesitem.NotesItem
}

func New(dispatchPtr *ole.IDispatch) NotesRichTextItem {
	return NotesRichTextItem{notesitem.New(dispatchPtr)}
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_PARENT_PROPERTY_OBJECT.html */
/* Moved from NotesEmbeddedObject. */
func NotesEmbeddedObjectParent(e notesembeddedobject.NotesEmbeddedObject) (NotesRichTextItem, error) {
	dispatchPtr, err := e.Com().GetObjectProperty("Parent")
	return New(dispatchPtr), err
}

/* --------------------------------- Properties --------------------------------- */
/* https://help.hcl-software.com/dom_designer/10.0.1/basic/H_EMBEDDEDOBJECTS_PROPERTY_RTITEM.html */
func (r NotesRichTextItem) EmbeddedObjects() ([]notesembeddedobject.NotesEmbeddedObject, error) {
	return com.GetObjectArrayProperty(r.Com(), notesembeddedobject.New, "EmbeddedObjects")
}

/* --------------------------------- Methods ------------------------------------ */
/* https://help.hcl-software.com/dom_designer/10.0.1/basic/H_ADDNEWLINE_METHOD.html */
type addNewLineParams struct {
	n              *domino.Integer
	forceParagraph *domino.Boolean
}

type addNewLineParam func(*addNewLineParams)

func WithAddNewLineN(n domino.Integer) addNewLineParam {
	return func(c *addNewLineParams) {
		c.n = &n
	}
}

func WithAddNewLineForceParagraph(forceParagraph domino.Boolean) addNewLineParam {
	return func(c *addNewLineParams) {
		c.forceParagraph = &forceParagraph
	}
}

func (r NotesRichTextItem) AddNewLine(params ...addNewLineParam) error {
	paramsStruct := &addNewLineParams{}
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
	_, err := r.Com().CallMethod("AddNewLine", paramsOrdered...)
	return err
}

/* https://help.hcl-software.com/dom_designer/10.0.1/basic/H_ADDPAGEBREAK_METHOD_290_ABOUT.html */
type addPageBreakParams struct {
	notesRichTextParagraphStyle *notesrichtextparagraphstyle.NotesRichTextParagraphStyle
}

type addPageBreakParam func(*addPageBreakParams)

func WithAddPageBreakNotesRichTextParagraphStyle(notesRichTextParagraphStyle notesrichtextparagraphstyle.NotesRichTextParagraphStyle) addPageBreakParam {
	return func(c *addPageBreakParams) {
		c.notesRichTextParagraphStyle = &notesRichTextParagraphStyle
	}
}

func (r NotesRichTextItem) AddPageBreak(params ...addPageBreakParam) error {
	paramsStruct := &addPageBreakParams{}
	paramsOrdered := []interface{}{}

	for _, p := range params {
		p(paramsStruct)
	}

	if paramsStruct.notesRichTextParagraphStyle != nil {
		paramsOrdered = append(paramsOrdered, *paramsStruct.notesRichTextParagraphStyle.Com().Dispatch())
	}
	_, err := r.Com().CallMethod("AddPageBreak", paramsOrdered...)
	return err
}

/* https://help.hcl-software.com/dom_designer/10.0.1/basic/H_ADDTAB_METHOD.html */
type addTabParams struct {
	n *domino.Integer
}

type addTabParam func(*addTabParams)

func WithAddTabN(n domino.Integer) addTabParam {
	return func(c *addTabParams) {
		c.n = &n
	}
}

func (r NotesRichTextItem) AddTab(params ...addTabParam) error {
	paramsStruct := &addTabParams{}
	paramsOrdered := []interface{}{}

	for _, p := range params {
		p(paramsStruct)
	}

	if paramsStruct.n != nil {
		paramsOrdered = append(paramsOrdered, *paramsStruct.n)
	}
	_, err := r.Com().CallMethod("AddTab", paramsOrdered...)
	return err
}

/* https://help.hcl-software.com/dom_designer/10.0.1/basic/H_APPENDPARAGRAPHSTYLE_METHOD_1636_ABOUT.html */
func (r NotesRichTextItem) AppendParagraphStyle(notesRichTextParagraphStyle notesrichtextparagraphstyle.NotesRichTextParagraphStyle) error {
	_, err := r.Com().CallMethod("AppendParagraphStyle", notesRichTextParagraphStyle.Com().Dispatch())
	return err
}

/* https://help.hcl-software.com/dom_designer/10.0.1/basic/H_APPENDRTITEM_METHOD.html */
func (r NotesRichTextItem) AppendRTItem(notesRichTextItem NotesRichTextItem) error {
	_, err := r.Com().CallMethod("AppendRTItem", notesRichTextItem.Com().Dispatch())
	return err
}

/* https://help.hcl-software.com/dom_designer/10.0.1/basic/H_APPENDSTYLE_METHOD.html */
func (r NotesRichTextItem) AppendStyle(notesRichTextStyle notesrichtextstyle.NotesRichTextStyle) error {
	_, err := r.Com().CallMethod("AppendStyle", notesRichTextStyle.Com().Dispatch())
	return err
}

/* https://help.hcl-software.com/dom_designer/10.0.1/basic/H_APPENDTABLE_METHOD_RTITEM.html */
type appendTableParams struct {
	labels         *[]domino.String
	leftMargin     *domino.Long
	rtpsStyleArray *[]notesrichtextparagraphstyle.NotesRichTextParagraphStyle
}

type appendTableParam func(*appendTableParams)

func WithAppendTableLabels(labels []domino.String) appendTableParam {
	return func(c *appendTableParams) {
		c.labels = &labels
	}
}

func WithAppendTableLeftMargin(leftMargin domino.Long) appendTableParam {
	return func(c *appendTableParams) {
		c.leftMargin = &leftMargin
	}
}

func WithAppendTableRtpsStyleArray(rtpsStyleArray []notesrichtextparagraphstyle.NotesRichTextParagraphStyle) appendTableParam {
	return func(c *appendTableParams) {
		c.rtpsStyleArray = &rtpsStyleArray
	}
}

func (r NotesRichTextItem) AppendTable(rows domino.Integer, columns domino.Integer, params ...appendTableParam) error {
	paramsStruct := &appendTableParams{}
	paramsOrdered := []interface{}{rows, columns}

	for _, p := range params {
		p(paramsStruct)
	}

	if paramsStruct.labels != nil {
		paramsOrdered = append(paramsOrdered, *paramsStruct.labels)
		if paramsStruct.leftMargin != nil {
			paramsOrdered = append(paramsOrdered, *paramsStruct.leftMargin)
			if paramsStruct.rtpsStyleArray != nil {
				paramsOrdered = append(paramsOrdered, domino.DispatchSlice(*paramsStruct.rtpsStyleArray))
			}
		}
	}
	_, err := r.Com().CallMethod("AppendTable", paramsOrdered...)
	return err
}

/* https://help.hcl-software.com/dom_designer/10.0.1/basic/H_APPENDTEXT_METHOD.html */
func (r NotesRichTextItem) AppendText(text domino.String) error {
	_, err := r.Com().CallMethod("AppendText", text)
	return err
}

/* https://help.hcl-software.com/dom_designer/10.0.1/basic/H_BEGININSERT_METHOD_RTITEM.html */
type beginInsertParams struct {
	after *domino.Boolean
}

type beginInsertParam func(*beginInsertParams)

func WithBeginInsertAfter(after domino.Boolean) beginInsertParam {
	return func(c *beginInsertParams) {
		c.after = &after
	}
}

func (r NotesRichTextItem) BeginInsert(element domino.NotesConnector, params ...beginInsertParam) error {
	paramsStruct := &beginInsertParams{}
	paramsOrdered := []interface{}{element.Com().Dispatch()}

	for _, p := range params {
		p(paramsStruct)
	}

	if paramsStruct.after != nil {
		paramsOrdered = append(paramsOrdered, *paramsStruct.after)
	}
	_, err := r.Com().CallMethod("BeginInsert", paramsOrdered...)
	return err
}

/* https://help.hcl-software.com/dom_designer/10.0.1/basic/H_BEGINSECTION_METHOD_RTITEM.html */
type beginSectionParams struct {
	titleStyle *notesrichtextstyle.NotesRichTextStyle
	barColor   *notescolorobject.NotesColorObject
	expand     *domino.Boolean
}

type beginSectionParam func(*beginSectionParams)

func WithBeginSectionTitleStyle(titleStyle notesrichtextstyle.NotesRichTextStyle) beginSectionParam {
	return func(c *beginSectionParams) {
		c.titleStyle = &titleStyle
	}
}

func WithBeginSectionBarColor(barColor notescolorobject.NotesColorObject) beginSectionParam {
	return func(c *beginSectionParams) {
		c.barColor = &barColor
	}
}

func WithBeginSectionExpand(expand domino.Boolean) beginSectionParam {
	return func(c *beginSectionParams) {
		c.expand = &expand
	}
}

func (r NotesRichTextItem) BeginSection(title domino.String, params ...beginSectionParam) error {
	paramsStruct := &beginSectionParams{}
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
	_, err := r.Com().CallMethod("BeginSection", paramsOrdered...)
	return err
}

/* https://help.hcl-software.com/dom_designer/10.0.1/basic/H_COMPACT_METHOD_RTITEM.html */
func (r NotesRichTextItem) Compact() error {
	_, err := r.Com().CallMethod("Compact")
	return err
}

/* https://help.hcl-software.com/dom_designer/10.0.1/basic/H_CONVERTOHTML_METHOD_NOTESRICHTEXTITEM.html */
func (r NotesRichTextItem) Converttohtml() error {
	_, err := r.Com().CallMethod("Converttohtml")
	return err
}

/* https://help.hcl-software.com/dom_designer/10.0.1/basic/H_CREATENAVIGATOR_METHOD_RTITEM.html */
func (r NotesRichTextItem) CreateNavigator() (notesrichtextnavigator.NotesRichTextNavigator, error) {
	dispatchPtr, err := r.Com().CallObjectMethod("CreateNavigator")
	return notesrichtextnavigator.New(dispatchPtr), err
}

/* https://help.hcl-software.com/dom_designer/10.0.1/basic/H_CREATERANGE_METHOD_RTITEM.html */
func (r NotesRichTextItem) CreateRange() (notesrichtextrange.NotesRichTextRange, error) {
	dispatchPtr, err := r.Com().CallObjectMethod("CreateRange")
	return notesrichtextrange.New(dispatchPtr), err
}

/* https://help.hcl-software.com/dom_designer/10.0.1/basic/H_EMBEDOBJECT_METHOD.html */
type embedObjectParams struct {
	name *domino.String
}

type embedObjectParam func(*embedObjectParams)

func WithEmbedObjectName(name domino.String) embedObjectParam {
	return func(c *embedObjectParams) {
		c.name = &name
	}
}

func (r NotesRichTextItem) EmbedObject(t notesembeddedobject.EmbedType, class domino.String, source domino.String, params ...embedObjectParam) (notesembeddedobject.NotesEmbeddedObject, error) {
	paramsStruct := &embedObjectParams{}
	paramsOrdered := []interface{}{t, class, source}

	for _, p := range params {
		p(paramsStruct)
	}

	if paramsStruct.name != nil {
		paramsOrdered = append(paramsOrdered, *paramsStruct.name)
	}
	dispatchPtr, err := r.Com().CallObjectMethod("EmbedObject", paramsOrdered...)
	return notesembeddedobject.New(dispatchPtr), err
}

/* https://help.hcl-software.com/dom_designer/10.0.1/basic/H_ENDINSERT_METHOD_RTITEM.html */
func (r NotesRichTextItem) EndInsert() error {
	_, err := r.Com().CallMethod("EndInsert")
	return err
}

/* https://help.hcl-software.com/dom_designer/10.0.1/basic/H_ENDSECTION_METHOD_RTITEM.html */
func (r NotesRichTextItem) EndSection() error {
	_, err := r.Com().CallMethod("EndSection")
	return err
}

/* https://help.hcl-software.com/dom_designer/10.0.1/basic/H_GETEMBEDDEDOBJECT_METHOD.html */
func (r NotesRichTextItem) GetEmbeddedObject(name domino.String) (notesembeddedobject.NotesEmbeddedObject, error) {
	dispatchPtr, err := r.Com().CallObjectMethod("GetEmbeddedObject", name)
	return notesembeddedobject.New(dispatchPtr), err
}

/* https://help.hcl-software.com/dom_designer/10.0.1/basic/H_GETFORMATTEDTEXT_METHOD.html */
func (r NotesRichTextItem) GetFormattedText(tabstrip domino.Boolean, lineLength domino.Integer) (domino.String, error) {
	val, err := r.Com().CallMethod("GetFormattedText", tabstrip, lineLength)
	return helpers.CastValue[domino.String](val), err
}

/* https://help.hcl-software.com/dom_designer/10.0.1/basic/H_GETHTMLREFERENCES_METHOD_NOTESRICHTEXTITEM.html */
func (r NotesRichTextItem) GetHTMLReferences() (domino.Variant, error) {
	val, err := r.Com().CallMethod("GetHTMLReferences")
	return helpers.CastValue[domino.Variant](val), err
}

/* https://help.hcl-software.com/dom_designer/10.0.1/basic/H_GETNOTESFONT_METHOD_RTITEM.html */
type getNotesFontParams struct {
	addOnFail *domino.Boolean
}

type getNotesFontParam func(*getNotesFontParams)

func WithGetNotesFontAddOnFail(addOnFail domino.Boolean) getNotesFontParam {
	return func(c *getNotesFontParams) {
		c.addOnFail = &addOnFail
	}
}

func (r NotesRichTextItem) GetNotesFont(faceName domino.String, params ...getNotesFontParam) (domino.Long, error) {
	paramsStruct := &getNotesFontParams{}
	paramsOrdered := []interface{}{faceName}

	for _, p := range params {
		p(paramsStruct)
	}

	if paramsStruct.addOnFail != nil {
		paramsOrdered = append(paramsOrdered, *paramsStruct.addOnFail)
	}
	val, err := r.Com().CallMethod("GetNotesFont", paramsOrdered...)
	return helpers.CastValue[domino.Long](val), err
}

/* https://help.hcl-software.com/dom_designer/10.0.1/basic/H_GETUNFORMATTEDTEXT_METHOD.html */
func (r NotesRichTextItem) GetUnformattedText() (domino.String, error) {
	val, err := r.Com().CallMethod("GetUnformattedText")
	return helpers.CastValue[domino.String](val), err
}

/* https://help.hcl-software.com/dom_designer/10.0.1/basic/H_UPDATE_METHOD_RTITEM.html */
func (r NotesRichTextItem) Update() error {
	_, err := r.Com().CallMethod("Update")
	return err
}
