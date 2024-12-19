/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_NOTESVIEWENTRYCOLLECTION_9327.html */
package domigo

import (
	"github.com/monstermichl/domigo/internal/helpers"

	ole "github.com/go-ole/go-ole"
)

type NotesViewEntryCollection struct {
	NotesStruct
}

func NewNotesViewEntryCollection(dispatchPtr *ole.IDispatch) NotesViewEntryCollection {
	return NotesViewEntryCollection{NewNotesStruct(dispatchPtr)}
}

/* --------------------------------- Properties --------------------------------- */
/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_COUNT_PROPERTY_8365.html */
func (v NotesViewEntryCollection) Count() (Long, error) {
	val, err := v.com().GetProperty("Count")
	return helpers.CastValue[Long](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_PARENT_PROPERTY_9184.html */
func (v NotesViewEntryCollection) Parent() (NotesView, error) {
	dispatchPtr, err := v.com().GetObjectProperty("Parent")
	return NewNotesView(dispatchPtr), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_QUERY_PROPERTY_8476.html */
func (v NotesViewEntryCollection) Query() (String, error) {
	val, err := v.com().GetProperty("Query")
	return helpers.CastValue[String](val), err
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
	paramsOrdered := []interface{}{addentry.com().Dispatch()}

	for _, p := range params {
		p(paramsStruct)
	}

	if paramsStruct.checkDups != nil {
		paramsOrdered = append(paramsOrdered, *paramsStruct.checkDups)
	}
	_, err := v.com().CallMethod("AddEntry", paramsOrdered...)
	return err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_CLONE_METHOD_VIEWENTRYCOLLECTION.html */
func (v NotesViewEntryCollection) Clone() error {
	_, err := v.com().CallMethod("Clone")
	return err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_CONTAINS_METHOD_VIEWENTRYCOLLECTION.html */
func (v NotesViewEntryCollection) Contains(inputNotes any) (Boolean, error) {
	/* TODO: Handle different input types. */
	val, err := v.com().CallMethod("Contains", inputNotes)
	return helpers.CastValue[Boolean](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_DELETEENTRY_METHOD_2345.html */
func (v NotesViewEntryCollection) DeleteEntry(deleteentry NotesViewEntry) error {
	_, err := v.com().CallMethod("DeleteEntry", deleteentry.com().Dispatch())
	return err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_FTSEARCH_METHOD_6772.html */
func (v NotesViewEntryCollection) FTSearch(query String, maxDocs Integer) error {
	_, err := v.com().CallMethod("FTSearch", query, maxDocs)
	return err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GETENTRY_METHOD_9111.html */
func (v NotesViewEntryCollection) GetEntry(entry any) (NotesViewEntry, error) {
	/* TODO: Handle different input types (see documentation). */
	dispatchPtr, err := v.com().CallObjectMethod("GetEntry", entry)
	return NewNotesViewEntry(dispatchPtr), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GETFIRSTENTRY_METHOD_3097.html */
func (v NotesViewEntryCollection) GetFirstEntry() (NotesViewEntry, error) {
	dispatchPtr, err := v.com().CallObjectMethod("GetFirstEntry")
	return NewNotesViewEntry(dispatchPtr), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GETLASTENTRY_METHOD_5374.html */
func (v NotesViewEntryCollection) GetLastEntry() (NotesViewEntry, error) {
	dispatchPtr, err := v.com().CallObjectMethod("GetLastEntry")
	return NewNotesViewEntry(dispatchPtr), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GETNEXTENTRY_6168.html */
func (v NotesViewEntryCollection) GetNextEntry(currententry NotesViewEntry) (NotesViewEntry, error) {
	dispatchPtr, err := v.com().CallObjectMethod("GetNextEntry", currententry.com().Dispatch())
	return NewNotesViewEntry(dispatchPtr), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GETNTHENTRY_METHOD_4753.html */
func (v NotesViewEntryCollection) GetNthEntry(index Long) (NotesViewEntry, error) {
	dispatchPtr, err := v.com().CallObjectMethod("GetNthEntry", index)
	return NewNotesViewEntry(dispatchPtr), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GETPREVENTRY_METHOD_2018.html */
func (v NotesViewEntryCollection) GetPrevEntry(currententry NotesViewEntry) (NotesViewEntry, error) {
	dispatchPtr, err := v.com().CallObjectMethod("GetPrevEntry", currententry.com().Dispatch())
	return NewNotesViewEntry(dispatchPtr), err
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
	_, err := v.com().CallMethod("Intersect", paramsOrdered...)
	return err
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
	_, err := v.com().CallMethod("MarkAllRead", paramsOrdered...)
	return err
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
	_, err := v.com().CallMethod("MarkAllUnread", paramsOrdered...)
	return err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_MERGE_METHOD_VIEWENTRYCOLLECTION.html */
func (v NotesViewEntryCollection) Merge(inputNotes Long) error {
	_, err := v.com().CallMethod("Merge", inputNotes)
	return err
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
	_, err := v.com().CallMethod("PutAllInFolder", paramsOrdered...)
	return err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_REMOVEALL_METHOD_9278.html */
func (v NotesViewEntryCollection) RemoveAll(force Boolean) error {
	_, err := v.com().CallMethod("RemoveAll", force)
	return err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_REMOVEALLFROMFOLDER_METHOD_1275.html */
func (v NotesViewEntryCollection) RemoveAllFromFolder(foldername String) error {
	_, err := v.com().CallMethod("RemoveAllFromFolder", foldername)
	return err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_STAMPALL_METHOD_1012.html */
func (v NotesViewEntryCollection) StampAll(itemname String, value any) error {
	_, err := v.com().CallMethod("StampAll", itemname, value)
	return err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_STAMPALLMULTI_METHOD.html */
func (v NotesViewEntryCollection) StampAllMulti(document NotesDocument) error {
	_, err := v.com().CallMethod("StampAllMulti", document.com().Dispatch())
	return err
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
	_, err := v.com().CallMethod("Subtract", paramsOrdered...)
	return err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_UPDATEALL_METHOD_9459.html */
func (v NotesViewEntryCollection) UpdateAll() error {
	_, err := v.com().CallMethod("UpdateAll")
	return err
}
