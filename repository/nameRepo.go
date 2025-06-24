package repository

import (
	"database/sql"
	"errors"
	"strconv"

	"api.fiber.practice/models"
	_ "github.com/lib/pq"
)

type NameSQL struct {
	DB *sql.DB
}

func (m NameSQL) DBGetAllName() ([]models.FullNameRet, error) {
	ret := []models.FullNameRet{}

	// Query Get All Names
	Q := `
	SELECT
		n.first_name,
		COALESCE(n.middle_name, '') AS middle_name,
		COALESCE(n.last_name, '') AS last_name
	FROM public.full_name n `

	rows, err := m.DB.Query(Q)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		data := models.FullNameRet{}
		err = rows.Scan(&data.FirstName, &data.MiddleName, &data.LastName)

		if err != nil {
			return nil, err
		}

		ret = append(ret, data)
	}

	return ret, nil
}

func (m NameSQL) DBGetFullName(reqID int) ([]models.FullNameRet, error) {
	ret := []models.FullNameRet{}

	// Query Get Name by Id
	Q := `
	SELECT
		n.first_name,
		COALESCE(n.middle_name, '') AS middle_name,
		COALESCE(n.last_name, '') AS last_name
	FROM public.full_name n `

	Q = Q + " WHERE n.name_id=" + strconv.Itoa(reqID)
	rows, err := m.DB.Query(Q)
	if err != nil {
		return nil, errors.New(Q + "error: " + err.Error())
	}

	for rows.Next() {
		data := models.FullNameRet{}
		err = rows.Scan(&data.FirstName, &data.MiddleName, &data.LastName)

		if err != nil {
			return nil, errors.New("error receiving data: " + err.Error())
		}

		ret = append(ret, data)
	}

	return ret, nil
}
