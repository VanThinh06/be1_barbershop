-- Create uuid-ossp extension
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

-- Create Provinces table
CREATE TABLE "Provinces" (
  "id" smallserial PRIMARY KEY,
  "name" varchar(50) UNIQUE NOT NULL
);

-- Create Districts table
CREATE TABLE "Districts" (
  "id" smallserial PRIMARY KEY,
  "name" varchar(50) NOT NULL,
  "province_id" int2 NOT NULL,
  FOREIGN KEY ("province_id") REFERENCES "Provinces" ("id")
);

-- Create index on province_id
CREATE INDEX ON "Districts" ("province_id");

-- Create Wards table
CREATE TABLE "Wards" (
  "id" smallserial PRIMARY KEY,
  "name" varchar(50) NOT NULL,
  "district_id" int2 NOT NULL,
  FOREIGN KEY ("district_id") REFERENCES "Districts" ("id")
);

-- Create index on district_id
CREATE INDEX ON "Wards" ("district_id");

-- Create Genders table
CREATE TABLE "Genders" (
  "id" smallserial PRIMARY KEY,
  "name" varchar(10) UNIQUE NOT NULL
);

-- Create Roles table
CREATE TABLE "Roles" (
  "id" smallserial PRIMARY KEY,
  "name" varchar(100) UNIQUE NOT NULL,
  "type" varchar(50) NOT NULL
);

-- Create index on name
CREATE INDEX ON "Roles" ("name");

-- Create Permissions table
CREATE TABLE "Permissions" (
  "id" smallserial PRIMARY KEY,
  "name" varchar(100) UNIQUE NOT NULL,
  "module" varchar(100) NOT NULL
);

-- Create index on name
CREATE INDEX ON "Permissions" ("name");

-- Create RolePermissions table
CREATE TABLE "RolePermissions" (
  "id" smallserial PRIMARY KEY, 
  "role_id" int2 NOT NULL,
  "permission_id" int2 NOT NULL,
  FOREIGN KEY ("role_id") REFERENCES "Roles" ("id"),
  FOREIGN KEY ("permission_id") REFERENCES "Permissions" ("id")
);

-- Create indexes on RolePermissions
CREATE INDEX ON "RolePermissions" ("role_id");
CREATE INDEX ON "RolePermissions" ("permission_id");

-- Create BarberShopChains table
CREATE TABLE "BarberShopChains" (
  "barber_shop_chain_id" uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
  "name" varchar(100) NOT NULL,
  "chain_identifier" varchar(50) UNIQUE NOT NULL,
  "founder" varchar(50) NOT NULL,
  "founding_date" date NOT NULL,
  "website" varchar(150)
);

-- Create indexes on BarberShopChains
CREATE INDEX ON "BarberShopChains" ("name");
CREATE INDEX ON "BarberShopChains" ("chain_identifier");

-- Create BarberTypes table
CREATE TABLE "BarberTypes" (
  "id" smallserial PRIMARY KEY,
  "name" VARCHAR(100) NOT NULL UNIQUE
);

-- Create indexes on BarberTypes
CREATE INDEX ON "BarberTypes" ("name");
