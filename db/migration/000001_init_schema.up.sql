CREATE TABLE "devices" (
  "device_name" varchar PRIMARY KEY,
  "expected" varchar NOT NULL,
  "price" bigint NOT NULL,
  "img_url" varchar NOT NULL,
  "source_url" varchar NOT NULL,
  "spec_score" int NOT NULL
);

CREATE INDEX ON "devices" ("price");

CREATE INDEX ON "devices" ("spec_score");

COMMENT ON COLUMN "devices"."expected" IS '"-" implies not specified, saved as :- Expected launch: Month 20xx';

COMMENT ON COLUMN "devices"."spec_score" IS '0 implies not specified';