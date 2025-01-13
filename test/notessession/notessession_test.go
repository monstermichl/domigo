package notessession_test

import (
	"testing"

	testhelpers "github.com/monstermichl/domigo/test/helpers"

	"github.com/monstermichl/domigo"
	"github.com/stretchr/testify/require"
)

var session domigo.NotesSession
var db domigo.NotesDatabase

/* https://pkg.go.dev/testing#hdr-Main */
func TestMain(m *testing.M) {
	testhelpers.Initialize(func(sessionTemp domigo.NotesSession, dbTemp domigo.NotesDatabase) (string, error) {
		session = sessionTemp
		db = dbTemp

		m.Run()
		return "", nil
	})
}

/* --------------------------------- Properties --------------------------------- */
/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ADDRESSBOOKS_PROPERTY.html */
func TestAddressBooks(t *testing.T) {
	_, err := session.AddressBooks()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_COMMONUSERNAME_PROPERTY.html */
func TestCommonUserName(t *testing.T) {
	_, err := session.CommonUserName()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_CONVERTMIME_PROPERTY.html */
func TestConvertMIME(t *testing.T) {
	_, err := session.ConvertMIME()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_CONVERTMIME_PROPERTY.html */
func TestSetConvertMIME(t *testing.T) {
	s, err := session.ConvertMIME()
	require.Nil(t, err)

	err = session.SetConvertMIME(s)
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_CURRENTAGENT_PROPERTY.html */
func TestCurrentAgent(t *testing.T) {
	_, err := session.CurrentAgent()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_CURRENTDATABASE_PROPERTY.html */
func TestCurrentDatabase(t *testing.T) {
	_, err := session.CurrentDatabase()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_DOCUMENTCONTEXT_PROPERTY.html */
func TestDocumentContext(t *testing.T) {
	_, err := session.DocumentContext()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_EFFECTIVEUSERNAME_PROPERTY.html */
func TestEffectiveUserName(t *testing.T) {
	_, err := session.EffectiveUserName()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_HTTPURL_PROPERTY_SESSION_COM.html */
func TestHttpURL(t *testing.T) {
	_, err := session.HttpURL()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_INTERNATIONAL_PROPERTY.html */
func TestInternational(t *testing.T) {
	_, err := session.International()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISONSERVER_PROPERTY.html */
func TestIsOnServer(t *testing.T) {
	_, err := session.IsOnServer()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_LASTEXITSTATUS_PROPERTY.html */
func TestLastExitStatus(t *testing.T) {
	_, err := session.LastExitStatus()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_LASTRUN_PROPERTY_SESSION.html */
func TestLastRun(t *testing.T) {
	_, err := session.LastRun()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_NOTESBUILDVERSION_PROPERTY_4714_ABOUT.html */
func TestNotesBuildVersion(t *testing.T) {
	_, err := session.NotesBuildVersion()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_NOTESURL_PROPERTY_SESSION_COM.html */
func TestNotesURL(t *testing.T) {
	_, err := session.NotesURL()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_NOTESVERSION_PROPERTY.html */
func TestNotesVersion(t *testing.T) {
	_, err := session.NotesVersion()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ORGDIRECTORYPATH_PROPERTY_SESSION.html */
func TestOrgDirectoryPath(t *testing.T) {
	_, err := session.OrgDirectoryPath()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_PLATFORM_PROPERTY.html */
func TestPlatform(t *testing.T) {
	_, err := session.Platform()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SAVEDDATA_PROPERTY.html */
func TestSavedData(t *testing.T) {
	_, err := session.SavedData()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SERVERNAME_PROPERTY_SESSION_COM.html */
func TestServerName(t *testing.T) {
	_, err := session.ServerName()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_URLDATABASE_PROPERTY_SESSION_COM.html */
func TestURLDatabase(t *testing.T) {
	_, err := session.URLDatabase()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_USERGROUPNAMELIST_PROPERTY_SESSION.html */
func TestUserGroupNameList(t *testing.T) {
	_, err := session.UserGroupNameList()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_USERNAME_PROPERTY.html */
func TestUserName(t *testing.T) {
	_, err := session.UserName()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_USERNAMELIST_PROPERTY_7482.html */
func TestUserNameList(t *testing.T) {
	_, err := session.UserNameList()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_USERNAMEOBJECT_PROPERTY_SESSION_COM.html */
func TestUserNameObject(t *testing.T) {
	_, err := session.UserNameObject()
	require.Nil(t, err)
}

/* --------------------------------- Methods ------------------------------------ */
/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_INITIALIZEUSINGNOTESUSERNAME_METHOD_SESSION_COM.html */
func TestInitializeUsingNotesUserName(t *testing.T) {
	session, err := domigo.InitializeUsingNotesUserName("")
	defer session.Release()

	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_CREATEADMINISTRATIONPROCESS_METHOD_SESSION.html */
func TestCreateAdministrationProcess(t *testing.T) {
	_, err := session.CreateAdministrationProcess("")
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_CREATECOLOROBJECT_METHOD_SESSION.html */
func TestCreateColorObject(t *testing.T) {
	_, err := session.CreateColorObject()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_CREATEDATERANGE_METHOD.html */
func TestCreateDateRange(t *testing.T) {
	_, err := session.CreateDateRange()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_CREATEDATETIME_METHOD.html */
func TestCreateDateTime(t *testing.T) {
	_, err := session.CreateDateTime("Yesterday")
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_CREATEDOMPARSER_METHOD_SESSION.html */
// func TestCreateDOMParser(t *testing.T) {
// 	_, err := session.CreateDOMParser()
// 	require.Nil(t, err)
// }

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_CREATEDXLEXPORTER_METHOD_SESSION.html */
// func TestCreateDXLExporter(t *testing.T) {
// 	_, err := session.CreateDXLExporter()
// 	require.Nil(t, err)
// }

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_CREATEDXLIMPORTER_METHOD_SESSION.html */
// func TestCreateDXLImporter(t *testing.T) {
// 	_, err := session.CreateDXLImporter( /* TODO: Pass test values. */ )
// 	require.Nil(t, err)
// }

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_CREATELOG_METHOD.html */
func TestCreateLog(t *testing.T) {
	_, err := session.CreateLog("testLog")
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_CREATENAME_METHOD.html */
func TestCreateName(t *testing.T) {
	_, err := session.CreateName("testName")
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_CREATENEWSLETTER_METHOD.html */
func TestCreateNewsletter(t *testing.T) {
	collection, err := db.CreateDocumentCollection()
	defer collection.Release()
	require.Nil(t, err)

	_, err = session.CreateNewsletter(collection)
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_CREATEREGISTRATION_METHOD_SESSION_COM.html */
func TestCreateRegistration(t *testing.T) {
	_, err := session.CreateRegistration()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_CREATERICHTEXTPARAGRAPHSTYLE_METHOD_8901_ABOUT.html */
func TestCreateRichTextParagraphStyle(t *testing.T) {
	_, err := session.CreateRichTextParagraphStyle()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_CREATERICHTEXTSTYLE_METHOD.html */
func TestCreateRichTextStyle(t *testing.T) {
	_, err := session.CreateRichTextStyle()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_CREATESAXPARSER_METHOD_SESSION.html */
// func TestCreateSAXParser(t *testing.T) {
// 	_, err := session.CreateSAXParser( /* TODO: Pass test values. */ )
// 	require.Nil(t, err)
// }

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_CREATESTREAM_METHOD_SESSION.html */
func TestCreateStream(t *testing.T) {
	_, err := session.CreateStream()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_CREATEXSLTRANSFORMER_METHOD_SESSION.html */
// func TestCreateXSLTransformer(t *testing.T) {
// 	_, err := session.CreateXSLTransformer( /* TODO: Pass test values. */ )
// 	require.Nil(t, err)
// }

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_EVALUATE_METHOD_SESSION_COM.html */
func TestEvaluate(t *testing.T) {
	doc, err := db.CreateDocument()
	defer doc.Release()
	require.Nil(t, err)

	_, err = session.Evaluate("SELECT @All", doc)
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_FREERESOURCESEARCH_METHOD.html */
func TestFreeResourceSearch(t *testing.T) {
	start, err := session.CreateDateTime("Yesterday")
	defer start.Release()
	require.Nil(t, err)

	end, err := session.CreateDateTime("Today")
	defer end.Release()
	require.Nil(t, err)

	_, err = session.FreeResourceSearch(start, end, "", 1)
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_FREETIMESEARCH_METHOD.html */
func TestFreeTimeSearch(t *testing.T) {
	dateRange, err := session.CreateDateRange()
	defer dateRange.Release()
	require.Nil(t, err)

	_, err = session.FreeTimeSearch(dateRange, 60, []string{})
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GETDATABASE_METHOD.html */
func TestGetDatabase(t *testing.T) {
	_, err := session.GetDatabase("", "GoInterface.nsf")
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GETCALENDAR_METHOD.html */
func TestGetCalendar(t *testing.T) {
	_, err := session.GetCalendar(db)
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GETDBDIRECTORY_METHOD.html */
func TestGetDbDirectory(t *testing.T) {
	_, err := session.GetDbDirectory("")
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GETDIRECTORY_METHOD.html */
func TestGetDirectory(t *testing.T) {
	_, err := session.GetDirectory()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SETENVIRONMENTVAR_METHOD.html */
func TestSetEnvironmentVar(t *testing.T) {
	err := session.SetEnvironmentVar("testEnvVar", "testValue")
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GETENVIRONMENTSTRING_METHOD.html */
func TestGetEnvironmentString(t *testing.T) {
	_, err := session.GetEnvironmentString("testEnvVar")
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GETENVIRONMENTVALUE_METHOD.html */
func TestGetEnvironmentValue(t *testing.T) {
	_, err := session.GetEnvironmentValue("testEnvVar")
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GETPROPERTYBROKER_METHOD.html */
func TestGetPropertyBroker(t *testing.T) {
	_, err := session.GetPropertyBroker()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GETUSERPOLICYSETTINGS_METHOD_SESSION.html */
// func TestGetUserPolicySettings(t *testing.T) {
// 	_, err := session.GetUserPolicySettings("", "Anonymous", 4)
// 	require.Nil(t, err)
// }

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_HASHPASSWORD_METHOD_SESSION.html */
func TestHashPassword(t *testing.T) {
	_, err := session.HashPassword("abc")
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_RESETUSERPASSWORD_METHOD_SESSION_COM.html */
func TestResetUserPassword(t *testing.T) {
	err := session.ResetUserPassword("", "Anonymous", "")
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_RESOLVE_METHOD_SESSION_COM.html */
func TestResolve(t *testing.T) {
	err := session.Resolve("testurl")
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SENDCONSOLECOMMAND_METHOD_SESSION.html */
func TestSendConsoleCommand(t *testing.T) {
	_, err := session.SendConsoleCommand("", "")
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_UPDATEPROCESSEDDOC_METHOD.html */
func TestUpdateProcessedDoc(t *testing.T) {
	doc, err := db.CreateDocument()
	defer doc.Release()
	require.Nil(t, err)

	_, err = doc.Save(true, false)
	require.Nil(t, err)

	err = session.UpdateProcessedDoc(doc)
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_USEDOUBLEASPOINTER_METHOD_SESSION.html */
func TestUseDoubleAsPointer(t *testing.T) {
	err := session.UseDoubleAsPointer()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_VERIFYPASSWORD_METHOD_SESSION.html */
func TestVerifyPassword(t *testing.T) {
	password := "abc"
	hashed, err := session.HashPassword(password)
	require.Nil(t, err)

	_, err = session.VerifyPassword(password, hashed)
	require.Nil(t, err)
}
