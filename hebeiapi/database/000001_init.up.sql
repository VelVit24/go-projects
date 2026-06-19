create table authors (
    avatar text,
    id serial primary key,
    name text not null
);
create table categories (
    id serial primary key,
    name text not null,
    slug text not null unique
);
create table posts (
    author_id int references authors(id),
    blog_picture_path_linux text,
    blog_picture_path_windows text,
    category_id int references categories(id),
    content text,
    excerpt text,
    id serial primary key,
    published_at timestamp default(now()),
    slug text not null unique,
    title text not null
);
create table tags (
    id serial primary key,
    name text not null,
    slug text not null unique
);
create table post_tags (
    post_id int references posts(id),
    tag_id int references tags(id),
    PRIMARY KEY(post_id, tag_id)
);
create table comments (
    content text,
    created_at timestamp default(now()),
    id serial primary key,
    name text,
    parent_id int,
    post_id int references posts(id)
);
CREATE TABLE loaders (
    auto_weight INTEGER,
    battery_type TEXT,
    brake_type TEXT,
    charging_time TEXT,
    engine_type TEXT,
    fork_length INTEGER,
    front_wheels TEXT,
    height INTEGER,
    hydraulic_lifting_engine TEXT,
    id SERIAL PRIMARY KEY,
    length INTEGER,
    lift_height INTEGER,
    lifting_angle INTEGER,
    lifting_cylinder TEXT,
    longmen_frame_material INTEGER,
    max_lift_weight INTEGER,
    name TEXT NOT NULL,
    picture_path_linux TEXT,
    picture_path_windows TEXT,
    price INTEGER,
    rear_wheels TEXT,
    steering_mode TEXT,
    turning_radius INTEGER,
    voltage INTEGER,
    wheel_axis TEXT,
    width INTEGER,
    working_hours TEXT
);
CREATE TABLE manual_loaders (
    brake_type TEXT,
    control TEXT,
    drive_gear TEXT,
    fork_length INTEGER,
    fork_width INTEGER,
    id SERIAL PRIMARY KEY,
    length INTEGER,
    lifting_speed INTEGER,
    max_lift_weight INTEGER,
    max_speed NUMERIC,
    name TEXT NOT NULL,
    picture_path_linux TEXT,
    picture_path_windows TEXT,
    price INTEGER
);