create table video
(
    video_id             bigint auto_increment,
    user_id      bigint       not null,
    title          varchar(255) not null,
    created_at     bigint       not null,
    play_url		varchar(255) not null,
    cover_url 		varchar(255) not null,
    favorite_count	bigint       not null,
    comment_count	bigint       not null,
    is_favorite		boolean not null,
    primary key (video_id)
);


create table user
(
    user_id          bigint auto_increment,
    username         varchar(64)  not null,
    follow_count     bigint       not null,
    is_follow		 boolean not null,
    avatar			 varchar(255) not null,
    background_image varchar(255) not null,
    signature		 varchar(255) not null,
    total_favorited	 bigint       not null,
    work_count		 bigint       not null,
    favorite_count	 bigint       not null,
    primary key (user_id)
);


create table comment
(
	comment_id             bigint auto_increment,
	user_id		   bigint       not null,
	content         varchar(255) not null,
	created_at     bigint       not null,
	primary key (comment_id)
);
