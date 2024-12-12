/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_NOTESCOLOROBJECT_CLASS.html */
package notescolorobject

import (
	"domigo/domino"
	"domigo/helpers"

	ole "github.com/go-ole/go-ole"
)

type NotesColorObject struct {
	domino.NotesStruct
}

func New(dispatchPtr *ole.IDispatch) NotesColorObject {
	return NotesColorObject{domino.NewNotesStruct(dispatchPtr)}
}

/* --------------------------------- Properties --------------------------------- */
/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_BLUE_PROPERTY_COLOR.html */
func (c NotesColorObject) Blue() (domino.Long, error) {
	val, err := c.Com().GetProperty("Blue")
	return helpers.CastValue[domino.Long](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GREEN_PROPERTY_COLOR.html */
func (c NotesColorObject) Green() (domino.Long, error) {
	val, err := c.Com().GetProperty("Green")
	return helpers.CastValue[domino.Long](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_HUE_PROPERTY_COLOR.html */
func (c NotesColorObject) Hue() (domino.Long, error) {
	val, err := c.Com().GetProperty("Hue")
	return helpers.CastValue[domino.Long](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_LUMINANCE_PROPERTY_COLOR.html */
func (c NotesColorObject) Luminance() (domino.Long, error) {
	val, err := c.Com().GetProperty("Luminance")
	return helpers.CastValue[domino.Long](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_NOTESCOLOR_PROPERTY_COLOR.html */
func (c NotesColorObject) NotesColor() (domino.Long, error) {
	val, err := c.Com().GetProperty("NotesColor")
	return helpers.CastValue[domino.Long](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_NOTESCOLOR_PROPERTY_COLOR.html */
func (c NotesColorObject) SetNotesColor(v domino.Long) error {
	return c.Com().PutProperty("NotesColor", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_RED_PROPERTY_COLOR.html */
func (c NotesColorObject) Red() (domino.Long, error) {
	val, err := c.Com().GetProperty("Red")
	return helpers.CastValue[domino.Long](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SATURATION_PROPERTY_COLOR.html */
func (c NotesColorObject) Saturation() (domino.Long, error) {
	val, err := c.Com().GetProperty("Saturation")
	return helpers.CastValue[domino.Long](val), err
}

/* --------------------------------- Methods ------------------------------------ */
/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SETHSL_METHOD_COLOR.html */
func (c NotesColorObject) SetHSL(hue domino.Long, saturation domino.Long, luminance domino.Long) (domino.Long, error) {
	val, err := c.Com().CallMethod("SetHSL", hue, saturation, luminance)
	return helpers.CastValue[domino.Long](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SETRGB_METHOD_COLOR.html */
func (c NotesColorObject) SetRGB(red domino.Long, green domino.Long, blue domino.Long) (domino.Long, error) {
	val, err := c.Com().CallMethod("SetRGB", red, green, blue)
	return helpers.CastValue[domino.Long](val), err
}
