-- Create ServiceCategories table
CREATE TABLE "ServiceCategories" (
  "id" serial2 PRIMARY KEY,
  "name" varchar(50) NOT NULL,
  "barber_shop_id" uuid,
  FOREIGN KEY ("barber_shop_id") REFERENCES "BarberShops" ("id")
);

-- Create indexes
CREATE UNIQUE INDEX ON "ServiceCategories" ("name", "barber_shop_id");
CREATE INDEX ON "ServiceCategories" ("name");
CREATE INDEX ON "ServiceCategories" ("barber_shop_id");


-- Create CategoryPositions table
CREATE TABLE "CategoryPositions" (
  "barber_shop_id" uuid NOT NULL,
  "category_id" serial2 NOT NULL,
  "position" int2 NOT NULL,
  "visible" boolean NOT NULL DEFAULT true,
  PRIMARY KEY ("barber_shop_id", "category_id"),
  FOREIGN KEY ("barber_shop_id") REFERENCES "BarberShops" ("id"),
  FOREIGN KEY ("category_id") REFERENCES "ServiceCategories" ("id")
);

-- Create indexes
CREATE INDEX  ON "CategoryPositions" ("barber_shop_id");
CREATE INDEX  ON "CategoryPositions" ("category_id");


-- Create BarberShopServices table
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
  "is_active" BOOLEAN NOT NULL DEFAULT FALSE,
  FOREIGN KEY ("barber_shop_id") REFERENCES "BarberShops" ("id"),
  FOREIGN KEY ("category_id") REFERENCES "ServiceCategories" ("id"),
  FOREIGN KEY ("gender_id") REFERENCES "Genders" ("id")
);

-- Create indexes
CREATE INDEX ON "BarberShopServices" ("barber_shop_id");
CREATE INDEX ON "BarberShopServices" ("category_id");
CREATE INDEX ON "BarberShopServices" ("gender_id");
CREATE UNIQUE INDEX ON "BarberShopServices" ("barber_shop_id", "category_id", "name");


-- Create ComboServices table
CREATE TABLE "ComboServices" (
  "id" uuid PRIMARY KEY DEFAULT (uuid_generate_v4()),
  "barber_shop_id" uuid NOT NULL,
  "name" varchar(100) NOT NULL,
  "gender_id" int2 NOT NULL,
  "timer" int2 NOT NULL DEFAULT 0,
  "price" real NOT NULL DEFAULT 0,
  "discount_price" real, 
  "discount_start_time" timestamp, 
  "discount_end_time" timestamp, 
  "description" varchar(500),
  "image_url" varchar(120),
  "is_active" BOOLEAN NOT NULL DEFAULT FALSE,
  FOREIGN KEY ("barber_shop_id") REFERENCES "BarberShops" ("id"),
  FOREIGN KEY ("gender_id") REFERENCES "Genders" ("id")
);

-- Create indexes
CREATE INDEX ON "ComboServices" ("barber_shop_id");
CREATE INDEX ON "ComboServices" ("gender_id");
CREATE UNIQUE INDEX ON "ComboServices" ("barber_shop_id", "name");


-- Create ComboServiceItems table
CREATE TABLE "ComboServiceItems" (
  "id" uuid PRIMARY KEY DEFAULT (uuid_generate_v4()),
  "combo_service_id" uuid NOT NULL,
  "barber_shop_service_id" uuid NOT NULL,
  FOREIGN KEY ("combo_service_id") REFERENCES "ComboServices" ("id"),
  FOREIGN KEY ("barber_shop_service_id") REFERENCES "BarberShopServices" ("id")
);

-- Create indexes
CREATE INDEX ON "ComboServiceItems" ("combo_service_id");
CREATE INDEX ON "ComboServiceItems" ("barber_shop_service_id");



-- CREATE TABLE "BarberShopServices_Appointments" (
--   "BarberShopServices_id" uuid,
--   "AppointmentsService_id" uuid,
--   PRIMARY KEY ("BarberShopServices_id", "AppointmentsService_id"),
--   FOREIGN KEY ("BarberShopServices_id") REFERENCES "BarberShopServices" ("id"),
--   FOREIGN KEY ("AppointmentsService_id") REFERENCES "Appointments" ("id")
-- );


-- CREATE TABLE "BarberShopReviews" (
--   "id" uuid PRIMARY KEY DEFAULT (uuid_generate_v4()),
--   "customer_id" uuid NOT NULL,
--   "barber_shop_id" uuid NOT NULL,
--   "rating" int2 NOT NULL,
--   "comment" varchar(240),
--   "create_at" timestamptz NOT NULL DEFAULT (now()),
--   "update_at" timestamptz NOT NULL DEFAULT '0001-01-01 00:00:00Z',
--   FOREIGN KEY ("customer_id") REFERENCES "Customers" ("id"),
--   FOREIGN KEY ("barber_shop_id") REFERENCES "BarberShops" ("id")
-- );

-- CREATE INDEX ON "BarberShopReviews" ("customer_id");
-- CREATE INDEX ON "BarberShopReviews" ("barber_shop_id");

-- CREATE TABLE "BarberReviews" (
--   "id" uuid PRIMARY KEY DEFAULT (uuid_generate_v4()),
--   "barber_id" uuid NOT NULL,
--   "customer_id" uuid NOT NULL,
--   "rating" int2 NOT NULL,
--   "comment" varchar(240),
--   "create_at" timestamptz NOT NULL DEFAULT (now()),
--   "update_at" timestamptz NOT NULL DEFAULT '0001-01-01 00:00:00Z',
--   FOREIGN KEY ("barber_id") REFERENCES "Barbers" ("id"),
--   FOREIGN KEY ("customer_id") REFERENCES "Customers" ("id")
-- );

-- CREATE INDEX ON "BarberReviews" ("barber_id");
-- CREATE INDEX ON "BarberReviews" ("customer_id");


-- CREATE TABLE "Appointments" (
--   "id" uuid PRIMARY KEY DEFAULT (uuid_generate_v4()),
--   "barber_shop_id" uuid NOT NULL,
--   "customer_id" uuid NOT NULL,
--   "barber_id" uuid NOT NULL,
--   "appointment_date_time" timestamptz NOT NULL,
--   "status" int2 NOT NULL,
--   "create_at" timestamptz NOT NULL DEFAULT (now()),
--   "update_at" timestamptz NOT NULL DEFAULT '0001-01-01 00:00:00Z',
--   FOREIGN KEY ("barber_shop_id") REFERENCES "BarberShops" ("id"),
--   FOREIGN KEY ("customer_id") REFERENCES "Customers" ("id"),
--   FOREIGN KEY ("barber_id") REFERENCES "Barbers" ("id")
-- );

-- CREATE INDEX ON "Appointments" ("barber_id");
-- CREATE INDEX ON "Appointments" ("customer_id");
-- CREATE INDEX ON "Appointments" ("barber_shop_id");
-- CREATE UNIQUE INDEX ON "Appointments" ("barber_id", "appointment_date_time");

-- CREATE OR REPLACE FUNCTION check_appointment_conflict()
-- RETURNS TRIGGER AS $$
-- BEGIN
--   IF EXISTS (
--     SELECT 1
--     FROM "Appointments"
--     WHERE "barber_shop_id" = NEW."barber_shop_id"
--       AND "barber_id" = NEW."barber_id"
--       AND (
--         (NEW."appointment_date_time" + NEW."timer" * interval '1 minute') BETWEEN "appointment_date_time" AND "appointment_date_time" + NEW."timer" * interval '1 minute'
--         OR NEW."appointment_date_time" BETWEEN "appointment_date_time" AND "appointment_date_time" + NEW."timer" * interval '1 minute'
--       )
--   ) THEN
--     RAISE EXCEPTION 'Appointment conflict: Another appointment exists in this time slot.';
--   END IF;
--   RETURN NEW;
-- END;
-- $$ LANGUAGE plpgsql;

-- CREATE TRIGGER check_appointment_conflict_trigger
-- BEFORE INSERT ON "Appointments"
-- FOR EACH ROW EXECUTE FUNCTION check_appointment_conflict();