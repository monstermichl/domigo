/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_NOTESREPLICATION_CLASS_1289.html */
package domigo

import (
	"github.com/monstermichl/domigo/internal/com"

	ole "github.com/go-ole/go-ole"
)

type NotesReplication struct {
	NotesStruct
}

func newNotesReplication(dispatchPtr *ole.IDispatch) NotesReplication {
	return NotesReplication{newNotesStruct(dispatchPtr)}
}

/* --------------------------------- Properties --------------------------------- */
/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ABSTRACT_PROPERTY_2067.html */
func (r NotesReplication) Abstract() (Boolean, error) {
	return getComProperty[Boolean](r, "Abstract")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ABSTRACT_PROPERTY_2067.html */
func (r NotesReplication) SetAbstract(v Boolean) error {
	return putComProperty(r, "Abstract", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_CUTOFFDATE_PROPERTY_296.html */
func (r NotesReplication) CutoffDate() (Time, error) {
	return getComProperty[Time](r, "CutoffDate")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_CUTOFFDELETE_PROPERTY_8613.html */
func (r NotesReplication) CutoffDelete() (Boolean, error) {
	return getComProperty[Boolean](r, "CutoffDelete")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_CUTOFFDELETE_PROPERTY_8613.html */
func (r NotesReplication) SetCutoffDelete(v Boolean) error {
	return putComProperty(r, "CutoffDelete", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_CUTOFFINTERVAL_PROPERTY_6968.html */
func (r NotesReplication) CutoffInterval() (Long, error) {
	return getComProperty[Long](r, "CutoffInterval")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_CUTOFFINTERVAL_PROPERTY_6968.html */
func (r NotesReplication) SetCutoffInterval(v Long) error {
	return putComProperty(r, "CutoffInterval", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_DISABLED_PROPERTY_7391.html */
func (r NotesReplication) Disabled() (Boolean, error) {
	return getComProperty[Boolean](r, "Disabled")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_DISABLED_PROPERTY_7391.html */
func (r NotesReplication) SetDisabled(v Boolean) error {
	return putComProperty(r, "Disabled", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_DONTSENDLOCALSECURITYUPDATES_PROPERTY_REP.html */
func (r NotesReplication) DontSendLocalSecurityUpdates() (Boolean, error) {
	return getComProperty[Boolean](r, "DontSendLocalSecurityUpdates")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_DONTSENDLOCALSECURITYUPDATES_PROPERTY_REP.html */
func (r NotesReplication) SetDontSendLocalSecurityUpdates(v Boolean) error {
	return putComProperty(r, "DontSendLocalSecurityUpdates", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_IGNOREDELETES_PROPERTY_9218.html */
func (r NotesReplication) IgnoreDeletes() (Boolean, error) {
	return getComProperty[Boolean](r, "IgnoreDeletes")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_IGNOREDELETES_PROPERTY_9218.html */
func (r NotesReplication) SetIgnoreDeletes(v Boolean) error {
	return putComProperty(r, "IgnoreDeletes", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_IGNOREDESTDELETES_PROPERTY_3069.html */
func (r NotesReplication) IgnoreDestDeletes() (Boolean, error) {
	return getComProperty[Boolean](r, "IgnoreDestDeletes")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_IGNOREDESTDELETES_PROPERTY_3069.html */
func (r NotesReplication) SetIgnoreDestDeletes(v Boolean) error {
	return putComProperty(r, "IgnoreDestDeletes", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_PRIORITY_PROPERTY_6071.html */
func (r NotesReplication) Priority() (Long, error) {
	return getComProperty[Long](r, "Priority")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_PRIORITY_PROPERTY_6071.html */
func (r NotesReplication) SetPriority(v Long) error {
	return putComProperty(r, "Priority", v)
}

/* --------------------------------- Methods ------------------------------------ */
/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_CLEARHISTORY_METHOD_8188.html */
func (r NotesReplication) ClearHistory() error {
	return callComVoidMethod(r, "ClearHistory")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GETENTRY_METHOD_REP.html */
type notesReplicationGetEntryParams struct {
	createflag *Boolean
}

type notesReplicationGetEntryParam func(*notesReplicationGetEntryParams)

func WithNotesReplicationGetEntryCreateflag(createflag Boolean) notesReplicationGetEntryParam {
	return func(c *notesReplicationGetEntryParams) {
		c.createflag = &createflag
	}
}

func (r NotesReplication) GetEntry(source String, destination String, params ...notesReplicationGetEntryParam) (NotesReplicationEntry, error) {
	paramsStruct := &notesReplicationGetEntryParams{}
	paramsOrdered := []interface{}{source, destination}

	for _, p := range params {
		p(paramsStruct)
	}

	if paramsStruct.createflag != nil {
		paramsOrdered = append(paramsOrdered, *paramsStruct.createflag)
	}
	return callComObjectMethod(r, newNotesReplicationEntry, "GetEntry", paramsOrdered...)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GETENTRIES_METHOD.html */
func (r NotesReplication) GetEntries() ([]NotesReplicationEntry, error) {
	return com.CallObjectArrayMethod(r.com(), newNotesReplicationEntry, "GetEntries")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_RESET_METHOD_9955.html */
func (r NotesReplication) Reset() error {
	return callComVoidMethod(r, "Reset")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SAVE_METHOD_3588.html */
func (r NotesReplication) Save() error {
	return callComVoidMethod(r, "Save")
}
