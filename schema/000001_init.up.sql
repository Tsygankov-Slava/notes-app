CREATE TABLE users
(
    id serial not null unique,
    username varchar(255) not null unique,
    password_hash varchar(255) not null
);

CREATE TABLE notes
(
    id serial not null unique,
    title varchar(255) not null,
    description varchar(255)
);

CREATE TABLE users_notes
(
    id serial not null unique,
    user_id int references users (id) on delete cascade not null,
    notes_id int references notes (id) on delete cascade not null
);
