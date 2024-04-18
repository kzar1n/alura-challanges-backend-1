create table if not exists aluraflix.videos (
    id int primary key auto_increment,
    title varchar(255) not null,
    description text not null,
    url varchar(255) not null
);