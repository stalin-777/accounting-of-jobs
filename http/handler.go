package http

import (
	"net/http"

	"github.com/labstack/echo/v4"
	aoj "github.com/stalin-777/accounting-of-jobs"
	"github.com/stalin-777/accounting-of-jobs/logger"
)

type Handler struct {
	WorkplaceService aoj.WorkplaceService
}

//Workplace - get workplace by ID
func (h *Handler) Workplace(c echo.Context) error {

	id, err := getIDFromRequest(c.Param("id"))
	if err != nil {
		logger.Warn(err)
		return respondWithErrorStatus(c, http.StatusBadRequest, err.Error())
	}

	workplace, err := h.WorkplaceService.Workplace(id)
	if err != nil {
		logger.Warn(err)
		return respondWithErrorStatus(c, http.StatusBadRequest, err.Error())
	}

	logger.Infof("Successful attempt to get a workplace. ID: %v", id)
	return respondWithData(c, workplace)
}

//Workplaces - get list workplaces
func (h *Handler) Workplaces(c echo.Context) error {

	workplaces, err := h.WorkplaceService.Workplaces()
	if err != nil {
		logger.Warn(err)
		return respondWithErrorStatus(c, http.StatusBadRequest, err.Error())
	}

	logger.Info("Successful attempt to get a list of workplaces")
	return respondWithData(c, workplaces)
}

//CreateWorkplace - create a new row in DB
func (h *Handler) CreateWorkplace(c echo.Context) error {

	workplace := &aoj.Workplace{}

	err := c.Bind(workplace)
	if err != nil {
		logger.Warn(err)
		return respondWithErrorStatus(c, http.StatusBadRequest, err.Error())
	}

	//fmt.Println(c.RealIP())

	err = h.WorkplaceService.CreateWorkplace(workplace)
	if err != nil {
		logger.Warn(err)
		return respondWithErrorStatus(c, http.StatusBadRequest, err.Error())
	}

	logger.Infof("Successful attempt to create a workplace. ID:%v", workplace.ID)
	return respondWithData(c, workplace)
}

//UpdateWorkplace - update workplace data in DB by ID
func (h *Handler) UpdateWorkplace(c echo.Context) error {

	id, err := getIDFromRequest(c.Param("id"))
	if err != nil {
		logger.Warn(err)
		return respondWithErrorStatus(c, http.StatusBadRequest, err.Error())
	}

	workplace := &aoj.Workplace{ID: id}

	err = c.Bind(workplace)
	if err != nil {
		logger.Warn(err)
		return respondWithErrorStatus(c, http.StatusBadRequest, err.Error())
	}

	err = h.WorkplaceService.UpdateWorkplace(workplace)
	if err != nil {
		logger.Warn(err)
		return respondWithErrorStatus(c, http.StatusBadRequest, err.Error())
	}

	logger.Infof("Successful attempt to delete a workplace. ID:%v, Username:%s", id, workplace.Username)
	return respondWithData(c, workplace)
}

//DeleteWorkplace - delete workplace by ID
func (h *Handler) DeleteWorkplace(c echo.Context) error {

	id, err := getIDFromRequest(c.Param("id"))
	if err != nil {
		logger.Warn(err)
		return respondWithErrorStatus(c, http.StatusBadRequest, err.Error())
	}

	err = h.WorkplaceService.DeleteWorkplace(id)
	if err != nil {
		logger.Warn(err)
		return respondWithErrorStatus(c, http.StatusBadRequest, err.Error())
	}

	logger.Infof("Successful attempt to delete a workplace. ID:%v", id)
	return respondWithNoData(c)
}
