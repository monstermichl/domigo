/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_NOTESAGENT_CLASS.html */
package notesagent_test

import (
	"fmt"
	"testing"

	"github.com/monstermichl/domigo/domino/notesagent"
	"github.com/monstermichl/domigo/domino/notessession"
	testhelpers "github.com/monstermichl/domigo/test/helpers"

	"github.com/stretchr/testify/require"
)

var agent notesagent.NotesAgent

/* https://pkg.go.dev/testing#hdr-Main */
func TestMain(m *testing.M) {
	var info string

	session, err := notessession.Initialize()
	defer session.Release()

	defer func() {
		fmt.Println(err)
		fmt.Println(info)
	}()

	if err != nil {
		info = "Session could not be initialized"
		return
	}

	db, err := testhelpers.CreateTestDatabase(session)
	defer db.Release()
	defer db.Remove()

	if err != nil {
		info = "Database could not be created"
		return
	}
	err = db.SetIsDesignLockingEnabled(true)

	if err != nil {
		info = "Design locking could not be enabled"
		return
	}
	agents, _ := db.Agents()

	if err != nil {
		info = "Agents could not be retrieved"
		return
	}

	if len(agents) == 0 {
		info = "No agents found"
		return
	}

	for _, a := range agents {
		defer a.Release()
	}
	agent = agents[0]

	m.Run()
}

/* --------------------------------- Properties --------------------------------- */
/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_COMMENT_PROPERTY.html */
func TestComment(t *testing.T) {
	_, err := agent.Comment()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_COMMONOWNER_PROPERTY.html */
func TestCommonOwner(t *testing.T) {
	_, err := agent.CommonOwner()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_HASRUNSINCEMODIFIED_PROPERTY_AGENT.html */
func TestHasRunSinceModified(t *testing.T) {
	_, err := agent.HasRunSinceModified()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_HTTPURL_PROPERTY_AGENT_COM.html */
func TestHttpURL(t *testing.T) {
	_, err := agent.HttpURL()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISACTIVATABLE_PROPERTY_AGENT.html */
func TestIsActivatable(t *testing.T) {
	_, err := agent.IsActivatable()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISENABLED_PROPERTY.html */
func TestIsEnabled(t *testing.T) {
	_, err := agent.IsEnabled()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISENABLED_PROPERTY.html */
func TestSetIsEnabled(t *testing.T) {
	s, err := agent.IsEnabled()
	require.Nil(t, err)

	err = agent.SetIsEnabled(s)
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISNOTESAGENT_PROPERTY_4410_ABOUT.html */
func TestIsNotesAgent(t *testing.T) {
	_, err := agent.IsNotesAgent()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISPUBLIC_PROPERTY.html */
func TestIsPublic(t *testing.T) {
	_, err := agent.IsPublic()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISWEBAGENT_PROPERTY_5149_ABOUT.html */
func TestIsWebAgent(t *testing.T) {
	_, err := agent.IsWebAgent()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_LASTRUN_PROPERTY_AGENT.html */
func TestLastRun(t *testing.T) {
	_, err := agent.LastRun()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_LOCKHOLDERS_PROPERTY_AGENT.html */
func TestLockHolders(t *testing.T) {
	_, err := agent.LockHolders()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_NAME_PROPERTY_AGENT.html */
func TestName(t *testing.T) {
	_, err := agent.Name()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_NOTESURL_PROPERTY_AGENT_COM.html */
func TestNotesURL(t *testing.T) {
	_, err := agent.NotesURL()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ONBEHALFOF_PROPERTY_AGENT.html */
func TestOnBehalfOf(t *testing.T) {
	_, err := agent.OnBehalfOf()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_OWNER_PROPERTY.html */
func TestOwner(t *testing.T) {
	_, err := agent.Owner()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_PARAMETERDOCID_PROPERTY_AGENT.html */
func TestParameterDocID(t *testing.T) {
	_, err := agent.ParameterDocID()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_PROHIBITDESIGNUPDATE_PROPERTY.html */
func TestProhibitDesignUpdate(t *testing.T) {
	_, err := agent.ProhibitDesignUpdate()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_PROHIBITDESIGNUPDATE_PROPERTY.html */
func TestSetProhibitDesignUpdate(t *testing.T) {
	s, err := agent.ProhibitDesignUpdate()
	require.Nil(t, err)

	err = agent.SetProhibitDesignUpdate(s)
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_QUERY_PROPERTY_AGENT.html */
func TestQuery(t *testing.T) {
	_, err := agent.Query()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SERVERNAME_PROPERTY.html */
func TestServerName(t *testing.T) {
	_, err := agent.ServerName()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SERVERNAME_PROPERTY.html */
func TestSetServerName(t *testing.T) {
	s, err := agent.ServerName()
	require.Nil(t, err)

	err = agent.SetServerName(s)
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_TARGET_PROPERTY_1466_ABOUT.html */
func TestTarget(t *testing.T) {
	_, err := agent.Target()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_TRIGGER_PROPERTY_9708_ABOUT.html */
func TestTrigger(t *testing.T) {
	_, err := agent.Trigger()
	require.Nil(t, err)
}

/* --------------------------------- Methods ------------------------------------ */
/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_LOCK_METHOD_AGENT.html */
func TestLock(t *testing.T) {
	_, err := agent.Lock()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_LOCKPROVISIONAL_METHOD_AGENT.html */
func TestLockProvisional(t *testing.T) {
	_, err := agent.LockProvisional()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_RUN_METHOD_6415.html */
func TestRun(t *testing.T) {
	_, err := agent.Run()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_RUNONSERVER_METHOD_5924_ABOUT.html */
func TestRunOnServer(t *testing.T) {
	_, err := agent.RunOnServer()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SAVE_METHOD_AGENT.html */
func TestSave(t *testing.T) {
	err := agent.Save()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_UNLOCK_METHOD_AGENT.html */
func TestUnLock(t *testing.T) {
	err := agent.UnLock()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_REMOVE_METHOD_AGENT.html */
func TestRemove(t *testing.T) {
	err := agent.Remove()
	require.Nil(t, err)
}
