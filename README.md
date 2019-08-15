# بِسْمِ اللَّهِ الرَّحْمَٰنِ الرَّحِيمِ 

# hijri-calendar
Hijri calendar - التقويم الهجري en Go.

**Go Hijri Calendar v0.4.2** provides functionality for conversion among Hijri  to Gregorian calendars

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
var ht time.Time = time.Date(2018, time.June, 20, 12, 1, 1, 0, hijri.Hidjaz)

// Get a new instance of time.Time
d := hijri.New(ht)

// Get the date in Gregorian calendar
fmt.Println(d.Date())
```

3- Convert Hijri calendar to Gregorian calendar.

```go
// Create a new instance of hijri.Time
var ht hijri.Time = hijri.Date(1436, hijri.JoumadaAlOula, 7, 12, 59, 59, 0, hijri.Hidjaz)

// Get a new instance of time.Time
t := ht.HijriTime()

// Get the date in Gregorian calendar
fmt.Println(t.Date()) // output: 2015 February 26
```
