CREATE TABLE IF NOT EXISTS "Migrations" (
  "Id" VARCHAR(50) NOT NULL,
  "Applied" TIMESTAMP NOT NULL DEFAULT (CURRENT_TIMESTAMP)
);

CREATE OR REPLACE PROCEDURE add_migration (id VARCHAR(50))
LANGUAGE SQL
AS 
$$
  INSERT INTO "Migrations" ("Id") VALUES (id);
$$;

CREATE OR REPLACE PROCEDURE remove_migration (id VARCHAR(50))
LANGUAGE SQL
AS 
$$
  DELETE FROM "Migrations" as m WHERE m."Id" = id;
$$;

CREATE TABLE "Notebooks" (
  "Id" int GENERATED ALWAYS AS IDENTITY,
  "Name" VARCHAR(200) NOT NULL,
  "Author" VARCHAR(200) NOT NULL
);

-- Add to migration history
CALL add_migration ('0001');