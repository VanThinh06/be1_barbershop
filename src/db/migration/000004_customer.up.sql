-- Create Customers table
CREATE TABLE "Customers" (
  "id" uuid PRIMARY KEY DEFAULT (uuid_generate_v4()),
  "name" varchar(50) NOT NULL,
  "email" varchar(50) UNIQUE NOT NULL,
  "phone" varchar(15) UNIQUE,
  "gender_id" int2,
  "hashed_password" varchar(150),
  "avatar" varchar(120),
  "is_social_auth" bool DEFAULT false,
  "password_changed_at" timestamptz DEFAULT NULL,
  "create_at" timestamptz NOT NULL DEFAULT (now()),
  FOREIGN KEY ("gender_id") REFERENCES "Genders" ("id")
);

-- Create indexes on Customers
CREATE INDEX ON "Customers" ("phone");
CREATE INDEX ON "Customers" ("email");

-- Create SessionsCustomer table
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

-- Create index on SessionsCustomer
CREATE INDEX ON "SessionsCustomer" ("customer_id");

-- Create OTPRequestsCustomer table
CREATE TABLE "OTPRequestsCustomer" (
  "id" uuid PRIMARY KEY DEFAULT (uuid_generate_v4()),
  "customer_id" uuid NOT NULL,
  "otp" varchar(6) NOT NULL,
  "requested_at" timestamptz NOT NULL DEFAULT (now()),
  "is_confirm" bool NOT NULL DEFAULT false,
  FOREIGN KEY ("customer_id") REFERENCES "Customers" ("id")
);

-- Create index on OTPRequestsCustomer
CREATE INDEX ON "OTPRequestsCustomer" ("customer_id");
