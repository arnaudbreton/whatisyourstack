package models

type Language struct {
	Name string
}

type Technology struct {
	Languages []Language
	//Technologies []Technology
}

func (t *Technology) AddLanguage(l Language) bool {
	if t.Languages == nil {
		t.Languages = make([]Language, 0, 5)
		t.Languages = append(t.Languages, l)

		return true
	}

	for _, tl := range t.Languages {
        if tl == l {
            return false
        }
    }

    if len(t.Languages)+1 > cap(t.Languages) {
    	newLanguages := make([]Language, len(t.Languages)+1, 2*cap(t.Languages))
	    for i := range t.Languages {
	        newLanguages[i] = t.Languages[i]
	    }
	    t.Languages = newLanguages
    }
    
    t.Languages = append(t.Languages, l)

    return true
}

func (t *Technology) RemoveLanguage(l Language) bool {
	if t.Languages == nil {
		return false
	}

	for i, tl := range t.Languages {
        if tl == l {
        	t.Languages = t.Languages[:i+copy(t.Languages[i:], t.Languages[i+1:])]
			return true
        }
    }

    return false
}

type Company struct {
	Name string
}

type Stack struct {
	Company Company
	Technologies []Technology
}