ALTER TABLE "Customers"
ADD COLUMN is_social_auth BOOLEAN DEFAULT false;

ALTER TABLE "SessionsCustomer"
ADD COLUMN  timezone VARCHAR NOT NULL;

CREATE INDEX idx_customers_email ON "Customers" (email);