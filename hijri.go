// Hijri Calendar v0.4.4

// Copyright (c) 2016-2019 K.I.A.Derouiche

// This source code is licensed under Apache-2.0 license that can be found in the LICENSE file
// Package hjiri provides functionality for implementation of Islamic(Hijri) Calendar.

package hijri

import (
	"math"
	"time"

	"go-hijricalendar/ummalquradb"
)

// A Month specifies a month of the year starting from Mouharram = 1.
type Month int

// A Weekday specifies a day of the week starting.
type Weekday int

// A Time represents a moment in time in Hijri Calendar.
type Time struct {
	year  int
	month Month
	day   int
	hour  int
	min   int
	sec   int
	nsec  int
	loc   *time.Location
	wday  Weekday
}

// List of months in Hijri calendar. Starting from Mouharram = 1
const (
	Mouharram Month = 1 + iota
	Safar
	RabiaAlAwal
	RabiaAthThani
	JoumadaAlOula
	JoumadaAthThania
	Rajab
	Chaabane
	Ramadan
	Chawwal
	DhouAlQida
	DhouAlHijja
)

// List of arabic days in a week.
const (
	Alsabt Weekday = iota
	Alahad
	Alithnayn
	Altholathae
	Alalrbiae
	Alhamiss
	Aljomoaa
)

var months = [...]string{
	"محرم",         //"Mouharram"
	"صفر",          //"Safar"
	"ربيع الأول",   //"RabiaAlAwal"
	"ربيع الثاني",  //"RabiaAthThani"
	"جمادي الأولى", //"JoumadaAlOula"
	"جمادي الآخرة", //"JoumadaAlhThania"
	"رجب",          //"Rajab"
	"شعبان",        //"Chaabane"
	"رمضان",        //"Ramadan"
	"شوال",         //"Chawwal"
	"ذو القعدة",    //"DhouAlQida"
	"ذو الحجة",     //"DhouAlHijja"
}

var days = [...]string{
	"السبت",    //"Alsabt"
	"الاحد",    //"Alahad"
	"الاثنين",  //"Alithnayn"
	"الثلاثاء", //"Altholathae"
	"الاربعاء", //"Alalrbiae"
	"الخميس",   //"Alhamiss"
	"الجمعة",   //"Aljomoaa"
}

var sdays = [...]string{
	"S",
	"A",
	"I",
	"T",
	"A",
	"K",
	"J",
}

// {days, leap_days, days_before_start}
var CalendarCount = [12][3]int{
	{30, 30, 0},   // Mouharram
	{29, 29, 30},  // Safar
	{30, 30, 60},  // RabiaAlAwal
	{29, 29, 90},  // RabiaAthThani
	{30, 30, 120}, // JoumadaAlOula
	{29, 29, 150}, // JoumadaAlhThania
	{30, 30, 180}, // Rajab
	{29, 29, 210}, // Chaabane
	{29, 30, 240}, // Ramadan(exception)
	{29, 29, 270}, // Chawwal
	{30, 30, 300}, // DhouAlQida
	{29, 30, 330}, // DhouAlHijja (exception)
}

// Returns the Hijri name of the month.
func (m Month) String() string {
	return months[m-1]
}

// Pointers to time.Location for UmmAlQura time zones.
func UmmAlQura() *time.Location {
	loc, err := time.LoadLocation("Asia/UmmAlQura")
	if err != nil {
		loc = time.FixedZone("Asia/UmmAlQura", 10800) // UTC + 03:00
	}
	return loc
}

func getWeekday(wd time.Weekday) Weekday {
	switch wd {
	case time.Saturday:
		return Alsabt
	case time.Sunday:
		return Alahad
	case time.Monday:
		return Alithnayn
	case time.Tuesday:
		return Altholathae
	case time.Wednesday:
		return Alalrbiae
	case time.Thursday:
		return Alhamiss
	case time.Friday:
		return Aljomoaa
	}
	return 0
}

// Date returns the year, month, day of t.
func (t Time) Date() (int, Month, int) {
	return t.year, t.month, t.day
}

// String returns the Hijri name of the day in week
func (wd Weekday) String() string {
	return days[wd]
}

// Short returns the Hijri short name of the day in week
func (wd Weekday) Short() string {
	return sdays[wd]
}

// Unix returns the number of seconds since January 1, 1979 UTC.
func (t Time) Unix() int64 {
	return t.HijriTime().Unix()
}

// UnixNano seturns the number of nanoseconds since January 1, 1970 UTC.
func (t Time) UnixNano() int64 {
	return t.HijriTime().UnixNano()
}

// Location returns a pointer to time.Location of t.
func (t Time) Location() *time.Location {
	return t.loc
}

// Clock returns the hour, minute, and second within the day specified by t.
func (t Time) Clock() (hour, min, sec int) {
	return t.hour, t.min, t.sec
}

// Year returns the year of t
func (t Time) Year() int {
	return t.year
}

// Month returns the month of t
func (t Time) Month() Month {
	return t.month
}

// Day returns the Days of t
func (t Time) Day() int {
	return t.day
}

// Hour returns the hour of t in the range [0, 23].
func (t Time) Hour() int {
	return t.hour
}

// Minute returns the minute offset of t in the range [0, 59].
func (t Time) Minute() int {
	return t.min
}

// Second returns the seconds offset of t in the range [0, 59].
func (t Time) Second() int {
	return t.sec
}

// Nanosecond returns the nanoseconds offset of t in the range [0, 999999999].
func (t Time) Nanosecond() int {
	return t.nsec
}

//Weekday returns the weekday of t
func (t Time) Weekday() Weekday {
	return t.wday
}

// Add returns a new instance of Time for t+d.
func (t Time) Add(d time.Duration) Time {
	return New(t.HijriTime().Add(d))
}

// AddDate returns a new instance of Time for
func (t Time) AddDate(years, months, days int) Time {
	return New(t.HijriTime().AddDate(years, months, days))
}

// Yesterday returns a new instance of Time representing a day before the day of t
func (t Time) Yesterday() Time {
	return t.AddDate(0, 0, -1)
}

// Tomorrow returns a new instance of Time representing a day after the day of t.
func (t Time) Tomorrow() Time {
	return t.AddDate(0, 0, 1)
}

// Now returns a new instance of Time corresponding to the current time.
// loc is a pointer to time.Location and must not be nil.
func Now(loc *time.Location) Time {
	if loc == nil {
		panic("Hijri: the Location must not be nil in call to Now")
	}
	return New(time.Now().In(loc))
}

// Zone returns the zone name and its offset in seconds east of UTC of t.
func (t Time) Zone() (string, int) {
	return t.HijriTime().Zone()
}

// New converts Gregorian calendar to Hijri calendar and
// returns a new instance of Time corresponding to the time of t.
// t is an instance of time.Time in Gregorian calendar.
func New(t time.Time) Time {
	hit := new(Time)
	hit.Kcalendar(t)

	return *hit
}

// http://www.coderanch.com/t/534271/java/java/Gregorian-Hijri-Dates-Converter-JAVA
//Gegorean To Hijri

func (t *Time) Kcalendar(tx time.Time) {

	var year, month, day int
	var epochastro float64 = 1948084

	t.nsec = tx.Nanosecond()
	t.sec = tx.Second()
	t.min = tx.Minute()
	t.hour = tx.Hour()
	t.loc = tx.Location()
	t.wday = getWeekday(tx.Weekday())

	gy, gm, gday := tx.Date()
	vgm := int(gm)

	if vgm < 3 {
		gy--
		vgm += 12
	}

	a := math.Floor(float64(gy) / 100.)
	b := 2 - a + math.Floor(a/4.)

	if gy < 1583 {
		b = 0
	}
	if gy == 1582 {
		if vgm > 10 {
			b = -10
		}
		if vgm == 10 {
			b = 0
			if gday > 4 {
				b = -10
			}
		}
	}

	jd := math.Floor(365.25*(float64(gy)+4716)) + math.Floor(30.6001*(float64(vgm)+1)) + float64(gday) + b - 1524

	b = 0
	if jd > 2299160 {
		a = math.Floor((jd - 1867216.25) / 36524.25)
		b = 1 + a - math.Floor(a/4.)
	}
	bb := jd + b + 1524
	cc := math.Floor((bb - 122.1) / 365.25)
	dd := math.Floor(365.25 * cc)
	ee := math.Floor((bb - dd) / 30.6001)
	gday = int((bb - dd) - math.Floor(30.6001*ee))
	month = int(ee - 1)
	if ee > 13 {
		cc++
		month = int(ee - 13)
	}
	year = int(cc - 4716)

	iyear := 10631. / 30.

	shift1 := 8.01 / 60.

	z := jd - epochastro
	cyc := math.Floor(z / 10631.)
	z = z - 10631*cyc
	j := math.Floor((z - shift1) / iyear)
	year = int(30*cyc + j)
	z = z - math.Floor(j*iyear+shift1)
	month = int(math.Floor((z + 28.5001) / 29.5))
	if month == 13 {
		month = 12
	}
	day = int(z - math.Floor(29.5001*float64(month)-29))

	t.year = year
	t.month = Month(month)
	t.day = day
}

// Unix returns a new instance of HijriDate from unix timestamp.
// sec seconds and nsec nanoseconds since January 1, 1970 UTC.
// loc is a pointer to time.Location and must not be nil.
func Unix(sec, nsec int64, loc *time.Location) Time {
	if loc == nil {
		panic("hijri: the location must not be nil in call to Unix")
	}
	return New(time.Unix(sec, nsec).In(loc))
}

// SetUnix sets t to represent the corresponding unix timestamp of
// sec seconds and nsec nanoseconds since January 1, 1970 UTC.
// loc is a pointer to time.Location and must not be nil.
func (t *Time) SetUnix(sec, nsec int64, loc *time.Location) Time {
	if loc == nil {
		panic("hijri: the location must not be nil in call to SetUnix")
	}
	return New(time.Unix(sec, nsec).In(loc))
}

// Hijri To Gregorian
func Date(year int, month Month, day, hour, min, sec, nsec int, loc *time.Location) Time {
	if loc == nil {
		panic("hijri: the Location must not be nil in call to Date")
	}

	t := new(Time)
	t.DateChange(year, month, day, hour, min, sec, nsec, loc)

	return *t
}

func (t *Time) DateChange(year int, month Month, day, hour, min, sec, nsec int, loc *time.Location) {
	if loc == nil {
		panic("hijri: the Location must not be nil in call to Change")
	}

	t.year = year
	t.month = month
	t.day = day
	t.hour = hour
	t.min = min
	t.sec = sec
	t.nsec = nsec
	t.loc = loc
}

func GetJdnHijri(year, month, day int) int {
	iy := year
	im := month
	id := day
	ii := iy - 1
	iln := (ii * 12) + 1 + (im - 1)
	i := iln - 16260
	mcjdn := id + ummalquradb.GetUmmalquradb(i)
	cjdn := mcjdn + 2400000 //int64

	return cjdn
}

//source from: http://keith-wood.name/calendars.html
func (t Time) HijriTime() time.Time {
	var e, a, b, c, d, z float64
	var day, year, month float64

	julianDate := GetJdnHijri(t.year, int(t.month), t.day)

	z = math.Floor(float64(julianDate) + 0.5)
	a = math.Floor((z - 1867216.25) / 36524.25)
	a = z + 1 + a - math.Floor(a/4)
	b = a + 1524
	c = math.Floor((b - 122.1) / 365.25)
	d = math.Floor(365.25 * c)
	e = math.Floor((b - d) / 30.6001)
	day = b - d - math.Floor(e*30.6001)
	if e > 13.5 {
		month = e - 13.0
	} else {
		month = e - 1.0
	}
	if month > 2.5 {
		year = c - 4716.0
	} else {
		year = c - 4715.0
	}
	if year <= 0.0 {
		year -= 1.0
	}
	return time.Date(int(year), time.Month(int(month)), int(day), t.hour, t.min, t.sec, t.nsec, t.loc)
}

// End Hijri to Gregorian
