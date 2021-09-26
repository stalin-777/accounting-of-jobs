package postgres

import (
	"context"
	stderrors "errors"

	"github.com/jackc/pgconn"
	"github.com/jackc/pgerrcode"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/juju/errors"
	aoj "github.com/stalin-777/accounting-of-jobs"
)

type WorkplaceService struct {
	DB *pgxpool.Pool
}

const (
	dbFindWorkplace = `
	SELECT 
		id,
		username,
		hostname,
		ip
	FROM public.workplace
	WHERE id = $1`

	dbFindWorkplaces = `
	SELECT 
		id,
		username,
		hostname,
		ip
	FROM public.workplace`

	dbCreateWorkplace = `
	INSERT INTO public.workplace (
		username,
		hostname,
		ip
	) 
	VALUES ($1, $2, $3)
	RETURNING id`

	dbUpdateWorkplace = `
	UPDATE public.workplace
	SET 
		username = $1,
		hostname = $2,
		ip = $3
	`

	dbDeleteWorkplace = `
	DELETE FROM public.workplace
	WHERE id = $1
	RETURNING id;`
)

var ErrConstraintPgx = errors.New("Пользователь с таким именем уже существует. Обновить данные пользователя? Y/N")

func (s *WorkplaceService) Workplace(id int) (*aoj.Workplace, error) {

	var w aoj.Workplace

	err := s.DB.QueryRow(context.Background(), dbFindWorkplace, id).Scan(
		&w.ID,
		&w.Username,
		&w.Hostname,
		&w.IP,
	)
	if err != nil {
		return nil, err
	}

	return &w, nil
}

func (s *WorkplaceService) Workplaces() ([]*aoj.Workplace, error) {

	rows, err := s.DB.Query(context.Background(), dbFindWorkplaces)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var result []*aoj.Workplace

	for rows.Next() {

		var w aoj.Workplace

		rows.Scan(
			&w.ID,
			&w.Username,
			&w.Hostname,
			&w.IP,
		)
		if err != nil {
			return nil, err
		}

		result = append(result, &w)
	}

	if len(result) == 0 {
		return nil, errors.New("Пустой список рабочих мест")
	}

	return result, nil
}

func (s *WorkplaceService) CreateWorkplace(w *aoj.Workplace) error {

	err := s.DB.QueryRow(context.Background(), dbCreateWorkplace,
		w.Username,
		w.Hostname,
		w.IP,
	).Scan(&w.ID)
	if err != nil {

		var pgErr *pgconn.PgError
		if stderrors.As(err, &pgErr) {
			if pgErr.Code == pgerrcode.UniqueViolation && pgErr.ConstraintName == "workplace_username_key" {
				return ErrConstraintPgx
			}
		}
		return err
	}

	return nil
}

func (s *WorkplaceService) UpdateWorkplace(w *aoj.Workplace) error {

	returning := ` RETURNING id`
	sql := dbUpdateWorkplace
	sqlArgs := []interface{}{
		w.Username,
		w.Hostname,
		w.IP,
	}
	if w.ID == 0 {
		sql += " WHERE username = $4" + returning
		sqlArgs = append(sqlArgs, w.Username)
	} else {
		sql += " WHERE id = $4" + returning
		sqlArgs = append(sqlArgs, w.ID)
	}

	err := s.DB.QueryRow(context.Background(), sql, sqlArgs...).Scan(&w.ID)
	if err != nil {
		return err
	}

	return nil
}

func (s *WorkplaceService) DeleteWorkplace(id int) error {

	//Протестировать и заменить на Exec
	err := s.DB.QueryRow(context.Background(), dbDeleteWorkplace, id).Scan(&id)
	if err != nil {
		return err
	}

	return nil
}
