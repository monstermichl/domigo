/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_NOTESDOCUMENTCOLLECTION_CLASS.html */
package notesdocumentcollection

import (
	"domigo/domino"
	"domigo/domino/notesdatetime"
	"domigo/domino/notesdocument"
	"domigo/helpers"

	ole "github.com/go-ole/go-ole"
)

type NotesDocumentCollection struct {
	domino.NotesStruct
}

func New(dispatchPtr *ole.IDispatch) NotesDocumentCollection {
	return NotesDocumentCollection{domino.NewNotesStruct(dispatchPtr)}
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_RESPONSES_PROPERTY.html */
/* Moved from NotesDocument. */
func NotesDocumentResponses(d notesdocument.NotesDocument) (NotesDocumentCollection, error) {
	dispatchPtr, err := d.Com().GetObjectProperty("Responses")
	return New(dispatchPtr), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ATTACHVCARD_METHOD.html */
/* Moved from NotesDocument. */
func NotesDocumentAttachVCard(d notesdocument.NotesDocument, clientADTObject NotesDocumentCollection) error {
	_, err := d.Com().CallMethod("AttachVCard", clientADTObject)
	return err
}

/* --------------------------------- Properties --------------------------------- */
/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_COUNT_PROPERTY.html */
func (d NotesDocumentCollection) Count() (domino.Long, error) {
	val, err := d.Com().GetProperty("Count")
	return helpers.CastValue[domino.Long](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISSORTED_PROPERTY_COLLECTION.html */
func (d NotesDocumentCollection) IsSorted() (domino.Boolean, error) {
	val, err := d.Com().GetProperty("IsSorted")
	return helpers.CastValue[domino.Boolean](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_QUERY_PROPERTY_COLLECTION.html */
func (d NotesDocumentCollection) Query() (domino.String, error) {
	val, err := d.Com().GetProperty("Query")
	return helpers.CastValue[domino.String](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_UNTILTIME_PROPERTY_COLLECTION.html */
func (d NotesDocumentCollection) UntilTime() (notesdatetime.NotesDateTime, error) {
	dispatchPtr, err := d.Com().GetObjectProperty("UntilTime")
	return notesdatetime.New(dispatchPtr), err
}

/* --------------------------------- Methods ------------------------------------ */
/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ADDDOCUMENT_METHOD_7116_ABOUT.html */
type addDocumentParams struct {
	checkDups *domino.Boolean
}

type addDocumentParam func(*addDocumentParams)

func WithAddDocumentCheckDups(checkDups domino.Boolean) addDocumentParam {
	return func(c *addDocumentParams) {
		c.checkDups = &checkDups
	}
}

func (d NotesDocumentCollection) AddDocument(document notesdocument.NotesDocument, params ...addDocumentParam) error {
	paramsStruct := &addDocumentParams{}
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
func (d NotesDocumentCollection) ContainsNoteID(id domino.String) error {
	_, err := d.Com().CallMethod("Contains", id)
	return err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_CONTAINS_METHOD_COLLECTION.html */
func (d NotesDocumentCollection) ContainsNotesDocument(doc notesdocument.NotesDocument) error {
	_, err := d.Com().CallMethod("Contains", doc.Com().Dispatch())
	return err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_CONTAINS_METHOD_COLLECTION.html */
func (d NotesDocumentCollection) ContainsNotesDocumentCollection(col NotesDocumentCollection) error {
	_, err := d.Com().CallMethod("Contains", col.Com().Dispatch())
	return err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_DELETEDOCUMENT_METHOD_7984_ABOUT.html */
func (d NotesDocumentCollection) DeleteDocument(document notesdocument.NotesDocument) error {
	_, err := d.Com().CallMethod("DeleteDocument", document.Com().Dispatch())
	return err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_FTSEARCH_METHOD_COLLECTION.html */
func (d NotesDocumentCollection) FTSearch(query domino.String, maxDocs domino.Integer) error {
	_, err := d.Com().CallMethod("FTSearch", query, maxDocs)
	return err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GETDOCUMENT_METHOD_DOCCOLL.html */
func (d NotesDocumentCollection) GetDocument(document notesdocument.NotesDocument) (notesdocument.NotesDocument, error) {
	dispatchPtr, err := d.Com().CallObjectMethod("GetDocument", document.Com().Dispatch())
	return notesdocument.New(dispatchPtr), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GETFIRSTDOCUMENT_METHOD_COLLECTION.html */
func (d NotesDocumentCollection) GetFirstDocument() (notesdocument.NotesDocument, error) {
	dispatchPtr, err := d.Com().CallObjectMethod("GetFirstDocument")
	return notesdocument.New(dispatchPtr), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GETLASTDOCUMENT_METHOD_COLLECTION.html */
func (d NotesDocumentCollection) GetLastDocument() (notesdocument.NotesDocument, error) {
	dispatchPtr, err := d.Com().CallObjectMethod("GetLastDocument")
	return notesdocument.New(dispatchPtr), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GETNEXTDOCUMENT_METHOD_COLLECTION.html */
func (d NotesDocumentCollection) GetNextDocument(document notesdocument.NotesDocument) (notesdocument.NotesDocument, error) {
	dispatchPtr, err := d.Com().CallObjectMethod("GetNextDocument", document.Com().Dispatch())
	return notesdocument.New(dispatchPtr), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GETNTHDOCUMENT_METHOD_COLLECTION.html */
func (d NotesDocumentCollection) GetNthDocument(n domino.Long) (notesdocument.NotesDocument, error) {
	dispatchPtr, err := d.Com().CallObjectMethod("GetNthDocument", n)
	return notesdocument.New(dispatchPtr), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GETPREVDOCUMENT_METHOD_COLLECTION.html */
func (d NotesDocumentCollection) GetPrevDocument(document notesdocument.NotesDocument) (notesdocument.NotesDocument, error) {
	dispatchPtr, err := d.Com().CallObjectMethod("GetPrevDocument", document.Com().Dispatch())
	return notesdocument.New(dispatchPtr), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_INTERSECT_METHOD_COLLECTION.html */
/* TODO: Handle different input types (see documentation). */
func (d NotesDocumentCollection) IntersectNoteID(id domino.String) error {
	_, err := d.Com().CallMethod("Intersect", id)
	return err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_INTERSECT_METHOD_COLLECTION.html */
func (d NotesDocumentCollection) IntersectNotesDocument(doc notesdocument.NotesDocument) error {
	_, err := d.Com().CallMethod("Intersect", doc.Com().Dispatch())
	return err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_INTERSECT_METHOD_COLLECTION.html */
func (d NotesDocumentCollection) IntersectNotesDocumentCollection(col NotesDocumentCollection) error {
	_, err := d.Com().CallMethod("Intersect", col.Com().Dispatch())
	return err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_MARKALLREAD_DOCCOLLECTION.html */
type markAllReadParams struct {
	username *domino.String
}

type markAllReadParam func(*markAllReadParams)

func WithMarkAllReadUsername(username domino.String) markAllReadParam {
	return func(c *markAllReadParams) {
		c.username = &username
	}
}

func (d NotesDocumentCollection) MarkAllRead(params ...markAllReadParam) error {
	paramsStruct := &markAllReadParams{}
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
type markAllUnreadParams struct {
	username *domino.String
}

type markAllUnreadParam func(*markAllUnreadParams)

func WithMarkAllUnreadUsername(username domino.String) markAllUnreadParam {
	return func(c *markAllUnreadParams) {
		c.username = &username
	}
}

func (d NotesDocumentCollection) MarkAllUnread(params ...markAllUnreadParam) error {
	paramsStruct := &markAllUnreadParams{}
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
func (d NotesDocumentCollection) MergeNoteID(id domino.String) error {
	_, err := d.Com().CallMethod("Merge", id)
	return err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_MERGE_METHOD_COLLECTION.html */
func (d NotesDocumentCollection) MergeNotesDocument(doc notesdocument.NotesDocument) error {
	_, err := d.Com().CallMethod("Merge", doc.Com().Dispatch())
	return err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_MERGE_METHOD_COLLECTION.html */
func (d NotesDocumentCollection) MergeNotesDocumentCollection(col NotesDocumentCollection) error {
	_, err := d.Com().CallMethod("Merge", col.Com().Dispatch())
	return err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_PUTALLINFOLDER_METHOD.html */
type putAllInFolderParams struct {
	createonfail *domino.Boolean
}

type putAllInFolderParam func(*putAllInFolderParams)

func WithPutAllInFolderCreateonfail(createonfail domino.Boolean) putAllInFolderParam {
	return func(c *putAllInFolderParams) {
		c.createonfail = &createonfail
	}
}

func (d NotesDocumentCollection) PutAllInFolder(folderName domino.String, params ...putAllInFolderParam) error {
	paramsStruct := &putAllInFolderParams{}
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
func (d NotesDocumentCollection) RemoveAll(force domino.Boolean) error {
	_, err := d.Com().CallMethod("RemoveAll", force)
	return err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_REMOVEALLFROMFOLDER_METHOD.html */
func (d NotesDocumentCollection) RemoveAllFromFolder(folderName domino.String) error {
	_, err := d.Com().CallMethod("RemoveAllFromFolder", folderName)
	return err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_STAMPALL_METHOD.html */
func (d NotesDocumentCollection) StampAll(itemname domino.String, value any) error {
	_, err := d.Com().CallMethod("StampAll", itemname, value)
	return err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_STAMPALLMULTI_METHOD_DOCCOL.html */
func (d NotesDocumentCollection) StampAllMulti(document notesdocument.NotesDocument) error {
	_, err := d.Com().CallMethod("StampAllMulti", document.Com().Dispatch())
	return err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SUBTRACT_METHOD_COLLECTION.html */
/* TODO: Handle different input types (see documentation). */
func (d NotesDocumentCollection) SubtractNoteID(id domino.String) error {
	_, err := d.Com().CallMethod("Subtract", id)
	return err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SUBTRACT_METHOD_COLLECTION.html */
func (d NotesDocumentCollection) SubtractNotesDocument(doc notesdocument.NotesDocument) error {
	_, err := d.Com().CallMethod("Subtract", doc.Com().Dispatch())
	return err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SUBTRACT_METHOD_COLLECTION.html */
func (d NotesDocumentCollection) SubtractNotesDocumentCollection(col NotesDocumentCollection) error {
	_, err := d.Com().CallMethod("Subtract", col.Com().Dispatch())
	return err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_UPDATEALL_METHOD.html */
func (d NotesDocumentCollection) UpdateAll() error {
	_, err := d.Com().CallMethod("UpdateAll")
	return err
}
