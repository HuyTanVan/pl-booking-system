-- name: CreateTeam :one
INSERT INTO teams (
    name,
    stadium_id,
    logo
)
VALUES ($1, $2, $3)
RETURNING *;

-- name: ListTeams :many
SELECT * FROM teams
ORDER BY id
LIMIT $1
OFFSET $2;

-- name: GetTeamByID :one
SELECT * FROM teams
WHERE id = $1 LIMIT 1;

-- name: UpdateTeam :exec
UPDATE teams
SET name = coalesce(sqlc.narg('name'), name),
    stadium_id = coalesce(sqlc.narg('stadium_id'), stadium_id),
    logo = coalesce(sqlc.narg('logo'), logo)
WHERE id=$1;