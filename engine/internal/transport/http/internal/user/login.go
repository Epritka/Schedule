package user

import (
	httpassistant "engine/infrastructure/http-assistant"
	"net/http"
)

type loginRequest struct {
	Login     string `json:"login"`
	GroupName string `json:"groupName"`
}

func (h *handlers) Login(w http.ResponseWriter, r *http.Request) {
	lr := loginRequest{}

	err := httpassistant.ParseBody(r.Body, &lr)
	if err != nil {
		httpassistant.Reply(w, r, http.StatusBadRequest, nil)
		return
	}

	id := h.userUseCase.Update(lr.Login, lr.GroupName)

	userInfo := map[string]int{
		"groupId": *id,
	}

	httpassistant.Reply(w, r, http.StatusCreated, &userInfo)
}
