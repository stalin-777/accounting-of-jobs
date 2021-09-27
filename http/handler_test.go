package http

import (
	"errors"
	"log"
	"net"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	aoj "github.com/stalin-777/accounting-of-jobs"
	"github.com/stalin-777/accounting-of-jobs/logger"
	"github.com/stalin-777/accounting-of-jobs/mock"
	"github.com/stretchr/testify/assert"
)

var (
	// mockDB = map[int]*aoj.Workplace{
	// 	1: &aoj.Workplace{1, "localhost", net.ParseIP("127.0.0.1"), "user1"},
	// }
	workplaceJSON = `{"data":{"id":1,"hostname":"localhost","ip":"127.0.0.1","username":"user1"},"success":true}
`
)

func TestHandler_Workplace(t *testing.T) {

	// Setup
	err := logger.InitZapLogger(
		"../logs",
		"log.log",
		10,
		0,
		0,
	)
	if err != nil {
		log.Fatal(err)
	}

	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/workplaces/:id")
	c.SetParamNames("id")
	c.SetParamValues("1")

	wps := &mock.WorkplaceService{}
	wps.WorkplaceFn = func(id int) (*aoj.Workplace, error) {
		if id == 1 {
			return &aoj.Workplace{1, "localhost", net.ParseIP("127.0.0.1"), "user1"}, nil
		}
		return nil, errors.New("Some error")
	}
	h := &Handler{WorkplaceService: wps}

	// Assertions
	if assert.NoError(t, h.Workplace(c)) {
		if !wps.WorkplaceInvoked {
			t.Fatal("expected Workplace() to be invoked")
		}
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, workplaceJSON, rec.Body.String())

	}
}
