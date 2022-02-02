package models

type Status2 struct {
	Action   string   `json:"action,omitempty"`
	RigID    string   `json:"rigId,omitempty"`
	Options  []string `json:"options,omitempty"`
	DeviceID string   `json:"deviceId,omitempty"`
}
