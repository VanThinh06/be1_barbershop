-- CREATE EXTENSION IF NOT EXISTS "postgis";
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
CREATE TABLE "BarberShops" (
  "shop_id" uuid PRIMARY KEY DEFAULT (uuid_generate_v4()),
  "owner_id" uuid NOT NULL,
  "chain_id" uuid,
  "name" varchar NOT NULL,
  "facility" integer NOT NULL DEFAULT 1,
  "address" varchar NOT NULL,
  "coordinates" GEOGRAPHY(Point, 4326) NOT NULL,
  "image" varchar,
  "list_image" varchar[],
  "status" integer NOT NULL DEFAULT 1,
  "rate" float,
  "start_time" TIME NOT NULL DEFAULT '08:00:00'::TIME,
  "end_time" TIME NOT NULL DEFAULT '21:00:00'::TIME,
  "break_time" TIME NOT NULL DEFAULT '12:00:00'::TIME,
  "break_minutes" integer NOT NULL DEFAULT 0,
  "interval_scheduler" integer NOT NULL DEFAULT 30,
  "is_reputation" bool DEFAULT false,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "updated_at" timestamptz
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
  "haircut" bool NOT NULL DEFAULT false,
  "hashed_password" varchar NOT NULL,
  "avatar" varchar,
  "status" integer DEFAULT 1,
  "password_changed_at" timestamptz NOT NULL DEFAULT '0001-01-01 00:00:00Z',
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "updated_at" timestamptz
);

CREATE TABLE "SessionsBarbers" (
  "id" uuid PRIMARY KEY DEFAULT (uuid_generate_v4()),
  "barber_id" uuid NOT NULL,
  "refresh_token" varchar NOT NULL,
  "user_agent" varchar NOT NULL,
  "client_ip" varchar NOT NULL,
  "fcm_device" varchar NOT NULL,
  "is_blocked" bool NOT NULL DEFAULT false,
  "timezone" varchar NOT NULL,
  "expires_at" timestamptz NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "ServiceCategories" (
  "id" uuid PRIMARY KEY DEFAULT (uuid_generate_v4()),
  "chain_id" uuid,
  "shop_id" uuid,
  "gender" integer NOT NULL DEFAULT 1,
  "name" varchar NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "updated_at" timestamptz,
  "hidden" boolean NOT NULL DEFAULT false
);

CREATE TABLE "Services" (
  "id" uuid PRIMARY KEY DEFAULT (uuid_generate_v4()),
  "category_id" uuid NOT NULL,
  "shop_id" uuid,
  "name" varchar NOT NULL,
  "timer" integer,
  "price" real,
  "description" varchar,
  "image" varchar,
  "hidden" boolean NOT NULL DEFAULT false,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "updated_at" timestamptz
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
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "updated_at" timestamptz
);

CREATE TABLE "SessionsCustomers" (
  "id" uuid PRIMARY KEY DEFAULT (uuid_generate_v4()),
  "customer_id" uuid NOT NULL,
  "refresh_token" varchar NOT NULL,
  "user_agent" varchar NOT NULL,
  "client_ip" varchar NOT NULL,
  "fcm_device" varchar NOT NULL,
  "timezone" varchar NOT NULL,
  "coordinates" GEOGRAPHY(Point, 4326),
  "is_blocked" bool NOT NULL DEFAULT false,
  "expires_at" timestamptz NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "Appointments" (
  "appointment_id" uuid PRIMARY KEY DEFAULT (uuid_generate_v4()),
  "barbershops_id" uuid NOT NULL,
  "customer_id" uuid NOT NULL,
  "barber_id" uuid NOT NULL,
  "appointment_datetime" timestamptz NOT NULL,
  "timer" integer NOT NULL,
  "status" integer NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "updated_at" timestamptz
);

CREATE OR REPLACE FUNCTION check_appointment_conflict()
RETURNS TRIGGER AS $$
BEGIN
  IF EXISTS (
    SELECT 1
    FROM "Appointments"
    WHERE "barbershops_id" = NEW."barbershops_id"
      AND "barber_id" = NEW."barber_id"
      AND (
        (NEW."appointment_datetime" + NEW."timer" * interval '1 minute') BETWEEN "appointment_datetime" AND "appointment_datetime" + NEW."timer" * interval '1 minute'
        OR NEW."appointment_datetime" BETWEEN "appointment_datetime" AND "appointment_datetime" + NEW."timer" * interval '1 minute'
      )
  ) THEN
    RAISE EXCEPTION 'Appointment conflict: Another appointment exists in this time slot.';
  END IF;
  RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER check_appointment_conflict_trigger
BEFORE INSERT ON "Appointments"
FOR EACH ROW EXECUTE FUNCTION check_appointment_conflict();

CREATE TABLE "Chains" (
  "chain_id" uuid PRIMARY KEY DEFAULT (uuid_generate_v4()),
  "owner_id" uuid UNIQUE NOT NULL,
  "name" varchar NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "updated_at" timestamptz
);

CREATE INDEX ON "BarberShops" ("owner_id");

CREATE INDEX ON "BarberShops" ("chain_id");

CREATE INDEX ON "BarberShops" ("name");

CREATE UNIQUE INDEX ON "BarberShops" ("chain_id", "facility");

CREATE UNIQUE INDEX ON "BarberShops" ("owner_id", "facility");

CREATE INDEX ON "Barbers" ("shop_id");

CREATE INDEX ON "Barbers" ("nick_name");

CREATE INDEX ON "Barbers" ("email");

CREATE INDEX ON "Barbers" ("phone");

CREATE INDEX ON "ServiceCategories" ("shop_id");

CREATE INDEX ON "ServiceCategories" ("chain_id");

CREATE UNIQUE INDEX ON "ServiceCategories" ("chain_id", "name");

CREATE UNIQUE INDEX ON "ServiceCategories" ("shop_id", "name");

CREATE INDEX ON "Services" ("category_id");

CREATE INDEX ON "Services" ("shop_id");

CREATE UNIQUE INDEX ON "Services" ("category_id", "name");

CREATE UNIQUE INDEX ON "Services" ("shop_id", "name");

CREATE INDEX ON "Customers" ("email");

CREATE INDEX ON "Customers" ("phone");

CREATE INDEX ON "Appointments" ("barber_id");

CREATE INDEX ON "Appointments" ("customer_id");

CREATE UNIQUE INDEX ON "Appointments" ("barber_id", "appointment_datetime");

ALTER TABLE "BarberShops" ADD FOREIGN KEY ("owner_id") REFERENCES "Barbers" ("barber_id");

ALTER TABLE "BarberShops" ADD FOREIGN KEY ("chain_id") REFERENCES "Chains" ("chain_id");

ALTER TABLE "Barbers" ADD FOREIGN KEY ("shop_id") REFERENCES "BarberShops" ("shop_id");

ALTER TABLE "Barbers" ADD FOREIGN KEY ("manager_id") REFERENCES "Barbers" ("barber_id");

ALTER TABLE "SessionsBarbers" ADD FOREIGN KEY ("barber_id") REFERENCES "Barbers" ("barber_id");

ALTER TABLE "ServiceCategories" ADD FOREIGN KEY ("chain_id") REFERENCES "Chains" ("chain_id");

ALTER TABLE "ServiceCategories" ADD FOREIGN KEY ("shop_id") REFERENCES "BarberShops" ("shop_id");

ALTER TABLE "Services" ADD FOREIGN KEY ("category_id") REFERENCES "ServiceCategories" ("id");

ALTER TABLE "Services" ADD FOREIGN KEY ("shop_id") REFERENCES "BarberShops" ("shop_id");

ALTER TABLE "SessionsCustomers" ADD FOREIGN KEY ("customer_id") REFERENCES "Customers" ("customer_id");

ALTER TABLE "Appointments" ADD FOREIGN KEY ("barbershops_id") REFERENCES "BarberShops" ("shop_id");

ALTER TABLE "Appointments" ADD FOREIGN KEY ("customer_id") REFERENCES "Customers" ("customer_id");

ALTER TABLE "Appointments" ADD FOREIGN KEY ("barber_id") REFERENCES "Barbers" ("barber_id");

CREATE TABLE "Services_Appointments" (
  "Services_id" uuid,
  "Appointments_service_id" uuid,
  PRIMARY KEY ("Services_id", "Appointments_service_id")
);

ALTER TABLE "Services_Appointments" ADD FOREIGN KEY ("Services_id") REFERENCES "Services" ("id");

ALTER TABLE "Services_Appointments" ADD FOREIGN KEY ("Appointments_service_id") REFERENCES "Appointments" ("appointment_id");

ALTER TABLE "Chains" ADD FOREIGN KEY ("owner_id") REFERENCES "Barbers" ("barber_id");

