/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_NOTESCALENDARNOTICE_CLASS.html */
package notescalendarnotice_test

import (
	"fmt"
	"testing"

	"github.com/monstermichl/domigo"
	testhelpers "github.com/monstermichl/domigo/test/helpers"

	"github.com/stretchr/testify/require"
)

var calendarnotice domigo.NotesCalendarNotice

/* https://pkg.go.dev/testing#hdr-Main */
func TestMain(m *testing.M) {
	var info string

	session, err := domigo.Initialize()
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

	calendar, err := session.GetCalendar(db)
	defer calendar.Release()

	if err != nil {
		info = "Calendar could not be retrieved"
		return
	}

	calendarentry, err := calendar.CreateEntry("BEGIN:VCALENDAR;PRODID:-//xyz Corp//NONSGML PDA Calendar Version 1.0//EN;VERSION:2.0;EGIN:VEVENT;DTSTAMP:19960704T120000Z;UID:uid1@example.com;ORGANIZER:mailto:jsmith@example.com;DTSTART:19960918T143000Z;DTEND:19960920T220000Z;STATUS:CONFIRMED;CATEGORIES:CONFERENCE;SUMMARY:Networld+Interop Conference;DESCRIPTION:Networld+Interop Conference and Exhibit\nAtlanta World Congress Center\nAtlanta, Georgia;END:VEVENT;END:VCALENDAR")
	defer calendarentry.Release()

	if err != nil {
		info = "Calendar entry could not be created"
		return
	}
	notices, err := calendarentry.GetNotices()

	for _, n := range notices {
		defer n.Release()
	}

	if err != nil {
		info = "Notices could not be retrieved"
		return
	} else if len(notices) == 0 {
		info = "No notices found"
		return
	}
	calendarnotice = notices[0]

	m.Run()
}

/* --------------------------------- Properties --------------------------------- */
/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_NOTEID_PROPERTY_CALNOTICE.html */
func TestNoteID(t *testing.T) {
	_, err := calendarnotice.NoteID()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_UNID_PROPERTY_CALNOTICE.html */
func TestUNID(t *testing.T) {
	_, err := calendarnotice.UNID()
	require.Nil(t, err)
}

/* --------------------------------- Methods ------------------------------------ */
/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ACCEPTCOUNTER_METHOD_CALNOTICE.html */
func TestAcceptCounter(t *testing.T) {
	err := calendarnotice.AcceptCounter("Accepting")
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_COUNTER_METHOD_CALNOTICE.html */
func TestCounter(t *testing.T) {
	err := calendarnotice.Counter("Counting")
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_DECLINE_METHOD_CALNOTICE.html */
func TestDecline(t *testing.T) {
	err := calendarnotice.Decline("Can't participate")
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_DECLINECOUNTER_METHOD_CALNOTICE.html */
func TestDeclineCounter(t *testing.T) {
	err := calendarnotice.DeclineCounter("Declining")
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_DELEGATE_METHOD_CALNOTICE.html */
func TestDelegate(t *testing.T) {
	err := calendarnotice.Delegate("Please join for me", "test@test.com")
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GETASDOCUMENT_METHOD_CALNOTICE.html */
func TestGetAsDocument(t *testing.T) {
	_, err := calendarnotice.GetAsDocument()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GETOUTSTANDINGINVITATIONS_METHOD_CALNOTICE.html */
func TestGetOutstandingInvitations(t *testing.T) {
	_, err := calendarnotice.GetOutstandingInvitations()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_READ_METHOD_CALNOTICE.html */
func TestRead(t *testing.T) {
	_, err := calendarnotice.Read()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_REMOVECANCELLED_METHOD_CALNOTICE.html */
func TestRemoveCancelled(t *testing.T) {
	err := calendarnotice.RemoveCancelled()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_REQUESTINFO_METHOD_CALNOTICE.html */
func TestRequestInfo(t *testing.T) {
	err := calendarnotice.RequestInfo("Need more info")
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SENDUPDATEDINFO_METHOD_CALNOTICE.html */
func TestSendUpdatedInfo(t *testing.T) {
	err := calendarnotice.SendUpdatedInfo("Updating stuff")
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_TENTATIVELYACCEPT_METHOD_CALNOTICE.html */
func TestTentativelyAccept(t *testing.T) {
	err := calendarnotice.TentativelyAccept("As of now")
	require.Nil(t, err)
}
