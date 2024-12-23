/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_NOTESRICHTEXTNAVIGATOR_CLASS.html */
package domigo

import (
	ole "github.com/go-ole/go-ole"
)

type NotesRichTextNavigatorRtElementType Long

const (
	NOTESRICHTEXTNAVIGATOR_RTELEM_TYPE_DOCLINK        NotesRichTextNavigatorRtElementType = 5
	NOTESRICHTEXTNAVIGATOR_RTELEM_TYPE_FILEATTACHMENT NotesRichTextNavigatorRtElementType = 8
	NOTESRICHTEXTNAVIGATOR_RTELEM_TYPE_OLE            NotesRichTextNavigatorRtElementType = 9
	NOTESRICHTEXTNAVIGATOR_RTELEM_TYPE_SECTION        NotesRichTextNavigatorRtElementType = 6
	NOTESRICHTEXTNAVIGATOR_RTELEM_TYPE_TABLE          NotesRichTextNavigatorRtElementType = 1
	NOTESRICHTEXTNAVIGATOR_RTELEM_TYPE_TABLECELL      NotesRichTextNavigatorRtElementType = 7
	NOTESRICHTEXTNAVIGATOR_RTELEM_TYPE_TEXTPARAGRAPH  NotesRichTextNavigatorRtElementType = 4
	NOTESRICHTEXTNAVIGATOR_RTELEM_TYPE_TEXTRUN        NotesRichTextNavigatorRtElementType = 3
)

type NotesRichTextNavigator struct {
	NotesStruct
}

func NewNotesRichTextNavigator(dispatchPtr *ole.IDispatch) NotesRichTextNavigator {
	return NotesRichTextNavigator{NewNotesStruct(dispatchPtr)}
}

/* --------------------------------- Properties --------------------------------- */
/* --------------------------------- Methods ------------------------------------ */
/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_CLONE_METHOD_RTNAV.html */
func (r NotesRichTextNavigator) Clone() (NotesRichTextNavigator, error) {
	return callComObjectMethod(r, NewNotesRichTextNavigator, "Clone")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_FINDFIRSTELEMENT_METHOD_RTNAV.html */
func (r NotesRichTextNavigator) FindFirstElement(elementType NotesRichTextNavigatorRtElementType) (Boolean, error) {
	return callComMethod[Boolean](r, "FindFirstElement", elementType)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_FINDFIRSTSTRING_METHOD_RTNAV.html */
type notesRichTextNavigatorFindFirstStringParams struct {
	options *Long
}

type notesRichTextNavigatorFindFirstStringParam func(*notesRichTextNavigatorFindFirstStringParams)

func WithNotesRichTextNavigatorFindFirstStringOptions(options Long) notesRichTextNavigatorFindFirstStringParam {
	return func(c *notesRichTextNavigatorFindFirstStringParams) {
		c.options = &options
	}
}

func (r NotesRichTextNavigator) FindFirstString(target String, params ...notesRichTextNavigatorFindFirstStringParam) (Boolean, error) {
	paramsStruct := &notesRichTextNavigatorFindFirstStringParams{}
	paramsOrdered := []interface{}{target}

	for _, p := range params {
		p(paramsStruct)
	}

	if paramsStruct.options != nil {
		paramsOrdered = append(paramsOrdered, *paramsStruct.options)
	}
	return callComMethod[Boolean](r, "FindFirstString", paramsOrdered...)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_FINDLASTELEMENT_METHOD_RTNAV.html */
func (r NotesRichTextNavigator) FindLastElement(elementType NotesRichTextNavigatorRtElementType) (Boolean, error) {
	return callComMethod[Boolean](r, "FindLastElement", elementType)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_FINDNEXTELEMENT_METHOD_RTNAV.html */
type notesRichTextNavigatorFindNextElementParams struct {
	elementType *NotesRichTextNavigatorRtElementType
	occurrence  *Long
}

type notesRichTextNavigatorFindNextElementParam func(*notesRichTextNavigatorFindNextElementParams)

func WithNotesRichTextNavigatorFindNextElementType(elementType NotesRichTextNavigatorRtElementType) notesRichTextNavigatorFindNextElementParam {
	return func(c *notesRichTextNavigatorFindNextElementParams) {
		c.elementType = &elementType
	}
}

func WithNotesRichTextNavigatorFindNextElementOccurrence(occurrence Long) notesRichTextNavigatorFindNextElementParam {
	return func(c *notesRichTextNavigatorFindNextElementParams) {
		c.occurrence = &occurrence
	}
}

func (r NotesRichTextNavigator) FindNextElement(params ...notesRichTextNavigatorFindNextElementParam) (Boolean, error) {
	paramsStruct := &notesRichTextNavigatorFindNextElementParams{}
	paramsOrdered := []interface{}{}

	for _, p := range params {
		p(paramsStruct)
	}

	if paramsStruct.elementType != nil {
		paramsOrdered = append(paramsOrdered, *paramsStruct.elementType)
		if paramsStruct.occurrence != nil {
			paramsOrdered = append(paramsOrdered, *paramsStruct.occurrence)
		}
	}
	return callComMethod[Boolean](r, "FindNextElement", paramsOrdered...)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_FINDNEXTSTRING_METHOD_RTNAV.html */
type notesRichTextNavigatorFindNextStringParams struct {
	options *Long
}

type notesRichTextNavigatorFindNextStringParam func(*notesRichTextNavigatorFindNextStringParams)

func WithNotesRichTextNavigatorFindNextStringOptions(options Long) notesRichTextNavigatorFindNextStringParam {
	return func(c *notesRichTextNavigatorFindNextStringParams) {
		c.options = &options
	}
}

func (r NotesRichTextNavigator) FindNextString(target String, params ...notesRichTextNavigatorFindNextStringParam) (Boolean, error) {
	paramsStruct := &notesRichTextNavigatorFindNextStringParams{}
	paramsOrdered := []interface{}{target}

	for _, p := range params {
		p(paramsStruct)
	}

	if paramsStruct.options != nil {
		paramsOrdered = append(paramsOrdered, *paramsStruct.options)
	}
	return callComMethod[Boolean](r, "FindNextString", paramsOrdered...)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_FINDNTHELEMENT_METHOD_RTNAV.html */
type notesRichTextNavigatorFindNthElementParams struct {
	occurrence *Long
}

type notesRichTextNavigatorFindNthElementParam func(*notesRichTextNavigatorFindNthElementParams)

func WithNotesRichTextNavigatorFindNthElementOccurrence(occurrence Long) notesRichTextNavigatorFindNthElementParam {
	return func(c *notesRichTextNavigatorFindNthElementParams) {
		c.occurrence = &occurrence
	}
}

func (r NotesRichTextNavigator) FindNthElement(elementType NotesRichTextNavigatorRtElementType, params ...notesRichTextNavigatorFindNthElementParam) (Boolean, error) {
	paramsStruct := &notesRichTextNavigatorFindNthElementParams{}
	paramsOrdered := []interface{}{elementType}

	for _, p := range params {
		p(paramsStruct)
	}

	if paramsStruct.occurrence != nil {
		paramsOrdered = append(paramsOrdered, *paramsStruct.occurrence)
	}
	return callComMethod[Boolean](r, "FindNthElement", paramsOrdered...)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GETELEMENT_METHOD_RTNAV.html */
/* TODO: Implement later. */
// func (r NotesRichTextNavigator) GetElement() (Object, error) {
// 	dispatchPtr, err := CallComObjectMethod(r, "GetElement")
// 	return unknown.New(dispatchPtr), err
// }

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GETFIRSTELEMENT_METHOD_RTNAV.html */
/* TODO: Implement later. */
// func (r NotesRichTextNavigator) GetFirstElement(elementType RtElementType) (Object, error) {
// 	dispatchPtr, err := CallComObjectMethod(r, "GetFirstElement", elementType)
// 	return unknown.New(dispatchPtr), err
// }

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GETLASTELEMENT_METHOD_RTNAV.html */
/* TODO: Implement later. */
// func (r NotesRichTextNavigator) GetLastElement(elementType RtElementType) (Object, error) {
// 	dispatchPtr, err := CallComObjectMethod(r, "GetLastElement", elementType)
// 	return unknown.New(dispatchPtr), err
// }

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GETNEXTELEMENT_METHOD_RTNAV.html */
/* TODO: Implement later. */
// type notesRichTextNavigatorGetNextElementParams struct {
// 	elementType *RtElementType
// 	occurrence  *Long
// }

// type notesRichTextNavigatorGetNextElementParam func(*notesRichTextNavigatorGetNextElementParams)

// func WithNotesRichTextNavigatorGetNextElementType(elementType RtElementType) notesRichTextNavigatorGetNextElementParam {
// 	return func(c *notesRichTextNavigatorGetNextElementParams) {
// 		c.elementType = &elementType
// 	}
// }

// func WithNotesRichTextNavigatorGetNextElementOccurrence(occurrence Long) notesRichTextNavigatorGetNextElementParam {
// 	return func(c *notesRichTextNavigatorGetNextElementParams) {
// 		c.occurrence = &occurrence
// 	}
// }

// func (r NotesRichTextNavigator) GetNextElement(params ...notesRichTextNavigatorGetNextElementParam) (Object, error) {
// 	paramsStruct := &notesRichTextNavigatorGetNextElementParams{}
// 	paramsOrdered := []interface{}{}

// 	for _, p := range params {
// 		p(paramsStruct)
// 	}

// 	if paramsStruct.elementType != nil {
// 		paramsOrdered = append(paramsOrdered, *paramsStruct.elementType)
// 		if paramsStruct.occurrence != nil {
// 			paramsOrdered = append(paramsOrdered, *paramsStruct.occurrence)
// 		}
// 	}
// 	dispatchPtr, err := CallComObjectMethod(r, "GetNextElement", paramsOrdered...)
// 	return New(Object, error)(dispatchPtr), err
// }

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GETNTHELEMENT_METHOD_RTNAV.html */
/* TODO: Implement later. */
// func (r NotesRichTextNavigator) GetNthElement(elementType RtElementType, occurrence Long) (Object, error) {
// 	dispatchPtr, err := CallComObjectMethod(r, "GetNthElement", elementType, occurrence)
// 	return New(Object, error)(dispatchPtr), err
// }

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SETCHAROFFSET_METHOD_RTNAV.html */
func (r NotesRichTextNavigator) SetCharOffset(offset Integer) error {
	return callComVoidMethod(r, "SetCharOffset", offset)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SETPOSITION_METHOD_RTNAV.html */
func (r NotesRichTextNavigator) SetPosition(element notesStruct) error {
	return callComVoidMethod(r, "SetPosition", element)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SETPOSITIONATEND_METHOD_RTNAV.html */
func (r NotesRichTextNavigator) SetPositionAtEnd(element notesStruct) error {
	return callComVoidMethod(r, "SetPositionAtEnd", element)
}
