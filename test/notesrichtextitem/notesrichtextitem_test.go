/* https://help.hcl-software.com/dom_designer/10.0.1/basic/H_NOTESRICHTEXTITEM_CLASS.html */
package notesrichtextitem_test

import (
	"domigo/domino/notesdatabase"
	"domigo/domino/notesdocument"
	"domigo/domino/notesembeddedobject"
	"domigo/domino/notesrichtextitem"
	"domigo/domino/notessession"
	testhelpers "domigo/test/helpers"
	"testing"

	"github.com/stretchr/testify/require"
)

const EMBEDDED_OBJECT_NAME = "EmbeddedTestObject"

var session notessession.NotesSession
var database notesdatabase.NotesDatabase
var doc notesdocument.NotesDocument
var richtextitem notesrichtextitem.NotesRichTextItem

/* https://pkg.go.dev/testing#hdr-Main */
func TestMain(m *testing.M) {
	session, _ = notessession.Initialize()
	database, _ = testhelpers.CreateTestDatabase(session)
	doc, _ = database.CreateDocument()
	richtextitem, _ = doc.CreateRichTextItem("TestItem")

	defer richtextitem.Release()
	defer doc.Release()
	defer database.Release()
	defer database.Remove()
	defer session.Release()

	m.Run()
}

/* --------------------------------- Properties --------------------------------- */
/* https://help.hcl-software.com/dom_designer/10.0.1/basic/H_EMBEDDEDOBJECTS_PROPERTY_RTITEM.html */
func TestEmbeddedObjects(t *testing.T) {
	_, err := richtextitem.EmbeddedObjects()
	require.Nil(t, err)
}

/* --------------------------------- Methods ------------------------------------ */
/* https://help.hcl-software.com/dom_designer/10.0.1/basic/H_ADDNEWLINE_METHOD.html */
func TestAddNewLine(t *testing.T) {
	err := richtextitem.AddNewLine()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/10.0.1/basic/H_ADDPAGEBREAK_METHOD_290_ABOUT.html */
func TestAddPageBreak(t *testing.T) {
	err := richtextitem.AddPageBreak()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/10.0.1/basic/H_ADDTAB_METHOD.html */
func TestAddTab(t *testing.T) {
	err := richtextitem.AddTab()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/10.0.1/basic/H_APPENDPARAGRAPHSTYLE_METHOD_1636_ABOUT.html */
func TestAppendParagraphStyle(t *testing.T) {
	style, err := session.CreateRichTextParagraphStyle()

	require.Nil(t, err, "Could not create style")
	defer style.Release()

	err = richtextitem.AppendParagraphStyle(style)
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/10.0.1/basic/H_APPENDRTITEM_METHOD.html */
func TestAppendRTItem(t *testing.T) {
	rti, err := doc.CreateRichTextItem("NewRichTextItem")
	require.Nil(t, err, "Could not create RichText item")

	err = richtextitem.AppendRTItem(rti)
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/10.0.1/basic/H_APPENDSTYLE_METHOD.html */
func TestAppendStyle(t *testing.T) {
	style, err := session.CreateRichTextStyle()

	require.Nil(t, err, "Could not create style")
	defer style.Release()

	err = richtextitem.AppendStyle(style)
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/10.0.1/basic/H_APPENDTABLE_METHOD_RTITEM.html */
func TestAppendTable(t *testing.T) {
	err := richtextitem.AppendTable(1, 1)
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/10.0.1/basic/H_APPENDTEXT_METHOD.html */
func TestAppendText(t *testing.T) {
	err := richtextitem.AppendText("test text")
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/10.0.1/basic/H_BEGININSERT_METHOD_RTITEM.html */
func TestBeginInsert(t *testing.T) {
	o, err := richtextitem.EmbedObject(notesembeddedobject.EMBED_ATTACHMENT, "", "notesrichtextitem_test.go")
	require.Nil(t, err, "Could not create embedded object")

	err = richtextitem.BeginInsert(o)
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/10.0.1/basic/H_BEGINSECTION_METHOD_RTITEM.html */
func TestBeginSection(t *testing.T) {
	err := richtextitem.BeginSection("Test section")
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/10.0.1/basic/H_COMPACT_METHOD_RTITEM.html */
func TestCompact(t *testing.T) {
	err := richtextitem.Compact()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/10.0.1/basic/H_CONVERTOHTML_METHOD_NOTESRICHTEXTITEM.html */
func TestConverttohtml(t *testing.T) {
	err := richtextitem.Converttohtml()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/10.0.1/basic/H_CREATENAVIGATOR_METHOD_RTITEM.html */
func TestCreateNavigator(t *testing.T) {
	_, err := richtextitem.CreateNavigator()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/10.0.1/basic/H_CREATERANGE_METHOD_RTITEM.html */
func TestCreateRange(t *testing.T) {
	_, err := richtextitem.CreateRange()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/10.0.1/basic/H_EMBEDOBJECT_METHOD.html */
func TestEmbedObject(t *testing.T) {
	_, err := richtextitem.EmbedObject(notesembeddedobject.EMBED_OBJECT, "", "notesrichtextitem_test.go", notesrichtextitem.WithEmbedObjectName(EMBEDDED_OBJECT_NAME))
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/10.0.1/basic/H_ENDINSERT_METHOD_RTITEM.html */
func TestEndInsert(t *testing.T) {
	err := richtextitem.EndInsert()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/10.0.1/basic/H_ENDSECTION_METHOD_RTITEM.html */
func TestEndSection(t *testing.T) {
	err := richtextitem.EndSection()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/10.0.1/basic/H_GETEMBEDDEDOBJECT_METHOD.html */
func TestGetEmbeddedObject(t *testing.T) {
	_, err := richtextitem.GetEmbeddedObject(EMBEDDED_OBJECT_NAME)
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/10.0.1/basic/H_GETFORMATTEDTEXT_METHOD.html */
func TestGetFormattedText(t *testing.T) {
	_, err := richtextitem.GetFormattedText(true, 50)
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/10.0.1/basic/H_GETHTMLREFERENCES_METHOD_NOTESRICHTEXTITEM.html */
func TestGetHTMLReferences(t *testing.T) {
	_, err := richtextitem.GetHTMLReferences()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/10.0.1/basic/H_GETNOTESFONT_METHOD_RTITEM.html */
func TestGetNotesFont(t *testing.T) {
	_, err := richtextitem.GetNotesFont("Arial")
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/10.0.1/basic/H_GETUNFORMATTEDTEXT_METHOD.html */
func TestGetUnformattedText(t *testing.T) {
	_, err := richtextitem.GetUnformattedText()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/10.0.1/basic/H_UPDATE_METHOD_RTITEM.html */
func TestUpdate(t *testing.T) {
	err := richtextitem.Update()
	require.Nil(t, err)
}
