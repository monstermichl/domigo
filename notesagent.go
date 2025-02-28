/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_NOTESAGENT_CLASS.html */
package domigo

import (
	ole "github.com/go-ole/go-ole"
)

type NotesAgent struct {
	NotesStruct
}

func newNotesAgent(dispatchPtr *ole.IDispatch) NotesAgent {
	return NotesAgent{newNotesStruct(dispatchPtr)}
}

/* --------------------------------- Properties --------------------------------- */
/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_COMMENT_PROPERTY.html */
func (a NotesAgent) Comment() (String, error) {
	return getComProperty[String](a, "Comment")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_COMMONOWNER_PROPERTY.html */
func (a NotesAgent) CommonOwner() (String, error) {
	return getComProperty[String](a, "CommonOwner")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_HASRUNSINCEMODIFIED_PROPERTY_AGENT.html */
func (a NotesAgent) HasRunSinceModified() (Boolean, error) {
	return getComProperty[Boolean](a, "HasRunSinceModified")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_HTTPURL_PROPERTY_AGENT_COM.html */
func (a NotesAgent) HttpURL() (String, error) {
	return getComProperty[String](a, "HttpURL")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISACTIVATABLE_PROPERTY_AGENT.html */
func (a NotesAgent) IsActivatable() (Boolean, error) {
	return getComProperty[Boolean](a, "IsActivatable")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISENABLED_PROPERTY.html */
func (a NotesAgent) IsEnabled() (Boolean, error) {
	return getComProperty[Boolean](a, "IsEnabled")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISENABLED_PROPERTY.html */
func (a NotesAgent) SetIsEnabled(v Boolean) error {
	return putComProperty(a, "IsEnabled", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISNOTESAGENT_PROPERTY_4410_ABOUT.html */
func (a NotesAgent) IsNotesAgent() (Boolean, error) {
	return getComProperty[Boolean](a, "IsNotesAgent")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISPUBLIC_PROPERTY.html */
func (a NotesAgent) IsPublic() (Boolean, error) {
	return getComProperty[Boolean](a, "IsPublic")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISWEBAGENT_PROPERTY_5149_ABOUT.html */
func (a NotesAgent) IsWebAgent() (Boolean, error) {
	return getComProperty[Boolean](a, "IsWebAgent")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_LASTRUN_PROPERTY_AGENT.html */
func (a NotesAgent) LastRun() (Time, error) {
	return getComProperty[Time](a, "LastRun")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_LOCKHOLDERS_PROPERTY_AGENT.html */
func (a NotesAgent) LockHolders() ([]String, error) {
	return getComArrayProperty[String](a, "LockHolders")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_NAME_PROPERTY_AGENT.html */
func (a NotesAgent) Name() (String, error) {
	return getComProperty[String](a, "Name")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_NOTESURL_PROPERTY_AGENT_COM.html */
func (a NotesAgent) NotesURL() (String, error) {
	return getComProperty[String](a, "NotesURL")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ONBEHALFOF_PROPERTY_AGENT.html */
func (a NotesAgent) OnBehalfOf() (String, error) {
	return getComProperty[String](a, "OnBehalfOf")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_OWNER_PROPERTY.html */
func (a NotesAgent) Owner() (String, error) {
	return getComProperty[String](a, "Owner")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_PARAMETERDOCID_PROPERTY_AGENT.html */
func (a NotesAgent) ParameterDocID() (String, error) {
	return getComProperty[String](a, "ParameterDocID")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_PARENT_PROPERTY_AGENT.html */
func (a NotesAgent) Parent() (NotesDatabase, error) {
	return getComObjectProperty(a, newNotesDatabase, "Parent")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_PROHIBITDESIGNUPDATE_PROPERTY.html */
func (a NotesAgent) ProhibitDesignUpdate() (Boolean, error) {
	return getComProperty[Boolean](a, "ProhibitDesignUpdate")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_PROHIBITDESIGNUPDATE_PROPERTY.html */
func (a NotesAgent) SetProhibitDesignUpdate(v Boolean) error {
	return putComProperty(a, "ProhibitDesignUpdate", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_QUERY_PROPERTY_AGENT.html */
func (a NotesAgent) Query() (String, error) {
	return getComProperty[String](a, "Query")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SERVERNAME_PROPERTY.html */
func (a NotesAgent) ServerName() (String, error) {
	return getComProperty[String](a, "ServerName")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SERVERNAME_PROPERTY.html */
func (a NotesAgent) SetServerName(v String) error {
	return putComProperty(a, "ServerName", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_TARGET_PROPERTY_1466_ABOUT.html */
func (a NotesAgent) Target() (Long, error) {
	return getComProperty[Long](a, "Target")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_TRIGGER_PROPERTY_9708_ABOUT.html */
func (a NotesAgent) Trigger() (Long, error) {
	return getComProperty[Long](a, "Trigger")
}

/* --------------------------------- Methods ------------------------------------ */
/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GETPERFORMANCEDOCUMENT_METHOD_AGENT.html */
func (a NotesAgent) GetPerformanceDocument() (NotesAgent, error) {
	return callComObjectMethod(a, newNotesAgent, "GetPerformanceDocument")
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
	return callComMethod[Boolean](a, "Lock", paramsOrdered...)
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
	return callComMethod[Boolean](a, "LockProvisional", paramsOrdered...)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_REMOVE_METHOD_AGENT.html */
func (a NotesAgent) Remove() error {
	return callComVoidMethod(a, "Remove")
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
	return callComMethod[Integer](a, "Run", paramsOrdered...)
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
	return callComMethod[Integer](a, "RunOnServer", paramsOrdered...)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SAVE_METHOD_AGENT.html */
func (a NotesAgent) Save() error {
	return callComVoidMethod(a, "Save")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_UNLOCK_METHOD_AGENT.html */
func (a NotesAgent) UnLock() error {
	return callComVoidMethod(a, "UnLock")
}
