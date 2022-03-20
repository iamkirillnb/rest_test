package repos

import (
	"github.com/iamkirillnb/rus_law/internal/entities"
	"github.com/iamkirillnb/rus_law/pkg/logging"
	"github.com/jmoiron/sqlx"
)

type DbLaw struct {
	DB     *sqlx.DB
	Logger *logging.Logger
}

func NewDbLaw(db *sqlx.DB, logger *logging.Logger) DbLaw {
	return DbLaw{
		DB:     db,
		Logger: logger,
	}
}

func (d *DbLaw) GetAll() ([]*entities.FormData,error) {
	const qry = `select * from data`

	data := []*entities.FormData{}
	err := d.DB.Select(&data, qry)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (d *DbLaw) WriteData(data *entities.FormData) error {
	const qry = `insert into data (name, email, question) values (:name, :email, :question)`

	_, err := d.DB.NamedExec(qry, data)
	if err != nil {
		return err
	}
	return nil
}
