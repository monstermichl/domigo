/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_NOTESCOLOROBJECT_CLASS.html */
package domigo

import (
	ole "github.com/go-ole/go-ole"
)

type NotesColorObject struct {
	NotesStruct
}

func newNotesColorObject(dispatchPtr *ole.IDispatch) NotesColorObject {
	return NotesColorObject{newNotesStruct(dispatchPtr)}
}

/* --------------------------------- Properties --------------------------------- */
/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_BLUE_PROPERTY_COLOR.html */
func (c NotesColorObject) Blue() (Long, error) {
	return getComProperty[Long](c, "Blue")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GREEN_PROPERTY_COLOR.html */
func (c NotesColorObject) Green() (Long, error) {
	return getComProperty[Long](c, "Green")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_HUE_PROPERTY_COLOR.html */
func (c NotesColorObject) Hue() (Long, error) {
	return getComProperty[Long](c, "Hue")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_LUMINANCE_PROPERTY_COLOR.html */
func (c NotesColorObject) Luminance() (Long, error) {
	return getComProperty[Long](c, "Luminance")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_NOTESCOLOR_PROPERTY_COLOR.html */
func (c NotesColorObject) NotesColor() (Long, error) {
	return getComProperty[Long](c, "NotesColor")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_NOTESCOLOR_PROPERTY_COLOR.html */
func (c NotesColorObject) SetNotesColor(v Long) error {
	return putComProperty(c, "NotesColor", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_RED_PROPERTY_COLOR.html */
func (c NotesColorObject) Red() (Long, error) {
	return getComProperty[Long](c, "Red")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SATURATION_PROPERTY_COLOR.html */
func (c NotesColorObject) Saturation() (Long, error) {
	return getComProperty[Long](c, "Saturation")
}

/* --------------------------------- Methods ------------------------------------ */
/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SETHSL_METHOD_COLOR.html */
func (c NotesColorObject) SetHSL(hue Long, saturation Long, luminance Long) (Long, error) {
	return callComMethod[Long](c, "SetHSL", hue, saturation, luminance)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SETRGB_METHOD_COLOR.html */
func (c NotesColorObject) SetRGB(red Long, green Long, blue Long) (Long, error) {
	return callComMethod[Long](c, "SetRGB", red, green, blue)
}
