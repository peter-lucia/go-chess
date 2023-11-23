package ui

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"net/http"
	"strconv"
)

type MoveCallback func(MoveRequest) (bool, BoardPosition, error)

type MoveResponse struct {
	Request       MoveRequest
	BoardPosition BoardPosition
	Success       string `json:"success" xml:"success"`
}

type MoveRequest struct {
	Start            string        `json:"start"`
	End              string        `json:"end"`
	NewBoardPosition BoardPosition `json:"board"`
}

type BoardPosition struct {
	// typical start:
	// {wR wN wB wQ wK wB wN wR wP wP wP wP wP wP wP wP ... bP bP bP bP bP bP bP bP bR bN bB bQ bK bB bN bR}
	A1 string `json:"a1"`
	B1 string `json:"b1"`
	C1 string `json:"c1"`
	D1 string `json:"d1"`
	E1 string `json:"e1"`
	F1 string `json:"f1"`
	G1 string `json:"g1"`
	H1 string `json:"h1"`

	A2 string `json:"a2"`
	B2 string `json:"b2"`
	C2 string `json:"c2"`
	D2 string `json:"d2"`
	E2 string `json:"e2"`
	F2 string `json:"f2"`
	G2 string `json:"g2"`
	H2 string `json:"h2"`

	A3 string `json:"a3"`
	B3 string `json:"b3"`
	C3 string `json:"c3"`
	D3 string `json:"d3"`
	E3 string `json:"e3"`
	F3 string `json:"f3"`
	G3 string `json:"g3"`
	H3 string `json:"h3"`

	A4 string `json:"a4"`
	B4 string `json:"b4"`
	C4 string `json:"c4"`
	D4 string `json:"d4"`
	E4 string `json:"e4"`
	F4 string `json:"f4"`
	G4 string `json:"g4"`
	H4 string `json:"h4"`

	A5 string `json:"a5"`
	B5 string `json:"b5"`
	C5 string `json:"c5"`
	D5 string `json:"d5"`
	E5 string `json:"e5"`
	F5 string `json:"f5"`
	G5 string `json:"g5"`
	H5 string `json:"h5"`

	B6 string `json:"b6"`
	C6 string `json:"c6"`
	D6 string `json:"d6"`
	E6 string `json:"e6"`
	F6 string `json:"f6"`
	G6 string `json:"g6"`
	H6 string `json:"h6"`

	A7 string `json:"a7"`
	B7 string `json:"b7"`
	C7 string `json:"c7"`
	D7 string `json:"d7"`
	E7 string `json:"e7"`
	F7 string `json:"f7"`
	G7 string `json:"g7"`
	H7 string `json:"h7"`

	A8 string `json:"a8"`
	B8 string `json:"b8"`
	C8 string `json:"c8"`
	D8 string `json:"d8"`
	E8 string `json:"e8"`
	F8 string `json:"f8"`
	G8 string `json:"g8"`
	H8 string `json:"h8"`
}

func handleMove(c echo.Context, cb MoveCallback) error {

	// extract the json body parameters by binding to a move request struct
	var mr MoveRequest
	err := c.Bind(&mr)

	if err != nil {
		return c.String(http.StatusBadRequest, "Bad Request")
	}

	success, newPosition, err := cb(mr)

	if err != nil {
		moveResponse := &MoveResponse{
			Request:       mr,
			Success:       strconv.FormatBool(false),
			BoardPosition: newPosition,
		}
		return c.JSON(http.StatusInternalServerError, moveResponse)
	}
	moveResponse := &MoveResponse{
		Request:       mr,
		Success:       strconv.FormatBool(success),
		BoardPosition: newPosition,
	}
	return c.JSON(http.StatusOK, moveResponse)
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
