// Code generated by sqlc. DO NOT EDIT.

package querytest

import (
	"context"
	"database/sql"
	"fmt"
)

type DBTX interface {
	ExecContext(context.Context, string, ...interface{}) (sql.Result, error)
	PrepareContext(context.Context, string) (*sql.Stmt, error)
	QueryContext(context.Context, string, ...interface{}) (*sql.Rows, error)
	QueryRowContext(context.Context, string, ...interface{}) *sql.Row
}

func New(db DBTX) *Queries {
	return &Queries{db: db}
}

func Prepare(ctx context.Context, db DBTX) (*Queries, error) {
	q := Queries{db: db}
	var err error
	if q.deleteUsersByNameStmt, err = db.PrepareContext(ctx, deleteUsersByName); err != nil {
		return nil, fmt.Errorf("error preparing query DeleteUsersByName: %w", err)
	}
	if q.getUserByIDStmt, err = db.PrepareContext(ctx, getUserByID); err != nil {
		return nil, fmt.Errorf("error preparing query GetUserByID: %w", err)
	}
	if q.insertNewUserStmt, err = db.PrepareContext(ctx, insertNewUser); err != nil {
		return nil, fmt.Errorf("error preparing query InsertNewUser: %w", err)
	}
	if q.insertNewUserWithResultStmt, err = db.PrepareContext(ctx, insertNewUserWithResult); err != nil {
		return nil, fmt.Errorf("error preparing query InsertNewUserWithResult: %w", err)
	}
	if q.listUsersStmt, err = db.PrepareContext(ctx, listUsers); err != nil {
		return nil, fmt.Errorf("error preparing query ListUsers: %w", err)
	}
	return &q, nil
}

func (q *Queries) Close() error {
	var err error
	if q.deleteUsersByNameStmt != nil {
		if cerr := q.deleteUsersByNameStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing deleteUsersByNameStmt: %w", cerr)
		}
	}
	if q.getUserByIDStmt != nil {
		if cerr := q.getUserByIDStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getUserByIDStmt: %w", cerr)
		}
	}
	if q.insertNewUserStmt != nil {
		if cerr := q.insertNewUserStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing insertNewUserStmt: %w", cerr)
		}
	}
	if q.insertNewUserWithResultStmt != nil {
		if cerr := q.insertNewUserWithResultStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing insertNewUserWithResultStmt: %w", cerr)
		}
	}
	if q.listUsersStmt != nil {
		if cerr := q.listUsersStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing listUsersStmt: %w", cerr)
		}
	}
	return err
}

func (q *Queries) exec(ctx context.Context, stmt *sql.Stmt, query string, args ...interface{}) (sql.Result, error) {
	switch {
	case stmt != nil && q.tx != nil:
		return q.tx.StmtContext(ctx, stmt).ExecContext(ctx, args...)
	case stmt != nil:
		return stmt.ExecContext(ctx, args...)
	default:
		return q.db.ExecContext(ctx, query, args...)
	}
}

func (q *Queries) query(ctx context.Context, stmt *sql.Stmt, query string, args ...interface{}) (*sql.Rows, error) {
	switch {
	case stmt != nil && q.tx != nil:
		return q.tx.StmtContext(ctx, stmt).QueryContext(ctx, args...)
	case stmt != nil:
		return stmt.QueryContext(ctx, args...)
	default:
		return q.db.QueryContext(ctx, query, args...)
	}
}

func (q *Queries) queryRow(ctx context.Context, stmt *sql.Stmt, query string, args ...interface{}) *sql.Row {
	switch {
	case stmt != nil && q.tx != nil:
		return q.tx.StmtContext(ctx, stmt).QueryRowContext(ctx, args...)
	case stmt != nil:
		return stmt.QueryRowContext(ctx, args...)
	default:
		return q.db.QueryRowContext(ctx, query, args...)
	}
}

type Queries struct {
	db                          DBTX
	tx                          *sql.Tx
	deleteUsersByNameStmt       *sql.Stmt
	getUserByIDStmt             *sql.Stmt
	insertNewUserStmt           *sql.Stmt
	insertNewUserWithResultStmt *sql.Stmt
	listUsersStmt               *sql.Stmt
}

func (q *Queries) WithTx(tx *sql.Tx) *Queries {
	return &Queries{
		db:                          tx,
		tx:                          tx,
		deleteUsersByNameStmt:       q.deleteUsersByNameStmt,
		getUserByIDStmt:             q.getUserByIDStmt,
		insertNewUserStmt:           q.insertNewUserStmt,
		insertNewUserWithResultStmt: q.insertNewUserWithResultStmt,
		listUsersStmt:               q.listUsersStmt,
	}
}