/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_NOTESSESSION_CLASS.html */
package notessession

import (
	"domigo/domino"
	"domigo/domino/com"
	"domigo/domino/notesadministrationprocess"
	"domigo/domino/notesagent"
	"domigo/domino/notescolorobject"
	"domigo/domino/notesdatabase"
	"domigo/domino/notesdaterange"
	"domigo/domino/notesdatetime"
	"domigo/domino/notesdbdirectory"
	"domigo/domino/notesdirectory"
	"domigo/domino/notesdocument"
	"domigo/domino/notesdocumentcollection"
	"domigo/domino/noteslog"
	"domigo/domino/notesname"
	"domigo/domino/notesnewsletter"
	"domigo/domino/notespropertybroker"
	"domigo/domino/notesregistration"
	"domigo/domino/notesrichtextparagraphstyle"
	"domigo/domino/notesrichtextstyle"
	"domigo/domino/notesstream"
	"domigo/helpers"
	"errors"

	ole "github.com/go-ole/go-ole"
)

type NotesSession struct {
	domino.NotesStruct
}

func New(dispatchPtr *ole.IDispatch) NotesSession {
	return NotesSession{domino.NewNotesStruct(dispatchPtr)}
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_PARENT_PROPERTY_DB.html */
/* Moved from NotesDatabase. */
func NotesDatabaseParent(d notesdatabase.NotesDatabase) (NotesSession, error) {
	dispatchPtr, err := d.Com().GetObjectProperty("Parent")
	return New(dispatchPtr), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_PARENT_PROPERTY_DATERANGE_COM.html */
/* Moved from NotesDateRange. */
func NotesDateRangeParent(d notesdaterange.NotesDateRange) (NotesSession, error) {
	dispatchPtr, err := d.Com().GetObjectProperty("Parent")
	return New(dispatchPtr), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_PARENT_PROPERTY_DATETIME_COM.html */
/* Moved from NotesDateTime. */
func NotesDateTimeParent(d notesdatetime.NotesDateTime) (NotesSession, error) {
	dispatchPtr, err := d.Com().GetObjectProperty("Parent")
	return New(dispatchPtr), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_PARENT_PROPERTY_DBDIRECTORY_COM.html */
/* Moved from NotesDbDirectory. */
func NotesDbDirectoryParent(d notesdbdirectory.NotesDbDirectory) (NotesSession, error) {
	dispatchPtr, err := d.Com().GetObjectProperty("Parent")
	return New(dispatchPtr), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_PARENT_PROPERTY_LOG_COM.html */
/* Moved from NotesLog. */
func NotesLogParent(l noteslog.NotesLog) (NotesSession, error) {
	dispatchPtr, err := l.Com().GetObjectProperty("Parent")
	return New(dispatchPtr), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_PARENT_PROPERTY_NEWSLETTER_COM.html */
/* Moved from NotesNewsletter. */
func NotesNewsletterParent(n notesnewsletter.NotesNewsletter) (NotesSession, error) {
	dispatchPtr, err := n.Com().GetObjectProperty("Parent")
	return New(dispatchPtr), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_PARENT_PROPERTY_RICHTEXTSTYLE_COM.html */
/* Moved from NotesRichTextStyle */
func NotesRichTextStyleParent(r notesrichtextstyle.NotesRichTextStyle) (NotesSession, error) {
	dispatchPtr, err := r.Com().GetObjectProperty("Parent")
	return New(dispatchPtr), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_PARENT_PROPERTY_STREAM.html */
/* Moved from NotesStream. */
func NotesStreamParent(s notesstream.NotesStream) (NotesSession, error) {
	dispatchPtr, err := s.Com().GetObjectProperty("Parent")
	return New(dispatchPtr), err
}

func initialize(params ...interface{}) (NotesSession, error) {
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
	session = New(com.Dispatch())
	_, err = session.Com().CallObjectMethod("Initialize", params...)

	if err != nil {
		return session, errors.New("session could not be initialized")
	}
	return session, err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_INITIALIZE_METHOD_SESSION_COM.html */
type initializeParams struct {
	password *domino.String
}

type initializeParam func(*initializeParams)

func WithInitializePassword(password domino.String) initializeParam {
	return func(c *initializeParams) {
		c.password = &password
	}
}

func Initialize(params ...initializeParam) (NotesSession, error) {
	s := &initializeParams{}
	paramsOrdered := []interface{}{}

	for _, p := range params {
		p(s)
	}

	if s.password != nil {
		paramsOrdered = append(paramsOrdered, *s.password)
	}
	return initialize(paramsOrdered...)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_INITIALIZEUSINGNOTESUSERNAME_METHOD_SESSION_COM.html */
type initializeUsingNotesUserNameParams struct {
	password *domino.String
}

type initializeUsingNotesUserNameParam func(*initializeUsingNotesUserNameParams)

func WithInitializeUsingNotesUserNamePassword(password domino.String) initializeUsingNotesUserNameParam {
	return func(c *initializeUsingNotesUserNameParams) {
		c.password = &password
	}
}

func InitializeUsingNotesUserName(name domino.String, params ...initializeUsingNotesUserNameParam) (NotesSession, error) {
	s := &initializeUsingNotesUserNameParams{}
	paramsOrdered := []interface{}{name}

	for _, p := range params {
		p(s)
	}

	if s.password != nil {
		paramsOrdered = append(paramsOrdered, *s.password)
	}
	return initialize(paramsOrdered...)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_RESETUSERPASSWORD_METHOD_SESSION_COM.html */
type resetUserPasswordParams struct {
	downloadcount *domino.Integer
}

type resetUserPasswordParam func(*resetUserPasswordParams)

func WithResetUserPasswordDownloadcount(downloadcount domino.Integer) resetUserPasswordParam {
	return func(c *resetUserPasswordParams) {
		c.downloadcount = &downloadcount
	}
}

func (s NotesSession) ResetUserPassword(servername domino.String, username domino.String, password domino.String, params ...resetUserPasswordParam) error {
	paramsStruct := &resetUserPasswordParams{}
	paramsOrdered := []interface{}{servername, username, password}

	for _, p := range params {
		p(paramsStruct)
	}

	if paramsStruct.downloadcount != nil {
		paramsOrdered = append(paramsOrdered, *paramsStruct.downloadcount)
	}
	_, err := s.Com().CallMethod("ResetUserPassword", paramsOrdered...)
	return err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_RESOLVE_METHOD_SESSION_COM.html */
func (s NotesSession) Resolve(url domino.String) error {
	_, err := s.Com().CallMethod("Resolve", url)
	return err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SENDCONSOLECOMMAND_METHOD_SESSION.html */
func (s NotesSession) SendConsoleCommand(serverName domino.String, consoleCommand domino.String) (domino.String, error) {
	val, err := s.Com().CallMethod("SendConsoleCommand", serverName, consoleCommand)
	return helpers.CastValue[domino.String](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SETENVIRONMENTVAR_METHOD.html */
type setEnvironmentVarParams struct {
	issystemvar *domino.Boolean
}

type setEnvironmentVarParam func(*setEnvironmentVarParams)

func WithSetEnvironmentVarIssystemvar(issystemvar domino.Boolean) setEnvironmentVarParam {
	return func(c *setEnvironmentVarParams) {
		c.issystemvar = &issystemvar
	}
}

func (s NotesSession) SetEnvironmentVar(name domino.String, valueV domino.Variant, params ...setEnvironmentVarParam) error {
	paramsStruct := &setEnvironmentVarParams{}
	paramsOrdered := []interface{}{name, valueV}

	for _, p := range params {
		p(paramsStruct)
	}

	if paramsStruct.issystemvar != nil {
		paramsOrdered = append(paramsOrdered, *paramsStruct.issystemvar)
	}
	_, err := s.Com().CallMethod("SetEnvironmentVar", paramsOrdered...)
	return err
}

/* --------------------------------- Properties --------------------------------- */
/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ADDRESSBOOKS_PROPERTY.html */
func (s NotesSession) AddressBooks() ([]notesdatabase.NotesDatabase, error) {
	return com.GetObjectArrayProperty(s.Com(), notesdatabase.New, "AddressBooks")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_COMMONUSERNAME_PROPERTY.html */
func (s NotesSession) CommonUserName() (domino.String, error) {
	val, err := s.Com().GetProperty("CommonUserName")
	return helpers.CastValue[domino.String](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_CONVERTMIME_PROPERTY.html */
func (s NotesSession) ConvertMIME() (domino.Boolean, error) {
	val, err := s.Com().GetProperty("ConvertMIME")
	return helpers.CastValue[domino.Boolean](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_CONVERTMIME_PROPERTY.html */
func (s NotesSession) SetConvertMIME(v domino.Boolean) error {
	return s.Com().PutProperty("ConvertMIME", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_CURRENTAGENT_PROPERTY.html */
func (s NotesSession) CurrentAgent() (notesagent.NotesAgent, error) {
	dispatchPtr, err := s.Com().GetObjectProperty("CurrentAgent")
	return notesagent.New(dispatchPtr), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_CURRENTDATABASE_PROPERTY.html */
func (s NotesSession) CurrentDatabase() (notesdatabase.NotesDatabase, error) {
	dispatchPtr, err := s.Com().GetObjectProperty("CurrentDatabase")
	return notesdatabase.New(dispatchPtr), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_DOCUMENTCONTEXT_PROPERTY.html */
func (s NotesSession) DocumentContext() (notesdocument.NotesDocument, error) {
	dispatchPtr, err := s.Com().GetObjectProperty("DocumentContext")
	return notesdocument.New(dispatchPtr), err
}

/* TODO: Access type for EffectiveUserName could not be evaluated, check yourself if getter/setter is needed. */
/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_EFFECTIVEUSERNAME_PROPERTY.html */
func (s NotesSession) EffectiveUserName() (domino.String, error) {
	val, err := s.Com().GetProperty("EffectiveUserName")
	return helpers.CastValue[domino.String](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_EFFECTIVEUSERNAME_PROPERTY.html */
func (s NotesSession) SetEffectiveUserName(v domino.String) error {
	return s.Com().PutProperty("EffectiveUserName", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_HTTPURL_PROPERTY_SESSION_COM.html */
func (s NotesSession) HttpURL() (domino.String, error) {
	val, err := s.Com().GetProperty("HttpURL")
	return helpers.CastValue[domino.String](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_INTERNATIONAL_PROPERTY.html */
func (s NotesSession) International() (domino.Integer, error) {
	val, err := s.Com().GetProperty("International")
	return helpers.CastValue[domino.Integer](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISONSERVER_PROPERTY.html */
func (s NotesSession) IsOnServer() (domino.Boolean, error) {
	val, err := s.Com().GetProperty("IsOnServer")
	return helpers.CastValue[domino.Boolean](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_LASTEXITSTATUS_PROPERTY.html */
func (s NotesSession) LastExitStatus() (domino.Long, error) {
	val, err := s.Com().GetProperty("LastExitStatus")
	return helpers.CastValue[domino.Long](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_LASTRUN_PROPERTY_SESSION.html */
func (s NotesSession) LastRun() (domino.Time, error) {
	val, err := s.Com().GetProperty("LastRun")
	return helpers.CastValue[domino.Time](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_NOTESBUILDVERSION_PROPERTY_4714_ABOUT.html */
func (s NotesSession) NotesBuildVersion() (domino.Long, error) {
	val, err := s.Com().GetProperty("NotesBuildVersion")
	return helpers.CastValue[domino.Long](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_NOTESURL_PROPERTY_SESSION_COM.html */
func (s NotesSession) NotesURL() (domino.String, error) {
	val, err := s.Com().GetProperty("NotesURL")
	return helpers.CastValue[domino.String](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_NOTESVERSION_PROPERTY.html */
func (s NotesSession) NotesVersion() (domino.String, error) {
	val, err := s.Com().GetProperty("NotesVersion")
	return helpers.CastValue[domino.String](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ORGDIRECTORYPATH_PROPERTY_SESSION.html */
func (s NotesSession) OrgDirectoryPath() (domino.String, error) {
	val, err := s.Com().GetProperty("OrgDirectoryPath")
	return helpers.CastValue[domino.String](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_PLATFORM_PROPERTY.html */
func (s NotesSession) Platform() (domino.String, error) {
	val, err := s.Com().GetProperty("Platform")
	return helpers.CastValue[domino.String](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SAVEDDATA_PROPERTY.html */
func (s NotesSession) SavedData() (notesdocument.NotesDocument, error) {
	dispatchPtr, err := s.Com().GetObjectProperty("SavedData")
	return notesdocument.New(dispatchPtr), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SERVERNAME_PROPERTY_SESSION_COM.html */
func (s NotesSession) ServerName() (domino.String, error) {
	val, err := s.Com().GetProperty("ServerName")
	return helpers.CastValue[domino.String](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_URLDATABASE_PROPERTY_SESSION_COM.html */
func (s NotesSession) URLDatabase() (notesdatabase.NotesDatabase, error) {
	dispatchPtr, err := s.Com().GetObjectProperty("URLDatabase")
	return notesdatabase.New(dispatchPtr), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_USERGROUPNAMELIST_PROPERTY_SESSION.html */
func (s NotesSession) UserGroupNameList() ([]notesname.NotesName, error) {
	return com.GetObjectArrayProperty(s.Com(), notesname.New, "UserGroupNameList")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_USERNAME_PROPERTY.html */
func (s NotesSession) UserName() (domino.String, error) {
	val, err := s.Com().GetProperty("UserName")
	return helpers.CastValue[domino.String](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_USERNAMELIST_PROPERTY_7482.html */
func (s NotesSession) UserNameList() ([]notesname.NotesName, error) {
	return com.GetObjectArrayProperty(s.Com(), notesname.New, "UserNameList")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_USERNAMEOBJECT_PROPERTY_SESSION_COM.html */
func (s NotesSession) UserNameObject() (notesname.NotesName, error) {
	dispatchPtr, err := s.Com().GetObjectProperty("UserNameObject")
	return notesname.New(dispatchPtr), err
}

/* --------------------------------- Methods ------------------------------------ */
/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_CREATEADMINISTRATIONPROCESS_METHOD_SESSION.html */
func (s NotesSession) CreateAdministrationProcess(server domino.String) (notesadministrationprocess.NotesAdministrationProcess, error) {
	dispatchPtr, err := s.Com().CallObjectMethod("CreateAdministrationProcess", server)
	return notesadministrationprocess.New(dispatchPtr), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_CREATECOLOROBJECT_METHOD_SESSION.html */
func (s NotesSession) CreateColorObject() (notescolorobject.NotesColorObject, error) {
	dispatchPtr, err := s.Com().CallObjectMethod("CreateColorObject")
	return notescolorobject.New(dispatchPtr), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_CREATEDATERANGE_METHOD.html */
func (s NotesSession) CreateDateRange() (notesdaterange.NotesDateRange, error) {
	dispatchPtr, err := s.Com().CallObjectMethod("CreateDateRange")
	return notesdaterange.New(dispatchPtr), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_CREATEDATETIME_METHOD.html */
func (s NotesSession) CreateDateTime(dateTime domino.String) (notesdatetime.NotesDateTime, error) {
	dispatchPtr, err := s.Com().CallObjectMethod("CreateDateTime", dateTime)
	return notesdatetime.New(dispatchPtr), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_CREATEDOMPARSER_METHOD_SESSION.html */
/* TODO: Implement this function correctly (see documentation). */
// func (s NotesSession) CreateDOMParser(input domino.Unknown, output domino.Unknown) (notesdomparser.NotesDOMParser, error) {
// 	dispatchPtr, err := s.Com().CallObjectMethod("CreateDOMParser", input, output)
// 	return notesdomparser.New(dispatchPtr), err
// }

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_CREATEDXLEXPORTER_METHOD_SESSION.html */
/* TODO: Implement this function correctly (see documentation). */
// func (s NotesSession) CreateDXLExporter(input domino.Unknown, output domino.Unknown) (notesdxlexporter.NotesDXLExporter, error) {
// 	dispatchPtr, err := s.Com().CallObjectMethod("CreateDXLExporter", input, output)
// 	return notesdxlexporter.New(dispatchPtr), err
// }

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_CREATEDXLIMPORTER_METHOD_SESSION.html */
/* TODO: Implement this function correctly (see documentation). */
// func (s NotesSession) CreateDXLImporter(input domino.Unknown, output domino.Unknown) (notesdxlimporter.NotesDXLImporter, error) {
// 	dispatchPtr, err := s.Com().CallObjectMethod("CreateDXLImporter", input, output)
// 	return notesdxlimporter.New(dispatchPtr), err
// }

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_CREATELOG_METHOD.html */
func (s NotesSession) CreateLog(programName domino.String) (noteslog.NotesLog, error) {
	dispatchPtr, err := s.Com().CallObjectMethod("CreateLog", programName)
	return noteslog.New(dispatchPtr), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_CREATENAME_METHOD.html */
type createNameParams struct {
	language *domino.String
}

type createNameParam func(*createNameParams)

func WithCreateNameLanguage(language domino.String) createNameParam {
	return func(c *createNameParams) {
		c.language = &language
	}
}

func (s NotesSession) CreateName(name domino.String, params ...createNameParam) (notesname.NotesName, error) {
	paramsStruct := &createNameParams{}
	paramsOrdered := []interface{}{name}

	for _, p := range params {
		p(paramsStruct)
	}

	if paramsStruct.language != nil {
		paramsOrdered = append(paramsOrdered, *paramsStruct.language)
	}
	dispatchPtr, err := s.Com().CallObjectMethod("CreateName", paramsOrdered...)
	return notesname.New(dispatchPtr), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_CREATENEWSLETTER_METHOD.html */
func (s NotesSession) CreateNewsletter(notesDocumentCollection notesdocumentcollection.NotesDocumentCollection) (notesnewsletter.NotesNewsletter, error) {
	dispatchPtr, err := s.Com().CallObjectMethod("CreateNewsletter", notesDocumentCollection.Com().Dispatch())
	return notesnewsletter.New(dispatchPtr), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_CREATEREGISTRATION_METHOD_SESSION_COM.html */
func (s NotesSession) CreateRegistration() (notesregistration.NotesRegistration, error) {
	dispatchPtr, err := s.Com().CallObjectMethod("CreateRegistration")
	return notesregistration.New(dispatchPtr), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_CREATERICHTEXTPARAGRAPHSTYLE_METHOD_8901_ABOUT.html */
func (s NotesSession) CreateRichTextParagraphStyle() (notesrichtextparagraphstyle.NotesRichTextParagraphStyle, error) {
	dispatchPtr, err := s.Com().CallObjectMethod("CreateRichTextParagraphStyle")
	return notesrichtextparagraphstyle.New(dispatchPtr), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_CREATERICHTEXTSTYLE_METHOD.html */
func (s NotesSession) CreateRichTextStyle() (notesrichtextstyle.NotesRichTextStyle, error) {
	dispatchPtr, err := s.Com().CallObjectMethod("CreateRichTextStyle")
	return notesrichtextstyle.New(dispatchPtr), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_CREATESAXPARSER_METHOD_SESSION.html */
/* TODO: Implement this function correctly (see documentation). */
// func (s NotesSession) CreateSAXParser(input domino.Unknown, output domino.Unknown) (notessaxparser.NotesSAXParser, error) {
// 	dispatchPtr, err := s.Com().CallObjectMethod("CreateSAXParser", input, output)
// 	return notessaxparser.New(dispatchPtr), err
// }

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_CREATESTREAM_METHOD_SESSION.html */
func (s NotesSession) CreateStream() (notesstream.NotesStream, error) {
	dispatchPtr, err := s.Com().CallObjectMethod("CreateStream")
	return notesstream.New(dispatchPtr), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_CREATEXSLTRANSFORMER_METHOD_SESSION.html */
/* TODO: Implement this function correctly (see documentation). */
// func (s NotesSession) CreateXSLTransformer(input domino.Unknown, styleSheet domino.Unknown, output domino.Unknown) (notesxsltransformer.NotesXSLTransformer, error) {
// 	dispatchPtr, err := s.Com().CallObjectMethod("CreateXSLTransformer", input, styleSheet, output)
// 	return notesxsltransformer.New(dispatchPtr), err
// }

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_EVALUATE_METHOD_SESSION_COM.html */
func (s NotesSession) Evaluate(formula domino.String, doc notesdocument.NotesDocument) (domino.Variant, error) {
	val, err := s.Com().CallMethod("Evaluate", formula, doc.Com().Dispatch())
	return helpers.CastValue[domino.Variant](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_FREERESOURCESEARCH_METHOD.html */
type freeResourceSearchParams struct {
	User          *domino.String
	Apptunid      *domino.String
	Server        *domino.String
	Outputversion *domino.Integer
}

type freeResourceSearchParam func(*freeResourceSearchParams)

func WithFreeResourceSearchUser(User domino.String) freeResourceSearchParam {
	return func(c *freeResourceSearchParams) {
		c.User = &User
	}
}

func WithFreeResourceSearchApptunid(Apptunid domino.String) freeResourceSearchParam {
	return func(c *freeResourceSearchParams) {
		c.Apptunid = &Apptunid
	}
}

func WithFreeResourceSearchServer(Server domino.String) freeResourceSearchParam {
	return func(c *freeResourceSearchParams) {
		c.Server = &Server
	}
}

func WithFreeResourceSearchOutputversion(Outputversion domino.Integer) freeResourceSearchParam {
	return func(c *freeResourceSearchParams) {
		c.Outputversion = &Outputversion
	}
}

func (s NotesSession) FreeResourceSearch(start notesdatetime.NotesDateTime, end notesdatetime.NotesDateTime, site domino.String, resourceType domino.Integer, Maxresults domino.Integer, Mincapacity domino.Integer, params ...freeResourceSearchParam) ([]domino.String, error) {
	paramsStruct := &freeResourceSearchParams{}
	paramsOrdered := []interface{}{start.Com().Dispatch(), end.Com().Dispatch(), site, resourceType, Maxresults, Mincapacity}

	for _, p := range params {
		p(paramsStruct)
	}

	if paramsStruct.User != nil {
		paramsOrdered = append(paramsOrdered, *paramsStruct.User)
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
	vals, err := s.Com().CallArrayMethod("FreeResourceSearch", paramsOrdered...)
	return helpers.CastSlice[domino.String](vals), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_FREETIMESEARCH_METHOD.html */
func (s NotesSession) FreeTimeSearch(window notesdaterange.NotesDateRange, duration domino.Integer, names []domino.String, firstfit domino.Boolean) (notesdaterange.NotesDateRange, error) {
	dispatchPtr, err := s.Com().CallObjectMethod("FreeTimeSearch", window.Com().Dispatch(), duration, names, firstfit)
	return notesdaterange.New(dispatchPtr), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GETDATABASE_METHOD.html */
type getDatabaseParams struct {
	createonfail *domino.Boolean
}

type getDatabaseParam func(*getDatabaseParams)

func WithGetDatabaseCreateonfail(createonfail domino.Boolean) getDatabaseParam {
	return func(c *getDatabaseParams) {
		c.createonfail = &createonfail
	}
}

func (s NotesSession) GetDatabase(server domino.String, dbfile domino.String, params ...getDatabaseParam) (notesdatabase.NotesDatabase, error) {
	paramsStruct := &getDatabaseParams{}
	paramsOrdered := []interface{}{server, dbfile}

	for _, p := range params {
		p(paramsStruct)
	}

	if paramsStruct.createonfail != nil {
		paramsOrdered = append(paramsOrdered, *paramsStruct.createonfail)
	}
	dispatchPtr, err := s.Com().CallObjectMethod("GetDatabase", paramsOrdered...)
	return notesdatabase.New(dispatchPtr), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GETCALENDAR_METHOD.html */
func (s NotesSession) GetCalendar() error {
	_, err := s.Com().CallMethod("GetCalendar")
	return err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GETDBDIRECTORY_METHOD.html */
func (s NotesSession) GetDbDirectory(serverName domino.String) (notesdbdirectory.NotesDbDirectory, error) {
	dispatchPtr, err := s.Com().CallObjectMethod("GetDbDirectory", serverName)
	return notesdbdirectory.New(dispatchPtr), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GETDIRECTORY_METHOD.html */
type getDirectoryParams struct {
	serverName *domino.String
}

type getDirectoryParam func(*getDirectoryParams)

func WithGetDirectoryServerName(serverName domino.String) getDirectoryParam {
	return func(c *getDirectoryParams) {
		c.serverName = &serverName
	}
}

func (s NotesSession) GetDirectory(params ...getDirectoryParam) (notesdirectory.NotesDirectory, error) {
	paramsStruct := &getDirectoryParams{}
	paramsOrdered := []interface{}{}

	for _, p := range params {
		p(paramsStruct)
	}

	if paramsStruct.serverName != nil {
		paramsOrdered = append(paramsOrdered, *paramsStruct.serverName)
	}
	dispatchPtr, err := s.Com().CallObjectMethod("GetDirectory", paramsOrdered...)
	return notesdirectory.New(dispatchPtr), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GETENVIRONMENTSTRING_METHOD.html */
type getEnvironmentStringParams struct {
	system *domino.Boolean
}

type getEnvironmentStringParam func(*getEnvironmentStringParams)

func WithGetEnvironmentStringSystem(system domino.Boolean) getEnvironmentStringParam {
	return func(c *getEnvironmentStringParams) {
		c.system = &system
	}
}

func (s NotesSession) GetEnvironmentString(name domino.String, params ...getEnvironmentStringParam) (domino.String, error) {
	paramsStruct := &getEnvironmentStringParams{}
	paramsOrdered := []interface{}{name}

	for _, p := range params {
		p(paramsStruct)
	}

	if paramsStruct.system != nil {
		paramsOrdered = append(paramsOrdered, *paramsStruct.system)
	}
	val, err := s.Com().CallMethod("GetEnvironmentString", paramsOrdered...)
	return helpers.CastValue[domino.String](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GETENVIRONMENTVALUE_METHOD.html */
type getEnvironmentValueParams struct {
	system *domino.Boolean
}

type getEnvironmentValueParam func(*getEnvironmentValueParams)

func WithGetEnvironmentValueSystem(system domino.Boolean) getEnvironmentValueParam {
	return func(c *getEnvironmentValueParams) {
		c.system = &system
	}
}

func (s NotesSession) GetEnvironmentValue(name domino.String, params ...getEnvironmentValueParam) (domino.Variant, error) {
	paramsStruct := &getEnvironmentValueParams{}
	paramsOrdered := []interface{}{name}

	for _, p := range params {
		p(paramsStruct)
	}

	if paramsStruct.system != nil {
		paramsOrdered = append(paramsOrdered, *paramsStruct.system)
	}
	val, err := s.Com().CallMethod("GetEnvironmentValue", paramsOrdered...)
	return helpers.CastValue[domino.Variant](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GETPROPERTYBROKER_METHOD.html */
func (s NotesSession) GetPropertyBroker() (notespropertybroker.NotesPropertyBroker, error) {
	dispatchPtr, err := s.Com().CallObjectMethod("GetPropertyBroker")
	return notespropertybroker.New(dispatchPtr), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GETUSERPOLICYSETTINGS_METHOD_SESSION.html */
type getUserPolicySettingsParams struct {
	explicitPolicy *domino.String
	reserved       *domino.String
}

type getUserPolicySettingsParam func(*getUserPolicySettingsParams)

func WithGetUserPolicySettingsExplicitPolicy(explicitPolicy domino.String) getUserPolicySettingsParam {
	return func(c *getUserPolicySettingsParams) {
		c.explicitPolicy = &explicitPolicy
	}
}

func WithGetUserPolicySettingsReserved(reserved domino.String) getUserPolicySettingsParam {
	return func(c *getUserPolicySettingsParams) {
		c.reserved = &reserved
	}
}

func (s NotesSession) GetUserPolicySettings(server domino.String, name domino.String, policyType domino.Boolean, params ...getUserPolicySettingsParam) (notesdocument.NotesDocument, error) {
	paramsStruct := &getUserPolicySettingsParams{}
	paramsOrdered := []interface{}{server, name, policyType}

	for _, p := range params {
		p(paramsStruct)
	}

	if paramsStruct.explicitPolicy != nil {
		paramsOrdered = append(paramsOrdered, *paramsStruct.explicitPolicy)
		if paramsStruct.reserved != nil {
			paramsOrdered = append(paramsOrdered, *paramsStruct.reserved)
		}
	}
	dispatchPtr, err := s.Com().CallObjectMethod("GetUserPolicySettings", paramsOrdered...)
	return notesdocument.New(dispatchPtr), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_HASHPASSWORD_METHOD_SESSION.html */
func (s NotesSession) HashPassword(password domino.String) (domino.String, error) {
	val, err := s.Com().CallMethod("HashPassword", password)
	return helpers.CastValue[domino.String](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_UPDATEPROCESSEDDOC_METHOD.html */
func (s NotesSession) UpdateProcessedDoc(notesDocument notesdocument.NotesDocument) error {
	_, err := s.Com().CallMethod("UpdateProcessedDoc", notesDocument.Com().Dispatch())
	return err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_USEDOUBLEASPOINTER_METHOD_SESSION.html */
func (s NotesSession) UseDoubleAsPointer() error {
	_, err := s.Com().CallMethod("UseDoubleAsPointer")
	return err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_VERIFYPASSWORD_METHOD_SESSION.html */
func (s NotesSession) VerifyPassword(password domino.String, hashedPassword domino.String) (domino.Boolean, error) {
	val, err := s.Com().CallMethod("VerifyPassword", password, hashedPassword)
	return helpers.CastValue[domino.Boolean](val), err
}
