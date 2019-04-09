package handler

import (
	"net/http"
	"number-server/app/domain/model"
	"number-server/app/usecase"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{}

type NumberHandler interface {
	Add(w http.ResponseWriter, r *http.Request)
}

type numberHandler struct {
	numberUseCase usecase.NumberUseCase
}

func NewNumberHandler(numberUseCase usecase.NumberUseCase) NumberHandler {
	return &numberHandler{
		numberUseCase,
	}
}

func (h *numberHandler) Add(w http.ResponseWriter, r *http.Request) {
	upgrader.CheckOrigin = func(r *http.Request) bool { return true }
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		panic(err)
	}

	defer ws.Close()

	for {
		_, n, err := ws.ReadMessage()
		if err != nil {
			break
		}
		number := &model.Number{
			Value: string(n),
		}

		if err := h.numberUseCase.ReadNumber(number); err != nil {
			break
		}
		return
	}
}
