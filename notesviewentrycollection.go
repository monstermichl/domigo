/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_NOTESVIEWENTRYCOLLECTION_9327.html */
package domigo

import (
	ole "github.com/go-ole/go-ole"
)

type NotesViewEntryCollection struct {
	NotesStruct
}

func newNotesViewEntryCollection(dispatchPtr *ole.IDispatch) NotesViewEntryCollection {
	return NotesViewEntryCollection{newNotesStruct(dispatchPtr)}
}

/* --------------------------------- Properties --------------------------------- */
/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_COUNT_PROPERTY_8365.html */
func (v NotesViewEntryCollection) Count() (Long, error) {
	return getComProperty[Long](v, "Count")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_PARENT_PROPERTY_9184.html */
func (v NotesViewEntryCollection) Parent() (NotesView, error) {
	return getComObjectProperty(v, newNotesView, "Parent")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_QUERY_PROPERTY_8476.html */
func (v NotesViewEntryCollection) Query() (String, error) {
	return getComProperty[String](v, "Query")
}

/* --------------------------------- Methods ------------------------------------ */
/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ADDENTRY_METHOD_3272.html */
type notesViewEntryCollectionAddEntryParams struct {
	checkDups *Boolean
}

type notesViewEntryCollectionAddEntryParam func(*notesViewEntryCollectionAddEntryParams)

func WithNotesViewEntryCollectionAddEntryCheckDups(checkDups Boolean) notesViewEntryCollectionAddEntryParam {
	return func(c *notesViewEntryCollectionAddEntryParams) {
		c.checkDups = &checkDups
	}
}

func (v NotesViewEntryCollection) AddEntry(addentry NotesViewEntry, params ...notesViewEntryCollectionAddEntryParam) error {
	paramsStruct := &notesViewEntryCollectionAddEntryParams{}
	paramsOrdered := []interface{}{addentry}

	for _, p := range params {
		p(paramsStruct)
	}

	if paramsStruct.checkDups != nil {
		paramsOrdered = append(paramsOrdered, *paramsStruct.checkDups)
	}
	return callComVoidMethod(v, "AddEntry", paramsOrdered...)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_CLONE_METHOD_VIEWENTRYCOLLECTION.html */
func (v NotesViewEntryCollection) Clone() error {
	return callComVoidMethod(v, "Clone")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_CONTAINS_METHOD_VIEWENTRYCOLLECTION.html */
func (v NotesViewEntryCollection) Contains(inputNotes any) (Boolean, error) {
	/* TODO: Handle different input types. */
	return callComMethod[Boolean](v, "Contains", inputNotes)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_DELETEENTRY_METHOD_2345.html */
func (v NotesViewEntryCollection) DeleteEntry(deleteentry NotesViewEntry) error {
	return callComVoidMethod(v, "DeleteEntry", deleteentry)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_FTSEARCH_METHOD_6772.html */
func (v NotesViewEntryCollection) FTSearch(query String, maxDocs Integer) error {
	return callComVoidMethod(v, "FTSearch", query, maxDocs)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GETENTRY_METHOD_9111.html */
func (v NotesViewEntryCollection) GetEntry(entry any) (NotesViewEntry, error) {
	/* TODO: Handle different input types (see documentation). */
	return callComObjectMethod(v, newNotesViewEntry, "GetEntry", entry)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GETFIRSTENTRY_METHOD_3097.html */
func (v NotesViewEntryCollection) GetFirstEntry() (NotesViewEntry, error) {
	return callComObjectMethod(v, newNotesViewEntry, "GetFirstEntry")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GETLASTENTRY_METHOD_5374.html */
func (v NotesViewEntryCollection) GetLastEntry() (NotesViewEntry, error) {
	return callComObjectMethod(v, newNotesViewEntry, "GetLastEntry")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GETNEXTENTRY_6168.html */
func (v NotesViewEntryCollection) GetNextEntry(currententry NotesViewEntry) (NotesViewEntry, error) {
	return callComObjectMethod(v, newNotesViewEntry, "GetNextEntry", currententry)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GETNTHENTRY_METHOD_4753.html */
func (v NotesViewEntryCollection) GetNthEntry(index Long) (NotesViewEntry, error) {
	return callComObjectMethod(v, newNotesViewEntry, "GetNthEntry", index)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GETPREVENTRY_METHOD_2018.html */
func (v NotesViewEntryCollection) GetPrevEntry(currententry NotesViewEntry) (NotesViewEntry, error) {
	return callComObjectMethod(v, newNotesViewEntry, "GetPrevEntry", currententry)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_INTERSECT_METHOD_VIEWENTRYCOLLECTION.html */
type notesViewEntryCollectionIntersectParams struct {
	maintainOrder *Boolean
}

type notesViewEntryCollectionIntersectParam func(*notesViewEntryCollectionIntersectParams)

func WithNotesViewEntryCollectionIntersectMaintainOrder(maintainOrder Boolean) notesViewEntryCollectionIntersectParam {
	return func(c *notesViewEntryCollectionIntersectParams) {
		c.maintainOrder = &maintainOrder
	}
}

func (v NotesViewEntryCollection) Intersect(inputNotes Integer, params ...notesViewEntryCollectionIntersectParam) error {
	paramsStruct := &notesViewEntryCollectionIntersectParams{}
	paramsOrdered := []interface{}{inputNotes}

	for _, p := range params {
		p(paramsStruct)
	}

	if paramsStruct.maintainOrder != nil {
		paramsOrdered = append(paramsOrdered, *paramsStruct.maintainOrder)
	}
	return callComVoidMethod(v, "Intersect", paramsOrdered...)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_MARKALLREAD_VIEWCOLLECTION.html */
type notesViewEntryCollectionMarkAllReadParams struct {
	username *String
}

type notesViewEntryCollectionMarkAllReadParam func(*notesViewEntryCollectionMarkAllReadParams)

func WithNotesViewEntryCollectionMarkAllReadUsername(username String) notesViewEntryCollectionMarkAllReadParam {
	return func(c *notesViewEntryCollectionMarkAllReadParams) {
		c.username = &username
	}
}

func (v NotesViewEntryCollection) MarkAllRead(params ...notesViewEntryCollectionMarkAllReadParam) error {
	paramsStruct := &notesViewEntryCollectionMarkAllReadParams{}
	paramsOrdered := []interface{}{}

	for _, p := range params {
		p(paramsStruct)
	}

	if paramsStruct.username != nil {
		paramsOrdered = append(paramsOrdered, *paramsStruct.username)
	}
	return callComVoidMethod(v, "MarkAllRead", paramsOrdered...)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_MARKALLUNREAD_VIEWCOLLECTION.html */
type notesViewEntryCollectionMarkAllUnreadParams struct {
	username *String
}

type notesViewEntryCollectionMarkAllUnreadParam func(*notesViewEntryCollectionMarkAllUnreadParams)

func WithNotesViewEntryCollectionMarkAllUnreadUsername(username String) notesViewEntryCollectionMarkAllUnreadParam {
	return func(c *notesViewEntryCollectionMarkAllUnreadParams) {
		c.username = &username
	}
}

func (v NotesViewEntryCollection) MarkAllUnread(params ...notesViewEntryCollectionMarkAllUnreadParam) error {
	paramsStruct := &notesViewEntryCollectionMarkAllUnreadParams{}
	paramsOrdered := []interface{}{}

	for _, p := range params {
		p(paramsStruct)
	}

	if paramsStruct.username != nil {
		paramsOrdered = append(paramsOrdered, *paramsStruct.username)
	}
	return callComVoidMethod(v, "MarkAllUnread", paramsOrdered...)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_MERGE_METHOD_VIEWENTRYCOLLECTION.html */
func (v NotesViewEntryCollection) Merge(inputNotes Long) error {
	return callComVoidMethod(v, "Merge", inputNotes)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_PUTALLINFOLDER_METHOD_9442.html */
type notesViewEntryCollectionPutAllInFolderParams struct {
	createonfail *Boolean
}

type notesViewEntryCollectionPutAllInFolderParam func(*notesViewEntryCollectionPutAllInFolderParams)

func WithNotesViewEntryCollectionPutAllInFolderCreateonfail(createonfail Boolean) notesViewEntryCollectionPutAllInFolderParam {
	return func(c *notesViewEntryCollectionPutAllInFolderParams) {
		c.createonfail = &createonfail
	}
}

func (v NotesViewEntryCollection) PutAllInFolder(foldername String, params ...notesViewEntryCollectionPutAllInFolderParam) error {
	paramsStruct := &notesViewEntryCollectionPutAllInFolderParams{}
	paramsOrdered := []interface{}{foldername}

	for _, p := range params {
		p(paramsStruct)
	}

	if paramsStruct.createonfail != nil {
		paramsOrdered = append(paramsOrdered, *paramsStruct.createonfail)
	}
	return callComVoidMethod(v, "PutAllInFolder", paramsOrdered...)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_REMOVEALL_METHOD_9278.html */
func (v NotesViewEntryCollection) RemoveAll(force Boolean) error {
	return callComVoidMethod(v, "RemoveAll", force)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_REMOVEALLFROMFOLDER_METHOD_1275.html */
func (v NotesViewEntryCollection) RemoveAllFromFolder(foldername String) error {
	return callComVoidMethod(v, "RemoveAllFromFolder", foldername)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_STAMPALL_METHOD_1012.html */
func (v NotesViewEntryCollection) StampAll(itemname String, value any) error {
	return callComVoidMethod(v, "StampAll", itemname, value)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_STAMPALLMULTI_METHOD.html */
func (v NotesViewEntryCollection) StampAllMulti(document NotesDocument) error {
	return callComVoidMethod(v, "StampAllMulti", document)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SUBTRACT_METHOD_VIEWENTRYCOLLECTION.html */
type notesViewEntryCollectionSubtractParams struct {
	maintainOrder *Boolean
}

type notesViewEntryCollectionSubtractParam func(*notesViewEntryCollectionSubtractParams)

func WithNotesViewEntryCollectionSubtractMaintainOrder(maintainOrder Boolean) notesViewEntryCollectionSubtractParam {
	return func(c *notesViewEntryCollectionSubtractParams) {
		c.maintainOrder = &maintainOrder
	}
}

func (v NotesViewEntryCollection) Subtract(inputNotes Long, params ...notesViewEntryCollectionSubtractParam) error {
	paramsStruct := &notesViewEntryCollectionSubtractParams{}
	paramsOrdered := []interface{}{inputNotes}

	for _, p := range params {
		p(paramsStruct)
	}

	if paramsStruct.maintainOrder != nil {
		paramsOrdered = append(paramsOrdered, *paramsStruct.maintainOrder)
	}
	return callComVoidMethod(v, "Subtract", paramsOrdered...)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_UPDATEALL_METHOD_9459.html */
func (v NotesViewEntryCollection) UpdateAll() error {
	return callComVoidMethod(v, "UpdateAll")
}
