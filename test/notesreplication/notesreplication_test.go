/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_NOTESREPLICATION_CLASS_1289.html */
package notesreplication_test

import (
	"testing"

	"github.com/monstermichl/domigo"
	testhelpers "github.com/monstermichl/domigo/test/helpers"

	"github.com/stretchr/testify/require"
)

var replication domigo.NotesReplication

/* https://pkg.go.dev/testing#hdr-Main */
func TestMain(m *testing.M) {
	testhelpers.Initialize(func(session domigo.NotesSession, db domigo.NotesDatabase) (string, error) {
		var err error
		replication, err = db.ReplicationInfo()
		defer replication.Release()

		if err != nil {
			return "Replaction could not be created", err
		}
		m.Run()
		return "", nil
	})
}

/* --------------------------------- Properties --------------------------------- */
/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ABSTRACT_PROPERTY_2067.html */
func TestAbstract(t *testing.T) {
	_, err := replication.Abstract()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ABSTRACT_PROPERTY_2067.html */
func TestSetAbstract(t *testing.T) {
	s, err := replication.Abstract()
	require.Nil(t, err)

	err = replication.SetAbstract(s)
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_CUTOFFDATE_PROPERTY_296.html */
func TestCutoffDate(t *testing.T) {
	_, err := replication.CutoffDate()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_CUTOFFDELETE_PROPERTY_8613.html */
func TestCutoffDelete(t *testing.T) {
	_, err := replication.CutoffDelete()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_CUTOFFDELETE_PROPERTY_8613.html */
func TestSetCutoffDelete(t *testing.T) {
	s, err := replication.CutoffDelete()
	require.Nil(t, err)

	err = replication.SetCutoffDelete(s)
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_CUTOFFINTERVAL_PROPERTY_6968.html */
func TestCutoffInterval(t *testing.T) {
	_, err := replication.CutoffInterval()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_CUTOFFINTERVAL_PROPERTY_6968.html */
func TestSetCutoffInterval(t *testing.T) {
	s, err := replication.CutoffInterval()
	require.Nil(t, err)

	err = replication.SetCutoffInterval(s)
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_DISABLED_PROPERTY_7391.html */
func TestDisabled(t *testing.T) {
	_, err := replication.Disabled()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_DISABLED_PROPERTY_7391.html */
func TestSetDisabled(t *testing.T) {
	s, err := replication.Disabled()
	require.Nil(t, err)

	err = replication.SetDisabled(s)
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_DONTSENDLOCALSECURITYUPDATES_PROPERTY_REP.html */
func TestDontSendLocalSecurityUpdates(t *testing.T) {
	_, err := replication.DontSendLocalSecurityUpdates()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_DONTSENDLOCALSECURITYUPDATES_PROPERTY_REP.html */
func TestSetDontSendLocalSecurityUpdates(t *testing.T) {
	s, err := replication.DontSendLocalSecurityUpdates()
	require.Nil(t, err)

	err = replication.SetDontSendLocalSecurityUpdates(s)
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_IGNOREDELETES_PROPERTY_9218.html */
func TestIgnoreDeletes(t *testing.T) {
	_, err := replication.IgnoreDeletes()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_IGNOREDELETES_PROPERTY_9218.html */
func TestSetIgnoreDeletes(t *testing.T) {
	s, err := replication.IgnoreDeletes()
	require.Nil(t, err)

	err = replication.SetIgnoreDeletes(s)
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_IGNOREDESTDELETES_PROPERTY_3069.html */
func TestIgnoreDestDeletes(t *testing.T) {
	_, err := replication.IgnoreDestDeletes()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_IGNOREDESTDELETES_PROPERTY_3069.html */
func TestSetIgnoreDestDeletes(t *testing.T) {
	s, err := replication.IgnoreDestDeletes()
	require.Nil(t, err)

	err = replication.SetIgnoreDestDeletes(s)
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_PRIORITY_PROPERTY_6071.html */
func TestPriority(t *testing.T) {
	_, err := replication.Priority()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_PRIORITY_PROPERTY_6071.html */
func TestSetPriority(t *testing.T) {
	s, err := replication.Priority()
	require.Nil(t, err)

	err = replication.SetPriority(s)
	require.Nil(t, err)
}

/* --------------------------------- Methods ------------------------------------ */
/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_CLEARHISTORY_METHOD_8188.html */
func TestClearHistory(t *testing.T) {
	err := replication.ClearHistory()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GETENTRY_METHOD_REP.html */
func TestGetEntry(t *testing.T) {
	_, err := replication.GetEntry("", "")
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GETENTRIES_METHOD.html */
func TestGetEntries(t *testing.T) {
	_, err := replication.GetEntries()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_RESET_METHOD_9955.html */
func TestReset(t *testing.T) {
	err := replication.Reset()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SAVE_METHOD_3588.html */
func TestSave(t *testing.T) {
	err := replication.Save()
	require.Nil(t, err)
}
