package models

import (
    // "fmt"
    "errors"
    "encoding/json"
)
type Language struct {
    Id      int64 `json:"-"`
    Name    string
}

type LanguageApi struct {
    Name        string
}

type Technology struct {
    Id          int64
    Name        string
    languages   map[string]Language
    //Technologies []Technology
}

type TechnologyApi struct {
    Name        string
    Languages   []LanguageApi
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

func NewTechnologyFromApi(techApi *TechnologyApi) (*Technology, error) {
    if techApi == nil {
        return nil, errors.New("Empty Tech Api")
    }

    tech, err := NewTechnology(techApi.Name)

    if err != nil {
        return nil, err
    }

    for _, langApi := range techApi.Languages {
        if 0 == len(langApi.Name) {
            return nil, errors.New("Empty Lang Api")
        }
        langApi := &Language{Name:langApi.Name}
        tech.AddLanguage(langApi)
    }

    return tech, nil
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
    Id      int64 `json:"-"`
    Name    string
}

type CompanyApi struct {
    Name    string
}

type Stack struct {
    Id              int64
    Company         Company
    technologies    map[string]Technology
}

type StackApi struct {
    Company CompanyApi
    Technologies []TechnologyApi
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

func NewStackFromApi(stackApi *StackApi) (*Stack, error) {
    company := &Company{Name:stackApi.Company.Name}
    stack, err := NewStack(company)

    if err != nil {
        return nil, err
    }

    for _, techApi := range stackApi.Technologies {
        tech, err := NewTechnologyFromApi(&techApi)

        if err != nil {
            return nil, err
        }
        stack.AddTechnology(tech)
    }

    return stack, nil
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