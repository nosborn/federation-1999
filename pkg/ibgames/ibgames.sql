.bail on

PRAGMA foreign_keys = ON;
PRAGMA journal_mode = WAL;

BEGIN EXCLUSIVE TRANSACTION;

CREATE TABLE IF NOT EXISTS accounts (
    name TEXT NOT NULL, -- CHAR(32)
    name_key TEXT, -- CHAR(32)
    uid INTEGER PRIMARY KEY AUTOINCREMENT, -- SERIAL NOT NULL
    encrypt TEXT NOT NULL, -- CHAR(112)
    schange TEXT, -- DATETIME YEAR TO MINUTE
    acct_expire INT, -- INTERVAL DAY(3) TO MINUTE
    slogin TEXT, -- DATETIME YEAR TO MINUTE
    ulogin TEXT, -- DATETIME YEAR TO MINUTE
    sucip TEXT, -- CHAR(15)
    nunsuclog INTEGER DEFAULT 0, -- SMALLINT
    unsucip TEXT, -- CHAR(15)
    email TEXT, -- CHAR(48) NOT NULL
    email_key TEXT, -- CHAR(48) NOT NULL
    signup TEXT DEFAULT CURRENT_DATE, -- DATE DEFAULT TODAY
    status TEXT DEFAULT "A", -- CHAR(1)
    complimentary TEXT DEFAULT "N", -- CHAR(1)
    minutes INTEGER DEFAULT 0,
    admin_uid INT,
    admin_date TEXT, -- DATE
    bulk_mail TEXT DEFAULT "Y", -- CHAR(1)
    bad_check TEXT DEFAULT "N", -- CHAR(1)
    buddy_uid INT,
    buddy_payment TEXT DEFAULT "N", -- CHAR(1)

    UNIQUE (name_key),

    CHECK (bad_check IN ('N' ,'Y' )),
    CHECK (buddy_payment IN ('N' ,'Y' )),
    CHECK (bulk_mail IN ('N' ,'Y' )),
    CHECK (complimentary IN ('N' ,'Y' )),
    CHECK (nunsuclog >= 0 ),
    CHECK (status IN ('A' ,'S' ,'X' )),
    CHECK (uid <= 2147483647), -- int32 max value

    FOREIGN KEY (admin_uid) REFERENCES accounts(uid),
    FOREIGN KEY (buddy_uid) REFERENCES accounts(uid)
) STRICT;

CREATE INDEX IF NOT EXISTS ac_email_idx ON accounts (email_key);

CREATE TRIGGER IF NOT EXISTS prevent_uid_overflow
BEFORE INSERT ON accounts
WHEN (SELECT seq FROM sqlite_sequence WHERE name = 'accounts') >= 2147483647
BEGIN
    SELECT RAISE(FAIL, 'UID limit reached');
END;

-- CREATE TABLE bank_accounts
-- CREATE TABLE checks

CREATE TABLE IF NOT EXISTS cookies (
  sid TEXT PRIMARY KEY, -- CHAR(32),
  ip_address TEXT NOT NULL, -- CHAR(15)
  uid INTEGER NOT NULL,
  expire INTEGER NOT NULL,

  CHECK (expire > 0)
) STRICT;

CREATE TABLE IF NOT EXISTS sessions (
    sid INTEGER PRIMARY KEY AUTOINCREMENT, -- SERIAL NOT NULL
    product INTEGER NOT NULL, -- SMALLINT
    uid INTEGER NOT NULL,
    ip_address TEXT NOT NULL, -- CHAR(15)
    begin TEXT DEFAULT CURRENT_TIMESTAMP, -- DATETIME YEAR TO MINUTE DEFAULT CURRENT YEAR TO MINUTE
    end TEXT DEFAULT CURRENT_TIMESTAMP, -- DATETIME YEAR TO MINUTE DEFAULT CURRENT YEAR TO MINUTE
    minutes INTEGER NOT NULL, -- SMALLINT NOT NULL

    CHECK (product >= 0),
    CHECK (sid <= 2147483647), -- int32 max value

    FOREIGN KEY (uid) REFERENCES accounts(uid) ON DELETE CASCADE
) STRICT;

CREATE TRIGGER IF NOT EXISTS prevent_sid_overflow
BEFORE INSERT ON sessions
WHEN (SELECT seq FROM sqlite_sequence WHERE name = 'sessions') >= 2147483647
BEGIN
    SELECT RAISE(FAIL, 'SID limit reached');
END;

-- CREATE TABLE credits
-- CREATE TABLE exchange_rate
-- CREATE TABLE netbanx

COMMIT TRANSACTION;
