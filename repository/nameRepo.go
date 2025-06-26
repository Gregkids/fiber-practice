package repository

import (
	"context"
	"database/sql"
	"strconv"

	"api.fiber.practice/models"
	_ "github.com/lib/pq"
)

type NameSQL struct {
	DB *sql.DB
}

func (sql NameSQL) DBGetAllName() ([]models.FullNameRet, error) {
	ret := []models.FullNameRet{}

	// Query Get All Names
	Q := `
	SELECT
		n.name_id,
		n.first_name,
		COALESCE(n.middle_name, '') AS middle_name,
		COALESCE(n.last_name, '') AS last_name
	FROM public.full_name n `

	rows, err := sql.DB.Query(Q)
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

func (sql NameSQL) DBGetFullName(reqID int) ([]models.FullNameRet, error) {
	ret := []models.FullNameRet{}

	// Query Get Name by Id
	Q := `
	SELECT
		n.name_id,
		n.first_name,
		COALESCE(n.middle_name, '') AS middle_name,
		COALESCE(n.last_name, '') AS last_name
	FROM public.full_name n `

	Q = Q + " WHERE n.name_id=" + strconv.Itoa(reqID)
	row := sql.DB.QueryRow(Q, reqID)
	data := models.FullNameRet{}
	err := row.Scan(&data.NameID, &data.FirstName, &data.MiddleName, &data.LastName)

	if err != nil {
		return nil, err
	}

	ret = append(ret, data)

	return ret, nil
}

func (sql NameSQL) DBCreateName(req *models.FullNameReq, reqID int) error {
	ctx := context.Background()
	tx, err := sql.DB.BeginTx(ctx, nil)
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
