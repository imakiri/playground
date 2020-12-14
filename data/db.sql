-- main.users definition

-- Drop table

-- DROP TABLE main.users;

CREATE TABLE main.users
(
    id        int4        NOT NULL GENERATED ALWAYS AS IDENTITY,
    login     varchar(45) NOT NULL,
    "name"    varchar(45) NOT NULL,
    avatar    bytea       NOT NULL,
    pass_hash bytea       NOT NULL,
    CONSTRAINT users_pk PRIMARY KEY (id),
    CONSTRAINT users_un UNIQUE (login, id)
);
CREATE UNIQUE INDEX users_id_idx ON main.users USING btree (id);
CREATE UNIQUE INDEX users_login_idx ON main.users USING btree (login);


-- main."data" definition

-- Drop table

-- DROP TABLE main."data";

CREATE TABLE main."data"
(
    id      int4         NOT NULL GENERATED ALWAYS AS IDENTITY,
    "date"  timestamp(0) NOT NULL,
    query   varchar(45)  NOT NULL,
    picture bytea        NOT NULL,
    user_id int4         NOT NULL,
    CONSTRAINT data_pk PRIMARY KEY (id),
    CONSTRAINT data_un UNIQUE (id),
    CONSTRAINT data_fk FOREIGN KEY (user_id) REFERENCES main.users (id) ON UPDATE CASCADE ON DELETE CASCADE
);
CREATE INDEX data_date_idx ON main.data USING btree (date);
CREATE UNIQUE INDEX data_id_idx ON main.data USING btree (id);
CREATE INDEX data_user_id_idx ON main.data USING btree (user_id);


-- main.permissions definition

-- Drop table

-- DROP TABLE main.permissions;

CREATE TABLE main.permissions
(
    user_id         int4 NOT NULL,
    some_permission bool NOT NULL,
    CONSTRAINT permissions_pk PRIMARY KEY (user_id),
    CONSTRAINT permissions_fk FOREIGN KEY (user_id) REFERENCES main.users (id) ON UPDATE CASCADE ON DELETE CASCADE
);