/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_NOTESREPLICATION_CLASS_1289.html */
package domigo

import (
	"github.com/monstermichl/domigo/domino/com"
	"github.com/monstermichl/domigo/helpers"

	ole "github.com/go-ole/go-ole"
)

type NotesReplication struct {
	NotesStruct
}

func NewNotesReplication(dispatchPtr *ole.IDispatch) NotesReplication {
	return NotesReplication{NewNotesStruct(dispatchPtr)}
}

/* --------------------------------- Properties --------------------------------- */
/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ABSTRACT_PROPERTY_2067.html */
func (r NotesReplication) Abstract() (Boolean, error) {
	val, err := r.com().GetProperty("Abstract")
	return helpers.CastValue[Boolean](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ABSTRACT_PROPERTY_2067.html */
func (r NotesReplication) SetAbstract(v Boolean) error {
	return r.com().PutProperty("Abstract", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_CUTOFFDATE_PROPERTY_296.html */
func (r NotesReplication) CutoffDate() (Time, error) {
	val, err := r.com().GetProperty("CutoffDate")
	return helpers.CastValue[Time](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_CUTOFFDELETE_PROPERTY_8613.html */
func (r NotesReplication) CutoffDelete() (Boolean, error) {
	val, err := r.com().GetProperty("CutoffDelete")
	return helpers.CastValue[Boolean](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_CUTOFFDELETE_PROPERTY_8613.html */
func (r NotesReplication) SetCutoffDelete(v Boolean) error {
	return r.com().PutProperty("CutoffDelete", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_CUTOFFINTERVAL_PROPERTY_6968.html */
func (r NotesReplication) CutoffInterval() (Long, error) {
	val, err := r.com().GetProperty("CutoffInterval")
	return helpers.CastValue[Long](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_CUTOFFINTERVAL_PROPERTY_6968.html */
func (r NotesReplication) SetCutoffInterval(v Long) error {
	return r.com().PutProperty("CutoffInterval", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_DISABLED_PROPERTY_7391.html */
func (r NotesReplication) Disabled() (Boolean, error) {
	val, err := r.com().GetProperty("Disabled")
	return helpers.CastValue[Boolean](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_DISABLED_PROPERTY_7391.html */
func (r NotesReplication) SetDisabled(v Boolean) error {
	return r.com().PutProperty("Disabled", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_DONTSENDLOCALSECURITYUPDATES_PROPERTY_REP.html */
func (r NotesReplication) DontSendLocalSecurityUpdates() (Boolean, error) {
	val, err := r.com().GetProperty("DontSendLocalSecurityUpdates")
	return helpers.CastValue[Boolean](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_DONTSENDLOCALSECURITYUPDATES_PROPERTY_REP.html */
func (r NotesReplication) SetDontSendLocalSecurityUpdates(v Boolean) error {
	return r.com().PutProperty("DontSendLocalSecurityUpdates", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_IGNOREDELETES_PROPERTY_9218.html */
func (r NotesReplication) IgnoreDeletes() (Boolean, error) {
	val, err := r.com().GetProperty("IgnoreDeletes")
	return helpers.CastValue[Boolean](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_IGNOREDELETES_PROPERTY_9218.html */
func (r NotesReplication) SetIgnoreDeletes(v Boolean) error {
	return r.com().PutProperty("IgnoreDeletes", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_IGNOREDESTDELETES_PROPERTY_3069.html */
func (r NotesReplication) IgnoreDestDeletes() (Boolean, error) {
	val, err := r.com().GetProperty("IgnoreDestDeletes")
	return helpers.CastValue[Boolean](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_IGNOREDESTDELETES_PROPERTY_3069.html */
func (r NotesReplication) SetIgnoreDestDeletes(v Boolean) error {
	return r.com().PutProperty("IgnoreDestDeletes", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_PRIORITY_PROPERTY_6071.html */
func (r NotesReplication) Priority() (Long, error) {
	val, err := r.com().GetProperty("Priority")
	return helpers.CastValue[Long](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_PRIORITY_PROPERTY_6071.html */
func (r NotesReplication) SetPriority(v Long) error {
	return r.com().PutProperty("Priority", v)
}

/* --------------------------------- Methods ------------------------------------ */
/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_CLEARHISTORY_METHOD_8188.html */
func (r NotesReplication) ClearHistory() error {
	_, err := r.com().CallMethod("ClearHistory")
	return err
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
	dispatchPtr, err := r.com().CallObjectMethod("GetEntry", paramsOrdered...)
	return NewNotesReplicationEntry(dispatchPtr), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GETENTRIES_METHOD.html */
func (r NotesReplication) GetEntries() ([]NotesReplicationEntry, error) {
	return com.CallObjectArrayMethod(r.com(), NewNotesReplicationEntry, "GetEntries")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_RESET_METHOD_9955.html */
func (r NotesReplication) Reset() error {
	_, err := r.com().CallMethod("Reset")
	return err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SAVE_METHOD_3588.html */
func (r NotesReplication) Save() error {
	_, err := r.com().CallMethod("Save")
	return err
}
