/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_NOTESDBDIRECTORY_CLASS.html */
package notesdbdirectory

import (
	"domigo/domino"
	"domigo/domino/notesdatabase"
	"domigo/domino/notesdatetime"
	"domigo/helpers"

	ole "github.com/go-ole/go-ole"
)

type FileType = domino.Integer

const (
	FILETYPE_DATABASE           FileType = 1247
	FILETYPE_NOTES_DATABASE     FileType = 1247
	FILETYPE_TEMPLATE           FileType = 1248
	FILETYPE_REPLICA_CANDIDATE  FileType = 1245
	FILETYPE_TEMPLATE_CANDIDATE FileType = 1246
)

type NotesDbDirectory struct {
	domino.NotesStruct
}

func New(dispatchPtr *ole.IDispatch) NotesDbDirectory {
	return NotesDbDirectory{domino.NewNotesStruct(dispatchPtr)}
}

/* --------------------------------- Properties --------------------------------- */
/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_NAME_PROPERTY_DIRECTORY.html */
func (d NotesDbDirectory) Name() (domino.String, error) {
	val, err := d.Com().GetProperty("Name")
	return helpers.CastValue[domino.String](val), err
}

/* --------------------------------- Methods ------------------------------------ */
/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_CREATEDATABASE_METHOD_DBDIRECTORY_COM.html */
type createDatabaseParams struct {
	open *domino.Boolean
}

type createDatabaseParam func(*createDatabaseParams)

func WithCreateDatabaseOpen(open domino.Boolean) createDatabaseParam {
	return func(c *createDatabaseParams) {
		c.open = &open
	}
}

func (d NotesDbDirectory) CreateDatabase(dbfile domino.String, params ...createDatabaseParam) (notesdatabase.NotesDatabase, error) {
	paramsStruct := &createDatabaseParams{}
	paramsOrdered := []interface{}{dbfile}

	for _, p := range params {
		p(paramsStruct)
	}

	if paramsStruct.open != nil {
		paramsOrdered = append(paramsOrdered, *paramsStruct.open)
	}
	dispatchPtr, err := d.Com().CallObjectMethod("CreateDatabase", paramsOrdered...)
	return notesdatabase.New(dispatchPtr), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GETFIRSTDATABASE_METHOD.html */
func (d NotesDbDirectory) GetFirstDatabase(fileType FileType) (notesdatabase.NotesDatabase, error) {
	dispatchPtr, err := d.Com().CallObjectMethod("GetFirstDatabase", fileType)
	return notesdatabase.New(dispatchPtr), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GETNEXTDATABASE_METHOD.html */
func (d NotesDbDirectory) GetNextDatabase() (notesdatabase.NotesDatabase, error) {
	dispatchPtr, err := d.Com().CallObjectMethod("GetNextDatabase")
	return notesdatabase.New(dispatchPtr), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_OPENDATABASE_METHOD_DBDIRECTORY_COM.html */
type openDatabaseParams struct {
	open *domino.Boolean
}

type openDatabaseParam func(*openDatabaseParams)

func WithOpenDatabaseOpen(open domino.Boolean) openDatabaseParam {
	return func(c *openDatabaseParams) {
		c.open = &open
	}
}

func (d NotesDbDirectory) OpenDatabase(dbfile domino.String, params ...openDatabaseParam) (notesdatabase.NotesDatabase, error) {
	paramsStruct := &openDatabaseParams{}
	paramsOrdered := []interface{}{dbfile}

	for _, p := range params {
		p(paramsStruct)
	}

	if paramsStruct.open != nil {
		paramsOrdered = append(paramsOrdered, *paramsStruct.open)
	}
	dispatchPtr, err := d.Com().CallObjectMethod("OpenDatabase", paramsOrdered...)
	return notesdatabase.New(dispatchPtr), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_OPENDATABASEBYREPLICAID_METHOD_DBDIRECTORY_COM.html */
func (d NotesDbDirectory) OpenDatabaseByReplicaID(rid domino.String) (notesdatabase.NotesDatabase, error) {
	dispatchPtr, err := d.Com().CallObjectMethod("OpenDatabaseByReplicaID", rid)
	return notesdatabase.New(dispatchPtr), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_OPENDATABASEIFMODIFIED_METHOD_DBDIRECTORY_COM.html */
func (d NotesDbDirectory) OpenDatabaseIfModified(dbfile domino.String, notesDateTime notesdatetime.NotesDateTime) (notesdatabase.NotesDatabase, error) {
	dispatchPtr, err := d.Com().CallObjectMethod("OpenDatabaseIfModified", dbfile, notesDateTime.Com().Dispatch())
	return notesdatabase.New(dispatchPtr), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_OPENMAILDATABASE_METHOD_DBDIRECTORY_COM.html */
func (d NotesDbDirectory) OpenMailDatabase() (notesdatabase.NotesDatabase, error) {
	dispatchPtr, err := d.Com().CallObjectMethod("OpenMailDatabase")
	return notesdatabase.New(dispatchPtr), err
}
