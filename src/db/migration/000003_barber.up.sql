-- Create Barbers table
CREATE TABLE "Barbers" (
  "barber_id" serial PRIMARY KEY,
  "shop_id" INT,
  "barber_code" int,
  "gender_id" int2,
  "phone" varchar(15) UNIQUE NOT NULL,
  "nick_name" varchar(50) NOT NULL,
  "email" varchar(50) UNIQUE,
  "full_name" varchar(50) NOT NULL,
  "avatar_url" varchar(120),
  "barber_type_id" INT2 NOT NULL,
  "birth_date" DATE,
  "member_date" DATE,
  "create_at" timestamptz NOT NULL DEFAULT (now()),
  FOREIGN KEY ("shop_id") REFERENCES "BarberShops" ("shop_id"),
  FOREIGN KEY ("gender_id") REFERENCES "Genders" ("id"),
  FOREIGN KEY ("barber_type_id") REFERENCES "BarberTypes" ("id"),
  UNIQUE ("nick_name", "shop_id")
);

-- Create indexes on Barbers
CREATE INDEX ON "Barbers" ("phone");
CREATE INDEX ON "Barbers" ("email");
CREATE INDEX ON "Barbers" ("nick_name");

-- Create BarberRoles table
CREATE TABLE "BarberRoles" (
  "id" serial PRIMARY KEY,
  "barber_id" int NOT NULL,
  "barber_shop_id" int NOT NULL,
  "role_id" int2 NOT NULL,
  FOREIGN KEY ("barber_id") REFERENCES "Barbers" ("barber_id"),
  FOREIGN KEY ("barber_shop_id") REFERENCES "BarberShops" ("shop_id"),
  FOREIGN KEY ("role_id") REFERENCES "Roles" ("id"),
  UNIQUE ("barber_id", "barber_shop_id", "role_id")
);

-- Create indexes on BarberRoles
CREATE INDEX ON "BarberRoles" ("barber_id");
CREATE INDEX ON "BarberRoles" ("barber_shop_id");
CREATE INDEX ON "BarberRoles" ("role_id");


CREATE TABLE "Users" (
  "user_id" serial PRIMARY KEY,
  "phone" varchar(15) UNIQUE NOT NULL,
  "email" varchar(50) UNIQUE,
  "hashed_password" varchar(150),
  "is_active" BOOLEAN NOT NULL DEFAULT TRUE,
  "role_id" int2 NOT NULL,
  "barber_id" INT,
  "password_changed_at" timestamptz DEFAULT NULL,
  "created_at" timestamptz NOT NULL DEFAULT now(),
  FOREIGN KEY ("role_id") REFERENCES "Roles" ("id"),
  FOREIGN KEY ("barber_id") REFERENCES "Barbers" ("barber_id")
);
CREATE INDEX ON "Users" ("phone");
CREATE INDEX ON "Users" ("email");
CREATE INDEX ON "Users" ("role_id");
CREATE INDEX ON "Users" ("barber_id");

-- Create SessionsBarber table
CREATE TABLE "SessionsBarber" (
  "id" serial PRIMARY KEY,
  "barber_id" int NOT NULL,
  "refresh_token" varchar NOT NULL,
  "user_agent" varchar NOT NULL,
  "client_ip" varchar NOT NULL,
  "fcm_device" varchar NOT NULL,
  "is_blocked" bool NOT NULL DEFAULT false,  
  "expires_at" timestamptz NOT NULL,
  "create_at" timestamptz NOT NULL DEFAULT (now()),
  FOREIGN KEY ("barber_id") REFERENCES "Barbers" ("barber_id")
);

-- Create index on SessionsBarber
CREATE INDEX ON "SessionsBarber" ("barber_id");

-- Create OTPRequests table
CREATE TABLE "OTPRequests" (
  "id" serial PRIMARY KEY,
  "barber_id" int NOT NULL,
  "otp" varchar(6) NOT NULL,
  "requested_at" timestamptz NOT NULL DEFAULT (now()),
  "is_confirm" bool NOT NULL DEFAULT false,
  FOREIGN KEY ("barber_id") REFERENCES "Barbers" ("barber_id")
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