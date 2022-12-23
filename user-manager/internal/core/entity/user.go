package entity

type UserFilters struct {
	Ids             []int `json:"ids"`
	TelegramUserIds []int `json:"telegramUserId"`

	Limit  int `json:"limit"`
	Offset int `json:"offset"`
}

type User struct {
	Id             int `json:"id"`
	TelegramUserId int `json:"telegramUserId"`
}
