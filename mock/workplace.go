package mock

import (
	aoj "github.com/stalin-777/accounting-of-jobs"
)

type WorkplaceService struct {
	FindWorkplaceFn      func(id int) (*aoj.Workplace, error)
	FindWorkplaceInvoked bool

	FindWorkplacesFn      func() ([]*aoj.Workplace, error)
	FindWorkplacesInvoked bool

	CreateWorkplaceFn      func(w *aoj.Workplace) error
	CreateWorkplaceInvoked bool

	UpdateWorkplaceFn      func(w *aoj.Workplace) error
	UpdateWorkplaceInvoked bool

	DeleteWorkplaceFn      func(id int) error
	DeleteWorkplaceInvoked bool
}

func (s *WorkplaceService) FindWorkplace(id int) (*aoj.Workplace, error) {

	s.FindWorkplaceInvoked = true
	return s.FindWorkplaceFn(id)
}

func (s *WorkplaceService) FindWorkplaces() ([]*aoj.Workplace, error) {

	s.FindWorkplacesInvoked = true
	return s.FindWorkplacesFn()
}
func (s *WorkplaceService) CreateWorkplace(wp *aoj.Workplace) error {

	s.CreateWorkplaceInvoked = true
	return s.CreateWorkplaceFn(wp)
}
func (s *WorkplaceService) UpdateWorkplace(wp *aoj.Workplace) error {

	s.FindWorkplaceInvoked = true
	return s.UpdateWorkplaceFn(wp)
}
func (s *WorkplaceService) DeleteWorkplace(id int) error {

	s.FindWorkplaceInvoked = true
	return s.DeleteWorkplaceFn(id)
}
