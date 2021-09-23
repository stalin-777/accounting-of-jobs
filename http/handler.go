package http

import (
	"fmt"
	"net/http"

	"github.com/jackc/pgx"
	"github.com/labstack/echo/v4"
	aoj "github.com/stalin-777/accounting-of-jobs"
)

type Handler struct {
	WorkplaceService aoj.WorkplaceService
}

//Workplace - get workplace by ID
func (h *Handler) Workplace(c echo.Context) error {

	id, err := getIDFromRequest(c.Param("id"))
	if err != nil {
		return respondWithErrorStatus(c, http.StatusBadRequest, err.Error())
	}

	workplace, err := h.WorkplaceService.Workplace(id)
	if err != nil {
		return respondWithErrorStatus(c, http.StatusBadRequest, err.Error())
	}

	return respondWithData(c, workplace)
}

//Workplaces - get list workplaces
func (h *Handler) Workplaces(c echo.Context) error {

	workplaces, err := h.WorkplaceService.Workplaces()
	if err != nil {
		return respondWithErrorStatus(c, http.StatusBadRequest, err.Error())
	}

	return respondWithData(c, workplaces)
}

//CreateWorkplace - create a new row in DB
func (h *Handler) CreateWorkplace(c echo.Context) error {

	workplace := &aoj.Workplace{}
	err := c.Bind(workplace)
	if err != nil {
		return respondWithErrorStatus(c, http.StatusBadRequest, err.Error())
	}

	//fmt.Println(c.RealIP())

	err = h.WorkplaceService.CreateWorkplace(workplace)
	if err != nil {
		if pgerr, ok := err.(pgx.PgError); ok {
			fmt.Println(pgerr.ConstraintName)
			if pgerr.ConstraintName == "workplace_username_key" {
				respondWithErrorStatus(c, http.StatusBadRequest, "Не братан")
			}
		}
		return respondWithErrorStatus(c, http.StatusBadRequest, err.Error())
	}

	return respondWithData(c, workplace)
}

//UpdateWorkplace - update workplace data in DB by ID
func (h *Handler) UpdateWorkplace(c echo.Context) error {

	id, err := getIDFromRequest(c.Param("id"))
	if err != nil {
		return respondWithErrorStatus(c, http.StatusBadRequest, err.Error())
	}

	workplace := &aoj.Workplace{ID: id}
	err = c.Bind(workplace)
	if err != nil {
		return respondWithErrorStatus(c, http.StatusBadRequest, err.Error())
	}

	err = h.WorkplaceService.UpdateWorkplace(workplace)
	if err != nil {
		return respondWithErrorStatus(c, http.StatusBadRequest, err.Error())
	}

	return respondWithData(c, workplace)
}

//DeleteWorkplace - delete workplace by ID
func (h *Handler) DeleteWorkplace(c echo.Context) error {

	id, err := getIDFromRequest(c.Param("id"))
	if err != nil {
		return respondWithErrorStatus(c, http.StatusBadRequest, err.Error())
	}

	err = h.WorkplaceService.DeleteWorkplace(id)
	if err != nil {
		return respondWithErrorStatus(c, http.StatusBadRequest, err.Error())
	}

	return respondWithNoData(c)
}
