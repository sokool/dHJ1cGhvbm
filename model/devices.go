package model

type Query struct {
	ID    ID
	Brand Brand
}

type Devices interface {
	Add(Device) error
	Update(Device)
	Search(Query) ([]Device, error)
}

// device in memory representation of devices, just for tests
type devices map[ID]*Device

func NewDevices() Devices {
	return make(devices)
}

func (d devices) Add(device Device) error {
	//TODO implement me
	panic("implement me")
}

func (d devices) Update(device Device) {
	//TODO implement me
	panic("implement me")
}

func (d devices) Search(query Query) ([]Device, error) {
	//TODO implement me
	panic("implement me")
}
