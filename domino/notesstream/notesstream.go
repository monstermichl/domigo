package notesstream

import (
	"github.com/monstermichl/domigo/domino"
	"github.com/monstermichl/domigo/helpers"

	ole "github.com/go-ole/go-ole"
)

type NotesStream struct {
	domino.NotesStruct
}

func New(dispatchPtr *ole.IDispatch) NotesStream {
	return NotesStream{domino.NewNotesStruct(dispatchPtr)}
}

/* --------------------------------- Properties --------------------------------- */
/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_BYTES_PROPERTY_STREAM.html */
func (s NotesStream) Bytes() (domino.Long, error) {
	val, err := s.Com().GetProperty("Bytes")
	return helpers.CastValue[domino.Long](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_CHARSET_PROPERTY_STREAM.html */
func (s NotesStream) Charset() (domino.String, error) {
	val, err := s.Com().GetProperty("Charset")
	return helpers.CastValue[domino.String](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISEOS_PROPERTY_STREAM.html */
func (s NotesStream) IsEOS() (domino.Boolean, error) {
	val, err := s.Com().GetProperty("IsEOS")
	return helpers.CastValue[domino.Boolean](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISREADONLY_PROPERTY_STREAM.html */
func (s NotesStream) IsReadOnly() (domino.Boolean, error) {
	val, err := s.Com().GetProperty("IsReadOnly")
	return helpers.CastValue[domino.Boolean](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_POSITION_PROPERTY_STREAM.html */
func (s NotesStream) Position() (domino.Long, error) {
	val, err := s.Com().GetProperty("Position")
	return helpers.CastValue[domino.Long](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_POSITION_PROPERTY_STREAM.html */
func (s NotesStream) SetPosition(v domino.Long) error {
	return s.Com().PutProperty("Position", v)
}

/* --------------------------------- Methods ------------------------------------ */
/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_CLOSE_METHOD_STREAM.html */
func (s NotesStream) Close() error {
	_, err := s.Com().CallMethod("Close")
	return err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_OPEN_METHOD_STREAM.html */
func (s NotesStream) Open(pathname domino.String, charset domino.String) (domino.Boolean, error) {
	val, err := s.Com().CallMethod("Open", pathname, charset)
	return helpers.CastValue[domino.Boolean](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_READ_METHOD_STREAM.html */
func (s NotesStream) Read(length domino.Long) ([]domino.Byte, error) {
	vals, err := s.Com().CallArrayMethod("Read", length)
	return helpers.CastSlice[domino.Byte](vals), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_READTEXT_METHOD_STREAM.html */
func (s NotesStream) ReadText(oneLine domino.Long, eol domino.Long) (domino.String, error) {
	val, err := s.Com().CallMethod("ReadText", oneLine, eol)
	return helpers.CastValue[domino.String](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_TRUNCATE_METHOD_STREAM.html */
func (s NotesStream) Truncate() error {
	_, err := s.Com().CallMethod("Truncate")
	return err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_WRITE_METHOD_STREAM.html */
func (s NotesStream) Write(buffer []domino.Byte) (domino.Byte, error) {
	val, err := s.Com().CallMethod("Write", buffer)
	return helpers.CastValue[domino.Byte](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_WRITETEXT_METHOD_STREAM.html */
func (s NotesStream) WriteText(text domino.String, eol domino.Long) (domino.Byte, error) {
	val, err := s.Com().CallMethod("WriteText", text, eol)
	return helpers.CastValue[domino.Byte](val), err
}
