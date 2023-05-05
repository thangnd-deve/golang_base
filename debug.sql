CREATE TABLE users (
    id BIGINT PRIMARY KEY ,
    owner VARCHAR(255) NOT NULL,
    balance BIGINT NOT NULL,
    currency BIGINT NOT NULL,
    created_at TIMESTAMP not NULL DEFAULT(now())
);

create table entries (
    id bigint PRIMARY KEY,
    user_id bigint  not NULL,
    amount BIGINT not NULL,
    created_at TIMESTAMP not NULL DEFAULT(now())
);

CREATE TABLE transfers (
    id bigint PRIMARY KEY,
    from_user_id bigint not NULL,
    to_user_id bigint not NULL,
    amount BIGINT not NULL,
    created_at TIMESTAMP not NULL DEFAULT(now())
);