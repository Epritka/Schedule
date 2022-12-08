package entity

import "unicode"

type UserFilters struct {
	Ids      []int  `json:"ids"`
	Email    string `json:"email"`
	FullName string `json:"fullName"`

	Limit  int `json:"limit"`
	Offset int `json:"offset"`
}

type User struct {
	Id           int          `json:"id"`
	AuthSourceId int          `json:"authSourceId"`
	Email        string       `json:"email"`
	FirstName    string       `json:"firstName"`
	LastName     string       `json:"lastName"`
	Password     string       `json:"password"`
	IsSuperuser  bool         `json:"isSuperuser"`
	IsActive     bool         `json:"isActive"`
	Roles        []Role       `json:"roles"`
	Permissions  []Permission `json:"permissions"`
	AuthSource   AuthSource   `json:"authSource"`
}

func (user *User) CheckPermission(permission Permission) bool {
	if user.IsSuperuser {
		return true
	}

	for _, perm := range user.GetPermissions() {
		if perm == permission {
			return true
		}
	}
	return false
}

func (user *User) GetPermissions() []Permission {
	list := []Permission{}
	keys := make(map[Permission]bool)

	for _, role := range user.Roles {
		for _, perm := range role.Permissions {
			if _, value := keys[perm]; !value {
				keys[perm] = true
				list = append(list, perm)
			}
		}
	}

	return append(list, user.Permissions...)
}

func (user *User) ValidPassword(password string) bool {
	var (
		hasMinLen = false
		hasUpper  = false
		hasLower  = false
		hasNumber = false
	)
	if len(password) >= 7 {
		hasMinLen = true
	}
	for _, char := range password {
		switch {
		case unicode.IsUpper(char):
			hasUpper = true
		case unicode.IsLower(char):
			hasLower = true
		case unicode.IsNumber(char):
			hasNumber = true
		}
	}
	return hasMinLen && hasUpper && hasLower && hasNumber
}
