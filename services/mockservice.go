package services

import (
	//"errors"
)

//ManagerImplementationMock - Implementation of interface
type ManagerImplementationMock struct{}

//MockProvider - Mock to execute functional tests
type MockProvider struct {
	ManagerImplementationMock
}

// DBIndexData - Get Index Data
func (p MockProvider) DBIndexData() {

	//ConnectElasticSearch()
	//return errors.New("test")
}