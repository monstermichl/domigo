/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_NOTESMIMEENTITY_CLASS_OVERVIEW.html */
package domigo

import (
	"github.com/monstermichl/domigo/internal/com"
	"github.com/monstermichl/domigo/internal/helpers"

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
	val, err := getComProperty(m, "BoundaryEnd")
	return helpers.CastValue[String](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_BOUNDARYSTART_PROPERTY_MIMEENTITY.html */
func (m NotesMIMEEntity) BoundaryStart() (String, error) {
	val, err := getComProperty(m, "BoundaryStart")
	return helpers.CastValue[String](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_CHARSET_PROPERTY_MIMEENTITY.html */
func (m NotesMIMEEntity) Charset() (String, error) {
	val, err := getComProperty(m, "Charset")
	return helpers.CastValue[String](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_CONTENTASTEXT_PROPERTY_MIMEENTITY.html */
func (m NotesMIMEEntity) ContentAsText() (String, error) {
	val, err := getComProperty(m, "ContentAsText")
	return helpers.CastValue[String](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_CONTENTSUBTYPE_PROPERTY_MIMEENTITY.html */
func (m NotesMIMEEntity) ContentSubType() (String, error) {
	val, err := getComProperty(m, "ContentSubType")
	return helpers.CastValue[String](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_CONTENTTYPE_PROPERTY_MIMEENTITY.html */
func (m NotesMIMEEntity) ContentType() (String, error) {
	val, err := getComProperty(m, "ContentType")
	return helpers.CastValue[String](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ENCODING_PROPERTY_MIMEENTITY.html */
func (m NotesMIMEEntity) Encoding() (NotesMIMEEntityEncoding, error) {
	val, err := getComProperty(m, "Encoding")
	return helpers.CastValue[NotesMIMEEntityEncoding](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_HEADEROBJECTS_PROPERTY_MIMEENTITY.html */
func (m NotesMIMEEntity) HeaderObjects() ([]NotesMIMEHeader, error) {
	return com.GetObjectArrayProperty(m.com(), NewNotesMIMEHeader, "HeaderObjects")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_HEADERS_PROPERTY_MIMEENTITY.html */
func (m NotesMIMEEntity) Headers() (String, error) {
	val, err := getComProperty(m, "Headers")
	return helpers.CastValue[String](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_PREAMBLE_PROPERTY_MIMEENTITY.html */
func (m NotesMIMEEntity) Preamble() (String, error) {
	val, err := getComProperty(m, "Preamble")
	return helpers.CastValue[String](val), err
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
	dispatchPtr, err := callComObjectMethod(m, "CreateChildEntity", paramsOrdered...)
	return NewNotesMIMEEntity(dispatchPtr), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_CREATEHEADER_METHOD_MIMEENTITY.html */
func (m NotesMIMEEntity) CreateHeader(headerName String) (NotesMIMEHeader, error) {
	dispatchPtr, err := callComObjectMethod(m, "CreateHeader", headerName)
	return NewNotesMIMEHeader(dispatchPtr), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_CREATEPARENTENTITY_METHOD_MIMEENTITY.html */
func (m NotesMIMEEntity) CreateParentEntity() (NotesMIMEEntity, error) {
	dispatchPtr, err := callComObjectMethod(m, "CreateParentEntity")
	return NewNotesMIMEEntity(dispatchPtr), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_DECODECONTENT_METHOD_MIMEENTITY.html */
func (m NotesMIMEEntity) DecodeContent() error {
	_, err := callComMethod(m, "DecodeContent")
	return err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ENCODECONTENT_METHOD_MIMEENTITY.html */
func (m NotesMIMEEntity) EncodeContent(encoding NotesMIMEEntityEncoding) error {
	_, err := callComMethod(m, "EncodeContent", encoding)
	return err
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
	paramsOrdered := []interface{}{stream.com().Dispatch()}

	for _, p := range params {
		p(paramsStruct)
	}

	if paramsStruct.decoded != nil {
		paramsOrdered = append(paramsOrdered, *paramsStruct.decoded)
	}
	_, err := callComMethod(m, "GetContentAsBytes", paramsOrdered...)
	return err
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
	paramsOrdered := []interface{}{stream.com().Dispatch()}

	for _, p := range params {
		p(paramsStruct)
	}

	if paramsStruct.decoded != nil {
		paramsOrdered = append(paramsOrdered, *paramsStruct.decoded)
	}
	_, err := callComMethod(m, "GetContentAsText", paramsOrdered...)
	return err
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
	paramsOrdered := []interface{}{stream.com().Dispatch()}

	for _, p := range params {
		p(paramsStruct)
	}

	if paramsStruct.headerFilters != nil {
		paramsOrdered = append(paramsOrdered, *paramsStruct.headerFilters)
		if paramsStruct.inclusive != nil {
			paramsOrdered = append(paramsOrdered, *paramsStruct.inclusive)
		}
	}
	_, err := callComMethod(m, "GetEntityAsText", paramsOrdered...)
	return err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GETFIRSTCHILDENTITY_METHOD_MIMEENTITY.html */
func (m NotesMIMEEntity) GetFirstChildEntity() (NotesMIMEEntity, error) {
	dispatchPtr, err := callComObjectMethod(m, "GetFirstChildEntity")
	return NewNotesMIMEEntity(dispatchPtr), err
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
	dispatchPtr, err := callComObjectMethod(m, "GetNextEntity", paramsOrdered...)
	return NewNotesMIMEEntity(dispatchPtr), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GETNEXTSIBLING_METHOD_MIMEENTITY.html */
func (m NotesMIMEEntity) GetNextSibling() (NotesMIMEEntity, error) {
	dispatchPtr, err := callComObjectMethod(m, "GetNextSibling")
	return NewNotesMIMEEntity(dispatchPtr), err
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
	dispatchPtr, err := callComObjectMethod(m, "GetNthHeader", paramsOrdered...)
	return NewNotesMIMEHeader(dispatchPtr), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GETPARENTENTITY_METHOD_MIMEENTITY.html */
func (m NotesMIMEEntity) GetParentEntity() (NotesMIMEEntity, error) {
	dispatchPtr, err := callComObjectMethod(m, "GetParentEntity")
	return NewNotesMIMEEntity(dispatchPtr), err
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
	dispatchPtr, err := callComObjectMethod(m, "GetPrevEntity", paramsOrdered...)
	return NewNotesMIMEEntity(dispatchPtr), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GETPREVSIBLING_METHOD_MIMEENTITY.html */
func (m NotesMIMEEntity) GetPrevSibling() (NotesMIMEEntity, error) {
	dispatchPtr, err := callComObjectMethod(m, "GetPrevSibling")
	return NewNotesMIMEEntity(dispatchPtr), err
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
	val, err := callComMethod(m, "GetSomeHeaders", paramsOrdered...)
	return helpers.CastValue[String](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_REMOVE_METHOD_MIMEENTITY.html */
func (m NotesMIMEEntity) Remove() error {
	_, err := callComMethod(m, "Remove")
	return err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SETCONTENTFROMBYTES_METHOD_MIMEENTITY.html */
func (m NotesMIMEEntity) SetContentFromBytes(stream NotesStream, contentType String, encoding NotesMIMEEntityEncoding) error {
	_, err := callComMethod(m, "SetContentFromBytes", stream.com().Dispatch(), contentType, encoding)
	return err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SETCONTENTFROMTEXT_METHOD_MIMEENTITY.html */
func (m NotesMIMEEntity) SetContentFromText(stream NotesStream, contentType String, encoding NotesMIMEEntityEncoding) error {
	_, err := callComMethod(m, "SetContentFromText", stream.com().Dispatch(), contentType, encoding)
	return err
}
