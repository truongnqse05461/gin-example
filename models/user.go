package models

type User struct {
	ID         string `json:"id"`
	Name       string `json:"name"`
	Email      string `json:"email"`
	Phone      string `json:"phone"`
	Password   string `json:"password"`
	LastUpdate int64
}

func (u *User) BuildInsetArgs() []interface{} {
	return []interface{}{u.ID, u.Name, u.Email, u.Phone, u.Password, u.LastUpdate}
}
