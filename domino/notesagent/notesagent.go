/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_NOTESAGENT_CLASS.html */
package notesagent

import (
	"domigo/domino"
	"domigo/helpers"

	ole "github.com/go-ole/go-ole"
)

type NotesAgent struct {
	domino.NotesStruct
}

func New(dispatchPtr *ole.IDispatch) NotesAgent {
	return NotesAgent{domino.NewNotesStruct(dispatchPtr)}
}

/* --------------------------------- Properties --------------------------------- */
/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_COMMENT_PROPERTY.html */
func (a NotesAgent) Comment() (domino.String, error) {
	val, err := a.Com().GetProperty("Comment")
	return helpers.CastValue[domino.String](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_COMMONOWNER_PROPERTY.html */
func (a NotesAgent) CommonOwner() (domino.String, error) {
	val, err := a.Com().GetProperty("CommonOwner")
	return helpers.CastValue[domino.String](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_HASRUNSINCEMODIFIED_PROPERTY_AGENT.html */
func (a NotesAgent) HasRunSinceModified() (domino.Boolean, error) {
	val, err := a.Com().GetProperty("HasRunSinceModified")
	return helpers.CastValue[domino.Boolean](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_HTTPURL_PROPERTY_AGENT_COM.html */
func (a NotesAgent) HttpURL() (domino.String, error) {
	val, err := a.Com().GetProperty("HttpURL")
	return helpers.CastValue[domino.String](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISACTIVATABLE_PROPERTY_AGENT.html */
func (a NotesAgent) IsActivatable() (domino.Boolean, error) {
	val, err := a.Com().GetProperty("IsActivatable")
	return helpers.CastValue[domino.Boolean](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISENABLED_PROPERTY.html */
func (a NotesAgent) IsEnabled() (domino.Boolean, error) {
	val, err := a.Com().GetProperty("IsEnabled")
	return helpers.CastValue[domino.Boolean](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISENABLED_PROPERTY.html */
func (a NotesAgent) SetIsEnabled(v domino.Boolean) error {
	return a.Com().PutProperty("IsEnabled", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISNOTESAGENT_PROPERTY_4410_ABOUT.html */
func (a NotesAgent) IsNotesAgent() (domino.Boolean, error) {
	val, err := a.Com().GetProperty("IsNotesAgent")
	return helpers.CastValue[domino.Boolean](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISPUBLIC_PROPERTY.html */
func (a NotesAgent) IsPublic() (domino.Boolean, error) {
	val, err := a.Com().GetProperty("IsPublic")
	return helpers.CastValue[domino.Boolean](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISWEBAGENT_PROPERTY_5149_ABOUT.html */
func (a NotesAgent) IsWebAgent() (domino.Boolean, error) {
	val, err := a.Com().GetProperty("IsWebAgent")
	return helpers.CastValue[domino.Boolean](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_LASTRUN_PROPERTY_AGENT.html */
func (a NotesAgent) LastRun() (domino.Time, error) {
	val, err := a.Com().GetProperty("LastRun")
	return helpers.CastValue[domino.Time](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_LOCKHOLDERS_PROPERTY_AGENT.html */
func (a NotesAgent) LockHolders() ([]domino.String, error) {
	vals, err := a.Com().GetArrayProperty("LockHolders")
	return helpers.CastSlice[domino.String](vals), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_NAME_PROPERTY_AGENT.html */
func (a NotesAgent) Name() (domino.String, error) {
	val, err := a.Com().GetProperty("Name")
	return helpers.CastValue[domino.String](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_NOTESURL_PROPERTY_AGENT_COM.html */
func (a NotesAgent) NotesURL() (domino.String, error) {
	val, err := a.Com().GetProperty("NotesURL")
	return helpers.CastValue[domino.String](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ONBEHALFOF_PROPERTY_AGENT.html */
func (a NotesAgent) OnBehalfOf() (domino.String, error) {
	val, err := a.Com().GetProperty("OnBehalfOf")
	return helpers.CastValue[domino.String](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_OWNER_PROPERTY.html */
func (a NotesAgent) Owner() (domino.String, error) {
	val, err := a.Com().GetProperty("Owner")
	return helpers.CastValue[domino.String](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_PARAMETERDOCID_PROPERTY_AGENT.html */
func (a NotesAgent) ParameterDocID() (domino.String, error) {
	val, err := a.Com().GetProperty("ParameterDocID")
	return helpers.CastValue[domino.String](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_PROHIBITDESIGNUPDATE_PROPERTY.html */
func (a NotesAgent) ProhibitDesignUpdate() (domino.Boolean, error) {
	val, err := a.Com().GetProperty("ProhibitDesignUpdate")
	return helpers.CastValue[domino.Boolean](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_PROHIBITDESIGNUPDATE_PROPERTY.html */
func (a NotesAgent) SetProhibitDesignUpdate(v domino.Boolean) error {
	return a.Com().PutProperty("ProhibitDesignUpdate", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_QUERY_PROPERTY_AGENT.html */
func (a NotesAgent) Query() (domino.String, error) {
	val, err := a.Com().GetProperty("Query")
	return helpers.CastValue[domino.String](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SERVERNAME_PROPERTY.html */
func (a NotesAgent) ServerName() (domino.String, error) {
	val, err := a.Com().GetProperty("ServerName")
	return helpers.CastValue[domino.String](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SERVERNAME_PROPERTY.html */
func (a NotesAgent) SetServerName(v domino.String) error {
	return a.Com().PutProperty("ServerName", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_TARGET_PROPERTY_1466_ABOUT.html */
func (a NotesAgent) Target() (domino.Long, error) {
	val, err := a.Com().GetProperty("Target")
	return helpers.CastValue[domino.Long](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_TRIGGER_PROPERTY_9708_ABOUT.html */
func (a NotesAgent) Trigger() (domino.Long, error) {
	val, err := a.Com().GetProperty("Trigger")
	return helpers.CastValue[domino.Long](val), err
}

/* --------------------------------- Methods ------------------------------------ */
/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_LOCK_METHOD_AGENT.html */
type lockParams struct {
	name          *[]domino.String
	provisionalOK *domino.Boolean
}

type lockParam func(*lockParams)

func WithLockName(name []domino.String) lockParam {
	return func(c *lockParams) {
		c.name = &name
	}
}

func WithLockProvisionalOK(provisionalOK domino.Boolean) lockParam {
	return func(c *lockParams) {
		c.provisionalOK = &provisionalOK
	}
}

func (a NotesAgent) Lock(params ...lockParam) (domino.Boolean, error) {
	paramsStruct := &lockParams{}
	paramsOrdered := []interface{}{}

	for _, p := range params {
		p(paramsStruct)
	}

	if paramsStruct.name != nil {
		paramsOrdered = append(paramsOrdered, *paramsStruct.name)
		if paramsStruct.provisionalOK != nil {
			paramsOrdered = append(paramsOrdered, *paramsStruct.provisionalOK)
		}
	}
	val, err := a.Com().CallMethod("Lock", paramsOrdered...)
	return helpers.CastValue[domino.Boolean](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_LOCKPROVISIONAL_METHOD_AGENT.html */
type lockProvisionalParams struct {
	name *[]domino.String
}

type lockProvisionalParam func(*lockProvisionalParams)

func WithLockProvisionalName(name []domino.String) lockProvisionalParam {
	return func(c *lockProvisionalParams) {
		c.name = &name
	}
}

func (a NotesAgent) LockProvisional(params ...lockProvisionalParam) (domino.Boolean, error) {
	paramsStruct := &lockProvisionalParams{}
	paramsOrdered := []interface{}{}

	for _, p := range params {
		p(paramsStruct)
	}

	if paramsStruct.name != nil {
		paramsOrdered = append(paramsOrdered, *paramsStruct.name)
	}
	val, err := a.Com().CallMethod("LockProvisional", paramsOrdered...)
	return helpers.CastValue[domino.Boolean](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_REMOVE_METHOD_AGENT.html */
func (a NotesAgent) Remove() error {
	_, err := a.Com().CallMethod("Remove")
	return err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_RUN_METHOD_6415.html */
type runParams struct {
	noteID *domino.String
}

type runParam func(*runParams)

func WithRunNoteID(noteID domino.String) runParam {
	return func(c *runParams) {
		c.noteID = &noteID
	}
}

func (a NotesAgent) Run(params ...runParam) (domino.Integer, error) {
	paramsStruct := &runParams{}
	paramsOrdered := []interface{}{}

	for _, p := range params {
		p(paramsStruct)
	}

	if paramsStruct.noteID != nil {
		paramsOrdered = append(paramsOrdered, *paramsStruct.noteID)
	}
	val, err := a.Com().CallMethod("Run", paramsOrdered...)
	return helpers.CastValue[domino.Integer](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_RUNONSERVER_METHOD_5924_ABOUT.html */
type runOnServerParams struct {
	noteID *domino.String
}

type runOnServerParam func(*runOnServerParams)

func WithRunOnServerNoteID(noteID domino.String) runOnServerParam {
	return func(c *runOnServerParams) {
		c.noteID = &noteID
	}
}

func (a NotesAgent) RunOnServer(params ...runOnServerParam) (domino.Integer, error) {
	paramsStruct := &runOnServerParams{}
	paramsOrdered := []interface{}{}

	for _, p := range params {
		p(paramsStruct)
	}

	if paramsStruct.noteID != nil {
		paramsOrdered = append(paramsOrdered, *paramsStruct.noteID)
	}
	val, err := a.Com().CallMethod("RunOnServer", paramsOrdered...)
	return helpers.CastValue[domino.Integer](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SAVE_METHOD_AGENT.html */
func (a NotesAgent) Save() error {
	_, err := a.Com().CallMethod("Save")
	return err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_UNLOCK_METHOD_AGENT.html */
func (a NotesAgent) UnLock() error {
	_, err := a.Com().CallMethod("UnLock")
	return err
}
