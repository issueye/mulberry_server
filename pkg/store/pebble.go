package store

import pebbledb "github.com/cockroachdb/pebble"

type PebbleStore struct {
	db   *pebbledb.DB
	mode *pebbledb.WriteOptions
}

func NewPebbleStore(path string) (Store, error) {
	store := &PebbleStore{}
	store.mode = pebbledb.Sync
	db, err := pebbledb.Open(path, &pebbledb.Options{})
	if err != nil {
		panic(err)
	}
	store.db = db
	return store, nil
}

func (p *PebbleStore) Get(key string) ([]byte, error) {
	data, _, err := p.db.Get([]byte(key))
	return data, err
}

func (p *PebbleStore) Set(key string, value []byte) error {
	return p.db.Set([]byte(key), value, p.mode)
}

func (p *PebbleStore) Del(key string) error {
	return p.db.Delete([]byte(key), p.mode)
}

func (p *PebbleStore) Close() error {
	return p.db.Close()
}

func (p *PebbleStore) ForEach(prefix string, fn func(key string, value []byte) error) error {
	keyUpperBound := func(b []byte) []byte {
		end := make([]byte, len(b))
		copy(end, b)
		for i := len(end) - 1; i >= 0; i-- {
			end[i] = end[i] + 1
			if end[i] != 0 {
				return end[:i+1]
			}
		}
		return nil // no upper-bound
	}

	prefixIterOptions := func(prefix []byte) *pebbledb.IterOptions {
		return &pebbledb.IterOptions{
			LowerBound: prefix,
			UpperBound: keyUpperBound(prefix),
		}
	}

	iter, err := p.db.NewIter(prefixIterOptions([]byte(prefix)))
	if err != nil {
		return err
	}

	defer iter.Close()

	for iter.First(); iter.Valid(); iter.Next() {
		err := fn(string(iter.Key()), iter.Value())
		if err != nil {
			return err
		}
	}

	return nil
}

func (p *PebbleStore) Count(prefix string) (int, error) {
	count := 0
	err := p.ForEach(prefix, func(key string, value []byte) error {
		count++
		return nil
	})
	return count, err
}

func (p *PebbleStore) Clear(prefix string) error {
	return p.ForEach(prefix, func(key string, value []byte) error {
		return p.db.Delete([]byte(key), p.mode)
	})
}
