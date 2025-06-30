package repository

import (
	"context"
	"database/sql"
	"errors"

	"api.fiber.practice/models"
	_ "github.com/lib/pq"
)

type NameSQL struct {
	DB *sql.DB
}

func (q *NameSQL) DBGetAllName() ([]models.FullNameRet, error) {
	ret := []models.FullNameRet{}

	// Query Get All Names
	query := `
	SELECT
		n.name_id,
		n.first_name,
		COALESCE(n.middle_name, '') AS middle_name,
		COALESCE(n.last_name, '') AS last_name
	FROM public.full_name n `

	rows, err := q.DB.Query(query)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		data := models.FullNameRet{}
		err = rows.Scan(&data.NameID, &data.FirstName, &data.MiddleName, &data.LastName)

		if err != nil {
			return nil, err
		}

		ret = append(ret, data)
	}

	return ret, nil
}

func (q *NameSQL) DBGetFullName(reqID int) ([]models.FullNameRet, error) {
	ret := []models.FullNameRet{}

	// Query Get Name by Id
	query := `
	SELECT
		n.name_id,
		n.first_name,
		COALESCE(n.middle_name, '') AS middle_name,
		COALESCE(n.last_name, '') AS last_name
	FROM public.full_name n `

	query = query + " WHERE n.name_id=$1"
	data := models.FullNameRet{}
	err := q.DB.QueryRow(query, reqID).Scan(&data.NameID, &data.FirstName, &data.MiddleName, &data.LastName)

	if err == sql.ErrNoRows {
		return nil, errors.New("data not found")
	} else if err != nil {
		return nil, err
	}

	ret = append(ret, data)

	return ret, nil
}

func (q *NameSQL) DBCreateName(req *models.FullNameReq, reqID int) error {
	ctx := context.Background()
	tx, err := q.DB.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	// Query Insert Name
	Q := `
	INSERT INTO public.full_name
		(name_id, first_name, middle_name, last_name)
	VALUES
		($1, $2, $3, $4); 
	`

	_, err = tx.Exec(Q, reqID, req.FirstName, req.MiddleName, req.LastName)
	if err != nil {
		tx.Rollback()
		return err
	}

	err = tx.Commit()
	if err != nil {
		return err
	} else {
		return nil
	}
}

func (q *NameSQL) DBUpdateName(req *models.FullNameReq, reqID int) error {
	ctx := context.Background()
	tx, err := q.DB.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	// Query Insert Name
	query := `
	UPDATE public.full_name
	SET
		first_name=$2, 
		middle_name=$3, 
		last_name=$4
	WHERE name_id=$1; 
	`
	_, err = tx.Exec(query, reqID, req.FirstName, req.MiddleName, req.LastName)
	if err == sql.ErrNoRows {
		tx.Rollback()
		return errors.New("data not found")
	} else if err != nil {
		tx.Rollback()
		return err
	}

	return nil
}
