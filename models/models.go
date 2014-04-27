package models

import (
    // "fmt"
    "errors"
    "encoding/json"
)
type Language struct {
    Name string
}

type Technology struct {
    Name string
    languages map[string]Language
    //Technologies []Technology
}

func NewTechnology(name string) (*Technology, error) {
    if len(name) == 0 {
        return nil, errors.New("Empty name")
    }
    return &Technology{
        Name: name, 
        languages: make(map[string]Language),
    }, nil
}

func (t *Technology) AddLanguage(l *Language) error {
    if l == nil || len(l.Name) == 0 {
        return errors.New("Empty language")
    }

    t.languages[l.Name] = *l
    return nil
}

func (t *Technology) RemoveLanguage(l *Language) error {
    if len(l.Name) == 0 || &l == nil {
        return errors.New("Empty language")
    }

    delete(t.languages, l.Name)
    return nil
}

func (t *Technology) GetLanguages() []Language {
    l := make([]Language, 0, len(t.languages))

    for  _, language := range t.languages {
       l = append(l, language)
    }

    return l
}

func (t *Technology) MarshalJSON() ([]byte, error) {
    return json.Marshal(map[string]interface{}{
        "name": t.Name,
        "languages": t.GetLanguages(),
    })
}

type Company struct {
    Name string
}

type Stack struct {
    Company Company
    technologies map[string]Technology
}

func NewStack(company *Company) (*Stack, error) {
    if company == nil || len(company.Name) == 0 {
        return nil, errors.New("Empty company")
    }

    return &Stack{
            Company: *company, 
            technologies: make(map[string]Technology),
        }, nil
}

func (s *Stack) AddTechnology(t *Technology) error {
    if t == nil || len(t.Name) == 0 {
        return errors.New("Empty Technology")
    } 

    s.technologies[t.Name] = *t
    return nil
}

func (s *Stack) RemoveTechnology(t *Technology) error {
    if t == nil || len(t.Name) == 0 {
        return errors.New("Empty Technology")
    } 

    delete(s.technologies, t.Name)
    return nil
}

func (s *Stack) GetTechnologies() []Technology {
    t := make([]Technology, 0, len(s.technologies))

    for  _, technology := range s.technologies {
       t = append(t, technology)
    }

    return t
}

func (s *Stack) MarshalJSON() ([]byte, error) {
    return json.Marshal(map[string]interface{}{
        "company": s.Company,
        "technologies": s.GetTechnologies(),
    })
}