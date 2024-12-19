/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_NOTESCOLOROBJECT_CLASS.html */
package domigo

import (
	"github.com/monstermichl/domigo/helpers"

	ole "github.com/go-ole/go-ole"
)

type NotesColorObject struct {
	NotesStruct
}

func NewNotesColorObject(dispatchPtr *ole.IDispatch) NotesColorObject {
	return NotesColorObject{NewNotesStruct(dispatchPtr)}
}

/* --------------------------------- Properties --------------------------------- */
/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_BLUE_PROPERTY_COLOR.html */
func (c NotesColorObject) Blue() (Long, error) {
	val, err := c.com().GetProperty("Blue")
	return helpers.CastValue[Long](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GREEN_PROPERTY_COLOR.html */
func (c NotesColorObject) Green() (Long, error) {
	val, err := c.com().GetProperty("Green")
	return helpers.CastValue[Long](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_HUE_PROPERTY_COLOR.html */
func (c NotesColorObject) Hue() (Long, error) {
	val, err := c.com().GetProperty("Hue")
	return helpers.CastValue[Long](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_LUMINANCE_PROPERTY_COLOR.html */
func (c NotesColorObject) Luminance() (Long, error) {
	val, err := c.com().GetProperty("Luminance")
	return helpers.CastValue[Long](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_NOTESCOLOR_PROPERTY_COLOR.html */
func (c NotesColorObject) NotesColor() (Long, error) {
	val, err := c.com().GetProperty("NotesColor")
	return helpers.CastValue[Long](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_NOTESCOLOR_PROPERTY_COLOR.html */
func (c NotesColorObject) SetNotesColor(v Long) error {
	return c.com().PutProperty("NotesColor", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_RED_PROPERTY_COLOR.html */
func (c NotesColorObject) Red() (Long, error) {
	val, err := c.com().GetProperty("Red")
	return helpers.CastValue[Long](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SATURATION_PROPERTY_COLOR.html */
func (c NotesColorObject) Saturation() (Long, error) {
	val, err := c.com().GetProperty("Saturation")
	return helpers.CastValue[Long](val), err
}

/* --------------------------------- Methods ------------------------------------ */
/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SETHSL_METHOD_COLOR.html */
func (c NotesColorObject) SetHSL(hue Long, saturation Long, luminance Long) (Long, error) {
	val, err := c.com().CallMethod("SetHSL", hue, saturation, luminance)
	return helpers.CastValue[Long](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SETRGB_METHOD_COLOR.html */
func (c NotesColorObject) SetRGB(red Long, green Long, blue Long) (Long, error) {
	val, err := c.com().CallMethod("SetRGB", red, green, blue)
	return helpers.CastValue[Long](val), err
}
