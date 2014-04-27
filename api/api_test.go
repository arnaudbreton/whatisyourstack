package api
/*
import (
    "testing"
    "github.com/stretchr/testify/assert"
    "net/http"
    "net/http/httptest"
    // "fmt"
   )

func Test_GetStacks(t *testing.T) {
	recorder := httptest.NewRecorder()	

	m := NewApp()
	req, _ := http.NewRequest("GET", "http://localhost:3000/stacks", nil)
	m.ServeHTTP(recorder, req)

	assert.Equal(t, 200, recorder.Code)
	assert.Equal(t, 1, len(recorder.Header()["Content-Type"]))
	assert.Equal(t, "application/json", recorder.Header()["Content-Type"][0])
	assert.Equal(t, "", recorder.Body.String())
}*/