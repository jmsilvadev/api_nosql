package tests

import (
	"api-nosql/services"
	"log"
	"strconv"
	"strings"
	"testing"
	"time"

	"gotest.tools/assert"
)

var vIndexName = "test-" + strconv.FormatInt(time.Now().UnixNano(), 10)
var currManager = services.GetManager()
var prefSequence = ".sequence-"

func TestCreateIndexWithWrongDataType(t *testing.T) {

	var params = make([]services.ParamIndex, 2)

	params[0].Field = "campo1"
	params[0].Type = "number"
	params[1].Field = "campo2"
	params[1].Type = "txklzjvhklzjhvklfdzext"
	respIndex, respErrorIndex, err := currManager.DBCreateIndex(vIndexName, params)
	assert.NilError(t, err)
	assert.Assert(t, respErrorIndex.Status >= 400)
	assert.Assert(t, respIndex.Acknowledged == false)
}

func TestCreateIndex(t *testing.T) {

	var params = make([]services.ParamIndex, 16)

	params[0].Field = "campo1"
	params[0].Type = "text"
	params[1].Field = "campo2"
	params[1].Type = "integer"
	params[2].Field = "campo3"
	params[2].Type = "short"
	params[3].Field = "campo4"
	params[3].Type = "byte"
	params[4].Field = "campo5"
	params[4].Type = "double"
	params[5].Field = "campo6"
	params[5].Type = "float"
	params[6].Field = "campo7"
	params[6].Type = "half_float"
	params[7].Field = "campo9"
	params[7].Type = "date"
	params[8].Field = "campo10"
	params[8].Type = "date_nanos"
	params[9].Field = "campo11"
	params[9].Type = "boolean"
	params[10].Field = "campo12"
	params[10].Type = "binary"
	params[11].Field = "campo13"
	params[11].Type = "integer_range"
	params[12].Field = "campo14"
	params[12].Type = "float_range"
	params[13].Field = "campo15"
	params[13].Type = "long_range"
	params[14].Field = "campo16"
	params[14].Type = "double_range"
	params[15].Field = "campo17"
	params[15].Type = "date_range"

	respIndex, respError, err := currManager.DBCreateIndex(vIndexName, params)
	assert.NilError(t, err)
	assert.Assert(t, respIndex.Acknowledged)
	assert.Assert(t, respError.Status == 0)
}

func TestListIndexes(t *testing.T) {

	respIndex, err := currManager.DBListIndexes()
	assert.NilError(t, err)
	assert.Assert(t, len(respIndex) > 0)
}

func TestGetLastID(t *testing.T) {

	respIndex, err := currManager.DBGetLastID(vIndexName)
	assert.NilError(t, err)
	assert.Assert(t, respIndex >= 0)
}

func TestSaveInIndexWithWrondDataType(t *testing.T) {

	json := `{"campo1":"Eve","campo2":"teste em campo integer com texto"}`
	respSave, respErrorSave, err := currManager.DBInsertDocument(vIndexName, json)

	assert.NilError(t, err)
	assert.Assert(t, respSave.Shards.Successful == 0)
	assert.Assert(t, respErrorSave.Status == 400)
}

func TestSaveInIndex(t *testing.T) {

	//Gravação 1 - Create
	json := `{"name":"Bob SquarePaints","Age":1}`
	respSave, respError, err := currManager.DBInsertDocument(vIndexName, json)

	assert.NilError(t, err)
	assert.Assert(t, respSave.Shards.Successful >= 0)
	assert.Assert(t, respError.Status == 0)

	//Gravação 2 - Update
	json = `{"name":"Bob SquarePaints","Age":1,"Parents":["Patrick","Squidward"]}`
	respSave, respError, err = currManager.DBUpdateDocument(vIndexName, json, 1)

	assert.NilError(t, err)
	assert.Assert(t, respSave.Shards.Successful >= 0)
	assert.Assert(t, respError.Status == 0)

	//Gravação 3 - Create
	json = `{"name":"Squidward","Age":190}`
	respSave, respError, err = currManager.DBInsertDocument(vIndexName, json)

	assert.NilError(t, err)
	assert.Assert(t, respSave.Shards.Successful >= 0)
	assert.Assert(t, respError.Status == 0)

	//Gravação 4 - Create
	json = `{"name":"Patrick","Age":1}`
	respSave, respError, err = currManager.DBInsertDocument(vIndexName, json)

	//Gravação 5 - Create
	json = `{"name":"Squidward Clone","Age":190}`
	respSave, respError, err = currManager.DBInsertDocument(vIndexName, json)

	assert.NilError(t, err)
	assert.Assert(t, respSave.Shards.Successful >= 0)
	assert.Assert(t, respError.Status == 0)

}

func TestDBFind(t *testing.T) {

	respIndex, err := currManager.DBFind(vIndexName, 2)

	for index, value := range respIndex {
		assert.Assert(t, index != "")
		assert.Assert(t, value != nil)
	}

	assert.NilError(t, err)

}

func TestDBQuerySimple(t *testing.T) {

	time.Sleep(500 * time.Millisecond)
	var queryString = make([]services.QueryString, 1)

	queryString[0].Field = "name"
	queryString[0].Value = "*ward*"
	respIndex, err := currManager.DBQuery(vIndexName, queryString)
	log.Println("JSON: " + respIndex)
	assert.Assert(t, strings.Contains(respIndex, "id"))
	assert.NilError(t, err)

}

func TestDBQueryCompound(t *testing.T) {

	time.Sleep(500 * time.Millisecond)
	var queryString = make([]services.QueryString, 2)

	queryString[0].Field = "name"
	queryString[0].Value = "*ward*"
	queryString[1].Field = "Age"
	queryString[1].Value = "*190*"

	respIndex, err := currManager.DBQuery(vIndexName, queryString)
	log.Println("JSON: " + respIndex)
	assert.Assert(t, strings.Contains(respIndex, "id"))
	assert.NilError(t, err)

}

/*
func TestDeleteIndex(t *testing.T) {

	respIndex, err := currManager.DBDeleteIndex(vIndexName)
	assert.NilError(t, err)
	assert.Assert(t, respIndex.Acknowledged)

	respIndex, err = currManager.DBDeleteIndex(prefSequence + vIndexName)
	assert.NilError(t, err)
	assert.Assert(t, respIndex.Acknowledged)
}*/
