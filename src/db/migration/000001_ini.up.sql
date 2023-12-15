-- CREATE EXTENSION IF NOT EXISTS "postgis";
-- CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
CREATE TABLE "BarberShops" (
  "shop_id" uuid PRIMARY KEY DEFAULT (uuid_generate_v4()),
  "owner_id" uuid NOT NULL,
  "name" varchar NOT NULL,
  "facility" integer NOT NULL DEFAULT 1,
  "address" varchar NOT NULL,
  "image" varchar,
  "list_image" varchar[],
  "status" integer NOT NULL DEFAULT 1,
  "coordinates" GEOGRAPHY(Point, 4326) NOT NULL,
  "rate" float,
  "is_reputation" bool DEFAULT false,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "update_at" timestamptz
);

CREATE TABLE "Barbers" (
  "barber_id" uuid PRIMARY KEY DEFAULT (uuid_generate_v4()),
  "shop_id" uuid,
  "manager_id" uuid,
  "nick_name" varchar NOT NULL,
  "full_name" varchar NOT NULL,
  "phone" varchar UNIQUE NOT NULL,
  "email" varchar UNIQUE NOT NULL,
  "gender" integer NOT NULL,
  "role" integer NOT NULL,
  "hashed_password" varchar NOT NULL,
  "avatar" varchar,
  "status" integer DEFAULT 1,
  "password_changed_at" timestamptz NOT NULL DEFAULT '0001-01-01 00:00:00Z',
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "update_at" timestamptz
);

CREATE TABLE "SessionsBarber" (
  "id" uuid PRIMARY KEY DEFAULT (uuid_generate_v4()),
  "barber_id" uuid NOT NULL,
  "refresh_token" varchar NOT NULL,
  "user_agent" varchar NOT NULL,
  "client_ip" varchar NOT NULL,
  "fcm_device" varchar NOT NULL,
  "is_blocked" bool NOT NULL DEFAULT false,
  "timezone" varchar NOT NULL,
  "expires_at" timestamptz NOT NULL,
  "create_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "ServiceCategory" (
  "id" uuid PRIMARY KEY DEFAULT (uuid_generate_v4()),
  "shop_id" uuid NOT NULL,
  "name" varchar NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "update_at" timestamptz
);

CREATE TABLE "Services" (
  "id" uuid PRIMARY KEY DEFAULT (uuid_generate_v4()),
  "category_id" uuid NOT NULL,
  "name" varchar NOT NULL,
  "timer" integer,
  "price" real,
  "description" varchar,
  "image" varchar,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "update_at" timestamptz
);

CREATE TABLE "Customers" (
  "customer_id" uuid PRIMARY KEY DEFAULT (uuid_generate_v4()),
  "name" varchar NOT NULL,
  "email" varchar UNIQUE NOT NULL,
  "phone" varchar UNIQUE,
  "gender" integer NOT NULL DEFAULT 1,
  "hashed_password" varchar,
  "avatar" varchar,
  "is_social_auth" bool DEFAULT false,
  "password_changed_at" timestamptz NOT NULL DEFAULT '0001-01-01 00:00:00Z',
  "create_at" timestamptz NOT NULL DEFAULT (now()),
  "update_at" timestamptz
);

CREATE TABLE "SessionsCustomer" (
  "id" uuid PRIMARY KEY DEFAULT (uuid_generate_v4()),
  "customer_id" uuid NOT NULL,
  "refresh_token" varchar NOT NULL,
  "user_agent" varchar NOT NULL,
  "client_ip" varchar NOT NULL,
  "fcm_device" varchar NOT NULL,
  "timezone" varchar NOT NULL,
  "coordinates" GEOGRAPHY(Point, 4326) NOT NULL,
  "is_blocked" bool NOT NULL DEFAULT false,
  "expires_at" timestamptz NOT NULL,
  "create_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "Appointments" (
  "appointment_id" uuid PRIMARY KEY DEFAULT (uuid_generate_v4()),
  "barbershops_id" uuid NOT NULL,
  "customer_id" uuid NOT NULL,
  "barber_id" uuid NOT NULL,
  "service_id" uuid UNIQUE NOT NULL,
  "appointment_datetime" timestamptz NOT NULL,
  "status" integer NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "update_at" timestamptz
);

CREATE INDEX ON "BarberShops" ("owner_id");

CREATE INDEX ON "BarberShops" ("name");

CREATE INDEX ON "Barbers" ("barber_id");

CREATE INDEX ON "Barbers" ("shop_id");

CREATE INDEX ON "Barbers" ("nick_name");

CREATE INDEX ON "ServiceCategory" ("id");

CREATE INDEX ON "ServiceCategory" ("shop_id");

CREATE INDEX ON "Services" ("id");

CREATE INDEX ON "Services" ("category_id");

CREATE INDEX ON "Customers" ("customer_id");

CREATE INDEX ON "Customers" ("email");

CREATE INDEX ON "Appointments" ("appointment_id");

CREATE INDEX ON "Appointments" ("barber_id");

CREATE INDEX ON "Appointments" ("customer_id");

ALTER TABLE "BarberShops" ADD FOREIGN KEY ("owner_id") REFERENCES "Barbers" ("barber_id");

ALTER TABLE "Barbers" ADD FOREIGN KEY ("shop_id") REFERENCES "BarberShops" ("shop_id");

ALTER TABLE "Barbers" ADD FOREIGN KEY ("manager_id") REFERENCES "Barbers" ("barber_id");

ALTER TABLE "SessionsBarber" ADD FOREIGN KEY ("barber_id") REFERENCES "Barbers" ("barber_id");

ALTER TABLE "ServiceCategory" ADD FOREIGN KEY ("shop_id") REFERENCES "BarberShops" ("shop_id");

ALTER TABLE "Services" ADD FOREIGN KEY ("category_id") REFERENCES "ServiceCategory" ("id");

ALTER TABLE "SessionsCustomer" ADD FOREIGN KEY ("customer_id") REFERENCES "Customers" ("customer_id");

ALTER TABLE "Appointments" ADD FOREIGN KEY ("barbershops_id") REFERENCES "BarberShops" ("shop_id");

ALTER TABLE "Appointments" ADD FOREIGN KEY ("customer_id") REFERENCES "Customers" ("customer_id");

ALTER TABLE "Appointments" ADD FOREIGN KEY ("barber_id") REFERENCES "Barbers" ("barber_id");

CREATE TABLE "Services_Appointments" (
  "Services_id" uuid,
  "Appointments_service_id" uuid,
  PRIMARY KEY ("Services_id", "Appointments_service_id")
);

ALTER TABLE "Services_Appointments" ADD FOREIGN KEY ("Services_id") REFERENCES "Services" ("id");

ALTER TABLE "Services_Appointments" ADD FOREIGN KEY ("Appointments_service_id") REFERENCES "Appointments" ("service_id");

