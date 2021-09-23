package accountingofjobs

import "net"

type Workplace struct {
	ID       int    `json:"id"`
	Hostname string `json:"hostname"`
	IP       net.IP `json:"ip"`
	Username string `json:"username"`
}

type WorkplaceService interface {
	Workplace(id int) (*Workplace, error)
	Workplaces() ([]*Workplace, error)
	CreateWorkplace(w *Workplace) error
	UpdateWorkplace(w *Workplace) error
	DeleteWorkplace(id int) error
}
