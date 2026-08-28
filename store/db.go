package store

import (
	"encoding/json"
	"fmt"
	"go.etcd.io/bbolt"
	"sync"
	"training41/domain"
)

var buckets = []string{"records", "users", "events", "audits", "trainings"}

type DB struct {
	b  *bbolt.DB
	mu sync.RWMutex
}

func Open(path string) (*DB, error) {
	b, e := bbolt.Open(path, 0600, nil)
	if e != nil {
		return nil, e
	}
	d := &DB{b: b}
	e = d.init()
	if e != nil {
		b.Close()
		return nil, e
	}
	return d, nil
}
func (d *DB) init() error {
	return d.b.Update(func(tx *bbolt.Tx) error {
		for _, n := range buckets {
			if _, e := tx.CreateBucketIfNotExists([]byte(n)); e != nil {
				return e
			}
		}
		return nil
	})
}
func (d *DB) Close() error         { return d.b.Close() }
func encode(v any) ([]byte, error) { return json.Marshal(v) }
func (d *DB) put(bucket, key string, v any) error {
	raw, e := encode(v)
	if e != nil {
		return e
	}
	return d.b.Update(func(tx *bbolt.Tx) error { return tx.Bucket([]byte(bucket)).Put([]byte(key), raw) })
}
func (d *DB) get(bucket, key string, out any) error {
	return d.b.View(func(tx *bbolt.Tx) error {
		v := tx.Bucket([]byte(bucket)).Get([]byte(key))
		if v == nil {
			return fmt.Errorf("%s %s not found", bucket, key)
		}
		return json.Unmarshal(v, out)
	})
}
func (d *DB) SaveRecord(r domain.Record) error {
	d.mu.Lock()
	defer d.mu.Unlock()
	return d.put("records", r.ID, r)
}
func (d *DB) GetRecord(id string) (domain.Record, error) {
	var r domain.Record
	e := d.get("records", id, &r)
	return r, e
}
func (d *DB) SaveUser(u domain.User) error { return d.put("users", u.ID, u) }
func (d *DB) GetUser(id string) (domain.User, error) {
	var u domain.User
	e := d.get("users", id, &u)
	return u, e
}
func (d *DB) SaveEvent(v domain.Event) error       { return d.put("events", v.ID, v) }
func (d *DB) SaveAudit(v domain.Audit) error       { return d.put("audits", v.ID, v) }
func (d *DB) SaveTraining(v domain.Training) error { return d.put("trainings", v.ID, v) }
func (d *DB) GetTraining(id string) (domain.Training, error) {
	var t domain.Training
	e := d.get("trainings", id, &t)
	return t, e
}
func (d *DB) ListRecords() ([]domain.Record, error) {
	out := []domain.Record{}
	e := d.b.View(func(tx *bbolt.Tx) error {
		return tx.Bucket([]byte("records")).ForEach(func(_, v []byte) error {
			var r domain.Record
			if e := json.Unmarshal(v, &r); e != nil {
				return e
			}
			out = append(out, r)
			return nil
		})
	})
	return out, e
}
