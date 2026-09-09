package domain

type TrialClass struct {
	ID             int64  `json:"id"`
	Name           string `json:"name"`
	Quota          int64  `json:"quota"`
	AvailableSlots int64  `json:"available_slots"`
}
