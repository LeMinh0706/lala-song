-- +goose Up
-- +goose StatementBegin
CREATE TABLE "role" (
  "id" serial PRIMARY KEY,
  "name" varchar NOT NULL
);

CREATE TABLE "users" (
  "id" uuid PRIMARY KEY,
  "username" varchar UNIQUE NOT NULL,
  "password" varchar NOT NULL,
  "fullname" varchar NOT NULL,
  "gender" int NOT NULL,
  "avt" varchar NOT NULL,
  "role_id" int NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "favorite" (
  "user_id" uuid NOT NULL,
  "song_id" uuid NOT NULL,
  PRIMARY KEY ("user_id", "song_id")
);

CREATE TABLE "singers" (
  "id" bigserial PRIMARY KEY,
  "fullname" varchar NOT NULL,
  "image_url" varchar NOT NULL,
  "is_deleted" bool NOT NULL DEFAULT false
);

CREATE TABLE "album" (
  "id" bigserial PRIMARY KEY,
  "name" varchar NOT NULL,
  "image_url" varchar NOT NULL,
  "is_deleted" bool NOT NULL DEFAULT false,
  "singer_id" bigint NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "singer_song" (
  "song_id" uuid NOT NULL,
  "singer_id" bigint NOT NULL,
  PRIMARY KEY ("singer_id", "song_id")
);

CREATE TABLE "songs" (
  "id" uuid PRIMARY KEY,
  "name" varchar NOT NULL,
  "song_file" varchar NOT NULL,
  "lyric_file" varchar NOT NULL,
  "is_deleted" bool NOT NULL DEFAULT false,
  "album_id" bigint NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "song_genre" (
  "genres_id" bigint NOT NULL,
  "song_id" uuid NOT NULL,
  PRIMARY KEY ("genres_id", "song_id")
);

CREATE TABLE "genres" (
  "id" bigserial PRIMARY KEY,
  "name" varchar NOT NULL,
  "image_url" varchar NOT NULL
);

CREATE INDEX ON "users" ("username");

CREATE INDEX ON "favorite" ("user_id");

CREATE INDEX ON "favorite" ("song_id");

CREATE INDEX ON "album" ("singer_id");

CREATE INDEX ON "singer_song" ("singer_id");

CREATE INDEX ON "singer_song" ("song_id");

CREATE INDEX ON "songs" ("album_id");

CREATE INDEX ON "song_genre" ("genres_id");

CREATE INDEX ON "song_genre" ("song_id");

ALTER TABLE "users" ADD FOREIGN KEY ("role_id") REFERENCES "role" ("id");

ALTER TABLE "favorite" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

ALTER TABLE "favorite" ADD FOREIGN KEY ("song_id") REFERENCES "songs" ("id");

ALTER TABLE "album" ADD FOREIGN KEY ("singer_id") REFERENCES "singers" ("id");

ALTER TABLE "singer_song" ADD FOREIGN KEY ("song_id") REFERENCES "songs" ("id");

ALTER TABLE "singer_song" ADD FOREIGN KEY ("singer_id") REFERENCES "singers" ("id");

ALTER TABLE "songs" ADD FOREIGN KEY ("album_id") REFERENCES "album" ("id");

ALTER TABLE "song_genre" ADD FOREIGN KEY ("genres_id") REFERENCES "genres" ("id");

ALTER TABLE "song_genre" ADD FOREIGN KEY ("song_id") REFERENCES "songs" ("id");

INSERT INTO "role" VALUES (1, 'Admin'),(2, 'User');
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

DROP TABLE IF EXISTS "role" CASCADE;
DROP TABLE IF EXISTS "users" CASCADE;
DROP TABLE IF EXISTS "favorite" CASCADE;
DROP TABLE IF EXISTS "singers" CASCADE;
DROP TABLE IF EXISTS "singer_song" CASCADE;
DROP TABLE IF EXISTS "album" CASCADE;
DROP TABLE IF EXISTS "songs" CASCADE;
DROP TABLE IF EXISTS "song_genre" CASCADE;
DROP TABLE IF EXISTS "genres" CASCADE;

-- +goose StatementEnd
