ALTER TABLE "Customers"
    ADD COLUMN is_social_auth BOOLEAN DEFAULT false,
    ALTER COLUMN hashed_password DROP NOT NULL,
    ALTER COLUMN phone DROP NOT NULL;

ALTER TABLE "SessionsCustomer"
ADD COLUMN  timezone VARCHAR(255) NOT NULL;

CREATE INDEX idx_customers_email ON "Customers" (email);