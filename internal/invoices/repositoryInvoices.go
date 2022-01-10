package invoices

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/NahuelGScerca/HackthonGo/internal/models"
)

type Repository interface {
	Store(ctx context.Context, entidad models.Invoices) error
	Get(ctx context.Context, id int) (models.Invoices, error)
}

type repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) Repository {
	return &repository{
		db: db,
	}
}

func (r *repository) Get(ctx context.Context, id int) (models.Invoices, error) {
	query := "SELECT * FROM invoices WHERE id = ?;"
	row := r.db.QueryRow(query, id)
	b := models.Invoices{}
	err := row.Scan(&b.ID, &b.Datetime, &b.IdCustomer, &b.Total)
	if err != nil {
		return models.Invoices{}, err
	}

	return b, nil
}

func (r *repository) Store(ctx context.Context, entidad models.Invoices) error {
	// fmt.Println("ENTIDAD:  ", entidad)
	// return nil
	query := "INSERT INTO invoices(id,datetime,idCustomer,total) VALUES (?,?,?,?)"
	stmt, err := r.db.Prepare(query)
	if err != nil {
		return err
	}

	res, err := stmt.Exec(&entidad.ID, &entidad.Datetime, &entidad.IdCustomer, &entidad.Total)
	if err != nil {
		return err
	}

	if num, err := res.RowsAffected(); num > 0 && err == nil {
		defer stmt.Close()

		return nil
	}

	return fmt.Errorf("Error insert")
}
