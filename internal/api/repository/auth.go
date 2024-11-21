package repository

type User struct {
	Id       int    `json:"id"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
	Role     string `json:"role"`
}

func (r *Repository) GetUserByEmail(email string) (*User, error) {
	u := new(User)

	res := r.database.Conn.QueryRow("SELECT * FROM user WHERE email = ?", email)

	err := res.Scan(&u.Id, &u.Email, &u.Password, &u.Role)

	if err != nil {
		return nil, err
	}

	return u, nil
}

func (r *Repository) SaveUser(user struct {
	email    string
	password string
	role     string
}) (*User, error) {
	insert, err := r.database.Conn.Query("INSERT INTO user (email, password, role) VALUES (?, ?, ?)", user.email, user.password, user.role)

	if err != nil {
		return nil, err
	}

	defer insert.Close()

	return &User{}, nil
}
