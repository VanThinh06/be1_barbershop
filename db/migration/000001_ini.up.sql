CREATE TABLE "users" (
  "username" varchar PRIMARY KEY,
  "full_name" varchar,
  "email" varchar UNIQUE NOT NULL,
  "hashed_password" varchar NOT NULL,
  "password_changed_at" timestamptz NOT NULL DEFAULT '0001-01-01 00:00:00Z',
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "barber" (
  "id" uuid UNIQUE PRIMARY KEY DEFAULT (uuid_generate_v4()),
  "name_id" varchar UNIQUE NOT NULL,
  "store_id" uuid,
  "store_manager" uuid[],
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "update_at" timestamptz
);

CREATE TABLE "store" (
  "id" uuid UNIQUE PRIMARY KEY DEFAULT (uuid_generate_v4()),
  "name_id" varchar UNIQUE NOT NULL,
  "name_store" varchar NOT NULL,
  "location" integer,
  "image" varchar,
  "list_image" varchar[],
  "manager_id" uuid UNIQUE,
  "employee_id" uuid[],
  "status" varchar NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "update_at" timestamptz
);

CREATE INDEX ON "barber" ("name_id");

CREATE INDEX ON "store" ("name_store");

CREATE INDEX ON "store" ("name_id");

ALTER TABLE "barber" ADD FOREIGN KEY ("name_id") REFERENCES "users" ("username");

ALTER TABLE "store" ADD FOREIGN KEY ("manager_id") REFERENCES "barber" ("id");
