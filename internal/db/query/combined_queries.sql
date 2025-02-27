/* a query getting data needed for homepage  */
-- name: FetchHomePageDetails :many
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

/* a query getting all tickets belong to a match  */
-- name: ListTicketsOfMatch :many
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

/* a query getting all tickets belong to a match  */
-- name: ListTicketsOfMatch :many