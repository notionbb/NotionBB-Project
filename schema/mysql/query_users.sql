CREATE TABLE `users` (
	`uid` int not null AUTO_INCREMENT,
	`name` varchar(100) not null,
	`password` varchar(100) not null,
	`salt` varchar(80) DEFAULT '' not null,
	`group` int not null,
	`active` boolean DEFAULT 0 not null,
	`is_super_admin` boolean DEFAULT 0 not null,
	`createdAt` datetime not null,
	`lastActiveAt` datetime not null,
	`session` varchar(200) DEFAULT '' not null,
	`last_ip` varchar(200) DEFAULT '' not null,
	`profile_comments` int DEFAULT 0 not null,
	`who_can_convo` int DEFAULT 0 not null,
	`enable_embeds` int DEFAULT -1 not null,
	`email` varchar(200) DEFAULT '' not null,
	`avatar` varchar(100) DEFAULT '' not null,
	`message` text not null,
	`url_prefix` varchar(20) DEFAULT '' not null,
	`url_name` varchar(100) DEFAULT '' not null,
	`level` smallint DEFAULT 0 not null,
	`score` int DEFAULT 0 not null,
	`posts` int DEFAULT 0 not null,
	`bigposts` int DEFAULT 0 not null,
	`megaposts` int DEFAULT 0 not null,
	`topics` int DEFAULT 0 not null,
	`liked` int DEFAULT 0 not null,
	`oldestItemLikedCreatedAt` datetime not null,
	`lastLiked` datetime not null,
	`temp_group` int DEFAULT 0 not null,
	primary key(`uid`),
	unique(`name`)
) CHARSET=utf8mb4 COLLATE utf8mb4_general_ci;