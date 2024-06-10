package model

type Storage struct {
	data map[string]interface{}
}

func (s *Storage) Save(key string, value interface{}) {
	s.data[key] = value
}

func (s *Storage) Get(key string) interface{} {
	return s.data[key]
}

func (s *Storage) Remove(key string) {
	delete(s.data, key)
}

var StorageData Storage

func init() {
	data := make(map[string]interface{})

	StorageData = Storage{
		data: data,
	}
}
