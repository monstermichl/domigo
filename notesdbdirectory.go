/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_NOTESDBDIRECTORY_CLASS.html */
package domigo

import (
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
	return getComProperty[String](d, "Name")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_PARENT_PROPERTY_DBDIRECTORY_COM.html */
func (d NotesDbDirectory) Parent() (NotesSession, error) {
	return getComObjectProperty(d, NewNotesSession, "Parent")
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
	return callComObjectMethod(d, NewNotesDatabase, "CreateDatabase", paramsOrdered...)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GETFIRSTDATABASE_METHOD.html */
func (d NotesDbDirectory) GetFirstDatabase(fileType NotesDbDirectoryFileType) (NotesDatabase, error) {
	return callComObjectMethod(d, NewNotesDatabase, "GetFirstDatabase", fileType)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GETNEXTDATABASE_METHOD.html */
func (d NotesDbDirectory) GetNextDatabase() (NotesDatabase, error) {
	return callComObjectMethod(d, NewNotesDatabase, "GetNextDatabase")
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
	return callComObjectMethod(d, NewNotesDatabase, "OpenDatabase", paramsOrdered...)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_OPENDATABASEBYREPLICAID_METHOD_DBDIRECTORY_COM.html */
func (d NotesDbDirectory) OpenDatabaseByReplicaID(rid String) (NotesDatabase, error) {
	return callComObjectMethod(d, NewNotesDatabase, "OpenDatabaseByReplicaID", rid)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_OPENDATABASEIFMODIFIED_METHOD_DBDIRECTORY_COM.html */
func (d NotesDbDirectory) OpenDatabaseIfModified(dbfile String, notesDateTime NotesDateTime) (NotesDatabase, error) {
	return callComObjectMethod(d, NewNotesDatabase, "OpenDatabaseIfModified", dbfile, notesDateTime)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_OPENMAILDATABASE_METHOD_DBDIRECTORY_COM.html */
func (d NotesDbDirectory) OpenMailDatabase() (NotesDatabase, error) {
	return callComObjectMethod(d, NewNotesDatabase, "OpenMailDatabase")
}
