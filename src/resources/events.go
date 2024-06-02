package resources

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/danielronalds/go-restful-api/src/db"
	"github.com/labstack/echo/v4"
)

type Event struct {
	Id          int    `db:"id"`
	Name        string `db:"name"`
	Description string `db:"description"`
	Startdate   string `db:"startdate"`
	Enddate     string `db:"enddate"`
}

type PostedEvent struct {
	Name        string `json:"Name" form:"Name" query:"Name"`
	Description string `json:"Description" form:"Description" query:"Description"`
	Startdate   string `json:"Startdate" form:"Startdate" query:"Startdate"`
	Enddate     string `json:"Enddate" form:"Enddate" query:"Enddate"`
}

func GetEvents(c echo.Context) error {
	pg := db.GetDatabase()

	events := []Event{}

	err := pg.Select(&events, "SELECT * FROM api.Events LIMIT 100")
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, events)
}

func GetEventById(c echo.Context) error {
	idStr := c.Param("id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	pg := db.GetDatabase()

	event := Event{}

	query := fmt.Sprintf("SELECT * FROM api.Events WHERE Id = %v", id)
	err = pg.Get(&event, query)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, event)
}

func CreateEvent(c echo.Context) error {
	postedEvent := new(PostedEvent)

	if err := c.Bind(&postedEvent); err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	pg := db.GetDatabase()

	query := `INSERT INTO api.Events (Name, Description, Startdate, Enddate) VALUES($1, $2, $3, $4)`

	result, err := pg.Exec(query, postedEvent.Name, postedEvent.Description, postedEvent.Startdate, postedEvent.Enddate)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	rowsAffected, err := result.RowsAffected()

	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	return c.String(http.StatusOK, fmt.Sprintf("Affected %v row", rowsAffected))
}
