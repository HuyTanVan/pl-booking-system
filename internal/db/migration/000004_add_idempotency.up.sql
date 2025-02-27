-- CREATE TYPE header_pair AS (
--    "name" VARCHAR,
--    "value" BYTEA
-- );

CREATE TABLE idempotency (
   "user_id" INTEGER NOT NULL,
   "idempotency_key" VARCHAR NOT NULL,
   "response_status_code" SMALLINT NOT NULL,
   -- "response_headers" header_pair[] NOT NULL,
   "response_headers" BYTEA NOT NULL,
   "response_body" BYTEA NOT NULL,
   "created_at" timestamptz NOT NULL,
   PRIMARY KEY(user_id, idempotency_key)
);

ALTER TABLE "idempotency" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");