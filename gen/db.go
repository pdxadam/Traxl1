// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0

package dbTraxl

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
	if q.createUsersStmt, err = db.PrepareContext(ctx, createUsers); err != nil {
		return nil, fmt.Errorf("error preparing query CreateUsers: %w", err)
	}
	if q.deleteInstanceStmt, err = db.PrepareContext(ctx, deleteInstance); err != nil {
		return nil, fmt.Errorf("error preparing query DeleteInstance: %w", err)
	}
	if q.deleteTopicStmt, err = db.PrepareContext(ctx, deleteTopic); err != nil {
		return nil, fmt.Errorf("error preparing query DeleteTopic: %w", err)
	}
	if q.deleteTopicInstancesStmt, err = db.PrepareContext(ctx, deleteTopicInstances); err != nil {
		return nil, fmt.Errorf("error preparing query DeleteTopicInstances: %w", err)
	}
	if q.deleteUserInstancesStmt, err = db.PrepareContext(ctx, deleteUserInstances); err != nil {
		return nil, fmt.Errorf("error preparing query DeleteUserInstances: %w", err)
	}
	if q.deleteUserTopicsStmt, err = db.PrepareContext(ctx, deleteUserTopics); err != nil {
		return nil, fmt.Errorf("error preparing query DeleteUserTopics: %w", err)
	}
	if q.deleteUsersStmt, err = db.PrepareContext(ctx, deleteUsers); err != nil {
		return nil, fmt.Errorf("error preparing query DeleteUsers: %w", err)
	}
	if q.getUserStmt, err = db.PrepareContext(ctx, getUser); err != nil {
		return nil, fmt.Errorf("error preparing query GetUser: %w", err)
	}
	if q.getUserByNameStmt, err = db.PrepareContext(ctx, getUserByName); err != nil {
		return nil, fmt.Errorf("error preparing query GetUserByName: %w", err)
	}
	if q.getUserInstancesStmt, err = db.PrepareContext(ctx, getUserInstances); err != nil {
		return nil, fmt.Errorf("error preparing query GetUserInstances: %w", err)
	}
	if q.getUserTopicsStmt, err = db.PrepareContext(ctx, getUserTopics); err != nil {
		return nil, fmt.Errorf("error preparing query GetUserTopics: %w", err)
	}
	if q.insertInstanceStmt, err = db.PrepareContext(ctx, insertInstance); err != nil {
		return nil, fmt.Errorf("error preparing query InsertInstance: %w", err)
	}
	if q.insertTopicStmt, err = db.PrepareContext(ctx, insertTopic); err != nil {
		return nil, fmt.Errorf("error preparing query InsertTopic: %w", err)
	}
	if q.listInstancesStmt, err = db.PrepareContext(ctx, listInstances); err != nil {
		return nil, fmt.Errorf("error preparing query ListInstances: %w", err)
	}
	if q.listTopicsStmt, err = db.PrepareContext(ctx, listTopics); err != nil {
		return nil, fmt.Errorf("error preparing query ListTopics: %w", err)
	}
	if q.listUsersStmt, err = db.PrepareContext(ctx, listUsers); err != nil {
		return nil, fmt.Errorf("error preparing query ListUsers: %w", err)
	}
	if q.updateInstanceStmt, err = db.PrepareContext(ctx, updateInstance); err != nil {
		return nil, fmt.Errorf("error preparing query UpdateInstance: %w", err)
	}
	if q.updateTopicStmt, err = db.PrepareContext(ctx, updateTopic); err != nil {
		return nil, fmt.Errorf("error preparing query UpdateTopic: %w", err)
	}
	return &q, nil
}

func (q *Queries) Close() error {
	var err error
	if q.createUsersStmt != nil {
		if cerr := q.createUsersStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing createUsersStmt: %w", cerr)
		}
	}
	if q.deleteInstanceStmt != nil {
		if cerr := q.deleteInstanceStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing deleteInstanceStmt: %w", cerr)
		}
	}
	if q.deleteTopicStmt != nil {
		if cerr := q.deleteTopicStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing deleteTopicStmt: %w", cerr)
		}
	}
	if q.deleteTopicInstancesStmt != nil {
		if cerr := q.deleteTopicInstancesStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing deleteTopicInstancesStmt: %w", cerr)
		}
	}
	if q.deleteUserInstancesStmt != nil {
		if cerr := q.deleteUserInstancesStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing deleteUserInstancesStmt: %w", cerr)
		}
	}
	if q.deleteUserTopicsStmt != nil {
		if cerr := q.deleteUserTopicsStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing deleteUserTopicsStmt: %w", cerr)
		}
	}
	if q.deleteUsersStmt != nil {
		if cerr := q.deleteUsersStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing deleteUsersStmt: %w", cerr)
		}
	}
	if q.getUserStmt != nil {
		if cerr := q.getUserStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getUserStmt: %w", cerr)
		}
	}
	if q.getUserByNameStmt != nil {
		if cerr := q.getUserByNameStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getUserByNameStmt: %w", cerr)
		}
	}
	if q.getUserInstancesStmt != nil {
		if cerr := q.getUserInstancesStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getUserInstancesStmt: %w", cerr)
		}
	}
	if q.getUserTopicsStmt != nil {
		if cerr := q.getUserTopicsStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getUserTopicsStmt: %w", cerr)
		}
	}
	if q.insertInstanceStmt != nil {
		if cerr := q.insertInstanceStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing insertInstanceStmt: %w", cerr)
		}
	}
	if q.insertTopicStmt != nil {
		if cerr := q.insertTopicStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing insertTopicStmt: %w", cerr)
		}
	}
	if q.listInstancesStmt != nil {
		if cerr := q.listInstancesStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing listInstancesStmt: %w", cerr)
		}
	}
	if q.listTopicsStmt != nil {
		if cerr := q.listTopicsStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing listTopicsStmt: %w", cerr)
		}
	}
	if q.listUsersStmt != nil {
		if cerr := q.listUsersStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing listUsersStmt: %w", cerr)
		}
	}
	if q.updateInstanceStmt != nil {
		if cerr := q.updateInstanceStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing updateInstanceStmt: %w", cerr)
		}
	}
	if q.updateTopicStmt != nil {
		if cerr := q.updateTopicStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing updateTopicStmt: %w", cerr)
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
	db                       DBTX
	tx                       *sql.Tx
	createUsersStmt          *sql.Stmt
	deleteInstanceStmt       *sql.Stmt
	deleteTopicStmt          *sql.Stmt
	deleteTopicInstancesStmt *sql.Stmt
	deleteUserInstancesStmt  *sql.Stmt
	deleteUserTopicsStmt     *sql.Stmt
	deleteUsersStmt          *sql.Stmt
	getUserStmt              *sql.Stmt
	getUserByNameStmt        *sql.Stmt
	getUserInstancesStmt     *sql.Stmt
	getUserTopicsStmt        *sql.Stmt
	insertInstanceStmt       *sql.Stmt
	insertTopicStmt          *sql.Stmt
	listInstancesStmt        *sql.Stmt
	listTopicsStmt           *sql.Stmt
	listUsersStmt            *sql.Stmt
	updateInstanceStmt       *sql.Stmt
	updateTopicStmt          *sql.Stmt
}

func (q *Queries) WithTx(tx *sql.Tx) *Queries {
	return &Queries{
		db:                       tx,
		tx:                       tx,
		createUsersStmt:          q.createUsersStmt,
		deleteInstanceStmt:       q.deleteInstanceStmt,
		deleteTopicStmt:          q.deleteTopicStmt,
		deleteTopicInstancesStmt: q.deleteTopicInstancesStmt,
		deleteUserInstancesStmt:  q.deleteUserInstancesStmt,
		deleteUserTopicsStmt:     q.deleteUserTopicsStmt,
		deleteUsersStmt:          q.deleteUsersStmt,
		getUserStmt:              q.getUserStmt,
		getUserByNameStmt:        q.getUserByNameStmt,
		getUserInstancesStmt:     q.getUserInstancesStmt,
		getUserTopicsStmt:        q.getUserTopicsStmt,
		insertInstanceStmt:       q.insertInstanceStmt,
		insertTopicStmt:          q.insertTopicStmt,
		listInstancesStmt:        q.listInstancesStmt,
		listTopicsStmt:           q.listTopicsStmt,
		listUsersStmt:            q.listUsersStmt,
		updateInstanceStmt:       q.updateInstanceStmt,
		updateTopicStmt:          q.updateTopicStmt,
	}
}
