package repository

import (
	"database/sql"

	"api.fiber.practice/models"
)

type NameSQL struct {
	DB *sql.DB
}

func (m NameSQL) DBGetAllName(req models.FullNameReq) ([]models.FullNameRet, error) {
	ret := []models.FullNameRet{}

	// Query Get All Names
	Q := `
	SELECT
		first_name,
		middle_name,
		last_name
	FROM full_name n
	WHERE 1=1 `

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

	if err != nil {
		return nil, err
	}

	return ret, nil
}

func (m NameSQL) DBGetFullName(reqID int) ([]models.FullNameRet, error) {
	return nil, nil
}
