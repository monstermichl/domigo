/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_NOTESLOG_CLASS.html */
package domigo

import (
	ole "github.com/go-ole/go-ole"
)

type NotesLogEventType = Long

const (
	NOTESLOG_EV_ALARM    NotesLogEventType = 8
	NOTESLOG_EV_COMM     NotesLogEventType = 1
	NOTESLOG_EV_MAIL     NotesLogEventType = 3
	NOTESLOG_EV_MISC     NotesLogEventType = 6
	NOTESLOG_EV_REPLICA  NotesLogEventType = 4
	NOTESLOG_EV_RESOURCE NotesLogEventType = 5
	NOTESLOG_EV_SECURITY NotesLogEventType = 2
	NOTESLOG_EV_SERVER   NotesLogEventType = 7
	NOTESLOG_EV_UNKNOWN  NotesLogEventType = 0
	NOTESLOG_EV_UPDATE   NotesLogEventType = 9
)

type NotesLogSeverity = Long

const (
	SEV_FAILURE  NotesLogSeverity = 2
	SEV_FATAL    NotesLogSeverity = 1
	SEV_NORMAL   NotesLogSeverity = 5
	SEV_WARNING1 NotesLogSeverity = 3
	SEV_WARNING2 NotesLogSeverity = 4
)

type NotesLog struct {
	NotesStruct
}

func newNotesLog(dispatchPtr *ole.IDispatch) NotesLog {
	return NotesLog{newNotesStruct(dispatchPtr)}
}

/* --------------------------------- Properties --------------------------------- */
/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_LOGACTIONS_PROPERTY.html */
func (l NotesLog) LogActions() (Boolean, error) {
	return getComProperty[Boolean](l, "LogActions")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_LOGACTIONS_PROPERTY.html */
func (l NotesLog) SetLogActions(v Boolean) error {
	return putComProperty(l, "LogActions", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_LOGERRORS_PROPERTY.html */
func (l NotesLog) LogErrors() (Boolean, error) {
	return getComProperty[Boolean](l, "LogErrors")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_LOGERRORS_PROPERTY.html */
func (l NotesLog) SetLogErrors(v Boolean) error {
	return putComProperty(l, "LogErrors", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_NUMACTIONS_PROPERTY.html */
func (l NotesLog) NumActions() (Long, error) {
	return getComProperty[Long](l, "NumActions")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_NUMERRORS_PROPERTY.html */
func (l NotesLog) NumErrors() (Long, error) {
	return getComProperty[Long](l, "NumErrors")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_OVERWRITEFILE_PROPERTY.html */
func (l NotesLog) OverwriteFile() (Boolean, error) {
	return getComProperty[Boolean](l, "OverwriteFile")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_OVERWRITEFILE_PROPERTY.html */
func (l NotesLog) SetOverwriteFile(v Boolean) error {
	return putComProperty(l, "OverwriteFile", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_PARENT_PROPERTY_LOG_COM.html */
func (l NotesLog) Parent() (NotesSession, error) {
	return getComObjectProperty(l, newNotesSession, "Parent")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_PROGRAMNAME_PROPERTY.html */
func (l NotesLog) ProgramName() (String, error) {
	return getComProperty[String](l, "ProgramName")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_PROGRAMNAME_PROPERTY.html */
func (l NotesLog) SetProgramName(v String) error {
	return putComProperty(l, "ProgramName", v)
}

/* --------------------------------- Methods ------------------------------------ */
/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_CLOSE_METHOD_LOG.html */
func (l NotesLog) Close() error {
	return callComVoidMethod(l, "Close")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_LOGACTION_METHOD.html */
func (l NotesLog) LogAction(description String) error {
	return callComVoidMethod(l, "LogAction", description)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_LOGERROR_METHOD.html */
func (l NotesLog) LogError(code Integer, description String) error {
	return callComVoidMethod(l, "LogError", code, description)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_LOGEVENT_METHOD.html */
func (l NotesLog) LogEvent(message String, queuename String, eventType NotesLogEventType, severity NotesLogSeverity) error {
	return callComVoidMethod(l, "LogEvent", message, queuename, eventType, severity)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_OPENAGENTLOG_METHOD.html */
func (l NotesLog) OpenAgentLog() error {
	return callComVoidMethod(l, "OpenAgentLog")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_OPENFILELOG_METHOD.html */
func (l NotesLog) OpenFileLog(path String) error {
	return callComVoidMethod(l, "OpenFileLog", path)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_OPENMAILLOG_METHOD.html */
func (l NotesLog) OpenMailLog(recipientsV []Single, subject String) error {
	return callComVoidMethod(l, "OpenMailLog", recipientsV, subject)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_OPENNOTESLOG_METHOD.html */
func (l NotesLog) OpenNotesLog(server String, dbfile String) error {
	return callComVoidMethod(l, "OpenNotesLog", server, dbfile)
}
