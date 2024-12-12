/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_NOTESLOG_CLASS.html */
package noteslog

import (
	"domigo/domino"
	"domigo/helpers"

	ole "github.com/go-ole/go-ole"
)

type NotesLog struct {
	domino.NotesStruct
}

func New(dispatchPtr *ole.IDispatch) NotesLog {
	return NotesLog{domino.NewNotesStruct(dispatchPtr)}
}

/* --------------------------------- Properties --------------------------------- */
/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_LOGACTIONS_PROPERTY.html */
func (l NotesLog) LogActions() (domino.Boolean, error) {
	val, err := l.Com().GetProperty("LogActions")
	return helpers.CastValue[domino.Boolean](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_LOGACTIONS_PROPERTY.html */
func (l NotesLog) SetLogActions(v domino.Boolean) error {
	return l.Com().PutProperty("LogActions", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_LOGERRORS_PROPERTY.html */
func (l NotesLog) LogErrors() (domino.Boolean, error) {
	val, err := l.Com().GetProperty("LogErrors")
	return helpers.CastValue[domino.Boolean](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_LOGERRORS_PROPERTY.html */
func (l NotesLog) SetLogErrors(v domino.Boolean) error {
	return l.Com().PutProperty("LogErrors", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_NUMACTIONS_PROPERTY.html */
func (l NotesLog) NumActions() (domino.Long, error) {
	val, err := l.Com().GetProperty("NumActions")
	return helpers.CastValue[domino.Long](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_NUMERRORS_PROPERTY.html */
func (l NotesLog) NumErrors() (domino.Long, error) {
	val, err := l.Com().GetProperty("NumErrors")
	return helpers.CastValue[domino.Long](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_OVERWRITEFILE_PROPERTY.html */
func (l NotesLog) OverwriteFile() (domino.Boolean, error) {
	val, err := l.Com().GetProperty("OverwriteFile")
	return helpers.CastValue[domino.Boolean](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_OVERWRITEFILE_PROPERTY.html */
func (l NotesLog) SetOverwriteFile(v domino.Boolean) error {
	return l.Com().PutProperty("OverwriteFile", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_PROGRAMNAME_PROPERTY.html */
func (l NotesLog) ProgramName() (domino.String, error) {
	val, err := l.Com().GetProperty("ProgramName")
	return helpers.CastValue[domino.String](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_PROGRAMNAME_PROPERTY.html */
func (l NotesLog) SetProgramName(v domino.String) error {
	return l.Com().PutProperty("ProgramName", v)
}

/* --------------------------------- Methods ------------------------------------ */
/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_CLOSE_METHOD_LOG.html */
func (l NotesLog) Close() error {
	_, err := l.Com().CallMethod("Close")
	return err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_LOGACTION_METHOD.html */
func (l NotesLog) LogAction(description domino.String) error {
	_, err := l.Com().CallMethod("LogAction", description)
	return err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_LOGERROR_METHOD.html */
func (l NotesLog) LogError(code domino.Integer, description domino.String) error {
	_, err := l.Com().CallMethod("LogError", code, description)
	return err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_LOGEVENT_METHOD.html */
func (l NotesLog) LogEvent(message domino.String, queuename domino.String, eventType domino.Integer, severity domino.Integer) error {
	_, err := l.Com().CallMethod("LogEvent", message, queuename, eventType, severity)
	return err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_OPENAGENTLOG_METHOD.html */
func (l NotesLog) OpenAgentLog() error {
	_, err := l.Com().CallMethod("OpenAgentLog")
	return err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_OPENFILELOG_METHOD.html */
func (l NotesLog) OpenFileLog(path domino.String) error {
	_, err := l.Com().CallMethod("OpenFileLog", path)
	return err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_OPENMAILLOG_METHOD.html */
func (l NotesLog) OpenMailLog(recipientsV []domino.Single, subject domino.String) error {
	_, err := l.Com().CallMethod("OpenMailLog", recipientsV, subject)
	return err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_OPENNOTESLOG_METHOD.html */
func (l NotesLog) OpenNotesLog(server domino.String, dbfile domino.String) error {
	_, err := l.Com().CallMethod("OpenNotesLog", server, dbfile)
	return err
}
