create table authors (
    avatar text not null unique,
    id serial primary key,
    name text
);
create table categories (
    id serial primary key,
    name text not null unique,
    slug text not null
);
create table posts (
    author_id int references authors(id),
    blogPicturePathLinux text,
    blogPicturePathWindows text,
    category_slug int references categories(slug),
    content text,
    id serial primary key,
    publishedAt timestamp default(now()),
    slug text not null,
    tags int[],
    title text not null
);
create table tags (
    id serial primary key,
    name text not null,
    slug text not null
);
create table comments (
    content text,
    createdAt timestamp default(now()),
    id serail primary key,
    name text,
    parentId int,
    post_slug int references posts(slug)
);
