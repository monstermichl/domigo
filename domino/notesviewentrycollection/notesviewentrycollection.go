package notesviewentrycollection

import (
	"github.com/monstermichl/domigo/domino"
	"github.com/monstermichl/domigo/domino/notesdocument"
	"github.com/monstermichl/domigo/domino/notesviewentry"

	"github.com/monstermichl/domigo/helpers"

	ole "github.com/go-ole/go-ole"
)

type NotesViewEntryCollection struct {
	domino.NotesStruct
}

func New(dispatchPtr *ole.IDispatch) NotesViewEntryCollection {
	return NotesViewEntryCollection{domino.NewNotesStruct(dispatchPtr)}
}

/* --------------------------------- Properties --------------------------------- */
/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_COUNT_PROPERTY_8365.html */
func (v NotesViewEntryCollection) Count() (domino.Long, error) {
	val, err := v.Com().GetProperty("Count")
	return helpers.CastValue[domino.Long](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_QUERY_PROPERTY_8476.html */
func (v NotesViewEntryCollection) Query() (domino.String, error) {
	val, err := v.Com().GetProperty("Query")
	return helpers.CastValue[domino.String](val), err
}

/* --------------------------------- Methods ------------------------------------ */
/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ADDENTRY_METHOD_3272.html */
func (v NotesViewEntryCollection) AddEntry(addentry notesviewentry.NotesViewEntry, checkDups domino.Boolean) error {
	_, err := v.Com().CallMethod("AddEntry", addentry.Com().Dispatch(), checkDups)
	return err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_CLONE_METHOD_VIEWENTRYCOLLECTION.html */
func (v NotesViewEntryCollection) Clone() error {
	_, err := v.Com().CallMethod("Clone")
	return err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_CONTAINS_METHOD_VIEWENTRYCOLLECTION.html */
func (v NotesViewEntryCollection) Contains(inputNotes any) (domino.Boolean, error) {
	/* TODO: Handle different input types. */
	val, err := v.Com().CallMethod("Contains", inputNotes)
	return helpers.CastValue[domino.Boolean](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_DELETEENTRY_METHOD_2345.html */
func (v NotesViewEntryCollection) DeleteEntry(deleteentry notesviewentry.NotesViewEntry) error {
	_, err := v.Com().CallMethod("DeleteEntry", deleteentry.Com().Dispatch())
	return err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_FTSEARCH_METHOD_6772.html */
func (v NotesViewEntryCollection) FTSearch(query domino.String, maxDocs domino.Integer) error {
	_, err := v.Com().CallMethod("FTSearch", query, maxDocs)
	return err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GETENTRY_METHOD_9111.html */
func (v NotesViewEntryCollection) GetEntry(entry any) (notesviewentry.NotesViewEntry, error) {
	/* TODO: Handle different input types (see documentation). */
	dispatchPtr, err := v.Com().CallObjectMethod("GetEntry", entry)
	return notesviewentry.New(dispatchPtr), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GETFIRSTENTRY_METHOD_3097.html */
func (v NotesViewEntryCollection) GetFirstEntry() (notesviewentry.NotesViewEntry, error) {
	dispatchPtr, err := v.Com().CallObjectMethod("GetFirstEntry")
	return notesviewentry.New(dispatchPtr), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GETLASTENTRY_METHOD_5374.html */
func (v NotesViewEntryCollection) GetLastEntry() (notesviewentry.NotesViewEntry, error) {
	dispatchPtr, err := v.Com().CallObjectMethod("GetLastEntry")
	return notesviewentry.New(dispatchPtr), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GETNEXTENTRY_6168.html */
func (v NotesViewEntryCollection) GetNextEntry(currententry notesviewentry.NotesViewEntry) (notesviewentry.NotesViewEntry, error) {
	dispatchPtr, err := v.Com().CallObjectMethod("GetNextEntry", currententry.Com().Dispatch())
	return notesviewentry.New(dispatchPtr), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GETNTHENTRY_METHOD_4753.html */
func (v NotesViewEntryCollection) GetNthEntry(index domino.Long) (notesviewentry.NotesViewEntry, error) {
	dispatchPtr, err := v.Com().CallObjectMethod("GetNthEntry", index)
	return notesviewentry.New(dispatchPtr), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GETPREVENTRY_METHOD_2018.html */
func (v NotesViewEntryCollection) GetPrevEntry(currententry notesviewentry.NotesViewEntry) (notesviewentry.NotesViewEntry, error) {
	dispatchPtr, err := v.Com().CallObjectMethod("GetPrevEntry", currententry.Com().Dispatch())
	return notesviewentry.New(dispatchPtr), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_INTERSECT_METHOD_VIEWENTRYCOLLECTION.html */
func (v NotesViewEntryCollection) Intersect(inputNotes domino.Integer, maintainOrder domino.Boolean) error {
	_, err := v.Com().CallMethod("Intersect", inputNotes, maintainOrder)
	return err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_MARKALLREAD_VIEWCOLLECTION.html */
func (v NotesViewEntryCollection) MarkAllRead(username domino.String) error {
	_, err := v.Com().CallMethod("MarkAllRead", username)
	return err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_MARKALLUNREAD_VIEWCOLLECTION.html */
func (v NotesViewEntryCollection) MarkAllUnread(username domino.String) error {
	_, err := v.Com().CallMethod("MarkAllUnread", username)
	return err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_MERGE_METHOD_VIEWENTRYCOLLECTION.html */
func (v NotesViewEntryCollection) Merge(inputNotes domino.Long) error {
	_, err := v.Com().CallMethod("Merge", inputNotes)
	return err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_PUTALLINFOLDER_METHOD_9442.html */
func (v NotesViewEntryCollection) PutAllInFolder(foldername domino.String, createonfail domino.Boolean) error {
	_, err := v.Com().CallMethod("PutAllInFolder", foldername, createonfail)
	return err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_REMOVEALL_METHOD_9278.html */
func (v NotesViewEntryCollection) RemoveAll(force domino.Boolean) error {
	_, err := v.Com().CallMethod("RemoveAll", force)
	return err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_REMOVEALLFROMFOLDER_METHOD_1275.html */
func (v NotesViewEntryCollection) RemoveAllFromFolder(foldername domino.String) error {
	_, err := v.Com().CallMethod("RemoveAllFromFolder", foldername)
	return err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_STAMPALL_METHOD_1012.html */
func (v NotesViewEntryCollection) StampAll(itemname domino.String, value any) error {
	_, err := v.Com().CallMethod("StampAll", itemname, value)
	return err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_STAMPALLMULTI_METHOD.html */
func (v NotesViewEntryCollection) StampAllMulti(document notesdocument.NotesDocument) error {
	_, err := v.Com().CallMethod("StampAllMulti", document.Com().Dispatch())
	return err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SUBTRACT_METHOD_VIEWENTRYCOLLECTION.html */
func (v NotesViewEntryCollection) Subtract(inputNotes domino.Long, maintainOrder domino.Boolean) error {
	_, err := v.Com().CallMethod("Subtract", inputNotes, maintainOrder)
	return err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_UPDATEALL_METHOD_9459.html */
func (v NotesViewEntryCollection) UpdateAll() error {
	_, err := v.Com().CallMethod("UpdateAll")
	return err
}
