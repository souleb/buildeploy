CREATE TABLE "pipeline" (
  "id" SERIAL PRIMARY KEY,
  "name" varchar(50),
  "status" int,
  "created_at" timestamptz NOT NULL DEFAULT NOW() ,
  "updated_at" timestamptz NOT NULL DEFAULT NOW() ,
  "deleted_at" timestamptz
);

CREATE TABLE "workflow" (
  "id" SERIAL PRIMARY KEY,
  "name" varchar(50),
  "pipeline_id" int
);

CREATE TABLE "job" (
  "id" SERIAL PRIMARY KEY,
  "name" varchar(50),
  "workflow_id" int,
  "steps" varchar,
  "env" varchar(255),
  "branches" varchar(255),
  "needs" varchar,
  "status" int
);

CREATE TABLE "job_docker" (
  "id" SERIAL PRIMARY KEY,
  "job_id" int,
  "image" varchar(255),
  "tags" varchar(255)
);

CREATE TABLE "job_machine" (
  "id" SERIAL PRIMARY KEY,
  "job_id" int,
  "os" varchar(50),
  "cpus" varchar(50),
  "memory" varchar(255)
);

ALTER TABLE "workflow" ADD FOREIGN KEY ("pipeline_id") REFERENCES "pipeline" ("id");

ALTER TABLE "job" ADD FOREIGN KEY ("workflow_id") REFERENCES "workflow" ("id");

ALTER TABLE "job_docker" ADD FOREIGN KEY ("job_id") REFERENCES "job" ("id");

ALTER TABLE "job_machine" ADD FOREIGN KEY ("job_id") REFERENCES "job" ("id");


CREATE OR REPLACE FUNCTION trigger_set_timestamp()
RETURNS TRIGGER AS $$
BEGIN
  NEW.updated_at = NOW();
  RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER set_timestamp
BEFORE UPDATE ON pipeline
FOR EACH ROW
EXECUTE PROCEDURE trigger_set_timestamp();
