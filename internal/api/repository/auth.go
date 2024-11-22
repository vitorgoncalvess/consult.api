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

func (r *Repository) SaveUser(user User) (*User, error) {
	insert, err := r.database.Conn.Exec("INSERT INTO user (email, password, role) VALUES (?, ?, ?)", user.Email, user.Password, user.Role)

	if err != nil {
		return nil, err
	}

	id, err := insert.LastInsertId()

	if err != nil {
		return nil, err
	}

	user.Id = int(id)

	return &user, nil
}
