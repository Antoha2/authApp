package repository

import (
	"context"
	"log"

	models "github.com/Antoha2/auth"
	//"github.com/mattn/go-sqlite3"
)

// SaveUser saves user to db.
func (r *RepAuth) SaveUser(ctx context.Context, email string, passHash []byte) (int64, error) {
	const op = "storage.sqlite.SaveUser"

	log.Println("Простенький зaпрос на добавление пользователя")
	/*
		// Простенький зaпрос на добавление пользователя
		stmt, err := s.db.Prepare("INSERT INTO users(email, pass_hash) VALUES(?, ?)")

		if err != nil {
			return 0, fmt.Errorf("%s: %w", op, err)
		}

		// Выполняем запрос, передав параметры
		res, err := stmt.ExecContext(ctx, email, passHash)
		if err != nil {
			var sqliteErr sqlite3.Error

			// Небольшое кунг-фу для выявления ошибки ErrConstraintUnique
			// (см. подробности ниже)
			if errors.As(err, &sqliteErr) && sqliteErr.ExtendedCode == sqlite3.ErrConstraintUnique {
				return 0, fmt.Errorf("%s: %w", op, storage.ErrUserExists)
			}

			return 0, fmt.Errorf("%s: %w", op, err)
		}

		// Получаем ID созданной записи
		id, err := res.LastInsertId()
		if err != nil {
			return 0, fmt.Errorf("%s: %w", op, err)
		}
	*/
	return 1, nil
}

// User returns user by email.
func (r *RepAuth) User(ctx context.Context, email string) (models.User, error) {
	const op = "storage.sqlite.User"
	user := new(models.User)
	log.Println("Простенький зaпрос на поиск пользователя")
	/*
	   stmt, err := s.db.Prepare("SELECT id, email, pass_hash FROM users WHERE email = ?")
	   if err != nil {
	       return models.User{}, fmt.Errorf("%s: %w", op, err)
	   }

	   row := stmt.QueryRowContext(ctx, email)

	   var user models.User
	   err = row.Scan(&user.ID, &user.Email, &user.PassHash)
	   if err != nil {
	       if errors.Is(err, sql.ErrNoRows) {
	           return models.User{}, fmt.Errorf("%s: %w", op, storage.ErrUserNotFound)
	       }

	       return models.User{}, fmt.Errorf("%s: %w", op, err)
	   }
	*/
	return *user, nil
}

// App returns app by id.
func (r *RepAuth) App(ctx context.Context, id int) (models.App, error) {
	const op = "storage.sqlite.App"
	app := new(models.App)
	log.Println("Простенький зaпрос на поиск app")
	/*
	   stmt, err := s.db.Prepare("SELECT id, name, secret FROM apps WHERE id = ?")
	   if err != nil {
	       return models.App{}, fmt.Errorf("%s: %w", op, err)
	   }

	   row := stmt.QueryRowContext(ctx, id)

	   var app models.App
	   err = row.Scan(&app.ID, &app.Name, &app.Secret)
	   if err != nil {
	       if errors.Is(err, sql.ErrNoRows) {
	           return models.App{}, fmt.Errorf("%s: %w", op, storage.ErrAppNotFound)
	       }

	       return models.App{}, fmt.Errorf("%s: %w", op, err)
	   }
	*/
	return *app, nil
}
