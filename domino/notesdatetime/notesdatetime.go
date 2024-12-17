/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_NOTESDATETIME_CLASS.html */
package notesdatetime

import (
	"github.com/monstermichl/domigo/domino"
	"github.com/monstermichl/domigo/helpers"

	ole "github.com/go-ole/go-ole"
)

type NotesDateTime struct {
	domino.NotesStruct
}

func New(dispatchPtr *ole.IDispatch) NotesDateTime {
	return NotesDateTime{domino.NewNotesStruct(dispatchPtr)}
}

/* --------------------------------- Properties --------------------------------- */
/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_DATEONLY_PROPERTY.html */
func (d NotesDateTime) DateOnly() (domino.String, error) {
	val, err := d.Com().GetProperty("DateOnly")
	return helpers.CastValue[domino.String](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_GMTTIME_PROPERTY.html */
func (d NotesDateTime) GMTTime() (domino.String, error) {
	val, err := d.Com().GetProperty("GMTTime")
	return helpers.CastValue[domino.String](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISDST_PROPERTY.html */
func (d NotesDateTime) IsDST() (domino.Boolean, error) {
	val, err := d.Com().GetProperty("IsDST")
	return helpers.CastValue[domino.Boolean](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ISVALIDDATE_PROPERTY_DATETIME.html */
func (d NotesDateTime) IsValidDate() (domino.Boolean, error) {
	val, err := d.Com().GetProperty("IsValidDate")
	return helpers.CastValue[domino.Boolean](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_LOCALTIME_PROPERTY.html */
func (d NotesDateTime) LocalTime() (domino.String, error) {
	val, err := d.Com().GetProperty("LocalTime")
	return helpers.CastValue[domino.String](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_LOCALTIME_PROPERTY.html */
func (d NotesDateTime) SetLocalTime(v domino.String) error {
	return d.Com().PutProperty("LocalTime", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_LSGMTTIME_PROPERTY.html */
func (d NotesDateTime) LSGMTTime() (domino.Time, error) {
	val, err := d.Com().GetProperty("LSGMTTime")
	return helpers.CastValue[domino.Time](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_LSLOCALTIME_PROPERTY.html */
func (d NotesDateTime) LSLocalTime() (domino.Time, error) {
	val, err := d.Com().GetProperty("LSLocalTime")
	return helpers.CastValue[domino.Time](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_LSLOCALTIME_PROPERTY.html */
func (d NotesDateTime) SetLSLocalTime(v domino.Time) error {
	return d.Com().PutProperty("LSLocalTime", v)
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_TIMEONLY_PROPERTY.html */
func (d NotesDateTime) TimeOnly() (domino.String, error) {
	val, err := d.Com().GetProperty("TimeOnly")
	return helpers.CastValue[domino.String](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_TIMEZONE_PROPERTY.html */
func (d NotesDateTime) TimeZone() (domino.Long, error) {
	val, err := d.Com().GetProperty("TimeZone")
	return helpers.CastValue[domino.Long](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ZONETIME_PROPERTY.html */
func (d NotesDateTime) ZoneTime() (domino.String, error) {
	val, err := d.Com().GetProperty("ZoneTime")
	return helpers.CastValue[domino.String](val), err
}

/* --------------------------------- Methods ------------------------------------ */
/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ADJUSTDAY_METHOD.html */
type adjustDayParams struct {
	preserveLocalTime *domino.Boolean
}

type adjustDayParam func(*adjustDayParams)

func WithAdjustDayPreserveLocalTime(preserveLocalTime domino.Boolean) adjustDayParam {
	return func(c *adjustDayParams) {
		c.preserveLocalTime = &preserveLocalTime
	}
}

func (d NotesDateTime) AdjustDay(n domino.Integer, params ...adjustDayParam) error {
	s := &adjustDayParams{}
	paramsOrdered := []interface{}{n}

	for _, p := range params {
		p(s)
	}

	if s.preserveLocalTime != nil {
		paramsOrdered = append(paramsOrdered, *s.preserveLocalTime)
	}
	_, err := d.Com().CallMethod("AdjustDay", paramsOrdered...)
	return err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ADJUSTHOUR_METHOD.html */
type adjustHourParams struct {
	preserveLocalTime *domino.Boolean
}

type adjustHourParam func(*adjustHourParams)

func WithAdjustHourPreserveLocalTime(preserveLocalTime domino.Boolean) adjustHourParam {
	return func(c *adjustHourParams) {
		c.preserveLocalTime = &preserveLocalTime
	}
}

func (d NotesDateTime) AdjustHour(n domino.Integer, params ...adjustHourParam) error {
	s := &adjustHourParams{}
	paramsOrdered := []interface{}{n}

	for _, p := range params {
		p(s)
	}

	if s.preserveLocalTime != nil {
		paramsOrdered = append(paramsOrdered, *s.preserveLocalTime)
	}
	_, err := d.Com().CallMethod("AdjustHour", paramsOrdered...)
	return err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ADJUSTMINUTE_METHOD.html */
type adjustMinuteParams struct {
	preserveLocalTime *domino.Boolean
}

type adjustMinuteParam func(*adjustMinuteParams)

func WithAdjustMinutePreserveLocalTime(preserveLocalTime domino.Boolean) adjustMinuteParam {
	return func(c *adjustMinuteParams) {
		c.preserveLocalTime = &preserveLocalTime
	}
}

func (d NotesDateTime) AdjustMinute(n domino.Integer, params ...adjustMinuteParam) error {
	s := &adjustMinuteParams{}
	paramsOrdered := []interface{}{n}

	for _, p := range params {
		p(s)
	}

	if s.preserveLocalTime != nil {
		paramsOrdered = append(paramsOrdered, *s.preserveLocalTime)
	}
	_, err := d.Com().CallMethod("AdjustMinute", paramsOrdered...)
	return err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ADJUSTMONTH_METHOD.html */
type adjustMonthParams struct {
	preserveLocalTime *domino.Boolean
}

type adjustMonthParam func(*adjustMonthParams)

func WithAdjustMonthPreserveLocalTime(preserveLocalTime domino.Boolean) adjustMonthParam {
	return func(c *adjustMonthParams) {
		c.preserveLocalTime = &preserveLocalTime
	}
}

func (d NotesDateTime) AdjustMonth(n domino.Integer, params ...adjustMonthParam) error {
	s := &adjustMonthParams{}
	paramsOrdered := []interface{}{n}

	for _, p := range params {
		p(s)
	}

	if s.preserveLocalTime != nil {
		paramsOrdered = append(paramsOrdered, *s.preserveLocalTime)
	}
	_, err := d.Com().CallMethod("AdjustMonth", paramsOrdered...)
	return err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ADJUSTSECOND_METHOD.html */
type adjustSecondParams struct {
	preserveLocalTime *domino.Boolean
}

type adjustSecondParam func(*adjustSecondParams)

func WithAdjustSecondPreserveLocalTime(preserveLocalTime domino.Boolean) adjustSecondParam {
	return func(c *adjustSecondParams) {
		c.preserveLocalTime = &preserveLocalTime
	}
}

func (d NotesDateTime) AdjustSecond(n domino.Integer, params ...adjustSecondParam) error {
	s := &adjustSecondParams{}
	paramsOrdered := []interface{}{n}

	for _, p := range params {
		p(s)
	}

	if s.preserveLocalTime != nil {
		paramsOrdered = append(paramsOrdered, *s.preserveLocalTime)
	}
	_, err := d.Com().CallMethod("AdjustSecond", paramsOrdered...)
	return err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_ADJUSTYEAR_METHOD.html */
type adjustYearParams struct {
	preserveLocalTime *domino.Boolean
}

type adjustYearParam func(*adjustYearParams)

func WithAdjustYearPreserveLocalTime(preserveLocalTime domino.Boolean) adjustYearParam {
	return func(c *adjustYearParams) {
		c.preserveLocalTime = &preserveLocalTime
	}
}

func (d NotesDateTime) AdjustYear(n domino.Integer, params ...adjustYearParam) error {
	s := &adjustYearParams{}
	paramsOrdered := []interface{}{n}

	for _, p := range params {
		p(s)
	}

	if s.preserveLocalTime != nil {
		paramsOrdered = append(paramsOrdered, *s.preserveLocalTime)
	}
	_, err := d.Com().CallMethod("AdjustYear", paramsOrdered...)
	return err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_CONVERTTOZONE_METHOD.html */
func (d NotesDateTime) ConvertToZone(newzone domino.Integer, dst domino.Boolean) error {
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
func (d NotesDateTime) TimeDifference(notesDateTime NotesDateTime) (domino.Long, error) {
	val, err := d.Com().CallMethod("TimeDifference", notesDateTime.Com().Dispatch())
	return helpers.CastValue[domino.Long](val), err
}

/* https://help.hcl-software.com/dom_designer/14.0.0/basic/H_TIMEDIFFERENCEDOUBLE_METHOD.html */
func (d NotesDateTime) TimeDifferenceDouble(notesDateTime NotesDateTime) (domino.Double, error) {
	val, err := d.Com().CallMethod("TimeDifferenceDouble", notesDateTime.Com().Dispatch())
	return helpers.CastValue[domino.Double](val), err
}
