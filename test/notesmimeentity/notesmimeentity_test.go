/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_NOTESMIMEENTITY_CLASS_OVERVIEW.html */
package notesmimeentity_test

import (
	"testing"

	"github.com/monstermichl/domigo"
	testhelpers "github.com/monstermichl/domigo/test/helpers"

	"github.com/stretchr/testify/require"
)

var stream domigo.NotesStream
var mimeentity domigo.NotesMIMEEntity

/* https://pkg.go.dev/testing#hdr-Main */
func TestMain(m *testing.M) {
	testhelpers.Initialize(func(session domigo.NotesSession, db domigo.NotesDatabase) (string, error) {
		doc, err := db.CreateDocument()
		defer doc.Release()

		if err != nil {
			return "Document could not be created", err
		}
		stream, err = session.CreateStream()
		defer stream.Release()

		if err != nil {
			return "Stream could not be created", err
		}
		mimeentity, err = doc.CreateMIMEEntity()
		defer mimeentity.Release()

		if err != nil {
			return "MIME entity could not be created", err
		}
		m.Run()
		return "", nil
	})
}

/* --------------------------------- Properties --------------------------------- */
/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_BOUNDARYEND_PROPERTY_MIMEENTITY.html */
func TestBoundaryEnd(t *testing.T) {
	_, err := mimeentity.BoundaryEnd()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_BOUNDARYSTART_PROPERTY_MIMEENTITY.html */
func TestBoundaryStart(t *testing.T) {
	_, err := mimeentity.BoundaryStart()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_CHARSET_PROPERTY_MIMEENTITY.html */
func TestCharset(t *testing.T) {
	_, err := mimeentity.Charset()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_CONTENTASTEXT_PROPERTY_MIMEENTITY.html */
func TestContentAsText(t *testing.T) {
	_, err := mimeentity.ContentAsText()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_CONTENTSUBTYPE_PROPERTY_MIMEENTITY.html */
func TestContentSubType(t *testing.T) {
	_, err := mimeentity.ContentSubType()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_CONTENTTYPE_PROPERTY_MIMEENTITY.html */
func TestContentType(t *testing.T) {
	_, err := mimeentity.ContentType()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ENCODING_PROPERTY_MIMEENTITY.html */
func TestEncoding(t *testing.T) {
	_, err := mimeentity.Encoding()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_HEADEROBJECTS_PROPERTY_MIMEENTITY.html */
func TestHeaderObjects(t *testing.T) {
	_, err := mimeentity.HeaderObjects()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_HEADERS_PROPERTY_MIMEENTITY.html */
func TestHeaders(t *testing.T) {
	_, err := mimeentity.Headers()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_PREAMBLE_PROPERTY_MIMEENTITY.html */
func TestPreamble(t *testing.T) {
	_, err := mimeentity.Preamble()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_PREAMBLE_PROPERTY_MIMEENTITY.html */
func TestSetPreamble(t *testing.T) {
	s, err := mimeentity.Preamble()
	require.Nil(t, err)

	err = mimeentity.SetPreamble(s)
	require.Nil(t, err)
}

/* --------------------------------- Methods ------------------------------------ */
/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_CREATECHILDENTITY_METHOD_MIMEENTITY.html */
func TestCreateChildEntity(t *testing.T) {
	_, err := mimeentity.CreateChildEntity( /* TODO: Pass test values. */ )
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_CREATEHEADER_METHOD_MIMEENTITY.html */
func TestCreateHeader(t *testing.T) {
	_, err := mimeentity.CreateHeader("Test Header")
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_CREATEPARENTENTITY_METHOD_MIMEENTITY.html */
func TestCreateParentEntity(t *testing.T) {
	_, err := mimeentity.CreateParentEntity()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_DECODECONTENT_METHOD_MIMEENTITY.html */
func TestDecodeContent(t *testing.T) {
	err := mimeentity.DecodeContent()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ENCODECONTENT_METHOD_MIMEENTITY.html */
func TestEncodeContent(t *testing.T) {
	err := mimeentity.EncodeContent(domigo.NOTESMIMEENTITY_ENC_NONE)
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GETCONTENTASBYTES_METHOD_MIMEENTITY.html */
func TestGetContentAsBytes(t *testing.T) {
	err := mimeentity.GetContentAsBytes(stream)
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GETCONTENTASTEXT_METHOD_MIMEENTITY.html */
func TestGetContentAsText(t *testing.T) {
	err := mimeentity.GetContentAsText(stream)
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GETENTITYASTEXT_METHOD_MIMEENTITY.html */
func TestGetEntityAsText(t *testing.T) {
	err := mimeentity.GetEntityAsText(stream)
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GETFIRSTCHILDENTITY_METHOD_MIMEENTITY.html */
func TestGetFirstChildEntity(t *testing.T) {
	_, err := mimeentity.GetFirstChildEntity()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GETNEXTENTITY_METHOD_MIMEENTITY.html */
func TestGetNextEntity(t *testing.T) {
	_, err := mimeentity.GetNextEntity()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GETNEXTSIBLING_METHOD_MIMEENTITY.html */
func TestGetNextSibling(t *testing.T) {
	_, err := mimeentity.GetNextSibling()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GETNTHHEADER_METHOD_MIMEENTITY.html */
func TestGetNthHeader(t *testing.T) {
	_, err := mimeentity.GetNthHeader("test")
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GETPARENTENTITY_METHOD_MIMEENTITY.html */
func TestGetParentEntity(t *testing.T) {
	_, err := mimeentity.GetParentEntity()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GETPREVENTITY_METHOD_MIMEENTITY.html */
func TestGetPrevEntity(t *testing.T) {
	_, err := mimeentity.GetPrevEntity()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GETPREVSIBLING_METHOD_MIMEENTITY.html */
func TestGetPrevSibling(t *testing.T) {
	_, err := mimeentity.GetPrevSibling()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GETSOMEHEADERS_METHOD_MIMEENTITY.html */
func TestGetSomeHeaders(t *testing.T) {
	_, err := mimeentity.GetSomeHeaders()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SETCONTENTFROMBYTES_METHOD_MIMEENTITY.html */
func TestSetContentFromBytes(t *testing.T) {
	err := mimeentity.SetContentFromBytes(stream, "text/plain", domigo.NOTESMIMEENTITY_ENC_BASE64)
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SETCONTENTFROMTEXT_METHOD_MIMEENTITY.html */
func TestSetContentFromText(t *testing.T) {
	err := mimeentity.SetContentFromText(stream, "text/plain", domigo.NOTESMIMEENTITY_ENC_BASE64)
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_REMOVE_METHOD_MIMEENTITY.html */
func TestRemove(t *testing.T) {
	err := mimeentity.Remove()
	require.Nil(t, err)
}
