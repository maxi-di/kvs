package kvs

import (
	"encoding/json"
	"fmt"
	"os"
	"path"
	"strings"

	"github.com/sirupsen/logrus"
)

type FileStorage struct {
	location string
	logger   *logrus.Logger
}

func NewFileStorage(location string, logger *logrus.Logger) (*FileStorage, error) {

	if err := os.MkdirAll(location, 0755); err != nil {
		logger.Fatalf("can't create '%s' dir", location)
	}

	j := &FileStorage{
		location: location,
		logger:   logger,
	}
	return j, nil
}

func (s *FileStorage) ListDB() []string {
	return findFiles(s.location, "")
}

func fromJSON(dbLocation string) ([]StorageRecord, error) {
	targets := []StorageRecord{}
	txt, err := os.ReadFile(dbLocation)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(txt, &targets)
	if err != nil {
		return nil, err
	}
	return targets, nil
}

func (s *FileStorage) GetKeys(db string) ([]string, error) {
	targets, err := fromJSON(path.Join(s.location, db))
	if err != nil {
		return []string{}, nil
	}
	result := make([]string, 0)
	for _, v := range targets {
		result = append(result, v.Key)
	}
	return result, nil
}

func (s *FileStorage) GetValue(db, key string) (string, error) {
	targets, err := fromJSON(path.Join(s.location, db))
	if err != nil {
		return "", nil
	}
	record, err := getStorageRecord(targets, key)
	if err != nil {
		return "", err
	}
	return record.Value, nil
}

func (s *FileStorage) existDB(db string) bool {
	_, err := os.Open(path.Join(s.location, db))
	return err == nil
}

func (s *FileStorage) Insert(db, key, value string) error {

	targets, _ := fromJSON(path.Join(s.location, db))

	targets = append(targets, StorageRecord{Key: key, Value: value})
	txt, err := json.MarshalIndent(targets, "", "    ")
	if err != nil {
		return err
	}
	err = os.WriteFile(path.Join(s.location, db), txt, 0755)
	if err != nil {
		return err
	}
	return nil
}

func (s *FileStorage) makeDBName(db string) string {
	return path.Join(s.location, db)
}

func (s *FileStorage) NewDB(db string) error {
	if s.existDB(db) {
		return fmt.Errorf("db already exists")
	}
	_, err := os.Create(s.makeDBName(db))
	return err
}

func (s *FileStorage) RemoveDB(name string) error {
	return os.Remove(path.Join(s.location, name))
}

func (s *FileStorage) Name() string {
	return s.location
}

func findFiles(dir string, pattern string) []string {
	var a []string
	fileInfos, _ := os.ReadDir(dir)
	for _, v := range fileInfos {
		if !v.IsDir() && strings.Contains(v.Name(), pattern) {
			a = append(a, v.Name())
		}
	}
	return a
}
