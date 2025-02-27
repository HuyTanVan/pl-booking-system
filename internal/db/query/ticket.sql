-- name: CreateTicket :one
INSERT INTO tickets (
    match_id,
    seat_id,
    price,
    is_available,
    created_at,
    updated_at
)
VALUES ($1, $2, $3, $4, $5, $6)
RETURNING *;

-- name: ListTickets :many
SELECT * FROM tickets
WHERE match_id = $1
ORDER BY price
LIMIT $2
OFFSET $3;

-- name: ListMinPriceOfTicketsByMatch :many
SELECT
	t.match_id AS match_id, 
	min(t.price) AS min_price
FROM tickets t
LEFT JOIN matches m on m.id = t.match_id
group by t.match_id
ORDER BY min_price
;
-- select t.*,
-- 		s.stadium_id,
-- 		s.block,	
-- 		s.row,
-- 		s.is_available
-- from 
-- tickets AS t
-- left JOIN seats AS s ON t.seat_id = s.id;
-- name: ListTicketsWithDetails :many
SELECT
	t.id,
	t.match_id,
	t.seat_id,
	t.price,
	t.is_available AS ticket_avalable,
	home_team.name AS home_team,
	away_team.name AS away_team,
	m.match_date AS match_date,
	sta.location AS stadium_location,
	s.block,
	s.row,
	s.seat_column
FROM tickets t 
LEFT JOIN matches AS m ON t.match_id = m.id
LEFT JOIN teams AS home_team ON m.home_team_id = home_team.id
LEFT JOIN teams AS away_team ON m.away_team_id = away_team.id
LEFT JOIN seats AS s ON t.seat_id = s.id
LEFT JOIN stadiums AS sta ON s.stadium_id = sta.id
WHERE t.match_id = $1
ORDER BY t.price
LIMIT $2
OFFSET $3;

-- name: GetTicketByID :one
SELECT * FROM tickets
WHERE id = $1 LIMIT 1;

-- name: GetTicketWithDetails :one
SELECT
	t.id,
	t.match_id,
	t.seat_id,
	t.price,
	t.is_available AS ticket_avalable,
	home_team.name AS home_team,
	away_team.name AS away_team,
	m.match_date AS match_date,
	sta.location AS stadium_location,
	s.block,
	s.row,
	s.seat_column
FROM tickets t 
LEFT JOIN matches AS m ON t.match_id = m.id
LEFT JOIN teams AS home_team ON m.home_team_id = home_team.id
LEFT JOIN teams AS away_team ON m.away_team_id = away_team.id
LEFT JOIN seats AS s ON t.seat_id = s.id
LEFT JOIN stadiums AS sta ON s.stadium_id = sta.id
WHERE t.id = $1
LIMIT 1;
-- name: UpdateTicket :exec
UPDATE tickets
SET match_id = coalesce(sqlc.narg('match_id'), match_id),
    seat_id = coalesce(sqlc.narg('seat_id'), seat_id),
    price = coalesce(sqlc.narg('price'), price),
    is_available = coalesce(sqlc.narg('is_available'), is_available)
WHERE id=$1;