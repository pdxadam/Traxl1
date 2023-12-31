// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0
// source: query.sql

package dbTraxl

import (
	"context"
	"time"
)

const createUsers = `-- name: CreateUsers :one
INSERT INTO dbTraxl.users (userName, PasswordHash, name )
VALUES ($1,
        $2,
        $3) RETURNING pkuser, username, passwordhash, name, config, createdts, isenabled
`

type CreateUsersParams struct {
	Username     string `db:"username" json:"username"`
	Passwordhash string `db:"passwordhash" json:"passwordhash"`
	Name         string `db:"name" json:"name"`
}

// insert new user
func (q *Queries) CreateUsers(ctx context.Context, arg CreateUsersParams) (DbtraxlUser, error) {
	row := q.queryRow(ctx, q.createUsersStmt, createUsers, arg.Username, arg.Passwordhash, arg.Name)
	var i DbtraxlUser
	err := row.Scan(
		&i.Pkuser,
		&i.Username,
		&i.Passwordhash,
		&i.Name,
		&i.Config,
		&i.Createdts,
		&i.Isenabled,
	)
	return i, err
}

const deleteInstance = `-- name: DeleteInstance :exec
DELETE
FROM dbTraxl.instances
WHERE PkInstance = $1
`

// delete a particular instance
func (q *Queries) DeleteInstance(ctx context.Context, pkinstance int64) error {
	_, err := q.exec(ctx, q.deleteInstanceStmt, deleteInstance, pkinstance)
	return err
}

const deleteTopic = `-- name: DeleteTopic :exec
DELETE
FROM dbTraxl.topics
WHERE PkTopic = $1
`

// Delete a particular topic
func (q *Queries) DeleteTopic(ctx context.Context, pktopic int64) error {
	_, err := q.exec(ctx, q.deleteTopicStmt, deleteTopic, pktopic)
	return err
}

const deleteTopicInstances = `-- name: DeleteTopicInstances :exec
DELETE 
FROM dbTraxl.instances
WHERE FkTopic = $1
`

// delete the instances for a particular topic
func (q *Queries) DeleteTopicInstances(ctx context.Context, fktopic int64) error {
	_, err := q.exec(ctx, q.deleteTopicInstancesStmt, deleteTopicInstances, fktopic)
	return err
}

const deleteUserInstances = `-- name: DeleteUserInstances :exec
DELETE 
FROM dbTraxl.instances
WHERE FkUser = $1
`

// delete all instances for a particular user
func (q *Queries) DeleteUserInstances(ctx context.Context, fkuser int64) error {
	_, err := q.exec(ctx, q.deleteUserInstancesStmt, deleteUserInstances, fkuser)
	return err
}

const deleteUserTopics = `-- name: DeleteUserTopics :exec
DELETE 
FROM dbTraxl.topics
WHERE FkUser = $1
`

// delete a user's topics
func (q *Queries) DeleteUserTopics(ctx context.Context, fkuser int64) error {
	_, err := q.exec(ctx, q.deleteUserTopicsStmt, deleteUserTopics, fkuser)
	return err
}

const deleteUsers = `-- name: DeleteUsers :exec
DELETE
FROM dbTraxl.users
WHERE PkUser = $1
`

// delete a particular user
func (q *Queries) DeleteUsers(ctx context.Context, pkuser int64) error {
	_, err := q.exec(ctx, q.deleteUsersStmt, deleteUsers, pkuser)
	return err
}

const getUser = `-- name: GetUser :one
SELECT pkuser, username, passwordhash, name, config, createdts, isenabled
FROM dbtraxl.users
WHERE PkUser = $1
`

// get user of a particular userid
func (q *Queries) GetUser(ctx context.Context, pkuser int64) (DbtraxlUser, error) {
	row := q.queryRow(ctx, q.getUserStmt, getUser, pkuser)
	var i DbtraxlUser
	err := row.Scan(
		&i.Pkuser,
		&i.Username,
		&i.Passwordhash,
		&i.Name,
		&i.Config,
		&i.Createdts,
		&i.Isenabled,
	)
	return i, err
}

const getUserByName = `-- name: GetUserByName :one
SELECT pkuser, username, passwordhash, name, config, createdts, isenabled
FROM dbTraxl.users
WHERE username = $1
`

// get a particular user by username
func (q *Queries) GetUserByName(ctx context.Context, username string) (DbtraxlUser, error) {
	row := q.queryRow(ctx, q.getUserByNameStmt, getUserByName, username)
	var i DbtraxlUser
	err := row.Scan(
		&i.Pkuser,
		&i.Username,
		&i.Passwordhash,
		&i.Name,
		&i.Config,
		&i.Createdts,
		&i.Isenabled,
	)
	return i, err
}

const getUserInstances = `-- name: GetUserInstances :many
SELECT u.PkUser, t.PkTopic, t.TopicName, i.PkInstance, i.Start_Date
FROM dbTraxl.users u, dbTraxl.topics t, dbTraxl.instances i
WHERE u.PkUser = t.FkUser
AND t.PkTopic = i.FkTopic
AND u.PkUser = $1
`

type GetUserInstancesRow struct {
	Pkuser     int64     `db:"pkuser" json:"pkuser"`
	Pktopic    int64     `db:"pktopic" json:"pktopic"`
	Topicname  string    `db:"topicname" json:"topicname"`
	Pkinstance int64     `db:"pkinstance" json:"pkinstance"`
	StartDate  time.Time `db:"start_date" json:"startDate"`
}

// get all the instances
func (q *Queries) GetUserInstances(ctx context.Context, pkuser int64) ([]GetUserInstancesRow, error) {
	rows, err := q.query(ctx, q.getUserInstancesStmt, getUserInstances, pkuser)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetUserInstancesRow
	for rows.Next() {
		var i GetUserInstancesRow
		if err := rows.Scan(
			&i.Pkuser,
			&i.Pktopic,
			&i.Topicname,
			&i.Pkinstance,
			&i.StartDate,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getUserTopics = `-- name: GetUserTopics :many
SELECT u.PkUser, t.PkTopic, t.TopicName
FROM dbtraxl.users u, 
     dbtraxl.topics t
WHERE u.PkUser = w.FkUser
AND U.PkUser = $1
`

type GetUserTopicsRow struct {
	Pkuser    int64  `db:"pkuser" json:"pkuser"`
	Pktopic   int64  `db:"pktopic" json:"pktopic"`
	Topicname string `db:"topicname" json:"topicname"`
}

// get topics for a particular user
func (q *Queries) GetUserTopics(ctx context.Context, pkuser int64) ([]GetUserTopicsRow, error) {
	rows, err := q.query(ctx, q.getUserTopicsStmt, getUserTopics, pkuser)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetUserTopicsRow
	for rows.Next() {
		var i GetUserTopicsRow
		if err := rows.Scan(&i.Pkuser, &i.Pktopic, &i.Topicname); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const insertInstance = `-- name: InsertInstance :one
INSERT INTO dbTraxl.instances (Start_Date, FkTopic, FkUser)
VALUES ($1, $2, $3) RETURNING PkInstance
`

type InsertInstanceParams struct {
	StartDate time.Time `db:"start_date" json:"startDate"`
	Fktopic   int64     `db:"fktopic" json:"fktopic"`
	Fkuser    int64     `db:"fkuser" json:"fkuser"`
}

// insert instance for topic and user of particular ids and returns new instance id
func (q *Queries) InsertInstance(ctx context.Context, arg InsertInstanceParams) (int64, error) {
	row := q.queryRow(ctx, q.insertInstanceStmt, insertInstance, arg.StartDate, arg.Fktopic, arg.Fkuser)
	var pkinstance int64
	err := row.Scan(&pkinstance)
	return pkinstance, err
}

const insertTopic = `-- name: InsertTopic :one
INSERT INTO dbTraxl.topics (TopicName, FkUser)
VALUES ($1, $2) RETURNING PkTopic
`

type InsertTopicParams struct {
	Topicname string `db:"topicname" json:"topicname"`
	Fkuser    int64  `db:"fkuser" json:"fkuser"`
}

// insert topic for a user of a particular id and returns the new topic id
func (q *Queries) InsertTopic(ctx context.Context, arg InsertTopicParams) (int64, error) {
	row := q.queryRow(ctx, q.insertTopicStmt, insertTopic, arg.Topicname, arg.Fkuser)
	var pktopic int64
	err := row.Scan(&pktopic)
	return pktopic, err
}

const listInstances = `-- name: ListInstances :many
SELECT pkinstance, fktopic, fkuser, start_date
FROM dbtraxl.instances
ORDER BY pkinstance
`

// get all instances ordered by id
func (q *Queries) ListInstances(ctx context.Context) ([]DbtraxlInstance, error) {
	rows, err := q.query(ctx, q.listInstancesStmt, listInstances)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []DbtraxlInstance
	for rows.Next() {
		var i DbtraxlInstance
		if err := rows.Scan(
			&i.Pkinstance,
			&i.Fktopic,
			&i.Fkuser,
			&i.StartDate,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listTopics = `-- name: ListTopics :many
SELECT pktopic, topicname, fkuser
FROM dbtraxl.topics
ORDER BY topicname
`

// get all topics ordered by topic name
func (q *Queries) ListTopics(ctx context.Context) ([]DbtraxlTopic, error) {
	rows, err := q.query(ctx, q.listTopicsStmt, listTopics)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []DbtraxlTopic
	for rows.Next() {
		var i DbtraxlTopic
		if err := rows.Scan(&i.Pktopic, &i.Topicname, &i.Fkuser); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listUsers = `-- name: ListUsers :many
SELECT pkuser, username, passwordhash, name, config, createdts, isenabled
FROM dbtraxl.users
ORDER BY userName
`

// get all users ordered by user name
func (q *Queries) ListUsers(ctx context.Context) ([]DbtraxlUser, error) {
	rows, err := q.query(ctx, q.listUsersStmt, listUsers)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []DbtraxlUser
	for rows.Next() {
		var i DbtraxlUser
		if err := rows.Scan(
			&i.Pkuser,
			&i.Username,
			&i.Passwordhash,
			&i.Name,
			&i.Config,
			&i.Createdts,
			&i.Isenabled,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateInstance = `-- name: UpdateInstance :one
UPDATE dbTraxl.instances
SET Start_Date = $1, FkTopic = $2
WHERE PkInstance = $3 RETURNING PkInstance
`

type UpdateInstanceParams struct {
	StartDate  time.Time `db:"start_date" json:"startDate"`
	Fktopic    int64     `db:"fktopic" json:"fktopic"`
	Pkinstance int64     `db:"pkinstance" json:"pkinstance"`
}

// updates instance timestamp and top of particular ID
func (q *Queries) UpdateInstance(ctx context.Context, arg UpdateInstanceParams) (int64, error) {
	row := q.queryRow(ctx, q.updateInstanceStmt, updateInstance, arg.StartDate, arg.Fktopic, arg.Pkinstance)
	var pkinstance int64
	err := row.Scan(&pkinstance)
	return pkinstance, err
}

const updateTopic = `-- name: UpdateTopic :one
UPDATE dbTraxl.topics
SET TopicName = $1
WHERE PkTopic = $2 RETURNING PkTopic
`

type UpdateTopicParams struct {
	Topicname string `db:"topicname" json:"topicname"`
	Pktopic   int64  `db:"pktopic" json:"pktopic"`
}

// update a particular id
func (q *Queries) UpdateTopic(ctx context.Context, arg UpdateTopicParams) (int64, error) {
	row := q.queryRow(ctx, q.updateTopicStmt, updateTopic, arg.Topicname, arg.Pktopic)
	var pktopic int64
	err := row.Scan(&pktopic)
	return pktopic, err
}
