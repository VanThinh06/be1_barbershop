CREATE TABLE "sessions" (
  "id" uuid PRIMARY KEY,
  "username" varchar UNIQUE NOT NULL,
  "refresh_token" varchar NOT NULL,
  "user_agent" VARCHAR NOT NULL,
  "client_ip" VARCHAR NOT NULL,
  "is_blocked" BOOLEAN NOT NULL DEFAULT FALSE,
  "expires_at" timestamptz NOT NULL,
  "update_at" timestamptz NOT NULL DEFAULT (now())
);
ALTER TABLE "sessions" ADD FOREIGN KEY ("username") REFERENCES "users" ("username");