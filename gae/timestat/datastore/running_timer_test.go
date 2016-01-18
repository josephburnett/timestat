package datastore

import (
	"testing"
	"time"
	m "timestat/model"
)

var sampleTimeString = "Mon Jan 2 15:04:05 -0700 MST 2006"
var sampleTime, _ = time.Parse(sampleTimeString, sampleTimeString)
var sampleTimer = &m.RunningTimer{Start: sampleTime}

func TestDimensionIDGeneral(t *testing.T) {
	id, err := DimensionID(sampleTimer, m.General)
	assertNoError(t, err)
	assertID(t, "general", id)
}

func TestDimensionIDDay(t *testing.T) {
	id, err := DimensionID(sampleTimer, m.Day)
	assertNoError(t, err)
	assertID(t, "2006-01-02", id)
}

func TestDimensionIDWeek(t *testing.T) {
	id, err := DimensionID(sampleTimer, m.Week)
	assertNoError(t, err)
	assertID(t, "2006-W01", id)
}

func TestDimensionIDMonth(t *testing.T) {
	id, err := DimensionID(sampleTimer, m.Month)
	assertNoError(t, err)
	assertID(t, "january", id)
}

func TestDimensionIDYear(t *testing.T) {
	id, err := DimensionID(sampleTimer, m.Year)
	assertNoError(t, err)
	assertID(t, "2006", id)
}

func TestDimensionIDDayOfTheWeek(t *testing.T) {
	id, err := DimensionID(sampleTimer, m.DayOfTheWeek)
	assertNoError(t, err)
	assertID(t, "monday", id)
}

func TestDimensionIDTenMinuteTime(t *testing.T) {
	id, err := DimensionID(sampleTimer, m.TenMinuteTime)
	assertNoError(t, err)
	assertID(t, "15-00", id)
}

func TestDimensionIDTenMinuteTimeAndDayOfTheWeek(t *testing.T) {
	id, err := DimensionID(sampleTimer, m.TenMinuteTimeAndDayOfTheWeek)
	assertNoError(t, err)
	assertID(t, "monday-15-00", id)
}

func TestDimensionIDInvalid(t *testing.T) {
	_, err := DimensionID(sampleTimer, m.Dimension("invalid"))
	if err == nil {
		t.Errorf("Expected error but found none.")
	}
}

func assertNoError(t *testing.T, err error) {
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
}

func assertID(t *testing.T, expect, actual string) {
	if expect != actual {
		t.Errorf("Expected %v but got: %v", expect, actual)
	}
}
