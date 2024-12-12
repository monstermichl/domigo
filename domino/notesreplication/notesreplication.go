/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_NOTESREPLICATION_CLASS_1289.html */
package notesreplication

import (
	"domigo/domino"
	"domigo/domino/com"
	"domigo/domino/notesreplicationentry"
	"domigo/helpers"

	ole "github.com/go-ole/go-ole"
)

type NotesReplication struct {
	domino.NotesStruct
}

func New(dispatchPtr *ole.IDispatch) NotesReplication {
	return NotesReplication{domino.NewNotesStruct(dispatchPtr)}
}

/* --------------------------------- Properties --------------------------------- */
/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ABSTRACT_PROPERTY_2067.html */
func (r NotesReplication) Abstract() (domino.Boolean, error) {
	val, err := r.Com().GetProperty("Abstract")
	return helpers.CastValue[domino.Boolean](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ABSTRACT_PROPERTY_2067.html */
func (r NotesReplication) SetAbstract(v domino.Boolean) error {
	return r.Com().PutProperty("Abstract", v)
}

/* TODO: Access type for CutoffDate could not be evaluated, check yourself if getter/setter is needed. */
/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_CUTOFFDATE_PROPERTY_296.html */
func (r NotesReplication) CutoffDate() (domino.Time, error) {
	val, err := r.Com().GetProperty("CutoffDate")
	return helpers.CastValue[domino.Time](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_CUTOFFDATE_PROPERTY_296.html */
func (r NotesReplication) SetCutoffDate(v domino.Time) error {
	return r.Com().PutProperty("CutoffDate", v)
}

/* TODO: Access type for CutoffDelete could not be evaluated, check yourself if getter/setter is needed. */
/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_CUTOFFDELETE_PROPERTY_8613.html */
func (r NotesReplication) CutoffDelete() (domino.Boolean, error) {
	val, err := r.Com().GetProperty("CutoffDelete")
	return helpers.CastValue[domino.Boolean](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_CUTOFFDELETE_PROPERTY_8613.html */
func (r NotesReplication) SetCutoffDelete(v domino.Boolean) error {
	return r.Com().PutProperty("CutoffDelete", v)
}

/* TODO: Access type for CutoffInterval could not be evaluated, check yourself if getter/setter is needed. */
/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_CUTOFFINTERVAL_PROPERTY_6968.html */
func (r NotesReplication) CutoffInterval() (domino.Long, error) {
	val, err := r.Com().GetProperty("CutoffInterval")
	return helpers.CastValue[domino.Long](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_CUTOFFINTERVAL_PROPERTY_6968.html */
func (r NotesReplication) SetCutoffInterval(v domino.Long) error {
	return r.Com().PutProperty("CutoffInterval", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_DISABLED_PROPERTY_7391.html */
func (r NotesReplication) Disabled() (domino.Boolean, error) {
	val, err := r.Com().GetProperty("Disabled")
	return helpers.CastValue[domino.Boolean](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_DISABLED_PROPERTY_7391.html */
func (r NotesReplication) SetDisabled(v domino.Boolean) error {
	return r.Com().PutProperty("Disabled", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_DONTSENDLOCALSECURITYUPDATES_PROPERTY_REP.html */
func (r NotesReplication) DontSendLocalSecurityUpdates() (domino.Boolean, error) {
	val, err := r.Com().GetProperty("DontSendLocalSecurityUpdates")
	return helpers.CastValue[domino.Boolean](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_DONTSENDLOCALSECURITYUPDATES_PROPERTY_REP.html */
func (r NotesReplication) SetDontSendLocalSecurityUpdates(v domino.Boolean) error {
	return r.Com().PutProperty("DontSendLocalSecurityUpdates", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_IGNOREDELETES_PROPERTY_9218.html */
func (r NotesReplication) IgnoreDeletes() (domino.Boolean, error) {
	val, err := r.Com().GetProperty("IgnoreDeletes")
	return helpers.CastValue[domino.Boolean](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_IGNOREDELETES_PROPERTY_9218.html */
func (r NotesReplication) SetIgnoreDeletes(v domino.Boolean) error {
	return r.Com().PutProperty("IgnoreDeletes", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_IGNOREDESTDELETES_PROPERTY_3069.html */
func (r NotesReplication) IgnoreDestDeletes() (domino.Boolean, error) {
	val, err := r.Com().GetProperty("IgnoreDestDeletes")
	return helpers.CastValue[domino.Boolean](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_IGNOREDESTDELETES_PROPERTY_3069.html */
func (r NotesReplication) SetIgnoreDestDeletes(v domino.Boolean) error {
	return r.Com().PutProperty("IgnoreDestDeletes", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_PRIORITY_PROPERTY_6071.html */
func (r NotesReplication) Priority() (domino.Long, error) {
	val, err := r.Com().GetProperty("Priority")
	return helpers.CastValue[domino.Long](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_PRIORITY_PROPERTY_6071.html */
func (r NotesReplication) SetPriority(v domino.Long) error {
	return r.Com().PutProperty("Priority", v)
}

/* --------------------------------- Methods ------------------------------------ */
/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_CLEARHISTORY_METHOD_8188.html */
func (r NotesReplication) ClearHistory() error {
	_, err := r.Com().CallMethod("ClearHistory")
	return err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GETENTRY_METHOD_REP.html */
func (r NotesReplication) GetEntry(source domino.String, destination domino.String, createflag domino.Boolean) (notesreplicationentry.NotesReplicationEntry, error) {
	params := helpers.OptionalParams(source, destination, createflag)

	dispatchPtr, err := r.Com().CallObjectMethod("GetEntry", params...)
	return notesreplicationentry.New(dispatchPtr), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GETENTRIES_METHOD.html */
func (r NotesReplication) GetEntries() ([]notesreplicationentry.NotesReplicationEntry, error) {
	return com.GetObjectArrayProperty(r.Com(), notesreplicationentry.New, "GetEntries")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_RESET_METHOD_9955.html */
func (r NotesReplication) Reset() error {
	_, err := r.Com().CallMethod("Reset")
	return err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SAVE_METHOD_3588.html */
func (r NotesReplication) Save() error {
	_, err := r.Com().CallMethod("Save")
	return err
}
