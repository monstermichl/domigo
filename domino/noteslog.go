/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_NOTESLOG_CLASS.html */
package domigo

import (
	"github.com/monstermichl/domigo/helpers"

	ole "github.com/go-ole/go-ole"
)

type NotesLogEventType Long

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

type NotesLogSeverity Long

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

func NewNotesLog(dispatchPtr *ole.IDispatch) NotesLog {
	return NotesLog{NewNotesStruct(dispatchPtr)}
}

/* --------------------------------- Properties --------------------------------- */
/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_LOGACTIONS_PROPERTY.html */
func (l NotesLog) LogActions() (Boolean, error) {
	val, err := l.Com().GetProperty("LogActions")
	return helpers.CastValue[Boolean](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_LOGACTIONS_PROPERTY.html */
func (l NotesLog) SetLogActions(v Boolean) error {
	return l.Com().PutProperty("LogActions", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_LOGERRORS_PROPERTY.html */
func (l NotesLog) LogErrors() (Boolean, error) {
	val, err := l.Com().GetProperty("LogErrors")
	return helpers.CastValue[Boolean](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_LOGERRORS_PROPERTY.html */
func (l NotesLog) SetLogErrors(v Boolean) error {
	return l.Com().PutProperty("LogErrors", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_NUMACTIONS_PROPERTY.html */
func (l NotesLog) NumActions() (Long, error) {
	val, err := l.Com().GetProperty("NumActions")
	return helpers.CastValue[Long](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_NUMERRORS_PROPERTY.html */
func (l NotesLog) NumErrors() (Long, error) {
	val, err := l.Com().GetProperty("NumErrors")
	return helpers.CastValue[Long](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_OVERWRITEFILE_PROPERTY.html */
func (l NotesLog) OverwriteFile() (Boolean, error) {
	val, err := l.Com().GetProperty("OverwriteFile")
	return helpers.CastValue[Boolean](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_OVERWRITEFILE_PROPERTY.html */
func (l NotesLog) SetOverwriteFile(v Boolean) error {
	return l.Com().PutProperty("OverwriteFile", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_PARENT_PROPERTY_LOG_COM.html */
func (l NotesLog) Parent() (NotesSession, error) {
	dispatchPtr, err := l.Com().GetObjectProperty("Parent")
	return NewNotesSession(dispatchPtr), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_PROGRAMNAME_PROPERTY.html */
func (l NotesLog) ProgramName() (String, error) {
	val, err := l.Com().GetProperty("ProgramName")
	return helpers.CastValue[String](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_PROGRAMNAME_PROPERTY.html */
func (l NotesLog) SetProgramName(v String) error {
	return l.Com().PutProperty("ProgramName", v)
}

/* --------------------------------- Methods ------------------------------------ */
/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_CLOSE_METHOD_LOG.html */
func (l NotesLog) Close() error {
	_, err := l.Com().CallMethod("Close")
	return err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_LOGACTION_METHOD.html */
func (l NotesLog) LogAction(description String) error {
	_, err := l.Com().CallMethod("LogAction", description)
	return err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_LOGERROR_METHOD.html */
func (l NotesLog) LogError(code Integer, description String) error {
	_, err := l.Com().CallMethod("LogError", code, description)
	return err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_LOGEVENT_METHOD.html */
func (l NotesLog) LogEvent(message String, queuename String, eventType NotesLogEventType, severity NotesLogSeverity) error {
	_, err := l.Com().CallMethod("LogEvent", message, queuename, eventType, severity)
	return err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_OPENAGENTLOG_METHOD.html */
func (l NotesLog) OpenAgentLog() error {
	_, err := l.Com().CallMethod("OpenAgentLog")
	return err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_OPENFILELOG_METHOD.html */
func (l NotesLog) OpenFileLog(path String) error {
	_, err := l.Com().CallMethod("OpenFileLog", path)
	return err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_OPENMAILLOG_METHOD.html */
func (l NotesLog) OpenMailLog(recipientsV []Single, subject String) error {
	_, err := l.Com().CallMethod("OpenMailLog", recipientsV, subject)
	return err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_OPENNOTESLOG_METHOD.html */
func (l NotesLog) OpenNotesLog(server String, dbfile String) error {
	_, err := l.Com().CallMethod("OpenNotesLog", server, dbfile)
	return err
}
