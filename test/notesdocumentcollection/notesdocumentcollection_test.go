/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_NOTESDOCUMENTCOLLECTION_CLASS.html */
package notesdocumentcollection_test

import (
	"testing"

	domigo "github.com/monstermichl/domigo/domino"
	testhelpers "github.com/monstermichl/domigo/test/helpers"

	"github.com/stretchr/testify/require"
)

const TEST_FOLDER_NAME = "test-folder"

var database domigo.NotesDatabase
var documentcollection domigo.NotesDocumentCollection

/* https://pkg.go.dev/testing#hdr-Main */
func TestMain(m *testing.M) {
	session, _ := domigo.Initialize()
	database, _ = testhelpers.CreateTestDatabase(session)
	documentcollection, _ = database.CreateDocumentCollection()

	defer database.Release()
	defer database.Remove()
	defer session.Release()

	m.Run()
}

/* --------------------------------- Properties --------------------------------- */
/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_COUNT_PROPERTY.html */
func TestCount(t *testing.T) {
	_, err := documentcollection.Count()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISSORTED_PROPERTY_COLLECTION.html */
func TestIsSorted(t *testing.T) {
	_, err := documentcollection.IsSorted()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_QUERY_PROPERTY_COLLECTION.html */
func TestQuery(t *testing.T) {
	_, err := documentcollection.Query()
	require.Nil(t, err)
}

/* TODO: Access type for UntilTime could not be evaluated, check yourself if getter/setter is needed. */
/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_UNTILTIME_PROPERTY_COLLECTION.html */
func TestUntilTime(t *testing.T) {
	_, err := documentcollection.UntilTime()
	require.Nil(t, err)
}

/* --------------------------------- Methods ------------------------------------ */
/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ADDDOCUMENT_METHOD_7116_ABOUT.html */
func TestAddDocument(t *testing.T) {
	doc, err := database.CreateDocument()
	require.Nil(t, err)

	err = documentcollection.AddDocument(doc)
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_CLONE_METHOD_COLLECTION.html */
func TestClone(t *testing.T) {
	err := documentcollection.Clone()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_CONTAINS_METHOD_COLLECTION.html */
func TestContainsNoteID(t *testing.T) {
	doc, err := database.CreateDocument()
	defer doc.Release()
	require.Nil(t, err)

	id, err := doc.NoteID()
	require.Nil(t, err)

	_, err = documentcollection.Contains(id)
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_CONTAINS_METHOD_COLLECTION.html */
func TestContainsNotesDocument(t *testing.T) {
	doc, err := database.CreateDocument()
	defer doc.Release()
	require.Nil(t, err)

	_, err = documentcollection.Contains(doc)
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_CONTAINS_METHOD_COLLECTION.html */
func TestContainsNotesDocumentCollection(t *testing.T) {
	col, err := database.CreateDocument()
	defer col.Release()
	require.Nil(t, err)

	_, err = documentcollection.Contains(col)
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_DELETEDOCUMENT_METHOD_7984_ABOUT.html */
func TestDeleteDocument(t *testing.T) {
	doc, err := database.CreateDocument()
	defer doc.Release()
	require.Nil(t, err)

	err = documentcollection.AddDocument(doc)
	require.Nil(t, err)

	err = documentcollection.DeleteDocument(doc)
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_FTSEARCH_METHOD_COLLECTION.html */
// func TestFTSearch(t *testing.T) {
// 	err := documentcollection.FTSearch( /* TODO: Pass test values. */ )
// 	require.Nil(t, err)
// }

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GETDOCUMENT_METHOD_DOCCOLL.html */
func TestGetDocument(t *testing.T) {
	doc, err := database.CreateDocument()
	defer doc.Release()
	require.Nil(t, err)

	err = documentcollection.AddDocument(doc)
	require.Nil(t, err)

	_, err = documentcollection.GetDocument(doc)
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GETFIRSTDOCUMENT_METHOD_COLLECTION.html */
func TestGetFirstDocument(t *testing.T) {
	_, err := documentcollection.GetFirstDocument()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GETLASTDOCUMENT_METHOD_COLLECTION.html */
func TestGetLastDocument(t *testing.T) {
	_, err := documentcollection.GetLastDocument()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GETNEXTDOCUMENT_METHOD_COLLECTION.html */
func TestGetNextDocument(t *testing.T) {
	firstDoc, err := database.CreateDocument()
	defer firstDoc.Release()
	require.Nil(t, err)

	err = documentcollection.AddDocument(firstDoc)
	require.Nil(t, err)

	secondDoc, err := database.CreateDocument()
	defer secondDoc.Release()
	require.Nil(t, err)

	err = documentcollection.AddDocument(secondDoc)
	require.Nil(t, err)

	_, err = documentcollection.GetNextDocument(firstDoc)
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GETNTHDOCUMENT_METHOD_COLLECTION.html */
func TestGetNthDocument(t *testing.T) {
	_, err := documentcollection.GetNthDocument(1)
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GETPREVDOCUMENT_METHOD_COLLECTION.html */
func TestGetPrevDocument(t *testing.T) {
	firstDoc, err := database.CreateDocument()
	defer firstDoc.Release()
	require.Nil(t, err)

	err = documentcollection.AddDocument(firstDoc)
	require.Nil(t, err)

	secondDoc, err := database.CreateDocument()
	defer secondDoc.Release()
	require.Nil(t, err)

	err = documentcollection.AddDocument(secondDoc)
	require.Nil(t, err)

	_, err = documentcollection.GetPrevDocument(secondDoc)
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_INTERSECT_METHOD_COLLECTION.html */
func TestIntersectNoteID(t *testing.T) {
	doc, err := database.CreateDocument()
	defer doc.Release()
	require.Nil(t, err)

	id, err := doc.NoteID()
	require.Nil(t, err)

	err = documentcollection.Intersect(id)
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_INTERSECT_METHOD_COLLECTION.html */
func TestIntersectNotesDocument(t *testing.T) {
	doc, err := database.CreateDocument()
	require.Nil(t, err)

	err = documentcollection.Intersect(doc)
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_INTERSECT_METHOD_COLLECTION.html */
func TestIntersectNotesDocumentCollection(t *testing.T) {
	col, err := database.CreateDocumentCollection()
	require.Nil(t, err)

	err = documentcollection.Intersect(col)
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_MARKALLREAD_DOCCOLLECTION.html */
func TestMarkAllRead(t *testing.T) {
	err := documentcollection.MarkAllRead()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_MARKALLUNREAD_DOCCOLLECTION.html */
func TestMarkAllUnread(t *testing.T) {
	err := documentcollection.MarkAllUnread()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_MERGE_METHOD_COLLECTION.html */
func TestMergeNoteID(t *testing.T) {
	doc, err := database.CreateDocument()
	require.Nil(t, err)

	id, err := doc.NoteID()
	require.Nil(t, err)

	err = documentcollection.Merge(id)
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_MERGE_METHOD_COLLECTION.html */
func TestMergeNotesDocument(t *testing.T) {
	doc, err := database.CreateDocument()
	require.Nil(t, err)

	err = documentcollection.Merge(doc)
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_MERGE_METHOD_COLLECTION.html */
func TestMergeNotesDocumentCollection(t *testing.T) {
	col, err := database.CreateDocumentCollection()
	require.Nil(t, err)

	err = documentcollection.Merge(col)
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_PUTALLINFOLDER_METHOD.html */
func TestPutAllInFolder(t *testing.T) {
	err := database.EnableFolder(TEST_FOLDER_NAME)
	require.Nil(t, err)

	err = documentcollection.PutAllInFolder(TEST_FOLDER_NAME)
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_REMOVEALL_METHOD.html */
func TestRemoveAll(t *testing.T) {
	err := documentcollection.RemoveAll(false)
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_REMOVEALLFROMFOLDER_METHOD.html */
func TestRemoveAllFromFolder(t *testing.T) {
	err := database.EnableFolder(TEST_FOLDER_NAME)
	require.Nil(t, err)

	err = documentcollection.RemoveAllFromFolder(TEST_FOLDER_NAME)
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_STAMPALL_METHOD.html */
func TestStampAll(t *testing.T) {
	err := documentcollection.StampAll("name", "test")
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_STAMPALLMULTI_METHOD_DOCCOL.html */
func TestStampAllMulti(t *testing.T) {
	doc, err := database.CreateDocument()
	require.Nil(t, err)

	err = documentcollection.StampAllMulti(doc)
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SUBTRACT_METHOD_COLLECTION.html */
func TestSubtractNoteID(t *testing.T) {
	doc, err := database.CreateDocument()
	require.Nil(t, err)

	id, err := doc.NoteID()
	require.Nil(t, err)

	err = documentcollection.Subtract(id)
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SUBTRACT_METHOD_COLLECTION.html */
func TestSubtractNotesDocument(t *testing.T) {
	doc, err := database.CreateDocument()
	require.Nil(t, err)

	err = documentcollection.Subtract(doc)
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SUBTRACT_METHOD_COLLECTION.html */
func TestSubtractNotesDocumentCollection(t *testing.T) {
	col, err := database.CreateDocumentCollection()
	require.Nil(t, err)

	err = documentcollection.Subtract(col)
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_UPDATEALL_METHOD.html */
func TestUpdateAll(t *testing.T) {
	err := documentcollection.UpdateAll()
	require.Nil(t, err)
}
