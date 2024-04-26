package model

type BaconRequest struct {
	InstanceID string `json:"instance_id"`
	Bacon      Bacon  `json:"bacon"`
}

type Bacon struct {
	BaconName   string `json:"bacon_name"`
	Description string `json:"description"`
}
