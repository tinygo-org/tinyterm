package tinyterm

import (
	"strconv"
	"testing"
)

func Test_parseUint8(t *testing.T) {
	testCases := []struct {
		str string
	}{
		{str: "64"},
		{str: "255"},
		{str: "8"},
	}
	for _, tc := range testCases {
		t.Run(tc.str, func(t *testing.T) {
			expected, xerr := strconv.Atoi(tc.str)
			parsed, err := parseUint8([]byte(tc.str))
			if (xerr != nil && err == nil) || (err != nil && xerr == nil) {
				t.Errorf("xerr: %v, err: %v", xerr, err)
			}
			if expected != int(parsed) {
				t.Errorf("expected: %v, parsed: %v", expected, parsed)
			}
		})
	}
}

func Test_parseSGR(t *testing.T) {
	testCases := []struct {
		desc               string
		str                string
		count              int
		v1, v2, v3, v4, v5 uint8
	}{
		{
			desc:  "fiveValues1",
			str:   "38;2;255;64;8",
			count: 5, v1: 38, v2: 2, v3: 255, v4: 64, v5: 8,
		},
		{
			desc:  "threeValues1",
			str:   "38;5;128",
			count: 3, v1: 38, v2: 5, v3: 128,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			pb := buffer{}
			for _, c := range []byte(tc.str) {
				pb.WriteByte(c)
			}
			v, count := pb.parseSGR()
			v1 := v[0]
			v2 := v[1]
			v3 := v[2]
			v4 := v[3]
			v5 := v[4]
			if v1 != tc.v1 || v2 != tc.v2 || v3 != tc.v3 ||
				v4 != tc.v4 || v5 != tc.v5 || count != tc.count {
				t.Errorf(
					"expected: %v, %v, %v, %v, %v, %d; actual: %v, %v, %v, %v, %v, %d",
					tc.v1, tc.v2, tc.v3, tc.v4, tc.v5, tc.count, v1, v2, v3, v4, v5, count,
				)
			}
		})
	}
}
