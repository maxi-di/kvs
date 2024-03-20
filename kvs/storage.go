package kvs

type StorageRecord struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type StorageRecords []StorageRecord

type KvStorage interface {
	NewDB(s string) error
	RemoveDB(s string) error
	ListDB() []string
	// Open(db string) error
	Name() string
	Insert(db, key, value string) error
	// Update(db, key, value string) error
	GetValue(db, key string) (string, error)
	GetKeys(db string) ([]string, error)
}

func getStorageRecord(records StorageRecords, key string) (StorageRecord, error) {
	for _, v := range records {
		if v.Key == key {
			return v, nil
		}
	}
	return StorageRecord{}, nil
}
