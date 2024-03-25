drop database video_show;
create database video_show;

use video_show;

-- insert into  video_path values(null,'image1')
-- insert into  video_path values(null,'image2');
-- select * from  video_path where id>=1 limit 1;
-- 
-- You must first delete data from child table and then delete the image;
-- delete from  like where imageid=1;delete from  video_path where id=1;
-- video_path table 
drop table video_path;
create table video_path(
    id bigint unsigned primary key auto_increment,
    name varchar(256) not null
);
-- User table 
-- insert into  user values(null,'user1','1234');
-- delete from  like where userid=1;delete from  user where id=1;
drop table user;
create table user(
    id int unsigned primary key auto_increment,
    name varchar(256) unique not null,
    password varchar(256) not null,
    viewposition bigint unsigned
);
-- User like image table
-- delete from  like where userid=1;
-- delete from  like where imageid=1;
drop table collect;
create table collect(
    id bigint unsigned primary key auto_increment,
    userid int unsigned  not null,
    videoid bigint unsigned  not null,
    index userid_idx (userid),
    index videoid_idx (videoid)
);