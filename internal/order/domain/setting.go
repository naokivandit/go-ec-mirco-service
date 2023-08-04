package domain

type Setting struct {
	ID             int  `json:"id"`
	MentenanceMode bool `json:"mentenance_mode"`
}

func (s Setting) IsMaintenance() bool {
	return s.MentenanceMode
}
