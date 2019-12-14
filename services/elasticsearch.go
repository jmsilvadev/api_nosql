package services

import (
	"api-nosql/config"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"reflect"
	"strconv"
	"strings"
)

var opType string
var prefSequence = ".sequence-"

// DBListIndexes - List All Indexes in Elastic
func (p *ElasticSearchProvider) DBListIndexes() ([]ResponseListIndex, error) {

	var respArray []ResponseListIndex

	resp, err := http.Get("http://" + config.Host + ":" + config.ManagerPort + "/_cat/indices?format=json&h=index,health,docs.count")
	if err != nil {
		return respArray, err
	}
	defer resp.Body.Close()

	respBody, err := ioutil.ReadAll(resp.Body)
	decoder := json.NewDecoder(strings.NewReader(string(respBody)))
	err = decoder.Decode(&respArray)

	return respArray, err
}

// bdNewIndex (In CamelCase because is a private method) - Create an New Index
func (p *ElasticSearchProvider) bdNewIndex(indexName string, params []ParamIndex) (ResponseIndex, ResponseErrorIndex, error) {

	var respArray ResponseIndex
	var respErrorArray ResponseErrorIndex
	var body io.Reader

	if len(params) > 0 {
		fields := make(map[string]Type)
		for index := range params {
			vType := Type{params[index].Type}
			fields[params[index].Field] = vType
		}

		data := PayloadCreateIndex{
			Settings{
				NumberOfShards: 1,
			},
			Mappings{
				Properties: fields,
			},
		}

		payloadBytes, err := json.Marshal(data)

		if err != nil {
			return respArray, respErrorArray, err
		}

		body = bytes.NewReader(payloadBytes)
	}

	req, err := http.NewRequest("PUT", "http://"+config.Host+":"+config.ManagerPort+"/"+strings.ToLower(indexName), body)
	if err != nil {
		return respArray, respErrorArray, err
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return respArray, respErrorArray, err
	}
	defer resp.Body.Close()

	respBody, err := ioutil.ReadAll(resp.Body)

	decoder := json.NewDecoder(strings.NewReader(string(respBody)))

	if strings.Contains(string(respBody), "error") {
		err = decoder.Decode(&respErrorArray)
	} else {
		err = decoder.Decode(&respArray)
	}

	return respArray, respErrorArray, err
}

// DBCreateIndex - Create an Index and a Sequence to control IDs
func (p *ElasticSearchProvider) DBCreateIndex(indexName string, params []ParamIndex) (ResponseIndex, ResponseErrorIndex, error) {

	respIndex, respError, err := p.bdNewIndex(indexName, params)
	if err != nil || respError.Status >= 400 {
		return respIndex, respError, err
	}

	json := `{"last_id": 0 }`
	p.dbSave(prefSequence+indexName, json, 1)

	return respIndex, respError, err
}

// DBDeleteIndex - Delete an Index
func (p *ElasticSearchProvider) DBDeleteIndex(indexName string) (ResponseIndex, error) {

	var respArray ResponseIndex

	req, err := http.NewRequest("DELETE", "http://"+config.Host+":"+config.ManagerPort+"/"+strings.ToLower(indexName), nil)
	if err != nil {
		return respArray, err
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return respArray, err
	}
	defer resp.Body.Close()
	respBody, err := ioutil.ReadAll(resp.Body)

	decoder := json.NewDecoder(strings.NewReader(string(respBody)))
	err = decoder.Decode(&respArray)

	return respArray, err
}

// dbSave (In CamelCase because is a private method) - Save in an Index
func (p *ElasticSearchProvider) dbSave(indexName string, contentDocument string, Identifier int) (ResponseSaveDocument, ResponseErrorIndex, error) {

	var respArray ResponseSaveDocument
	var respErrorArray ResponseErrorIndex

	if opType == "insert" {
		opType = "?op_type=create"
	}

	payloadBytes := []byte(contentDocument)

	body := bytes.NewReader(payloadBytes)

	req, err := http.NewRequest("PUT", "http://"+config.Host+":"+config.ManagerPort+"/"+strings.ToLower(indexName)+"/_doc/"+strconv.Itoa(Identifier)+opType, body)
	if err != nil {
		return respArray, respErrorArray, err
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return respArray, respErrorArray, err
	}
	defer resp.Body.Close()
	respBody, err := ioutil.ReadAll(resp.Body)

	decoder := json.NewDecoder(strings.NewReader(string(respBody)))

	if strings.Contains(string(respBody), "error") {
		err = decoder.Decode(&respErrorArray)
	} else {
		err = decoder.Decode(&respArray)
	}

	return respArray, respErrorArray, err
}

//DBInsertDocument - Force a create new Document
func (p *ElasticSearchProvider) DBInsertDocument(indexName string, contentDocument string) (ResponseSaveDocument, ResponseErrorIndex, error) {
	opType = "insert"
	vID, err := p.DBGetLastID(indexName)

	respArray, respError, err := p.dbSave(indexName, contentDocument, vID+1)
	if err == nil && respError.Status == 0 {
		opType = ""
		json := `{"last_id":` + strconv.Itoa(vID+1) + `}`
		p.dbSave(prefSequence+indexName, json, 1)
	}
	return respArray, respError, err
}

//DBUpdateDocument - Force a update an existent Document
func (p *ElasticSearchProvider) DBUpdateDocument(indexName string, contentDocument string, Identifier int) (ResponseSaveDocument, ResponseErrorIndex, error) {
	opType = ""
	return p.dbSave(indexName, contentDocument, Identifier)
}

// DBGetLastID - Get Last Id in an Index
func (p *ElasticSearchProvider) DBGetLastID(indexName string) (int, error) {
	var respArray ResponseLastId
	resp, err := http.Get("http://" + config.Host + ":" + config.ManagerPort + "/" + strings.ToLower(prefSequence+indexName) + "/_source/1")
	if err != nil {
		return 0, err
	}
	resp.Header.Set("Content-Type", "application/json")

	defer resp.Body.Close()

	respBody, err := ioutil.ReadAll(resp.Body)

	decoder := json.NewDecoder(strings.NewReader(string(respBody)))
	err = decoder.Decode(&respArray)

	return respArray.LastID, err
}

// DBFind - Get a Document by an Id in an Index
func (p *ElasticSearchProvider) DBFind(indexName string, Identifier int) (map[string]interface{}, error) {

	var respArray = map[string]interface{}{}

	resp, err := http.Get("http://" + config.Host + ":" + config.ManagerPort + "/" + strings.ToLower(indexName) + "/_source/" + strconv.Itoa(Identifier))
	if err != nil {
		return nil, err
	}
	resp.Header.Set("Content-Type", "application/json")

	defer resp.Body.Close()

	respBody, err := ioutil.ReadAll(resp.Body)

	err = json.Unmarshal(respBody, &respArray)
	if err != nil {
		return nil, err
	}

	return respArray, err
}

// DBQuery - Get a Document by an Id in an Index
func (p *ElasticSearchProvider) DBQuery(indexName string, queryString []QueryString) (string, error) {

	var respArray ResponseQuery
	var processedArray = []map[string]interface{}{}
	var processedQuery string
	var respJSON string

	for index, value := range queryString {
		if index == 0 && value.Operator != "" {
			value.Operator = ""
		}
		processedQuery = processedQuery + value.Operator + "(" + value.Field + ": " + value.Value + ") "
	}

	fullQuery := `
	{
		"query": {
			"query_string": {
				"query": "` + processedQuery + `"
			}
		}
	}`

	payloadBytes := []byte(fullQuery)

	body := bytes.NewReader(payloadBytes)

	req, err := http.NewRequest("GET", "http://"+config.Host+":"+config.ManagerPort+"/"+strings.ToLower(indexName)+"/_search", body)
	if err != nil {
		return "", err
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	respBody, err := ioutil.ReadAll(resp.Body)

	decoder := json.NewDecoder(strings.NewReader(string(respBody)))
	err = decoder.Decode(&respArray)

	if err != nil || respArray.Hits.Total.Value == 0 {
		return "", err
	}

	processedArray = respArray.Hits.Hits
	respJSON = "["
	for _, valueProcessed := range processedArray {
		respJSON = respJSON + "{"
		for index, value := range valueProcessed {
			if index == "_id" {
				switch str := value.(type) {
				case string:
					respJSON = respJSON + `"id": "` + str + `",`
				case int:
					respJSON = respJSON + `"id": "` + strconv.Itoa(str) + `",`
				case int64:
					respJSON = respJSON + `"id": "` + strconv.FormatInt(str, 10) + `",`
				}
			} else if index == "_source" {
				source := reflect.ValueOf(value)
				if source.Kind() == reflect.Map {
					for _, sourceIndex := range source.MapKeys() {
						sourceValue := source.MapIndex(sourceIndex)
						switch str := sourceValue.Interface().(type) {
						case string:
							respJSON = respJSON + `"` + sourceIndex.Interface().(string) + `": "` + str + `",`
						case int:
							respJSON = respJSON + `"` + sourceIndex.Interface().(string) + `": "` + strconv.Itoa(str) + `",`
						case int64:
							respJSON = respJSON + `"` + sourceIndex.Interface().(string) + `": "` + strconv.FormatInt(str, 10) + `",`
						case float64:
							respJSON = respJSON + `"` + sourceIndex.Interface().(string) + `": "` + fmt.Sprintf("%f", str) + `",`
						}
					}
				}
			}
		}
		respJSON = respJSON + "},"
	}
	respJSON = respJSON + "]"
	respJSON = strings.ReplaceAll(respJSON, ",}", "}")
	respJSON = strings.ReplaceAll(respJSON, ",]", "]")
	return respJSON, err
}
