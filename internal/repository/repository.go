package repository

import (
	"GabrielMSosa/crud-api/internal/domain"
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/go-sql-driver/mysql"
)

var (
	ErrRestrictFK    = errors.New("error Locality no exist")
	ErrGenericDriver = errors.New("error bad query")
	ErrNotFound      = errors.New("seller not found")
)

type Repository interface {
	GetAll(ctx context.Context) ([]domain.Seller, error)
	GetById(ctx context.Context, id int) (domain.Seller, error)
	Exists(ctx context.Context, cid int) bool
	Save(ctx context.Context, s domain.Seller) (int, error)
	Update(ctx context.Context, s domain.Seller) error
	Delete(ctx context.Context, id int) error
}

type repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) Repository {
	return &repository{
		db: db,
	}
}

// Retorna todos las filas de la db
func (r *repository) GetAll(ctx context.Context) ([]domain.Seller, error) {
	query := "SELECT * FROM sellers"
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	var sellers []domain.Seller
	for rows.Next() {
		s := domain.Seller{}
		_ = rows.Scan(&s.ID, &s.CID, &s.CompanyName, &s.Address, &s.Telephone, &s.LocalityId)
		sellers = append(sellers, s)
	}
	return sellers, nil
}

func (r *repository) GetById(ctx context.Context, id int) (domain.Seller, error) {
	query := "SELECT * FROM sellers WHERE id=?;"
	row := r.db.QueryRow(query, id)
	s := domain.Seller{}
	err := row.Scan(&s.ID, &s.CID, &s.CompanyName, &s.Address, &s.Telephone, &s.LocalityId)
	if err != nil {
		return domain.Seller{}, err
	}

	return s, nil
}
func (r *repository) Exists(ctx context.Context, cid int) bool {
	query := "SELECT cid FROM sellers WHERE cid=?;"
	row := r.db.QueryRow(query, cid)
	err := row.Scan(&cid)
	return err == nil
}
func (r *repository) Save(ctx context.Context, s domain.Seller) (int, error) {
	query := "INSERT INTO sellers(cid,company_name,address,telephone,locality_id) VALUES(?,?,?,?,?)"
	stmt, err := r.db.Prepare(query)
	if err != nil {
		return 0, err
	}
	defer stmt.Close()
	res, err := stmt.Exec(s.CID, s.CompanyName, s.Address, s.Telephone, s.LocalityId)
	//capturamos otros errores que SQL a veces no nos muestra
	driverErr, ok := err.(*mysql.MySQLError)
	if ok {
		fmt.Println("Error method save exec query:", driverErr.Number, " || ", driverErr.Message)
		switch {
		case driverErr.Number == 1452:
			return 0, ErrRestrictFK
		default:
			return 0, ErrGenericDriver
		}
	}
	if err != nil {
		return 0, err
	}
	id, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}
	return int(id), nil
}
func (r *repository) Update(ctx context.Context, s domain.Seller) error {
	query := "UPDATE sellers SET cid=?, company_name=?,address=?,telephone=?, locality_id=? WHERE id=?"
	stmt, err := r.db.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()
	res, err := stmt.Exec(s.CID, s.CompanyName, s.Address, s.Telephone, s.LocalityId, s.ID)
	//capturamos otros errores que SQL a veces no nos muestra
	driverErr, ok := err.(*mysql.MySQLError)
	if ok {
		fmt.Println("Error method save exec query:", driverErr.Number, " || ", driverErr.Message)
		switch {
		case driverErr.Number == 1452:
			return ErrRestrictFK
		default:
			return ErrGenericDriver
		}
	}
	if err != nil {
		return err
	}
	_, err = res.RowsAffected()
	if err != nil {
		return err
	}
	return nil
}
func (r *repository) Delete(ctx context.Context, id int) error {
	query := "DELETE FROM sellers WHERE id=?"
	stmt, err := r.db.Prepare(query)
	if err != nil {
		return err
	}
	res, err := stmt.Exec(id)
	if err != nil {
		return err
	}
	affect, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if affect < 1 {
		return ErrNotFound
	}
	return nil
}
