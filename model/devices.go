package model

type Query struct {
	ID    ID
	Brand Brand
}

type Devices interface {
	Add(*Device) error
	Update(*Device) error
	Delete(ID) error
	Search(Query) ([]Device, error)
}

// device in memory representation of devices, just for tests
type devices map[ID]*Device

func NewDevices() Devices {
	return make(devices)
}

func (r devices) Delete(id ID) error {
	delete(r, id)
	return nil
}

func (r devices) Add(d *Device) (err error) {
	d.ID, err = NewID(sequence + 1)
	if err != nil {
		return err
	}

	r[d.ID] = d
	sequence++

	return nil
}

func (r devices) Update(d *Device) error {
	m, found := r[d.ID]
	if !found {
		return Err("device: not found")
	}

	if d.Brand != m.Brand {
		m.Brand = d.Brand
	}

	if d.Name != m.Name {
		m.Name = d.Name
	}

	d = m

	return nil
}

func (r devices) Search(q Query) ([]Device, error) {
	var o []Device
	for _, d := range r {
		if !q.ID.IsZero() && q.ID != d.ID {
			continue
		}

		if !q.Brand.IsZero() && q.Brand != d.Brand {
			continue
		}

		o = append(o, *d)
	}

	return o, nil
}

var sequence int
