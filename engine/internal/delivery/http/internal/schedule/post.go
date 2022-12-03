package schedule

import (
	httpassistant "engine/infrastructure/http-assistant"
	"engine/internal/core/entity"
	"net/http"
)

func (h *handlers) UploadNewSchedule(w http.ResponseWriter, r *http.Request) {
	schedules := []entity.Schedule{}

	err := httpassistant.ParseBody(r.Body, &schedules)
	if err != nil {
		httpassistant.Reply(w, r, http.StatusBadRequest, err)
		return
	}

	err = h.scheduleUseCase.Create(schedules)
	if err != nil {
		httpassistant.Reply(w, r, http.StatusBadRequest, err)
		return
	}

	httpassistant.Reply(w, r, http.StatusCreated, err)
}
