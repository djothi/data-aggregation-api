package dcim

// Reference is a nested NetBox object, only its identifier is used.
type Reference struct {
	ID int `json:"id" validate:"required"`
}

type NetworkDevice struct {
	Hostname     string `json:"name" validate:"required"`
	SerialNumber string `json:"serial" validate:"omitempty"`
	Tags         []struct {
		Name string `json:"name" validate:"required"`
	} `json:"tags" validate:"omitempty"`

	// DeviceType and Role identify the port layout of the device.
	// NetBox exposes the role as "role" since 3.6 and as "device_role" before,
	// both are accepted: see RoleID().
	DeviceType *Reference `json:"device_type" validate:"omitempty"`
	Role       *Reference `json:"role"        validate:"omitempty"`
	DeviceRole *Reference `json:"device_role" validate:"omitempty"`
}

// DeviceTypeID returns the NetBox identifier of the device type, 0 when unknown.
func (d *NetworkDevice) DeviceTypeID() int {
	if d.DeviceType == nil {
		return 0
	}
	return d.DeviceType.ID
}

// RoleID returns the NetBox identifier of the device role, 0 when unknown.
func (d *NetworkDevice) RoleID() int {
	switch {
	case d.Role != nil:
		return d.Role.ID
	case d.DeviceRole != nil:
		return d.DeviceRole.ID
	default:
		return 0
	}
}
