/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_NOTESAGENT_CLASS.html */
package domigo

import (
	"github.com/monstermichl/domigo/internal/helpers"

	ole "github.com/go-ole/go-ole"
)

type NotesAgent struct {
	NotesStruct
}

func NewNotesAgent(dispatchPtr *ole.IDispatch) NotesAgent {
	return NotesAgent{NewNotesStruct(dispatchPtr)}
}

/* --------------------------------- Properties --------------------------------- */
/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_COMMENT_PROPERTY.html */
func (a NotesAgent) Comment() (String, error) {
	val, err := getComProperty(a, "Comment")
	return helpers.CastValue[String](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_COMMONOWNER_PROPERTY.html */
func (a NotesAgent) CommonOwner() (String, error) {
	val, err := getComProperty(a, "CommonOwner")
	return helpers.CastValue[String](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_HASRUNSINCEMODIFIED_PROPERTY_AGENT.html */
func (a NotesAgent) HasRunSinceModified() (Boolean, error) {
	val, err := getComProperty(a, "HasRunSinceModified")
	return helpers.CastValue[Boolean](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_HTTPURL_PROPERTY_AGENT_COM.html */
func (a NotesAgent) HttpURL() (String, error) {
	val, err := getComProperty(a, "HttpURL")
	return helpers.CastValue[String](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISACTIVATABLE_PROPERTY_AGENT.html */
func (a NotesAgent) IsActivatable() (Boolean, error) {
	val, err := getComProperty(a, "IsActivatable")
	return helpers.CastValue[Boolean](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISENABLED_PROPERTY.html */
func (a NotesAgent) IsEnabled() (Boolean, error) {
	val, err := getComProperty(a, "IsEnabled")
	return helpers.CastValue[Boolean](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISENABLED_PROPERTY.html */
func (a NotesAgent) SetIsEnabled(v Boolean) error {
	return putComProperty(a, "IsEnabled", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISNOTESAGENT_PROPERTY_4410_ABOUT.html */
func (a NotesAgent) IsNotesAgent() (Boolean, error) {
	val, err := getComProperty(a, "IsNotesAgent")
	return helpers.CastValue[Boolean](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISPUBLIC_PROPERTY.html */
func (a NotesAgent) IsPublic() (Boolean, error) {
	val, err := getComProperty(a, "IsPublic")
	return helpers.CastValue[Boolean](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISWEBAGENT_PROPERTY_5149_ABOUT.html */
func (a NotesAgent) IsWebAgent() (Boolean, error) {
	val, err := getComProperty(a, "IsWebAgent")
	return helpers.CastValue[Boolean](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_LASTRUN_PROPERTY_AGENT.html */
func (a NotesAgent) LastRun() (Time, error) {
	val, err := getComProperty(a, "LastRun")
	return helpers.CastValue[Time](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_LOCKHOLDERS_PROPERTY_AGENT.html */
func (a NotesAgent) LockHolders() ([]String, error) {
	vals, err := getComArrayProperty(a, "LockHolders")
	return helpers.CastSlice[String](vals), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_NAME_PROPERTY_AGENT.html */
func (a NotesAgent) Name() (String, error) {
	val, err := getComProperty(a, "Name")
	return helpers.CastValue[String](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_NOTESURL_PROPERTY_AGENT_COM.html */
func (a NotesAgent) NotesURL() (String, error) {
	val, err := getComProperty(a, "NotesURL")
	return helpers.CastValue[String](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ONBEHALFOF_PROPERTY_AGENT.html */
func (a NotesAgent) OnBehalfOf() (String, error) {
	val, err := getComProperty(a, "OnBehalfOf")
	return helpers.CastValue[String](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_OWNER_PROPERTY.html */
func (a NotesAgent) Owner() (String, error) {
	val, err := getComProperty(a, "Owner")
	return helpers.CastValue[String](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_PARAMETERDOCID_PROPERTY_AGENT.html */
func (a NotesAgent) ParameterDocID() (String, error) {
	val, err := getComProperty(a, "ParameterDocID")
	return helpers.CastValue[String](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_PARENT_PROPERTY_AGENT.html */
func (a NotesAgent) Parent() (NotesDatabase, error) {
	dispatchPtr, err := getComObjectProperty(a, "Parent")
	return NewNotesDatabase(dispatchPtr), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_PROHIBITDESIGNUPDATE_PROPERTY.html */
func (a NotesAgent) ProhibitDesignUpdate() (Boolean, error) {
	val, err := getComProperty(a, "ProhibitDesignUpdate")
	return helpers.CastValue[Boolean](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_PROHIBITDESIGNUPDATE_PROPERTY.html */
func (a NotesAgent) SetProhibitDesignUpdate(v Boolean) error {
	return putComProperty(a, "ProhibitDesignUpdate", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_QUERY_PROPERTY_AGENT.html */
func (a NotesAgent) Query() (String, error) {
	val, err := getComProperty(a, "Query")
	return helpers.CastValue[String](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SERVERNAME_PROPERTY.html */
func (a NotesAgent) ServerName() (String, error) {
	val, err := getComProperty(a, "ServerName")
	return helpers.CastValue[String](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SERVERNAME_PROPERTY.html */
func (a NotesAgent) SetServerName(v String) error {
	return putComProperty(a, "ServerName", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_TARGET_PROPERTY_1466_ABOUT.html */
func (a NotesAgent) Target() (Long, error) {
	val, err := getComProperty(a, "Target")
	return helpers.CastValue[Long](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_TRIGGER_PROPERTY_9708_ABOUT.html */
func (a NotesAgent) Trigger() (Long, error) {
	val, err := getComProperty(a, "Trigger")
	return helpers.CastValue[Long](val), err
}

/* --------------------------------- Methods ------------------------------------ */
/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GETPERFORMANCEDOCUMENT_METHOD_AGENT.html */
func (a NotesAgent) GetPerformanceDocument() (NotesAgent, error) {
	dispatchPtr, err := callComObjectMethod(a, "GetPerformanceDocument")
	return NewNotesAgent(dispatchPtr), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_LOCK_METHOD_AGENT.html */
type notesAgentLockParams struct {
	name          *[]String
	provisionalOK *Boolean
}

type notesAgentLockParam func(*notesAgentLockParams)

func WithNotesAgentLockName(name []String) notesAgentLockParam {
	return func(c *notesAgentLockParams) {
		c.name = &name
	}
}

func WithNotesAgentLockProvisionalOK(provisionalOK Boolean) notesAgentLockParam {
	return func(c *notesAgentLockParams) {
		c.provisionalOK = &provisionalOK
	}
}

func (a NotesAgent) Lock(params ...notesAgentLockParam) (Boolean, error) {
	paramsStruct := &notesAgentLockParams{}
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
	val, err := callComMethod(a, "Lock", paramsOrdered...)
	return helpers.CastValue[Boolean](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_LOCKPROVISIONAL_METHOD_AGENT.html */
type notesAgentLockProvisionalParams struct {
	name *[]String
}

type notesAgentLockProvisionalParam func(*notesAgentLockProvisionalParams)

func WithNotesAgentLockProvisionalName(name []String) notesAgentLockProvisionalParam {
	return func(c *notesAgentLockProvisionalParams) {
		c.name = &name
	}
}

func (a NotesAgent) LockProvisional(params ...notesAgentLockProvisionalParam) (Boolean, error) {
	paramsStruct := &notesAgentLockProvisionalParams{}
	paramsOrdered := []interface{}{}

	for _, p := range params {
		p(paramsStruct)
	}

	if paramsStruct.name != nil {
		paramsOrdered = append(paramsOrdered, *paramsStruct.name)
	}
	val, err := callComMethod(a, "LockProvisional", paramsOrdered...)
	return helpers.CastValue[Boolean](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_REMOVE_METHOD_AGENT.html */
func (a NotesAgent) Remove() error {
	_, err := callComMethod(a, "Remove")
	return err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_RUN_METHOD_6415.html */
type notesAgentRunParams struct {
	noteID *String
}

type notesAgentRunParam func(*notesAgentRunParams)

func WithNotesAgentRunNoteID(noteID String) notesAgentRunParam {
	return func(c *notesAgentRunParams) {
		c.noteID = &noteID
	}
}

func (a NotesAgent) Run(params ...notesAgentRunParam) (Integer, error) {
	paramsStruct := &notesAgentRunParams{}
	paramsOrdered := []interface{}{}

	for _, p := range params {
		p(paramsStruct)
	}

	if paramsStruct.noteID != nil {
		paramsOrdered = append(paramsOrdered, *paramsStruct.noteID)
	}
	val, err := callComMethod(a, "Run", paramsOrdered...)
	return helpers.CastValue[Integer](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_RUNONSERVER_METHOD_5924_ABOUT.html */
type notesAgentRunOnServerParams struct {
	noteID *String
}

type notesAgentRunOnServerParam func(*notesAgentRunOnServerParams)

func WithNotesAgentRunOnServerNoteID(noteID String) notesAgentRunOnServerParam {
	return func(c *notesAgentRunOnServerParams) {
		c.noteID = &noteID
	}
}

func (a NotesAgent) RunOnServer(params ...notesAgentRunOnServerParam) (Integer, error) {
	paramsStruct := &notesAgentRunOnServerParams{}
	paramsOrdered := []interface{}{}

	for _, p := range params {
		p(paramsStruct)
	}

	if paramsStruct.noteID != nil {
		paramsOrdered = append(paramsOrdered, *paramsStruct.noteID)
	}
	val, err := callComMethod(a, "RunOnServer", paramsOrdered...)
	return helpers.CastValue[Integer](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SAVE_METHOD_AGENT.html */
func (a NotesAgent) Save() error {
	_, err := callComMethod(a, "Save")
	return err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_UNLOCK_METHOD_AGENT.html */
func (a NotesAgent) UnLock() error {
	_, err := callComMethod(a, "UnLock")
	return err
}
