package kvs

type StorageRecord struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type KvStorage interface {
	NewDB(s string) error
	RemoveDB(s string) error
	ListDB() []string
	Name() string
	Insert(db, key, value string) error
	GetValue(db, key string) (string, error)
	GetKeys(db string) ([]string, error)
}
