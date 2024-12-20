/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_NOTESCALENDAR_CLASS.html */
package notescalendar_test

import (
	"testing"

	"github.com/monstermichl/domigo"
	testhelpers "github.com/monstermichl/domigo/test/helpers"

	"github.com/stretchr/testify/require"
)

var start domigo.NotesDateTime
var end domigo.NotesDateTime
var calendar domigo.NotesCalendar

/* https://pkg.go.dev/testing#hdr-Main */
func TestMain(m *testing.M) {
	testhelpers.Initialize(func(session domigo.NotesSession, db domigo.NotesDatabase) (string, error) {
		var err error

		start, err = session.CreateDateTime("Yesterday")
		defer start.Release()

		if err != nil {
			return "Start date could not be created", err
		}

		end, err = session.CreateDateTime("Today")
		defer end.Release()

		if err != nil {
			return "End date could not be created", err
		}

		calendar, err = session.GetCalendar(db)
		defer calendar.Release()

		if err != nil {
			return "NotesCalendar could not be created", err
		}
		m.Run()
		return "", nil
	})
}

/* --------------------------------- Properties --------------------------------- */
/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ENTRIESPROCESSED_PROPERTY_CAL.html */
func TestEntriesProcessed(t *testing.T) {
	_, err := calendar.EntriesProcessed()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_READRANGEMASK1_PROPERTY_CAL.html */
func TestReadRangeMask1(t *testing.T) {
	_, err := calendar.ReadRangeMask1()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_READRANGEMASK1_PROPERTY_CAL.html */
func TestSetReadRangeMask1(t *testing.T) {
	s, err := calendar.ReadRangeMask1()
	require.Nil(t, err)

	err = calendar.SetReadRangeMask1(s)
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_READRANGEMASK2_PROPERTY_CAL.html */
func TestReadRangeMask2(t *testing.T) {
	_, err := calendar.ReadRangeMask2()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_READRANGEMASK2_PROPERTY_CAL.html */
func TestSetReadRangeMask2(t *testing.T) {
	s, err := calendar.ReadRangeMask2()
	require.Nil(t, err)

	err = calendar.SetReadRangeMask2(s)
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_READXLOTUSPROPSOUTPUTLEVEL_PROPERTY_CAL.html */
func TestReadXLotusPropsOutputLevel(t *testing.T) {
	_, err := calendar.ReadXLotusPropsOutputLevel()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_READXLOTUSPROPSOUTPUTLEVEL_PROPERTY_CAL.html */
func TestSetReadXLotusPropsOutputLevel(t *testing.T) {
	s, err := calendar.ReadXLotusPropsOutputLevel()
	require.Nil(t, err)

	err = calendar.SetReadXLotusPropsOutputLevel(s)
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_UNTILTIME_PROPERTY_CAL.html */
func TestUntilTime(t *testing.T) {
	_, err := calendar.UntilTime()
	require.Nil(t, err)
}

/* --------------------------------- Methods ------------------------------------ */
/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_CREATEENTRY_METHOD_CAL.html */
func TestCreateEntry(t *testing.T) {
	_, err := calendar.CreateEntry("BEGIN:VCALENDAR;PRODID:-//xyz Corp//NONSGML PDA Calendar Version 1.0//EN;VERSION:2.0;EGIN:VEVENT;DTSTAMP:19960704T120000Z;UID:uid1@example.com;ORGANIZER:mailto:jsmith@example.com;DTSTART:19960918T143000Z;DTEND:19960920T220000Z;STATUS:CONFIRMED;CATEGORIES:CONFERENCE;SUMMARY:Networld+Interop Conference;DESCRIPTION:Networld+Interop Conference and Exhibit\nAtlanta World Congress Center\nAtlanta, Georgia;END:VEVENT;END:VCALENDAR")
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GETENTRIES_METHOD_CAL.html */
func TestGetEntries(t *testing.T) {
	_, err := calendar.GetEntries(start, end)
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GETENTRY_METHOD_CAL.html */
func TestGetEntry(t *testing.T) {
	_, err := calendar.GetEntry("test")
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GETENTRYBYNOTEID_METHOD_CAL.html */
func TestGetEntryByNoteID(t *testing.T) {
	_, err := calendar.GetEntryByNoteID("test")
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GETENTRYBYUNID_METHOD_CAL.html */
func TestGetEntryByUNID(t *testing.T) {
	_, err := calendar.GetEntryByUNID("test")
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GETNEWINVITATIONS_METHOD_CAL.html */
func TestGetNewInvitations(t *testing.T) {
	_, err := calendar.GetNewInvitations()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GETNOTICEBYUNID_METHOD_CAL.html */
func TestGetNoticeByUNID(t *testing.T) {
	_, err := calendar.GetNoticeByUNID("test")
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GETRECURRENCEID_METHOD_CAL.html */
func TestGetRecurrenceID(t *testing.T) {
	_, err := calendar.GetRecurrenceID()
	require.Nil(t, err)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_READRANGE_METHOD_CAL.html */
func TestReadRange(t *testing.T) {
	_, err := calendar.ReadRange(start, end)
	require.Nil(t, err)
}
