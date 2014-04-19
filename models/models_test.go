package models

import (
    "testing"
    // "fmt"
   )

func Test_AddLanguage(t *testing.T) {
	tech := Technology{}
	golang := Language{"Go"}

	result := tech.AddLanguage(golang) 
	if !result {
		t.Error("The language should have been added")
		return
	}

	if result && len(tech.Languages) == 0 {
		t.Error("The value has been added but is not in the array")
		return
	}

	if tech.Languages[0] != golang {
		t.Error("Wrong language added")
		return
	}

	t.Log("Language added")

}

func Test_RemoveLanguage(t *testing.T) {
	tech := Technology{}
	golang := Language{"Go"}
	python := Language{"Python"}

	tech.AddLanguage(golang) 
	tech.AddLanguage(python) 

	tech.RemoveLanguage(golang)

	if len(tech.Languages) != 1 {
		t.Error("The language has not been removed")
		return
	}

	if tech.Languages[0] != python {
		t.Error("Wrong language removed")
	}

	t.Log("Language removed")
}