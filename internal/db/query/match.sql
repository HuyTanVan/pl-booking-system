-- name: CreateMatch :one
INSERT INTO matches (
    home_team_id,
    away_team_id,
    stadium_id,
    match_date,
    session,
    status
)
VALUES ($1, $2, $3, $4, $5, $6)
RETURNING *;

-- name: ListMatches :many
SELECT * FROM matches
ORDER BY match_date
LIMIT $1
OFFSET $2;

-- name: ListMatchesWithDetails :many
SELECT
  m.*,
  home_team.name AS home_team_name,
  away_team.name AS away_team_name,
  stadium.name AS stadium_name,
  COUNT(CASE WHEN ticket.is_available = true THEN 1 END) AS available,
  TO_CHAR(MIN(CASE WHEN ticket.price > 0 THEN ticket.price ELSE 0 END)::float, 'FM999999999.00') AS min_price
FROM matches m
LEFT JOIN teams AS home_team ON m.home_team_id = home_team.id
LEFT JOIN teams AS away_team ON m.away_team_id = away_team.id
LEFT JOIN stadiums AS stadium ON m.stadium_id = stadium.id
LEFT JOIN tickets AS ticket ON m.id = ticket.match_id
WHERE m.match_date >= CURRENT_DATE
  AND m.match_date <= 
      CASE
        WHEN sqlc.narg('range') = 'a' THEN CURRENT_DATE + INTERVAL '90 DAY'
        WHEN sqlc.narg('range') = 't' THEN CURRENT_DATE + INTERVAL '1 DAY'
        WHEN sqlc.narg('range') = 'w' THEN CURRENT_DATE + INTERVAL '7 DAY'
        WHEN sqlc.narg('range') = 'm' THEN CURRENT_DATE + INTERVAL '30 DAY'
        ELSE COALESCE(sqlc.narg('match_date'), CURRENT_DATE)  -- Default to today if no custom date
      END
GROUP BY m.id, home_team.name, away_team.name, stadium.name
ORDER BY 
      CASE
        WHEN sqlc.narg('sortby') = 'd' THEN match_date
        -- WHEN sqlc.narg('sortby') = 'p' THEN min_price
        ELSE COALESCE(sqlc.narg('match_date'), m.match_date)
      END
LIMIT $1
OFFSET $2;


-- -- name: ListMatchesWithDetails :many
-- SELECT
--   m.*,
--   home_team.name AS home_team_name,
--   away_team.name AS away_team_name,
--   stadium.name AS stadium_name
-- FROM matches m
-- LEFT JOIN teams AS home_team ON m.home_team_id = home_team.id
-- LEFT JOIN teams AS away_team ON m.away_team_id = away_team.id
-- LEFT JOIN stadiums AS stadium ON m.stadium_id = stadium.id
-- ORDER BY m.match_date
-- LIMIT $1
-- OFFSET $2;

-- name: GetMatchByID :one
SELECT * FROM matches
WHERE id=$1 LIMIT 1;

-- name: UpdateMatch :exec
UPDATE matches
SET home_team_id = coalesce(sqlc.narg('home_team_id'), home_team_id),
    away_team_id = coalesce(sqlc.narg('away_team_id'), away_team_id),
    stadium_id = coalesce(sqlc.narg('stadium_id'), stadium_id),
    match_date = coalesce(sqlc.narg('match_date'), match_date),
    session = coalesce(sqlc.narg('session'), session),
    status = coalesce(sqlc.narg('status'), status)
WHERE id=$1;