CREATE TABLE IF NOT EXISTS users
(
    id         serial primary key,
    username   varchar(255) not null unique,
    password   varchar(255) not null,
    created_at timestamp    not null default now()
);

CREATE TABLE IF NOT EXISTS accounts
(
    id         serial primary key,
    balance    int       not null default 0,
    created_at timestamp not null default now()
);

CREATE TABLE IF NOT EXISTS products
(
    id   serial primary key,
    name varchar(255) not null unique
);

CREATE TABLE IF NOT EXISTS reservations
(
    id         serial primary key,
    account_id int       not null,
    product_id int       not null,
    order_id   int       not null unique,
    amount     int       not null,
    created_at timestamp not null default now(),
    foreign key (account_id) references accounts (id),
    foreign key (product_id) references products (id)
);

CREATE TABLE IF NOT EXISTS operations
(
    id             serial primary key,
    account_id     int          not null,
    amount         int          not null,
    operation_type varchar(255) not null,
    created_at     timestamp    not null default now(),
    product_id     int                   default null,
    order_id       int                   default null,
    description    varchar(255)          default null,
    foreign key (account_id) references accounts (id)
);