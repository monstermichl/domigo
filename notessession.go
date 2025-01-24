/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_NOTESSESSION_CLASS.html */
package domigo

import (
	"errors"

	"github.com/monstermichl/domigo/internal/com"

	ole "github.com/go-ole/go-ole"
)

type NotesSessionPolicy = Long

const (
	NOTESSESSION_POLICYSETTINGS_ARCHIVE      NotesSessionPolicy = 2
	NOTESSESSION_POLICYSETTINGS_DESKTOP      NotesSessionPolicy = 4
	NOTESSESSION_POLICYSETTINGS_REGISTRATION NotesSessionPolicy = 0
	NOTESSESSION_POLICYSETTINGS_SECURITY     NotesSessionPolicy = 3
	NOTESSESSION_POLICYSETTINGS_SETUP        NotesSessionPolicy = 1
)

type NotesSession struct {
	NotesStruct
}

func newNotesSession(dispatchPtr *ole.IDispatch) NotesSession {
	return NotesSession{newNotesStruct(dispatchPtr)}
}

func initialize(params ...any) (NotesSession, error) {
	var session NotesSession
	com, err := com.CreateObject("Lotus.NotesSession")

	/* Make sure session is released if something happened (https://stackoverflow.com/a/20507179). */
	defer func() {
		if err != nil {
			session.Release()
		}
	}()

	if err != nil {
		return session, errors.New("session could not be initialized")
	}
	session = newNotesSession(com.Dispatch())
	err = callComVoidMethod(session, "Initialize", params...)

	if err != nil {
		return session, errors.New("session could not be initialized")
	}
	return session, err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_INITIALIZE_METHOD_SESSION_COM.html */
type notesSessionInitializeParams struct {
	password *String
}

type notesSessionInitializeParam func(*notesSessionInitializeParams)

func WithNotesSessionInitializePassword(password String) notesSessionInitializeParam {
	return func(c *notesSessionInitializeParams) {
		c.password = &password
	}
}

func Initialize(params ...notesSessionInitializeParam) (NotesSession, error) {
	s := &notesSessionInitializeParams{}
	paramsOrdered := []any{}

	for _, p := range params {
		p(s)
	}

	if s.password != nil {
		paramsOrdered = append(paramsOrdered, *s.password)
	}
	return initialize(paramsOrdered...)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_INITIALIZEUSINGNOTESUSERNAME_METHOD_SESSION_COM.html */
type notesSessionInitializeUsingNotesUserNameParams struct {
	password *String
}

type notesSessionInitializeUsingNotesUserNameParam func(*notesSessionInitializeUsingNotesUserNameParams)

func WithNotesSessionInitializeUsingNotesUserNamePassword(password String) notesSessionInitializeUsingNotesUserNameParam {
	return func(c *notesSessionInitializeUsingNotesUserNameParams) {
		c.password = &password
	}
}

func InitializeUsingNotesUserName(name String, params ...notesSessionInitializeUsingNotesUserNameParam) (NotesSession, error) {
	s := &notesSessionInitializeUsingNotesUserNameParams{}
	paramsOrdered := []any{name}

	for _, p := range params {
		p(s)
	}

	if s.password != nil {
		paramsOrdered = append(paramsOrdered, *s.password)
	}
	return initialize(paramsOrdered...)
}

/* --------------------------------- Properties --------------------------------- */
/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ADDRESSBOOKS_PROPERTY.html */
func (s NotesSession) AddressBooks() ([]NotesDatabase, error) {
	return com.GetObjectArrayProperty(s.com(), newNotesDatabase, "AddressBooks")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_COMMONUSERNAME_PROPERTY.html */
func (s NotesSession) CommonUserName() (String, error) {
	return getComProperty[String](s, "CommonUserName")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_CONVERTMIME_PROPERTY.html */
func (s NotesSession) ConvertMIME() (Boolean, error) {
	return getComProperty[Boolean](s, "ConvertMIME")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_CONVERTMIME_PROPERTY.html */
func (s NotesSession) SetConvertMIME(v Boolean) error {
	return putComProperty(s, "ConvertMIME", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_CURRENTAGENT_PROPERTY.html */
func (s NotesSession) CurrentAgent() (NotesAgent, error) {
	return getComObjectProperty(s, newNotesAgent, "CurrentAgent")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_CURRENTDATABASE_PROPERTY.html */
func (s NotesSession) CurrentDatabase() (NotesDatabase, error) {
	return getComObjectProperty(s, newNotesDatabase, "CurrentDatabase")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_DOCUMENTCONTEXT_PROPERTY.html */
func (s NotesSession) DocumentContext() (NotesDocument, error) {
	return getComObjectProperty(s, newNotesDocument, "DocumentContext")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_EFFECTIVEUSERNAME_PROPERTY.html */
func (s NotesSession) EffectiveUserName() (String, error) {
	return getComProperty[String](s, "EffectiveUserName")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_HTTPURL_PROPERTY_SESSION_COM.html */
func (s NotesSession) HttpURL() (String, error) {
	return getComProperty[String](s, "HttpURL")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_INTERNATIONAL_PROPERTY.html */
func (s NotesSession) International() (NotesInternational, error) {
	return getComObjectProperty(s, newNotesInternational, "International")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISONSERVER_PROPERTY.html */
func (s NotesSession) IsOnServer() (Boolean, error) {
	return getComProperty[Boolean](s, "IsOnServer")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_LASTEXITSTATUS_PROPERTY.html */
func (s NotesSession) LastExitStatus() (Long, error) {
	return getComProperty[Long](s, "LastExitStatus")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_LASTRUN_PROPERTY_SESSION.html */
func (s NotesSession) LastRun() (Time, error) {
	return getComProperty[Time](s, "LastRun")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_NOTESBUILDVERSION_PROPERTY_4714_ABOUT.html */
func (s NotesSession) NotesBuildVersion() (Long, error) {
	return getComProperty[Long](s, "NotesBuildVersion")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_NOTESURL_PROPERTY_SESSION_COM.html */
func (s NotesSession) NotesURL() (String, error) {
	return getComProperty[String](s, "NotesURL")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_NOTESVERSION_PROPERTY.html */
func (s NotesSession) NotesVersion() (String, error) {
	return getComProperty[String](s, "NotesVersion")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ORGDIRECTORYPATH_PROPERTY_SESSION.html */
func (s NotesSession) OrgDirectoryPath() (String, error) {
	return getComProperty[String](s, "OrgDirectoryPath")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_PLATFORM_PROPERTY.html */
func (s NotesSession) Platform() (String, error) {
	return getComProperty[String](s, "Platform")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SAVEDDATA_PROPERTY.html */
func (s NotesSession) SavedData() (NotesDocument, error) {
	return getComObjectProperty(s, newNotesDocument, "SavedData")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SERVERNAME_PROPERTY_SESSION_COM.html */
func (s NotesSession) ServerName() (String, error) {
	return getComProperty[String](s, "ServerName")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_URLDATABASE_PROPERTY_SESSION_COM.html */
func (s NotesSession) URLDatabase() (NotesDatabase, error) {
	return getComObjectProperty(s, newNotesDatabase, "URLDatabase")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_USERGROUPNAMELIST_PROPERTY_SESSION.html */
func (s NotesSession) UserGroupNameList() ([]NotesName, error) {
	return com.GetObjectArrayProperty(s.com(), newNotesName, "UserGroupNameList")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_USERNAME_PROPERTY.html */
func (s NotesSession) UserName() (String, error) {
	return getComProperty[String](s, "UserName")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_USERNAMELIST_PROPERTY_7482.html */
func (s NotesSession) UserNameList() ([]NotesName, error) {
	return com.GetObjectArrayProperty(s.com(), newNotesName, "UserNameList")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_USERNAMEOBJECT_PROPERTY_SESSION_COM.html */
func (s NotesSession) UserNameObject() (NotesName, error) {
	return getComObjectProperty(s, newNotesName, "UserNameObject")
}

/* --------------------------------- Methods ------------------------------------ */
/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_CREATEADMINISTRATIONPROCESS_METHOD_SESSION.html */
func (s NotesSession) CreateAdministrationProcess(server String) (NotesAdministrationProcess, error) {
	return callComObjectMethod(s, newNotesAdministrationProcess, "CreateAdministrationProcess", server)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_CREATECOLOROBJECT_METHOD_SESSION.html */
func (s NotesSession) CreateColorObject() (NotesColorObject, error) {
	return callComObjectMethod(s, newNotesColorObject, "CreateColorObject")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_CREATEDATERANGE_METHOD.html */
func (s NotesSession) CreateDateRange() (NotesDateRange, error) {
	return callComObjectMethod(s, newNotesDateRange, "CreateDateRange")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_CREATEDATETIME_METHOD.html */
func (s NotesSession) CreateDateTime(dateTime String) (NotesDateTime, error) {
	return callComObjectMethod(s, newNotesDateTime, "CreateDateTime", dateTime)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_CREATEDXLEXPORTER_METHOD_SESSION.html */
func (s NotesSession) CreateDXLExporter() (NotesDXLExporter, error) {
	return callComObjectMethod(s, newNotesDXLExporter, "CreateDXLExporter")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_CREATEDXLIMPORTER_METHOD_SESSION.html */
func (s NotesSession) CreateDXLImporter() (NotesDXLImporter, error) {
	return callComObjectMethod(s, newNotesDXLImporter, "CreateDXLImporter")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_CREATELOG_METHOD.html */
func (s NotesSession) CreateLog(programName String) (NotesLog, error) {
	return callComObjectMethod(s, newNotesLog, "CreateLog", programName)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_CREATENAME_METHOD.html */
type notesSessionCreateNameParams struct {
	language *String
}

type notesSessionCreateNameParam func(*notesSessionCreateNameParams)

func WithNotesSessionCreateNameLanguage(language String) notesSessionCreateNameParam {
	return func(c *notesSessionCreateNameParams) {
		c.language = &language
	}
}

func (s NotesSession) CreateName(name String, params ...notesSessionCreateNameParam) (NotesName, error) {
	paramsStruct := &notesSessionCreateNameParams{}
	paramsOrdered := []any{name}

	for _, p := range params {
		p(paramsStruct)
	}

	if paramsStruct.language != nil {
		paramsOrdered = append(paramsOrdered, *paramsStruct.language)
	}
	return callComObjectMethod(s, newNotesName, "CreateName", paramsOrdered...)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_CREATENEWSLETTER_METHOD.html */
func (s NotesSession) CreateNewsletter(notesDocumentCollection NotesDocumentCollection) (NotesNewsletter, error) {
	return callComObjectMethod(s, newNotesNewsletter, "CreateNewsletter", notesDocumentCollection)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_CREATEREGISTRATION_METHOD_SESSION_COM.html */
func (s NotesSession) CreateRegistration() (NotesRegistration, error) {
	return callComObjectMethod(s, newNotesRegistration, "CreateRegistration")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_CREATERICHTEXTPARAGRAPHSTYLE_METHOD_8901_ABOUT.html */
func (s NotesSession) CreateRichTextParagraphStyle() (NotesRichTextParagraphStyle, error) {
	return callComObjectMethod(s, newNotesRichTextParagraphStyle, "CreateRichTextParagraphStyle")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_CREATERICHTEXTSTYLE_METHOD.html */
func (s NotesSession) CreateRichTextStyle() (NotesRichTextStyle, error) {
	return callComObjectMethod(s, newNotesRichTextStyle, "CreateRichTextStyle")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_CREATESAXPARSER_METHOD_SESSION.html */
/* TODO: Implement this function correctly (see documentation). */
// type notesSessionCreateSAXParserParams struct {
//     input *String
//     output *Object
// }

// type notesSessionCreateSAXParserParam func(*notesSessionCreateSAXParserParams)

// func WithNotesSessionCreateSAXParserInput(input String) notesSessionCreateSAXParserParam {
//     return func(c *notesSessionCreateSAXParserParams) {
//         c.input = &input
//     }
// }

// func WithNotesSessionCreateSAXParserOutput(output Object) notesSessionCreateSAXParserParam {
//     return func(c *notesSessionCreateSAXParserParams) {
//         c.output = &output
//     }
// }

// func (s NotesSession) CreateSAXParser(params ...notesSessionCreateSAXParserParam) (Object, error) {
//     paramsStruct := &notesSessionCreateSAXParserParams{}
//     paramsOrdered := []any{}

//     for _, p := range params {
//         p(paramsStruct)
//     }

//     if paramsStruct.input != nil {
//         paramsOrdered = append(paramsOrdered, *paramsStruct.input)
//             if paramsStruct.output != nil {
//                 paramsOrdered = append(paramsOrdered, *paramsStruct.output)
//             }
//     }
//     dispatchPtr, err := CallComObjectMethod(s, "CreateSAXParser", paramsOrdered...)
//     return New(dispatchPtr), err
// }

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_CREATESTREAM_METHOD_SESSION.html */
func (s NotesSession) CreateStream() (NotesStream, error) {
	return callComObjectMethod(s, newNotesStream, "CreateStream")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_CREATEXSLTRANSFORMER_METHOD_SESSION.html */
/* TODO: Implement this function correctly (see documentation). */
// type notesSessionCreateXSLTransformerParams struct {
//     input *String
//     styleSheet *Object
//     output *Object
// }

// type notesSessionCreateXSLTransformerParam func(*notesSessionCreateXSLTransformerParams)

// func WithNotesSessionCreateXSLTransformerInput(input String) notesSessionCreateXSLTransformerParam {
//     return func(c *notesSessionCreateXSLTransformerParams) {
//         c.input = &input
//     }
// }

// func WithNotesSessionCreateXSLTransformerStyleSheet(styleSheet Object) notesSessionCreateXSLTransformerParam {
//     return func(c *notesSessionCreateXSLTransformerParams) {
//         c.styleSheet = &styleSheet
//     }
// }

// func WithNotesSessionCreateXSLTransformerOutput(output Object) notesSessionCreateXSLTransformerParam {
//     return func(c *notesSessionCreateXSLTransformerParams) {
//         c.output = &output
//     }
// }

// func (s NotesSession) CreateXSLTransformer(params ...notesSessionCreateXSLTransformerParam) (Object, error) {
//     paramsStruct := &notesSessionCreateXSLTransformerParams{}
//     paramsOrdered := []any{}

//     for _, p := range params {
//         p(paramsStruct)
//     }

//     if paramsStruct.input != nil {
//         paramsOrdered = append(paramsOrdered, *paramsStruct.input)
//             if paramsStruct.styleSheet != nil {
//                 paramsOrdered = append(paramsOrdered, *paramsStruct.styleSheet)
//                         if paramsStruct.output != nil {
//                             paramsOrdered = append(paramsOrdered, *paramsStruct.output)
//                         }
//             }
//     }
//     dispatchPtr, err := CallComObjectMethod(s, "CreateXSLTransformer", paramsOrdered...)
//     return New(dispatchPtr), err
// }

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_EVALUATE_METHOD_SESSION_COM.html */
func (s NotesSession) Evaluate(formula String, doc NotesDocument) (Variant, error) {
	return callComMethod[Variant](s, "Evaluate", formula, doc)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_FREERESOURCESEARCH_METHOD.html */
type notesSessionFreeResourceSearchParams struct {
	Maxresults    *Integer
	User          *String
	Mincapacity   *Integer
	Apptunid      *String
	Server        *String
	Outputversion *Integer
}

type notesSessionFreeResourceSearchParam func(*notesSessionFreeResourceSearchParams)

func WithNotesSessionFreeResourceSearchMaxresults(Maxresults Integer) notesSessionFreeResourceSearchParam {
	return func(c *notesSessionFreeResourceSearchParams) {
		c.Maxresults = &Maxresults
	}
}

func WithNotesSessionFreeResourceSearchUser(User String) notesSessionFreeResourceSearchParam {
	return func(c *notesSessionFreeResourceSearchParams) {
		c.User = &User
	}
}

func WithNotesSessionFreeResourceSearchMincapacity(Mincapacity Integer) notesSessionFreeResourceSearchParam {
	return func(c *notesSessionFreeResourceSearchParams) {
		c.Mincapacity = &Mincapacity
	}
}

func WithNotesSessionFreeResourceSearchApptunid(Apptunid String) notesSessionFreeResourceSearchParam {
	return func(c *notesSessionFreeResourceSearchParams) {
		c.Apptunid = &Apptunid
	}
}

func WithNotesSessionFreeResourceSearchServer(Server String) notesSessionFreeResourceSearchParam {
	return func(c *notesSessionFreeResourceSearchParams) {
		c.Server = &Server
	}
}

func WithNotesSessionFreeResourceSearchOutputversion(Outputversion Integer) notesSessionFreeResourceSearchParam {
	return func(c *notesSessionFreeResourceSearchParams) {
		c.Outputversion = &Outputversion
	}
}

func (s NotesSession) FreeResourceSearch(Start NotesDateTime, End NotesDateTime, Site String, Type Integer, params ...notesSessionFreeResourceSearchParam) ([]String, error) {
	paramsStruct := &notesSessionFreeResourceSearchParams{}
	paramsOrdered := []any{Start, End, Site, Type}

	for _, p := range params {
		p(paramsStruct)
	}

	if paramsStruct.Maxresults != nil {
		paramsOrdered = append(paramsOrdered, *paramsStruct.Maxresults)
		if paramsStruct.User != nil {
			paramsOrdered = append(paramsOrdered, *paramsStruct.User)
			if paramsStruct.Mincapacity != nil {
				paramsOrdered = append(paramsOrdered, *paramsStruct.Mincapacity)
				if paramsStruct.Apptunid != nil {
					paramsOrdered = append(paramsOrdered, *paramsStruct.Apptunid)
					if paramsStruct.Server != nil {
						paramsOrdered = append(paramsOrdered, *paramsStruct.Server)
						if paramsStruct.Outputversion != nil {
							paramsOrdered = append(paramsOrdered, *paramsStruct.Outputversion)
						}
					}
				}
			}
		}
	}
	return callComArrayMethod[String](s, "FreeResourceSearch", paramsOrdered...)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_FREETIMESEARCH_METHOD.html */
type notesSessionFreeTimeSearchParams struct {
	firstfit *Boolean
}

type notesSessionFreeTimeSearchParam func(*notesSessionFreeTimeSearchParams)

func WithNotesSessionFreeTimeSearchFirstfit(firstfit Boolean) notesSessionFreeTimeSearchParam {
	return func(c *notesSessionFreeTimeSearchParams) {
		c.firstfit = &firstfit
	}
}

func (s NotesSession) FreeTimeSearch(window NotesDateRange, duration Integer, names []String, params ...notesSessionFreeTimeSearchParam) (NotesDateRange, error) {
	paramsStruct := &notesSessionFreeTimeSearchParams{}
	paramsOrdered := []any{window, duration, names}

	for _, p := range params {
		p(paramsStruct)
	}

	if paramsStruct.firstfit != nil {
		paramsOrdered = append(paramsOrdered, *paramsStruct.firstfit)
	}
	return callComObjectMethod(s, newNotesDateRange, "FreeTimeSearch", paramsOrdered...)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GETDATABASE_METHOD.html */
type notesSessionGetDatabaseParams struct {
	createonfail *Boolean
}

type notesSessionGetDatabaseParam func(*notesSessionGetDatabaseParams)

func WithNotesSessionGetDatabaseCreateonfail(createonfail Boolean) notesSessionGetDatabaseParam {
	return func(c *notesSessionGetDatabaseParams) {
		c.createonfail = &createonfail
	}
}

func (s NotesSession) GetDatabase(server String, dbfile String, params ...notesSessionGetDatabaseParam) (NotesDatabase, error) {
	paramsStruct := &notesSessionGetDatabaseParams{}
	paramsOrdered := []any{server, dbfile}

	for _, p := range params {
		p(paramsStruct)
	}

	if paramsStruct.createonfail != nil {
		paramsOrdered = append(paramsOrdered, *paramsStruct.createonfail)
	}
	return callComObjectMethod(s, newNotesDatabase, "GetDatabase", paramsOrdered...)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GETCALENDAR_METHOD.html */
func (s NotesSession) GetCalendar(mailDatabase NotesDatabase) (NotesCalendar, error) {
	return callComObjectMethod(s, newNotesCalendar, "GetCalendar", mailDatabase)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GETDBDIRECTORY_METHOD.html */
func (s NotesSession) GetDbDirectory(serverName String) (NotesDbDirectory, error) {
	return callComObjectMethod(s, newNotesDbDirectory, "GetDbDirectory", serverName)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GETDIRECTORY_METHOD.html */
type notesSessionGetDirectoryParams struct {
	serverName *String
}

type notesSessionGetDirectoryParam func(*notesSessionGetDirectoryParams)

func WithNotesSessionGetDirectoryServerName(serverName String) notesSessionGetDirectoryParam {
	return func(c *notesSessionGetDirectoryParams) {
		c.serverName = &serverName
	}
}

func (s NotesSession) GetDirectory(params ...notesSessionGetDirectoryParam) (NotesDirectory, error) {
	paramsStruct := &notesSessionGetDirectoryParams{}
	paramsOrdered := []any{}

	for _, p := range params {
		p(paramsStruct)
	}

	if paramsStruct.serverName != nil {
		paramsOrdered = append(paramsOrdered, *paramsStruct.serverName)
	}
	return callComObjectMethod(s, newNotesDirectory, "GetDirectory", paramsOrdered...)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GETENVIRONMENTSTRING_METHOD.html */
type notesSessionGetEnvironmentStringParams struct {
	system *Boolean
}

type notesSessionGetEnvironmentStringParam func(*notesSessionGetEnvironmentStringParams)

func WithNotesSessionGetEnvironmentStringSystem(system Boolean) notesSessionGetEnvironmentStringParam {
	return func(c *notesSessionGetEnvironmentStringParams) {
		c.system = &system
	}
}

func (s NotesSession) GetEnvironmentString(name String, params ...notesSessionGetEnvironmentStringParam) (String, error) {
	paramsStruct := &notesSessionGetEnvironmentStringParams{}
	paramsOrdered := []any{name}

	for _, p := range params {
		p(paramsStruct)
	}

	if paramsStruct.system != nil {
		paramsOrdered = append(paramsOrdered, *paramsStruct.system)
	}
	return callComMethod[String](s, "GetEnvironmentString", paramsOrdered...)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GETENVIRONMENTVALUE_METHOD.html */
type notesSessionGetEnvironmentValueParams struct {
	system *Boolean
}

type notesSessionGetEnvironmentValueParam func(*notesSessionGetEnvironmentValueParams)

func WithNotesSessionGetEnvironmentValueSystem(system Boolean) notesSessionGetEnvironmentValueParam {
	return func(c *notesSessionGetEnvironmentValueParams) {
		c.system = &system
	}
}

func (s NotesSession) GetEnvironmentValue(name String, params ...notesSessionGetEnvironmentValueParam) (any, error) {
	paramsStruct := &notesSessionGetEnvironmentValueParams{}
	paramsOrdered := []any{name}

	for _, p := range params {
		p(paramsStruct)
	}

	if paramsStruct.system != nil {
		paramsOrdered = append(paramsOrdered, *paramsStruct.system)
	}
	return callComAnyMethod(s, "GetEnvironmentValue", paramsOrdered...)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GETPROPERTYBROKER_METHOD.html */
func (s NotesSession) GetPropertyBroker() (NotesPropertyBroker, error) {
	return callComObjectMethod(s, newNotesPropertyBroker, "GetPropertyBroker")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GETUSERPOLICYSETTINGS_METHOD_SESSION.html */
type notesSessionGetUserPolicySettingsParams struct {
	explicitPolicy *String
	reserved       *String
}

type notesSessionGetUserPolicySettingsParam func(*notesSessionGetUserPolicySettingsParams)

func WithNotesSessionGetUserPolicySettingsExplicitPolicy(explicitPolicy String) notesSessionGetUserPolicySettingsParam {
	return func(c *notesSessionGetUserPolicySettingsParams) {
		c.explicitPolicy = &explicitPolicy
	}
}

func WithNotesSessionGetUserPolicySettingsReserved(reserved String) notesSessionGetUserPolicySettingsParam {
	return func(c *notesSessionGetUserPolicySettingsParams) {
		c.reserved = &reserved
	}
}

func (s NotesSession) GetUserPolicySettings(server String, name String, policyType NotesSessionPolicy, params ...notesSessionGetUserPolicySettingsParam) (NotesDocument, error) {
	paramsStruct := &notesSessionGetUserPolicySettingsParams{}
	paramsOrdered := []any{server, name, policyType}

	for _, p := range params {
		p(paramsStruct)
	}

	if paramsStruct.explicitPolicy != nil {
		paramsOrdered = append(paramsOrdered, *paramsStruct.explicitPolicy)
		if paramsStruct.reserved != nil {
			paramsOrdered = append(paramsOrdered, *paramsStruct.reserved)
		}
	}
	return callComObjectMethod(s, newNotesDocument, "GetUserPolicySettings", paramsOrdered...)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_HASHPASSWORD_METHOD_SESSION.html */
func (s NotesSession) HashPassword(password String) (String, error) {
	return callComMethod[String](s, "HashPassword", password)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_RESETUSERPASSWORD_METHOD_SESSION_COM.html */
type notesSessionResetUserPasswordParams struct {
	downloadcount *Integer
}

type notesSessionResetUserPasswordParam func(*notesSessionResetUserPasswordParams)

func WithNotesSessionResetUserPasswordDownloadcount(downloadcount Integer) notesSessionResetUserPasswordParam {
	return func(c *notesSessionResetUserPasswordParams) {
		c.downloadcount = &downloadcount
	}
}

func (s NotesSession) ResetUserPassword(servername String, username String, password String, params ...notesSessionResetUserPasswordParam) error {
	paramsStruct := &notesSessionResetUserPasswordParams{}
	paramsOrdered := []any{servername, username, password}

	for _, p := range params {
		p(paramsStruct)
	}

	if paramsStruct.downloadcount != nil {
		paramsOrdered = append(paramsOrdered, *paramsStruct.downloadcount)
	}
	return callComVoidMethod(s, "ResetUserPassword", paramsOrdered...)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_RESOLVE_METHOD_SESSION_COM.html */
func (s NotesSession) Resolve(url String) error {
	return callComVoidMethod(s, "Resolve", url)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SENDCONSOLECOMMAND_METHOD_SESSION.html */
func (s NotesSession) SendConsoleCommand(serverName String, consoleCommand String) (String, error) {
	return callComMethod[String](s, "SendConsoleCommand", serverName, consoleCommand)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SETENVIRONMENTVAR_METHOD.html */
type notesSessionSetEnvironmentVarParams struct {
	issystemvar *Boolean
}

type notesSessionSetEnvironmentVarParam func(*notesSessionSetEnvironmentVarParams)

func WithNotesSessionSetEnvironmentVarIssystemvar(issystemvar Boolean) notesSessionSetEnvironmentVarParam {
	return func(c *notesSessionSetEnvironmentVarParams) {
		c.issystemvar = &issystemvar
	}
}

func (s NotesSession) SetEnvironmentVar(name String, valueV any, params ...notesSessionSetEnvironmentVarParam) error {
	paramsStruct := &notesSessionSetEnvironmentVarParams{}
	paramsOrdered := []any{name, valueV}

	for _, p := range params {
		p(paramsStruct)
	}

	if paramsStruct.issystemvar != nil {
		paramsOrdered = append(paramsOrdered, *paramsStruct.issystemvar)
	}
	return callComVoidMethod(s, "SetEnvironmentVar", paramsOrdered...)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_UPDATEPROCESSEDDOC_METHOD.html */
func (s NotesSession) UpdateProcessedDoc(notesDocument NotesDocument) error {
	return callComVoidMethod(s, "UpdateProcessedDoc", notesDocument)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_USEDOUBLEASPOINTER_METHOD_SESSION.html */
func (s NotesSession) UseDoubleAsPointer() error {
	return callComVoidMethod(s, "UseDoubleAsPointer")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_VERIFYPASSWORD_METHOD_SESSION.html */
func (s NotesSession) VerifyPassword(password String, hashedPassword String) (Boolean, error) {
	return callComMethod[Boolean](s, "VerifyPassword", password, hashedPassword)
}
