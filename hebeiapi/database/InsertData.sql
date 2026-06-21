-- AUTHORS

INSERT INTO authors (avatar, name) VALUES
('avatars/alex.png', 'Alex'),
('avatars/max.png', 'Max'),
('avatars/anna.png', 'Anna');


-- CATEGORIES

INSERT INTO categories (name, slug) VALUES
('Technology', 'technology'),
('News', 'news'),
('Equipment', 'equipment');


-- POSTS

INSERT INTO posts (
    author_id,
    blog_picture_path_linux,
    blog_picture_path_windows,
    category_id,
    content,
    excerpt,
    slug,
    title
) VALUES
(
    1,
    '/images/linux1.png',
    'C:\\images\\linux1.png',
    1,
    'Linux is an operating system widely used by developers for backend development, programming and server administration.',
    'Learn why Linux is popular among developers and how to start using it.',
    'linux-development',
    'Linux for developers'
),
(
    2,
    '/images/loader.png',
    'C:\\images\\loader.png',
    3,
    'Modern loaders overview with specifications, features and comparison of different models.',
    'Overview of modern loaders and tips for choosing the right equipment.',
    'modern-loaders',
    'How to choose loader'
),
(
    3,
    '/images/postgres.png',
    'C:\\images\\postgres.png',
    1,
    'PostgreSQL database design principles, tables, relations and optimization techniques.',
    'A basic guide to designing databases with PostgreSQL.',
    'postgres-design',
    'Building database schema'
);


-- TAGS

INSERT INTO tags (name, slug) VALUES
('Linux', 'linux'),
('Go', 'go'),
('PostgreSQL', 'postgresql'),
('Machines', 'machines');


-- POST TAGS

INSERT INTO post_tags (post_id, tag_id) VALUES
(1,1),
(1,2),
(1,3),
(2,4),
(3,3);


-- COMMENTS
-- post 1

INSERT INTO comments (
    content,
    name,
    post_id,
    parent_id
) VALUES
(
    'Great article!',
    'Mike',
    1,
    0
),
(
    'I agree, Linux is very useful',
    'John',
    1,
    1
),
(
    'Especially for backend development',
    'Kate',
    1,
    2
),
(
    'Nice explanation',
    'Bob',
    1,
    0
);


-- LOADERS

INSERT INTO loaders (
    auto_weight,
    battery_type,
    brake_type,
    charging_time,
    engine_type,
    fork_length,
    front_wheels,
    height,
    hydraulic_lifting_engine,
    length,
    lift_height,
    lifting_angle,
    lifting_cylinder,
    longmen_frame_material,
    max_lift_weight,
    name,
    picture_path_linux,
    picture_path_windows,
    price,
    rear_wheels,
    steering_mode,
    turning_radius,
    voltage,
    wheel_axis,
    width,
    working_hours
) VALUES
(
    1200,
    'Lithium',
    'Electric',
    '5 hours',
    'Electric motor',
    1200,
    'Rubber',
    2100,
    'Hydraulic',
    2500,
    3000,
    45,
    'Double cylinder',
    1,
    2000,
    'Electric Loader X1',
    '/loaders/x1.png',
    'C:\\loaders\\x1.png',
    15000,
    'Rubber',
    'Electric steering',
    2500,
    48,
    'Standard',
    1200,
    '8 hours'
),
(
    1500,
    'Lead Acid',
    'Mechanical',
    '8 hours',
    'Diesel',
    1500,
    'Steel',
    2300,
    'Hydraulic',
    2800,
    3500,
    50,
    'Single cylinder',
    2,
    2500,
    'Heavy Loader Pro',
    '/loaders/pro.png',
    'C:\\loaders\\pro.png',
    22000,
    'Steel',
    'Manual',
    3000,
    72,
    'Reinforced',
    1400,
    '10 hours'
);


-- MANUAL LOADERS

INSERT INTO manual_loaders (
    brake_type,
    control,
    drive_gear,
    fork_length,
    fork_width,
    length,
    lifting_speed,
    max_lift_weight,
    max_speed,
    name,
    picture_path_linux,
    picture_path_windows,
    price
) VALUES
(
    'Mechanical',
    'Handle',
    'Manual',
    1150,
    550,
    1800,
    120,
    1500,
    5.5,
    'Manual Loader Basic',
    '/manual/basic.png',
    'C:\\manual\\basic.png',
    900
),
(
    'Hydraulic',
    'Lever',
    'Hydraulic',
    1300,
    600,
    2000,
    150,
    2000,
    6.0,
    'Manual Loader Pro',
    '/manual/pro.png',
    'C:\\manual\\pro.png',
    1500
);