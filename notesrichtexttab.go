/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_RICHTEXTTAB_CLASS.html */
package domigo

import (
	ole "github.com/go-ole/go-ole"
)

type NotesRichTextTabType = Long

const (
	NOTESRICHTEXTTAB_TAB_CENTER  NotesRichTextTabType = 3
	NOTESRICHTEXTTAB_TAB_DECIMAL NotesRichTextTabType = 2
	NOTESRICHTEXTTAB_TAB_LEFT    NotesRichTextTabType = 0
	NOTESRICHTEXTTAB_TAB_RIGHT   NotesRichTextTabType = 1
)

type NotesRichTextTab struct {
	NotesStruct
}

func NewNotesRichTextTab(dispatchPtr *ole.IDispatch) NotesRichTextTab {
	return NotesRichTextTab{NewNotesStruct(dispatchPtr)}
}

/* --------------------------------- Properties --------------------------------- */
/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_POSITION_PROPERTY_4695.html */
func (r NotesRichTextTab) Position() (Long, error) {
	return getComProperty[Long](r, "Position")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_TYPE_PROPERTY_7073.html */
func (r NotesRichTextTab) Type() (NotesRichTextTabType, error) {
	return getComProperty[NotesRichTextTabType](r, "Type")
}

/* --------------------------------- Methods ------------------------------------ */
/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_CLEAR_METHOD_1029.html */
func (r NotesRichTextTab) Clear() error {
	return callComVoidMethod(r, "Clear")
}
