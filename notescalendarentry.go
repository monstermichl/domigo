/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_NOTESCALENDARENTRY_CLASS.html */
package domigo

import (
	"github.com/monstermichl/domigo/internal/com"
	"github.com/monstermichl/domigo/internal/helpers"

	ole "github.com/go-ole/go-ole"
)

type NotesCalendarEntry struct {
	NotesStruct
}

func NewNotesCalendarEntry(dispatchPtr *ole.IDispatch) NotesCalendarEntry {
	return NotesCalendarEntry{NewNotesStruct(dispatchPtr)}
}

/* --------------------------------- Properties --------------------------------- */
/* TODO: Add UID property. */

/* --------------------------------- Methods ------------------------------------ */
/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ADDINVITEES_METHOD_CALENTRY.html */
type notesCalendarEntryAddInviteesParams struct {
	comments *String
	flags    *Integer
	scope    *Integer
	recurId  *String
}

type notesCalendarEntryAddInviteesParam func(*notesCalendarEntryAddInviteesParams)

func WithNotesCalendarEntryAddInviteesComments(comments String) notesCalendarEntryAddInviteesParam {
	return func(c *notesCalendarEntryAddInviteesParams) {
		c.comments = &comments
	}
}

func WithNotesCalendarEntryAddInviteesFlags(flags Integer) notesCalendarEntryAddInviteesParam {
	return func(c *notesCalendarEntryAddInviteesParams) {
		c.flags = &flags
	}
}

func WithNotesCalendarEntryAddInviteesScope(scope Integer) notesCalendarEntryAddInviteesParam {
	return func(c *notesCalendarEntryAddInviteesParams) {
		c.scope = &scope
	}
}

func WithNotesCalendarEntryAddInviteesRecurId(recurId String) notesCalendarEntryAddInviteesParam {
	return func(c *notesCalendarEntryAddInviteesParams) {
		c.recurId = &recurId
	}
}

func (c NotesCalendarEntry) AddInvitees(requiredNames []String, optionalNames []String, fyiNames []String, params ...notesCalendarEntryAddInviteesParam) error {
	paramsStruct := &notesCalendarEntryAddInviteesParams{}
	paramsOrdered := []interface{}{requiredNames, optionalNames, fyiNames}

	for _, p := range params {
		p(paramsStruct)
	}

	if paramsStruct.comments != nil {
		paramsOrdered = append(paramsOrdered, *paramsStruct.comments)
		if paramsStruct.flags != nil {
			paramsOrdered = append(paramsOrdered, *paramsStruct.flags)
			if paramsStruct.scope != nil {
				paramsOrdered = append(paramsOrdered, *paramsStruct.scope)
				if paramsStruct.recurId != nil {
					paramsOrdered = append(paramsOrdered, *paramsStruct.recurId)
				}
			}
		}
	}
	_, err := callComMethod(c, "AddInvitees", paramsOrdered...)
	return err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_CANCEL_METHOD_CALENTRY.html */
type notesCalendarEntryCancelParams struct {
	scope   *Integer
	recurId *String
}

type notesCalendarEntryCancelParam func(*notesCalendarEntryCancelParams)

func WithNotesCalendarEntryCancelScope(scope Integer) notesCalendarEntryCancelParam {
	return func(c *notesCalendarEntryCancelParams) {
		c.scope = &scope
	}
}

func WithNotesCalendarEntryCancelRecurId(recurId String) notesCalendarEntryCancelParam {
	return func(c *notesCalendarEntryCancelParams) {
		c.recurId = &recurId
	}
}

func (c NotesCalendarEntry) Cancel(comments String, params ...notesCalendarEntryCancelParam) error {
	paramsStruct := &notesCalendarEntryCancelParams{}
	paramsOrdered := []interface{}{comments}

	for _, p := range params {
		p(paramsStruct)
	}

	if paramsStruct.scope != nil {
		paramsOrdered = append(paramsOrdered, *paramsStruct.scope)
		if paramsStruct.recurId != nil {
			paramsOrdered = append(paramsOrdered, *paramsStruct.recurId)
		}
	}
	_, err := callComMethod(c, "Cancel", paramsOrdered...)
	return err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_COUNTER_METHOD_CALENTRY.html */
type notesCalendarEntryCounterParams struct {
	keepPlaceholder *Boolean
	scope           *Integer
	recurId         *String
}

type notesCalendarEntryCounterParam func(*notesCalendarEntryCounterParams)

func WithNotesCalendarEntryCounterKeepPlaceholder(keepPlaceholder Boolean) notesCalendarEntryCounterParam {
	return func(c *notesCalendarEntryCounterParams) {
		c.keepPlaceholder = &keepPlaceholder
	}
}

func WithNotesCalendarEntryCounterScope(scope Integer) notesCalendarEntryCounterParam {
	return func(c *notesCalendarEntryCounterParams) {
		c.scope = &scope
	}
}

func WithNotesCalendarEntryCounterRecurId(recurId String) notesCalendarEntryCounterParam {
	return func(c *notesCalendarEntryCounterParams) {
		c.recurId = &recurId
	}
}

func (c NotesCalendarEntry) Counter(comments String, start NotesDateTime, end NotesDateTime, params ...notesCalendarEntryCounterParam) error {
	paramsStruct := &notesCalendarEntryCounterParams{}
	paramsOrdered := []interface{}{comments, start.com().Dispatch(), end.com().Dispatch()}

	for _, p := range params {
		p(paramsStruct)
	}

	if paramsStruct.keepPlaceholder != nil {
		paramsOrdered = append(paramsOrdered, *paramsStruct.keepPlaceholder)
		if paramsStruct.scope != nil {
			paramsOrdered = append(paramsOrdered, *paramsStruct.scope)
			if paramsStruct.recurId != nil {
				paramsOrdered = append(paramsOrdered, *paramsStruct.recurId)
			}
		}
	}
	_, err := callComMethod(c, "Counter", paramsOrdered...)
	return err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_DECLINE_METHOD_CALENTRY.html */
type notesCalendarEntryDeclineParams struct {
	keepInformed *Boolean
	scope        *Integer
	recurId      *String
}

type notesCalendarEntryDeclineParam func(*notesCalendarEntryDeclineParams)

func WithNotesCalendarEntryDeclineKeepInformed(keepInformed Boolean) notesCalendarEntryDeclineParam {
	return func(c *notesCalendarEntryDeclineParams) {
		c.keepInformed = &keepInformed
	}
}

func WithNotesCalendarEntryDeclineScope(scope Integer) notesCalendarEntryDeclineParam {
	return func(c *notesCalendarEntryDeclineParams) {
		c.scope = &scope
	}
}

func WithNotesCalendarEntryDeclineRecurId(recurId String) notesCalendarEntryDeclineParam {
	return func(c *notesCalendarEntryDeclineParams) {
		c.recurId = &recurId
	}
}

func (c NotesCalendarEntry) Decline(comments String, params ...notesCalendarEntryDeclineParam) error {
	paramsStruct := &notesCalendarEntryDeclineParams{}
	paramsOrdered := []interface{}{comments}

	for _, p := range params {
		p(paramsStruct)
	}

	if paramsStruct.keepInformed != nil {
		paramsOrdered = append(paramsOrdered, *paramsStruct.keepInformed)
		if paramsStruct.scope != nil {
			paramsOrdered = append(paramsOrdered, *paramsStruct.scope)
			if paramsStruct.recurId != nil {
				paramsOrdered = append(paramsOrdered, *paramsStruct.recurId)
			}
		}
	}
	_, err := callComMethod(c, "Decline", paramsOrdered...)
	return err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_DELEGATE_METHOD_CALENTRY.html */
type notesCalendarEntryDelegateParams struct {
	keepInformed *Boolean
	scope        *Integer
	recurId      *String
}

type notesCalendarEntryDelegateParam func(*notesCalendarEntryDelegateParams)

func WithNotesCalendarEntryDelegateKeepInformed(keepInformed Boolean) notesCalendarEntryDelegateParam {
	return func(c *notesCalendarEntryDelegateParams) {
		c.keepInformed = &keepInformed
	}
}

func WithNotesCalendarEntryDelegateScope(scope Integer) notesCalendarEntryDelegateParam {
	return func(c *notesCalendarEntryDelegateParams) {
		c.scope = &scope
	}
}

func WithNotesCalendarEntryDelegateRecurId(recurId String) notesCalendarEntryDelegateParam {
	return func(c *notesCalendarEntryDelegateParams) {
		c.recurId = &recurId
	}
}

func (c NotesCalendarEntry) Delegate(commentsToOrganizer String, delegateTo String, params ...notesCalendarEntryDelegateParam) error {
	paramsStruct := &notesCalendarEntryDelegateParams{}
	paramsOrdered := []interface{}{commentsToOrganizer, delegateTo}

	for _, p := range params {
		p(paramsStruct)
	}

	if paramsStruct.keepInformed != nil {
		paramsOrdered = append(paramsOrdered, *paramsStruct.keepInformed)
		if paramsStruct.scope != nil {
			paramsOrdered = append(paramsOrdered, *paramsStruct.scope)
			if paramsStruct.recurId != nil {
				paramsOrdered = append(paramsOrdered, *paramsStruct.recurId)
			}
		}
	}
	_, err := callComMethod(c, "Delegate", paramsOrdered...)
	return err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GETASDOCUMENT_METHOD_CALENTRY.html */
type notesCalendarEntryGetAsDocumentParams struct {
	flags   *Integer
	recurId *String
}

type notesCalendarEntryGetAsDocumentParam func(*notesCalendarEntryGetAsDocumentParams)

func WithNotesCalendarEntryGetAsDocumentFlags(flags Integer) notesCalendarEntryGetAsDocumentParam {
	return func(c *notesCalendarEntryGetAsDocumentParams) {
		c.flags = &flags
	}
}

func WithNotesCalendarEntryGetAsDocumentRecurId(recurId String) notesCalendarEntryGetAsDocumentParam {
	return func(c *notesCalendarEntryGetAsDocumentParams) {
		c.recurId = &recurId
	}
}

func (c NotesCalendarEntry) GetAsDocument(params ...notesCalendarEntryGetAsDocumentParam) (NotesDocument, error) {
	paramsStruct := &notesCalendarEntryGetAsDocumentParams{}
	paramsOrdered := []interface{}{}

	for _, p := range params {
		p(paramsStruct)
	}

	if paramsStruct.flags != nil {
		paramsOrdered = append(paramsOrdered, *paramsStruct.flags)
		if paramsStruct.recurId != nil {
			paramsOrdered = append(paramsOrdered, *paramsStruct.recurId)
		}
	}
	dispatchPtr, err := callComObjectMethod(c, "GetAsDocument", paramsOrdered...)
	return NewNotesDocument(dispatchPtr), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GETNOTICES_METHOD_CALENTRY.html */
func (c NotesCalendarEntry) GetNotices() ([]NotesCalendarNotice, error) {
	return com.CallObjectArrayMethod(c.com(), NewNotesCalendarNotice, "GetNotices")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_MODIFYINVITEES_METHOD_CALENTRY.html */
type notesCalendarEntryModifyInviteesParams struct {
	comments *String
	flags    *Integer
	scope    *Integer
	recurId  *String
}

type notesCalendarEntryModifyInviteesParam func(*notesCalendarEntryModifyInviteesParams)

func WithNotesCalendarEntryModifyInviteesComments(comments String) notesCalendarEntryModifyInviteesParam {
	return func(c *notesCalendarEntryModifyInviteesParams) {
		c.comments = &comments
	}
}

func WithNotesCalendarEntryModifyInviteesFlags(flags Integer) notesCalendarEntryModifyInviteesParam {
	return func(c *notesCalendarEntryModifyInviteesParams) {
		c.flags = &flags
	}
}

func WithNotesCalendarEntryModifyInviteesScope(scope Integer) notesCalendarEntryModifyInviteesParam {
	return func(c *notesCalendarEntryModifyInviteesParams) {
		c.scope = &scope
	}
}

func WithNotesCalendarEntryModifyInviteesRecurId(recurId String) notesCalendarEntryModifyInviteesParam {
	return func(c *notesCalendarEntryModifyInviteesParams) {
		c.recurId = &recurId
	}
}

func (c NotesCalendarEntry) ModifyInvitees(requiredNames []String, optionalNames []String, fyiNames []String, removeNames []String, params ...notesCalendarEntryModifyInviteesParam) error {
	paramsStruct := &notesCalendarEntryModifyInviteesParams{}
	paramsOrdered := []interface{}{requiredNames, optionalNames, fyiNames, removeNames}

	for _, p := range params {
		p(paramsStruct)
	}

	if paramsStruct.comments != nil {
		paramsOrdered = append(paramsOrdered, *paramsStruct.comments)
		if paramsStruct.flags != nil {
			paramsOrdered = append(paramsOrdered, *paramsStruct.flags)
			if paramsStruct.scope != nil {
				paramsOrdered = append(paramsOrdered, *paramsStruct.scope)
				if paramsStruct.recurId != nil {
					paramsOrdered = append(paramsOrdered, *paramsStruct.recurId)
				}
			}
		}
	}
	_, err := callComMethod(c, "ModifyInvitees", paramsOrdered...)
	return err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_READ_METHOD_CALENTRY.html */
type notesCalendarEntryReadParams struct {
	recurId *String
}

type notesCalendarEntryReadParam func(*notesCalendarEntryReadParams)

func WithNotesCalendarEntryReadRecurId(recurId String) notesCalendarEntryReadParam {
	return func(c *notesCalendarEntryReadParams) {
		c.recurId = &recurId
	}
}

func (c NotesCalendarEntry) Read(params ...notesCalendarEntryReadParam) (String, error) {
	paramsStruct := &notesCalendarEntryReadParams{}
	paramsOrdered := []interface{}{}

	for _, p := range params {
		p(paramsStruct)
	}

	if paramsStruct.recurId != nil {
		paramsOrdered = append(paramsOrdered, *paramsStruct.recurId)
	}
	val, err := callComMethod(c, "Read", paramsOrdered...)
	return helpers.CastValue[String](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_REMOVE_METHOD_CALENTRY.html */
type notesCalendarEntryRemoveParams struct {
	scope   *Integer
	recurId *String
}

type notesCalendarEntryRemoveParam func(*notesCalendarEntryRemoveParams)

func WithNotesCalendarEntryRemoveScope(scope Integer) notesCalendarEntryRemoveParam {
	return func(c *notesCalendarEntryRemoveParams) {
		c.scope = &scope
	}
}

func WithNotesCalendarEntryRemoveRecurId(recurId String) notesCalendarEntryRemoveParam {
	return func(c *notesCalendarEntryRemoveParams) {
		c.recurId = &recurId
	}
}

func (c NotesCalendarEntry) Remove(comments String, params ...notesCalendarEntryRemoveParam) error {
	paramsStruct := &notesCalendarEntryRemoveParams{}
	paramsOrdered := []interface{}{comments}

	for _, p := range params {
		p(paramsStruct)
	}

	if paramsStruct.scope != nil {
		paramsOrdered = append(paramsOrdered, *paramsStruct.scope)
		if paramsStruct.recurId != nil {
			paramsOrdered = append(paramsOrdered, *paramsStruct.recurId)
		}
	}
	_, err := callComMethod(c, "Remove", paramsOrdered...)
	return err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_REMOVEINVITEES_METHOD_CALENTRY.html */
type notesCalendarEntryRemoveInviteesParams struct {
	comments *String
	flags    *Integer
	scope    *Integer
	recurId  *String
}

type notesCalendarEntryRemoveInviteesParam func(*notesCalendarEntryRemoveInviteesParams)

func WithNotesCalendarEntryRemoveInviteesComments(comments String) notesCalendarEntryRemoveInviteesParam {
	return func(c *notesCalendarEntryRemoveInviteesParams) {
		c.comments = &comments
	}
}

func WithNotesCalendarEntryRemoveInviteesFlags(flags Integer) notesCalendarEntryRemoveInviteesParam {
	return func(c *notesCalendarEntryRemoveInviteesParams) {
		c.flags = &flags
	}
}

func WithNotesCalendarEntryRemoveInviteesScope(scope Integer) notesCalendarEntryRemoveInviteesParam {
	return func(c *notesCalendarEntryRemoveInviteesParams) {
		c.scope = &scope
	}
}

func WithNotesCalendarEntryRemoveInviteesRecurId(recurId String) notesCalendarEntryRemoveInviteesParam {
	return func(c *notesCalendarEntryRemoveInviteesParams) {
		c.recurId = &recurId
	}
}

func (c NotesCalendarEntry) RemoveInvitees(names []String, params ...notesCalendarEntryRemoveInviteesParam) error {
	paramsStruct := &notesCalendarEntryRemoveInviteesParams{}
	paramsOrdered := []interface{}{names}

	for _, p := range params {
		p(paramsStruct)
	}

	if paramsStruct.comments != nil {
		paramsOrdered = append(paramsOrdered, *paramsStruct.comments)
		if paramsStruct.flags != nil {
			paramsOrdered = append(paramsOrdered, *paramsStruct.flags)
			if paramsStruct.scope != nil {
				paramsOrdered = append(paramsOrdered, *paramsStruct.scope)
				if paramsStruct.recurId != nil {
					paramsOrdered = append(paramsOrdered, *paramsStruct.recurId)
				}
			}
		}
	}
	_, err := callComMethod(c, "RemoveInvitees", paramsOrdered...)
	return err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_REQUESTINFO_METHOD_CALENTRY.html */
func (c NotesCalendarEntry) RequestInfo(comments String) error {
	_, err := callComMethod(c, "RequestInfo", comments)
	return err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_TENTATIVELYACCEPT_METHOD_CALENTRY.html */
type notesCalendarEntryTentativelyAcceptParams struct {
	scope   *Integer
	recurId *String
}

type notesCalendarEntryTentativelyAcceptParam func(*notesCalendarEntryTentativelyAcceptParams)

func WithNotesCalendarEntryTentativelyAcceptScope(scope Integer) notesCalendarEntryTentativelyAcceptParam {
	return func(c *notesCalendarEntryTentativelyAcceptParams) {
		c.scope = &scope
	}
}

func WithNotesCalendarEntryTentativelyAcceptRecurId(recurId String) notesCalendarEntryTentativelyAcceptParam {
	return func(c *notesCalendarEntryTentativelyAcceptParams) {
		c.recurId = &recurId
	}
}

func (c NotesCalendarEntry) TentativelyAccept(comments String, params ...notesCalendarEntryTentativelyAcceptParam) error {
	paramsStruct := &notesCalendarEntryTentativelyAcceptParams{}
	paramsOrdered := []interface{}{comments}

	for _, p := range params {
		p(paramsStruct)
	}

	if paramsStruct.scope != nil {
		paramsOrdered = append(paramsOrdered, *paramsStruct.scope)
		if paramsStruct.recurId != nil {
			paramsOrdered = append(paramsOrdered, *paramsStruct.recurId)
		}
	}
	_, err := callComMethod(c, "TentativelyAccept", paramsOrdered...)
	return err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_UPDATE_METHOD_CALENTRY.html */
type notesCalendarEntryUpdateParams struct {
	comments *String
	flags    *Integer
	recurId  *String
}

type notesCalendarEntryUpdateParam func(*notesCalendarEntryUpdateParams)

func WithNotesCalendarEntryUpdateComments(comments String) notesCalendarEntryUpdateParam {
	return func(c *notesCalendarEntryUpdateParams) {
		c.comments = &comments
	}
}

func WithNotesCalendarEntryUpdateFlags(flags Integer) notesCalendarEntryUpdateParam {
	return func(c *notesCalendarEntryUpdateParams) {
		c.flags = &flags
	}
}

func WithNotesCalendarEntryUpdateRecurId(recurId String) notesCalendarEntryUpdateParam {
	return func(c *notesCalendarEntryUpdateParams) {
		c.recurId = &recurId
	}
}

func (c NotesCalendarEntry) Update(icalentry String, params ...notesCalendarEntryUpdateParam) error {
	paramsStruct := &notesCalendarEntryUpdateParams{}
	paramsOrdered := []interface{}{icalentry}

	for _, p := range params {
		p(paramsStruct)
	}

	if paramsStruct.comments != nil {
		paramsOrdered = append(paramsOrdered, *paramsStruct.comments)
		if paramsStruct.flags != nil {
			paramsOrdered = append(paramsOrdered, *paramsStruct.flags)
			if paramsStruct.recurId != nil {
				paramsOrdered = append(paramsOrdered, *paramsStruct.recurId)
			}
		}
	}
	_, err := callComMethod(c, "Update", paramsOrdered...)
	return err
}
