CREATE TABLE "barber" (
  "username" varchar PRIMARY KEY,
  "full_name" varchar NOT NULL,
  "email" varchar UNIQUE NOT NULL,
  "hashed_password" varchar NOT NULL,
  "avatar" varchar,
  "role" varchar,
  "status" varchar,
  "is_manager" bool NOT NULL DEFAULT false,
  "store_work" uuid,
  "password_changed_at" timestamptz NOT NULL DEFAULT '0001-01-01 00:00:00Z',
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "update_at" timestamptz
);

CREATE TABLE "sessions_barber" (
  "id" uuid PRIMARY KEY DEFAULT (uuid_generate_v4()),
  "username" varchar NOT NULL,
  "refresh_token" varchar NOT NULL,
  "user_agent" varchar NOT NULL,
  "client_ip" varchar NOT NULL,
  "fcm_device" varchar NOT NULL,
  "is_blocked" bool NOT NULL DEFAULT false,
  "expires_at" timestamptz NOT NULL,
  "create_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "store" (
  "id" uuid PRIMARY KEY DEFAULT (uuid_generate_v4()),
  "name_id" varchar NOT NULL,
  "name_store" varchar NOT NULL,
  "location" real NOT NULL,
  "address" varchar NOT NULL,
  "image" varchar,
  "list_image" varchar[],
  "manager_id" varchar[],
  "employee_id" varchar[],
  "status" integer NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "update_at" timestamptz
);

CREATE TABLE "service_category" (
  "id" uuid PRIMARY KEY DEFAULT (uuid_generate_v4()),
  "store_id" uuid NOT NULL,
  "work" varchar UNIQUE NOT NULL,
  "description" varchar,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "update_at" timestamptz
);

CREATE TABLE "service" (
  "id" uuid PRIMARY KEY DEFAULT (uuid_generate_v4()),
  "service_category_id" uuid NOT NULL,
  "work" varchar UNIQUE NOT NULL,
  "timer" integer,
  "price" real NOT NULL,
  "description" varchar,
  "image" varchar,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "update_at" timestamptz
);

CREATE INDEX ON "barber" ("username");

CREATE INDEX ON "store" ("name_store");

CREATE INDEX ON "store" ("name_id");

CREATE INDEX ON "store" ("manager_id");

CREATE INDEX ON "service_category" ("store_id");

CREATE INDEX ON "service" ("service_category_id");

ALTER TABLE "barber" ADD FOREIGN KEY ("store_work") REFERENCES "store" ("id");

ALTER TABLE "sessions_barber" ADD FOREIGN KEY ("username") REFERENCES "barber" ("username");

ALTER TABLE "service_category" ADD FOREIGN KEY ("store_id") REFERENCES "store" ("id");

ALTER TABLE "service" ADD FOREIGN KEY ("service_category_id") REFERENCES "service_category" ("id");
