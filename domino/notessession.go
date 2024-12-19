/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_NOTESSESSION_CLASS.html */
package domigo

import (
	"errors"

	"github.com/monstermichl/domigo/domino/com"
	"github.com/monstermichl/domigo/helpers"

	ole "github.com/go-ole/go-ole"
)

type NotesSessionPolicy Long

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

func NewNotesSession(dispatchPtr *ole.IDispatch) NotesSession {
	return NotesSession{NewNotesStruct(dispatchPtr)}
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
	session = NewNotesSession(com.Dispatch())
	_, err = session.com().CallObjectMethod("Initialize", params...)

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
	paramsOrdered := []interface{}{name}

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
	return com.GetObjectArrayProperty(s.com(), NewNotesDatabase, "AddressBooks")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_COMMONUSERNAME_PROPERTY.html */
func (s NotesSession) CommonUserName() (String, error) {
	val, err := s.com().GetProperty("CommonUserName")
	return helpers.CastValue[String](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_CONVERTMIME_PROPERTY.html */
func (s NotesSession) ConvertMIME() (Boolean, error) {
	val, err := s.com().GetProperty("ConvertMIME")
	return helpers.CastValue[Boolean](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_CONVERTMIME_PROPERTY.html */
func (s NotesSession) SetConvertMIME(v Boolean) error {
	return s.com().PutProperty("ConvertMIME", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_CURRENTAGENT_PROPERTY.html */
func (s NotesSession) CurrentAgent() (NotesAgent, error) {
	dispatchPtr, err := s.com().GetObjectProperty("CurrentAgent")
	return NewNotesAgent(dispatchPtr), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_CURRENTDATABASE_PROPERTY.html */
func (s NotesSession) CurrentDatabase() (NotesDatabase, error) {
	dispatchPtr, err := s.com().GetObjectProperty("CurrentDatabase")
	return NewNotesDatabase(dispatchPtr), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_DOCUMENTCONTEXT_PROPERTY.html */
func (s NotesSession) DocumentContext() (NotesDocument, error) {
	dispatchPtr, err := s.com().GetObjectProperty("DocumentContext")
	return NewNotesDocument(dispatchPtr), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_EFFECTIVEUSERNAME_PROPERTY.html */
func (s NotesSession) EffectiveUserName() (String, error) {
	val, err := s.com().GetProperty("EffectiveUserName")
	return helpers.CastValue[String](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_HTTPURL_PROPERTY_SESSION_COM.html */
func (s NotesSession) HttpURL() (String, error) {
	val, err := s.com().GetProperty("HttpURL")
	return helpers.CastValue[String](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_INTERNATIONAL_PROPERTY.html */
func (s NotesSession) International() (Integer, error) {
	val, err := s.com().GetProperty("International")
	return helpers.CastValue[Integer](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISONSERVER_PROPERTY.html */
func (s NotesSession) IsOnServer() (Boolean, error) {
	val, err := s.com().GetProperty("IsOnServer")
	return helpers.CastValue[Boolean](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_LASTEXITSTATUS_PROPERTY.html */
func (s NotesSession) LastExitStatus() (Long, error) {
	val, err := s.com().GetProperty("LastExitStatus")
	return helpers.CastValue[Long](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_LASTRUN_PROPERTY_SESSION.html */
func (s NotesSession) LastRun() (Time, error) {
	val, err := s.com().GetProperty("LastRun")
	return helpers.CastValue[Time](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_NOTESBUILDVERSION_PROPERTY_4714_ABOUT.html */
func (s NotesSession) NotesBuildVersion() (Long, error) {
	val, err := s.com().GetProperty("NotesBuildVersion")
	return helpers.CastValue[Long](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_NOTESURL_PROPERTY_SESSION_COM.html */
func (s NotesSession) NotesURL() (String, error) {
	val, err := s.com().GetProperty("NotesURL")
	return helpers.CastValue[String](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_NOTESVERSION_PROPERTY.html */
func (s NotesSession) NotesVersion() (String, error) {
	val, err := s.com().GetProperty("NotesVersion")
	return helpers.CastValue[String](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ORGDIRECTORYPATH_PROPERTY_SESSION.html */
func (s NotesSession) OrgDirectoryPath() (String, error) {
	val, err := s.com().GetProperty("OrgDirectoryPath")
	return helpers.CastValue[String](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_PLATFORM_PROPERTY.html */
func (s NotesSession) Platform() (String, error) {
	val, err := s.com().GetProperty("Platform")
	return helpers.CastValue[String](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SAVEDDATA_PROPERTY.html */
func (s NotesSession) SavedData() (NotesDocument, error) {
	dispatchPtr, err := s.com().GetObjectProperty("SavedData")
	return NewNotesDocument(dispatchPtr), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SERVERNAME_PROPERTY_SESSION_COM.html */
func (s NotesSession) ServerName() (String, error) {
	val, err := s.com().GetProperty("ServerName")
	return helpers.CastValue[String](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_URLDATABASE_PROPERTY_SESSION_COM.html */
func (s NotesSession) URLDatabase() (NotesDatabase, error) {
	dispatchPtr, err := s.com().GetObjectProperty("URLDatabase")
	return NewNotesDatabase(dispatchPtr), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_USERGROUPNAMELIST_PROPERTY_SESSION.html */
func (s NotesSession) UserGroupNameList() ([]NotesName, error) {
	return com.GetObjectArrayProperty(s.com(), NewNotesName, "UserGroupNameList")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_USERNAME_PROPERTY.html */
func (s NotesSession) UserName() (String, error) {
	val, err := s.com().GetProperty("UserName")
	return helpers.CastValue[String](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_USERNAMELIST_PROPERTY_7482.html */
func (s NotesSession) UserNameList() ([]NotesName, error) {
	return com.GetObjectArrayProperty(s.com(), NewNotesName, "UserNameList")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_USERNAMEOBJECT_PROPERTY_SESSION_COM.html */
func (s NotesSession) UserNameObject() (NotesName, error) {
	dispatchPtr, err := s.com().GetObjectProperty("UserNameObject")
	return NewNotesName(dispatchPtr), err
}

/* --------------------------------- Methods ------------------------------------ */
/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_CREATEADMINISTRATIONPROCESS_METHOD_SESSION.html */
func (s NotesSession) CreateAdministrationProcess(server String) (NotesAdministrationProcess, error) {
	dispatchPtr, err := s.com().CallObjectMethod("CreateAdministrationProcess", server)
	return NewNotesAdministrationProcess(dispatchPtr), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_CREATECOLOROBJECT_METHOD_SESSION.html */
func (s NotesSession) CreateColorObject() (NotesColorObject, error) {
	dispatchPtr, err := s.com().CallObjectMethod("CreateColorObject")
	return NewNotesColorObject(dispatchPtr), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_CREATEDATERANGE_METHOD.html */
func (s NotesSession) CreateDateRange() (NotesDateRange, error) {
	dispatchPtr, err := s.com().CallObjectMethod("CreateDateRange")
	return NewNotesDateRange(dispatchPtr), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_CREATEDATETIME_METHOD.html */
func (s NotesSession) CreateDateTime(dateTime String) (NotesDateTime, error) {
	dispatchPtr, err := s.com().CallObjectMethod("CreateDateTime", dateTime)
	return NewNotesDateTime(dispatchPtr), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_CREATEDXLEXPORTER_METHOD_SESSION.html */
/* TODO: Implement this function correctly (see documentation). */
// type notesSessionCreateDXLExporterParams struct {
//     input *Object
//     output *Object
// }

// type notesSessionCreateDXLExporterParam func(*notesSessionCreateDXLExporterParams)

// func WithNotesSessionCreateDXLExporterInput(input Object) notesSessionCreateDXLExporterParam {
//     return func(c *notesSessionCreateDXLExporterParams) {
//         c.input = &input
//     }
// }

// func WithNotesSessionCreateDXLExporterOutput(output Object) notesSessionCreateDXLExporterParam {
//     return func(c *notesSessionCreateDXLExporterParams) {
//         c.output = &output
//     }
// }

// func (s NotesSession) CreateDXLExporter(params ...notesSessionCreateDXLExporterParam) (Object, error) {
//     paramsStruct := &notesSessionCreateDXLExporterParams{}
//     paramsOrdered := []interface{}{}

//     for _, p := range params {
//         p(paramsStruct)
//     }

//     if paramsStruct.input != nil {
//         paramsOrdered = append(paramsOrdered, *paramsStruct.input)
//             /* TODO: output probably needs to be handled differently with COM. */
//             if paramsStruct.output != nil {
//                 paramsOrdered = append(paramsOrdered, *paramsStruct.output)
//             }
//     }
//     dispatchPtr, err := s.com().CallObjectMethod("CreateDXLExporter", paramsOrdered...)
//     return New(dispatchPtr), err
// }

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_CREATEDXLIMPORTER_METHOD_SESSION.html */
/* TODO: Implement this function correctly (see documentation). */
// type notesSessionCreateDXLImporterParams struct {
//     input *String
//     output *Object
// }

// type notesSessionCreateDXLImporterParam func(*notesSessionCreateDXLImporterParams)

// func WithNotesSessionCreateDXLImporterInput(input String) notesSessionCreateDXLImporterParam {
//     return func(c *notesSessionCreateDXLImporterParams) {
//         c.input = &input
//     }
// }

// func WithNotesSessionCreateDXLImporterOutput(output Object) notesSessionCreateDXLImporterParam {
//     return func(c *notesSessionCreateDXLImporterParams) {
//         c.output = &output
//     }
// }

// func (s NotesSession) CreateDXLImporter(params ...notesSessionCreateDXLImporterParam) (Object, error) {
//     paramsStruct := &notesSessionCreateDXLImporterParams{}
//     paramsOrdered := []interface{}{}

//     for _, p := range params {
//         p(paramsStruct)
//     }

//     if paramsStruct.input != nil {
//         paramsOrdered = append(paramsOrdered, *paramsStruct.input)
//             /* TODO: output probably needs to be handled differently with COM. */
//             if paramsStruct.output != nil {
//                 paramsOrdered = append(paramsOrdered, *paramsStruct.output)
//             }
//     }
//     dispatchPtr, err := s.com().CallObjectMethod("CreateDXLImporter", paramsOrdered...)
//     return New(dispatchPtr), err
// }

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_CREATELOG_METHOD.html */
func (s NotesSession) CreateLog(programName String) (NotesLog, error) {
	dispatchPtr, err := s.com().CallObjectMethod("CreateLog", programName)
	return NewNotesLog(dispatchPtr), err
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
	paramsOrdered := []interface{}{name}

	for _, p := range params {
		p(paramsStruct)
	}

	if paramsStruct.language != nil {
		paramsOrdered = append(paramsOrdered, *paramsStruct.language)
	}
	dispatchPtr, err := s.com().CallObjectMethod("CreateName", paramsOrdered...)
	return NewNotesName(dispatchPtr), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_CREATENEWSLETTER_METHOD.html */
func (s NotesSession) CreateNewsletter(notesDocumentCollection NotesDocumentCollection) (NotesNewsletter, error) {
	dispatchPtr, err := s.com().CallObjectMethod("CreateNewsletter", notesDocumentCollection.com().Dispatch())
	return NewNotesNewsletter(dispatchPtr), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_CREATEREGISTRATION_METHOD_SESSION_COM.html */
func (s NotesSession) CreateRegistration() (NotesRegistration, error) {
	dispatchPtr, err := s.com().CallObjectMethod("CreateRegistration")
	return NewNotesRegistration(dispatchPtr), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_CREATERICHTEXTPARAGRAPHSTYLE_METHOD_8901_ABOUT.html */
func (s NotesSession) CreateRichTextParagraphStyle() (NotesRichTextParagraphStyle, error) {
	dispatchPtr, err := s.com().CallObjectMethod("CreateRichTextParagraphStyle")
	return NewNotesRichTextParagraphStyle(dispatchPtr), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_CREATERICHTEXTSTYLE_METHOD.html */
func (s NotesSession) CreateRichTextStyle() (NotesRichTextStyle, error) {
	dispatchPtr, err := s.com().CallObjectMethod("CreateRichTextStyle")
	return NewNotesRichTextStyle(dispatchPtr), err
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
//     paramsOrdered := []interface{}{}

//     for _, p := range params {
//         p(paramsStruct)
//     }

//     if paramsStruct.input != nil {
//         paramsOrdered = append(paramsOrdered, *paramsStruct.input)
//             if paramsStruct.output != nil {
//                 paramsOrdered = append(paramsOrdered, *paramsStruct.output)
//             }
//     }
//     dispatchPtr, err := s.com().CallObjectMethod("CreateSAXParser", paramsOrdered...)
//     return New(dispatchPtr), err
// }

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_CREATESTREAM_METHOD_SESSION.html */
func (s NotesSession) CreateStream() (NotesStream, error) {
	dispatchPtr, err := s.com().CallObjectMethod("CreateStream")
	return NewNotesStream(dispatchPtr), err
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
//     paramsOrdered := []interface{}{}

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
//     dispatchPtr, err := s.com().CallObjectMethod("CreateXSLTransformer", paramsOrdered...)
//     return New(dispatchPtr), err
// }

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_EVALUATE_METHOD_SESSION_COM.html */
func (s NotesSession) Evaluate(formula String, doc NotesDocument) (Variant, error) {
	val, err := s.com().CallMethod("Evaluate", formula, doc.com().Dispatch())
	return helpers.CastValue[Variant](val), err
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
	paramsOrdered := []interface{}{Start.com().Dispatch(), End.com().Dispatch(), Site, Type}

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
	vals, err := s.com().CallArrayMethod("FreeResourceSearch", paramsOrdered...)
	return helpers.CastSlice[String](vals), err
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
	paramsOrdered := []interface{}{window.com().Dispatch(), duration, names}

	for _, p := range params {
		p(paramsStruct)
	}

	if paramsStruct.firstfit != nil {
		paramsOrdered = append(paramsOrdered, *paramsStruct.firstfit)
	}
	dispatchPtr, err := s.com().CallObjectMethod("FreeTimeSearch", paramsOrdered...)
	return NewNotesDateRange(dispatchPtr), err
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
	paramsOrdered := []interface{}{server, dbfile}

	for _, p := range params {
		p(paramsStruct)
	}

	if paramsStruct.createonfail != nil {
		paramsOrdered = append(paramsOrdered, *paramsStruct.createonfail)
	}
	dispatchPtr, err := s.com().CallObjectMethod("GetDatabase", paramsOrdered...)
	return NewNotesDatabase(dispatchPtr), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GETCALENDAR_METHOD.html */
func (s NotesSession) GetCalendar(mailDatabase NotesDatabase) (NotesCalendar, error) {
	dispatchPtr, err := s.com().CallObjectMethod("GetCalendar", mailDatabase.com().Dispatch())
	return NewNotesCalendar(dispatchPtr), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GETDBDIRECTORY_METHOD.html */
func (s NotesSession) GetDbDirectory(serverName String) (NotesDbDirectory, error) {
	dispatchPtr, err := s.com().CallObjectMethod("GetDbDirectory", serverName)
	return NewNotesDbDirectory(dispatchPtr), err
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
	paramsOrdered := []interface{}{}

	for _, p := range params {
		p(paramsStruct)
	}

	if paramsStruct.serverName != nil {
		paramsOrdered = append(paramsOrdered, *paramsStruct.serverName)
	}
	dispatchPtr, err := s.com().CallObjectMethod("GetDirectory", paramsOrdered...)
	return NewNotesDirectory(dispatchPtr), err
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
	paramsOrdered := []interface{}{name}

	for _, p := range params {
		p(paramsStruct)
	}

	if paramsStruct.system != nil {
		paramsOrdered = append(paramsOrdered, *paramsStruct.system)
	}
	val, err := s.com().CallMethod("GetEnvironmentString", paramsOrdered...)
	return helpers.CastValue[String](val), err
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
	paramsOrdered := []interface{}{name}

	for _, p := range params {
		p(paramsStruct)
	}

	if paramsStruct.system != nil {
		paramsOrdered = append(paramsOrdered, *paramsStruct.system)
	}
	return s.com().CallMethod("GetEnvironmentValue", paramsOrdered...)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GETPROPERTYBROKER_METHOD.html */
func (s NotesSession) GetPropertyBroker() (NotesPropertyBroker, error) {
	dispatchPtr, err := s.com().CallObjectMethod("GetPropertyBroker")
	return NewNotesPropertyBroker(dispatchPtr), err
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
	dispatchPtr, err := s.com().CallObjectMethod("GetUserPolicySettings", paramsOrdered...)
	return NewNotesDocument(dispatchPtr), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_HASHPASSWORD_METHOD_SESSION.html */
func (s NotesSession) HashPassword(password String) (String, error) {
	val, err := s.com().CallMethod("HashPassword", password)
	return helpers.CastValue[String](val), err
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
	paramsOrdered := []interface{}{servername, username, password}

	for _, p := range params {
		p(paramsStruct)
	}

	if paramsStruct.downloadcount != nil {
		paramsOrdered = append(paramsOrdered, *paramsStruct.downloadcount)
	}
	_, err := s.com().CallMethod("ResetUserPassword", paramsOrdered...)
	return err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_RESOLVE_METHOD_SESSION_COM.html */
func (s NotesSession) Resolve(url String) error {
	_, err := s.com().CallMethod("Resolve", url)
	return err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SENDCONSOLECOMMAND_METHOD_SESSION.html */
func (s NotesSession) SendConsoleCommand(serverName String, consoleCommand String) (String, error) {
	val, err := s.com().CallMethod("SendConsoleCommand", serverName, consoleCommand)
	return helpers.CastValue[String](val), err
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

func (s NotesSession) SetEnvironmentVar(name String, valueV Variant, params ...notesSessionSetEnvironmentVarParam) error {
	paramsStruct := &notesSessionSetEnvironmentVarParams{}
	paramsOrdered := []interface{}{name, valueV}

	for _, p := range params {
		p(paramsStruct)
	}

	if paramsStruct.issystemvar != nil {
		paramsOrdered = append(paramsOrdered, *paramsStruct.issystemvar)
	}
	_, err := s.com().CallMethod("SetEnvironmentVar", paramsOrdered...)
	return err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_UPDATEPROCESSEDDOC_METHOD.html */
func (s NotesSession) UpdateProcessedDoc(notesDocument NotesDocument) error {
	_, err := s.com().CallMethod("UpdateProcessedDoc", notesDocument.com().Dispatch())
	return err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_USEDOUBLEASPOINTER_METHOD_SESSION.html */
func (s NotesSession) UseDoubleAsPointer() error {
	_, err := s.com().CallMethod("UseDoubleAsPointer")
	return err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_VERIFYPASSWORD_METHOD_SESSION.html */
func (s NotesSession) VerifyPassword(password String, hashedPassword String) (Boolean, error) {
	val, err := s.com().CallMethod("VerifyPassword", password, hashedPassword)
	return helpers.CastValue[Boolean](val), err
}
