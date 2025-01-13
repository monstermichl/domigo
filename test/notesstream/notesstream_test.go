/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_NOTESSTREAM_CLASS.html */
package notesstream_test

import (
    "github.com/monstermichl/domigo"
    testhelpers "github.com/monstermichl/domigo/test/helpers"
    "testing"

    "github.com/stretchr/testify/require"
)

var stream domigo.NotesStream

/* https://pkg.go.dev/testing#hdr-Main */
func TestMain(m *testing.M) {
    testhelpers.Initialize(func(session domigo.NotesSession, db domigo.NotesDatabase) (string, error) {
        var err error
        stream, err = session.CreateStream()
        defer stream.Release()

        if err != nil {
            return "NotesStream could not be created", err
        }
        m.Run()
        return "", nil
    })
}

/* --------------------------------- Properties --------------------------------- */
/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_BYTES_PROPERTY_STREAM.html */
func TestBytes(t *testing.T) {
    _, err := stream.Bytes()
    require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_CHARSET_PROPERTY_STREAM.html */
func TestCharset(t *testing.T) {
    _, err := stream.Charset()
    require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISEOS_PROPERTY_STREAM.html */
func TestIsEOS(t *testing.T) {
    _, err := stream.IsEOS()
    require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISREADONLY_PROPERTY_STREAM.html */
func TestIsReadOnly(t *testing.T) {
    _, err := stream.IsReadOnly()
    require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_PARENT_PROPERTY_STREAM.html */
func TestParent(t *testing.T) {
    _, err := stream.Parent()
    require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_POSITION_PROPERTY_STREAM.html */
func TestPosition(t *testing.T) {
    _, err := stream.Position()
    require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_POSITION_PROPERTY_STREAM.html */
func TestSetPosition(t *testing.T) {
    s, err := stream.Position()
    require.Nil(t, err)

    err = stream.SetPosition(s)
    require.Nil(t, err)
}
/* --------------------------------- Methods ------------------------------------ */
/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_CLOSE_METHOD_STREAM.html */
func TestClose(t *testing.T) {
    err := stream.Close()
    require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_OPEN_METHOD_STREAM.html */
func TestOpen(t *testing.T) {
    _, err := stream.Open("testStream")
    require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_READ_METHOD_STREAM.html */
func TestRead(t *testing.T) {
    _, err := stream.Read()
    require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_READTEXT_METHOD_STREAM.html */
func TestReadText(t *testing.T) {
    _, err := stream.ReadText()
    require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_TRUNCATE_METHOD_STREAM.html */
func TestTruncate(t *testing.T) {
    err := stream.Truncate()
    require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_WRITE_METHOD_STREAM.html */
func TestWrite(t *testing.T) {
    _, err := stream.Write([]domigo.Byte{0, 0, 0, 0})
    require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_WRITETEXT_METHOD_STREAM.html */
func TestWriteText(t *testing.T) {
    _, err := stream.WriteText("testText")
    require.Nil(t, err)
}
