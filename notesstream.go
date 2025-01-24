/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_NOTESSTREAM_CLASS.html */
package domigo

import (
	ole "github.com/go-ole/go-ole"
)

type NotesStream struct {
	NotesStruct
}

func newNotesStream(dispatchPtr *ole.IDispatch) NotesStream {
	return NotesStream{newNotesStruct(dispatchPtr)}
}

/* --------------------------------- Properties --------------------------------- */
/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_BYTES_PROPERTY_STREAM.html */
func (s NotesStream) Bytes() (Long, error) {
	return getComProperty[Long](s, "Bytes")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_CHARSET_PROPERTY_STREAM.html */
func (s NotesStream) Charset() (String, error) {
	return getComProperty[String](s, "Charset")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISEOS_PROPERTY_STREAM.html */
func (s NotesStream) IsEOS() (Boolean, error) {
	return getComProperty[Boolean](s, "IsEOS")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISREADONLY_PROPERTY_STREAM.html */
func (s NotesStream) IsReadOnly() (Boolean, error) {
	return getComProperty[Boolean](s, "IsReadOnly")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_PARENT_PROPERTY_STREAM.html */
func (s NotesStream) Parent() (NotesSession, error) {
	return getComObjectProperty(s, newNotesSession, "Parent")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_POSITION_PROPERTY_STREAM.html */
func (s NotesStream) Position() (Long, error) {
	return getComProperty[Long](s, "Position")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_POSITION_PROPERTY_STREAM.html */
func (s NotesStream) SetPosition(v Long) error {
	return putComProperty(s, "Position", v)
}

/* --------------------------------- Methods ------------------------------------ */
/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_CLOSE_METHOD_STREAM.html */
func (s NotesStream) Close() error {
	return callComVoidMethod(s, "Close")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_OPEN_METHOD_STREAM.html */
type notesStreamOpenParams struct {
	charset *String
}

type notesStreamOpenParam func(*notesStreamOpenParams)

func WithNotesStreamOpenCharset(charset String) notesStreamOpenParam {
	return func(c *notesStreamOpenParams) {
		c.charset = &charset
	}
}

func (s NotesStream) Open(pathname String, params ...notesStreamOpenParam) (Boolean, error) {
	paramsStruct := &notesStreamOpenParams{}
	paramsOrdered := []interface{}{pathname}

	for _, p := range params {
		p(paramsStruct)
	}

	if paramsStruct.charset != nil {
		paramsOrdered = append(paramsOrdered, *paramsStruct.charset)
	}
	return callComMethod[Boolean](s, "Open", paramsOrdered...)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_READ_METHOD_STREAM.html */
type notesStreamReadParams struct {
	length *Long
}

type notesStreamReadParam func(*notesStreamReadParams)

func WithNotesStreamReadLength(length Long) notesStreamReadParam {
	return func(c *notesStreamReadParams) {
		c.length = &length
	}
}

func (s NotesStream) Read(params ...notesStreamReadParam) ([]Byte, error) {
	paramsStruct := &notesStreamReadParams{}
	paramsOrdered := []interface{}{}

	for _, p := range params {
		p(paramsStruct)
	}

	if paramsStruct.length != nil {
		paramsOrdered = append(paramsOrdered, *paramsStruct.length)
	}
	return callComArrayMethod[Byte](s, "Read", paramsOrdered...)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_READTEXT_METHOD_STREAM.html */
type notesStreamReadTextParams struct {
	oneLine *Long
	eol     *Long
}

type notesStreamReadTextParam func(*notesStreamReadTextParams)

func WithNotesStreamReadTextOneLine(oneLine Long) notesStreamReadTextParam {
	return func(c *notesStreamReadTextParams) {
		c.oneLine = &oneLine
	}
}

func WithNotesStreamReadTextEol(eol Long) notesStreamReadTextParam {
	return func(c *notesStreamReadTextParams) {
		c.eol = &eol
	}
}

func (s NotesStream) ReadText(params ...notesStreamReadTextParam) (String, error) {
	paramsStruct := &notesStreamReadTextParams{}
	paramsOrdered := []interface{}{}

	for _, p := range params {
		p(paramsStruct)
	}

	if paramsStruct.oneLine != nil {
		paramsOrdered = append(paramsOrdered, *paramsStruct.oneLine)
		if paramsStruct.eol != nil {
			paramsOrdered = append(paramsOrdered, *paramsStruct.eol)
		}
	}
	return callComMethod[String](s, "ReadText", paramsOrdered...)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_TRUNCATE_METHOD_STREAM.html */
func (s NotesStream) Truncate() error {
	return callComVoidMethod(s, "Truncate")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_WRITE_METHOD_STREAM.html */
func (s NotesStream) Write(buffer []Byte) (Long, error) {
	return callComMethod[Long](s, "Write", buffer)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_WRITETEXT_METHOD_STREAM.html */
type notesStreamWriteTextParams struct {
	eol *Long
}

type notesStreamWriteTextParam func(*notesStreamWriteTextParams)

func WithNotesStreamWriteTextEol(eol Long) notesStreamWriteTextParam {
	return func(c *notesStreamWriteTextParams) {
		c.eol = &eol
	}
}

func (s NotesStream) WriteText(text String, params ...notesStreamWriteTextParam) (Long, error) {
	paramsStruct := &notesStreamWriteTextParams{}
	paramsOrdered := []interface{}{text}

	for _, p := range params {
		p(paramsStruct)
	}

	if paramsStruct.eol != nil {
		paramsOrdered = append(paramsOrdered, *paramsStruct.eol)
	}
	return callComMethod[Long](s, "WriteText", paramsOrdered...)
}
