-- CREATE EXTENSION IF NOT EXISTS "postgis";
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
CREATE TABLE "Roles" (
  "id" serial  PRIMARY KEY,
  "name" varchar UNIQUE NOT NULL,
  "create_at" timestamptz NOT NULL DEFAULT (now()),
  "update_at" timestamptz NOT NULL DEFAULT '0001-01-01 00:00:00Z'
);

CREATE TABLE "BarberRoles" (
  "id" uuid PRIMARY KEY DEFAULT (uuid_generate_v4()),
  "barber_id" uuid NOT NULL,
  "barbershop_id" uuid NOT NULL,
  "role_id" integer NOT NULL,
  "create_at" timestamptz NOT NULL DEFAULT (now()),
  "update_at" timestamptz NOT NULL DEFAULT '0001-01-01 00:00:00Z'
);

CREATE TABLE "BarberShopChains" (
  "id" uuid PRIMARY KEY DEFAULT (uuid_generate_v4()),
  "name" varchar UNIQUE NOT NULL,
  "description" text,
  "founder" varchar NOT NULL,
  "founding_date" timestamptz NOT NULL ,
  "website" varchar,
  "create_at" timestamptz NOT NULL DEFAULT (now()),
  "update_at" timestamptz NOT NULL DEFAULT '0001-01-01 00:00:00Z'
);

CREATE TABLE "BarberShops" (
  "id" uuid PRIMARY KEY DEFAULT (uuid_generate_v4()),
  "barbershop_chain_id" uuid,
  "name" varchar NOT NULL,
  "is_main_branch" bool NOT NULL DEFAULT false,
  "branch_count" integer NOT NULL DEFAULT 1,
  "coordinates" GEOGRAPHY(Point, 4326) NOT NULL,
  "address" varchar NOT NULL,
  "image" varchar,
  "status" integer NOT NULL DEFAULT 1,
  "rate" float NOT NULL DEFAULT 0,
  "start_time" TIME NOT NULL DEFAULT '08:00:00'::TIME,
  "end_time" TIME NOT NULL DEFAULT '5:00:00'::TIME,
  "break_time" TIME NOT NULL DEFAULT '12:00:00'::TIME,
  "break_minutes" integer NOT NULL DEFAULT 60,
  "interval_scheduler" integer NOT NULL DEFAULT 30,
  "is_reputation" bool NOT NULL DEFAULT false,
  "is_verified" bool NOT NULL DEFAULT false,
  "create_at" timestamptz NOT NULL DEFAULT (now()),
  "update_at" timestamptz NOT NULL DEFAULT '0001-01-01 00:00:00Z'
);

CREATE TABLE "Barbers" (
  "id" uuid PRIMARY KEY DEFAULT (uuid_generate_v4()),
  "gender_id" integer,
  "email" varchar UNIQUE NOT NULL,
  "phone" varchar UNIQUE NOT NULL,
  "hashed_password" varchar NOT NULL,
  "nick_name" varchar UNIQUE NOT NULL,
  "full_name" varchar NOT NULL,
  "haircut" bool NOT NULL DEFAULT false,
  "avatar_url" varchar,
  "password_changed_at" timestamptz NOT NULL DEFAULT '0001-01-01 00:00:00Z',
  "create_at" timestamptz NOT NULL DEFAULT (now()),
  "update_at" timestamptz NOT NULL DEFAULT '0001-01-01 00:00:00Z'
);

CREATE TABLE "BarberManagers" (
  "barber_id" uuid NOT NULL,
  "manager_id" uuid NOT NULL,
  "create_at" timestamptz NOT NULL DEFAULT (now()),
  "update_at" timestamptz NOT NULL DEFAULT '0001-01-01 00:00:00Z'
);

CREATE TABLE "Genders" (
  "id" serial  PRIMARY KEY,
  "name" varchar UNIQUE NOT NULL
);

CREATE TABLE "SessionsBarber" (
  "id" uuid PRIMARY KEY DEFAULT (uuid_generate_v4()),
  "barber_id" uuid NOT NULL,
  "refresh_token" varchar NOT NULL,
  "user_agent" varchar NOT NULL,
  "client_ip" varchar NOT NULL,
  "fcm_device" varchar NOT NULL,
  "is_blocked" bool NOT NULL DEFAULT false,
  "expires_at" timestamptz NOT NULL,
  "create_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "ServiceCategories" (
  "id" uuid PRIMARY KEY DEFAULT (uuid_generate_v4()),
  "name" varchar NOT NULL,
  "is_global" bool NOT NULL DEFAULT false,
  "create_at" timestamptz NOT NULL DEFAULT (now()),
  "update_at" timestamptz NOT NULL DEFAULT '0001-01-01 00:00:00Z'
);

CREATE TABLE "BarberShopServiceCategories" (
  "id" uuid PRIMARY KEY DEFAULT (uuid_generate_v4()),
  "barbershop_chain_id" uuid,
  "barbershop_id" uuid,
  "service_category_id" uuid,
  "create_at" timestamptz NOT NULL DEFAULT (now()),
  "update_at" timestamptz NOT NULL DEFAULT '0001-01-01 00:00:00Z'
);

CREATE TABLE "BarberShopServices" (
  "id" uuid PRIMARY KEY DEFAULT (uuid_generate_v4()),
  "barbershop_category_id" uuid NOT NULL,
  "barbershop_chain_id" uuid,
  "barbershop_id" uuid,
  "gender_id" integer,
  "name" varchar NOT NULL,
  "timer" integer NOT NULL DEFAULT 0,
  "price" real NOT NULL DEFAULT 0,
  "description" varchar,
  "image" varchar,
  "is_custom" bool NOT NULL DEFAULT false,
  "is_removed" bool NOT NULL DEFAULT false,
  "create_at" timestamptz NOT NULL DEFAULT (now()),
  "update_at" timestamptz NOT NULL DEFAULT '0001-01-01 00:00:00Z'
);

CREATE TABLE "Customers" (
  "id" uuid PRIMARY KEY DEFAULT (uuid_generate_v4()),
  "name" varchar NOT NULL,
  "email" varchar UNIQUE NOT NULL,
  "phone" varchar UNIQUE,
  "gender" integer NOT NULL DEFAULT 1,
  "hashed_password" varchar,
  "avatar" varchar,
  "is_social_auth" bool DEFAULT false,
  "password_changed_at" timestamptz NOT NULL DEFAULT '0001-01-01 00:00:00Z',
  "create_at" timestamptz NOT NULL DEFAULT (now()),
  "update_at" timestamptz NOT NULL DEFAULT '0001-01-01 00:00:00Z'
);

CREATE TABLE "SessionsCustomer" (
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
  "create_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "Appointments" (
  "id" uuid PRIMARY KEY DEFAULT (uuid_generate_v4()),
  "barbershop_id" uuid NOT NULL,
  "customer_id" uuid NOT NULL,
  "barber_id" uuid NOT NULL,
  "service_id" uuid NOT NULL,
  "appointment_datetime" timestamptz NOT NULL,
  "status" integer NOT NULL,
  "create_at" timestamptz NOT NULL DEFAULT (now()),
  "update_at" timestamptz NOT NULL DEFAULT '0001-01-01 00:00:00Z'
);

CREATE OR REPLACE FUNCTION check_appointment_conflict()
RETURNS TRIGGER AS $$
BEGIN
  IF EXISTS (
    SELECT 1
    FROM "Appointments"
    WHERE "barbershop_id" = NEW."barbershop_id"
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

CREATE TABLE "BarberShopReviews" (
  "id" uuid PRIMARY KEY DEFAULT (uuid_generate_v4()),
  "customer_id" uuid NOT NULL,
  "barbershop_id" uuid NOT NULL,
  "rating" integer NOT NULL,
  "comment" varchar,
  "create_at" timestamptz NOT NULL DEFAULT (now()),
  "update_at" timestamptz NOT NULL DEFAULT '0001-01-01 00:00:00Z'
);

CREATE TABLE "BarberReviews" (
  "id" uuid PRIMARY KEY DEFAULT (uuid_generate_v4()),
  "barber_id" uuid NOT NULL,
  "customer_id" uuid NOT NULL,
  "rating" integer NOT NULL,
  "comment" varchar,
  "create_at" timestamptz NOT NULL DEFAULT (now()),
  "update_at" timestamptz NOT NULL DEFAULT '0001-01-01 00:00:00Z'
);

CREATE INDEX ON "BarberRoles" ("barber_id");

CREATE INDEX ON "BarberRoles" ("barbershop_id");

CREATE INDEX ON "BarberRoles" ("role_id");

CREATE INDEX ON "BarberShopChains" ("name");

CREATE INDEX ON "BarberShops" ("barbershop_chain_id");

CREATE INDEX ON "BarberShops" ("name");

CREATE UNIQUE INDEX ON "BarberShops" ("barbershop_chain_id", "branch_count");

CREATE INDEX ON "Barbers" ("phone");
CREATE INDEX ON "Barbers" ("email");
CREATE INDEX ON "Barbers" ("nick_name");

CREATE INDEX ON "BarberManagers" ("barber_id");

CREATE INDEX ON "BarberManagers" ("manager_id");

CREATE UNIQUE INDEX ON "BarberManagers" ("manager_id", "barber_id");

CREATE INDEX ON "ServiceCategories" ("name");

CREATE UNIQUE INDEX ON "BarberShopServiceCategories" ("barbershop_chain_id", "service_category_id");

CREATE UNIQUE INDEX ON "BarberShopServiceCategories" ("barbershop_id", "service_category_id");

CREATE INDEX ON "BarberShopServices" ("barbershop_category_id");

CREATE UNIQUE INDEX ON "BarberShopServices" ("barbershop_category_id", "name");

CREATE INDEX ON "Customers" ("phone");

CREATE INDEX ON "Customers" ("email");

CREATE INDEX ON "Appointments" ("barber_id");

CREATE INDEX ON "Appointments" ("customer_id");

CREATE INDEX ON "Appointments" ("barbershop_id");

CREATE UNIQUE INDEX ON "Appointments" ("barber_id", "appointment_datetime");

CREATE INDEX ON "BarberShopReviews" ("customer_id");

CREATE INDEX ON "BarberShopReviews" ("barbershop_id");

CREATE INDEX ON "BarberReviews" ("barber_id");

CREATE INDEX ON "BarberReviews" ("customer_id");

ALTER TABLE "BarberRoles" ADD FOREIGN KEY ("barber_id") REFERENCES "Barbers" ("id");

ALTER TABLE "BarberRoles" ADD FOREIGN KEY ("barbershop_id") REFERENCES "BarberShops" ("id");

ALTER TABLE "BarberRoles" ADD FOREIGN KEY ("role_id") REFERENCES "Roles" ("id");

ALTER TABLE "BarberShops" ADD FOREIGN KEY ("barbershop_chain_id") REFERENCES "BarberShopChains" ("id");

ALTER TABLE "Barbers" ADD FOREIGN KEY ("gender_id") REFERENCES "Genders" ("id");

ALTER TABLE "BarberManagers" ADD FOREIGN KEY ("barber_id") REFERENCES "Barbers" ("id");

ALTER TABLE "BarberManagers" ADD FOREIGN KEY ("manager_id") REFERENCES "Barbers" ("id");

ALTER TABLE "SessionsBarber" ADD FOREIGN KEY ("barber_id") REFERENCES "Barbers" ("id");

ALTER TABLE "BarberShopServiceCategories" ADD FOREIGN KEY ("barbershop_chain_id") REFERENCES "BarberShopChains" ("id");

ALTER TABLE "BarberShopServiceCategories" ADD FOREIGN KEY ("barbershop_id") REFERENCES "BarberShops" ("id");

ALTER TABLE "BarberShopServiceCategories" ADD FOREIGN KEY ("service_category_id") REFERENCES "ServiceCategories" ("id");

ALTER TABLE "BarberShopServices" ADD FOREIGN KEY ("barbershop_category_id") REFERENCES "BarberShopServiceCategories" ("id");

ALTER TABLE "BarberShopServices" ADD FOREIGN KEY ("barbershop_chain_id") REFERENCES "BarberShopChains" ("id");

ALTER TABLE "BarberShopServices" ADD FOREIGN KEY ("barbershop_id") REFERENCES "BarberShops" ("id");

ALTER TABLE "BarberShopServices" ADD FOREIGN KEY ("gender_id") REFERENCES "Genders" ("id");

ALTER TABLE "SessionsCustomer" ADD FOREIGN KEY ("customer_id") REFERENCES "Customers" ("id");

ALTER TABLE "Appointments" ADD FOREIGN KEY ("barbershop_id") REFERENCES "BarberShops" ("id");

ALTER TABLE "Appointments" ADD FOREIGN KEY ("customer_id") REFERENCES "Customers" ("id");

ALTER TABLE "Appointments" ADD FOREIGN KEY ("barber_id") REFERENCES "Barbers" ("id");

CREATE TABLE "BarberShopServices_Appointments" (
  "BarberShopServices_id" uuid,
  "Appointments_service_id" uuid,
  PRIMARY KEY ("BarberShopServices_id", "Appointments_service_id")
);

ALTER TABLE "BarberShopServices_Appointments" ADD FOREIGN KEY ("BarberShopServices_id") REFERENCES "BarberShopServices" ("id");

ALTER TABLE "BarberShopServices_Appointments" ADD FOREIGN KEY ("Appointments_service_id") REFERENCES "Appointments" ("id");


ALTER TABLE "BarberShopReviews" ADD FOREIGN KEY ("customer_id") REFERENCES "Customers" ("id");

ALTER TABLE "BarberShopReviews" ADD FOREIGN KEY ("barbershop_id") REFERENCES "BarberShops" ("id");

ALTER TABLE "BarberReviews" ADD FOREIGN KEY ("barber_id") REFERENCES "Barbers" ("id");

ALTER TABLE "BarberReviews" ADD FOREIGN KEY ("customer_id") REFERENCES "Customers" ("id");
