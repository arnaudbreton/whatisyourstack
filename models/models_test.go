package models

import (
    "testing"
    "github.com/stretchr/testify/assert"
    // "fmt"
   )

func Test_NewTechnology(t *testing.T) {
    tech, err := NewTechnology("Test")

    assert.NotNil(t, tech)
    assert.Nil(t, err)
    assert.Equal(t, "Test", tech.Name)
    assert.NotNil(t, tech.GetLanguages())
    assert.Equal(t, 0, len(tech.GetLanguages()))

    techError, err2 := NewTechnology("")
    assert.Nil(t, techError)
    assert.NotNil(t, err2)
}

func Test_AddLanguage(t *testing.T) {
    tech, _ := NewTechnology("Test")
    golang := Language{"Go"}

    err := tech.AddLanguage(&golang) 
    assert.Nil(t, err)

    languages := tech.GetLanguages()
    assert.Equal(t, 1, len(languages))

    assert.Equal(t, golang, languages[0])

    err = tech.AddLanguage(&Language{}) 
    assert.NotNil(t, err)
}

func Test_RemoveLanguage(t *testing.T) {
    tech, _ := NewTechnology("Test")
    golang := Language{"Go"}
    python := Language{"Python"}
    ruby := Language{"Ruby"}

    tech.AddLanguage(&golang) 
    tech.AddLanguage(&python) 
    tech.AddLanguage(&ruby) 

    tech.RemoveLanguage(&python)

    languages := tech.GetLanguages()
    assert.Equal(t, 2, len(languages))
    assert.Equal(t, golang, languages[0])

    err := tech.RemoveLanguage(&Language{}) 
    assert.NotNil(t, err)
}

func Test_NewStack(t *testing.T) {
    company := Company{"Test company"}
    stack, err := NewStack(&company)

    assert.Nil(t, err)
    assert.NotNil(t, stack)
    assert.NotNil(t, stack.GetTechnologies())
    assert.Equal(t, 0, len(stack.GetTechnologies()))
    assert.Equal(t, company, stack.Company)

    stackError, err2 := NewStack(nil)

    assert.Nil(t, stackError)
    assert.NotNil(t, err2)
}

func Test_AddTechnology(t *testing.T) {
    company := Company{"Test company"}
    stack, _ := NewStack(&company)

    tech, _ := NewTechnology("Martini")
    golang := Language{"Go"}

    tech.AddLanguage(&golang)

    err := stack.AddTechnology(tech)
    assert.Nil(t, err)

    technologies := stack.GetTechnologies() 
    assert.Equal(t, 1, len(technologies))
    assert.Equal(t, tech.Name, technologies[0].Name)
}

func Test_RemoveTechnology(t *testing.T) {
    company := Company{"Test company"}
    stack, _ := NewStack(&company)

    tech, _ := NewTechnology("Martini")
    golang := Language{"Go"}

    tech.AddLanguage(&golang)

    err := stack.RemoveTechnology(tech)
    assert.Nil(t, err)

    technologies := stack.GetTechnologies() 
    assert.Equal(t, 0, len(technologies))
    if len(technologies) != 0 {
        t.Error("No technology removed")
    }

    err = stack.RemoveTechnology(&Technology{})
    assert.NotNil(t, err)
}