/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_RICHTEXTTAB_CLASS.html */
package notesrichtexttab

import (
	"domigo/domino"
	"domigo/helpers"

	ole "github.com/go-ole/go-ole"
)

type TabType = domino.Long

const (
	TAB_CENTER  TabType = 3
	TAB_DECIMAL TabType = 2
	TAB_LEFT    TabType = 0
	TAB_RIGHT   TabType = 1
)

type NotesRichTextTab struct {
	domino.NotesStruct
}

func New(dispatchPtr *ole.IDispatch) NotesRichTextTab {
	return NotesRichTextTab{domino.NewNotesStruct(dispatchPtr)}
}

/* --------------------------------- Properties --------------------------------- */
/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_POSITION_PROPERTY_4695.html */
func (r NotesRichTextTab) Position() (domino.Long, error) {
	val, err := r.Com().GetProperty("Position")
	return helpers.CastValue[domino.Long](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_TYPE_PROPERTY_7073.html */
func (r NotesRichTextTab) Type() (TabType, error) {
	val, err := r.Com().GetProperty("Type")
	return helpers.CastValue[TabType](val), err
}

/* --------------------------------- Methods ------------------------------------ */
/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_CLEAR_METHOD_1029.html */
func (r NotesRichTextTab) Clear() error {
	_, err := r.Com().CallMethod("Clear")
	return err
}
