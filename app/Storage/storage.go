package Storage

import "sync"

type Storage struct {
	data map[string]string
	mu   sync.RWMutex // Mutex для синхронизации доступа к данным
}

var storage *Storage

// New создаёт и возвращает новое хранилище
func New() *Storage {
	return &Storage{
		data: make(map[string]string),
	}
}

// Init инициализирует глобальное хранилище
func Init() {
	storage = New()
}

// GetInstance возвращает ссылку на глобальное хранилище
func GetInstance() *Storage {
	return storage
}

// Get получает значение по ключу
func (s *Storage) Get(key string) (string, bool) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	value, ok := s.data[key]
	return value, ok
}

// Set устанавливает значение по ключу
func (s *Storage) Set(key string, value string) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.data[key] = value
}
