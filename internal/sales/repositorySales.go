package sales

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/NahuelGScerca/HackthonGo/internal/models"
)

type Repository interface {
	Get(ctx context.Context, id int) (models.Sales, error)
	Store(ctx context.Context, entidad models.Sales) error
}

type repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) Repository {
	return &repository{
		db: db,
	}
}

func (r *repository) Get(ctx context.Context, id int) (models.Sales, error) {
	query := "SELECT * FROM sales WHERE id = ?;"
	row := r.db.QueryRow(query, id)
	b := models.Sales{}
	err := row.Scan(&b.ID, &b.IdInvoice, &b.IdProduct, &b.Quantity)
	if err != nil {
		return models.Sales{}, err
	}

	return b, nil
}

func (r *repository) Store(ctx context.Context, entidad models.Sales) error {
	// fmt.Println("ENTIDAD:  ", entidad)
	// return nil
	query := "INSERT INTO sales(id,idProduct,idinvoice,quantity) VALUES (?,?,?,?)"
	stmt, err := r.db.Prepare(query)
	if err != nil {
		return err
	}

	res, err := stmt.Exec(&entidad.ID, &entidad.IdProduct, &entidad.IdInvoice, &entidad.Quantity)
	if err != nil {
		return err
	}

	if num, err := res.RowsAffected(); num > 0 && err == nil {
		return nil
	}

	return fmt.Errorf("Error insert")
}

func (r *repository) Update(ctx context.Context, b models.Sales) error {
	query := "UPDATE sales SET idProduct=?, Idinvoice=? ,quantity=?  WHERE id=?"
	stmt, err := r.db.Prepare(query)
	if err != nil {
		return err
	}

	res, err := stmt.Exec(&b.IdProduct, &b.IdProduct, &b.Quantity, &b.ID)
	if err != nil {
		return err
	}

	_, err = res.RowsAffected()
	if err != nil {
		return err
	}

	return nil
}
