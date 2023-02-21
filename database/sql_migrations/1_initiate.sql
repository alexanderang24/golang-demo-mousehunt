-- +migrate Up
-- +migrate StatementBegin

CREATE TABLE trap
(
    id          SERIAL PRIMARY KEY,
    name        VARCHAR(256),
    description VARCHAR(256),
    min_power   INT,
    max_power   INT,
    price       INT,
    created_at  TIMESTAMP,
    updated_at  TIMESTAMP
);

CREATE TABLE location
(
    id          SERIAL PRIMARY KEY,
    name        VARCHAR(256),
    description VARCHAR(256),
    travel_cost INT,
    created_at  TIMESTAMP,
    updated_at  TIMESTAMP
);

CREATE TABLE mouse
(
    id          SERIAL PRIMARY KEY,
    name        VARCHAR(256),
    description VARCHAR(256),
    min_power   INT,
    max_power   INT,
    gold        INT,
    location_id SERIAL,
    created_at  TIMESTAMP,
    updated_at  TIMESTAMP,
    CONSTRAINT fk_mouse_location
        FOREIGN KEY (location_id)
            REFERENCES location (id)
);

CREATE TABLE "user"
(
    id          SERIAL PRIMARY KEY,
    username    VARCHAR(256),
    password    VARCHAR(256),
    role        VARCHAR(256),
    gold        INT,
    location_id SERIAL,
    trap_id     SERIAL,
    created_at  TIMESTAMP,
    updated_at  TIMESTAMP,
    CONSTRAINT fk_user_location
        FOREIGN KEY (location_id)
            REFERENCES location (id),
    CONSTRAINT fk_user_trap
        FOREIGN KEY (trap_id)
            REFERENCES trap (id)
);

CREATE TABLE hunt_history
(
    id          SERIAL PRIMARY KEY,
    user_id     SERIAL,
    mouse_id    SERIAL,
    location_id SERIAL,
    trap_id     SERIAL,
    created_at  TIMESTAMP,
    CONSTRAINT fk_hunt_history_user
        FOREIGN KEY (user_id)
            REFERENCES "user" (id),
    CONSTRAINT fk_hunt_history_mouse
        FOREIGN KEY (mouse_id)
            REFERENCES mouse (id),
    CONSTRAINT fk_hunt_history_location
        FOREIGN KEY (location_id)
            REFERENCES location (id),
    CONSTRAINT fk_hunt_history_trap
        FOREIGN KEY (trap_id)
            REFERENCES trap (id)
);


-- +migrate StatementEnd
