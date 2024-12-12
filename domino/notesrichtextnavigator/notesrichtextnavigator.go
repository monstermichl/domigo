/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_NOTESRICHTEXTNAVIGATOR_CLASS.html */
package notesrichtextnavigator

import (
	"domigo/domino"
	"domigo/helpers"

	ole "github.com/go-ole/go-ole"
)

type RtElementType = domino.Long

const (
	RTELEM_TYPE_DOCLINK        RtElementType = 5
	RTELEM_TYPE_FILEATTACHMENT RtElementType = 8
	RTELEM_TYPE_OLE            RtElementType = 9
	RTELEM_TYPE_SECTION        RtElementType = 6
	RTELEM_TYPE_TABLE          RtElementType = 1
	RTELEM_TYPE_TABLECELL      RtElementType = 7
	RTELEM_TYPE_TEXTPARAGRAPH  RtElementType = 4
	RTELEM_TYPE_TEXTRUN        RtElementType = 3
)

type NotesRichTextNavigator struct {
	domino.NotesStruct
}

func New(dispatchPtr *ole.IDispatch) NotesRichTextNavigator {
	return NotesRichTextNavigator{domino.NewNotesStruct(dispatchPtr)}
}

/* --------------------------------- Properties --------------------------------- */

/* --------------------------------- Methods ------------------------------------ */
/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_CLONE_METHOD_RTNAV.html */
func (r NotesRichTextNavigator) Clone() (NotesRichTextNavigator, error) {
	dispatchPtr, err := r.Com().CallObjectMethod("Clone")
	return New(dispatchPtr), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_FINDFIRSTELEMENT_METHOD_RTNAV.html */
func (r NotesRichTextNavigator) FindFirstElement(elementType RtElementType) (domino.Boolean, error) {
	val, err := r.Com().CallMethod("FindFirstElement", elementType)
	return helpers.CastValue[domino.Boolean](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_FINDFIRSTSTRING_METHOD_RTNAV.html */
type findFirstStringParams struct {
	options *domino.Long
}

type findFirstStringParam func(*findFirstStringParams)

func WithFindFirstStringOptions(options domino.Long) findFirstStringParam {
	return func(c *findFirstStringParams) {
		c.options = &options
	}
}

func (r NotesRichTextNavigator) FindFirstString(target domino.String, params ...findFirstStringParam) (domino.Boolean, error) {
	paramsStruct := &findFirstStringParams{}
	paramsOrdered := []interface{}{target}

	for _, p := range params {
		p(paramsStruct)
	}

	if paramsStruct.options != nil {
		paramsOrdered = append(paramsOrdered, *paramsStruct.options)
	}
	val, err := r.Com().CallMethod("FindFirstString", paramsOrdered...)
	return helpers.CastValue[domino.Boolean](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_FINDLASTELEMENT_METHOD_RTNAV.html */
func (r NotesRichTextNavigator) FindLastElement(elementType RtElementType) (domino.Boolean, error) {
	val, err := r.Com().CallMethod("FindLastElement", elementType)
	return helpers.CastValue[domino.Boolean](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_FINDNEXTELEMENT_METHOD_RTNAV.html */
type findNextElementParams struct {
	elementType *RtElementType
	occurrence  *domino.Long
}

type findNextElementParam func(*findNextElementParams)

func WithFindNextElementType(elementType RtElementType) findNextElementParam {
	return func(c *findNextElementParams) {
		c.elementType = &elementType
	}
}

func WithFindNextElementOccurrence(occurrence domino.Long) findNextElementParam {
	return func(c *findNextElementParams) {
		c.occurrence = &occurrence
	}
}

func (r NotesRichTextNavigator) FindNextElement(params ...findNextElementParam) (domino.Boolean, error) {
	paramsStruct := &findNextElementParams{}
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
	val, err := r.Com().CallMethod("FindNextElement", paramsOrdered...)
	return helpers.CastValue[domino.Boolean](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_FINDNEXTSTRING_METHOD_RTNAV.html */
type findNextStringParams struct {
	options *domino.Long
}

type findNextStringParam func(*findNextStringParams)

func WithFindNextStringOptions(options domino.Long) findNextStringParam {
	return func(c *findNextStringParams) {
		c.options = &options
	}
}

func (r NotesRichTextNavigator) FindNextString(target domino.String, params ...findNextStringParam) (domino.Boolean, error) {
	paramsStruct := &findNextStringParams{}
	paramsOrdered := []interface{}{target}

	for _, p := range params {
		p(paramsStruct)
	}

	if paramsStruct.options != nil {
		paramsOrdered = append(paramsOrdered, *paramsStruct.options)
	}
	val, err := r.Com().CallMethod("FindNextString", paramsOrdered...)
	return helpers.CastValue[domino.Boolean](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_FINDNTHELEMENT_METHOD_RTNAV.html */
type findNthElementParams struct {
	occurrence *domino.Long
}

type findNthElementParam func(*findNthElementParams)

func WithFindNthElementOccurrence(occurrence domino.Long) findNthElementParam {
	return func(c *findNthElementParams) {
		c.occurrence = &occurrence
	}
}

func (r NotesRichTextNavigator) FindNthElement(elementType RtElementType, params ...findNthElementParam) (domino.Boolean, error) {
	paramsStruct := &findNthElementParams{}
	paramsOrdered := []interface{}{elementType}

	for _, p := range params {
		p(paramsStruct)
	}

	if paramsStruct.occurrence != nil {
		paramsOrdered = append(paramsOrdered, *paramsStruct.occurrence)
	}
	val, err := r.Com().CallMethod("FindNthElement", paramsOrdered...)
	return helpers.CastValue[domino.Boolean](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GETELEMENT_METHOD_RTNAV.html */
/* TODO: Implement later. */
// func (r NotesRichTextNavigator) GetElement() (domino.Object, error) {
// 	dispatchPtr, err := r.Com().CallObjectMethod("GetElement")
// 	return unknown.New(dispatchPtr), err
// }

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GETFIRSTELEMENT_METHOD_RTNAV.html */
/* TODO: Implement later. */
// func (r NotesRichTextNavigator) GetFirstElement(elementType RtElementType) (domino.Object, error) {
// 	dispatchPtr, err := r.Com().CallObjectMethod("GetFirstElement", elementType)
// 	return unknown.New(dispatchPtr), err
// }

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GETLASTELEMENT_METHOD_RTNAV.html */
/* TODO: Implement later. */
// func (r NotesRichTextNavigator) GetLastElement(elementType RtElementType) (domino.Object, error) {
// 	dispatchPtr, err := r.Com().CallObjectMethod("GetLastElement", elementType)
// 	return unknown.New(dispatchPtr), err
// }

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GETNEXTELEMENT_METHOD_RTNAV.html */
/* TODO: Implement later. */
// type getNextElementParams struct {
// 	elementType *RtElementType
// 	occurrence  *domino.Long
// }

// type getNextElementParam func(*getNextElementParams)

// func WithGetNextElementType(elementType RtElementType) getNextElementParam {
// 	return func(c *getNextElementParams) {
// 		c.elementType = &elementType
// 	}
// }

// func WithGetNextElementOccurrence(occurrence domino.Long) getNextElementParam {
// 	return func(c *getNextElementParams) {
// 		c.occurrence = &occurrence
// 	}
// }

// func (r NotesRichTextNavigator) GetNextElement(params ...getNextElementParam) (any, error) {
// 	paramsStruct := &getNextElementParams{}
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
// 	dispatchPtr, err := r.Com().CallObjectMethod("GetNextElement", paramsOrdered...)
// 	return unknown.New(dispatchPtr), err
// }

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GETNTHELEMENT_METHOD_RTNAV.html */
/* TODO: Implement later. */
// type getNthElementParams struct {
// 	occurrence *domino.Long
// }

// type getNthElementParam func(*getNthElementParams)

// func WithGetNthElementOccurrence(occurrence domino.Long) getNthElementParam {
// 	return func(c *getNthElementParams) {
// 		c.occurrence = &occurrence
// 	}
// }

// func (r NotesRichTextNavigator) GetNthElement(elementType RtElementType, params ...getNthElementParam) (any, error) {
// 	paramsStruct := &getNthElementParams{}
// 	paramsOrdered := []interface{}{elementType}

// 	for _, p := range params {
// 		p(paramsStruct)
// 	}

// 	if paramsStruct.occurrence != nil {
// 		paramsOrdered = append(paramsOrdered, *paramsStruct.occurrence)
// 	}
// 	dispatchPtr, err := r.Com().CallObjectMethod("GetNthElement", paramsOrdered...)
// 	return unknown.New(dispatchPtr), err
// }

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SETCHAROFFSET_METHOD_RTNAV.html */
func (r NotesRichTextNavigator) SetCharOffset(offset domino.Integer) error {
	_, err := r.Com().CallMethod("SetCharOffset", offset)
	return err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SETPOSITION_METHOD_RTNAV.html */
func (r NotesRichTextNavigator) SetPosition(element domino.NotesConnector) error {
	_, err := r.Com().CallMethod("SetPosition", element.Com().Dispatch())
	return err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SETPOSITIONATEND_METHOD_RTNAV.html */
func (r NotesRichTextNavigator) SetPositionAtEnd(element domino.NotesConnector) error {
	_, err := r.Com().CallMethod("SetPositionAtEnd", element.Com().Dispatch())
	return err
}
