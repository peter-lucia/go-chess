package ui

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"net/http"
	"strconv"
)

type MoveCallback func(
	piece string,
	start string,
	end string) (bool, error)

type MoveResponse struct {
	Piece   string `json:"piece"   xml:"piece"`
	Start   string `json:"start"   xml:"start"`
	End     string `json:"end"     xml:"end"`
	Success string `json:"success" xml:"success"`
}

func handleMove(c echo.Context, cb MoveCallback) error {
	// in echo FormValue are json body parameters
	piece := c.FormValue("piece")
	start := c.FormValue("start")
	end := c.FormValue("end")

	success, err := cb(piece, start, end)

	if err != nil {
		moveResponse := &MoveResponse{
			Piece:   piece,
			Start:   start,
			End:     end,
			Success: strconv.FormatBool(false),
		}
		return c.JSON(http.StatusInternalServerError, moveResponse)
	}
	moveResponse := &MoveResponse{
		Piece:   piece,
		Start:   start,
		End:     end,
		Success: strconv.FormatBool(success),
	}
	if success {
		return c.JSON(http.StatusOK, moveResponse)
	} else {
		return c.JSON(http.StatusBadRequest, moveResponse)
	}
}

func StartUI(cb MoveCallback) {
	// Create a new instance of Echo
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Serve static files
	e.Static("/css", "ui/css/")
	e.Static("/js", "ui/js/")
	e.Static("/img", "ui/img/")
	e.File("/", "ui/index.html")

	e.POST(
		"/move",
		func(c echo.Context) error {
			return handleMove(c, cb)
		},
	)

	// Start the server
	e.Logger.Fatal(e.Start(":8080"))
}
