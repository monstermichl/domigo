/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_NOTESDATETIME_CLASS.html */
package domigo

import (
	"github.com/monstermichl/domigo/helpers"

	ole "github.com/go-ole/go-ole"
)

type NotesDateTime struct {
	NotesStruct
}

func NewNotesDateTime(dispatchPtr *ole.IDispatch) NotesDateTime {
	return NotesDateTime{NewNotesStruct(dispatchPtr)}
}

/* --------------------------------- Properties --------------------------------- */
/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_DATEONLY_PROPERTY.html */
func (d NotesDateTime) DateOnly() (String, error) {
	val, err := d.Com().GetProperty("DateOnly")
	return helpers.CastValue[String](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GMTTIME_PROPERTY.html */
func (d NotesDateTime) GMTTime() (String, error) {
	val, err := d.Com().GetProperty("GMTTime")
	return helpers.CastValue[String](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISDST_PROPERTY.html */
func (d NotesDateTime) IsDST() (Boolean, error) {
	val, err := d.Com().GetProperty("IsDST")
	return helpers.CastValue[Boolean](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISVALIDDATE_PROPERTY_DATETIME.html */
func (d NotesDateTime) IsValidDate() (Boolean, error) {
	val, err := d.Com().GetProperty("IsValidDate")
	return helpers.CastValue[Boolean](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_LOCALTIME_PROPERTY.html */
func (d NotesDateTime) LocalTime() (String, error) {
	val, err := d.Com().GetProperty("LocalTime")
	return helpers.CastValue[String](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_LOCALTIME_PROPERTY.html */
func (d NotesDateTime) SetLocalTime(v String) error {
	return d.Com().PutProperty("LocalTime", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_LSGMTTIME_PROPERTY.html */
func (d NotesDateTime) LSGMTTime() (Variant, error) {
	val, err := d.Com().GetProperty("LSGMTTime")
	return helpers.CastValue[Variant](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_LSLOCALTIME_PROPERTY.html */
func (d NotesDateTime) LSLocalTime() (Variant, error) {
	val, err := d.Com().GetProperty("LSLocalTime")
	return helpers.CastValue[Variant](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_LSLOCALTIME_PROPERTY.html */
func (d NotesDateTime) SetLSLocalTime(v Variant) error {
	return d.Com().PutProperty("LSLocalTime", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_PARENT_PROPERTY_DATETIME_COM.html */
func (d NotesDateTime) Parent() (NotesSession, error) {
	dispatchPtr, err := d.Com().GetObjectProperty("Parent")
	return NewNotesSession(dispatchPtr), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_TIMEONLY_PROPERTY.html */
func (d NotesDateTime) TimeOnly() (String, error) {
	val, err := d.Com().GetProperty("TimeOnly")
	return helpers.CastValue[String](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_TIMEZONE_PROPERTY.html */
func (d NotesDateTime) TimeZone() (Long, error) {
	val, err := d.Com().GetProperty("TimeZone")
	return helpers.CastValue[Long](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ZONETIME_PROPERTY.html */
func (d NotesDateTime) ZoneTime() (String, error) {
	val, err := d.Com().GetProperty("ZoneTime")
	return helpers.CastValue[String](val), err
}

/* --------------------------------- Methods ------------------------------------ */
/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ADJUSTDAY_METHOD.html */
type notesDateTimeAdjustDayParams struct {
	preserveLocalTime *Boolean
}

type notesDateTimeAdjustDayParam func(*notesDateTimeAdjustDayParams)

func WithNotesDateTimeAdjustDayPreserveLocalTime(preserveLocalTime Boolean) notesDateTimeAdjustDayParam {
	return func(c *notesDateTimeAdjustDayParams) {
		c.preserveLocalTime = &preserveLocalTime
	}
}

func (d NotesDateTime) AdjustDay(n Integer, params ...notesDateTimeAdjustDayParam) error {
	paramsStruct := &notesDateTimeAdjustDayParams{}
	paramsOrdered := []interface{}{n}

	for _, p := range params {
		p(paramsStruct)
	}

	if paramsStruct.preserveLocalTime != nil {
		paramsOrdered = append(paramsOrdered, *paramsStruct.preserveLocalTime)
	}
	_, err := d.Com().CallMethod("AdjustDay", paramsOrdered...)
	return err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ADJUSTHOUR_METHOD.html */
type notesDateTimeAdjustHourParams struct {
	preserveLocalTime *Boolean
}

type notesDateTimeAdjustHourParam func(*notesDateTimeAdjustHourParams)

func WithNotesDateTimeAdjustHourPreserveLocalTime(preserveLocalTime Boolean) notesDateTimeAdjustHourParam {
	return func(c *notesDateTimeAdjustHourParams) {
		c.preserveLocalTime = &preserveLocalTime
	}
}

func (d NotesDateTime) AdjustHour(n Integer, params ...notesDateTimeAdjustHourParam) error {
	paramsStruct := &notesDateTimeAdjustHourParams{}
	paramsOrdered := []interface{}{n}

	for _, p := range params {
		p(paramsStruct)
	}

	if paramsStruct.preserveLocalTime != nil {
		paramsOrdered = append(paramsOrdered, *paramsStruct.preserveLocalTime)
	}
	_, err := d.Com().CallMethod("AdjustHour", paramsOrdered...)
	return err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ADJUSTMINUTE_METHOD.html */
type notesDateTimeAdjustMinuteParams struct {
	preserveLocalTime *Boolean
}

type notesDateTimeAdjustMinuteParam func(*notesDateTimeAdjustMinuteParams)

func WithNotesDateTimeAdjustMinutePreserveLocalTime(preserveLocalTime Boolean) notesDateTimeAdjustMinuteParam {
	return func(c *notesDateTimeAdjustMinuteParams) {
		c.preserveLocalTime = &preserveLocalTime
	}
}

func (d NotesDateTime) AdjustMinute(n Integer, params ...notesDateTimeAdjustMinuteParam) error {
	paramsStruct := &notesDateTimeAdjustMinuteParams{}
	paramsOrdered := []interface{}{n}

	for _, p := range params {
		p(paramsStruct)
	}

	if paramsStruct.preserveLocalTime != nil {
		paramsOrdered = append(paramsOrdered, *paramsStruct.preserveLocalTime)
	}
	_, err := d.Com().CallMethod("AdjustMinute", paramsOrdered...)
	return err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ADJUSTMONTH_METHOD.html */
type notesDateTimeAdjustMonthParams struct {
	preserveLocalTime *Boolean
}

type notesDateTimeAdjustMonthParam func(*notesDateTimeAdjustMonthParams)

func WithNotesDateTimeAdjustMonthPreserveLocalTime(preserveLocalTime Boolean) notesDateTimeAdjustMonthParam {
	return func(c *notesDateTimeAdjustMonthParams) {
		c.preserveLocalTime = &preserveLocalTime
	}
}

func (d NotesDateTime) AdjustMonth(n Integer, params ...notesDateTimeAdjustMonthParam) error {
	paramsStruct := &notesDateTimeAdjustMonthParams{}
	paramsOrdered := []interface{}{n}

	for _, p := range params {
		p(paramsStruct)
	}

	if paramsStruct.preserveLocalTime != nil {
		paramsOrdered = append(paramsOrdered, *paramsStruct.preserveLocalTime)
	}
	_, err := d.Com().CallMethod("AdjustMonth", paramsOrdered...)
	return err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ADJUSTSECOND_METHOD.html */
type notesDateTimeAdjustSecondParams struct {
	preserveLocalTime *Boolean
}

type notesDateTimeAdjustSecondParam func(*notesDateTimeAdjustSecondParams)

func WithNotesDateTimeAdjustSecondPreserveLocalTime(preserveLocalTime Boolean) notesDateTimeAdjustSecondParam {
	return func(c *notesDateTimeAdjustSecondParams) {
		c.preserveLocalTime = &preserveLocalTime
	}
}

func (d NotesDateTime) AdjustSecond(n Integer, params ...notesDateTimeAdjustSecondParam) error {
	paramsStruct := &notesDateTimeAdjustSecondParams{}
	paramsOrdered := []interface{}{n}

	for _, p := range params {
		p(paramsStruct)
	}

	if paramsStruct.preserveLocalTime != nil {
		paramsOrdered = append(paramsOrdered, *paramsStruct.preserveLocalTime)
	}
	_, err := d.Com().CallMethod("AdjustSecond", paramsOrdered...)
	return err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ADJUSTYEAR_METHOD.html */
type notesDateTimeAdjustYearParams struct {
	preserveLocalTime *Boolean
}

type notesDateTimeAdjustYearParam func(*notesDateTimeAdjustYearParams)

func WithNotesDateTimeAdjustYearPreserveLocalTime(preserveLocalTime Boolean) notesDateTimeAdjustYearParam {
	return func(c *notesDateTimeAdjustYearParams) {
		c.preserveLocalTime = &preserveLocalTime
	}
}

func (d NotesDateTime) AdjustYear(n Integer, params ...notesDateTimeAdjustYearParam) error {
	paramsStruct := &notesDateTimeAdjustYearParams{}
	paramsOrdered := []interface{}{n}

	for _, p := range params {
		p(paramsStruct)
	}

	if paramsStruct.preserveLocalTime != nil {
		paramsOrdered = append(paramsOrdered, *paramsStruct.preserveLocalTime)
	}
	_, err := d.Com().CallMethod("AdjustYear", paramsOrdered...)
	return err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_CONVERTTOZONE_METHOD.html */
func (d NotesDateTime) ConvertToZone(newzone Integer, dst Boolean) error {
	_, err := d.Com().CallMethod("ConvertToZone", newzone, dst)
	return err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SETANYDATE_METHOD.html */
func (d NotesDateTime) SetAnyDate() error {
	_, err := d.Com().CallMethod("SetAnyDate")
	return err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SETANYTIME_METHOD.html */
func (d NotesDateTime) SetAnyTime() error {
	_, err := d.Com().CallMethod("SetAnyTime")
	return err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_SETNOW_METHOD.html */
func (d NotesDateTime) SetNow() error {
	_, err := d.Com().CallMethod("SetNow")
	return err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_TIMEDIFFERENCE_METHOD.html */
func (d NotesDateTime) TimeDifference(notesDateTime NotesDateTime) (Long, error) {
	val, err := d.Com().CallMethod("TimeDifference", notesDateTime.Com().Dispatch())
	return helpers.CastValue[Long](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_TIMEDIFFERENCEDOUBLE_METHOD.html */
func (d NotesDateTime) TimeDifferenceDouble(notesDateTime NotesDateTime) (Double, error) {
	val, err := d.Com().CallMethod("TimeDifferenceDouble", notesDateTime.Com().Dispatch())
	return helpers.CastValue[Double](val), err
}
