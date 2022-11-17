CREATE TABLE "user"
(
    id		    INTEGER PRIMARY KEY,
    balance     DECIMAL(19,4) NOT NULL DEFAULT 0.0 CONSTRAINT positive_balance CHECK (balance >= 0.0),
    reserve     DECIMAL(19,4) NOT NULL DEFAULT 0.0 CONSTRAINT positive_resereve CHECK (resereve >= 0.0)
);

CREATE TABLE "order"
(
    id		        INTEGER PRIMARY KEY,
    user_id         INTEGER,
    service_id      INTEGER,
    amount          DECIMAL(19,4) NOT NULL,
);