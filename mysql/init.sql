create schema fxmsdb character set utf8mb4 collate utf8mb4_bin;

create user fxms identified by 'fxms123456';
grant all privileges on fxmsdb.* to 'fxms'@'%' with grant option;