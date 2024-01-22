package domain

type Seller struct {
	ID          int    `json:"id,omitempty"`
	CID         int    `json:"cid,omitempty"`
	CompanyName string `json:"company_name,omitempty"`
	Address     string `json:"address,omitempty"`
	Telephone   string `json:"telephone,omitempty"`
	LocalityId  int    `json:"locality_id,omitempty"`
}
