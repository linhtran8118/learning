package main

import (
	"reflect"
	"testing"
)

func TestFindShortestPath(t *testing.T) {
	cases := []struct {
		name     string
		source   Location
		dest     Location
		method   string
		expected []Location
	}{
		{"Same location", Location{0, 0}, Location{0, 0}, "aaaa",[]Location{}},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			actual := findShortestPath(c.source, c.dest, c.method)
			if !reflect.DeepEqual(actual, c.expected) {
				t.Error("Wrong locations", actual)
			}
		})
	}
}

func TestNewShortestPath(t *testing.T) {
	m := &Map{}
	s := Location{0, 0}
	d := Location{1, 1}
	cases := []struct{
		name string
		strategy TravelStrategy
	}{
		{"car", Car{}},
		{"walking", Walking{}},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			m.setTravelStrategy(tc.strategy)
			actual := m.findShortestPath(s, d)
			if !reflect.DeepEqual(actual, []Location{}) {
				t.Error("Check wrong:", actual)
			}
		})
	}
}
