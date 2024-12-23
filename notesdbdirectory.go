/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_NOTESDBDIRECTORY_CLASS.html */
package domigo

import (
	"github.com/monstermichl/domigo/internal/helpers"

	ole "github.com/go-ole/go-ole"
)

type NotesDbDirectoryFileType = Integer

const (
	NOTESDBDIRECTORY_FILETYPE_DATABASE           NotesDbDirectoryFileType = 1247
	NOTESDBDIRECTORY_FILETYPE_NOTES_DATABASE     NotesDbDirectoryFileType = 1247
	NOTESDBDIRECTORY_FILETYPE_TEMPLATE           NotesDbDirectoryFileType = 1248
	NOTESDBDIRECTORY_FILETYPE_REPLICA_CANDIDATE  NotesDbDirectoryFileType = 1245
	NOTESDBDIRECTORY_FILETYPE_TEMPLATE_CANDIDATE NotesDbDirectoryFileType = 1246
)

type NotesDbDirectory struct {
	NotesStruct
}

func NewNotesDbDirectory(dispatchPtr *ole.IDispatch) NotesDbDirectory {
	return NotesDbDirectory{NewNotesStruct(dispatchPtr)}
}

/* --------------------------------- Properties --------------------------------- */
/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_NAME_PROPERTY_DIRECTORY.html */
func (d NotesDbDirectory) Name() (String, error) {
	val, err := getComProperty(d, "Name")
	return helpers.CastValue[String](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_PARENT_PROPERTY_DBDIRECTORY_COM.html */
func (d NotesDbDirectory) Parent() (NotesSession, error) {
	dispatchPtr, err := getComObjectProperty(d, "Parent")
	return NewNotesSession(dispatchPtr), err
}

/* --------------------------------- Methods ------------------------------------ */
/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_CREATEDATABASE_METHOD_DBDIRECTORY_COM.html */
type notesDbDirectoryCreateDatabaseParams struct {
	open *Boolean
}

type notesDbDirectoryCreateDatabaseParam func(*notesDbDirectoryCreateDatabaseParams)

func WithNotesDbDirectoryCreateDatabaseOpen(open Boolean) notesDbDirectoryCreateDatabaseParam {
	return func(c *notesDbDirectoryCreateDatabaseParams) {
		c.open = &open
	}
}

func (d NotesDbDirectory) CreateDatabase(dbfile String, params ...notesDbDirectoryCreateDatabaseParam) (NotesDatabase, error) {
	paramsStruct := &notesDbDirectoryCreateDatabaseParams{}
	paramsOrdered := []interface{}{dbfile}

	for _, p := range params {
		p(paramsStruct)
	}

	if paramsStruct.open != nil {
		paramsOrdered = append(paramsOrdered, *paramsStruct.open)
	}
	dispatchPtr, err := callComObjectMethod(d, "CreateDatabase", paramsOrdered...)
	return NewNotesDatabase(dispatchPtr), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GETFIRSTDATABASE_METHOD.html */
func (d NotesDbDirectory) GetFirstDatabase(fileType NotesDbDirectoryFileType) (NotesDatabase, error) {
	dispatchPtr, err := callComObjectMethod(d, "GetFirstDatabase", fileType)
	return NewNotesDatabase(dispatchPtr), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GETNEXTDATABASE_METHOD.html */
func (d NotesDbDirectory) GetNextDatabase() (NotesDatabase, error) {
	dispatchPtr, err := callComObjectMethod(d, "GetNextDatabase")
	return NewNotesDatabase(dispatchPtr), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_OPENDATABASE_METHOD_DBDIRECTORY_COM.html */
type notesDbDirectoryOpenDatabaseParams struct {
	open *Boolean
}

type notesDbDirectoryOpenDatabaseParam func(*notesDbDirectoryOpenDatabaseParams)

func WithNotesDbDirectoryOpenDatabaseOpen(open Boolean) notesDbDirectoryOpenDatabaseParam {
	return func(c *notesDbDirectoryOpenDatabaseParams) {
		c.open = &open
	}
}

func (d NotesDbDirectory) OpenDatabase(dbfile String, params ...notesDbDirectoryOpenDatabaseParam) (NotesDatabase, error) {
	paramsStruct := &notesDbDirectoryOpenDatabaseParams{}
	paramsOrdered := []interface{}{dbfile}

	for _, p := range params {
		p(paramsStruct)
	}

	if paramsStruct.open != nil {
		paramsOrdered = append(paramsOrdered, *paramsStruct.open)
	}
	dispatchPtr, err := callComObjectMethod(d, "OpenDatabase", paramsOrdered...)
	return NewNotesDatabase(dispatchPtr), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_OPENDATABASEBYREPLICAID_METHOD_DBDIRECTORY_COM.html */
func (d NotesDbDirectory) OpenDatabaseByReplicaID(rid String) (NotesDatabase, error) {
	dispatchPtr, err := callComObjectMethod(d, "OpenDatabaseByReplicaID", rid)
	return NewNotesDatabase(dispatchPtr), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_OPENDATABASEIFMODIFIED_METHOD_DBDIRECTORY_COM.html */
func (d NotesDbDirectory) OpenDatabaseIfModified(dbfile String, notesDateTime NotesDateTime) (NotesDatabase, error) {
	dispatchPtr, err := callComObjectMethod(d, "OpenDatabaseIfModified", dbfile, notesDateTime.com().Dispatch())
	return NewNotesDatabase(dispatchPtr), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_OPENMAILDATABASE_METHOD_DBDIRECTORY_COM.html */
func (d NotesDbDirectory) OpenMailDatabase() (NotesDatabase, error) {
	dispatchPtr, err := callComObjectMethod(d, "OpenMailDatabase")
	return NewNotesDatabase(dispatchPtr), err
}
