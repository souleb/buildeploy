CREATE TYPE "status" AS ENUM (
  'blocked',
  'failed',
  'queued',
  'running',
  'success'
);

CREATE TABLE "pipeline" (
  "id" SERIAL PRIMARY KEY,
  "name" varchar(50),
  "workflow_id" int,
  "status" status,
  "created_at" timestamptz,
  "deleted_at" timestamptz
);

CREATE TABLE "workflow" (
  "id" SERIAL PRIMARY KEY,
  "name" varchar(50)
);

CREATE TABLE "job" (
  "id" SERIAL PRIMARY KEY,
  "name" varchar(50),
  "workflow_id" int,
  "steps" varchar,
  "env" varchar(255),
  "branches" varchar(255),
  "needs" varchar,
  "status" status
);

CREATE TABLE "edgeList" (
  "id" SERIAL PRIMARY KEY,
  "edges" varchar,
  "job_id" int
);

CREATE TABLE "docker" (
  "id" SERIAL UNIQUE PRIMARY KEY,
  "job_id" int,
  "image" varchar(255),
  "tags" varchar(255)
);

CREATE TABLE "machine" (
  "id" SERIAL UNIQUE PRIMARY KEY,
  "job_id" int,
  "os" varchar(50),
  "cpus" varchar(50),
  "memory" varchar(255)
);

ALTER TABLE "pipeline" ADD FOREIGN KEY ("workflow_id") REFERENCES "workflow" ("id");

ALTER TABLE "job" ADD FOREIGN KEY ("workflow_id") REFERENCES "workflow" ("id");

ALTER TABLE "edgeList" ADD FOREIGN KEY ("job_id") REFERENCES "job" ("id");

ALTER TABLE "docker" ADD FOREIGN KEY ("job_id") REFERENCES "job" ("id");

ALTER TABLE "machine" ADD FOREIGN KEY ("job_id") REFERENCES "job" ("id");

CREATE INDEX "pipeline_btree" ON "pipeline" USING BTREE ("id");

CREATE UNIQUE INDEX ON "pipeline" ("id");

CREATE INDEX "edgelist_btree" ON "edgeList" USING BTREE ("job_id");

CREATE UNIQUE INDEX ON "edgeList" ("id");

CREATE INDEX "docker_btree" ON "docker" USING BTREE ("job_id");

CREATE UNIQUE INDEX ON "docker" ("id");

CREATE INDEX "machine_btree" ON "machine" USING BTREE ("job_id");

CREATE UNIQUE INDEX ON "machine" ("id");
