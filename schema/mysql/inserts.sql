ALTER TABLE `topics` ADD INDEX `i_parentID` (`parentID`);;
ALTER TABLE `replies` ADD INDEX `i_tid` (`tid`);;
ALTER TABLE `polls` ADD INDEX `i_parentID` (`parentID`);;
ALTER TABLE `likes` ADD INDEX `i_targetItem` (`targetItem`);;
ALTER TABLE `emails` ADD INDEX `i_uid` (`uid`);;
ALTER TABLE `attachments` ADD INDEX `i_originID` (`originID`);;
ALTER TABLE `attachments` ADD INDEX `i_path` (`path`);;
ALTER TABLE `activity_stream_matches` ADD INDEX `i_watcher` (`watcher`);;
INSERT INTO `sync`(`last_update`) VALUES (UTC_TIMESTAMP());
INSERT INTO `settings`(`name`,`content`,`type`,`constraints`) VALUES ('activation_type','1','list','1-3');
INSERT INTO `settings`(`name`,`content`,`type`) VALUES ('bigpost_min_words','250','int');
INSERT INTO `settings`(`name`,`content`,`type`) VALUES ('megapost_min_words','1000','int');
INSERT INTO `settings`(`name`,`content`,`type`) VALUES ('meta_desc','','html-attribute');
INSERT INTO `settings`(`name`,`content`,`type`) VALUES ('rapid_loading','1','bool');
INSERT INTO `settings`(`name`,`content`,`type`) VALUES ('google_site_verify','','html-attribute');
INSERT INTO `themes`(`uname`,`default`) VALUES ('cosora',1);
INSERT INTO `emails`(`email`,`uid`,`validated`) VALUES ('admin@localhost',1,1);
INSERT INTO `users_groups`(`name`,`permissions`,`plugin_perms`,`is_mod`,`is_admin`,`is_banned`,`tag`) VALUES ('Administrator','{"BanUsers":true,"ActivateUsers":true,"EditUser":true,"EditUserEmail":true,"EditUserPassword":true,"EditUserGroup":true,"EditUserGroupSuperMod":true,"EditGroup":true,"EditGroupLocalPerms":true,"EditGroupGlobalPerms":true,"EditGroupSuperMod":true,"ManageForums":true,"EditSettings":true,"ManageThemes":true,"ManagePlugins":true,"ViewAdminLogs":true,"ViewIPs":true,"UploadFiles":true,"UploadAvatars":true,"UseConvos":true,"UseConvosOnlyWithMod":true,"CreateProfileReply":true,"AutoEmbed":true,"AutoLink":true,"ViewTopic":true,"LikeItem":true,"CreateTopic":true,"EditTopic":true,"DeleteTopic":true,"CreateReply":true,"EditReply":true,"DeleteReply":true,"PinTopic":true,"CloseTopic":true,"MoveTopic":true}','{}',1,1,0,'Admin');
INSERT INTO `users_groups`(`name`,`permissions`,`plugin_perms`,`is_mod`,`is_admin`,`is_banned`,`tag`) VALUES ('Moderator','{"BanUsers":true,"ActivateUsers":true,"EditUser":true,"EditUserGroup":true,"ViewIPs":true,"UploadFiles":true,"UploadAvatars":true,"UseConvos":true,"UseConvosOnlyWithMod":true,"CreateProfileReply":true,"AutoEmbed":true,"AutoLink":true,"ViewTopic":true,"LikeItem":true,"CreateTopic":true,"EditTopic":true,"DeleteTopic":true,"CreateReply":true,"EditReply":true,"DeleteReply":true,"PinTopic":true,"CloseTopic":true,"MoveTopic":true}','{}',1,0,0,'Mod');
INSERT INTO `users_groups`(`name`,`permissions`,`plugin_perms`,`is_mod`,`is_admin`,`is_banned`,`tag`) VALUES ('Member','{"UploadFiles":true,"UploadAvatars":true,"UseConvos":true,"UseConvosOnlyWithMod":true,"CreateProfileReply":true,"AutoEmbed":true,"AutoLink":true,"ViewTopic":true,"LikeItem":true,"CreateTopic":true,"CreateReply":true}','{}',0,0,0,"");
INSERT INTO `users_groups`(`name`,`permissions`,`plugin_perms`,`is_mod`,`is_admin`,`is_banned`,`tag`) VALUES ('Banned','{"ViewTopic":true}','{}',0,0,1,"");
INSERT INTO `users_groups`(`name`,`permissions`,`plugin_perms`,`is_mod`,`is_admin`,`is_banned`,`tag`) VALUES ('Awaiting Activation','{"UseConvosOnlyWithMod":true,"ViewTopic":true}','{}',0,0,0,"");
INSERT INTO `users_groups`(`name`,`permissions`,`plugin_perms`,`is_mod`,`is_admin`,`is_banned`,`tag`) VALUES ('Not Loggedin','{"ViewTopic":true}','{}',0,0,0,'Guest');
INSERT INTO `forums`(`name`,`active`,`desc`,`tmpl`) VALUES ('Reports',0,'All the reports go here','');
INSERT INTO `forums`(`name`,`lastTopicID`,`lastReplyerID`,`desc`,`tmpl`) VALUES ('General',1,1,'A place for general discussions which don''t fit elsewhere','');
INSERT INTO `forums_permissions`(`gid`,`fid`,`permissions`) VALUES (1,1,'{"ViewTopic":true,"CreateReply":true,"CreateTopic":true,"PinTopic":true,"CloseTopic":true}');
INSERT INTO `forums_permissions`(`gid`,`fid`,`permissions`) VALUES (2,1,'{"ViewTopic":true,"CreateReply":true,"CloseTopic":true}');
INSERT INTO `forums_permissions`(`gid`,`fid`,`permissions`) VALUES (3,1,'{}');
INSERT INTO `forums_permissions`(`gid`,`fid`,`permissions`) VALUES (4,1,'{}');
INSERT INTO `forums_permissions`(`gid`,`fid`,`permissions`) VALUES (5,1,'{}');
INSERT INTO `forums_permissions`(`gid`,`fid`,`permissions`) VALUES (6,1,'{}');
INSERT INTO `forums_permissions`(`gid`,`fid`,`permissions`) VALUES (1,2,'{"ViewTopic":true,"CreateReply":true,"CreateTopic":true,"LikeItem":true,"EditTopic":true,"DeleteTopic":true,"EditReply":true,"DeleteReply":true,"PinTopic":true,"CloseTopic":true,"MoveTopic":true}');
INSERT INTO `forums_permissions`(`gid`,`fid`,`permissions`) VALUES (2,2,'{"ViewTopic":true,"CreateReply":true,"CreateTopic":true,"LikeItem":true,"EditTopic":true,"DeleteTopic":true,"EditReply":true,"DeleteReply":true,"PinTopic":true,"CloseTopic":true,"MoveTopic":true}');
INSERT INTO `forums_permissions`(`gid`,`fid`,`permissions`) VALUES (3,2,'{"ViewTopic":true,"CreateReply":true,"CreateTopic":true,"LikeItem":true}');
INSERT INTO `forums_permissions`(`gid`,`fid`,`permissions`) VALUES (4,2,'{"ViewTopic":true}');
INSERT INTO `forums_permissions`(`gid`,`fid`,`permissions`) VALUES (5,2,'{"ViewTopic":true}');
INSERT INTO `forums_permissions`(`gid`,`fid`,`permissions`) VALUES (6,2,'{"ViewTopic":true}');
INSERT INTO `topics`(`title`,`content`,`parsed_content`,`createdAt`,`lastReplyAt`,`lastReplyBy`,`createdBy`,`parentID`,`ip`) VALUES ('Test Topic','A topic automatically generated by the software.','A topic automatically generated by the software.',UTC_TIMESTAMP(),UTC_TIMESTAMP(),1,1,2,'::1');
INSERT INTO `replies`(`tid`,`content`,`parsed_content`,`createdAt`,`createdBy`,`lastUpdated`,`lastEdit`,`lastEditBy`,`ip`) VALUES (1,'A reply!','A reply!',UTC_TIMESTAMP(),1,UTC_TIMESTAMP(),0,0,'::1');
INSERT INTO `menus`() VALUES ();
INSERT INTO `menu_items`(`mid`,`name`,`htmlID`,`position`,`path`,`aria`,`tooltip`,`order`) VALUES (1,'{lang.menu_forums}','menu_forums','left','/forums/','{lang.menu_forums_aria}','{lang.menu_forums_tooltip}',0);
INSERT INTO `menu_items`(`mid`,`name`,`htmlID`,`cssClass`,`position`,`path`,`aria`,`tooltip`,`order`) VALUES (1,'{lang.menu_topics}','menu_topics','menu_topics','left','/topics/','{lang.menu_topics_aria}','{lang.menu_topics_tooltip}',1);
INSERT INTO `menu_items`(`mid`,`htmlID`,`cssClass`,`position`,`tmplName`,`order`) VALUES (1,'general_alerts','menu_alerts','right','menu_alerts',2);
INSERT INTO `menu_items`(`mid`,`name`,`cssClass`,`position`,`path`,`aria`,`tooltip`,`memberOnly`,`order`) VALUES (1,'{lang.menu_account}','menu_account','left','/user/edit/','{lang.menu_account_aria}','{lang.menu_account_tooltip}',1,3);
INSERT INTO `menu_items`(`mid`,`name`,`cssClass`,`position`,`path`,`aria`,`tooltip`,`memberOnly`,`order`) VALUES (1,'{lang.menu_profile}','menu_profile','left','{me.Link}','{lang.menu_profile_aria}','{lang.menu_profile_tooltip}',1,4);
INSERT INTO `menu_items`(`mid`,`name`,`cssClass`,`position`,`path`,`aria`,`tooltip`,`memberOnly`,`staffOnly`,`order`) VALUES (1,'{lang.menu_panel}','menu_panel menu_account','left','/panel/','{lang.menu_panel_aria}','{lang.menu_panel_tooltip}',1,1,5);
INSERT INTO `menu_items`(`mid`,`name`,`cssClass`,`position`,`path`,`aria`,`tooltip`,`memberOnly`,`order`) VALUES (1,'{lang.menu_logout}','menu_logout','left','/accounts/logout/?s={me.Session}','{lang.menu_logout_aria}','{lang.menu_logout_tooltip}',1,6);
INSERT INTO `menu_items`(`mid`,`name`,`cssClass`,`position`,`path`,`aria`,`tooltip`,`guestOnly`,`order`) VALUES (1,'{lang.menu_register}','menu_register','left','/accounts/create/','{lang.menu_register_aria}','{lang.menu_register_tooltip}',1,7);
INSERT INTO `menu_items`(`mid`,`name`,`cssClass`,`position`,`path`,`aria`,`tooltip`,`guestOnly`,`order`) VALUES (1,'{lang.menu_login}','menu_login','left','/accounts/login/','{lang.menu_login_aria}','{lang.menu_login_tooltip}',1,8);
