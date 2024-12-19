/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_NOTESSTREAM_CLASS.html */
package domigo

import (
	"github.com/monstermichl/domigo/helpers"

	ole "github.com/go-ole/go-ole"
)

type NotesStream struct {
	NotesStruct
}

func NewNotesStream(dispatchPtr *ole.IDispatch) NotesStream {
	return NotesStream{NewNotesStruct(dispatchPtr)}
}

/* --------------------------------- Properties --------------------------------- */
/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_BYTES_PROPERTY_STREAM.html */
func (s NotesStream) Bytes() (Long, error) {
	val, err := s.Com().GetProperty("Bytes")
	return helpers.CastValue[Long](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_CHARSET_PROPERTY_STREAM.html */
func (s NotesStream) Charset() (String, error) {
	val, err := s.Com().GetProperty("Charset")
	return helpers.CastValue[String](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISEOS_PROPERTY_STREAM.html */
func (s NotesStream) IsEOS() (Boolean, error) {
	val, err := s.Com().GetProperty("IsEOS")
	return helpers.CastValue[Boolean](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISREADONLY_PROPERTY_STREAM.html */
func (s NotesStream) IsReadOnly() (Boolean, error) {
	val, err := s.Com().GetProperty("IsReadOnly")
	return helpers.CastValue[Boolean](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_PARENT_PROPERTY_STREAM.html */
func (s NotesStream) Parent() (NotesSession, error) {
	dispatchPtr, err := s.Com().GetObjectProperty("Parent")
	return NewNotesSession(dispatchPtr), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_POSITION_PROPERTY_STREAM.html */
func (s NotesStream) Position() (Long, error) {
	val, err := s.Com().GetProperty("Position")
	return helpers.CastValue[Long](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_POSITION_PROPERTY_STREAM.html */
func (s NotesStream) SetPosition(v Long) error {
	return s.Com().PutProperty("Position", v)
}

/* --------------------------------- Methods ------------------------------------ */
/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_CLOSE_METHOD_STREAM.html */
func (s NotesStream) Close() error {
	_, err := s.Com().CallMethod("Close")
	return err
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
	val, err := s.Com().CallMethod("Open", paramsOrdered...)
	return helpers.CastValue[Boolean](val), err
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
	vals, err := s.Com().CallArrayMethod("Read", paramsOrdered...)
	return helpers.CastSlice[Byte](vals), err
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
	val, err := s.Com().CallMethod("ReadText", paramsOrdered...)
	return helpers.CastValue[String](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_TRUNCATE_METHOD_STREAM.html */
func (s NotesStream) Truncate() error {
	_, err := s.Com().CallMethod("Truncate")
	return err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_WRITE_METHOD_STREAM.html */
func (s NotesStream) Write(buffer []Byte) (Byte, error) {
	val, err := s.Com().CallMethod("Write", buffer)
	return helpers.CastValue[Byte](val), err
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

func (s NotesStream) WriteText(text String, params ...notesStreamWriteTextParam) (Byte, error) {
	paramsStruct := &notesStreamWriteTextParams{}
	paramsOrdered := []interface{}{text}

	for _, p := range params {
		p(paramsStruct)
	}

	if paramsStruct.eol != nil {
		paramsOrdered = append(paramsOrdered, *paramsStruct.eol)
	}
	val, err := s.Com().CallMethod("WriteText", paramsOrdered...)
	return helpers.CastValue[Byte](val), err
}
