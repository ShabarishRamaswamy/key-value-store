package models

type KeyValuePair struct {
	Key   interface{} `json:"key"`
	Value interface{} `json:"value"`
}

type GlobalKVStore map[interface{}]interface{}

const GlobalKVStoreName string = "store"
const ErrNoGlovalKVStore string = "error, no global key-value store"
const ErrNoValueInKVStore string = "error, no value in the key-value store"
const SQLITEDBName = "kvs.db"
const SQLITEDBFolderName = "databases"
const SQLITEDBPath = SQLITEDBFolderName + "/" + SQLITEDBName
