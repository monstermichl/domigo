/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_NOTESMIMEENTITY_CLASS_OVERVIEW.html */
package notesmimeentity

import (
	"domigo/domino"
	"domigo/domino/com"
	"domigo/domino/notesmimeheader"
	"domigo/domino/notesstream"
	"domigo/helpers"

	ole "github.com/go-ole/go-ole"
)

type Encoding = domino.Long

const (
	ENC_BASE64           Encoding = 1727
	ENC_EXTENSION        Encoding = 1731
	ENC_IDENTITY_7BIT    Encoding = 1728
	ENC_IDENTITY_8BIT    Encoding = 1729
	ENC_IDENTITY_BINARY  Encoding = 1730
	ENC_NONE             Encoding = 1725
	ENC_QUOTED_PRINTABLE Encoding = 1726
)

type NotesMIMEEntity struct {
	domino.NotesStruct
}

func New(dispatchPtr *ole.IDispatch) NotesMIMEEntity {
	return NotesMIMEEntity{domino.NewNotesStruct(dispatchPtr)}
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_PARENT_PROPERTY_MIMEHEADER.html */
/* Moved from NotesMIMEHeader. */
func NotesMIMEHeaderParent(m notesmimeheader.NotesMIMEHeader) (NotesMIMEEntity, error) {
	dispatchPtr, err := m.Com().GetObjectProperty("Parent")
	return New(dispatchPtr), err
}

/* --------------------------------- Properties --------------------------------- */
/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_BOUNDARYEND_PROPERTY_MIMEENTITY.html */
func (m NotesMIMEEntity) BoundaryEnd() (domino.String, error) {
	val, err := m.Com().GetProperty("BoundaryEnd")
	return helpers.CastValue[domino.String](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_BOUNDARYSTART_PROPERTY_MIMEENTITY.html */
func (m NotesMIMEEntity) BoundaryStart() (domino.String, error) {
	val, err := m.Com().GetProperty("BoundaryStart")
	return helpers.CastValue[domino.String](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_CHARSET_PROPERTY_MIMEENTITY.html */
func (m NotesMIMEEntity) Charset() (domino.String, error) {
	val, err := m.Com().GetProperty("Charset")
	return helpers.CastValue[domino.String](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_CONTENTASTEXT_PROPERTY_MIMEENTITY.html */
func (m NotesMIMEEntity) ContentAsText() (domino.String, error) {
	val, err := m.Com().GetProperty("ContentAsText")
	return helpers.CastValue[domino.String](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_CONTENTSUBTYPE_PROPERTY_MIMEENTITY.html */
func (m NotesMIMEEntity) ContentSubType() (domino.String, error) {
	val, err := m.Com().GetProperty("ContentSubType")
	return helpers.CastValue[domino.String](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_CONTENTTYPE_PROPERTY_MIMEENTITY.html */
func (m NotesMIMEEntity) ContentType() (domino.String, error) {
	val, err := m.Com().GetProperty("ContentType")
	return helpers.CastValue[domino.String](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ENCODING_PROPERTY_MIMEENTITY.html */
func (m NotesMIMEEntity) Encoding() (Encoding, error) {
	val, err := m.Com().GetProperty("Encoding")
	return helpers.CastValue[Encoding](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_HEADEROBJECTS_PROPERTY_MIMEENTITY.html */
func (m NotesMIMEEntity) HeaderObjects() ([]notesmimeheader.NotesMIMEHeader, error) {
	return com.GetObjectArrayProperty(m.Com(), notesmimeheader.New, "HeaderObjects")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_HEADERS_PROPERTY_MIMEENTITY.html */
func (m NotesMIMEEntity) Headers() (domino.String, error) {
	val, err := m.Com().GetProperty("Headers")
	return helpers.CastValue[domino.String](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_PREAMBLE_PROPERTY_MIMEENTITY.html */
func (m NotesMIMEEntity) Preamble() (domino.String, error) {
	val, err := m.Com().GetProperty("Preamble")
	return helpers.CastValue[domino.String](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_PREAMBLE_PROPERTY_MIMEENTITY.html */
func (m NotesMIMEEntity) SetPreamble(v domino.String) error {
	return m.Com().PutProperty("Preamble", v)
}

/* --------------------------------- Methods ------------------------------------ */
/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_CREATECHILDENTITY_METHOD_MIMEENTITY.html */
type createChildEntityParams struct {
	nextSibling *NotesMIMEEntity
}

type createChildEntityParam func(*createChildEntityParams)

func WithCreateChildEntityNextSibling(nextSibling NotesMIMEEntity) createChildEntityParam {
	return func(c *createChildEntityParams) {
		c.nextSibling = &nextSibling
	}
}

func (m NotesMIMEEntity) CreateChildEntity(params ...createChildEntityParam) (NotesMIMEEntity, error) {
	paramsStruct := &createChildEntityParams{}
	paramsOrdered := []interface{}{}

	for _, p := range params {
		p(paramsStruct)
	}

	if paramsStruct.nextSibling != nil {
		paramsOrdered = append(paramsOrdered, *paramsStruct.nextSibling)
	}
	dispatchPtr, err := m.Com().CallObjectMethod("CreateChildEntity", paramsOrdered...)
	return New(dispatchPtr), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_CREATEHEADER_METHOD_MIMEENTITY.html */
func (m NotesMIMEEntity) CreateHeader(headerName domino.String) (notesmimeheader.NotesMIMEHeader, error) {
	dispatchPtr, err := m.Com().CallObjectMethod("CreateHeader", headerName)
	return notesmimeheader.New(dispatchPtr), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_CREATEPARENTENTITY_METHOD_MIMEENTITY.html */
func (m NotesMIMEEntity) CreateParentEntity() (NotesMIMEEntity, error) {
	dispatchPtr, err := m.Com().CallObjectMethod("CreateParentEntity")
	return New(dispatchPtr), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_DECODECONTENT_METHOD_MIMEENTITY.html */
func (m NotesMIMEEntity) DecodeContent() error {
	_, err := m.Com().CallMethod("DecodeContent")
	return err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ENCODECONTENT_METHOD_MIMEENTITY.html */
func (m NotesMIMEEntity) EncodeContent(encoding Encoding) error {
	_, err := m.Com().CallMethod("EncodeContent", encoding)
	return err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GETCONTENTASBYTES_METHOD_MIMEENTITY.html */
type getContentAsBytesParams struct {
	decoded *domino.Boolean
}

type getContentAsBytesParam func(*getContentAsBytesParams)

func WithGetContentAsBytesDecoded(decoded domino.Boolean) getContentAsBytesParam {
	return func(c *getContentAsBytesParams) {
		c.decoded = &decoded
	}
}

func (m NotesMIMEEntity) GetContentAsBytes(stream notesstream.NotesStream, params ...getContentAsBytesParam) error {
	paramsStruct := &getContentAsBytesParams{}
	paramsOrdered := []interface{}{stream.Com().Dispatch()}

	for _, p := range params {
		p(paramsStruct)
	}

	if paramsStruct.decoded != nil {
		paramsOrdered = append(paramsOrdered, *paramsStruct.decoded)
	}
	_, err := m.Com().CallMethod("GetContentAsBytes", paramsOrdered...)
	return err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GETCONTENTASTEXT_METHOD_MIMEENTITY.html */
type getContentAsTextParams struct {
	decoded *domino.Boolean
}

type getContentAsTextParam func(*getContentAsTextParams)

func WithGetContentAsTextDecoded(decoded domino.Boolean) getContentAsTextParam {
	return func(c *getContentAsTextParams) {
		c.decoded = &decoded
	}
}

func (m NotesMIMEEntity) GetContentAsText(stream notesstream.NotesStream, params ...getContentAsTextParam) error {
	paramsStruct := &getContentAsTextParams{}
	paramsOrdered := []interface{}{stream.Com().Dispatch()}

	for _, p := range params {
		p(paramsStruct)
	}

	if paramsStruct.decoded != nil {
		paramsOrdered = append(paramsOrdered, *paramsStruct.decoded)
	}
	_, err := m.Com().CallMethod("GetContentAsText", paramsOrdered...)
	return err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GETENTITYASTEXT_METHOD_MIMEENTITY.html */
type getEntityAsTextParams struct {
	headerFilters *[]domino.String
	inclusive     *domino.Boolean
}

type getEntityAsTextParam func(*getEntityAsTextParams)

func WithGetEntityAsTextHeaderFilters(headerFilters []domino.String) getEntityAsTextParam {
	return func(c *getEntityAsTextParams) {
		c.headerFilters = &headerFilters
	}
}

func WithGetEntityAsTextInclusive(inclusive domino.Boolean) getEntityAsTextParam {
	return func(c *getEntityAsTextParams) {
		c.inclusive = &inclusive
	}
}

func (m NotesMIMEEntity) GetEntityAsText(stream notesstream.NotesStream, params ...getEntityAsTextParam) error {
	paramsStruct := &getEntityAsTextParams{}
	paramsOrdered := []interface{}{stream.Com().Dispatch()}

	for _, p := range params {
		p(paramsStruct)
	}

	if paramsStruct.headerFilters != nil {
		paramsOrdered = append(paramsOrdered, *paramsStruct.headerFilters)
		if paramsStruct.inclusive != nil {
			paramsOrdered = append(paramsOrdered, *paramsStruct.inclusive)
		}
	}
	_, err := m.Com().CallMethod("GetEntityAsText", paramsOrdered...)
	return err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GETFIRSTCHILDENTITY_METHOD_MIMEENTITY.html */
func (m NotesMIMEEntity) GetFirstChildEntity() (NotesMIMEEntity, error) {
	dispatchPtr, err := m.Com().CallObjectMethod("GetFirstChildEntity")
	return New(dispatchPtr), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GETNEXTENTITY_METHOD_MIMEENTITY.html */
type getNextEntityParams struct {
	search *domino.Integer
}

type getNextEntityParam func(*getNextEntityParams)

func WithGetNextEntitySearch(search domino.Integer) getNextEntityParam {
	return func(c *getNextEntityParams) {
		c.search = &search
	}
}

func (m NotesMIMEEntity) GetNextEntity(params ...getNextEntityParam) (NotesMIMEEntity, error) {
	paramsStruct := &getNextEntityParams{}
	paramsOrdered := []interface{}{}

	for _, p := range params {
		p(paramsStruct)
	}

	if paramsStruct.search != nil {
		paramsOrdered = append(paramsOrdered, *paramsStruct.search)
	}
	dispatchPtr, err := m.Com().CallObjectMethod("GetNextEntity", paramsOrdered...)
	return New(dispatchPtr), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GETNEXTSIBLING_METHOD_MIMEENTITY.html */
func (m NotesMIMEEntity) GetNextSibling() (NotesMIMEEntity, error) {
	dispatchPtr, err := m.Com().CallObjectMethod("GetNextSibling")
	return New(dispatchPtr), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GETNTHHEADER_METHOD_MIMEENTITY.html */
type getNthHeaderParams struct {
	instance *domino.Integer
}

type getNthHeaderParam func(*getNthHeaderParams)

func WithGetNthHeaderInstance(instance domino.Integer) getNthHeaderParam {
	return func(c *getNthHeaderParams) {
		c.instance = &instance
	}
}

func (m NotesMIMEEntity) GetNthHeader(headerName domino.String, params ...getNthHeaderParam) (notesmimeheader.NotesMIMEHeader, error) {
	paramsStruct := &getNthHeaderParams{}
	paramsOrdered := []interface{}{headerName}

	for _, p := range params {
		p(paramsStruct)
	}

	if paramsStruct.instance != nil {
		paramsOrdered = append(paramsOrdered, *paramsStruct.instance)
	}
	dispatchPtr, err := m.Com().CallObjectMethod("GetNthHeader", paramsOrdered...)
	return notesmimeheader.New(dispatchPtr), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GETPARENTENTITY_METHOD_MIMEENTITY.html */
func (m NotesMIMEEntity) GetParentEntity() (NotesMIMEEntity, error) {
	dispatchPtr, err := m.Com().CallObjectMethod("GetParentEntity")
	return New(dispatchPtr), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GETPREVENTITY_METHOD_MIMEENTITY.html */
type getPrevEntityParams struct {
	search *domino.Integer
}

type getPrevEntityParam func(*getPrevEntityParams)

func WithGetPrevEntitySearch(search domino.Integer) getPrevEntityParam {
	return func(c *getPrevEntityParams) {
		c.search = &search
	}
}

func (m NotesMIMEEntity) GetPrevEntity(params ...getPrevEntityParam) (NotesMIMEEntity, error) {
	paramsStruct := &getPrevEntityParams{}
	paramsOrdered := []interface{}{}

	for _, p := range params {
		p(paramsStruct)
	}

	if paramsStruct.search != nil {
		paramsOrdered = append(paramsOrdered, *paramsStruct.search)
	}
	dispatchPtr, err := m.Com().CallObjectMethod("GetPrevEntity", paramsOrdered...)
	return New(dispatchPtr), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GETPREVSIBLING_METHOD_MIMEENTITY.html */
func (m NotesMIMEEntity) GetPrevSibling() (NotesMIMEEntity, error) {
	dispatchPtr, err := m.Com().CallObjectMethod("GetPrevSibling")
	return New(dispatchPtr), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GETSOMEHEADERS_METHOD_MIMEENTITY.html */
type getSomeHeadersParams struct {
	headerFilters *[]domino.String
	inclusive     *domino.Boolean
}

type getSomeHeadersParam func(*getSomeHeadersParams)

func WithGetSomeHeadersHeaderFilters(headerFilters []domino.String) getSomeHeadersParam {
	return func(c *getSomeHeadersParams) {
		c.headerFilters = &headerFilters
	}
}

func WithGetSomeHeadersInclusive(inclusive domino.Boolean) getSomeHeadersParam {
	return func(c *getSomeHeadersParams) {
		c.inclusive = &inclusive
	}
}

func (m NotesMIMEEntity) GetSomeHeaders(params ...getSomeHeadersParam) (domino.String, error) {
	paramsStruct := &getSomeHeadersParams{}
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
	val, err := m.Com().CallMethod("GetSomeHeaders", paramsOrdered...)
	return helpers.CastValue[domino.String](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_REMOVE_METHOD_MIMEENTITY.html */
func (m NotesMIMEEntity) Remove() error {
	_, err := m.Com().CallMethod("Remove")
	return err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SETCONTENTFROMBYTES_METHOD_MIMEENTITY.html */
func (m NotesMIMEEntity) SetContentFromBytes(stream notesstream.NotesStream, contentType domino.String, encoding Encoding) error {
	_, err := m.Com().CallMethod("SetContentFromBytes", stream.Com().Dispatch(), contentType, encoding)
	return err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SETCONTENTFROMTEXT_METHOD_MIMEENTITY.html */
func (m NotesMIMEEntity) SetContentFromText(stream notesstream.NotesStream, contentType domino.String, encoding Encoding) error {
	_, err := m.Com().CallMethod("SetContentFromText", stream.Com().Dispatch(), contentType, encoding)
	return err
}
