/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_NOTESVIEWCOLUMN_CLASS.html */
package domigo

import (
	"github.com/monstermichl/domigo/helpers"

	ole "github.com/go-ole/go-ole"
)

type NotesViewColumn struct {
	NotesStruct
}

func NewNotesViewColumn(dispatchPtr *ole.IDispatch) NotesViewColumn {
	return NotesViewColumn{NewNotesStruct(dispatchPtr)}
}

/* --------------------------------- Properties --------------------------------- */
/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ALIGNMENT_PROPERTY_2919_ABOUT.html */
func (c NotesViewColumn) Alignment() (Integer, error) {
	val, err := c.Com().GetProperty("Alignment")
	return helpers.CastValue[Integer](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ALIGNMENT_PROPERTY_2919_ABOUT.html */
func (c NotesViewColumn) SetAlignment(v Integer) error {
	return c.Com().PutProperty("Alignment", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_DATEFMT_PROPERTY_4095_ABOUT.html */
func (c NotesViewColumn) DateFmt() (Integer, error) {
	val, err := c.Com().GetProperty("DateFmt")
	return helpers.CastValue[Integer](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_DATEFMT_PROPERTY_4095_ABOUT.html */
func (c NotesViewColumn) SetDateFmt(v Integer) error {
	return c.Com().PutProperty("DateFmt", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_FONTCOLOR_PROPERTY_9482_ABOUT.html */
func (c NotesViewColumn) FontColor() (Integer, error) {
	val, err := c.Com().GetProperty("FontColor")
	return helpers.CastValue[Integer](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_FONTCOLOR_PROPERTY_9482_ABOUT.html */
func (c NotesViewColumn) SetFontColor(v Integer) error {
	return c.Com().PutProperty("FontColor", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_FONTFACE_PROPERTY_5960_ABOUT.html */
func (c NotesViewColumn) FontFace() (String, error) {
	val, err := c.Com().GetProperty("FontFace")
	return helpers.CastValue[String](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_FONTFACE_PROPERTY_5960_ABOUT.html */
func (c NotesViewColumn) SetFontFace(v String) error {
	return c.Com().PutProperty("FontFace", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_FONTPOINTSIZE_PROPERTY_7819_ABOUT.html */
func (c NotesViewColumn) FontPointSize() (Integer, error) {
	val, err := c.Com().GetProperty("FontPointSize")
	return helpers.CastValue[Integer](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_FONTPOINTSIZE_PROPERTY_7819_ABOUT.html */
func (c NotesViewColumn) SetFontPointSize(v Integer) error {
	return c.Com().PutProperty("FontPointSize", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_FONTSTYLE_PROPERTY_2400_ABOUT.html */
func (c NotesViewColumn) FontStyle() (Integer, error) {
	val, err := c.Com().GetProperty("FontStyle")
	return helpers.CastValue[Integer](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_FONTSTYLE_PROPERTY_2400_ABOUT.html */
func (c NotesViewColumn) SetFontStyle(v Integer) error {
	return c.Com().PutProperty("FontStyle", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_FORMULA_PROPERTY.html */
func (c NotesViewColumn) Formula() (String, error) {
	val, err := c.Com().GetProperty("Formula")
	return helpers.CastValue[String](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_FORMULA_PROPERTY.html */
func (c NotesViewColumn) SetFormula(v String) error {
	return c.Com().PutProperty("Formula", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_HEADER_ALIGNMENT_PROPERTY_4751_ABOUT.html */
func (c NotesViewColumn) HeaderAlignment() (Integer, error) {
	val, err := c.Com().GetProperty("HeaderAlignment")
	return helpers.CastValue[Integer](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_HEADER_ALIGNMENT_PROPERTY_4751_ABOUT.html */
func (c NotesViewColumn) SetHeaderAlignment(v Integer) error {
	return c.Com().PutProperty("HeaderAlignment", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_HEADERFONTCOLOR_PROPERTY_VC.html */
func (c NotesViewColumn) HeaderFontColor() (Integer, error) {
	val, err := c.Com().GetProperty("HeaderFontColor")
	return helpers.CastValue[Integer](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_HEADERFONTCOLOR_PROPERTY_VC.html */
func (c NotesViewColumn) SetHeaderFontColor(v Integer) error {
	return c.Com().PutProperty("HeaderFontColor", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_HEADERFONTFACE_PROPERTY_VC.html */
func (c NotesViewColumn) HeaderFontFace() (String, error) {
	val, err := c.Com().GetProperty("HeaderFontFace")
	return helpers.CastValue[String](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_HEADERFONTFACE_PROPERTY_VC.html */
func (c NotesViewColumn) SetHeaderFontFace(v String) error {
	return c.Com().PutProperty("HeaderFontFace", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_HEADERFONTPOINTSIZE_PROPERTY_VC.html */
func (c NotesViewColumn) HeaderFontPointSize() (Integer, error) {
	val, err := c.Com().GetProperty("HeaderFontPointSize")
	return helpers.CastValue[Integer](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_HEADERFONTPOINTSIZE_PROPERTY_VC.html */
func (c NotesViewColumn) SetHeaderFontPointSize(v Integer) error {
	return c.Com().PutProperty("HeaderFontPointSize", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_HEADERFONTSTYLE_PROPERTY_VC.html */
func (c NotesViewColumn) HeaderFontStyle() (Integer, error) {
	val, err := c.Com().GetProperty("HeaderFontStyle")
	return helpers.CastValue[Integer](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_HEADERFONTSTYLE_PROPERTY_VC.html */
func (c NotesViewColumn) SetHeaderFontStyle(v Integer) error {
	return c.Com().PutProperty("HeaderFontStyle", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISACCENTINSENSITIVESORT_1949_ABOUT.html */
func (c NotesViewColumn) IsAccentSensitiveSort() (Boolean, error) {
	val, err := c.Com().GetProperty("IsAccentSensitiveSort")
	return helpers.CastValue[Boolean](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISACCENTINSENSITIVESORT_1949_ABOUT.html */
func (c NotesViewColumn) SetIsAccentSensitiveSort(v Boolean) error {
	return c.Com().PutProperty("IsAccentSensitiveSort", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISCASEINSENSITIVESORT_PROPERTY_1451_ABOUT.html */
func (c NotesViewColumn) IsCaseSensitiveSort() (Boolean, error) {
	val, err := c.Com().GetProperty("IsCaseSensitiveSort")
	return helpers.CastValue[Boolean](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISCASEINSENSITIVESORT_PROPERTY_1451_ABOUT.html */
func (c NotesViewColumn) SetIsCaseSensitiveSort(v Boolean) error {
	return c.Com().PutProperty("IsCaseSensitiveSort", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISCATEGORY_PROPERTY.html */
func (c NotesViewColumn) IsCategory() (Boolean, error) {
	val, err := c.Com().GetProperty("IsCategory")
	return helpers.CastValue[Boolean](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISFIELD_PROPERTY_3448_ABOUT.html */
func (c NotesViewColumn) IsField() (Boolean, error) {
	val, err := c.Com().GetProperty("IsField")
	return helpers.CastValue[Boolean](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISFONTBOLD_PROPERTY_VC.html */
func (c NotesViewColumn) IsFontBold() (Boolean, error) {
	val, err := c.Com().GetProperty("IsFontBold")
	return helpers.CastValue[Boolean](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISFONTBOLD_PROPERTY_VC.html */
func (c NotesViewColumn) SetIsFontBold(v Boolean) error {
	return c.Com().PutProperty("IsFontBold", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISFONTITALIC_PROPERTY_VC.html */
func (c NotesViewColumn) IsFontItalic() (Boolean, error) {
	val, err := c.Com().GetProperty("IsFontItalic")
	return helpers.CastValue[Boolean](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISFONTITALIC_PROPERTY_VC.html */
func (c NotesViewColumn) SetIsFontItalic(v Boolean) error {
	return c.Com().PutProperty("IsFontItalic", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISFONTSTRIKETHROUGH_PROPERTY_VC.html */
func (c NotesViewColumn) IsFontStrikethrough() (Boolean, error) {
	val, err := c.Com().GetProperty("IsFontStrikethrough")
	return helpers.CastValue[Boolean](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISFONTSTRIKETHROUGH_PROPERTY_VC.html */
func (c NotesViewColumn) SetIsFontStrikethrough(v Boolean) error {
	return c.Com().PutProperty("IsFontStrikethrough", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISFONTUNDERLINE_PROPERTY_VC.html */
func (c NotesViewColumn) IsFontUnderline() (Boolean, error) {
	val, err := c.Com().GetProperty("IsFontUnderline")
	return helpers.CastValue[Boolean](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISFONTUNDERLINE_PROPERTY_VC.html */
func (c NotesViewColumn) SetIsFontUnderline(v Boolean) error {
	return c.Com().PutProperty("IsFontUnderline", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISFORMULA_PROPERTY.html */
func (c NotesViewColumn) IsFormula() (Boolean, error) {
	val, err := c.Com().GetProperty("IsFormula")
	return helpers.CastValue[Boolean](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISHEADERFONTBOLD_PROPERTY_VC.html */
func (c NotesViewColumn) IsHeaderFontBold() (Boolean, error) {
	val, err := c.Com().GetProperty("IsHeaderFontBold")
	return helpers.CastValue[Boolean](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISHEADERFONTBOLD_PROPERTY_VC.html */
func (c NotesViewColumn) SetIsHeaderFontBold(v Boolean) error {
	return c.Com().PutProperty("IsHeaderFontBold", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISHEADERFONTITALIC_PROPERTY_VC.html */
func (c NotesViewColumn) IsHeaderFontItalic() (Boolean, error) {
	val, err := c.Com().GetProperty("IsHeaderFontItalic")
	return helpers.CastValue[Boolean](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISHEADERFONTITALIC_PROPERTY_VC.html */
func (c NotesViewColumn) SetIsHeaderFontItalic(v Boolean) error {
	return c.Com().PutProperty("IsHeaderFontItalic", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISHEADERFONTSTRIKETHROUGH_PROPERTY_VC.html */
func (c NotesViewColumn) IsHeaderFontStrikethrough() (Boolean, error) {
	val, err := c.Com().GetProperty("IsHeaderFontStrikethrough")
	return helpers.CastValue[Boolean](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISHEADERFONTSTRIKETHROUGH_PROPERTY_VC.html */
func (c NotesViewColumn) SetIsHeaderFontStrikethrough(v Boolean) error {
	return c.Com().PutProperty("IsHeaderFontStrikethrough", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISHEADERFONTUNDERLINE_PROPERTY_VC.html */
func (c NotesViewColumn) IsHeaderFontUnderline() (Boolean, error) {
	val, err := c.Com().GetProperty("IsHeaderFontUnderline")
	return helpers.CastValue[Boolean](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISHEADERFONTUNDERLINE_PROPERTY_VC.html */
func (c NotesViewColumn) SetIsHeaderFontUnderline(v Boolean) error {
	return c.Com().PutProperty("IsHeaderFontUnderline", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISHIDDEN_PROPERTY.html */
func (c NotesViewColumn) IsHidden() (Boolean, error) {
	val, err := c.Com().GetProperty("IsHidden")
	return helpers.CastValue[Boolean](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISHIDDEN_PROPERTY.html */
func (c NotesViewColumn) SetIsHidden(v Boolean) error {
	return c.Com().PutProperty("IsHidden", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISHIDEDETAIL_PROPERTY.html */
func (c NotesViewColumn) IsHideDetail() (Boolean, error) {
	val, err := c.Com().GetProperty("IsHideDetail")
	return helpers.CastValue[Boolean](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISHIDEDETAIL_PROPERTY.html */
func (c NotesViewColumn) SetIsHideDetail(v Boolean) error {
	return c.Com().PutProperty("IsHideDetail", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISICON_PROPERTY.html */
func (c NotesViewColumn) IsIcon() (Boolean, error) {
	val, err := c.Com().GetProperty("IsIcon")
	return helpers.CastValue[Boolean](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISNUMBERATTRIBPARENS_PROPERTY_VC.html */
func (c NotesViewColumn) IsNumberAttribParens() (Boolean, error) {
	val, err := c.Com().GetProperty("IsNumberAttribParens")
	return helpers.CastValue[Boolean](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISNUMBERATTRIBPARENS_PROPERTY_VC.html */
func (c NotesViewColumn) SetIsNumberAttribParens(v Boolean) error {
	return c.Com().PutProperty("IsNumberAttribParens", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISNUMBERATTRIBPERCENT_PROPERTY_VC.html */
func (c NotesViewColumn) IsNumberAttribPercent() (Boolean, error) {
	val, err := c.Com().GetProperty("IsNumberAttribPercent")
	return helpers.CastValue[Boolean](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISNUMBERATTRIBPERCENT_PROPERTY_VC.html */
func (c NotesViewColumn) SetIsNumberAttribPercent(v Boolean) error {
	return c.Com().PutProperty("IsNumberAttribPercent", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISNUMBERATTRIBPUNCTUATED_PROPERTY_VC.html */
func (c NotesViewColumn) IsNumberAttribPunctuated() (Boolean, error) {
	val, err := c.Com().GetProperty("IsNumberAttribPunctuated")
	return helpers.CastValue[Boolean](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISNUMBERATTRIBPUNCTUATED_PROPERTY_VC.html */
func (c NotesViewColumn) SetIsNumberAttribPunctuated(v Boolean) error {
	return c.Com().PutProperty("IsNumberAttribPunctuated", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISRESIZE_PROPERTY.html */
func (c NotesViewColumn) IsResize() (Boolean, error) {
	val, err := c.Com().GetProperty("IsResize")
	return helpers.CastValue[Boolean](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISRESIZE_PROPERTY.html */
func (c NotesViewColumn) SetIsResize(v Boolean) error {
	return c.Com().PutProperty("IsResize", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISRESORTASCENDING_PROPERTY.html */
func (c NotesViewColumn) IsResortAscending() (Boolean, error) {
	val, err := c.Com().GetProperty("IsResortAscending")
	return helpers.CastValue[Boolean](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISRESORTASCENDING_PROPERTY.html */
func (c NotesViewColumn) SetIsResortAscending(v Boolean) error {
	return c.Com().PutProperty("IsResortAscending", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISRESORTDESCENDING_PROPERTY.html */
func (c NotesViewColumn) IsResortDescending() (Boolean, error) {
	val, err := c.Com().GetProperty("IsResortDescending")
	return helpers.CastValue[Boolean](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISRESORTDESCENDING_PROPERTY.html */
func (c NotesViewColumn) SetIsResortDescending(v Boolean) error {
	return c.Com().PutProperty("IsResortDescending", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISRESORTTOVIEW_PROPERTY.html */
func (c NotesViewColumn) IsResortToView() (Boolean, error) {
	val, err := c.Com().GetProperty("IsResortToView")
	return helpers.CastValue[Boolean](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISRESORTTOVIEW_PROPERTY.html */
func (c NotesViewColumn) SetIsResortToView(v Boolean) error {
	return c.Com().PutProperty("IsResortToView", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISRESPONSE_PROPERTY_COLUMN.html */
func (c NotesViewColumn) IsResponse() (Boolean, error) {
	val, err := c.Com().GetProperty("IsResponse")
	return helpers.CastValue[Boolean](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISSECONDARYRESORT_PROPERTY.html */
func (c NotesViewColumn) IsSecondaryResort() (Boolean, error) {
	val, err := c.Com().GetProperty("IsSecondaryResort")
	return helpers.CastValue[Boolean](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISSECONDARYRESORT_PROPERTY.html */
func (c NotesViewColumn) SetIsSecondaryResort(v Boolean) error {
	return c.Com().PutProperty("IsSecondaryResort", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISSECONDARYRESORTDESCENDING_PROPERTY.html */
func (c NotesViewColumn) IsSecondaryResortDescending() (Boolean, error) {
	val, err := c.Com().GetProperty("IsSecondaryResortDescending")
	return helpers.CastValue[Boolean](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISSECONDARYRESORTDESCENDING_PROPERTY.html */
func (c NotesViewColumn) SetIsSecondaryResortDescending(v Boolean) error {
	return c.Com().PutProperty("IsSecondaryResortDescending", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISSHOWTWISTIE_PROPERTY.html */
func (c NotesViewColumn) IsShowTwistie() (Boolean, error) {
	val, err := c.Com().GetProperty("IsShowTwistie")
	return helpers.CastValue[Boolean](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISSHOWTWISTIE_PROPERTY.html */
func (c NotesViewColumn) SetIsShowTwistie(v Boolean) error {
	return c.Com().PutProperty("IsShowTwistie", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISSORTDESCENDING_PROPERTY.html */
func (c NotesViewColumn) IsSortDescending() (Boolean, error) {
	val, err := c.Com().GetProperty("IsSortDescending")
	return helpers.CastValue[Boolean](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISSORTDESCENDING_PROPERTY.html */
func (c NotesViewColumn) SetIsSortDescending(v Boolean) error {
	return c.Com().PutProperty("IsSortDescending", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISSORTED_PROPERTY_COLUMN.html */
func (c NotesViewColumn) IsSorted() (Boolean, error) {
	val, err := c.Com().GetProperty("IsSorted")
	return helpers.CastValue[Boolean](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISSORTED_PROPERTY_COLUMN.html */
func (c NotesViewColumn) SetIsSorted(v Boolean) error {
	return c.Com().PutProperty("IsSorted", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ITEMNAME_PROPERTY.html */
func (c NotesViewColumn) ItemName() (String, error) {
	val, err := c.Com().GetProperty("ItemName")
	return helpers.CastValue[String](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_LISTSEP_PROPERTY.html */
func (c NotesViewColumn) ListSep() (Integer, error) {
	val, err := c.Com().GetProperty("ListSep")
	return helpers.CastValue[Integer](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_LISTSEP_PROPERTY.html */
func (c NotesViewColumn) SetListSep(v Integer) error {
	return c.Com().PutProperty("ListSep", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_NUMBERATTRIB_PROPERTY.html */
func (c NotesViewColumn) NumberAttrib() (Integer, error) {
	val, err := c.Com().GetProperty("NumberAttrib")
	return helpers.CastValue[Integer](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_NUMBERATTRIB_PROPERTY.html */
func (c NotesViewColumn) SetNumberAttrib(v Integer) error {
	return c.Com().PutProperty("NumberAttrib", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_NUMBERDIGITS_PROPERTY.html */
func (c NotesViewColumn) NumberDigits() (Integer, error) {
	val, err := c.Com().GetProperty("NumberDigits")
	return helpers.CastValue[Integer](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_NUMBERDIGITS_PROPERTY.html */
func (c NotesViewColumn) SetNumberDigits(v Integer) error {
	return c.Com().PutProperty("NumberDigits", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_NUMBERFORMAT_PROPERTY.html */
func (c NotesViewColumn) NumberFormat() (Integer, error) {
	val, err := c.Com().GetProperty("NumberFormat")
	return helpers.CastValue[Integer](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_NUMBERFORMAT_PROPERTY.html */
func (c NotesViewColumn) SetNumberFormat(v Integer) error {
	return c.Com().PutProperty("NumberFormat", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_PARENT_PROPERTY_NOTESVIEWCOLUMN_CLASS.html */
func (c NotesViewColumn) Parent() (NotesView, error) {
	dispatchPtr, err := c.Com().GetObjectProperty("Parent")
	return NewNotesView(dispatchPtr), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_POSITION_PROPERTY.html */
func (c NotesViewColumn) Position() (Integer, error) {
	val, err := c.Com().GetProperty("Position")
	return helpers.CastValue[Integer](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_RESORTTOVIEWNAME_PROPERTY.html */
func (c NotesViewColumn) ResortToViewName() (String, error) {
	val, err := c.Com().GetProperty("ResortToViewName")
	return helpers.CastValue[String](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_RESORTTOVIEWNAME_PROPERTY.html */
func (c NotesViewColumn) SetResortToViewName(v String) error {
	return c.Com().PutProperty("ResortToViewName", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SECONDARYRESORTCOLUMNINDEX_PROPERTY_VC.html */
func (c NotesViewColumn) SecondaryResortColumnIndex() (Integer, error) {
	val, err := c.Com().GetProperty("SecondaryResortColumnIndex")
	return helpers.CastValue[Integer](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SECONDARYRESORTCOLUMNINDEX_PROPERTY_VC.html */
func (c NotesViewColumn) SetSecondaryResortColumnIndex(v Integer) error {
	return c.Com().PutProperty("SecondaryResortColumnIndex", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_TIMEDATEFMT_PROPERTY.html */
func (c NotesViewColumn) TimeDateFmt() (Integer, error) {
	val, err := c.Com().GetProperty("TimeDateFmt")
	return helpers.CastValue[Integer](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_TIMEDATEFMT_PROPERTY.html */
func (c NotesViewColumn) SetTimeDateFmt(v Integer) error {
	return c.Com().PutProperty("TimeDateFmt", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_TIMEFMT_PROPERTY.html */
func (c NotesViewColumn) TimeFmt() (Integer, error) {
	val, err := c.Com().GetProperty("TimeFmt")
	return helpers.CastValue[Integer](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_TIMEFMT_PROPERTY.html */
func (c NotesViewColumn) SetTimeFmt(v Integer) error {
	return c.Com().PutProperty("TimeFmt", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_TIMEZONEFMT.html */
func (c NotesViewColumn) TimeZoneFmt() (Integer, error) {
	val, err := c.Com().GetProperty("TimeZoneFmt")
	return helpers.CastValue[Integer](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_TIMEZONEFMT.html */
func (c NotesViewColumn) SetTimeZoneFmt(v Integer) error {
	return c.Com().PutProperty("TimeZoneFmt", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_TITLE_PROPERTY_COLUMN.html */
func (c NotesViewColumn) Title() (String, error) {
	val, err := c.Com().GetProperty("Title")
	return helpers.CastValue[String](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_TITLE_PROPERTY_COLUMN.html */
func (c NotesViewColumn) SetTitle(v String) error {
	return c.Com().PutProperty("Title", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_WIDTH_PROPERTY_VIEWCOLUMN.html */
func (c NotesViewColumn) Width() (Integer, error) {
	val, err := c.Com().GetProperty("Width")
	return helpers.CastValue[Integer](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_WIDTH_PROPERTY_VIEWCOLUMN.html */
func (c NotesViewColumn) SetWidth(v Integer) error {
	return c.Com().PutProperty("Width", v)
}

/* --------------------------------- Methods ------------------------------------ */
