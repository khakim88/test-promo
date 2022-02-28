package postgres

import (
	"database/sql"
	"fmt"

	"github.com/khakim88/test-promo/repository"

	_ "github.com/lib/pq"

	"github.com/khakim88/test-promo/common/logger"
)

type postgresConn struct {
	db *sql.DB
}

const (
	driverName = "postgres"
)

func NewPostgresConn(conf repository.DBConfiguration) (repository.DBReaderWriter, error) {
	connURL := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=disable",
		conf.DBUser,
		conf.DBPassword,
		conf.DBHost,
		conf.DBPort,
		conf.DBName)

	logger.Info(fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s",
		conf.DBUser,
		"**************",
		conf.DBHost,
		conf.DBPort,
		conf.DBName))

	db, err := sql.Open(driverName, connURL)
	if err != nil {
		logger.Error(err)
	}

	if conf.MaxConnection > 0 {
		db.SetMaxOpenConns(conf.MaxConnection)
	}
	if conf.MaxIdleConnection > 0 {
		db.SetMaxIdleConns(conf.MaxIdleConnection)
	}

	if err := db.Ping(); err != nil {
		logger.Error(err)
	}

	logger.Info("DB Says Pong!, DB connected")

	return &postgresConn{
		db: db,
	}, nil
}

func (p *postgresConn) Close() error {
	return p.db.Close()
}

// Transact ...
type Transact interface {
	Exec(query string, args ...interface{}) (sql.Result, error)
	Prepare(query string) (*sql.Stmt, error)
	Query(query string, args ...interface{}) (*sql.Rows, error)
	QueryRow(query string, args ...interface{}) *sql.Row
}

// TxCallback ...
type TxCallback func(Transact) error

// WithTransact ...
func WithTransact(db *sql.DB, callback TxCallback) (err error) {
	tx, err := db.Begin()
	if err != nil {
		return fmt.Errorf("transact begin: %v", err)
	}

	defer func() {
		if p := recover(); p != nil {
			tx.Rollback()
			panic(p)
		} else if err != nil {
			tx.Rollback()
		} else {
			err = tx.Commit()
		}
	}()

	err = callback(tx)
	if err != nil {
		return fmt.Errorf("error transact: %v", err)
	}
	return nil
}
