
CREATE TABLE "BarberShops" (
  "barber_shop_id" uuid PRIMARY KEY DEFAULT (uuid_generate_v4()),
  "barber_shop_chain_id" uuid,
  "code" varchar(20) UNIQUE NOT NULL,
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
  "create_at" timestamptz NOT NULL DEFAULT (now()),
  FOREIGN KEY ("barber_shop_chain_id") REFERENCES "BarberShopChains" ("barber_shop_chain_id"),
  FOREIGN KEY ("province_id") REFERENCES "Provinces" ("id"),
  FOREIGN KEY ("district_id") REFERENCES "Districts" ("id"),
  FOREIGN KEY ("ward_id") REFERENCES "Wards" ("id")
);

CREATE INDEX ON "BarberShops" ("barber_shop_chain_id");
CREATE INDEX ON "BarberShops" ("name");
CREATE INDEX ON "BarberShops" ("province_id");
CREATE INDEX ON "BarberShops" ("district_id");
CREATE INDEX ON "BarberShops" ("ward_id");
CREATE INDEX ON "BarberShops" ("code");


CREATE VIEW "view_barber_shops" AS
SELECT 
  bs.id,
  bs.barber_shop_chain_id,
  bs.name,
  p.name AS province_name,
  d.name AS district_name,
  w.name AS ward_name,
  bs.specific_location,
  bs.phone,
  bs.email,
  bs.website_url,
  bs.coordinates,
  bs.avatar_url,
  bs.cover_photo_url,
  bs.facade_photo_url,
  bs.representative_name,
  bs.citizen_id,
  bs.representative_phone,
  bs.representative_email,
  bs.representative_phone_other,
  bs.start_time_monday,
  bs.end_time_monday,
  bs.start_time_tuesday,
  bs.end_time_tuesday,
  bs.start_time_wednesday,
  bs.end_time_wednesday,
  bs.start_time_thursday,
  bs.end_time_thursday,
  bs.start_time_friday,
  bs.end_time_friday,
  bs.start_time_saturday,
  bs.end_time_saturday,
  bs.start_time_sunday,
  bs.end_time_sunday,
  bs.interval_scheduler,
  bs.is_main_branch,
  bs.is_reputation,
  bs.is_verified,
  bs.create_at
FROM "BarberShops" bs
JOIN "Provinces" p ON bs.province_id = p.id
JOIN "Districts" d ON bs.district_id = d.id
JOIN "Wards" w ON bs.ward_id = w.id;
