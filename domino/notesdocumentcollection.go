/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_NOTESDOCUMENTCOLLECTION_CLASS.html */
package domigo

import (
	"github.com/monstermichl/domigo/helpers"

	ole "github.com/go-ole/go-ole"
)

type NotesDocumentCollection struct {
	NotesStruct
}

func NewNotesDocumentCollection(dispatchPtr *ole.IDispatch) NotesDocumentCollection {
	return NotesDocumentCollection{NewNotesStruct(dispatchPtr)}
}

func (d NotesDocumentCollection) checkCombinableTypes(val any) error {
	return helpers.CheckTypeNames(val, []string{"string", "NotesDocument", "NotesDocumentCollection", "NotesViewEntry", "NotesViewEntryCollection"})
}

/* --------------------------------- Properties --------------------------------- */
/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_COUNT_PROPERTY.html */
func (d NotesDocumentCollection) Count() (Long, error) {
	val, err := d.Com().GetProperty("Count")
	return helpers.CastValue[Long](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISSORTED_PROPERTY_COLLECTION.html */
func (d NotesDocumentCollection) IsSorted() (Boolean, error) {
	val, err := d.Com().GetProperty("IsSorted")
	return helpers.CastValue[Boolean](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_PARENT_PROPERTY_COLLECTION.html */
func (d NotesDocumentCollection) Parent() (NotesDatabase, error) {
	dispatchPtr, err := d.Com().GetObjectProperty("Parent")
	return NewNotesDatabase(dispatchPtr), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_QUERY_PROPERTY_COLLECTION.html */
func (d NotesDocumentCollection) Query() (String, error) {
	val, err := d.Com().GetProperty("Query")
	return helpers.CastValue[String](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_UNTILTIME_PROPERTY_COLLECTION.html */
func (d NotesDocumentCollection) UntilTime() (NotesDateTime, error) {
	dispatchPtr, err := d.Com().GetObjectProperty("UntilTime")
	return NewNotesDateTime(dispatchPtr), err
}

/* --------------------------------- Methods ------------------------------------ */
/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ADDDOCUMENT_METHOD_7116_ABOUT.html */
type notesDocumentCollectionAddDocumentParams struct {
	checkDups *Boolean
}

type notesDocumentCollectionAddDocumentParam func(*notesDocumentCollectionAddDocumentParams)

func WithNotesDocumentCollectionAddDocumentCheckDups(checkDups Boolean) notesDocumentCollectionAddDocumentParam {
	return func(c *notesDocumentCollectionAddDocumentParams) {
		c.checkDups = &checkDups
	}
}

func (d NotesDocumentCollection) AddDocument(document NotesDocument, params ...notesDocumentCollectionAddDocumentParam) error {
	paramsStruct := &notesDocumentCollectionAddDocumentParams{}
	paramsOrdered := []interface{}{document.Com().Dispatch()}

	for _, p := range params {
		p(paramsStruct)
	}

	if paramsStruct.checkDups != nil {
		paramsOrdered = append(paramsOrdered, *paramsStruct.checkDups)
	}
	_, err := d.Com().CallMethod("AddDocument", paramsOrdered...)
	return err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_CLONE_METHOD_COLLECTION.html */
func (d NotesDocumentCollection) Clone() error {
	_, err := d.Com().CallMethod("Clone")
	return err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_CONTAINS_METHOD_COLLECTION.html */
/* TODO: Handle different input types (see documentation). */
func (d NotesDocumentCollection) Contains(inputNotes any) (Boolean, error) {
	err := d.checkCombinableTypes(inputNotes)

	if err != nil {
		return false, err
	}
	val, err := d.Com().CallMethod("Contains", inputNotes)
	return helpers.CastValue[Boolean](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_DELETEDOCUMENT_METHOD_7984_ABOUT.html */
func (d NotesDocumentCollection) DeleteDocument(document NotesDocument) error {
	_, err := d.Com().CallMethod("DeleteDocument", document.Com().Dispatch())
	return err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_FTSEARCH_METHOD_COLLECTION.html */
func (d NotesDocumentCollection) FTSearch(query String, maxDocs Integer) error {
	_, err := d.Com().CallMethod("FTSearch", query, maxDocs)
	return err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GETDOCUMENT_METHOD_DOCCOLL.html */
func (d NotesDocumentCollection) GetDocument(document NotesDocument) (NotesDocument, error) {
	dispatchPtr, err := d.Com().CallObjectMethod("GetDocument", document.Com().Dispatch())
	return NewNotesDocument(dispatchPtr), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GETFIRSTDOCUMENT_METHOD_COLLECTION.html */
func (d NotesDocumentCollection) GetFirstDocument() (NotesDocument, error) {
	dispatchPtr, err := d.Com().CallObjectMethod("GetFirstDocument")
	return NewNotesDocument(dispatchPtr), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GETLASTDOCUMENT_METHOD_COLLECTION.html */
func (d NotesDocumentCollection) GetLastDocument() (NotesDocument, error) {
	dispatchPtr, err := d.Com().CallObjectMethod("GetLastDocument")
	return NewNotesDocument(dispatchPtr), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GETNEXTDOCUMENT_METHOD_COLLECTION.html */
func (d NotesDocumentCollection) GetNextDocument(document NotesDocument) (NotesDocument, error) {
	dispatchPtr, err := d.Com().CallObjectMethod("GetNextDocument", document.Com().Dispatch())
	return NewNotesDocument(dispatchPtr), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GETNTHDOCUMENT_METHOD_COLLECTION.html */
func (d NotesDocumentCollection) GetNthDocument(n Long) (NotesDocument, error) {
	dispatchPtr, err := d.Com().CallObjectMethod("GetNthDocument", n)
	return NewNotesDocument(dispatchPtr), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GETPREVDOCUMENT_METHOD_COLLECTION.html */
func (d NotesDocumentCollection) GetPrevDocument(document NotesDocument) (NotesDocument, error) {
	dispatchPtr, err := d.Com().CallObjectMethod("GetPrevDocument", document.Com().Dispatch())
	return NewNotesDocument(dispatchPtr), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_INTERSECT_METHOD_COLLECTION.html */
/* TODO: Handle different input types (see documentation). */
func (d NotesDocumentCollection) Intersect(inputNotes any) error {
	err := d.checkCombinableTypes(inputNotes)

	if err != nil {
		return err
	}
	_, err = d.Com().CallMethod("Intersect", inputNotes)
	return err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_MARKALLREAD_DOCCOLLECTION.html */
type notesDocumentCollectionMarkAllReadParams struct {
	username *String
}

type notesDocumentCollectionMarkAllReadParam func(*notesDocumentCollectionMarkAllReadParams)

func WithNotesDocumentCollectionMarkAllReadUsername(username String) notesDocumentCollectionMarkAllReadParam {
	return func(c *notesDocumentCollectionMarkAllReadParams) {
		c.username = &username
	}
}

func (d NotesDocumentCollection) MarkAllRead(params ...notesDocumentCollectionMarkAllReadParam) error {
	paramsStruct := &notesDocumentCollectionMarkAllReadParams{}
	paramsOrdered := []interface{}{}

	for _, p := range params {
		p(paramsStruct)
	}

	if paramsStruct.username != nil {
		paramsOrdered = append(paramsOrdered, *paramsStruct.username)
	}
	_, err := d.Com().CallMethod("MarkAllRead", paramsOrdered...)
	return err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_MARKALLUNREAD_DOCCOLLECTION.html */
type notesDocumentCollectionMarkAllUnreadParams struct {
	username *String
}

type notesDocumentCollectionMarkAllUnreadParam func(*notesDocumentCollectionMarkAllUnreadParams)

func WithNotesDocumentCollectionMarkAllUnreadUsername(username String) notesDocumentCollectionMarkAllUnreadParam {
	return func(c *notesDocumentCollectionMarkAllUnreadParams) {
		c.username = &username
	}
}

func (d NotesDocumentCollection) MarkAllUnread(params ...notesDocumentCollectionMarkAllUnreadParam) error {
	paramsStruct := &notesDocumentCollectionMarkAllUnreadParams{}
	paramsOrdered := []interface{}{}

	for _, p := range params {
		p(paramsStruct)
	}

	if paramsStruct.username != nil {
		paramsOrdered = append(paramsOrdered, *paramsStruct.username)
	}
	_, err := d.Com().CallMethod("MarkAllUnread", paramsOrdered...)
	return err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_MERGE_METHOD_COLLECTION.html */
/* TODO: Handle different input types (see documentation). */
func (d NotesDocumentCollection) Merge(inputNotes any) error {
	err := d.checkCombinableTypes(inputNotes)

	if err != nil {
		return err
	}
	_, err = d.Com().CallMethod("Merge", inputNotes)
	return err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_PUTALLINFOLDER_METHOD.html */
type notesDocumentCollectionPutAllInFolderParams struct {
	createonfail *Boolean
}

type notesDocumentCollectionPutAllInFolderParam func(*notesDocumentCollectionPutAllInFolderParams)

func WithNotesDocumentCollectionPutAllInFolderCreateonfail(createonfail Boolean) notesDocumentCollectionPutAllInFolderParam {
	return func(c *notesDocumentCollectionPutAllInFolderParams) {
		c.createonfail = &createonfail
	}
}

func (d NotesDocumentCollection) PutAllInFolder(folderName String, params ...notesDocumentCollectionPutAllInFolderParam) error {
	paramsStruct := &notesDocumentCollectionPutAllInFolderParams{}
	paramsOrdered := []interface{}{folderName}

	for _, p := range params {
		p(paramsStruct)
	}

	if paramsStruct.createonfail != nil {
		paramsOrdered = append(paramsOrdered, *paramsStruct.createonfail)
	}
	_, err := d.Com().CallMethod("PutAllInFolder", paramsOrdered...)
	return err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_REMOVEALL_METHOD.html */
func (d NotesDocumentCollection) RemoveAll(force Boolean) error {
	_, err := d.Com().CallMethod("RemoveAll", force)
	return err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_REMOVEALLFROMFOLDER_METHOD.html */
func (d NotesDocumentCollection) RemoveAllFromFolder(folderName String) error {
	_, err := d.Com().CallMethod("RemoveAllFromFolder", folderName)
	return err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_STAMPALL_METHOD.html */
func (d NotesDocumentCollection) StampAll(itemname String, value any) error {
	_, err := d.Com().CallMethod("StampAll", itemname, value)
	return err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_STAMPALLMULTI_METHOD_DOCCOL.html */
func (d NotesDocumentCollection) StampAllMulti(document NotesDocument) error {
	_, err := d.Com().CallMethod("StampAllMulti", document.Com().Dispatch())
	return err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SUBTRACT_METHOD_COLLECTION.html */
/* TODO: Handle different input types (see documentation). */
func (d NotesDocumentCollection) Subtract(inputNotes any) error {
	err := d.checkCombinableTypes(inputNotes)

	if err != nil {
		return err
	}
	_, err = d.Com().CallMethod("Subtract", inputNotes)
	return err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_UPDATEALL_METHOD.html */
func (d NotesDocumentCollection) UpdateAll() error {
	_, err := d.Com().CallMethod("UpdateAll")
	return err
}
