// This file is auto-generated, DO NOT EDIT.
//
// Source:
//     Title: Hub Service
//     Version: 0.1.0
package generatortest

import (
	validation "github.com/go-ozzo/ozzo-validation"
)

// Interval is an enum.
type Interval string

var (
	IntervalDay     Interval = "Day"
	IntervalHour    Interval = "Hour"
	IntervalMonth   Interval = "Month"
	IntervalQuarter Interval = "Quarter"
	IntervalWeek    Interval = "Week"
	IntervalYear    Interval = "Year"

	// KnownInterval is the list of valid Interval
	KnownInterval = []Interval{
		IntervalDay,
		IntervalHour,
		IntervalMonth,
		IntervalQuarter,
		IntervalWeek,
		IntervalYear,
	}
	// KnownIntervalString is the list of valid Interval as string
	KnownIntervalString = []string{
		string(IntervalDay),
		string(IntervalHour),
		string(IntervalMonth),
		string(IntervalQuarter),
		string(IntervalWeek),
		string(IntervalYear),
	}

	// InKnownInterval is an ozzo-validator for Interval
	InKnownInterval = validation.In(
		IntervalDay,
		IntervalHour,
		IntervalMonth,
		IntervalQuarter,
		IntervalWeek,
		IntervalYear,
	)
)
