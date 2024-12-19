/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_NOTESQUERYRESULTSPROCESSOR_CLASS.html */
package notesqueryresultsprocessor_test

import (
	"fmt"
	"testing"

	"github.com/monstermichl/domigo"
	testhelpers "github.com/monstermichl/domigo/test/helpers"

	"github.com/stretchr/testify/require"
)

var database domigo.NotesDatabase
var queryresultsprocessor domigo.NotesQueryResultsProcessor

/* https://pkg.go.dev/testing#hdr-Main */
func TestMain(m *testing.M) {
	var info string

	session, err := domigo.Initialize()
	defer session.Release()

	defer func() {
		fmt.Println(err)
		fmt.Println(info)
	}()

	if err != nil {
		info = "Session could not be initialized"
		return
	}

	database, err = testhelpers.CreateTestDatabase(session)
	defer database.Release()
	defer database.Remove()

	if err != nil {
		info = "Database could not be created"
		return
	}

	queryresultsprocessor, err = database.GetQueryResultsProcessor()
	defer queryresultsprocessor.Release()

	if err != nil {
		info = "NotesQueryResultsProcessor could not be created"
		return
	}

	m.Run()
}

/* --------------------------------- Properties --------------------------------- */
/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_MAXENTRIES_PROPERTY.html */
func TestMaxEntries(t *testing.T) {
	_, err := queryresultsprocessor.MaxEntries()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_MAXENTRIES_PROPERTY.html */
func TestSetMaxEntries(t *testing.T) {
	s, err := queryresultsprocessor.MaxEntries()
	require.Nil(t, err)

	err = queryresultsprocessor.SetMaxEntries(s)
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_TIMEOUTSECQRP_PROPERTY.html */
func TestTimeOutSec(t *testing.T) {
	_, err := queryresultsprocessor.TimeOutSec()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_TIMEOUTSECQRP_PROPERTY.html */
func TestSetTimeOutSec(t *testing.T) {
	s, err := queryresultsprocessor.TimeOutSec()
	require.Nil(t, err)

	err = queryresultsprocessor.SetTimeOutSec(s)
	require.Nil(t, err)
}

/* --------------------------------- Methods ------------------------------------ */
/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ADDCOLLECTION_METHOD.html */
// func TestAddCollection(t *testing.T) {
// 	err := queryresultsprocessor.AddCollection( /* TODO: Pass test values. */ )
// 	require.Nil(t, err)
// }

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ADDCOLUMN_METHOD.html */
func TestAddColumn(t *testing.T) {
	err := queryresultsprocessor.AddColumn("TestColumn")
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ADDDOMINOQUERY_METHOD.html */
func TestAddDominoQuery(t *testing.T) {
	query, err := database.CreateDominoQuery()
	defer query.Release()
	require.Nil(t, err)

	err = queryresultsprocessor.AddDominoQuery(query, "@all", "TestReferenceName")
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ADDFORMULA_METHOD.html */
func TestAddFormula(t *testing.T) {
	err := queryresultsprocessor.AddFormula("@lowercase(\"\")", "TestColumn", "TestReferenceName")
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_EXECUTETOJSON_METHOD.html */
func TestExecuteToJSON(t *testing.T) {
	_, err := queryresultsprocessor.ExecuteToJSON()
	require.Nil(t, err)
}
