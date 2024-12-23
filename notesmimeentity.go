/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_NOTESMIMEENTITY_CLASS_OVERVIEW.html */
package domigo

import (
	"github.com/monstermichl/domigo/internal/com"

	ole "github.com/go-ole/go-ole"
)

type NotesMIMEEntityEncoding = Long

const (
	NOTESMIMEENTITY_ENC_BASE64           NotesMIMEEntityEncoding = 1727
	NOTESMIMEENTITY_ENC_EXTENSION        NotesMIMEEntityEncoding = 1731
	NOTESMIMEENTITY_ENC_IDENTITY_7BIT    NotesMIMEEntityEncoding = 1728
	NOTESMIMEENTITY_ENC_IDENTITY_8BIT    NotesMIMEEntityEncoding = 1729
	NOTESMIMEENTITY_ENC_IDENTITY_BINARY  NotesMIMEEntityEncoding = 1730
	NOTESMIMEENTITY_ENC_NONE             NotesMIMEEntityEncoding = 1725
	NOTESMIMEENTITY_ENC_QUOTED_PRINTABLE NotesMIMEEntityEncoding = 1726
)

type NotesMIMEEntity struct {
	NotesStruct
}

func NewNotesMIMEEntity(dispatchPtr *ole.IDispatch) NotesMIMEEntity {
	return NotesMIMEEntity{NewNotesStruct(dispatchPtr)}
}

/* --------------------------------- Properties --------------------------------- */
/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_BOUNDARYEND_PROPERTY_MIMEENTITY.html */
func (m NotesMIMEEntity) BoundaryEnd() (String, error) {
	return getComProperty[String](m, "BoundaryEnd")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_BOUNDARYSTART_PROPERTY_MIMEENTITY.html */
func (m NotesMIMEEntity) BoundaryStart() (String, error) {
	return getComProperty[String](m, "BoundaryStart")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_CHARSET_PROPERTY_MIMEENTITY.html */
func (m NotesMIMEEntity) Charset() (String, error) {
	return getComProperty[String](m, "Charset")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_CONTENTASTEXT_PROPERTY_MIMEENTITY.html */
func (m NotesMIMEEntity) ContentAsText() (String, error) {
	return getComProperty[String](m, "ContentAsText")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_CONTENTSUBTYPE_PROPERTY_MIMEENTITY.html */
func (m NotesMIMEEntity) ContentSubType() (String, error) {
	return getComProperty[String](m, "ContentSubType")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_CONTENTTYPE_PROPERTY_MIMEENTITY.html */
func (m NotesMIMEEntity) ContentType() (String, error) {
	return getComProperty[String](m, "ContentType")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ENCODING_PROPERTY_MIMEENTITY.html */
func (m NotesMIMEEntity) Encoding() (NotesMIMEEntityEncoding, error) {
	return getComProperty[NotesMIMEEntityEncoding](m, "Encoding")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_HEADEROBJECTS_PROPERTY_MIMEENTITY.html */
func (m NotesMIMEEntity) HeaderObjects() ([]NotesMIMEHeader, error) {
	return com.GetObjectArrayProperty(m.com(), NewNotesMIMEHeader, "HeaderObjects")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_HEADERS_PROPERTY_MIMEENTITY.html */
func (m NotesMIMEEntity) Headers() (String, error) {
	return getComProperty[String](m, "Headers")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_PREAMBLE_PROPERTY_MIMEENTITY.html */
func (m NotesMIMEEntity) Preamble() (String, error) {
	return getComProperty[String](m, "Preamble")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_PREAMBLE_PROPERTY_MIMEENTITY.html */
func (m NotesMIMEEntity) SetPreamble(v String) error {
	return putComProperty(m, "Preamble", v)
}

/* --------------------------------- Methods ------------------------------------ */
/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_CREATECHILDENTITY_METHOD_MIMEENTITY.html */
type notesMIMEEntityCreateChildEntityParams struct {
	nextSibling *NotesMIMEEntity
}

type notesMIMEEntityCreateChildEntityParam func(*notesMIMEEntityCreateChildEntityParams)

func WithNotesMIMEEntityCreateChildEntityNextSibling(nextSibling NotesMIMEEntity) notesMIMEEntityCreateChildEntityParam {
	return func(c *notesMIMEEntityCreateChildEntityParams) {
		c.nextSibling = &nextSibling
	}
}

func (m NotesMIMEEntity) CreateChildEntity(params ...notesMIMEEntityCreateChildEntityParam) (NotesMIMEEntity, error) {
	paramsStruct := &notesMIMEEntityCreateChildEntityParams{}
	paramsOrdered := []interface{}{}

	for _, p := range params {
		p(paramsStruct)
	}

	if paramsStruct.nextSibling != nil {
		paramsOrdered = append(paramsOrdered, *paramsStruct.nextSibling)
	}
	return callComObjectMethod(m, NewNotesMIMEEntity, "CreateChildEntity", paramsOrdered...)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_CREATEHEADER_METHOD_MIMEENTITY.html */
func (m NotesMIMEEntity) CreateHeader(headerName String) (NotesMIMEHeader, error) {
	return callComObjectMethod(m, NewNotesMIMEHeader, "CreateHeader", headerName)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_CREATEPARENTENTITY_METHOD_MIMEENTITY.html */
func (m NotesMIMEEntity) CreateParentEntity() (NotesMIMEEntity, error) {
	return callComObjectMethod(m, NewNotesMIMEEntity, "CreateParentEntity")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_DECODECONTENT_METHOD_MIMEENTITY.html */
func (m NotesMIMEEntity) DecodeContent() error {
	return callComVoidMethod(m, "DecodeContent")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ENCODECONTENT_METHOD_MIMEENTITY.html */
func (m NotesMIMEEntity) EncodeContent(encoding NotesMIMEEntityEncoding) error {
	return callComVoidMethod(m, "EncodeContent", encoding)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GETCONTENTASBYTES_METHOD_MIMEENTITY.html */
type notesMIMEEntityGetContentAsBytesParams struct {
	decoded *Boolean
}

type notesMIMEEntityGetContentAsBytesParam func(*notesMIMEEntityGetContentAsBytesParams)

func WithNotesMIMEEntityGetContentAsBytesDecoded(decoded Boolean) notesMIMEEntityGetContentAsBytesParam {
	return func(c *notesMIMEEntityGetContentAsBytesParams) {
		c.decoded = &decoded
	}
}

func (m NotesMIMEEntity) GetContentAsBytes(stream NotesStream, params ...notesMIMEEntityGetContentAsBytesParam) error {
	paramsStruct := &notesMIMEEntityGetContentAsBytesParams{}
	paramsOrdered := []interface{}{stream}

	for _, p := range params {
		p(paramsStruct)
	}

	if paramsStruct.decoded != nil {
		paramsOrdered = append(paramsOrdered, *paramsStruct.decoded)
	}
	return callComVoidMethod(m, "GetContentAsBytes", paramsOrdered...)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GETCONTENTASTEXT_METHOD_MIMEENTITY.html */
type notesMIMEEntityGetContentAsTextParams struct {
	decoded *Boolean
}

type notesMIMEEntityGetContentAsTextParam func(*notesMIMEEntityGetContentAsTextParams)

func WithNotesMIMEEntityGetContentAsTextDecoded(decoded Boolean) notesMIMEEntityGetContentAsTextParam {
	return func(c *notesMIMEEntityGetContentAsTextParams) {
		c.decoded = &decoded
	}
}

func (m NotesMIMEEntity) GetContentAsText(stream NotesStream, params ...notesMIMEEntityGetContentAsTextParam) error {
	paramsStruct := &notesMIMEEntityGetContentAsTextParams{}
	paramsOrdered := []interface{}{stream}

	for _, p := range params {
		p(paramsStruct)
	}

	if paramsStruct.decoded != nil {
		paramsOrdered = append(paramsOrdered, *paramsStruct.decoded)
	}
	return callComVoidMethod(m, "GetContentAsText", paramsOrdered...)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GETENTITYASTEXT_METHOD_MIMEENTITY.html */
type notesMIMEEntityGetEntityAsTextParams struct {
	headerFilters *[]String
	inclusive     *Boolean
}

type notesMIMEEntityGetEntityAsTextParam func(*notesMIMEEntityGetEntityAsTextParams)

func WithNotesMIMEEntityGetEntityAsTextHeaderFilters(headerFilters []String) notesMIMEEntityGetEntityAsTextParam {
	return func(c *notesMIMEEntityGetEntityAsTextParams) {
		c.headerFilters = &headerFilters
	}
}

func WithNotesMIMEEntityGetEntityAsTextInclusive(inclusive Boolean) notesMIMEEntityGetEntityAsTextParam {
	return func(c *notesMIMEEntityGetEntityAsTextParams) {
		c.inclusive = &inclusive
	}
}

func (m NotesMIMEEntity) GetEntityAsText(stream NotesStream, params ...notesMIMEEntityGetEntityAsTextParam) error {
	paramsStruct := &notesMIMEEntityGetEntityAsTextParams{}
	paramsOrdered := []interface{}{stream}

	for _, p := range params {
		p(paramsStruct)
	}

	if paramsStruct.headerFilters != nil {
		paramsOrdered = append(paramsOrdered, *paramsStruct.headerFilters)
		if paramsStruct.inclusive != nil {
			paramsOrdered = append(paramsOrdered, *paramsStruct.inclusive)
		}
	}
	return callComVoidMethod(m, "GetEntityAsText", paramsOrdered...)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GETFIRSTCHILDENTITY_METHOD_MIMEENTITY.html */
func (m NotesMIMEEntity) GetFirstChildEntity() (NotesMIMEEntity, error) {
	return callComObjectMethod(m, NewNotesMIMEEntity, "GetFirstChildEntity")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GETNEXTENTITY_METHOD_MIMEENTITY.html */
type notesMIMEEntityGetNextEntityParams struct {
	search *Integer
}

type notesMIMEEntityGetNextEntityParam func(*notesMIMEEntityGetNextEntityParams)

func WithNotesMIMEEntityGetNextEntitySearch(search Integer) notesMIMEEntityGetNextEntityParam {
	return func(c *notesMIMEEntityGetNextEntityParams) {
		c.search = &search
	}
}

func (m NotesMIMEEntity) GetNextEntity(params ...notesMIMEEntityGetNextEntityParam) (NotesMIMEEntity, error) {
	paramsStruct := &notesMIMEEntityGetNextEntityParams{}
	paramsOrdered := []interface{}{}

	for _, p := range params {
		p(paramsStruct)
	}

	if paramsStruct.search != nil {
		paramsOrdered = append(paramsOrdered, *paramsStruct.search)
	}
	return callComObjectMethod(m, NewNotesMIMEEntity, "GetNextEntity", paramsOrdered...)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GETNEXTSIBLING_METHOD_MIMEENTITY.html */
func (m NotesMIMEEntity) GetNextSibling() (NotesMIMEEntity, error) {
	return callComObjectMethod(m, NewNotesMIMEEntity, "GetNextSibling")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GETNTHHEADER_METHOD_MIMEENTITY.html */
type notesMIMEEntityGetNthHeaderParams struct {
	instance *Integer
}

type notesMIMEEntityGetNthHeaderParam func(*notesMIMEEntityGetNthHeaderParams)

func WithNotesMIMEEntityGetNthHeaderInstance(instance Integer) notesMIMEEntityGetNthHeaderParam {
	return func(c *notesMIMEEntityGetNthHeaderParams) {
		c.instance = &instance
	}
}

func (m NotesMIMEEntity) GetNthHeader(headerName String, params ...notesMIMEEntityGetNthHeaderParam) (NotesMIMEHeader, error) {
	paramsStruct := &notesMIMEEntityGetNthHeaderParams{}
	paramsOrdered := []interface{}{headerName}

	for _, p := range params {
		p(paramsStruct)
	}

	if paramsStruct.instance != nil {
		paramsOrdered = append(paramsOrdered, *paramsStruct.instance)
	}
	return callComObjectMethod(m, NewNotesMIMEHeader, "GetNthHeader", paramsOrdered...)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GETPARENTENTITY_METHOD_MIMEENTITY.html */
func (m NotesMIMEEntity) GetParentEntity() (NotesMIMEEntity, error) {
	return callComObjectMethod(m, NewNotesMIMEEntity, "GetParentEntity")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GETPREVENTITY_METHOD_MIMEENTITY.html */
type notesMIMEEntityGetPrevEntityParams struct {
	search *Integer
}

type notesMIMEEntityGetPrevEntityParam func(*notesMIMEEntityGetPrevEntityParams)

func WithNotesMIMEEntityGetPrevEntitySearch(search Integer) notesMIMEEntityGetPrevEntityParam {
	return func(c *notesMIMEEntityGetPrevEntityParams) {
		c.search = &search
	}
}

func (m NotesMIMEEntity) GetPrevEntity(params ...notesMIMEEntityGetPrevEntityParam) (NotesMIMEEntity, error) {
	paramsStruct := &notesMIMEEntityGetPrevEntityParams{}
	paramsOrdered := []interface{}{}

	for _, p := range params {
		p(paramsStruct)
	}

	if paramsStruct.search != nil {
		paramsOrdered = append(paramsOrdered, *paramsStruct.search)
	}
	return callComObjectMethod(m, NewNotesMIMEEntity, "GetPrevEntity", paramsOrdered...)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GETPREVSIBLING_METHOD_MIMEENTITY.html */
func (m NotesMIMEEntity) GetPrevSibling() (NotesMIMEEntity, error) {
	return callComObjectMethod(m, NewNotesMIMEEntity, "GetPrevSibling")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GETSOMEHEADERS_METHOD_MIMEENTITY.html */
type notesMIMEEntityGetSomeHeadersParams struct {
	headerFilters *[]String
	inclusive     *Boolean
}

type notesMIMEEntityGetSomeHeadersParam func(*notesMIMEEntityGetSomeHeadersParams)

func WithNotesMIMEEntityGetSomeHeadersHeaderFilters(headerFilters []String) notesMIMEEntityGetSomeHeadersParam {
	return func(c *notesMIMEEntityGetSomeHeadersParams) {
		c.headerFilters = &headerFilters
	}
}

func WithNotesMIMEEntityGetSomeHeadersInclusive(inclusive Boolean) notesMIMEEntityGetSomeHeadersParam {
	return func(c *notesMIMEEntityGetSomeHeadersParams) {
		c.inclusive = &inclusive
	}
}

func (m NotesMIMEEntity) GetSomeHeaders(params ...notesMIMEEntityGetSomeHeadersParam) (String, error) {
	paramsStruct := &notesMIMEEntityGetSomeHeadersParams{}
	paramsOrdered := []interface{}{}

	for _, p := range params {
		p(paramsStruct)
	}

	if paramsStruct.headerFilters != nil {
		paramsOrdered = append(paramsOrdered, *paramsStruct.headerFilters)
		if paramsStruct.inclusive != nil {
			paramsOrdered = append(paramsOrdered, *paramsStruct.inclusive)
		}
	}
	return callComMethod[String](m, "GetSomeHeaders", paramsOrdered...)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_REMOVE_METHOD_MIMEENTITY.html */
func (m NotesMIMEEntity) Remove() error {
	return callComVoidMethod(m, "Remove")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SETCONTENTFROMBYTES_METHOD_MIMEENTITY.html */
func (m NotesMIMEEntity) SetContentFromBytes(stream NotesStream, contentType String, encoding NotesMIMEEntityEncoding) error {
	return callComVoidMethod(m, "SetContentFromBytes", stream, contentType, encoding)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SETCONTENTFROMTEXT_METHOD_MIMEENTITY.html */
func (m NotesMIMEEntity) SetContentFromText(stream NotesStream, contentType String, encoding NotesMIMEEntityEncoding) error {
	return callComVoidMethod(m, "SetContentFromText", stream, contentType, encoding)
}
