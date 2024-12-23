/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_NOTESCALENDAR_CLASS.html */
package domigo

import (
	ole "github.com/go-ole/go-ole"
	"github.com/monstermichl/domigo/internal/com"
)

type NotesCalendar struct {
	NotesStruct
}

func NewNotesCalendar(dispatchPtr *ole.IDispatch) NotesCalendar {
	return NotesCalendar{NewNotesStruct(dispatchPtr)}
}

/* --------------------------------- Properties --------------------------------- */
/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ENTRIESPROCESSED_PROPERTY_CAL.html */
func (c NotesCalendar) EntriesProcessed() (Long, error) {
	return getComProperty[Long](c, "EntriesProcessed")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_READRANGEMASK1_PROPERTY_CAL.html */
func (c NotesCalendar) ReadRangeMask1() (Long, error) {
	return getComProperty[Long](c, "ReadRangeMask1")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_READRANGEMASK1_PROPERTY_CAL.html */
func (c NotesCalendar) SetReadRangeMask1(v Long) error {
	return putComProperty(c, "ReadRangeMask1", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_READRANGEMASK2_PROPERTY_CAL.html */
func (c NotesCalendar) ReadRangeMask2() (Long, error) {
	return getComProperty[Long](c, "ReadRangeMask2")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_READRANGEMASK2_PROPERTY_CAL.html */
func (c NotesCalendar) SetReadRangeMask2(v Long) error {
	return putComProperty(c, "ReadRangeMask2", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_READXLOTUSPROPSOUTPUTLEVEL_PROPERTY_CAL.html */
func (c NotesCalendar) ReadXLotusPropsOutputLevel() (Integer, error) {
	return getComProperty[Integer](c, "ReadXLotusPropsOutputLevel")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_READXLOTUSPROPSOUTPUTLEVEL_PROPERTY_CAL.html */
func (c NotesCalendar) SetReadXLotusPropsOutputLevel(v Integer) error {
	return putComProperty(c, "ReadXLotusPropsOutputLevel", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_UNTILTIME_PROPERTY_CAL.html */
func (c NotesCalendar) UntilTime() (NotesDateTime, error) {
	return getComObjectProperty(c, NewNotesDateTime, "UntilTime")
}

/* --------------------------------- Methods ------------------------------------ */
/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_CREATEENTRY_METHOD_CAL.html */
type notesCalendarCreateEntryParams struct {
	flags *Integer
}

type notesCalendarCreateEntryParam func(*notesCalendarCreateEntryParams)

func WithNotesCalendarCreateEntryFlags(flags Integer) notesCalendarCreateEntryParam {
	return func(c *notesCalendarCreateEntryParams) {
		c.flags = &flags
	}
}

func (c NotesCalendar) CreateEntry(icalEntry String, params ...notesCalendarCreateEntryParam) (NotesCalendarEntry, error) {
	paramsStruct := &notesCalendarCreateEntryParams{}
	paramsOrdered := []interface{}{icalEntry}

	for _, p := range params {
		p(paramsStruct)
	}

	if paramsStruct.flags != nil {
		paramsOrdered = append(paramsOrdered, *paramsStruct.flags)
	}
	return callComObjectMethod(c, NewNotesCalendarEntry, "CreateEntry", paramsOrdered...)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GETENTRIES_METHOD_CAL.html */
type notesCalendarGetEntriesParams struct {
	skipCount *Long
	maxReturn *Long
}

type notesCalendarGetEntriesParam func(*notesCalendarGetEntriesParams)

func WithNotesCalendarGetEntriesSkipCount(skipCount Long) notesCalendarGetEntriesParam {
	return func(c *notesCalendarGetEntriesParams) {
		c.skipCount = &skipCount
	}
}

func WithNotesCalendarGetEntriesMaxReturn(maxReturn Long) notesCalendarGetEntriesParam {
	return func(c *notesCalendarGetEntriesParams) {
		c.maxReturn = &maxReturn
	}
}

func (c NotesCalendar) GetEntries(start NotesDateTime, end NotesDateTime, params ...notesCalendarGetEntriesParam) (NotesCalendarEntry, error) {
	paramsStruct := &notesCalendarGetEntriesParams{}
	paramsOrdered := []interface{}{start, end}

	for _, p := range params {
		p(paramsStruct)
	}

	if paramsStruct.skipCount != nil {
		paramsOrdered = append(paramsOrdered, *paramsStruct.skipCount)
		if paramsStruct.maxReturn != nil {
			paramsOrdered = append(paramsOrdered, *paramsStruct.maxReturn)
		}
	}
	return callComObjectMethod(c, NewNotesCalendarEntry, "GetEntries", paramsOrdered...)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GETENTRY_METHOD_CAL.html */
func (c NotesCalendar) GetEntry(uid String) (NotesCalendarEntry, error) {
	return callComObjectMethod(c, NewNotesCalendarEntry, "GetEntry", uid)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GETENTRYBYNOTEID_METHOD_CAL.html */
func (c NotesCalendar) GetEntryByNoteID(noteid String) (NotesCalendarEntry, error) {
	return callComObjectMethod(c, NewNotesCalendarEntry, "GetEntryByNoteID", noteid)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GETENTRYBYUNID_METHOD_CAL.html */
func (c NotesCalendar) GetEntryByUNID(unid String) (NotesCalendarEntry, error) {
	return callComObjectMethod(c, NewNotesCalendarEntry, "GetEntryByUNID", unid)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GETNEWINVITATIONS_METHOD_CAL.html */
type notesCalendarGetNewInvitationsParams struct {
	start *NotesDateTime
	since *NotesDateTime
}

type notesCalendarGetNewInvitationsParam func(*notesCalendarGetNewInvitationsParams)

func WithNotesCalendarGetNewInvitationsStart(start NotesDateTime) notesCalendarGetNewInvitationsParam {
	return func(c *notesCalendarGetNewInvitationsParams) {
		c.start = &start
	}
}

func WithNotesCalendarGetNewInvitationsSince(since NotesDateTime) notesCalendarGetNewInvitationsParam {
	return func(c *notesCalendarGetNewInvitationsParams) {
		c.since = &since
	}
}

func (c NotesCalendar) GetNewInvitations(params ...notesCalendarGetNewInvitationsParam) ([]NotesCalendarNotice, error) {
	paramsStruct := &notesCalendarGetNewInvitationsParams{}
	paramsOrdered := []interface{}{}

	for _, p := range params {
		p(paramsStruct)
	}

	if paramsStruct.start != nil {
		paramsOrdered = append(paramsOrdered, *paramsStruct.start)
		if paramsStruct.since != nil {
			paramsOrdered = append(paramsOrdered, *paramsStruct.since)
		}
	}
	return com.CallObjectArrayMethod(c.com(), NewNotesCalendarNotice, "GetNewInvitations", paramsOrdered...)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GETNOTICEBYUNID_METHOD_CAL.html */
func (c NotesCalendar) GetNoticeByUNID(unid String) (NotesCalendarNotice, error) {
	return callComObjectMethod(c, NewNotesCalendarNotice, "GetNoticeByUNID", unid)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GETRECURRENCEID_METHOD_CAL.html */
func (c NotesCalendar) GetRecurrenceID() (String, error) {
	return callComMethod[String](c, "GetRecurrenceID")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_READRANGE_METHOD_CAL.html */
type notesCalendarReadRangeParams struct {
	skipCount *Long
	maxReturn *Long
}

type notesCalendarReadRangeParam func(*notesCalendarReadRangeParams)

func WithNotesCalendarReadRangeSkipCount(skipCount Long) notesCalendarReadRangeParam {
	return func(c *notesCalendarReadRangeParams) {
		c.skipCount = &skipCount
	}
}

func WithNotesCalendarReadRangeMaxReturn(maxReturn Long) notesCalendarReadRangeParam {
	return func(c *notesCalendarReadRangeParams) {
		c.maxReturn = &maxReturn
	}
}

func (c NotesCalendar) ReadRange(start NotesDateTime, end NotesDateTime, params ...notesCalendarReadRangeParam) (String, error) {
	paramsStruct := &notesCalendarReadRangeParams{}
	paramsOrdered := []interface{}{start, end}

	for _, p := range params {
		p(paramsStruct)
	}

	if paramsStruct.skipCount != nil {
		paramsOrdered = append(paramsOrdered, *paramsStruct.skipCount)
		if paramsStruct.maxReturn != nil {
			paramsOrdered = append(paramsOrdered, *paramsStruct.maxReturn)
		}
	}
	return callComMethod[String](c, "ReadRange", paramsOrdered...)
}
