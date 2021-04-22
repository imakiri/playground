-- DROP SCHEMA app;

CREATE SCHEMA app AUTHORIZATION postgres;
-- app.categories definition

-- Drop table

-- DROP TABLE app.categories;

CREATE TABLE app.categories
(
    category_uuid bpchar(21)   NOT NULL,
    "name"        varchar(121) NOT NULL,
    CONSTRAINT categories_un UNIQUE (category_uuid)
);
CREATE UNIQUE INDEX categories_category_uuid_idx ON app.categories USING btree (category_uuid);


-- app.users definition

-- Drop table

-- DROP TABLE app.users;

CREATE TABLE app.users
(
    user_uuid         bpchar(21)   NOT NULL,
    registration_date int8         NOT NULL,
    nick_name         varchar(32)  NOT NULL,
    full_name         varchar(121) NOT NULL,
    avatar512         bytea        NOT NULL,
    avatar256         bytea        NOT NULL,
    avatar128         bytea        NOT NULL,
    CONSTRAINT users_pk_nick UNIQUE (nick_name),
    CONSTRAINT users_un_uuid UNIQUE (user_uuid)
);


-- app.threads definition

-- Drop table

-- DROP TABLE app.threads;

CREATE TABLE app.threads
(
    thread_uuid    bpchar(21)     NOT NULL,
    category_uuid  bpchar(21)     NOT NULL,
    user_uuid      bpchar(21)     NOT NULL,
    "name"         varchar(121)   NOT NULL,
    date_added     int8           NOT NULL,
    date_last_edit int8           NOT NULL,
    "header"       varchar(65536) NOT NULL,
    CONSTRAINT threads_un UNIQUE (thread_uuid),
    CONSTRAINT threads_fk FOREIGN KEY (user_uuid) REFERENCES app.users (user_uuid),
    CONSTRAINT threads_fk_category FOREIGN KEY (category_uuid) REFERENCES app.categories (category_uuid)
);
CREATE INDEX threads_category_uuid_idx ON app.threads USING btree (category_uuid);
CREATE INDEX threads_user_uuid_idx ON app.threads USING btree (user_uuid);


-- app.posts definition

-- Drop table

-- DROP TABLE app.posts;

CREATE TABLE app.posts
(
    post_uuid      bpchar(36)     NOT NULL,
    thread_uuid    bpchar(21)     NOT NULL,
    user_uuid      bpchar(21)     NOT NULL,
    date_added     int8           NOT NULL,
    date_last_edit int8           NOT NULL,
    "content"      varchar(32768) NOT NULL,
    CONSTRAINT posts_un UNIQUE (post_uuid),
    CONSTRAINT posts_fk_thread FOREIGN KEY (thread_uuid) REFERENCES app.threads (thread_uuid),
    CONSTRAINT posts_fk_user FOREIGN KEY (user_uuid) REFERENCES app.users (user_uuid)
);
CREATE UNIQUE INDEX posts_post_uuid_idx ON app.posts USING btree (post_uuid);
CREATE INDEX posts_thread_uuid_idx ON app.posts USING btree (thread_uuid);
CREATE INDEX posts_user_uuid_idx ON app.posts USING btree (user_uuid);



CREATE OR REPLACE FUNCTION app.gen_random_bytes(integer)
    RETURNS bytea
    LANGUAGE c
    PARALLEL SAFE STRICT
AS
'$libdir/pgcrypto',
$function$pg_random_bytes$function$
;