package tests

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	/*"bytes"
	"context"
	"log"
	"api-nosql/handlers"
	"api-nosql/routes"*/

	"api-nosql/controllers"
	"api-nosql/models"

	"gotest.tools/assert"
)

func TestHealth(t *testing.T) {

	req, err := http.NewRequest(http.MethodGet, "/", nil)

	if err != nil {
		t.Fatal(err)
	}
	w := httptest.NewRecorder()
	handler := http.HandlerFunc(controllers.HealthInfo)

	handler.ServeHTTP(w, req)

	if status := w.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	expected := "You know, to interact with the NoSQL api. :)"
	assert.Equal(t, w.Body.String(), expected)
	assert.NilError(t, err)
}

func TestLicense(t *testing.T) {

	req, err := http.NewRequest(http.MethodGet, "/license", nil)
	if err != nil {
		t.Fatal(err)
	}
	w := httptest.NewRecorder()
	handler := http.HandlerFunc(controllers.ShowLicense)

	handler.ServeHTTP(w, req)

	if status := w.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	expected := "JMSilvaDev © 2019 | Todos os direitos reservados"
	decoder := json.NewDecoder(strings.NewReader(w.Body.String()))
	var m []models.License
	err = decoder.Decode(&m)
	assert.Equal(t, m[0].Name, expected)
	assert.NilError(t, err)
}

/*
func TestGetIndex(t *testing.T) {

	in := `{"query": ""}`
	rawIn := json.RawMessage(in)
	jsondataTest, err := rawIn.MarshalJSON()

	req, err := http.NewRequest(http.MethodPost, routes.BaseDir+"/", bytes.NewBuffer(jsondataTest))

	if err != nil {
		t.Fatal(err)
	}

	req.Header.Add("Content-Type", "application/json")

	ctx := req.Context()
	ctx = context.WithValue(req.Context(), handlers.XKey("indexName"), "test-index")
	req = req.WithContext(ctx)

	w := httptest.NewRecorder()
	handler := http.HandlerFunc(controllers.GetIndex)
	handler.ServeHTTP(w, req)

	decoder := json.NewDecoder(strings.NewReader(w.Body.String()))
	var i models.IndexResponse
	err = decoder.Decode(&i)

	log.Println("*******")
	log.Println(w.Body.String())
	log.Println("*******")

	if err != nil {
		log.Printf(w.Body.String())
		t.Fatal(err)
	}

	assert.NilError(t, err)
}*/
/*
func TestEPDeleteIndex(t *testing.T) {

	in := `{"query": ""}`
	rawIn := json.RawMessage(in)
	jsondataTest, err := rawIn.MarshalJSON()

	req, err := http.NewRequest(http.MethodDelete, routes.BaseDir+"/", bytes.NewBuffer(jsondataTest))
	req.Header.Add("Content-Type", "application/json")

	if err != nil {
		t.Fatal(err)
	}

	w := httptest.NewRecorder()
	handler := http.HandlerFunc(controllers.DeleteIndex)

	handler.ServeHTTP(w, req)

	expected := "Index Deleted Successfully"

	decoder := json.NewDecoder(strings.NewReader(w.Body.String()))
	var i models.IndexResponse
	err = decoder.Decode(&i)
	assert.Equal(t, i.Resp, expected)
	assert.NilError(t, err)
}
*/
