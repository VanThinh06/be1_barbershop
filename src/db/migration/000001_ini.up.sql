CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE "Provinces" (
  "id" serial2 PRIMARY KEY,
  "name" varchar(50) UNIQUE NOT NULL
);

CREATE TABLE "Districts" (
  "id" serial2 PRIMARY KEY,
  "name" varchar(50) NOT NULL,
  "province_id" int2 NOT NULL,
  FOREIGN KEY ("province_id") REFERENCES "Provinces" ("id")
);

CREATE INDEX ON "Districts" ("province_id");

CREATE TABLE "Wards" (
  "id" serial2 PRIMARY KEY,
  "name" varchar(50) NOT NULL,
  "district_id" int2 NOT NULL,
  FOREIGN KEY ("district_id") REFERENCES "Districts" ("id")
);

CREATE INDEX ON "Wards" ("district_id");

CREATE TABLE "Genders" (
  "id" serial2 PRIMARY KEY,
  "name" varchar(10) UNIQUE NOT NULL
);

CREATE TABLE "Roles" (
  "id" serial2 PRIMARY KEY,
  "name" varchar(100) UNIQUE NOT NULL,
  "type" varchar(50)
);

CREATE TABLE "Permissions" (
  "id" serial2 PRIMARY KEY,
  "name" varchar(100) UNIQUE NOT NULL,
  "description" varchar(255)
);

CREATE TABLE "RolePermissions" (
  "id" serial2 PRIMARY KEY, 
  "role_id" int2 NOT NULL,
  "permission_id" int2 NOT NULL,
  FOREIGN KEY ("role_id") REFERENCES "Roles" ("id"),
  FOREIGN KEY ("permission_id") REFERENCES "Permissions" ("id")
);

CREATE INDEX ON "RolePermissions" ("role_id");
CREATE INDEX ON "RolePermissions" ("permission_id");

CREATE TABLE "BarberShopChains" (
  "id" uuid PRIMARY KEY DEFAULT (uuid_generate_v4()),
  "name" varchar(100) NOT NULL,
  "chain_identifier" varchar(50) UNIQUE NOT NULL,
  "founder" varchar(50) NOT NULL,
  "founding_date" date NOT NULL,
  "website" varchar(150)
);

CREATE INDEX ON "BarberShopChains" ("name");

CREATE TABLE "BarberShops" (
  "id" uuid PRIMARY KEY DEFAULT (uuid_generate_v4()),
  "barber_shop_chain_id" uuid,
  "name" varchar(200) NOT NULL,
  "province_id" int2 NOT NULL,
  "district_id" int2 NOT NULL,
  "ward_id" int2 NOT NULL,
  "specific_location" varchar(200) NOT NULL,
  "phone" varchar(15) NOT NULL,
  "email" varchar(100) NOT NULL,
  "website_url" varchar(150),
  "coordinates" GEOGRAPHY(Point, 4326) NOT NULL,
  "avatar_url" varchar(120) NOT NULL,
  "cover_photo_url" varchar(120) NOT NULL,
  "facade_photo_url" varchar(120) NOT NULL,
  "representative_name" varchar(100) NOT NULL,
  "citizen_id" varchar(12) NOT NULL,
  "representative_phone" varchar(15) NOT NULL,
  "representative_email" varchar(100) NOT NULL,
  "representative_phone_other" varchar(15),
  "start_time_monday" TIME NOT NULL DEFAULT '00:00:00'::TIME,
  "end_time_monday" TIME NOT NULL DEFAULT '00:00:00'::TIME,
  "start_time_tuesday" TIME NOT NULL DEFAULT '00:00:00'::TIME,
  "end_time_tuesday" TIME NOT NULL DEFAULT '00:00:00'::TIME,
  "start_time_wednesday" TIME NOT NULL DEFAULT '00:00:00'::TIME,
  "end_time_wednesday" TIME NOT NULL DEFAULT '00:00:00'::TIME,
  "start_time_thursday" TIME NOT NULL DEFAULT '00:00:00'::TIME,
  "end_time_thursday" TIME NOT NULL DEFAULT '00:00:00'::TIME,
  "start_time_friday" TIME NOT NULL DEFAULT '00:00:00'::TIME,
  "end_time_friday" TIME NOT NULL DEFAULT '00:00:00'::TIME,
  "start_time_saturday" TIME NOT NULL DEFAULT '00:00:00'::TIME,
  "end_time_saturday" TIME NOT NULL DEFAULT '00:00:00'::TIME,
  "start_time_sunday" TIME NOT NULL DEFAULT '00:00:00'::TIME,
  "end_time_sunday" TIME NOT NULL DEFAULT '00:00:00'::TIME,
  "interval_scheduler" int2 NOT NULL DEFAULT 30,
  "is_main_branch" bool NOT NULL DEFAULT false,
  "is_reputation" bool NOT NULL DEFAULT false,
  "is_verified" bool NOT NULL DEFAULT false,
  "default_employee_password" varchar(25),
  "create_at" timestamptz NOT NULL DEFAULT (now()),
  FOREIGN KEY ("barber_shop_chain_id") REFERENCES "BarberShopChains" ("id"),
  FOREIGN KEY ("province_id") REFERENCES "Provinces" ("id"),
  FOREIGN KEY ("district_id") REFERENCES "Districts" ("id"),
  FOREIGN KEY ("ward_id") REFERENCES "Wards" ("id")
);

CREATE INDEX ON "BarberShops" ("barber_shop_chain_id");
CREATE INDEX ON "BarberShops" ("name");
CREATE INDEX ON "BarberShops" ("province_id");
CREATE INDEX ON "BarberShops" ("district_id");
CREATE INDEX ON "BarberShops" ("ward_id");

CREATE TABLE "Barbers" (
  "id" uuid PRIMARY KEY DEFAULT (uuid_generate_v4()),
  "gender_id" int2,
  "phone" varchar(15) UNIQUE NOT NULL,
  "nick_name" varchar(50) UNIQUE NOT NULL,
  "email" varchar(50) UNIQUE,
  "hashed_password" varchar(150),
  "full_name" varchar(50) NOT NULL,
  "haircut" bool NOT NULL DEFAULT FALSE,
  "work_status" BOOLEAN NOT NULL DEFAULT TRUE,
  "avatar_url" varchar(120),
  "password_changed_at" timestamptz NOT NULL DEFAULT '0001-01-01 00:00:00Z',
  "create_at" timestamptz NOT NULL DEFAULT (now()),
  FOREIGN KEY ("gender_id") REFERENCES "Genders" ("id")
);

CREATE INDEX ON "Barbers" ("phone");
CREATE INDEX ON "Barbers" ("email");
CREATE INDEX ON "Barbers" ("nick_name");

CREATE TABLE "BarberRoles" (
  "id" uuid PRIMARY KEY DEFAULT (uuid_generate_v4()),
  "barber_id" uuid NOT NULL,
  "barber_shop_id" uuid NOT NULL,
  "role_id" int2 NOT NULL,
  FOREIGN KEY ("barber_id") REFERENCES "Barbers" ("id"),
  FOREIGN KEY ("barber_shop_id") REFERENCES "BarberShops" ("id"),
  FOREIGN KEY ("role_id") REFERENCES "Roles" ("id")
);

CREATE INDEX ON "BarberRoles" ("barber_id");
CREATE INDEX ON "BarberRoles" ("barber_shop_id");
CREATE INDEX ON "BarberRoles" ("role_id");

CREATE TABLE "ServiceCategories" (
  "id" serial2 PRIMARY KEY,
  "name" varchar(50) UNIQUE NOT NULL,
  "barber_shop_id" uuid,
  FOREIGN KEY ("barber_shop_id") REFERENCES "BarberShops" ("id")
);

CREATE INDEX ON "ServiceCategories" ("name");
CREATE INDEX ON "ServiceCategories" ("barber_shop_id");

CREATE TABLE "BarberShopServices" (
  "id" uuid PRIMARY KEY DEFAULT (uuid_generate_v4()),
  "barber_shop_id" uuid NOT NULL,
  "category_id" int2 NOT NULL,
  "gender_id" int2 NOT NULL,
  "name" varchar(100) NOT NULL,
  "timer" int2 NOT NULL DEFAULT 0,
  "price" real NOT NULL DEFAULT 0,
  "discount_price" real, 
  "discount_start_time" timestamp, 
  "discount_end_time" timestamp, 
  "description" varchar(500),
  "image_url" varchar(120),
  "combo_services" text[],
  "is_active" BOOLEAN NOT NULL DEFAULT FALSE,
  FOREIGN KEY ("barber_shop_id") REFERENCES "BarberShops" ("id"),
  FOREIGN KEY ("category_id") REFERENCES "ServiceCategories" ("id"),
  FOREIGN KEY ("gender_id") REFERENCES "Genders" ("id")
);

CREATE INDEX ON "BarberShopServices" ("category_id");
CREATE INDEX ON "BarberShopServices" ("gender_id");
CREATE UNIQUE INDEX ON "BarberShopServices" ("category_id", "name", "gender_id");

CREATE TABLE "SessionsBarber" (
  "id" uuid PRIMARY KEY DEFAULT (uuid_generate_v4()),
  "barber_id" uuid NOT NULL,
  "refresh_token" varchar NOT NULL,
  "user_agent" varchar NOT NULL,
  "client_ip" varchar NOT NULL,
  "fcm_device" varchar NOT NULL,
  "is_blocked" bool NOT NULL DEFAULT false,  
  "expires_at" timestamptz NOT NULL,
  "create_at" timestamptz NOT NULL DEFAULT (now()),
  FOREIGN KEY ("barber_id") REFERENCES "Barbers" ("id")
);

CREATE INDEX ON "SessionsBarber" ("barber_id");

CREATE TABLE "Customers" (
  "id" uuid PRIMARY KEY DEFAULT (uuid_generate_v4()),
  "name" varchar(50) NOT NULL,
  "email" varchar(50) UNIQUE NOT NULL,
  "phone" varchar(15) UNIQUE,
  "gender_id" int2,
  "hashed_password" varchar(50),
  "avatar" varchar(120),
  "is_social_auth" bool DEFAULT false,
  "password_changed_at" timestamptz NOT NULL DEFAULT '0001-01-01 00:00:00Z',
  "create_at" timestamptz NOT NULL DEFAULT (now()),
  FOREIGN KEY ("gender_id") REFERENCES "Genders" ("id")
);

CREATE INDEX ON "Customers" ("phone");
CREATE INDEX ON "Customers" ("email");

CREATE TABLE "SessionsCustomer" (
  "id" uuid PRIMARY KEY DEFAULT (uuid_generate_v4()),
  "customer_id" uuid NOT NULL,
  "refresh_token" varchar NOT NULL,
  "user_agent" varchar NOT NULL,
  "client_ip" varchar NOT NULL,
  "fcm_device" varchar NOT NULL,
  "coordinates" GEOGRAPHY(Point, 4326),
  "is_blocked" bool NOT NULL DEFAULT false,
  "expires_at" timestamptz NOT NULL,
  "create_at" timestamptz NOT NULL DEFAULT (now()),
  FOREIGN KEY ("customer_id") REFERENCES "Customers" ("id")
);

CREATE INDEX ON "SessionsCustomer" ("customer_id");

CREATE TABLE "Appointments" (
  "id" uuid PRIMARY KEY DEFAULT (uuid_generate_v4()),
  "barber_shop_id" uuid NOT NULL,
  "customer_id" uuid NOT NULL,
  "barber_id" uuid NOT NULL,
  "appointment_date_time" timestamptz NOT NULL,
  "status" int2 NOT NULL,
  "create_at" timestamptz NOT NULL DEFAULT (now()),
  "update_at" timestamptz NOT NULL DEFAULT '0001-01-01 00:00:00Z',
  FOREIGN KEY ("barber_shop_id") REFERENCES "BarberShops" ("id"),
  FOREIGN KEY ("customer_id") REFERENCES "Customers" ("id"),
  FOREIGN KEY ("barber_id") REFERENCES "Barbers" ("id")
);

CREATE INDEX ON "Appointments" ("barber_id");
CREATE INDEX ON "Appointments" ("customer_id");
CREATE INDEX ON "Appointments" ("barber_shop_id");
CREATE UNIQUE INDEX ON "Appointments" ("barber_id", "appointment_date_time");

CREATE OR REPLACE FUNCTION check_appointment_conflict()
RETURNS TRIGGER AS $$
BEGIN
  IF EXISTS (
    SELECT 1
    FROM "Appointments"
    WHERE "barber_shop_id" = NEW."barber_shop_id"
      AND "barber_id" = NEW."barber_id"
      AND (
        (NEW."appointment_date_time" + NEW."timer" * interval '1 minute') BETWEEN "appointment_date_time" AND "appointment_date_time" + NEW."timer" * interval '1 minute'
        OR NEW."appointment_date_time" BETWEEN "appointment_date_time" AND "appointment_date_time" + NEW."timer" * interval '1 minute'
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
  "barber_shop_id" uuid NOT NULL,
  "rating" int2 NOT NULL,
  "comment" varchar(240),
  "create_at" timestamptz NOT NULL DEFAULT (now()),
  "update_at" timestamptz NOT NULL DEFAULT '0001-01-01 00:00:00Z',
  FOREIGN KEY ("customer_id") REFERENCES "Customers" ("id"),
  FOREIGN KEY ("barber_shop_id") REFERENCES "BarberShops" ("id")
);

CREATE INDEX ON "BarberShopReviews" ("customer_id");
CREATE INDEX ON "BarberShopReviews" ("barber_shop_id");

CREATE TABLE "BarberReviews" (
  "id" uuid PRIMARY KEY DEFAULT (uuid_generate_v4()),
  "barber_id" uuid NOT NULL,
  "customer_id" uuid NOT NULL,
  "rating" int2 NOT NULL,
  "comment" varchar(240),
  "create_at" timestamptz NOT NULL DEFAULT (now()),
  "update_at" timestamptz NOT NULL DEFAULT '0001-01-01 00:00:00Z',
  FOREIGN KEY ("barber_id") REFERENCES "Barbers" ("id"),
  FOREIGN KEY ("customer_id") REFERENCES "Customers" ("id")
);

CREATE INDEX ON "BarberReviews" ("barber_id");
CREATE INDEX ON "BarberReviews" ("customer_id");

CREATE TABLE "BarberShopServices_Appointments" (
  "BarberShopServices_id" uuid,
  "AppointmentsService_id" uuid,
  PRIMARY KEY ("BarberShopServices_id", "AppointmentsService_id"),
  FOREIGN KEY ("BarberShopServices_id") REFERENCES "BarberShopServices" ("id"),
  FOREIGN KEY ("AppointmentsService_id") REFERENCES "Appointments" ("id")
);

CREATE TABLE "OTPRequests" (
  "id" uuid PRIMARY KEY DEFAULT (uuid_generate_v4()),
  "barber_id" uuid NOT NULL,
  "otp" varchar(6) NOT NULL,
  "requested_at" timestamptz NOT NULL DEFAULT (now()),
  "is_confirm" bool NOT NULL DEFAULT false,
  FOREIGN KEY ("barber_id") REFERENCES "Barbers" ("id")
);

CREATE INDEX ON "OTPRequests" ("barber_id");
