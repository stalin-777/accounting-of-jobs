package mock

import (
	aoj "github.com/stalin-777/accounting-of-jobs"
)

type WorkplaceService struct {
	WorkplaceFn      func(id int) (*aoj.Workplace, error)
	WorkplaceInvoked bool

	WorkplacesFn      func() ([]*aoj.Workplace, error)
	WorkplacesInvoked bool

	CreateWorkplaceFn      func(w *aoj.Workplace) error
	CreateWorkplaceInvoked bool

	UpdateWorkplaceFn      func(w *aoj.Workplace) error
	UpdateWorkplaceInvoked bool

	DeleteWorkplaceFn      func(id int) error
	DeleteWorkplaceInvoked bool
}

func (s *WorkplaceService) Workplace(id int) (*aoj.Workplace, error) {

	s.WorkplaceInvoked = true
	return s.WorkplaceFn(id)
}

func (s *WorkplaceService) Workplaces() ([]*aoj.Workplace, error) {

	s.WorkplacesInvoked = true
	return s.WorkplacesFn()
}
func (s *WorkplaceService) CreateWorkplace(wp *aoj.Workplace) error {

	s.CreateWorkplaceInvoked = true
	return s.CreateWorkplaceFn(wp)
}
func (s *WorkplaceService) UpdateWorkplace(wp *aoj.Workplace) error {

	s.WorkplaceInvoked = true
	return s.UpdateWorkplaceFn(wp)
}
func (s *WorkplaceService) DeleteWorkplace(id int) error {

	s.WorkplaceInvoked = true
	return s.DeleteWorkplaceFn(id)
}
