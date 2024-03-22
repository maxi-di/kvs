package kvs

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path"
	"strings"

	"github.com/sirupsen/logrus"
)

type FileStorage struct {
	location string
	dbName   string
	logger   *logrus.Logger
	content  StorageContent
}

func NewFileStorage(location string, logger *logrus.Logger) (*FileStorage, error) {

	if err := os.MkdirAll(location, 0755); err != nil {
		logger.Fatalf("can't create '%s' dir", location)
	}

	j := &FileStorage{
		location: location,
		dbName:   "",
		logger:   logger,
		content: StorageContent{
			Version: 1,
			Records: []StorageRecord{},
		},
	}
	return j, nil
}

func newStorageContent() StorageContent {
	return StorageContent{
		Version: 1,
		Records: []StorageRecord{},
	}
}

func (s *FileStorage) ListDB() []string {
	return findFiles(s.location, "")
}

func fromJSON2(dbLocation string) (StorageContent, error) {
	content := StorageContent{}
	txt, err := os.ReadFile(dbLocation)
	if err != nil {
		return content, errors.New("can't open db file: " + err.Error())
	}
	err = json.Unmarshal(txt, &content)
	if err != nil {
		return content, errors.New("can't unmarshal db file: " + err.Error())
	}
	return content, nil
}

func (s *FileStorage) GetKeys() []string {
	result := make([]string, 0)
	for _, v := range s.content.Records {
		result = append(result, v.Key)
	}
	return result
}

func (s *FileStorage) GetValue(key string) (string, error) {
	record, _, exist := getStorageRecord(s.content.Records, key)
	if !exist {
		return "", errors.New("value not exist")
	}
	return record.Value, nil
}

func (s *FileStorage) existDB(db string) bool {
	_, err := os.Open(path.Join(s.location, db))
	return err == nil
}

func (s *FileStorage) Open(dbName string) error {
	content, err := fromJSON2(s.makeDBName(dbName))
	if err != nil {
		return err
	}
	s.content = content
	s.dbName = dbName
	s.logger.Info(s.content)
	return nil
}

func toJSON(content StorageContent) ([]byte, error) {
	txt, err := json.MarshalIndent(content, "", "    ")
	if err != nil {
		return nil, err
	}
	return txt, nil
}

func (s *FileStorage) Insert(key, value string, force bool) error {

	var err error = nil
	_, idx, exist := getStorageRecord(s.content.Records, key)

	if exist {
		if force {
			s.content.Records[idx] = StorageRecord{Key: key, Value: value}
			err = s.save()
		} else {
			err = errors.New("value exists, 'force' flag not provided")
		}
	} else {
		s.content.Records = append(s.content.Records, StorageRecord{Key: key, Value: value})
		err = s.save()
	}

	return err
}

func (s *FileStorage) makeDBName(db string) string {
	return path.Join(s.location, db)
}

func (s *FileStorage) save() error {
	if s.dbName == "" {
		return errors.New("db name not specified")
	}

	txt, err := toJSON(s.content)
	if err != nil {
		return err
	}

	err = os.WriteFile(s.makeDBName(s.dbName), txt, 0755)
	if err != nil {
		return err
	}
	return nil
}

func (s *FileStorage) NewDB(db string) error {
	if s.existDB(db) {
		return fmt.Errorf("db already exists")
	}

	s.content = newStorageContent()

	txt, err := toJSON(s.content)
	if err != nil {
		return err
	}

	err = os.WriteFile(s.makeDBName(db), txt, 0755)
	if err != nil {
		return err
	}
	return nil
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

func (s *FileStorage) RemoveValue(key string) error {
	_, idx, exist := getStorageRecord(s.content.Records, key)
	if exist {
		s.content.Records = append(s.content.Records[:idx], s.content.Records[idx+1:]...)
	}
	return s.save()
}
