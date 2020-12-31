//   Copyright 2021 echotrue

//   Licensed under the Apache License, Version 2.0 (the "License");
//   you may not use this file except in compliance with the License.
//   You may obtain a copy of the License at
//
//       http://www.apache.org/licenses/LICENSE-2.0
//
//   Unless required by applicable law or agreed to in writing, software
//   distributed under the License is distributed on an "AS IS" BASIS,
//   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//   See the License for the specific language governing permissions and
//   limitations under the License.

//  package `timkit` is a time toolkit for Golang reference PHP's library `Carbon`
package timkit

import (
	"strings"
	"sync"
	"time"
)

const (
	secondsPerMinute  = 60
	minutesPerHour    = 60
	hoursPerDay       = 24
	daysPerWeek       = 7
	monthsPerQuarter  = 3
	monthsPerYear     = 12
	yearsPerCenturies = 100
	yearsPerDecade    = 10
	weeksPerLongYear  = 53
	daysInLeapYear    = 366
	daysInNormalYear  = 365
	secondsInWeek     = 691200
)

// String formats for dates
const (
	DefaultFormat = "2006-01-02 15:04:05"
	DateFormat    = "2006-01-02"
	TimeFormat    = "15:04:05"
)

// The TimeKit type represents a time instance
// Provides a simple API extension for Time
type TimeKit struct {
	time.Time
	format      string
	weekendDays []time.Weekday
	weekStartAt time.Weekday
	weekEndAt   time.Weekday
	lock        sync.Mutex
}

// NewTimeKit return a pointer to a new NewTimeKit instance
func NewTimeKit(t time.Time) *TimeKit {
	return NewOptions(OptionSetTime(t))
}

type Option func(t *TimeKit)

func NewOptions(opt ...Option) *TimeKit {
	t := TimeKit{
		Time:        time.Now(),
		format:      DefaultFormat,
		weekendDays: []time.Weekday{time.Saturday, time.Sunday},
		weekStartAt: time.Monday,
		weekEndAt:   time.Sunday,
	}

	for _, o := range opt {
		o(&t)
	}
	return &t
}

// Now return a new TimeKit instance for current time in local
func Now() *TimeKit {
	return NewTimeKit(time.Now())
}

// NowWithLocation return a new TimeKit instance for current time in given location
func NowWithLocation(loc string) (*TimeKit, error) {
	l, err := time.LoadLocation(loc)
	if err != nil {
		return nil, err
	}
	return NewTimeKit(Now().In(l)), nil
}

// Parse return a new TimeKit instance from a string
// layout reference `DefaultFormat` or time.ANSIC
// value reference `2020-12-23 11:13:11`
// location default `UTC` .
func Parse(layout, value, location string) (*TimeKit, error) {
	l, err := time.LoadLocation(location)
	if err != nil {
		return nil, err
	}
	t, err := time.ParseInLocation(layout, value, l)
	if err != nil {
		return nil, err
	}
	return NewTimeKit(t), nil
}

// CreateFromTimestamp return a new TimeKit instance from a timestamp
// If the location is invalid , it return an error
func CreateFromTimestamp(timestamp int64, location string) (*TimeKit, error) {
	l, err := time.LoadLocation(location)
	if err != nil {
		return nil, err
	}
	return createFromTimestamp(timestamp, l), nil
}

func createFromTimestamp(timestamp int64, l *time.Location) *TimeKit {
	return NewTimeKit(time.Unix(timestamp, 0).In(l))
}

// OptionSetFormat format the time with this format string
func OptionSetFormat(format string) Option {
	return func(t *TimeKit) {
		t.format = format
	}
}

// OptionSetWeekStartAt set the start of week
func OptionSetWeekStartAt(s time.Weekday) Option {
	return func(t *TimeKit) {
		t.weekStartAt = s
	}
}

// OptionSetWeekEndAt set end of week
func OptionSetWeekEndAt(e time.Weekday) Option {
	return func(t *TimeKit) {
		t.weekStartAt = e
	}
}

// OptionSetTime set the time by given sec
func OptionSetTime(datetime time.Time) Option {
	return func(t *TimeKit) {
		t.Time = datetime
	}
}

// SetFormat format the time with this format string
func (tk *TimeKit) SetFormat(format string) {
	tk.lock.Lock()
	defer tk.lock.Unlock()
	tk.format = format
}

// SetTime set the time by given sec
func (tk *TimeKit) SetTime(t time.Time) {
	tk.lock.Lock()
	defer tk.lock.Unlock()
	tk.Time = t
}

// SetWeekStartsAt set the start of week
func (tk *TimeKit) SetWeekStartsAt(start time.Weekday) {
	tk.lock.Lock()
	defer tk.lock.Unlock()
	tk.weekStartAt = start
}

// SetWeekEndsAt set end of week
func (tk *TimeKit) SetWeekEndsAt(end time.Weekday) {
	tk.lock.Lock()
	defer tk.lock.Unlock()
	tk.weekEndAt = end
}

// SetTimestamp set the time by given sec todo delete
func (tk *TimeKit) SetTimestamp(sec int64) {
	tk.SetTime(time.Unix(sec, 0).In(tk.Location()))
}

// AddCenturies add centuries from the current time
func (tk *TimeKit) AddCenturies(centuries int) *TimeKit {
	tk.SetTime(tk.AddDate(yearsPerCenturies*centuries, 0, 0))
	return tk
}

// AddCentury add a century from the current time
func (tk *TimeKit) AddCentury() *TimeKit {
	return tk.AddCenturies(1)
}

// SubCenturies remove centuries from the current time
func (tk *TimeKit) SubCenturies(centuries int) *TimeKit {
	return tk.AddCenturies(-centuries)
}

// SubCentury remove a century from the current time
func (tk *TimeKit) SubCentury() *TimeKit {
	return tk.SubCenturies(1)
}

// AddDays add days from the current time
func (tk *TimeKit) AddDays(d int) *TimeKit {
	tk.SetTime(tk.AddDate(0, 0, d))
	return tk
}

// AddDays add a day from the current time
func (tk *TimeKit) AddDay() *TimeKit {
	return tk.AddDays(1)
}

// SubDays remove days from the current time
func (tk *TimeKit) SubDays(d int) *TimeKit {
	return tk.AddDays(-d)
}

// SubDay remove a day from the current time
func (tk *TimeKit) SubDay() *TimeKit {
	return tk.SubDays(1)
}

// AddHours add hours from the current time
func (tk *TimeKit) AddHours(h int) *TimeKit {
	d := time.Duration(h) * time.Hour
	tk.SetTime(tk.Add(d))
	return tk
}

// AddHour add an hour from the current time
func (tk *TimeKit) AddHour() *TimeKit {
	return tk.AddHours(1)
}

// SubHours remove hours from current time
func (tk *TimeKit) SubHours(h int) *TimeKit {
	return tk.AddHours(-h)
}

// SubHour remove an hours from the current time
func (tk *TimeKit) SubHour() *TimeKit {
	return tk.SubHours(1)
}

// AddMinutes add minutes from current time
func (tk *TimeKit) AddMinutes(m int) *TimeKit {
	d := time.Duration(m) * time.Minute
	tk.SetTime(tk.Add(d))
	return tk
}

// AddMinute add a minute from current time
func (tk *TimeKit) AddMinute() *TimeKit {
	return tk.AddMinutes(1)
}

// SubMinutes remove minutes from current time
func (tk *TimeKit) SubMinutes(m int) *TimeKit {
	return tk.AddMinutes(-m)
}

// SubMinute remove a minute from current time
func (tk *TimeKit) SubMinute() *TimeKit {
	return tk.SubMinutes(1)
}

// AddMonths add months from current time
func (tk *TimeKit) AddMonths(m int) *TimeKit {
	tk.SetTime(tk.AddDate(0, m, 0))
	return tk
}

// AddMonth add a month from current time
func (tk *TimeKit) AddMonth() *TimeKit {
	return tk.AddMonths(1)
}

// SubMonths remove months from current time
func (tk *TimeKit) SubMonths(m int) *TimeKit {
	return tk.AddMonths(-m)
}

// SubMonth remove a month from current time
func (tk *TimeKit) SubMonth() *TimeKit {
	return tk.SubMonths(1)
}

// AddMonthsNoOverflow add months from current time .
// Not overflowing in case the days of new month is less than the current one
func (tk *TimeKit) AddMonthsNoOverflow(m int) *TimeKit {
	newDate := tk.AddDate(0, m, 0)
	if tk.Day() != newDate.Day() {
		newDate = tk.LastDayOfPreMonth()
	}
	tk.SetTime(newDate)
	return tk
}

// AddMonthNoOverflow add a month with no overflow from current time
func (tk *TimeKit) AddMonthNoOverflow() *TimeKit {
	return tk.AddMonthsNoOverflow(1)
}

// SubMonthsNoOverflow remove months with no overflow from current time
func (tk *TimeKit) SubMonthsNoOverflow(m int) *TimeKit {
	return tk.AddMonthsNoOverflow(-m)
}

// SubMonthNoOverflow remove a month with no overflow from current time
func (tk *TimeKit) SubMonthNoOverflow() *TimeKit {
	return tk.SubMonthsNoOverflow(1)
}

// AddQuarters add quarters from current time
func (tk *TimeKit) AddQuarters(q int) *TimeKit {
	tk.SetTime(tk.AddDate(0, monthsPerQuarter*q, 0))
	return tk
}

// AddQuarter add a quarter from current time
func (tk *TimeKit) AddQuarter() *TimeKit {
	return tk.AddQuarters(1)
}

// SubQuarters remove quarters from current time
func (tk *TimeKit) SubQuarters(q int) *TimeKit {
	return tk.AddQuarters(-q)
}

// SubQuarter remove a quarter from current time
func (tk *TimeKit) SubQuarter() *TimeKit {
	return tk.SubQuarters(1)
}

// AddSeconds add seconds from current time
func (tk *TimeKit) AddSeconds(s int) *TimeKit {
	d := time.Duration(s) * time.Second
	tk.SetTime(tk.Add(d))
	return tk
}

// AddSecond add a second from current time
func (tk *TimeKit) AddSecond() *TimeKit {
	return tk.AddSeconds(1)
}

// SubSeconds remove seconds from current time
func (tk *TimeKit) SubSeconds(s int) *TimeKit {
	return tk.AddSeconds(-s)
}

// SubSecond remove a second from current time
func (tk *TimeKit) SubSecond() *TimeKit {
	return tk.SubSeconds(1)
}

// AddWeeks add weeks from current time
func (tk *TimeKit) AddWeeks(w int) *TimeKit {
	tk.SetTime(tk.AddDate(0, 0, daysPerWeek*w))
	return tk
}

// AddWeek add a week from current time
func (tk *TimeKit) AddWeek() *TimeKit {
	return tk.AddWeeks(1)
}

// SubWeeks remove weeks from current time
func (tk *TimeKit) SubWeeks(w int) *TimeKit {
	return tk.AddWeeks(-w)
}

// SubWeek remove a week from current time
func (tk *TimeKit) SubWeek() *TimeKit {
	return tk.SubWeeks(1)
}

// AddWeekdays add weekdays from current time
func (tk *TimeKit) AddWeekdays(wd int) *TimeKit {
	step := 1
	if wd < 0 {
		wd, step = -wd, -step
	}

	for wd > 0 {
		tk.AddDays(step)
		if tk.IsWeekday() {
			wd--
		}
	}
	return tk
}

// AddWeekday add a weekday fro current time
func (tk *TimeKit) AddWeekday() *TimeKit {
	return tk.AddWeekdays(1)
}

// SubWeekdays remove weekdays from current time
func (tk *TimeKit) SubWeekdays(wd int) *TimeKit {
	return tk.AddWeekdays(-wd)
}

// SubWeekday remove a weekday fro current time
func (tk *TimeKit) SubWeekday() *TimeKit {
	return tk.SubWeekdays(1)
}

// AddYears add years from current time
func (tk *TimeKit) AddYears(y int) *TimeKit {
	tk.SetTime(tk.AddDate(y, 0, 0))
	return tk
}

// AddYear add year from current time
func (tk *TimeKit) AddYear() *TimeKit {
	return tk.AddYears(1)
}

// SubYears remove years from current time
func (tk *TimeKit) SubYears(y int) *TimeKit {
	return tk.AddYears(-y)
}

// SubYear remove year from current time
func (tk *TimeKit) SubYear() *TimeKit {
	return tk.SubYears(1)
}

// DiffInSeconds return the difference in seconds
func (tk *TimeKit) DiffInSeconds(t *TimeKit, abs bool) int64 {
	if t == nil {
		t = createFromTimestamp(time.Now().Unix(), tk.Location())
	}
	diff := t.Timestamp() - tk.Timestamp()
	return absoluteValue(abs, diff)
}

// DiffInMinutes return the difference in minutes
func (tk *TimeKit) DiffInMinutes(t *TimeKit, abs bool) int64 {
	return tk.DiffInSeconds(t, abs) / secondsPerMinute
}

// DiffInHours return the difference in hours
func (tk *TimeKit) DiffInHours(t *TimeKit, abs bool) int64 {
	return tk.DiffInMinutes(t, abs) / minutesPerHour
}

// DiffInDays return the difference in days
func (tk *TimeKit) DiffInDays(t *TimeKit, abs bool) int64 {
	return tk.DiffInHours(t, abs) / hoursPerDay
}

// DiffInMonths return the difference in months
func (tk *TimeKit) DiffInMonths(t *TimeKit, abs bool) int64 {
	if t == nil {
		t = createFromTimestamp(time.Now().Unix(), tk.Location())
	}
	tkCopy := tk.Copy()
	tCopy := t.Copy()
	if tkCopy.Location() != tCopy.Location() {
		tkCopy.Time = tkCopy.In(time.UTC)
		tCopy.Time = tCopy.In(time.UTC)
	}
	return countMonthDifference(tkCopy.Time, tCopy.Time, abs)
}

// DiffDurationInString return the duration in string
func (tk *TimeKit) DiffDurationInString(t *TimeKit) string {
	if t == nil {
		t = createFromTimestamp(time.Now().Unix(), tk.Location())
	}
	return strings.Replace(tk.Sub(t.Time).String(), "-", "", 1)
}

func countMonthDifference(t1, t2 time.Time, abs bool) int64 {
	y1 := t1.Year()
	y2 := t2.Year()
	m1 := int(t1.Month())
	m2 := int(t2.Month())
	d1 := t1.Day()
	d2 := t2.Day()

	yearInterval := y1 - y2

	if m1 < m2 || m1 == m2 && d1 < d2 {
		yearInterval--
	}

	monthInterval := (m1 + 12) - m2
	if d1 < d2 {
		monthInterval--
	}
	monthInterval %= 12

	return absoluteValue(abs, int64(yearInterval*12+monthInterval))
}

type Filter func(*TimeKit) bool

// swap swap params
func swap(a, b *TimeKit) (*TimeKit, *TimeKit) {
	return b, a
}

// DiffFiltered return the difference by Unit `duration` and using a filter
//
func (tk *TimeKit) DiffFiltered(t *TimeKit, duration time.Duration, f Filter, abs bool) int64 {
	if t == nil {
		t = createFromTimestamp(time.Now().Unix(), tk.Location())
	}
	var diffNumber int64

	step := int64(duration.Seconds())
	start, end := tk.Copy(), t.Copy()
	inverse := false

	if start.After(end.Time) {
		start, end = swap(start, end)
		inverse = true
	}

	for start.DiffInSeconds(end, true)/step > 0 {
		if f(end) {
			diffNumber++
		}
		end.SetTime(end.Add(-duration))
	}

	if inverse {
		diffNumber = -diffNumber
	}

	return absoluteValue(abs, diffNumber)
}

// DiffInDaysFiltered return difference in days using filter
func (tk *TimeKit) DiffInDaysFiltered(t *TimeKit, f Filter, abs bool) int64 {
	return tk.DiffFiltered(t, time.Hour*hoursPerDay, f, abs)
}

// DiffInHoursFiltered return difference in hours using filter
func (tk *TimeKit) DiffInHoursFiltered(t *TimeKit, f Filter, abs bool) int64 {
	return tk.DiffFiltered(t, time.Hour, f, abs)
}

// StartOfCentury return the datetime of start of the century
func (tk *TimeKit) StartOfCentury() *TimeKit {
	year := tk.Year() - tk.Year()%yearsPerCenturies
	t := time.Date(year, time.January, 1, 0, 0, 0, 0, tk.Location())
	tk.SetTime(t)
	return tk
}

// EndOfCentury return the datetime of end of the century
func (tk *TimeKit) EndOfCentury() *TimeKit {
	year := tk.Year() - 1 - tk.Year()%yearsPerCenturies + yearsPerCenturies
	t := time.Date(year, time.December, 31, 23, 59, 59, 0, tk.Location())
	tk.SetTime(t)
	return tk
}

// StartOfYear return datetime of start of the year
func (tk *TimeKit) StartOfYear() *TimeKit {
	tk.SetTime(time.Date(tk.Year(), time.January, 1, 0, 0, 0, 0, tk.Location()))
	return tk
}

// EndOfYear return datetime of start of the year
func (tk *TimeKit) EndOfYear() *TimeKit {
	tk.SetTime(time.Date(tk.Year(), time.December, 31, 23, 59, 59, 0, tk.Location()))
	return tk
}

// StartOfQuarter return datetime of start of the quarter
func (tk *TimeKit) StartOfQuarter() *TimeKit {
	m := time.Month((tk.Quarter()-1)*monthsPerQuarter + 1)
	tk.SetTime(time.Date(tk.Year(), m, 1, 0, 0, 0, 0, tk.Location()))
	return tk
}

// EndOfQuarter return datetime of end of the quarter
func (tk *TimeKit) EndOfQuarter() *TimeKit {
	m := tk.Quarter() * monthsPerQuarter
	tk.SetTime(time.Date(tk.Year(), time.Month(m), 1, 23, 59, 59, 0, tk.Location()).AddDate(0, 1, -1))
	return tk
}

// StartOfMonth return datetime of start of the month
func (tk *TimeKit) StartOfMonth() *TimeKit {
	tk.SetTime(time.Date(tk.Year(), tk.Month(), 1, 0, 0, 0, 0, tk.Location()))
	return tk
}

// EndOfMonth return datetime of end of the month
func (tk *TimeKit) EndOfMonth() *TimeKit {
	tk.SetTime(time.Date(tk.Year(), tk.Month(), 1, 0, 0, 0, 0, tk.Location()).AddDate(0, 1, -1))
	return tk
}

// StartOfWeek return datetime of start of the Week
func (tk *TimeKit) StartOfWeek() *TimeKit {
	t := time.Date(tk.Year(), tk.Month(), tk.Day(), 0, 0, 0, 0, tk.Location())
	for t.Weekday() != tk.weekStartAt {
		t = t.AddDate(0, 0, -1)
	}
	tk.SetTime(t)
	return tk
}

// EndOfWeek return datetime of end of the Week
func (tk *TimeKit) EndOfWeek() *TimeKit {
	t := time.Date(tk.Year(), tk.Month(), tk.Day(), 23, 59, 59, 0, tk.Location())
	for t.Weekday() != tk.weekEndAt {
		t = t.AddDate(0, 0, 1)
	}
	tk.SetTime(t)
	return tk
}

// StartOfDay return datetime of start of the Day
func (tk *TimeKit) StartOfDay() *TimeKit {
	t := time.Date(tk.Year(), tk.Month(), tk.Day(), 0, 0, 0, 0, tk.Location())
	tk.SetTime(t)
	return tk
}

// EndOfDay return datetime of end of the Day
func (tk *TimeKit) EndOfDay() *TimeKit {
	t := time.Date(tk.Year(), tk.Month(), tk.Day(), 23, 59, 59, 0, tk.Location())
	tk.SetTime(t)
	return tk
}

// LastDayOfPreMonth return the last day of the previous month
func (tk *TimeKit) LastDayOfPreMonth() time.Time {
	return tk.AddDate(0, 0, -tk.Day())
}

// Timestamp returns the current time's seconds since Jan 1 1970 (Unix time).
func (tk *TimeKit) Timestamp() int64 {
	return tk.Unix()
}

// Copy return a new TimeKit instance of current instance
func (tk *TimeKit) Copy() *TimeKit {
	return createFromTimestamp(tk.Time.Unix(), tk.Location())
}

// IsWeekend whether the current time is a weekend day
func (tk *TimeKit) IsWeekend() bool {
	d := tk.Weekday()

	for _, v := range tk.weekendDays {
		if d == v {
			return true
		}
	}
	return false
}

// IsWeekday whether the current time is a weekday
func (tk *TimeKit) IsWeekday() bool {
	return !tk.IsWeekend()
}

// WeekendDays return the weekend days of the week
func (tk *TimeKit) WeekendDays() []time.Weekday {
	return tk.weekendDays
}

func (tk *TimeKit) Quarter() int {
	m := tk.Month()
	switch {
	case m < 4:
		return 1
	case m >= 4 && m < 7:
		return 2
	case m >= 7 && m < 10:
		return 3
	}
	return 4
}

// String gets the time string using the previously set format
func (tk *TimeKit) String() string {
	return tk.Format(tk.format)
}

// DateTimeString get the date string
func (tk *TimeKit) DateTimeString() string {
	return tk.Format(DefaultFormat)
}

// DateString get the date string
func (tk *TimeKit) DateString() string {
	return tk.Format(DateFormat)
}

// TimeString get the time string
func (tk *TimeKit) TimeString() string {
	return tk.Format(TimeFormat)
}

// absolute return the abs value if need
func absoluteValue(abs bool, value int64) int64 {
	if abs && value < 0 {
		return -value
	}
	return value
}
