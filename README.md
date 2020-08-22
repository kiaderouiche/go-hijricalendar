# بِسْمِ اللَّهِ الرَّحْمَٰنِ الرَّحِيمِ

# Go Hijri Calendar
Hijri calendar - التقويم الهجري en Go.

**Go Hijri Calendar v0.4.5** provides functionality for conversion among Hijri  to Gregorian calendars

## Remark
(Inspired by go-persian-calendar objectif unifying syntax and code)

## Installation

$ go get -u github.com/kiaderouiche/go-hijricalendar

## Getting started

1- Import the package `hijri`.

```go
import (
    "fmt"
    "time"
    "github.com/kiaderouiche/go-hijricalendar"
)
```

2- Convert Gregorian calendar to Hijri calendar.

```go
// Create a new instance of hijri.Time
var ht time.Time = time.Date(2018, time.June, 20, 12, 1, 1, 0, hijri.UmmAlQura())

// Get a new instance of time.Time
d := hijri.New(ht)

// Get the date in Gregorian calendar
fmt.Println(d.Date())
```

3- Convert Hijri calendar to Gregorian calendar.

```go
// Create a new instance of hijri.Time
var ht hijri.Time = hijri.Date(1436, hijri.JoumadaAlOula, 7, 12, 59, 59, 0, hijri.UmmAlQura())

// Get a new instance of time.Time
t := ht.HijriTime()

// Get the date in Gregorian calendar
fmt.Println(t.Date()) // output: 2015 February 26
```
4- Get time information.

```go
// Get a new instance of hijri.Time representing the current time
ht := hijri.Now(hijri.UmmAlQura())

// Get year, month, day
fmt.Printf("%v %v %v\n", ht.Year(), ht.Month(), ht.Day())

// Get hour, minute, second
fmt.Printf("Clock: %v\n", ht.Clock())
fmt.Printf("%v %v %v\n", ht.Hour(), ht.Minute(), ht.Second())

// Get Unix timestamp (the number of seconds since January 1, 1970 UTC)
fmt.Printf("Unix time: %v\n", ht.Unix())

// Get yesterday, today and tomorrow
fmt.Printf("Yesterday: %v\n", ht.Yesterday().Weekday())
fmt.Printf("Weekday: %v\n", ht.Weekday())
fmt.Printf("Tomorrow: %v\n", ht.Tomorrow().Weekday())
```
