-- Create Barbers table
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
  "default_password_encrypted" TEXT,
  "password_changed_at" timestamptz DEFAULT NULL,
  "create_at" timestamptz NOT NULL DEFAULT (now()),
  FOREIGN KEY ("gender_id") REFERENCES "Genders" ("id")
);

-- Create indexes on Barbers
CREATE INDEX ON "Barbers" ("phone");
CREATE INDEX ON "Barbers" ("email");
CREATE INDEX ON "Barbers" ("nick_name");

-- Create BarberRoles table
CREATE TABLE "BarberRoles" (
  "id" uuid PRIMARY KEY DEFAULT (uuid_generate_v4()),
  "barber_id" uuid NOT NULL,
  "barber_shop_id" uuid NOT NULL,
  "role_id" int2 NOT NULL,
  FOREIGN KEY ("barber_id") REFERENCES "Barbers" ("id"),
  FOREIGN KEY ("barber_shop_id") REFERENCES "BarberShops" ("id"),
  FOREIGN KEY ("role_id") REFERENCES "Roles" ("id"),
  UNIQUE ("barber_id", "barber_shop_id", "role_id")
);

-- Create indexes on BarberRoles
CREATE INDEX ON "BarberRoles" ("barber_id");
CREATE INDEX ON "BarberRoles" ("barber_shop_id");
CREATE INDEX ON "BarberRoles" ("role_id");

-- Create SessionsBarber table
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

-- Create index on SessionsBarber
CREATE INDEX ON "SessionsBarber" ("barber_id");

-- Create OTPRequests table
CREATE TABLE "OTPRequests" (
  "id" uuid PRIMARY KEY DEFAULT (uuid_generate_v4()),
  "barber_id" uuid NOT NULL,
  "otp" varchar(6) NOT NULL,
  "requested_at" timestamptz NOT NULL DEFAULT (now()),
  "is_confirm" bool NOT NULL DEFAULT false,
  FOREIGN KEY ("barber_id") REFERENCES "Barbers" ("id")
);

-- Create or replace function to clean old OTP requests
CREATE OR REPLACE FUNCTION clean_old_otp_requests()
RETURNS TRIGGER AS $$
BEGIN
    -- Xóa OTP không thuộc về ngày hiện tại
    DELETE FROM "OTPRequests"
    WHERE "barber_id" = NEW."barber_id" 
      AND "requested_at"::date <> CURRENT_DATE;

    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

-- Create trigger to invoke the function before insert
CREATE TRIGGER before_insert_otp_request
BEFORE INSERT ON "OTPRequests"
FOR EACH ROW
EXECUTE FUNCTION clean_old_otp_requests();