
CREATE TABLE "service" (
  "id" uuid PRIMARY KEY DEFAULT (uuid_generate_v4()),
  "store_id" uuid NOT NULL,
  "work" varchar UNIQUE NOT NULL,
  "timer" timestamptz NOT NULL,
  "price" numeric NOT NULL,
  "description" varchar,
  "image" varchar,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "update_at" timestamptz
);

CREATE TABLE "schedulerwork" (
  "id" uuid PRIMARY KEY DEFAULT (uuid_generate_v4()),
  "barber_id" uuid NOT NULL,
  "users_id" varchar NOT NULL,
  "timerstart" timestamptz NOT NULL,
  "timerend" timestamptz NOT NULL,
  "service" uuid[],
  "total_price" numeric NOT NULL DEFAULT 0,
  "status" integer NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "update_at" timestamptz
);


CREATE INDEX ON "service" ("store_id");

CREATE UNIQUE INDEX ON "schedulerwork" ("timerstart", "timerend");

ALTER TABLE "service" ADD FOREIGN KEY ("store_id") REFERENCES "store" ("id");

ALTER TABLE "schedulerwork" ADD FOREIGN KEY ("barber_id") REFERENCES "barber" ("id");

ALTER TABLE "schedulerwork" ADD FOREIGN KEY ("users_id") REFERENCES "users" ("username");