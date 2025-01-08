create table links
(
    id varchar(100) not null,
    title varchar(100) not null,
    url varchar(2048) not null,
    username varchar(100) not null,
    created_at bigint not null,
    updated_at bigint not null,
    primary key (id),
    foreign key fk_links_username (username) references users (username)
) engine = InnoDB;