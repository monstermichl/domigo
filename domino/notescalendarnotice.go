/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_NOTESCALENDARNOTICE_CLASS.html */
package domigo

import (
	"github.com/monstermichl/domigo/domino/com"
	"github.com/monstermichl/domigo/helpers"

	ole "github.com/go-ole/go-ole"
)

type NotesCalendarNotice struct {
	NotesStruct
}

func NewNotesCalendarNotice(dispatchPtr *ole.IDispatch) NotesCalendarNotice {
	return NotesCalendarNotice{NewNotesStruct(dispatchPtr)}
}

/* --------------------------------- Properties --------------------------------- */
/* TODO: Add OverwriteCheckEnabled property. */

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_NOTEID_PROPERTY_CALNOTICE.html */
func (c NotesCalendarNotice) NoteID() (String, error) {
	val, err := c.com().GetProperty("NoteID")
	return helpers.CastValue[String](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_UNID_PROPERTY_CALNOTICE.html */
func (c NotesCalendarNotice) UNID() (String, error) {
	val, err := c.com().GetProperty("UNID")
	return helpers.CastValue[String](val), err
}

/* --------------------------------- Methods ------------------------------------ */
/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ACCEPTCOUNTER_METHOD_CALNOTICE.html */
func (c NotesCalendarNotice) AcceptCounter(comments String) error {
	_, err := c.com().CallMethod("AcceptCounter", comments)
	return err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_COUNTER_METHOD_CALNOTICE.html */
type notesCalendarNoticeCounterParams struct {
	start *NotesDateTime
	end   *NotesDateTime
}

type notesCalendarNoticeCounterParam func(*notesCalendarNoticeCounterParams)

func WithNotesCalendarNoticeCounterStart(start NotesDateTime) notesCalendarNoticeCounterParam {
	return func(c *notesCalendarNoticeCounterParams) {
		c.start = &start
	}
}

func WithNotesCalendarNoticeCounterEnd(end NotesDateTime) notesCalendarNoticeCounterParam {
	return func(c *notesCalendarNoticeCounterParams) {
		c.end = &end
	}
}

func (c NotesCalendarNotice) Counter(comments String, params ...notesCalendarNoticeCounterParam) error {
	paramsStruct := &notesCalendarNoticeCounterParams{}
	paramsOrdered := []interface{}{comments}

	for _, p := range params {
		p(paramsStruct)
	}

	if paramsStruct.start != nil {
		paramsOrdered = append(paramsOrdered, *paramsStruct.start.com().Dispatch())
		if paramsStruct.end != nil {
			paramsOrdered = append(paramsOrdered, *paramsStruct.end.com().Dispatch())
		}
	}
	_, err := c.com().CallMethod("Counter", paramsOrdered...)
	return err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_DECLINE_METHOD_CALNOTICE.html */
type notesCalendarNoticeDeclineParams struct {
	keepInformed *Boolean
}

type notesCalendarNoticeDeclineParam func(*notesCalendarNoticeDeclineParams)

func WithNotesCalendarNoticeDeclineEnd(keepInformed Boolean) notesCalendarNoticeDeclineParam {
	return func(c *notesCalendarNoticeDeclineParams) {
		c.keepInformed = &keepInformed
	}
}

func (c NotesCalendarNotice) Decline(comments String, params ...notesCalendarNoticeDeclineParam) error {
	paramsStruct := &notesCalendarNoticeDeclineParams{}
	paramsOrdered := []interface{}{comments}

	for _, p := range params {
		p(paramsStruct)
	}

	if paramsStruct.keepInformed != nil {
		paramsOrdered = append(paramsOrdered, *paramsStruct.keepInformed)
	}
	_, err := c.com().CallMethod("Decline", paramsOrdered...)
	return err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_DECLINECOUNTER_METHOD_CALNOTICE.html */
func (c NotesCalendarNotice) DeclineCounter(comments String) error {
	_, err := c.com().CallMethod("DeclineCounter", comments)
	return err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_DELEGATE_METHOD_CALNOTICE.html */
type notesCalendarNoticeDelegateParams struct {
	keepInformed *Boolean
}

type notesCalendarNoticeDelegateParam func(*notesCalendarNoticeDelegateParams)

func WithNotesCalendarNoticeDelegateEnd(keepInformed Boolean) notesCalendarNoticeDelegateParam {
	return func(c *notesCalendarNoticeDelegateParams) {
		c.keepInformed = &keepInformed
	}
}

func (c NotesCalendarNotice) Delegate(commentsToOrganizer String, delegateTo String, params ...notesCalendarNoticeDelegateParam) error {
	paramsStruct := &notesCalendarNoticeDelegateParams{}
	paramsOrdered := []interface{}{commentsToOrganizer, delegateTo}

	for _, p := range params {
		p(paramsStruct)
	}

	if paramsStruct.keepInformed != nil {
		paramsOrdered = append(paramsOrdered, *paramsStruct.keepInformed)
	}
	_, err := c.com().CallMethod("Delegate", paramsOrdered...)
	return err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GETASDOCUMENT_METHOD_CALNOTICE.html */
func (c NotesCalendarNotice) GetAsDocument() (NotesDocument, error) {
	dispatchPtr, err := c.com().CallObjectMethod("GetAsDocument")
	return NewNotesDocument(dispatchPtr), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GETOUTSTANDINGINVITATIONS_METHOD_CALNOTICE.html */
func (c NotesCalendarNotice) GetOutstandingInvitations() ([]NotesCalendarNotice, error) {
	return com.CallObjectArrayMethod(c.com(), NewNotesCalendarNotice, "GetOutstandingInvitations")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_READ_METHOD_CALNOTICE.html */
func (c NotesCalendarNotice) Read() (String, error) {
	val, err := c.com().CallMethod("Read")
	return helpers.CastValue[String](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_REMOVECANCELLED_METHOD_CALNOTICE.html */
func (c NotesCalendarNotice) RemoveCancelled() error {
	_, err := c.com().CallMethod("RemoveCancelled")
	return err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_REQUESTINFO_METHOD_CALNOTICE.html */
func (c NotesCalendarNotice) RequestInfo(comments String) error {
	_, err := c.com().CallMethod("RequestInfo", comments)
	return err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SENDUPDATEDINFO_METHOD_CALNOTICE.html */
func (c NotesCalendarNotice) SendUpdatedInfo(comments String) error {
	_, err := c.com().CallMethod("SendUpdatedInfo", comments)
	return err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_TENTATIVELYACCEPT_METHOD_CALNOTICE.html */
func (c NotesCalendarNotice) TentativelyAccept(comments String) error {
	_, err := c.com().CallMethod("TentativelyAccept", comments)
	return err
}
