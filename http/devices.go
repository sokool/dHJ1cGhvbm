package http

import (
	"net/http"
	. "truphone/model"
)

type devices struct {
	handler
	repository Devices
}

func newDevices(r Devices) *devices {
	return &devices{repository: r}
}

func (d *devices) add(w http.ResponseWriter, r *http.Request) {
	var m Device
	if !d.read(w, r, &m) {
		return
	}

	if err := d.repository.Add(&m); err != nil {
		d.write(w, err)
		return
	}

	d.write(w, m)
}

func (d *devices) get(w http.ResponseWriter, r *http.Request) {
	i, ok := d.id(w, r)
	if !ok {
		return
	}

	d.one(w, Query{ID: i})
}

func (d *devices) search(w http.ResponseWriter, r *http.Request) {
	q, ok := d.query(w, r)
	if !ok {
		return
	}

	d.all(w, q)
}

func (d *devices) update(w http.ResponseWriter, r *http.Request) {
	var m Device
	var ok bool

	if m.ID, ok = d.id(w, r); !ok {
		return
	}

	if !d.read(w, r, &m) {
		return
	}

	if err := d.repository.Update(&m); err != nil {
		d.write(w, err)
		return
	}

	d.write(w, m)
}

func (d *devices) delete(w http.ResponseWriter, r *http.Request) {
	id, ok := d.id(w, r)
	if !ok {
		return
	}

	if err := d.repository.Delete(id); err != nil {
		d.write(w, err)
	}
}

func (d *devices) id(w http.ResponseWriter, r *http.Request) (ID, bool) {
	n, err := NewParseID(d.param(r, "device"))
	if err != nil {
		d.write(w, err)
	}
	return n, err == nil
}

func (d *devices) query(w http.ResponseWriter, r *http.Request) (Query, bool) {
	var m Query
	return m, d.url(w, r, &m)
}

func (d *devices) all(w http.ResponseWriter, q Query) {
	m, err := d.repository.Search(q)
	if err != nil {
		d.write(w, err)
		return
	}

	d.write(w, m)
}

func (d *devices) one(w http.ResponseWriter, q Query) {
	m, err := d.repository.Search(q)
	if err != nil {
		d.write(w, err)
		return
	}

	switch len(m) {
	case 0:
		d.write(w, nil)
	case 1:
		d.write(w, m[0])
	}
}
