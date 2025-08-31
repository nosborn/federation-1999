.bail on

PRAGMA foreign_keys = ON;
PRAGMA journal_mode = WAL;

BEGIN EXCLUSIVE TRANSACTION;

CREATE TABLE IF NOT EXISTS accounts (
  uid INTEGER PRIMARY KEY AUTOINCREMENT,
  data JSON NOT NULL,
  name TEXT GENERATED ALWAYS AS (json_extract(data, '$.name')) STORED NOT NULL,
  email TEXT GENERATED ALWAYS AS (json_extract(data, '$.email')) STORED,

  UNIQUE (name COLLATE NOCASE),
  UNIQUE (email COLLATE NOCASE)

  CHECK (uid <= 2147483647) -- int32 max value
);

CREATE TRIGGER IF NOT EXISTS prevent_uid_overflow
BEFORE INSERT ON accounts
WHEN (SELECT seq FROM sqlite_sequence WHERE name = 'accounts') >= 4294967295
BEGIN
    SELECT RAISE(FAIL, 'UID limit reached');
END;

CREATE TABLE IF NOT EXISTS sessions (
  sid INTEGER PRIMARY KEY AUTOINCREMENT,
  data JSON NOT NULL,
  uid INTEGER GENERATED ALWAYS AS (json_extract(data, '$.uid')) STORED NOT NULL

  CHECK (sid <= 2147483647) -- int32 max value
);

CREATE TRIGGER IF NOT EXISTS prevent_sid_overflow
BEFORE INSERT ON sessions
WHEN (SELECT seq FROM sqlite_sequence WHERE name = 'sessions') >= 4294967295
BEGIN
    SELECT RAISE(FAIL, 'SID limit reached');
END;

COMMIT TRANSACTION;
