/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_NOTESCALENDARENTRY_CLASS.html */
package notescalendarentry_test

import (
	"testing"

	"github.com/monstermichl/domigo"
	testhelpers "github.com/monstermichl/domigo/test/helpers"

	"github.com/stretchr/testify/require"
)

var session domigo.NotesSession
var calendarentry domigo.NotesCalendarEntry

/* https://pkg.go.dev/testing#hdr-Main */
func TestMain(m *testing.M) {
	testhelpers.Initialize(func(sessionTemp domigo.NotesSession, db domigo.NotesDatabase) (string, error) {
		var err error

		session = sessionTemp
		calendar, err := session.GetCalendar(db)
		defer calendar.Release()

		if err != nil {
			return "Calendar could not be retrieved", err
		}

		calendarentry, err = calendar.CreateEntry("BEGIN:VCALENDAR;PRODID:-//xyz Corp//NONSGML PDA Calendar Version 1.0//EN;VERSION:2.0;EGIN:VEVENT;DTSTAMP:19960704T120000Z;UID:uid1@example.com;ORGANIZER:mailto:jsmith@example.com;DTSTART:19960918T143000Z;DTEND:19960920T220000Z;STATUS:CONFIRMED;CATEGORIES:CONFERENCE;SUMMARY:Networld+Interop Conference;DESCRIPTION:Networld+Interop Conference and Exhibit\nAtlanta World Congress Center\nAtlanta, Georgia;END:VEVENT;END:VCALENDAR")
		defer calendarentry.Release()

		if err != nil {
			return "NotesCalendarEntry could not be created", err
		}
		m.Run()
		return "", nil
	})
}

/* --------------------------------- Properties --------------------------------- */
/* --------------------------------- Methods ------------------------------------ */
/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ADDINVITEES_METHOD_CALENTRY.html */
func TestAddInvitees(t *testing.T) {
	err := calendarentry.AddInvitees([]domigo.String{"test@test.com"}, []domigo.String{}, []domigo.String{})
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_CANCEL_METHOD_CALENTRY.html */
func TestCancel(t *testing.T) {
	err := calendarentry.Cancel("Meeting canceled")
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_COUNTER_METHOD_CALENTRY.html */
func TestCounter(t *testing.T) {
	start, err := session.CreateDateTime("Yesterday")
	require.Nil(t, err)

	end, err := session.CreateDateTime("Today")
	require.Nil(t, err)

	err = calendarentry.Counter("Counting", start, end)
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_DECLINE_METHOD_CALENTRY.html */
func TestDecline(t *testing.T) {
	err := calendarentry.Decline("Not attending")
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_DELEGATE_METHOD_CALENTRY.html */
func TestDelegate(t *testing.T) {
	err := calendarentry.Delegate("I'm delegating, bruh", "othertest@test.com")
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GETASDOCUMENT_METHOD_CALENTRY.html */
func TestGetAsDocument(t *testing.T) {
	_, err := calendarentry.GetAsDocument()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GETNOTICES_METHOD_CALENTRY.html */
func TestGetNotices(t *testing.T) {
	_, err := calendarentry.GetNotices()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_MODIFYINVITEES_METHOD_CALENTRY.html */
func TestModifyInvitees(t *testing.T) {
	err := calendarentry.ModifyInvitees([]domigo.String{"test2@test.com"}, []domigo.String{}, []domigo.String{}, []domigo.String{})
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_READ_METHOD_CALENTRY.html */
func TestRead(t *testing.T) {
	_, err := calendarentry.Read()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_REMOVEINVITEES_METHOD_CALENTRY.html */
func TestRemoveInvitees(t *testing.T) {
	err := calendarentry.RemoveInvitees([]domigo.String{"noone@test.com"})
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_REQUESTINFO_METHOD_CALENTRY.html */
func TestRequestInfo(t *testing.T) {
	err := calendarentry.RequestInfo("Give me some info")
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_TENTATIVELYACCEPT_METHOD_CALENTRY.html */
func TestTentativelyAccept(t *testing.T) {
	err := calendarentry.TentativelyAccept("Don't know yet")
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_UPDATE_METHOD_CALENTRY.html */
func TestUpdate(t *testing.T) {
	err := calendarentry.Update("BEGIN:VCALENDAR;PRODID:-//xyz Corp//NONSGML PDA Calendar Version 1.0//EN;VERSION:2.0;EGIN:VEVENT;DTSTAMP:19960704T120000Z;UID:uid1@example.com;ORGANIZER:mailto:jsmith@example.com;DTSTART:19960918T143000Z;DTEND:19960920T220000Z;STATUS:CONFIRMED;CATEGORIES:CONFERENCE;SUMMARY:Networld+Interop Conference;DESCRIPTION:Networld+Interop Conference and Exhibit\nAtlanta World Congress Center\nAtlanta, Georgia;END:VEVENT;END:VCALENDAR")
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_REMOVE_METHOD_CALENTRY.html */
func TestRemove(t *testing.T) {
	err := calendarentry.Remove("Removing whatever")
	require.Nil(t, err)
}
