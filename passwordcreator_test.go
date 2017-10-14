/*
Copyright (C) 2011 Andreas Sinz
Copyright (C) 2013 Adam Jimerson
Copyright (C) 2017 Luther Thompson

This program is free software; you can redistribute it and/or
modify it under the terms of the GNU General Public License, version 2,
as published by the Free Software Foundation.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with this program; if not, write to the Free Software
Foundation, Inc., 51 Franklin Street, Fifth Floor, Boston, MA  02110-1301, USA.
*/

package main

import (
	"testing"
)

const testLength = 8

func TestDigitsCharset(t *testing.T) {
	c, e := NewCreator(testLength, false, false, true, false, "")

	testDigits := "0123456789"
	if e != nil || c.characters != testDigits {
		t.Errorf("Characters not distinct.\nExpected \"%s\", but got \"%s\"", testDigits, c.characters)
	}
}

func TestSomeChars(t *testing.T) {
	c, err := NewCreator(testLength, true, false, true, false, ",.-_")

	testCharacters := "abcdefghijklmnopqrstuvwxyz0123456789,.-_"

	if err != nil || c.characters != testCharacters {
		t.Errorf("Characters not distinct.\nExpected \"%s\", but got \"%s\"", testCharacters, c.characters)
	}
}

func TestUniqueChars(t *testing.T) {
	expected := "ab"
	if c, err := NewCreator(
		testLength, false, false, false, false, "aaabbb",
	); c.characters != expected || err != nil {
		t.Errorf(
			"Characters not distinct.\nExpected \"%s\", but got \"%s\"",
			expected, c.characters,
		)
	}
}
