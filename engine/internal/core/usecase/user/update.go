package user

import "log"

func (usecase *UserUseCase) Update(login, groupName string) *int {
	scheduleRepo := usecase.repository.GetScheduleRepository()
	userRepo := usecase.repository.GetUserRepository()

	groupId := scheduleRepo.GetGroupId(groupName)

	if groupId == nil {
		return groupId
	}

	err := userRepo.Update(login, *groupId)

	if err != nil {
		log.Println(err)
	}

	return groupId
}
