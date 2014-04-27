package models

import (
    "testing"
    "github.com/stretchr/testify/assert"
    "encoding/json"
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
    golang := Language{Name:"Go"}

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
    golang := Language{Name:"Go"}
    python := Language{Name:"Python"}
    ruby := Language{Name:"Ruby"}

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

func Test_Technology_MarshalJSON(t *testing.T) {
    tech, _ := NewTechnology("Test")
    golang := &Language{Name:"Go"}
    python := &Language{Name:"Python"}

    tech.AddLanguage(golang) 
    tech.AddLanguage(python) 

    b, err := tech.MarshalJSON()

    assert.Nil(t, err)

    expectedMarshalling := make(map[string]interface{})
    expectedMarshalling["languages"] = append(make([]Language, 0), *golang, *python)
    expectedMarshalling["name"] = tech.Name

    marshalledExpected, _ := json.Marshal(expectedMarshalling)

    assert.Equal(t, marshalledExpected, b)
}

func Test_NewStack(t *testing.T) {
    company := Company{Name:"Test company"}
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
    company := Company{Name:"Test company"}
    stack, _ := NewStack(&company)

    tech, _ := NewTechnology("Martini")
    golang := Language{Name:"Go"}

    tech.AddLanguage(&golang)

    err := stack.AddTechnology(tech)
    assert.Nil(t, err)

    technologies := stack.GetTechnologies() 
    assert.Equal(t, 1, len(technologies))
    assert.Equal(t, tech.Name, technologies[0].Name)
}

func Test_RemoveTechnology(t *testing.T) {
    company := Company{Name:"Test company"}
    stack, _ := NewStack(&company)

    tech, _ := NewTechnology("Martini")
    golang := Language{Name:"Go"}

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

func Test_Stack_MarshalJSON(t *testing.T) {
    company := Company{Name:"Test company"}
    stack, _ := NewStack(&company)

    tech, _ := NewTechnology("Martini")
    golang := &Language{Name:"Go"}

    tech.AddLanguage(golang) 
    stack.AddTechnology(tech)

    b, err := stack.MarshalJSON()
    assert.Nil(t, err)

    expectedMarshalling := make(map[string]interface{})
    expectedMarshalling["technologies"] = stack.GetTechnologies()
    expectedMarshalling["company"] = company

    marshalledExpected, _ := json.Marshal(expectedMarshalling)

    assert.Equal(t, string(marshalledExpected), string(b))
}