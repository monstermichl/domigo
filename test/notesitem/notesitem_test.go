/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_NOTESITEM_CLASS.html */
package notesitem_test

import (
	"domigo/domino"
	"domigo/domino/notesdatabase"
	"domigo/domino/notesitem"
	"domigo/domino/notessession"
	testhelpers "domigo/test/helpers"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

const TEST_ITEM = "TestItem"
const TEST_VALUE = "test value"
const TEST_DATA_TYPE = "TestDataType"

var testDataTypeValue = []domino.Byte{0, 1, 2, 3}
var session notessession.NotesSession
var database notesdatabase.NotesDatabase
var item notesitem.NotesItem

/* https://pkg.go.dev/testing#hdr-Main */
func TestMain(m *testing.M) {
	session, _ = notessession.Initialize()
	database, _ = testhelpers.CreateTestDatabase(session)
	doc, _ := database.CreateDocument()
	item, _ = doc.ReplaceItemValue(TEST_ITEM, TEST_VALUE)

	defer item.Release()
	defer doc.Release()
	defer database.Release()
	defer database.Remove()
	defer session.Release()

	m.Run()
}

/* --------------------------------- Properties --------------------------------- */
/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_DATETIMEVALUE_PROPERTY.html */
func TestDateTimeValue(t *testing.T) {
	_, err := item.DateTimeValue()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_DATETIMEVALUE_PROPERTY.html */
func TestSetDateTimeValue(t *testing.T) {
	d, err := session.CreateDateTime("Today")
	require.Nil(t, err)

	err = item.SetDateTimeValue(d)
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISAUTHORS_PROPERTY.html */
func TestIsAuthors(t *testing.T) {
	_, err := item.IsAuthors()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISAUTHORS_PROPERTY.html */
func TestSetIsAuthors(t *testing.T) {
	s, err := item.IsAuthors()
	require.Nil(t, err)

	err = item.SetIsAuthors(s)
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISENCRYPTED_PROPERTY.html */
func TestIsEncrypted(t *testing.T) {
	_, err := item.IsEncrypted()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISENCRYPTED_PROPERTY.html */
func TestSetIsEncrypted(t *testing.T) {
	s, err := item.IsEncrypted()
	require.Nil(t, err)

	err = item.SetIsEncrypted(s)
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISNAMES_PROPERTY.html */
func TestIsNames(t *testing.T) {
	_, err := item.IsNames()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISNAMES_PROPERTY.html */
func TestSetIsNames(t *testing.T) {
	s, err := item.IsNames()
	require.Nil(t, err)

	err = item.SetIsNames(s)
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISPROTECTED_PROPERTY.html */
func TestIsProtected(t *testing.T) {
	_, err := item.IsProtected()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISPROTECTED_PROPERTY.html */
func TestSetIsProtected(t *testing.T) {
	s, err := item.IsProtected()
	require.Nil(t, err)

	err = item.SetIsProtected(s)
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISREADERS_PROPERTY.html */
func TestIsReaders(t *testing.T) {
	_, err := item.IsReaders()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISREADERS_PROPERTY.html */
func TestSetIsReaders(t *testing.T) {
	s, err := item.IsReaders()
	require.Nil(t, err)

	err = item.SetIsReaders(s)
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISSIGNED_PROPERTY_ITEM.html */
func TestIsSigned(t *testing.T) {
	_, err := item.IsSigned()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISSIGNED_PROPERTY_ITEM.html */
func TestSetIsSigned(t *testing.T) {
	s, err := item.IsSigned()
	require.Nil(t, err)

	err = item.SetIsSigned(s)
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISSUMMARY_PROPERTY.html */
func TestIsSummary(t *testing.T) {
	_, err := item.IsSummary()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISSUMMARY_PROPERTY.html */
func TestSetIsSummary(t *testing.T) {
	s, err := item.IsSummary()
	require.Nil(t, err)

	err = item.SetIsSummary(s)
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_LASTMODIFIED_PROPERTY_ITEM.html */
func TestLastModified(t *testing.T) {
	_, err := item.LastModified()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_NAME_PROPERTY_ITEM.html */
func TestName(t *testing.T) {
	_, err := item.Name()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SAVETODISK_PROPERTY.html */
func TestSaveToDisk(t *testing.T) {
	_, err := item.SaveToDisk()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SAVETODISK_PROPERTY.html */
func TestSetSaveToDisk(t *testing.T) {
	s, err := item.SaveToDisk()
	require.Nil(t, err)

	err = item.SetSaveToDisk(s)
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_TEXT_PROPERTY.html */
func TestText(t *testing.T) {
	_, err := item.Text()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_TYPE_PROPERTY_ITEM.html */
func TestType(t *testing.T) {
	tp, err := item.Type()

	require.Nil(t, err)
	require.Equal(t, notesitem.DATATYPE_TEXT, tp)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_VALUELENGTH_PROPERTY.html */
func TestValueLength(t *testing.T) {
	_, err := item.ValueLength()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_VALUES_PROPERTY.html */
func TestValues(t *testing.T) {
	_, err := item.Values()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_VALUES_PROPERTY.html */
func TestSetValues(t *testing.T) {
	err := item.SetValues(TEST_VALUE)
	require.Nil(t, err)
}

/* --------------------------------- Methods ------------------------------------ */
/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ABSTRACT_METHOD.html */
func TestAbstract(t *testing.T) {
	a, err := item.Abstract(6, true, false)

	require.Nil(t, err)
	require.Equal(t, "tst vl", a)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_APPENDTOTEXTLIST_METHOD.html */
func TestAppendToTextList(t *testing.T) {
	err := item.AppendToTextList([]domino.String{strings.Join([]string{TEST_VALUE, "2"}, "")})
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_CONTAINS_METHOD.html */
func TestContains(t *testing.T) {
	c, err := item.Contains(TEST_VALUE)

	require.Nil(t, err)
	require.True(t, c)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SETVALUECUSTOMDATABYTES_METHOD_ITEM.html */
func TestSetValueCustomDataBytes(t *testing.T) {
	err := item.SetValueCustomDataBytes(TEST_DATA_TYPE, testDataTypeValue)
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GETVALUECUSTOMDATABYTES_METHOD_ITEM.html */
func TestGetValueCustomDataBytes(t *testing.T) {
	b, err := item.GetValueCustomDataBytes(TEST_DATA_TYPE)

	require.Nil(t, err)
	require.Equal(t, testDataTypeValue, b)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GETVALUEDATETIMEARRAY_METHOD.html */
func TestGetValueDateTimeArray(t *testing.T) {
	_, err := item.GetValueDateTimeArray()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GETMIMEENTITY_METHOD_NOTESITEM.html */
func TestGetMIMEEntity(t *testing.T) {
	_, err := item.GetMIMEEntity()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_REMOVE_METHOD_ITEM.html */
func TestRemove(t *testing.T) {
	err := item.Remove()
	require.Nil(t, err)
}
