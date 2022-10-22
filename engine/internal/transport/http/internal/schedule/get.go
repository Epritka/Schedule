package schedule

import (
	httpassistant "engine/infrastructure/http-assistant"
	"net/http"
	"strconv"
	"time"
)

//	const days = {
//	    1: "Понедельник",
//	    2: "Вторник",
//	    3: "Среда",
//	    4: "Четверг",
//	    5: "Пятница",
//	    6: "Суббота",
//		7: "Воскресенье"
//	}
type dayRequest struct {
	GroupId int    `json:"groupId"`
	Date    string `json:"date"`
}

func (h *handlers) GetScheduleForDay(w http.ResponseWriter, r *http.Request) {
	dr := dayRequest{}

	err := httpassistant.ParseBody(r.Body, &dr)
	if err != nil {
		httpassistant.Reply(w, r, http.StatusBadRequest, err)
		return
	}

	i, err := strconv.ParseInt(dr.Date, 10, 64)
	if err != nil {
		httpassistant.Reply(w, r, http.StatusBadRequest, err)
		return
	}

	day, err := h.scheduleUseCase.GetDay(time.Unix(i, 0), dr.GroupId)
	if err != nil {
		httpassistant.Reply(w, r, http.StatusInternalServerError, err)
		return
	}

	httpassistant.Reply(w, r, http.StatusOK, &day)
}
