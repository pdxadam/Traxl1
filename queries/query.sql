-- name: ListUsers :many
-- get all users ordered by user name
SELECT *
FROM dbtraxl.users
ORDER BY userName;

-- name: ListTopics :many
-- get all topics ordered by topic name
SELECT *
FROM dbtraxl.topics
ORDER BY topicname;

-- name: ListInstances :many
-- get all instances ordered by id
SELECT *
FROM dbtraxl.instances
ORDER BY pkinstance;

-- name: GetUser :one
-- get user of a particular userid
SELECT *
FROM dbtraxl.users
WHERE PkUser = $1;

-- name: GetUserByName :one
-- get a particular user by username
SELECT *
FROM dbTraxl.users
WHERE username = $1;

-- name: GetUserTopics :many
-- get topics for a particular user
SELECT u.PkUser, t.PkTopic, t.TopicName
FROM dbtraxl.users u, 
     dbtraxl.topics t
WHERE u.PkUser = w.FkUser
AND U.PkUser = $1;

-- name: GetUserInstances :many
-- get all the instances
SELECT u.PkUser, t.PkTopic, t.TopicName, i.PkInstance, i.Start_Date
FROM dbTraxl.users u, dbTraxl.topics t, dbTraxl.instances i
WHERE u.PkUser = t.FkUser
AND t.PkTopic = i.FkTopic
AND u.PkUser = $1;

-- name: DeleteUsers :exec
-- delete a particular user
DELETE
FROM dbTraxl.users
WHERE PkUser = $1;

-- name: DeleteUserTopics :exec
-- delete a user's topics
DELETE 
FROM dbTraxl.topics
WHERE FkUser = $1;

-- name: DeleteTopic :exec
-- Delete a particular topic
DELETE
FROM dbTraxl.topics
WHERE PkTopic = $1;

-- name: DeleteUserInstances :exec
-- delete all instances for a particular user
DELETE 
FROM dbTraxl.instances
WHERE FkUser = $1;

-- name: DeleteTopicInstances :exec
-- delete the instances for a particular topic
DELETE 
FROM dbTraxl.instances
WHERE FkTopic = $1;

-- name: DeleteInstance :exec
-- delete a particular instance
DELETE
FROM dbTraxl.instances
WHERE PkInstance = $1;

-- name: InsertTopic :one
-- insert topic for a user of a particular id and returns the new topic id
INSERT INTO dbTraxl.topics (TopicName, FkUser)
VALUES ($1, $2) RETURNING PkTopic;

-- name: UpdateTopic :one
-- update a particular id
UPDATE dbTraxl.topics
SET TopicName = $1
WHERE PkTopic = $2 RETURNING PkTopic;

-- name: InsertInstance :one
-- insert instance for topic and user of particular ids and returns new instance id
INSERT INTO dbTraxl.instances (Start_Date, FkTopic, FkUser)
VALUES ($1, $2, $3) RETURNING PkInstance;

-- name: UpdateInstance :one
-- updates instance timestamp and top of particular ID
UPDATE dbTraxl.instances
SET Start_Date = $1, FkTopic = $2
WHERE PkInstance = $3 RETURNING PkInstance;

-- name: CreateUsers :one
-- insert new user
INSERT INTO dbTraxl.users (userName, PasswordHash, name )
VALUES ($1,
        $2,
        $3) RETURNING *;








