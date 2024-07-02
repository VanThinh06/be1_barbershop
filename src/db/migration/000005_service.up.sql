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


-- Create ServiceItems table
CREATE TABLE "ServiceItems" (
  "id" uuid PRIMARY KEY DEFAULT (uuid_generate_v4()),
  "barber_shop_id" uuid NOT NULL,
  "category_id" int2 NOT NULL,
  "gender_id" int2 NOT NULL,
  "name" varchar(100) NOT NULL,
  "timer" int2 NOT NULL DEFAULT 0,
  "price" real NOT NULL DEFAULT 0,
  "discount_price" real DEFAULT 0, 
  "discount_start_time" timestamp, 
  "discount_end_time" timestamp, 
  "description" varchar(500) DEFAULT '', 
  "image_url" varchar(120)  DEFAULT '', 
  "is_active" BOOLEAN NOT NULL DEFAULT FALSE,
  FOREIGN KEY ("barber_shop_id") REFERENCES "BarberShops" ("id"),
  FOREIGN KEY ("category_id") REFERENCES "ServiceCategories" ("id"),
  FOREIGN KEY ("gender_id") REFERENCES "Genders" ("id")
);

-- Create indexes
CREATE INDEX ON "ServiceItems" ("barber_shop_id");
CREATE INDEX ON "ServiceItems" ("category_id");
CREATE INDEX ON "ServiceItems" ("gender_id");
CREATE UNIQUE INDEX ON "ServiceItems" ("barber_shop_id", "category_id", "name");


-- Create ServicePackages table
CREATE TABLE "ServicePackages" (
  "id" uuid PRIMARY KEY DEFAULT (uuid_generate_v4()),
  "barber_shop_id" uuid NOT NULL,
  "name" varchar(100) NOT NULL,
  "gender_id" int2 NOT NULL,
  "timer" int2 NOT NULL DEFAULT 0,
  "price" real NOT NULL DEFAULT 0,
  "discount_price" real DEFAULT 0, 
  "discount_start_time" timestamp, 
  "discount_end_time" timestamp, 
  "description" varchar(500) DEFAULT '', 
  "image_url" varchar(120) DEFAULT '', 
  "is_active" BOOLEAN NOT NULL DEFAULT FALSE,
  FOREIGN KEY ("barber_shop_id") REFERENCES "BarberShops" ("id"),
  FOREIGN KEY ("gender_id") REFERENCES "Genders" ("id")
);

-- Create indexes
CREATE INDEX ON "ServicePackages" ("barber_shop_id");
CREATE INDEX ON "ServicePackages" ("gender_id");
CREATE UNIQUE INDEX ON "ServicePackages" ("barber_shop_id", "name");


-- Create ServicePackageItems table
CREATE TABLE "ServicePackageItems" (
  "id" uuid PRIMARY KEY DEFAULT (uuid_generate_v4()),
  "service_package_id" uuid NOT NULL,
  "service_item_id" uuid NOT NULL,
  FOREIGN KEY ("service_package_id") REFERENCES "ServicePackages" ("id"),
  FOREIGN KEY ("service_item_id") REFERENCES "ServiceItems" ("id")
);

-- Create indexes
CREATE INDEX ON "ServicePackageItems" ("service_package_id");
CREATE INDEX ON "ServicePackageItems" ("service_item_id");
CREATE UNIQUE INDEX ON "ServicePackageItems" ("service_package_id", "service_item_id");

-- Create view_service_packages view
CREATE VIEW "view_service_packages" AS
SELECT 
  cs.id,
  cs.barber_shop_id,
  cs.gender_id AS combo_service_gender,
  cs.name AS combo_service_name,
  cs.timer AS combo_service_timer,
  cs.price AS combo_service_price,
  cs.discount_price AS combo_service_discount_price,
  cs.discount_start_time AS combo_service_discount_start_time,
  cs.discount_end_time AS combo_service_discount_end_time,
  cs.description AS combo_service_description,
  cs.image_url AS combo_service_image_url,
  cs.is_active AS combo_service_is_active,
  array_agg(csi.service_item_id) AS service_item_ids
FROM 
  "ServicePackageItems" csi
JOIN 
  "ServicePackages" cs ON csi.service_package_id = cs.id
GROUP BY 
  cs.id, cs.barber_shop_id, cs.gender_id, cs.name, cs.timer, cs.price, cs.discount_price, 
  cs.discount_start_time, cs.discount_end_time, cs.description, cs.image_url, cs.is_active;


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
