/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_NOTESDOCUMENTCOLLECTION_CLASS.html */
package domigo

import (
	"github.com/monstermichl/domigo/internal/helpers"

	ole "github.com/go-ole/go-ole"
)

type NotesDocumentCollection struct {
	NotesStruct
}

func newNotesDocumentCollection(dispatchPtr *ole.IDispatch) NotesDocumentCollection {
	return NotesDocumentCollection{newNotesStruct(dispatchPtr)}
}

func (d NotesDocumentCollection) checkCombinableTypes(val any) error {
	return helpers.CheckTypeNames(val, []string{"string", "NotesDocument", "NotesDocumentCollection", "NotesViewEntry", "NotesViewEntryCollection"})
}

/* --------------------------------- Properties --------------------------------- */
/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_COUNT_PROPERTY.html */
func (d NotesDocumentCollection) Count() (Long, error) {
	return getComProperty[Long](d, "Count")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISSORTED_PROPERTY_COLLECTION.html */
func (d NotesDocumentCollection) IsSorted() (Boolean, error) {
	return getComProperty[Boolean](d, "IsSorted")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_PARENT_PROPERTY_COLLECTION.html */
func (d NotesDocumentCollection) Parent() (NotesDatabase, error) {
	return getComObjectProperty(d, newNotesDatabase, "Parent")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_QUERY_PROPERTY_COLLECTION.html */
func (d NotesDocumentCollection) Query() (String, error) {
	return getComProperty[String](d, "Query")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_UNTILTIME_PROPERTY_COLLECTION.html */
func (d NotesDocumentCollection) UntilTime() (NotesDateTime, error) {
	return getComObjectProperty(d, newNotesDateTime, "UntilTime")
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
	paramsOrdered := []interface{}{document}

	for _, p := range params {
		p(paramsStruct)
	}

	if paramsStruct.checkDups != nil {
		paramsOrdered = append(paramsOrdered, *paramsStruct.checkDups)
	}
	return callComVoidMethod(d, "AddDocument", paramsOrdered...)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_CLONE_METHOD_COLLECTION.html */
func (d NotesDocumentCollection) Clone() error {
	return callComVoidMethod(d, "Clone")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_CONTAINS_METHOD_COLLECTION.html */
/* TODO: Handle different input types (see documentation). */
func (d NotesDocumentCollection) Contains(inputNotes any) (Boolean, error) {
	err := d.checkCombinableTypes(inputNotes)

	if err != nil {
		return false, err
	}
	return callComMethod[Boolean](d, "Contains", inputNotes)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_DELETEDOCUMENT_METHOD_7984_ABOUT.html */
func (d NotesDocumentCollection) DeleteDocument(document NotesDocument) error {
	return callComVoidMethod(d, "DeleteDocument", document)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_FTSEARCH_METHOD_COLLECTION.html */
func (d NotesDocumentCollection) FTSearch(query String, maxDocs Integer) error {
	return callComVoidMethod(d, "FTSearch", query, maxDocs)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GETDOCUMENT_METHOD_DOCCOLL.html */
func (d NotesDocumentCollection) GetDocument(document NotesDocument) (NotesDocument, error) {
	return callComObjectMethod(d, newNotesDocument, "GetDocument", document)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GETFIRSTDOCUMENT_METHOD_COLLECTION.html */
func (d NotesDocumentCollection) GetFirstDocument() (NotesDocument, error) {
	return callComObjectMethod(d, newNotesDocument, "GetFirstDocument")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GETLASTDOCUMENT_METHOD_COLLECTION.html */
func (d NotesDocumentCollection) GetLastDocument() (NotesDocument, error) {
	return callComObjectMethod(d, newNotesDocument, "GetLastDocument")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GETNEXTDOCUMENT_METHOD_COLLECTION.html */
func (d NotesDocumentCollection) GetNextDocument(document NotesDocument) (NotesDocument, error) {
	return callComObjectMethod(d, newNotesDocument, "GetNextDocument", document)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GETNTHDOCUMENT_METHOD_COLLECTION.html */
func (d NotesDocumentCollection) GetNthDocument(n Long) (NotesDocument, error) {
	return callComObjectMethod(d, newNotesDocument, "GetNthDocument", n)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GETPREVDOCUMENT_METHOD_COLLECTION.html */
func (d NotesDocumentCollection) GetPrevDocument(document NotesDocument) (NotesDocument, error) {
	return callComObjectMethod(d, newNotesDocument, "GetPrevDocument", document)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_INTERSECT_METHOD_COLLECTION.html */
/* TODO: Handle different input types (see documentation). */
func (d NotesDocumentCollection) Intersect(inputNotes any) error {
	err := d.checkCombinableTypes(inputNotes)

	if err != nil {
		return err
	}
	return callComVoidMethod(d, "Intersect", inputNotes)
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
	return callComVoidMethod(d, "MarkAllRead", paramsOrdered...)
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
	return callComVoidMethod(d, "MarkAllUnread", paramsOrdered...)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_MERGE_METHOD_COLLECTION.html */
/* TODO: Handle different input types (see documentation). */
func (d NotesDocumentCollection) Merge(inputNotes any) error {
	err := d.checkCombinableTypes(inputNotes)

	if err != nil {
		return err
	}
	return callComVoidMethod(d, "Merge", inputNotes)
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
	return callComVoidMethod(d, "PutAllInFolder", paramsOrdered...)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_REMOVEALL_METHOD.html */
func (d NotesDocumentCollection) RemoveAll(force Boolean) error {
	return callComVoidMethod(d, "RemoveAll", force)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_REMOVEALLFROMFOLDER_METHOD.html */
func (d NotesDocumentCollection) RemoveAllFromFolder(folderName String) error {
	return callComVoidMethod(d, "RemoveAllFromFolder", folderName)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_STAMPALL_METHOD.html */
func (d NotesDocumentCollection) StampAll(itemname String, value any) error {
	return callComVoidMethod(d, "StampAll", itemname, value)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_STAMPALLMULTI_METHOD_DOCCOL.html */
func (d NotesDocumentCollection) StampAllMulti(document NotesDocument) error {
	return callComVoidMethod(d, "StampAllMulti", document)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SUBTRACT_METHOD_COLLECTION.html */
/* TODO: Handle different input types (see documentation). */
func (d NotesDocumentCollection) Subtract(inputNotes any) error {
	err := d.checkCombinableTypes(inputNotes)

	if err != nil {
		return err
	}
	return callComVoidMethod(d, "Subtract", inputNotes)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_UPDATEALL_METHOD.html */
func (d NotesDocumentCollection) UpdateAll() error {
	return callComVoidMethod(d, "UpdateAll")
}
