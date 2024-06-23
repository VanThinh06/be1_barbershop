
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
