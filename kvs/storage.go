package kvs

type StorageContent struct {
	Version int             `json:"version"`
	Records []StorageRecord `json:"records"`
}

type StorageRecord struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type Storage interface {
	NewDB(s string) error
	RemoveDB(s string) error
	ListDB() []string
	Open(db string) error
	Name() string
	Insert(key, value string, force bool) error
	GetValue(key string) (string, error)
	GetKeys() []string
	RemoveValue(v string) error
}

func getStorageRecord(records []StorageRecord, key string) (StorageRecord, int, bool) {
	for k, v := range records {
		if v.Key == key {
			return v, k, true
		}
	}
	return StorageRecord{}, -1, false
}
