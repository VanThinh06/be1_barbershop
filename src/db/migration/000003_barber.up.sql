
CREATE TABLE "Users" (
  "user_id" SERIAL PRIMARY KEY,
  "phone" VARCHAR(15) UNIQUE NOT NULL,
  "email" VARCHAR(50) UNIQUE,
  "hashed_password" VARCHAR(150) NOT NULL,
  "is_active" BOOLEAN NOT NULL DEFAULT TRUE,
  "role_id" SMALLINT NOT NULL,
  "password_changed_at" TIMESTAMPTZ DEFAULT NULL,
  "created_at" TIMESTAMPTZ NOT NULL DEFAULT NOW(),
  FOREIGN KEY ("role_id") REFERENCES "Roles" ("id")
);

CREATE INDEX ON "Users" ("role_id");

-- Create Barbers table
CREATE TABLE "Barbers" (
  "barber_id" SERIAL PRIMARY KEY,
  "shop_id" INT,
  "user_id" INT,
  "barber_code" INT,
  "gender_id" SMALLINT,
  "phone" VARCHAR(15) UNIQUE NOT NULL,
  "nick_name" VARCHAR(50),
  "email" VARCHAR(50) UNIQUE,
  "full_name" VARCHAR(50) NOT NULL,
  "avatar_url" VARCHAR(120),
  "barber_type_id" SMALLINT NOT NULL,
  "birth_date" DATE,
  "member_date" DATE,
  "created_at" TIMESTAMPTZ NOT NULL DEFAULT NOW(),
  FOREIGN KEY ("shop_id") REFERENCES "BarberShops" ("shop_id"),
  FOREIGN KEY ("gender_id") REFERENCES "Genders" ("id"),
  FOREIGN KEY ("barber_type_id") REFERENCES "BarberTypes" ("id"),
  FOREIGN KEY ("user_id") REFERENCES "Users" ("user_id"),
  UNIQUE ("nick_name", "shop_id")
);

CREATE INDEX ON "Barbers" ("barber_code");


-- Create UserRoles table
CREATE TABLE "UserRoles" (
  "id" SERIAL PRIMARY KEY,
  "user_id" INT NOT NULL,
  "barber_shop_id" INT NOT NULL,
  "role_id" SMALLINT NOT NULL,
  FOREIGN KEY ("user_id") REFERENCES "Users" ("user_id"),
  FOREIGN KEY ("barber_shop_id") REFERENCES "BarberShops" ("shop_id"),
  FOREIGN KEY ("role_id") REFERENCES "Roles" ("id")
);

-- Create indexes on Barber

Roles
CREATE INDEX ON "UserRoles" ("barber_id");
CREATE INDEX ON "UserRoles" ("barber_shop_id");


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