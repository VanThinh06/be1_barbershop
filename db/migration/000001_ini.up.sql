CREATE TABLE "users" (
  "username" varchar PRIMARY KEY,
  "full_name" varchar NOT NULL,
  "email" varchar UNIQUE NOT NULL,
  "image" varchar,
  "role" varchar,
  "location" real,
  "address" varchar,
  "hashed_password" varchar NOT NULL,
  "password_changed_at" timestamptz NOT NULL DEFAULT '0001-01-01 00:00:00Z',
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "barber" (
  "id" uuid UNIQUE PRIMARY KEY DEFAULT (uuid_generate_v4()),
  "name_id" varchar UNIQUE NOT NULL,
  "store_id" uuid,
  "status" varchar,
  "store_manager" uuid[],
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "update_at" timestamptz
);

CREATE TABLE "store" (
  "id" uuid PRIMARY KEY DEFAULT (uuid_generate_v4()),
  "name_id" varchar UNIQUE NOT NULL,
  "name_store" varchar NOT NULL,
  "location" real,
  "image" varchar,
  "list_image" varchar[],
  "manager_id" uuid,
  "employee_id" uuid[],
  "status" integer NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "update_at" timestamptz
);

CREATE TABLE "sessions" (
  "id" uuid PRIMARY KEY DEFAULT (uuid_generate_v4()),
  "username" varchar NOT NULL,
  "refresh_token" varchar NOT NULL,
  "user_agent" varchar NOT NULL,
  "client_ip" varchar NOT NULL,
  "is_blocked" bool NOT NULL DEFAULT false,
  "expires_at" timestamptz NOT NULL,
  "update_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "service" (
  "id" uuid PRIMARY KEY DEFAULT (uuid_generate_v4()),
  "store_id" uuid NOT NULL,
  "work" varchar UNIQUE NOT NULL,
  "timer" integer NOT NULL,
  "price" real NOT NULL,
  "description" varchar,
  "image" varchar,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "update_at" timestamptz
);

CREATE TABLE "schedulerwork" (
  "id" uuid PRIMARY KEY DEFAULT (uuid_generate_v4()),
  "barber_id" uuid  NOT NULL,
  "users_id" varchar  NOT NULL,
  "timerstart" timestamptz NOT NULL,
  "timerend" timestamptz NOT NULL,
  "service" uuid[],
  "total_price" real NOT NULL DEFAULT 0,
  "status" integer NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "update_at" timestamptz
);

CREATE INDEX ON "barber" ("name_id");

CREATE INDEX ON "store" ("name_store");

CREATE INDEX ON "store" ("name_id");

CREATE INDEX ON "store" ("manager_id");

CREATE INDEX ON "service" ("store_id");

CREATE UNIQUE INDEX ON "schedulerwork" ("timerstart", "timerend");

ALTER TABLE "barber" ADD FOREIGN KEY ("name_id") REFERENCES "users" ("username");

ALTER TABLE "store" ADD FOREIGN KEY ("manager_id") REFERENCES "barber" ("id");

ALTER TABLE "sessions" ADD FOREIGN KEY ("username") REFERENCES "users" ("username");

ALTER TABLE "service" ADD FOREIGN KEY ("store_id") REFERENCES "store" ("id");

ALTER TABLE "schedulerwork" ADD FOREIGN KEY ("barber_id") REFERENCES "barber" ("id");

ALTER TABLE "schedulerwork" ADD FOREIGN KEY ("users_id") REFERENCES "users" ("username");
