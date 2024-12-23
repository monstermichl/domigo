/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_NOTESVIEWCOLUMN_CLASS.html */
package domigo

import (
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
	return getComProperty[Integer](c, "Alignment")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ALIGNMENT_PROPERTY_2919_ABOUT.html */
func (c NotesViewColumn) SetAlignment(v Integer) error {
	return putComProperty(c, "Alignment", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_DATEFMT_PROPERTY_4095_ABOUT.html */
func (c NotesViewColumn) DateFmt() (Integer, error) {
	return getComProperty[Integer](c, "DateFmt")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_DATEFMT_PROPERTY_4095_ABOUT.html */
func (c NotesViewColumn) SetDateFmt(v Integer) error {
	return putComProperty(c, "DateFmt", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_FONTCOLOR_PROPERTY_9482_ABOUT.html */
func (c NotesViewColumn) FontColor() (Integer, error) {
	return getComProperty[Integer](c, "FontColor")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_FONTCOLOR_PROPERTY_9482_ABOUT.html */
func (c NotesViewColumn) SetFontColor(v Integer) error {
	return putComProperty(c, "FontColor", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_FONTFACE_PROPERTY_5960_ABOUT.html */
func (c NotesViewColumn) FontFace() (String, error) {
	return getComProperty[String](c, "FontFace")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_FONTFACE_PROPERTY_5960_ABOUT.html */
func (c NotesViewColumn) SetFontFace(v String) error {
	return putComProperty(c, "FontFace", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_FONTPOINTSIZE_PROPERTY_7819_ABOUT.html */
func (c NotesViewColumn) FontPointSize() (Integer, error) {
	return getComProperty[Integer](c, "FontPointSize")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_FONTPOINTSIZE_PROPERTY_7819_ABOUT.html */
func (c NotesViewColumn) SetFontPointSize(v Integer) error {
	return putComProperty(c, "FontPointSize", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_FONTSTYLE_PROPERTY_2400_ABOUT.html */
func (c NotesViewColumn) FontStyle() (Integer, error) {
	return getComProperty[Integer](c, "FontStyle")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_FONTSTYLE_PROPERTY_2400_ABOUT.html */
func (c NotesViewColumn) SetFontStyle(v Integer) error {
	return putComProperty(c, "FontStyle", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_FORMULA_PROPERTY.html */
func (c NotesViewColumn) Formula() (String, error) {
	return getComProperty[String](c, "Formula")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_FORMULA_PROPERTY.html */
func (c NotesViewColumn) SetFormula(v String) error {
	return putComProperty(c, "Formula", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_HEADER_ALIGNMENT_PROPERTY_4751_ABOUT.html */
func (c NotesViewColumn) HeaderAlignment() (Integer, error) {
	return getComProperty[Integer](c, "HeaderAlignment")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_HEADER_ALIGNMENT_PROPERTY_4751_ABOUT.html */
func (c NotesViewColumn) SetHeaderAlignment(v Integer) error {
	return putComProperty(c, "HeaderAlignment", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_HEADERFONTCOLOR_PROPERTY_VC.html */
func (c NotesViewColumn) HeaderFontColor() (Integer, error) {
	return getComProperty[Integer](c, "HeaderFontColor")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_HEADERFONTCOLOR_PROPERTY_VC.html */
func (c NotesViewColumn) SetHeaderFontColor(v Integer) error {
	return putComProperty(c, "HeaderFontColor", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_HEADERFONTFACE_PROPERTY_VC.html */
func (c NotesViewColumn) HeaderFontFace() (String, error) {
	return getComProperty[String](c, "HeaderFontFace")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_HEADERFONTFACE_PROPERTY_VC.html */
func (c NotesViewColumn) SetHeaderFontFace(v String) error {
	return putComProperty(c, "HeaderFontFace", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_HEADERFONTPOINTSIZE_PROPERTY_VC.html */
func (c NotesViewColumn) HeaderFontPointSize() (Integer, error) {
	return getComProperty[Integer](c, "HeaderFontPointSize")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_HEADERFONTPOINTSIZE_PROPERTY_VC.html */
func (c NotesViewColumn) SetHeaderFontPointSize(v Integer) error {
	return putComProperty(c, "HeaderFontPointSize", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_HEADERFONTSTYLE_PROPERTY_VC.html */
func (c NotesViewColumn) HeaderFontStyle() (Integer, error) {
	return getComProperty[Integer](c, "HeaderFontStyle")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_HEADERFONTSTYLE_PROPERTY_VC.html */
func (c NotesViewColumn) SetHeaderFontStyle(v Integer) error {
	return putComProperty(c, "HeaderFontStyle", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISACCENTINSENSITIVESORT_1949_ABOUT.html */
func (c NotesViewColumn) IsAccentSensitiveSort() (Boolean, error) {
	return getComProperty[Boolean](c, "IsAccentSensitiveSort")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISACCENTINSENSITIVESORT_1949_ABOUT.html */
func (c NotesViewColumn) SetIsAccentSensitiveSort(v Boolean) error {
	return putComProperty(c, "IsAccentSensitiveSort", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISCASEINSENSITIVESORT_PROPERTY_1451_ABOUT.html */
func (c NotesViewColumn) IsCaseSensitiveSort() (Boolean, error) {
	return getComProperty[Boolean](c, "IsCaseSensitiveSort")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISCASEINSENSITIVESORT_PROPERTY_1451_ABOUT.html */
func (c NotesViewColumn) SetIsCaseSensitiveSort(v Boolean) error {
	return putComProperty(c, "IsCaseSensitiveSort", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISCATEGORY_PROPERTY.html */
func (c NotesViewColumn) IsCategory() (Boolean, error) {
	return getComProperty[Boolean](c, "IsCategory")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISFIELD_PROPERTY_3448_ABOUT.html */
func (c NotesViewColumn) IsField() (Boolean, error) {
	return getComProperty[Boolean](c, "IsField")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISFONTBOLD_PROPERTY_VC.html */
func (c NotesViewColumn) IsFontBold() (Boolean, error) {
	return getComProperty[Boolean](c, "IsFontBold")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISFONTBOLD_PROPERTY_VC.html */
func (c NotesViewColumn) SetIsFontBold(v Boolean) error {
	return putComProperty(c, "IsFontBold", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISFONTITALIC_PROPERTY_VC.html */
func (c NotesViewColumn) IsFontItalic() (Boolean, error) {
	return getComProperty[Boolean](c, "IsFontItalic")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISFONTITALIC_PROPERTY_VC.html */
func (c NotesViewColumn) SetIsFontItalic(v Boolean) error {
	return putComProperty(c, "IsFontItalic", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISFONTSTRIKETHROUGH_PROPERTY_VC.html */
func (c NotesViewColumn) IsFontStrikethrough() (Boolean, error) {
	return getComProperty[Boolean](c, "IsFontStrikethrough")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISFONTSTRIKETHROUGH_PROPERTY_VC.html */
func (c NotesViewColumn) SetIsFontStrikethrough(v Boolean) error {
	return putComProperty(c, "IsFontStrikethrough", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISFONTUNDERLINE_PROPERTY_VC.html */
func (c NotesViewColumn) IsFontUnderline() (Boolean, error) {
	return getComProperty[Boolean](c, "IsFontUnderline")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISFONTUNDERLINE_PROPERTY_VC.html */
func (c NotesViewColumn) SetIsFontUnderline(v Boolean) error {
	return putComProperty(c, "IsFontUnderline", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISFORMULA_PROPERTY.html */
func (c NotesViewColumn) IsFormula() (Boolean, error) {
	return getComProperty[Boolean](c, "IsFormula")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISHEADERFONTBOLD_PROPERTY_VC.html */
func (c NotesViewColumn) IsHeaderFontBold() (Boolean, error) {
	return getComProperty[Boolean](c, "IsHeaderFontBold")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISHEADERFONTBOLD_PROPERTY_VC.html */
func (c NotesViewColumn) SetIsHeaderFontBold(v Boolean) error {
	return putComProperty(c, "IsHeaderFontBold", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISHEADERFONTITALIC_PROPERTY_VC.html */
func (c NotesViewColumn) IsHeaderFontItalic() (Boolean, error) {
	return getComProperty[Boolean](c, "IsHeaderFontItalic")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISHEADERFONTITALIC_PROPERTY_VC.html */
func (c NotesViewColumn) SetIsHeaderFontItalic(v Boolean) error {
	return putComProperty(c, "IsHeaderFontItalic", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISHEADERFONTSTRIKETHROUGH_PROPERTY_VC.html */
func (c NotesViewColumn) IsHeaderFontStrikethrough() (Boolean, error) {
	return getComProperty[Boolean](c, "IsHeaderFontStrikethrough")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISHEADERFONTSTRIKETHROUGH_PROPERTY_VC.html */
func (c NotesViewColumn) SetIsHeaderFontStrikethrough(v Boolean) error {
	return putComProperty(c, "IsHeaderFontStrikethrough", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISHEADERFONTUNDERLINE_PROPERTY_VC.html */
func (c NotesViewColumn) IsHeaderFontUnderline() (Boolean, error) {
	return getComProperty[Boolean](c, "IsHeaderFontUnderline")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISHEADERFONTUNDERLINE_PROPERTY_VC.html */
func (c NotesViewColumn) SetIsHeaderFontUnderline(v Boolean) error {
	return putComProperty(c, "IsHeaderFontUnderline", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISHIDDEN_PROPERTY.html */
func (c NotesViewColumn) IsHidden() (Boolean, error) {
	return getComProperty[Boolean](c, "IsHidden")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISHIDDEN_PROPERTY.html */
func (c NotesViewColumn) SetIsHidden(v Boolean) error {
	return putComProperty(c, "IsHidden", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISHIDEDETAIL_PROPERTY.html */
func (c NotesViewColumn) IsHideDetail() (Boolean, error) {
	return getComProperty[Boolean](c, "IsHideDetail")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISHIDEDETAIL_PROPERTY.html */
func (c NotesViewColumn) SetIsHideDetail(v Boolean) error {
	return putComProperty(c, "IsHideDetail", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISICON_PROPERTY.html */
func (c NotesViewColumn) IsIcon() (Boolean, error) {
	return getComProperty[Boolean](c, "IsIcon")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISNUMBERATTRIBPARENS_PROPERTY_VC.html */
func (c NotesViewColumn) IsNumberAttribParens() (Boolean, error) {
	return getComProperty[Boolean](c, "IsNumberAttribParens")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISNUMBERATTRIBPARENS_PROPERTY_VC.html */
func (c NotesViewColumn) SetIsNumberAttribParens(v Boolean) error {
	return putComProperty(c, "IsNumberAttribParens", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISNUMBERATTRIBPERCENT_PROPERTY_VC.html */
func (c NotesViewColumn) IsNumberAttribPercent() (Boolean, error) {
	return getComProperty[Boolean](c, "IsNumberAttribPercent")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISNUMBERATTRIBPERCENT_PROPERTY_VC.html */
func (c NotesViewColumn) SetIsNumberAttribPercent(v Boolean) error {
	return putComProperty(c, "IsNumberAttribPercent", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISNUMBERATTRIBPUNCTUATED_PROPERTY_VC.html */
func (c NotesViewColumn) IsNumberAttribPunctuated() (Boolean, error) {
	return getComProperty[Boolean](c, "IsNumberAttribPunctuated")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISNUMBERATTRIBPUNCTUATED_PROPERTY_VC.html */
func (c NotesViewColumn) SetIsNumberAttribPunctuated(v Boolean) error {
	return putComProperty(c, "IsNumberAttribPunctuated", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISRESIZE_PROPERTY.html */
func (c NotesViewColumn) IsResize() (Boolean, error) {
	return getComProperty[Boolean](c, "IsResize")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISRESIZE_PROPERTY.html */
func (c NotesViewColumn) SetIsResize(v Boolean) error {
	return putComProperty(c, "IsResize", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISRESORTASCENDING_PROPERTY.html */
func (c NotesViewColumn) IsResortAscending() (Boolean, error) {
	return getComProperty[Boolean](c, "IsResortAscending")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISRESORTASCENDING_PROPERTY.html */
func (c NotesViewColumn) SetIsResortAscending(v Boolean) error {
	return putComProperty(c, "IsResortAscending", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISRESORTDESCENDING_PROPERTY.html */
func (c NotesViewColumn) IsResortDescending() (Boolean, error) {
	return getComProperty[Boolean](c, "IsResortDescending")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISRESORTDESCENDING_PROPERTY.html */
func (c NotesViewColumn) SetIsResortDescending(v Boolean) error {
	return putComProperty(c, "IsResortDescending", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISRESORTTOVIEW_PROPERTY.html */
func (c NotesViewColumn) IsResortToView() (Boolean, error) {
	return getComProperty[Boolean](c, "IsResortToView")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISRESORTTOVIEW_PROPERTY.html */
func (c NotesViewColumn) SetIsResortToView(v Boolean) error {
	return putComProperty(c, "IsResortToView", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISRESPONSE_PROPERTY_COLUMN.html */
func (c NotesViewColumn) IsResponse() (Boolean, error) {
	return getComProperty[Boolean](c, "IsResponse")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISSECONDARYRESORT_PROPERTY.html */
func (c NotesViewColumn) IsSecondaryResort() (Boolean, error) {
	return getComProperty[Boolean](c, "IsSecondaryResort")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISSECONDARYRESORT_PROPERTY.html */
func (c NotesViewColumn) SetIsSecondaryResort(v Boolean) error {
	return putComProperty(c, "IsSecondaryResort", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISSECONDARYRESORTDESCENDING_PROPERTY.html */
func (c NotesViewColumn) IsSecondaryResortDescending() (Boolean, error) {
	return getComProperty[Boolean](c, "IsSecondaryResortDescending")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISSECONDARYRESORTDESCENDING_PROPERTY.html */
func (c NotesViewColumn) SetIsSecondaryResortDescending(v Boolean) error {
	return putComProperty(c, "IsSecondaryResortDescending", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISSHOWTWISTIE_PROPERTY.html */
func (c NotesViewColumn) IsShowTwistie() (Boolean, error) {
	return getComProperty[Boolean](c, "IsShowTwistie")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISSHOWTWISTIE_PROPERTY.html */
func (c NotesViewColumn) SetIsShowTwistie(v Boolean) error {
	return putComProperty(c, "IsShowTwistie", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISSORTDESCENDING_PROPERTY.html */
func (c NotesViewColumn) IsSortDescending() (Boolean, error) {
	return getComProperty[Boolean](c, "IsSortDescending")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISSORTDESCENDING_PROPERTY.html */
func (c NotesViewColumn) SetIsSortDescending(v Boolean) error {
	return putComProperty(c, "IsSortDescending", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISSORTED_PROPERTY_COLUMN.html */
func (c NotesViewColumn) IsSorted() (Boolean, error) {
	return getComProperty[Boolean](c, "IsSorted")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISSORTED_PROPERTY_COLUMN.html */
func (c NotesViewColumn) SetIsSorted(v Boolean) error {
	return putComProperty(c, "IsSorted", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ITEMNAME_PROPERTY.html */
func (c NotesViewColumn) ItemName() (String, error) {
	return getComProperty[String](c, "ItemName")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_LISTSEP_PROPERTY.html */
func (c NotesViewColumn) ListSep() (Integer, error) {
	return getComProperty[Integer](c, "ListSep")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_LISTSEP_PROPERTY.html */
func (c NotesViewColumn) SetListSep(v Integer) error {
	return putComProperty(c, "ListSep", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_NUMBERATTRIB_PROPERTY.html */
func (c NotesViewColumn) NumberAttrib() (Integer, error) {
	return getComProperty[Integer](c, "NumberAttrib")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_NUMBERATTRIB_PROPERTY.html */
func (c NotesViewColumn) SetNumberAttrib(v Integer) error {
	return putComProperty(c, "NumberAttrib", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_NUMBERDIGITS_PROPERTY.html */
func (c NotesViewColumn) NumberDigits() (Integer, error) {
	return getComProperty[Integer](c, "NumberDigits")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_NUMBERDIGITS_PROPERTY.html */
func (c NotesViewColumn) SetNumberDigits(v Integer) error {
	return putComProperty(c, "NumberDigits", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_NUMBERFORMAT_PROPERTY.html */
func (c NotesViewColumn) NumberFormat() (Integer, error) {
	return getComProperty[Integer](c, "NumberFormat")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_NUMBERFORMAT_PROPERTY.html */
func (c NotesViewColumn) SetNumberFormat(v Integer) error {
	return putComProperty(c, "NumberFormat", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_PARENT_PROPERTY_NOTESVIEWCOLUMN_CLASS.html */
func (c NotesViewColumn) Parent() (NotesView, error) {
	return getComObjectProperty(c, NewNotesView, "Parent")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_POSITION_PROPERTY.html */
func (c NotesViewColumn) Position() (Integer, error) {
	return getComProperty[Integer](c, "Position")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_RESORTTOVIEWNAME_PROPERTY.html */
func (c NotesViewColumn) ResortToViewName() (String, error) {
	return getComProperty[String](c, "ResortToViewName")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_RESORTTOVIEWNAME_PROPERTY.html */
func (c NotesViewColumn) SetResortToViewName(v String) error {
	return putComProperty(c, "ResortToViewName", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SECONDARYRESORTCOLUMNINDEX_PROPERTY_VC.html */
func (c NotesViewColumn) SecondaryResortColumnIndex() (Integer, error) {
	return getComProperty[Integer](c, "SecondaryResortColumnIndex")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SECONDARYRESORTCOLUMNINDEX_PROPERTY_VC.html */
func (c NotesViewColumn) SetSecondaryResortColumnIndex(v Integer) error {
	return putComProperty(c, "SecondaryResortColumnIndex", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_TIMEDATEFMT_PROPERTY.html */
func (c NotesViewColumn) TimeDateFmt() (Integer, error) {
	return getComProperty[Integer](c, "TimeDateFmt")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_TIMEDATEFMT_PROPERTY.html */
func (c NotesViewColumn) SetTimeDateFmt(v Integer) error {
	return putComProperty(c, "TimeDateFmt", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_TIMEFMT_PROPERTY.html */
func (c NotesViewColumn) TimeFmt() (Integer, error) {
	return getComProperty[Integer](c, "TimeFmt")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_TIMEFMT_PROPERTY.html */
func (c NotesViewColumn) SetTimeFmt(v Integer) error {
	return putComProperty(c, "TimeFmt", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_TIMEZONEFMT.html */
func (c NotesViewColumn) TimeZoneFmt() (Integer, error) {
	return getComProperty[Integer](c, "TimeZoneFmt")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_TIMEZONEFMT.html */
func (c NotesViewColumn) SetTimeZoneFmt(v Integer) error {
	return putComProperty(c, "TimeZoneFmt", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_TITLE_PROPERTY_COLUMN.html */
func (c NotesViewColumn) Title() (String, error) {
	return getComProperty[String](c, "Title")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_TITLE_PROPERTY_COLUMN.html */
func (c NotesViewColumn) SetTitle(v String) error {
	return putComProperty(c, "Title", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_WIDTH_PROPERTY_VIEWCOLUMN.html */
func (c NotesViewColumn) Width() (Integer, error) {
	return getComProperty[Integer](c, "Width")
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_WIDTH_PROPERTY_VIEWCOLUMN.html */
func (c NotesViewColumn) SetWidth(v Integer) error {
	return putComProperty(c, "Width", v)
}

/* --------------------------------- Methods ------------------------------------ */
