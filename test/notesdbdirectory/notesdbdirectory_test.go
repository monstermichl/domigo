/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_NOTESDBDIRECTORY_CLASS.html */
package notesdbdirectory_test

import (
	"testing"

	"github.com/monstermichl/domigo"
	testhelpers "github.com/monstermichl/domigo/test/helpers"

	"github.com/stretchr/testify/require"
)

var dbdirectory domigo.NotesDbDirectory

/* https://pkg.go.dev/testing#hdr-Main */
func TestMain(m *testing.M) {
	testhelpers.Initialize(func(session domigo.NotesSession, db domigo.NotesDatabase) (string, error) {
		var err error
		dbdirectory, _ = session.GetDbDirectory("")
		defer dbdirectory.Release()

		if err != nil {
			return "Db directory could not be retrieved", err
		}
		m.Run()
		return "", nil
	})
}

/* --------------------------------- Properties --------------------------------- */
/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_NAME_PROPERTY_DIRECTORY.html */
func TestName(t *testing.T) {
	_, err := dbdirectory.Name()
	require.Nil(t, err)
}

/* --------------------------------- Methods ------------------------------------ */
/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_CREATEDATABASE_METHOD_DBDIRECTORY_COM.html */
func TestCreateDatabase(t *testing.T) {
	db, err := dbdirectory.CreateDatabase(testhelpers.TestDatabaseName())
	defer db.Release()
	require.Nil(t, err)

	err = db.Remove()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GETFIRSTDATABASE_METHOD.html */
func TestGetFirstDatabase(t *testing.T) {
	_, err := dbdirectory.GetFirstDatabase(domigo.NOTESDBDIRECTORY_FILETYPE_DATABASE)
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GETNEXTDATABASE_METHOD.html */
func TestGetNextDatabase(t *testing.T) {
	_, err := dbdirectory.GetNextDatabase()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_OPENDATABASE_METHOD_DBDIRECTORY_COM.html */
func TestOpenDatabase(t *testing.T) {
	dbNew, err := dbdirectory.CreateDatabase(testhelpers.TestDatabaseName())
	defer dbNew.Release()
	require.Nil(t, err)

	path, err := dbNew.FilePath()
	require.Nil(t, err)

	db, err := dbdirectory.OpenDatabase(path)
	defer db.Release()
	defer db.Remove()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_OPENDATABASEBYREPLICAID_METHOD_DBDIRECTORY_COM.html */
// func TestOpenDatabaseByReplicaID(t *testing.T) {
// 	_, err := dbdirectory.OpenDatabaseByReplicaID( /* TODO: Pass test values. */ )
// 	require.Nil(t, err)
// }

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_OPENDATABASEIFMODIFIED_METHOD_DBDIRECTORY_COM.html */
// func TestOpenDatabaseIfModified(t *testing.T) {
// 	_, err := dbdirectory.OpenDatabaseIfModified( /* TODO: Pass test values. */ )
// 	require.Nil(t, err)
// }

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_OPENMAILDATABASE_METHOD_DBDIRECTORY_COM.html */
func TestOpenMailDatabase(t *testing.T) {
	_, err := dbdirectory.OpenMailDatabase()
	require.Nil(t, err)
}
