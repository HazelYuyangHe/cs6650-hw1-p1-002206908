package album

import "errors"

type MemoryStore struct {
	data []Album
}

func NewMemoryStore() *MemoryStore {
	return &MemoryStore{
		data: []Album{
			{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
			{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
			{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
		},
	}
}

func (s *MemoryStore) List() []Album { return s.data }

func (s *MemoryStore) Get(id string) (Album, error) {
	for _, a := range s.data {
		if a.ID == id {
			return a, nil
		}
	}
	return Album{}, errors.New("not found")
}

func (s *MemoryStore) Add(a Album) Album {
	s.data = append(s.data, a)
	return a
}

func (s *MemoryStore) SearchByTitle(title string) []Album {
	var out []Album
	t := lower(title)
	for _, a := range s.data {
		if contains(lower(a.Title), t) {
			out = append(out, a)
		}
	}
	return out
}

func lower(s string) string {
	b := []rune(s)
	for i, ch := range b {
		if 'A' <= ch && ch <= 'Z' {
			b[i] = ch + 32
		}
	}
	return string(b)
}
func contains(s, sub string) bool {
	if len(sub) == 0 { return true }
	for i := 0; i+len(sub) <= len(s); i++ {
		if s[i:i+len(sub)] == sub { return true }
	}
	return false
}
