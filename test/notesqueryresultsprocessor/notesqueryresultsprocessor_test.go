/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_NOTESQUERYRESULTSPROCESSOR_CLASS.html */
package notesqueryresultsprocessor_test

import (
	"testing"

	"github.com/monstermichl/domigo"
	testhelpers "github.com/monstermichl/domigo/test/helpers"

	"github.com/stretchr/testify/require"
)

var database domigo.NotesDatabase
var queryresultsprocessor domigo.NotesQueryResultsProcessor

/* https://pkg.go.dev/testing#hdr-Main */
func TestMain(m *testing.M) {
	testhelpers.Initialize(func(session domigo.NotesSession, db domigo.NotesDatabase) (string, error) {
		var err error
		database = db
		
		queryresultsprocessor, err = database.GetQueryResultsProcessor()
		defer queryresultsprocessor.Release()

		if err != nil {
			return "NotesQueryResultsProcessor could not be created", err
		}
		m.Run()
		return "", nil
	})
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
