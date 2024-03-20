package kvs

import (
	"encoding/json"
	"fmt"
	"os"
	"path"
	"strings"

	"github.com/sirupsen/logrus"
)

type JSONStorage struct {
	location string
	logger   *logrus.Logger
}

func NewJSONStorage(location string, logger *logrus.Logger) (*JSONStorage, error) {

	if location == "" {
		homeDir, err := os.UserHomeDir()
		if err != nil {
			return nil, fmt.Errorf("can't get user home dir")
		}
		location = path.Join(homeDir, ".local", "share", "kvs")
		if err := os.MkdirAll(location, 0755); err != nil {
			return nil, fmt.Errorf("can't create '%s' dir", location)
		}
	}

	j := &JSONStorage{
		location: location,
		logger:   logger,
	}
	return j, nil
}

func (s *JSONStorage) ListDB() []string {
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

func (s *JSONStorage) GetKeys(db string) ([]string, error) {
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

func (s *JSONStorage) GetValue(db, key string) (string, error) {
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

func (s *JSONStorage) existDB(db string) error {
	_, err := os.Open(path.Join(s.location, db))
	return err
}

func (s *JSONStorage) Insert(db, key, value string) error {
	if err := s.existDB(db); err != nil {
		return err
	}

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

func (s *JSONStorage) NewDB(db string) error {
	if err := s.existDB(db); err == nil {
		return fmt.Errorf("db already exists")
	}
	_, err := os.Create(path.Join(s.location, db))
	return err
}

func (s *JSONStorage) RemoveDB(name string) error {
	return os.Remove(path.Join(s.location, name))
}

func (s *JSONStorage) Name() string {
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
