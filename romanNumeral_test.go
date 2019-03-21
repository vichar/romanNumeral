package romanNumeral_test

import (
	"testing"

	"github.com/vichar/romanNumeral"
)

func Test(t *testing.T) {
	t.Run("test return the number as a string for number = 1 should return I", func(t *testing.T) {
		get := romanNumeral.Parse(1)
		want := "I"
		if want != get {
			t.Errorf("It should say %s but get %s", want, get)
		}
	})
	t.Run("test return the number as a string for number = 3 should return III", func(t *testing.T) {
		get := romanNumeral.Parse(3)
		want := "III"
		if want != get {
			t.Errorf("It should say %s but get %s", want, get)
		}
	})
	t.Run("test return the number as a string for number = 4 should return IV", func(t *testing.T) {
		get := romanNumeral.Parse(4)
		want := "IV"
		if want != get {
			t.Errorf("It should say %s but get %s", want, get)
		}
	})
	t.Run("Test return the number as a string for number = 5 should return V", func(t *testing.T) {
		get := romanNumeral.Parse(5)
		want := "V"
		if want != get {
			t.Errorf("It should say %s but get %s", want, get)
		}
	})
	t.Run("Test return the number as a string for number = 6 should return VI", func(t *testing.T) {
		get := romanNumeral.Parse(6)
		want := "VI"
		if want != get {
			t.Errorf("It should say %s but get %s", want, get)
		}
	})
	t.Run("Test return the number as a string for number = 7 should return VII", func(t *testing.T) {
		get := romanNumeral.Parse(7)
		want := "VII"
		if want != get {
			t.Errorf("It should say %s but get %s", want, get)
		}
	})
	t.Run("Test return the number as a string for number = 9 should return IX", func(t *testing.T) {
		get := romanNumeral.Parse(9)
		want := "IX"
		if want != get {
			t.Errorf("It should say %s but get %s", want, get)
		}
	})

	t.Run("Test return the number as a string for number = 10 should return X", func(t *testing.T) {
		get := romanNumeral.Parse(10)
		want := "X"
		if want != get {
			t.Errorf("It should say %s but get %s", want, get)
		}
	})
	t.Run("Test return the number as a string for number = 13 should return XIII", func(t *testing.T) {
		get := romanNumeral.Parse(13)
		want := "XIII"
		if want != get {
			t.Errorf("It should say %s but get %s", want, get)
		}
	})
	t.Run("Test return the number as a string for number = 20 should return XX", func(t *testing.T) {
		get := romanNumeral.Parse(19)
		want := "XX"
		if want != get {
			t.Errorf("It should say %s but get %s", want, get)
		}
	})

}
