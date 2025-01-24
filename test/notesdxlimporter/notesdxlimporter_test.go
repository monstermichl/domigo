/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_NOTESDXLIMPORTER_CLASS.html */
package notesdxlimporter_test

import (
	"testing"

	"github.com/monstermichl/domigo"
	testhelpers "github.com/monstermichl/domigo/test/helpers"

	"github.com/stretchr/testify/require"
)

var dxlimporter domigo.NotesDXLImporter
var db domigo.NotesDatabase

/* https://pkg.go.dev/testing#hdr-Main */
func TestMain(m *testing.M) {
	testhelpers.Initialize(func(session domigo.NotesSession, dbTemp domigo.NotesDatabase) (string, error) {
		var err error

		db = dbTemp
		dxlimporter, err = session.CreateDXLImporter()
		defer dxlimporter.Release()

		if err != nil {
			return "NotesDXLImporter could not be created", err
		}
		m.Run()
		return "", nil
	})
}

/* --------------------------------- Properties --------------------------------- */
/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ACLIMPORTOPTION_PROPERTY_IMPORTER.html */
func TestACLImportOption(t *testing.T) {
	_, err := dxlimporter.ACLImportOption()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ACLIMPORTOPTION_PROPERTY_IMPORTER.html */
func TestSetACLImportOption(t *testing.T) {
	s, err := dxlimporter.ACLImportOption()
	require.Nil(t, err)

	err = dxlimporter.SetACLImportOption(s)
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_COMPILELOTUSSCRIPT_PROPERTY_IMPORTER.html */
func TestCompileLotusScript(t *testing.T) {
	_, err := dxlimporter.CompileLotusScript()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_COMPILELOTUSSCRIPT_PROPERTY_IMPORTER.html */
func TestSetCompileLotusScript(t *testing.T) {
	s, err := dxlimporter.CompileLotusScript()
	require.Nil(t, err)

	err = dxlimporter.SetCompileLotusScript(s)
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_CREATEFTINDEX_PROPERTY_IMPORTER.html */
func TestCreateFTIndex(t *testing.T) {
	_, err := dxlimporter.CreateFTIndex()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_CREATEFTINDEX_PROPERTY_IMPORTER.html */
func TestSetCreateFTIndex(t *testing.T) {
	s, err := dxlimporter.CreateFTIndex()
	require.Nil(t, err)

	err = dxlimporter.SetCreateFTIndex(s)
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_DESIGNIMPORTOPTION_PROPERTY_IMPORTER.html */
func TestDesignImportOption(t *testing.T) {
	_, err := dxlimporter.DesignImportOption()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_DESIGNIMPORTOPTION_PROPERTY_IMPORTER.html */
func TestSetDesignImportOption(t *testing.T) {
	s, err := dxlimporter.DesignImportOption()
	require.Nil(t, err)

	err = dxlimporter.SetDesignImportOption(s)
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_DOCUMENTIMPORTOPTION_PROPERTY_IMPORTER.html */
func TestDocumentImportOption(t *testing.T) {
	_, err := dxlimporter.DocumentImportOption()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_DOCUMENTIMPORTOPTION_PROPERTY_IMPORTER.html */
func TestSetDocumentImportOption(t *testing.T) {
	s, err := dxlimporter.DocumentImportOption()
	require.Nil(t, err)

	err = dxlimporter.SetDocumentImportOption(s)
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_COUNT_PROPERTY_DXLIMPORTER.html */
func TestImportedNoteCount(t *testing.T) {
	_, err := dxlimporter.ImportedNoteCount()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_INPUTVALIDATIONOPTION_PROPERTY_IMPORTER.html */
func TestInputValidationOption(t *testing.T) {
	_, err := dxlimporter.InputValidationOption()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_INPUTVALIDATIONOPTION_PROPERTY_IMPORTER.html */
func TestSetInputValidationOption(t *testing.T) {
	s, err := dxlimporter.InputValidationOption()
	require.Nil(t, err)

	err = dxlimporter.SetInputValidationOption(s)
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_REPLACEDBPROPERTIES_PROPERTY_IMPORTER.html */
func TestReplaceDbProperties(t *testing.T) {
	_, err := dxlimporter.ReplaceDbProperties()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_REPLACEDBPROPERTIES_PROPERTY_IMPORTER.html */
func TestSetReplaceDbProperties(t *testing.T) {
	s, err := dxlimporter.ReplaceDbProperties()
	require.Nil(t, err)

	err = dxlimporter.SetReplaceDbProperties(s)
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_REPLICAREQUIREDFORREPLACEORUPDATE_PROPERTY_IMPORTER.html */
func TestReplicaRequiredForReplaceOrUpdate(t *testing.T) {
	_, err := dxlimporter.ReplicaRequiredForReplaceOrUpdate()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_REPLICAREQUIREDFORREPLACEORUPDATE_PROPERTY_IMPORTER.html */
func TestSetReplicaRequiredForReplaceOrUpdate(t *testing.T) {
	s, err := dxlimporter.ReplicaRequiredForReplaceOrUpdate()
	require.Nil(t, err)

	err = dxlimporter.SetReplicaRequiredForReplaceOrUpdate(s)
	require.Nil(t, err)
}

/* TODO: Access type for UnknownTokenLogOption could not be evaluated, check yourself if getter/setter is needed. */
/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_UNKNOWNTOKENLOGOPTION_PROPERTY_IMPORTER.html */
func TestUnknownTokenLogOption(t *testing.T) {
	_, err := dxlimporter.UnknownTokenLogOption()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_UNKNOWNTOKENLOGOPTION_PROPERTY_IMPORTER.html */
func TestSetUnknownTokenLogOption(t *testing.T) {
	s, err := dxlimporter.UnknownTokenLogOption()
	require.Nil(t, err)

	err = dxlimporter.SetUnknownTokenLogOption(s)
	require.Nil(t, err)
}

/* --------------------------------- Methods ------------------------------------ */
/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GETFIRSTNOTEID_METHOD_DXLIMPORTER.html */
func TestGetFirstImportedNoteID(t *testing.T) {
	_, err := dxlimporter.GetFirstImportedNoteID()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GETNEXTNOTEID_METHOD_DXLIMPORTER.html */
func TestGetNextImportedNoteID(t *testing.T) {
	_, err := dxlimporter.GetNextImportedNoteID("testNoteId")
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_IMPORT_METHOD_IMPORTER.html */
func TestImport(t *testing.T) {
	dxl := `<document form="Customers">
		<noteinfo noteid="942" unid="431A199A6FCC9C0985256A54005041A1" sequence="1"> 
			<created> 
				<datetime dst="true">20010522T103636,81-04</datetime>
			</created>
			<modified>
				<datetime dst="true">20010522T103709,78-04</datetime>
			</modified>
			<revised>
				<datetime dst="true">20010522T103709,77-04</datetime>
			</revised>
			<lastaccessed>
				<datetime dst="true">20010522T103709,77-04</datetime>
			</lastaccessed>
			<addedtofile> 
				<datetime dst="true">20010522T103709,77-04</datetime>
			</addedtofile>
		</noteinfo>
		<updatedby>
			<name>CN=Michelle Lally/OU=CAM/O=Acme</name>
		</updatedby>
		<item name="date">
			<datetime>20010601</datetime>
		</item>
		<item name="name">
			<text>Harry Hill</text>
		</item>
	</document>`

	err := dxlimporter.Import(dxl, db)
	require.Nil(t, err)
}
