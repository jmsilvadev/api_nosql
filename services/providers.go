package services

//Manager - Interface to provide polimorphism
type Manager interface {
	DBCreateIndex(indexName string, params []ParamIndex) (ResponseIndex, ResponseErrorIndex, error)
	bdNewIndex(indexName string, params []ParamIndex) (ResponseIndex, ResponseErrorIndex, error)
	bdNewSequence(sequenceName string, params []ParamIndex) (ResponseIndex, ResponseErrorIndex, error)
	DBDeleteIndex(indexName string) (ResponseIndex, error)
	DBListIndexes() ([]ResponseListIndex, error)
	DBGetLastID(indexName string) (int, error)
	dbSave(indexName string, contentDocument string, Identifier int) (ResponseSaveDocument, ResponseErrorIndex, error)
	DBInsertDocument(indexName string, contentDocument string) (ResponseSaveDocument, ResponseErrorIndex, error)
	DBUpdateDocument(indexName string, contentDocument string, Identifier int) (ResponseSaveDocument, ResponseErrorIndex, error)
	DBFind(indexName string, Identifier int) (map[string]interface{}, error)
	DBQuery(indexName string, queryString []QueryString) (string, error)
}

//ElasticSearchProvider - Provider of nosql
type ElasticSearchProvider struct {
	Manager
}

type ParamIndex struct {
	Field string
	Type  string
}

//Types for handlres Indexes
type PayloadCreateIndex struct {
	Settings Settings `json:"settings"`
	Mappings Mappings `json:"mappings"`
}
type Settings struct {
	NumberOfShards int `json:"number_of_shards"`
}

type Mappings struct {
	Properties map[string]Type `json:"properties"`
}

type Type struct {
	Type string `json:"type"`
}

//Types for Query
type QueryString struct {
	Field    string
	Value    string
	Operator string
}

//Response Types
type ResponseIndex struct {
	Acknowledged        bool   `json:"acknowledged"`
	Shards_acknowledged bool   `json:"shards_acknowledged"`
	Index               string `json:"index"`
}

type ResponseListIndex struct {
	IndexName string `json:"index"`
	Health    string `json:"health"`
	DocsCount string `json:"docs.count"`
}

type ResponseLastId struct {
	LastID int `json:"last_id"`
}

type ResponseSaveDocument struct {
	Index   string `json:"_index"`
	Type    string `json:"_type"`
	ID      string `json:"_id"`
	Version int    `json:"_version"`
	Result  string `json:"result"`
	Shards  struct {
		Total      int `json:"total"`
		Successful int `json:"successful"`
		Failed     int `json:"failed"`
	} `json:"_shards"`
	SeqNo       int `json:"_seq_no"`
	PrimaryTerm int `json:"_primary_term"`
}

type ResponseErrorIndex struct {
	Error struct {
		RootCause []struct {
			Type      string `json:"type"`
			Reason    string `json:"reason"`
			IndexUUID string `json:"index_uuid"`
			Index     string `json:"index"`
		} `json:"root_cause"`
		Type      string `json:"type"`
		Reason    string `json:"reason"`
		IndexUUID string `json:"index_uuid"`
		Index     string `json:"index"`
	} `json:"error"`
	Status int `json:"status"`
}

type ResponseQuery struct {
	Took     int  `json:"took"`
	TimedOut bool `json:"timed_out"`
	Shards   struct {
		Total      int `json:"total"`
		Successful int `json:"successful"`
		Skipped    int `json:"skipped"`
		Failed     int `json:"failed"`
	} `json:"_shards"`
	Hits struct {
		Total struct {
			Value    int    `json:"value"`
			Relation string `json:"relation"`
		} `json:"total"`
		MaxScore float64                  `json:"max_score"`
		Hits     []map[string]interface{} `json:"hits"`
	} `json:"hits"`
}

//GetManager Select manager to use
func GetManager() Manager {
	//Default Manager
	return &ElasticSearchProvider{}
}
