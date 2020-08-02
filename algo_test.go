package main

import (
	"testing"
)

func TestBase(t *testing.T) {
	var algo = BaseTypeAlgo

	var tTests = []struct {
		input    In
		expected Out
	}{
		{
			In{A: false, B: false, C: false, D: 0, E: 0, F: 0},
			Out{},
		},
		{
			In{A: false, B: false, C: true, D: 0, E: 0, F: 0},
			Out{},
		},
		{
			In{A: false, B: true, C: false, D: 0, E: 0, F: 0},
			Out{},
		},
		{
			In{A: false, B: true, C: true, D: 10, E: 50, F: 150},
			Out{H: "T", K: -40},
		},
		{
			In{A: true, B: false, C: false, D: 0, E: 0, F: 0},
			Out{},
		},
		{
			In{A: true, B: false, C: true, D: 0, E: 0, F: 0},
			Out{},
		},
		{
			In{A: true, B: true, C: false, D: 5.0, E: 10, F: 20},
			Out{H: "M", K: 10},
		},
		{
			In{A: true, B: true, C: true, D: 55.5, E: 210, F: 40},
			Out{H: "P", K: 425.5},
		},
	}

	for i, tt := range tTests {
		actual, err := tt.input.Process(algo)
		if actual != tt.expected {
			t.Errorf("Type='%s' (%d): input %+v expected %+v, actual %+v", algo, i, tt.input, tt.expected, actual)
		}

		if err == nil {
			if !(actual.H == "M" || actual.H == "P" || actual.H == "T") {
				t.Errorf("Type='%s' (%d): must throw error as H does not equal to either M,P,T", algo, i)
			}
		}
	}
}

func TestCustom1(t *testing.T) {
	var algo = Custom1TypeAlgo

	var tTests = []struct {
		input    In
		expected Out
	}{
		{
			In{A: false, B: false, C: false, D: 0, E: 0, F: 0},
			Out{},
		},
		{
			In{A: false, B: false, C: true, D: 0, E: 0, F: 0},
			Out{},
		},
		{
			In{A: false, B: true, C: false, D: 0, E: 0, F: 0},
			Out{},
		},
		{
			In{A: false, B: true, C: true, D: 10, E: 50, F: 150},
			Out{H: "T", K: -40},
		},
		{
			In{A: true, B: false, C: false, D: 0, E: 0, F: 0},
			Out{},
		},
		{
			In{A: true, B: false, C: true, D: 0, E: 0, F: 0},
			Out{},
		},
		{
			In{A: true, B: true, C: false, D: 5.0, E: 10, F: 20},
			Out{H: "M", K: 10},
		},
		{
			In{A: true, B: true, C: true, D: 55.5, E: 210, F: 40},
			Out{H: "P", K: 227.55},
		},
	}

	for i, tt := range tTests {
		actual, err := tt.input.Process(algo)
		if actual != tt.expected {
			t.Errorf("Type='%s' (%d): input %+v expected %+v, actual %+v", algo, i, tt.input, tt.expected, actual)
		}

		if err == nil {
			if !(actual.H == "M" || actual.H == "P" || actual.H == "T") {
				t.Errorf("Type='%s' (%d): must throw error as H does not equal to either M,P,T", algo, i)
			}
		}
	}
}

func TestCustom2(t *testing.T) {
	var algo = Custom2TypeAlgo

	var tTests = []struct {
		input    In
		expected Out
	}{
		{
			In{A: false, B: false, C: false, D: 0, E: 0, F: 0},
			Out{},
		},
		{
			In{A: false, B: false, C: true, D: 0, E: 0, F: 0},
			Out{},
		},
		{
			In{A: false, B: true, C: false, D: 0, E: 0, F: 0},
			Out{},
		},
		{
			In{A: false, B: true, C: true, D: 10, E: 50, F: 150},
			Out{H: "T", K: -40},
		},
		{
			In{A: true, B: false, C: false, D: 0, E: 0, F: 0},
			Out{},
		},
		{
			In{A: true, B: false, C: true, D: 50.5, E: 120, F: 60},
			Out{H: "M", K: 171.1},
		},
		{
			In{A: true, B: true, C: false, D: 5.25, E: 90, F: 10},
			Out{H: "T", K: 3.5},
		},
		{
			In{A: true, B: true, C: true, D: 55.5, E: 210, F: 40},
			Out{H: "P", K: 425.5},
		},
	}

	for i, tt := range tTests {
		actual, err := tt.input.Process(algo)
		if actual != tt.expected {
			t.Errorf("Type='%s' (%d): input %+v expected %+v, actual %+v", algo, i, tt.input, tt.expected, actual)
		}

		if err == nil {
			if !(actual.H == "M" || actual.H == "P" || actual.H == "T") {
				t.Errorf("Type='%s' (%d): must throw error as H does not equal to either M,P,T", algo, i)
			}
		}
	}
}
