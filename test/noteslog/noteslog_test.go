/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_NOTESLOG_CLASS.html */
package noteslog_test

import (
	"domigo/domino/noteslog"
	"domigo/domino/notessession"
	"testing"

	"github.com/stretchr/testify/require"
)

var log noteslog.NotesLog

/* https://pkg.go.dev/testing#hdr-Main */
func TestMain(m *testing.M) {
	session, _ := notessession.Initialize()
	log, _ = session.CreateLog("test")

	defer session.Release()

	m.Run()
}

/* --------------------------------- Properties --------------------------------- */
/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_LOGACTIONS_PROPERTY.html */
func TestLogActions(t *testing.T) {
	_, err := log.LogActions()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_LOGACTIONS_PROPERTY.html */
func TestSetLogActions(t *testing.T) {
	s, err := log.LogActions()
	require.Nil(t, err)

	err = log.SetLogActions(s)
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_LOGERRORS_PROPERTY.html */
func TestLogErrors(t *testing.T) {
	_, err := log.LogErrors()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_LOGERRORS_PROPERTY.html */
func TestSetLogErrors(t *testing.T) {
	s, err := log.LogErrors()
	require.Nil(t, err)

	err = log.SetLogErrors(s)
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_NUMACTIONS_PROPERTY.html */
func TestNumActions(t *testing.T) {
	_, err := log.NumActions()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_NUMERRORS_PROPERTY.html */
func TestNumErrors(t *testing.T) {
	_, err := log.NumErrors()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_OVERWRITEFILE_PROPERTY.html */
func TestOverwriteFile(t *testing.T) {
	_, err := log.OverwriteFile()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_OVERWRITEFILE_PROPERTY.html */
func TestSetOverwriteFile(t *testing.T) {
	s, err := log.OverwriteFile()
	require.Nil(t, err)

	err = log.SetOverwriteFile(s)
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_PROGRAMNAME_PROPERTY.html */
func TestProgramName(t *testing.T) {
	_, err := log.ProgramName()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_PROGRAMNAME_PROPERTY.html */
func TestSetProgramName(t *testing.T) {
	s, err := log.ProgramName()
	require.Nil(t, err)

	err = log.SetProgramName(s)
	require.Nil(t, err)
}

/* --------------------------------- Methods ------------------------------------ */
/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_CLOSE_METHOD_LOG.html */
func TestClose(t *testing.T) {
	err := log.Close()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_LOGACTION_METHOD.html */
func TestLogAction(t *testing.T) {
	err := log.LogAction("test-actions")
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_LOGERROR_METHOD.html */
func TestLogError(t *testing.T) {
	err := log.LogError(1, "test-error")
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_LOGEVENT_METHOD.html */
func TestLogEvent(t *testing.T) {
	err := log.LogEvent("test-event", "test-event-queue", 1, 1)
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_OPENAGENTLOG_METHOD.html */
func TestOpenAgentLog(t *testing.T) {
	err := log.OpenAgentLog()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_OPENFILELOG_METHOD.html */
// func TestOpenFileLog(t *testing.T) {
// 	err := log.OpenFileLog( /* TODO: Pass test values. */ )
// 	require.Nil(t, err)
// }

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_OPENMAILLOG_METHOD.html */
// func TestOpenMailLog(t *testing.T) {
// 	err := log.OpenMailLog( /* TODO: Pass test values. */ )
// 	require.Nil(t, err)
// }

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_OPENNOTESLOG_METHOD.html */
// func TestOpenNotesLog(t *testing.T) {
// 	err := log.OpenNotesLog( /* TODO: Pass test values. */ )
// 	require.Nil(t, err)
// }
