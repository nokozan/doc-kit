package models

type User struct {
	ID    int    `json:"id"`
	Email string `json:"email"`
}

func (u *User) Validate() bool {
	return u.ID > 0 && u.Email != "" && (len(u.Email) > 5)
}

func (u *User) GetInfo() string {
	return "User ID: " + string(u.ID) + ", Email: " + u.Email
}
