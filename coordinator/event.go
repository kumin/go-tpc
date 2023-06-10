package coordinator

type TranStatus int

const (
	Ok      TranStatus = 1
	Failure            = 2
)

type TPCEvent struct {
	TransID string     `json:"transaction_id,omitempty"`
	Status  TranStatus `json:"status,omitempty"`
	Type    string     `json:"type,omitempty`
	Payload []byte     `json:"payload,omitempty"`
}
