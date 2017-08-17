// +build !pgsql !sqlite !mssql

/* This file was generated by Gosora's Query Generator. Please try to avoid modifying this file, as it might change at any time. */

package main

import "log"
import "database/sql"

var get_user_stmt *sql.Stmt
var get_reply_stmt *sql.Stmt
var get_user_reply_stmt *sql.Stmt
var get_password_stmt *sql.Stmt
var get_settings_stmt *sql.Stmt
var get_setting_stmt *sql.Stmt
var get_full_setting_stmt *sql.Stmt
var get_full_settings_stmt *sql.Stmt
var get_groups_stmt *sql.Stmt
var get_forums_stmt *sql.Stmt
var get_forums_permissions_stmt *sql.Stmt
var get_plugins_stmt *sql.Stmt
var get_themes_stmt *sql.Stmt
var get_widgets_stmt *sql.Stmt
var is_plugin_active_stmt *sql.Stmt
var get_users_stmt *sql.Stmt
var get_users_offset_stmt *sql.Stmt
var is_theme_default_stmt *sql.Stmt
var get_modlogs_stmt *sql.Stmt
var get_modlogs_offset_stmt *sql.Stmt
var get_reply_tid_stmt *sql.Stmt
var get_topic_fid_stmt *sql.Stmt
var get_user_reply_uid_stmt *sql.Stmt
var has_liked_topic_stmt *sql.Stmt
var has_liked_reply_stmt *sql.Stmt
var get_user_name_stmt *sql.Stmt
var get_user_rank_stmt *sql.Stmt
var get_user_active_stmt *sql.Stmt
var get_user_group_stmt *sql.Stmt
var get_emails_by_user_stmt *sql.Stmt
var get_topic_basic_stmt *sql.Stmt
var get_activity_entry_stmt *sql.Stmt
var forum_entry_exists_stmt *sql.Stmt
var group_entry_exists_stmt *sql.Stmt
var get_forum_topics_offset_stmt *sql.Stmt
var get_topic_replies_offset_stmt *sql.Stmt
var get_topic_list_stmt *sql.Stmt
var get_topic_user_stmt *sql.Stmt
var get_topic_by_reply_stmt *sql.Stmt
var get_topic_replies_stmt *sql.Stmt
var get_forum_topics_stmt *sql.Stmt
var get_profile_replies_stmt *sql.Stmt
var get_watchers_stmt *sql.Stmt
var create_topic_stmt *sql.Stmt
var create_report_stmt *sql.Stmt
var create_reply_stmt *sql.Stmt
var create_action_reply_stmt *sql.Stmt
var create_like_stmt *sql.Stmt
var add_activity_stmt *sql.Stmt
var notify_one_stmt *sql.Stmt
var add_email_stmt *sql.Stmt
var create_profile_reply_stmt *sql.Stmt
var add_subscription_stmt *sql.Stmt
var create_forum_stmt *sql.Stmt
var add_forum_perms_to_forum_stmt *sql.Stmt
var add_plugin_stmt *sql.Stmt
var add_theme_stmt *sql.Stmt
var create_group_stmt *sql.Stmt
var add_modlog_entry_stmt *sql.Stmt
var add_adminlog_entry_stmt *sql.Stmt
var add_forum_perms_to_group_stmt *sql.Stmt
var add_replies_to_topic_stmt *sql.Stmt
var remove_replies_from_topic_stmt *sql.Stmt
var add_topics_to_forum_stmt *sql.Stmt
var remove_topics_from_forum_stmt *sql.Stmt
var update_forum_cache_stmt *sql.Stmt
var add_likes_to_topic_stmt *sql.Stmt
var add_likes_to_reply_stmt *sql.Stmt
var edit_topic_stmt *sql.Stmt
var edit_reply_stmt *sql.Stmt
var stick_topic_stmt *sql.Stmt
var unstick_topic_stmt *sql.Stmt
var update_last_ip_stmt *sql.Stmt
var update_session_stmt *sql.Stmt
var set_password_stmt *sql.Stmt
var set_avatar_stmt *sql.Stmt
var set_username_stmt *sql.Stmt
var change_group_stmt *sql.Stmt
var activate_user_stmt *sql.Stmt
var update_user_level_stmt *sql.Stmt
var increment_user_score_stmt *sql.Stmt
var increment_user_posts_stmt *sql.Stmt
var increment_user_bigposts_stmt *sql.Stmt
var increment_user_megaposts_stmt *sql.Stmt
var increment_user_topics_stmt *sql.Stmt
var edit_profile_reply_stmt *sql.Stmt
var delete_forum_stmt *sql.Stmt
var update_forum_stmt *sql.Stmt
var update_setting_stmt *sql.Stmt
var update_plugin_stmt *sql.Stmt
var update_plugin_install_stmt *sql.Stmt
var update_theme_stmt *sql.Stmt
var update_user_stmt *sql.Stmt
var update_group_perms_stmt *sql.Stmt
var update_group_rank_stmt *sql.Stmt
var update_group_stmt *sql.Stmt
var update_email_stmt *sql.Stmt
var verify_email_stmt *sql.Stmt
var delete_reply_stmt *sql.Stmt
var delete_topic_stmt *sql.Stmt
var delete_profile_reply_stmt *sql.Stmt
var delete_forum_perms_by_forum_stmt *sql.Stmt
var report_exists_stmt *sql.Stmt
var group_count_stmt *sql.Stmt
var modlog_count_stmt *sql.Stmt
var add_forum_perms_to_forum_admins_stmt *sql.Stmt
var add_forum_perms_to_forum_staff_stmt *sql.Stmt
var add_forum_perms_to_forum_members_stmt *sql.Stmt
var notify_watchers_stmt *sql.Stmt

func _gen_mysql() (err error) {
	if dev.DebugMode {
		log.Print("Building the generated statements")
	}
	
	log.Print("Preparing get_user statement.")
	get_user_stmt, err = db.Prepare("SELECT `name`,`group`,`is_super_admin`,`avatar`,`message`,`url_prefix`,`url_name`,`level` FROM `users` WHERE `uid` = ?")
	if err != nil {
		return err
	}
		
	log.Print("Preparing get_reply statement.")
	get_reply_stmt, err = db.Prepare("SELECT `tid`,`content`,`createdBy`,`createdAt`,`lastEdit`,`lastEditBy`,`ipaddress`,`likeCount` FROM `replies` WHERE `rid` = ?")
	if err != nil {
		return err
	}
		
	log.Print("Preparing get_user_reply statement.")
	get_user_reply_stmt, err = db.Prepare("SELECT `uid`,`content`,`createdBy`,`createdAt`,`lastEdit`,`lastEditBy`,`ipaddress` FROM `users_replies` WHERE `rid` = ?")
	if err != nil {
		return err
	}
		
	log.Print("Preparing get_password statement.")
	get_password_stmt, err = db.Prepare("SELECT `password`,`salt` FROM `users` WHERE `uid` = ?")
	if err != nil {
		return err
	}
		
	log.Print("Preparing get_settings statement.")
	get_settings_stmt, err = db.Prepare("SELECT `name`,`content`,`type` FROM `settings`")
	if err != nil {
		return err
	}
		
	log.Print("Preparing get_setting statement.")
	get_setting_stmt, err = db.Prepare("SELECT `content`,`type` FROM `settings` WHERE `name` = ?")
	if err != nil {
		return err
	}
		
	log.Print("Preparing get_full_setting statement.")
	get_full_setting_stmt, err = db.Prepare("SELECT `name`,`type`,`constraints` FROM `settings` WHERE `name` = ?")
	if err != nil {
		return err
	}
		
	log.Print("Preparing get_full_settings statement.")
	get_full_settings_stmt, err = db.Prepare("SELECT `name`,`content`,`type`,`constraints` FROM `settings`")
	if err != nil {
		return err
	}
		
	log.Print("Preparing get_groups statement.")
	get_groups_stmt, err = db.Prepare("SELECT `gid`,`name`,`permissions`,`plugin_perms`,`is_mod`,`is_admin`,`is_banned`,`tag` FROM `users_groups`")
	if err != nil {
		return err
	}
		
	log.Print("Preparing get_forums statement.")
	get_forums_stmt, err = db.Prepare("SELECT `fid`,`name`,`desc`,`active`,`preset`,`parentID`,`parentType`,`topicCount`,`lastTopic`,`lastTopicID`,`lastReplyer`,`lastReplyerID`,`lastTopicTime` FROM `forums` ORDER BY fid ASC")
	if err != nil {
		return err
	}
		
	log.Print("Preparing get_forums_permissions statement.")
	get_forums_permissions_stmt, err = db.Prepare("SELECT `gid`,`fid`,`permissions` FROM `forums_permissions` ORDER BY gid ASC,fid ASC")
	if err != nil {
		return err
	}
		
	log.Print("Preparing get_plugins statement.")
	get_plugins_stmt, err = db.Prepare("SELECT `uname`,`active`,`installed` FROM `plugins`")
	if err != nil {
		return err
	}
		
	log.Print("Preparing get_themes statement.")
	get_themes_stmt, err = db.Prepare("SELECT `uname`,`default` FROM `themes`")
	if err != nil {
		return err
	}
		
	log.Print("Preparing get_widgets statement.")
	get_widgets_stmt, err = db.Prepare("SELECT `position`,`side`,`type`,`active`,`location`,`data` FROM `widgets` ORDER BY position ASC")
	if err != nil {
		return err
	}
		
	log.Print("Preparing is_plugin_active statement.")
	is_plugin_active_stmt, err = db.Prepare("SELECT `active` FROM `plugins` WHERE `uname` = ?")
	if err != nil {
		return err
	}
		
	log.Print("Preparing get_users statement.")
	get_users_stmt, err = db.Prepare("SELECT `uid`,`name`,`group`,`active`,`is_super_admin`,`avatar` FROM `users`")
	if err != nil {
		return err
	}
		
	log.Print("Preparing get_users_offset statement.")
	get_users_offset_stmt, err = db.Prepare("SELECT `uid`,`name`,`group`,`active`,`is_super_admin`,`avatar` FROM `users` LIMIT ?,?")
	if err != nil {
		return err
	}
		
	log.Print("Preparing is_theme_default statement.")
	is_theme_default_stmt, err = db.Prepare("SELECT `default` FROM `themes` WHERE `uname` = ?")
	if err != nil {
		return err
	}
		
	log.Print("Preparing get_modlogs statement.")
	get_modlogs_stmt, err = db.Prepare("SELECT `action`,`elementID`,`elementType`,`ipaddress`,`actorID`,`doneAt` FROM `moderation_logs`")
	if err != nil {
		return err
	}
		
	log.Print("Preparing get_modlogs_offset statement.")
	get_modlogs_offset_stmt, err = db.Prepare("SELECT `action`,`elementID`,`elementType`,`ipaddress`,`actorID`,`doneAt` FROM `moderation_logs` LIMIT ?,?")
	if err != nil {
		return err
	}
		
	log.Print("Preparing get_reply_tid statement.")
	get_reply_tid_stmt, err = db.Prepare("SELECT `tid` FROM `replies` WHERE `rid` = ?")
	if err != nil {
		return err
	}
		
	log.Print("Preparing get_topic_fid statement.")
	get_topic_fid_stmt, err = db.Prepare("SELECT `parentID` FROM `topics` WHERE `tid` = ?")
	if err != nil {
		return err
	}
		
	log.Print("Preparing get_user_reply_uid statement.")
	get_user_reply_uid_stmt, err = db.Prepare("SELECT `uid` FROM `users_replies` WHERE `rid` = ?")
	if err != nil {
		return err
	}
		
	log.Print("Preparing has_liked_topic statement.")
	has_liked_topic_stmt, err = db.Prepare("SELECT `targetItem` FROM `likes` WHERE `sentBy` = ? AND `targetItem` = ? AND `targetType` = 'topics'")
	if err != nil {
		return err
	}
		
	log.Print("Preparing has_liked_reply statement.")
	has_liked_reply_stmt, err = db.Prepare("SELECT `targetItem` FROM `likes` WHERE `sentBy` = ? AND `targetItem` = ? AND `targetType` = 'replies'")
	if err != nil {
		return err
	}
		
	log.Print("Preparing get_user_name statement.")
	get_user_name_stmt, err = db.Prepare("SELECT `name` FROM `users` WHERE `uid` = ?")
	if err != nil {
		return err
	}
		
	log.Print("Preparing get_user_rank statement.")
	get_user_rank_stmt, err = db.Prepare("SELECT `group`,`is_super_admin` FROM `users` WHERE `uid` = ?")
	if err != nil {
		return err
	}
		
	log.Print("Preparing get_user_active statement.")
	get_user_active_stmt, err = db.Prepare("SELECT `active` FROM `users` WHERE `uid` = ?")
	if err != nil {
		return err
	}
		
	log.Print("Preparing get_user_group statement.")
	get_user_group_stmt, err = db.Prepare("SELECT `group` FROM `users` WHERE `uid` = ?")
	if err != nil {
		return err
	}
		
	log.Print("Preparing get_emails_by_user statement.")
	get_emails_by_user_stmt, err = db.Prepare("SELECT `email`,`validated`,`token` FROM `emails` WHERE `uid` = ?")
	if err != nil {
		return err
	}
		
	log.Print("Preparing get_topic_basic statement.")
	get_topic_basic_stmt, err = db.Prepare("SELECT `title`,`content` FROM `topics` WHERE `tid` = ?")
	if err != nil {
		return err
	}
		
	log.Print("Preparing get_activity_entry statement.")
	get_activity_entry_stmt, err = db.Prepare("SELECT `actor`,`targetUser`,`event`,`elementType`,`elementID` FROM `activity_stream` WHERE `asid` = ?")
	if err != nil {
		return err
	}
		
	log.Print("Preparing forum_entry_exists statement.")
	forum_entry_exists_stmt, err = db.Prepare("SELECT `fid` FROM `forums` WHERE `name` = '' ORDER BY fid ASC LIMIT 0,1")
	if err != nil {
		return err
	}
		
	log.Print("Preparing group_entry_exists statement.")
	group_entry_exists_stmt, err = db.Prepare("SELECT `gid` FROM `users_groups` WHERE `name` = '' ORDER BY gid ASC LIMIT 0,1")
	if err != nil {
		return err
	}
		
	log.Print("Preparing get_forum_topics_offset statement.")
	get_forum_topics_offset_stmt, err = db.Prepare("SELECT `tid`,`title`,`content`,`createdBy`,`is_closed`,`sticky`,`createdAt`,`lastReplyAt`,`lastReplyBy`,`parentID`,`postCount`,`likeCount` FROM `topics` WHERE `parentID` = ? ORDER BY sticky DESC,lastReplyAt DESC,createdBy DESC LIMIT ?,?")
	if err != nil {
		return err
	}
		
	log.Print("Preparing get_topic_replies_offset statement.")
	get_topic_replies_offset_stmt, err = db.Prepare("SELECT `replies`.`rid`,`replies`.`content`,`replies`.`createdBy`,`replies`.`createdAt`,`replies`.`lastEdit`,`replies`.`lastEditBy`,`users`.`avatar`,`users`.`name`,`users`.`group`,`users`.`url_prefix`,`users`.`url_name`,`users`.`level`,`replies`.`ipaddress`,`replies`.`likeCount`,`replies`.`actionType` FROM `replies` LEFT JOIN `users` ON `replies`.`createdBy` = `users`.`uid`  WHERE `tid` = ? LIMIT ?,?")
	if err != nil {
		return err
	}
		
	log.Print("Preparing get_topic_list statement.")
	get_topic_list_stmt, err = db.Prepare("SELECT `topics`.`tid`,`topics`.`title`,`topics`.`content`,`topics`.`createdBy`,`topics`.`is_closed`,`topics`.`sticky`,`topics`.`createdAt`,`topics`.`parentID`,`users`.`name`,`users`.`avatar` FROM `topics` LEFT JOIN `users` ON `topics`.`createdBy` = `users`.`uid`  ORDER BY topics.sticky DESC,topics.lastReplyAt DESC,topics.createdBy DESC")
	if err != nil {
		return err
	}
		
	log.Print("Preparing get_topic_user statement.")
	get_topic_user_stmt, err = db.Prepare("SELECT `topics`.`title`,`topics`.`content`,`topics`.`createdBy`,`topics`.`createdAt`,`topics`.`is_closed`,`topics`.`sticky`,`topics`.`parentID`,`topics`.`ipaddress`,`topics`.`postCount`,`topics`.`likeCount`,`users`.`name`,`users`.`avatar`,`users`.`group`,`users`.`url_prefix`,`users`.`url_name`,`users`.`level` FROM `topics` LEFT JOIN `users` ON `topics`.`createdBy` = `users`.`uid`  WHERE `tid` = ?")
	if err != nil {
		return err
	}
		
	log.Print("Preparing get_topic_by_reply statement.")
	get_topic_by_reply_stmt, err = db.Prepare("SELECT `topics`.`tid`,`topics`.`title`,`topics`.`content`,`topics`.`createdBy`,`topics`.`createdAt`,`topics`.`is_closed`,`topics`.`sticky`,`topics`.`parentID`,`topics`.`ipaddress`,`topics`.`postCount`,`topics`.`likeCount`,`topics`.`data` FROM `replies` LEFT JOIN `topics` ON `replies`.`tid` = `topics`.`tid`  WHERE `rid` = ?")
	if err != nil {
		return err
	}
		
	log.Print("Preparing get_topic_replies statement.")
	get_topic_replies_stmt, err = db.Prepare("SELECT `replies`.`rid`,`replies`.`content`,`replies`.`createdBy`,`replies`.`createdAt`,`replies`.`lastEdit`,`replies`.`lastEditBy`,`users`.`avatar`,`users`.`name`,`users`.`group`,`users`.`url_prefix`,`users`.`url_name`,`users`.`level`,`replies`.`ipaddress` FROM `replies` LEFT JOIN `users` ON `replies`.`createdBy` = `users`.`uid`  WHERE `tid` = ?")
	if err != nil {
		return err
	}
		
	log.Print("Preparing get_forum_topics statement.")
	get_forum_topics_stmt, err = db.Prepare("SELECT `topics`.`tid`,`topics`.`title`,`topics`.`content`,`topics`.`createdBy`,`topics`.`is_closed`,`topics`.`sticky`,`topics`.`createdAt`,`topics`.`lastReplyAt`,`topics`.`parentID`,`users`.`name`,`users`.`avatar` FROM `topics` LEFT JOIN `users` ON `topics`.`createdBy` = `users`.`uid`  WHERE `topics`.`parentID` = ? ORDER BY topics.sticky DESC,topics.lastReplyAt DESC,topics.createdBy DESC")
	if err != nil {
		return err
	}
		
	log.Print("Preparing get_profile_replies statement.")
	get_profile_replies_stmt, err = db.Prepare("SELECT `users_replies`.`rid`,`users_replies`.`content`,`users_replies`.`createdBy`,`users_replies`.`createdAt`,`users_replies`.`lastEdit`,`users_replies`.`lastEditBy`,`users`.`avatar`,`users`.`name`,`users`.`group` FROM `users_replies` LEFT JOIN `users` ON `users_replies`.`createdBy` = `users`.`uid`  WHERE `users_replies`.`uid` = ?")
	if err != nil {
		return err
	}
		
	log.Print("Preparing get_watchers statement.")
	get_watchers_stmt, err = db.Prepare("SELECT `activity_subscriptions`.`user` FROM `activity_stream` INNER JOIN `activity_subscriptions` ON `activity_subscriptions`.`targetType` = `activity_stream`.`elementType` AND `activity_subscriptions`.`targetID` = `activity_stream`.`elementID` AND `activity_subscriptions`.`user` != `activity_stream`.`actor`  WHERE `asid` = ?")
	if err != nil {
		return err
	}
		
	log.Print("Preparing create_topic statement.")
	create_topic_stmt, err = db.Prepare("INSERT INTO `topics`(`parentID`,`title`,`content`,`parsed_content`,`createdAt`,`lastReplyAt`,`lastReplyBy`,`ipaddress`,`words`,`createdBy`) VALUES (?,?,?,?,UTC_TIMESTAMP(),UTC_TIMESTAMP(),?,?,?,?)")
	if err != nil {
		return err
	}
		
	log.Print("Preparing create_report statement.")
	create_report_stmt, err = db.Prepare("INSERT INTO `topics`(`title`,`content`,`parsed_content`,`createdAt`,`lastReplyAt`,`createdBy`,`data`,`parentID`,`css_class`) VALUES (?,?,?,UTC_TIMESTAMP(),UTC_TIMESTAMP(),?,?,1,'report')")
	if err != nil {
		return err
	}
		
	log.Print("Preparing create_reply statement.")
	create_reply_stmt, err = db.Prepare("INSERT INTO `replies`(`tid`,`content`,`parsed_content`,`createdAt`,`ipaddress`,`words`,`createdBy`) VALUES (?,?,?,UTC_TIMESTAMP(),?,?,?)")
	if err != nil {
		return err
	}
		
	log.Print("Preparing create_action_reply statement.")
	create_action_reply_stmt, err = db.Prepare("INSERT INTO `replies`(`tid`,`actionType`,`ipaddress`,`createdBy`) VALUES (?,?,?,?)")
	if err != nil {
		return err
	}
		
	log.Print("Preparing create_like statement.")
	create_like_stmt, err = db.Prepare("INSERT INTO `likes`(`weight`,`targetItem`,`targetType`,`sentBy`) VALUES (?,?,?,?)")
	if err != nil {
		return err
	}
		
	log.Print("Preparing add_activity statement.")
	add_activity_stmt, err = db.Prepare("INSERT INTO `activity_stream`(`actor`,`targetUser`,`event`,`elementType`,`elementID`) VALUES (?,?,?,?,?)")
	if err != nil {
		return err
	}
		
	log.Print("Preparing notify_one statement.")
	notify_one_stmt, err = db.Prepare("INSERT INTO `activity_stream_matches`(`watcher`,`asid`) VALUES (?,?)")
	if err != nil {
		return err
	}
		
	log.Print("Preparing add_email statement.")
	add_email_stmt, err = db.Prepare("INSERT INTO `emails`(`email`,`uid`,`validated`,`token`) VALUES (?,?,?,?)")
	if err != nil {
		return err
	}
		
	log.Print("Preparing create_profile_reply statement.")
	create_profile_reply_stmt, err = db.Prepare("INSERT INTO `users_replies`(`uid`,`content`,`parsed_content`,`createdAt`,`createdBy`,`ipaddress`) VALUES (?,?,?,UTC_TIMESTAMP(),?,?)")
	if err != nil {
		return err
	}
		
	log.Print("Preparing add_subscription statement.")
	add_subscription_stmt, err = db.Prepare("INSERT INTO `activity_subscriptions`(`user`,`targetID`,`targetType`,`level`) VALUES (?,?,?,2)")
	if err != nil {
		return err
	}
		
	log.Print("Preparing create_forum statement.")
	create_forum_stmt, err = db.Prepare("INSERT INTO `forums`(`name`,`desc`,`active`,`preset`) VALUES (?,?,?,?)")
	if err != nil {
		return err
	}
		
	log.Print("Preparing add_forum_perms_to_forum statement.")
	add_forum_perms_to_forum_stmt, err = db.Prepare("INSERT INTO `forums_permissions`(`gid`,`fid`,`preset`,`permissions`) VALUES (?,?,?,?)")
	if err != nil {
		return err
	}
		
	log.Print("Preparing add_plugin statement.")
	add_plugin_stmt, err = db.Prepare("INSERT INTO `plugins`(`uname`,`active`,`installed`) VALUES (?,?,?)")
	if err != nil {
		return err
	}
		
	log.Print("Preparing add_theme statement.")
	add_theme_stmt, err = db.Prepare("INSERT INTO `themes`(`uname`,`default`) VALUES (?,?)")
	if err != nil {
		return err
	}
		
	log.Print("Preparing create_group statement.")
	create_group_stmt, err = db.Prepare("INSERT INTO `users_groups`(`name`,`tag`,`is_admin`,`is_mod`,`is_banned`,`permissions`) VALUES (?,?,?,?,?,?)")
	if err != nil {
		return err
	}
		
	log.Print("Preparing add_modlog_entry statement.")
	add_modlog_entry_stmt, err = db.Prepare("INSERT INTO `moderation_logs`(`action`,`elementID`,`elementType`,`ipaddress`,`actorID`,`doneAt`) VALUES (?,?,?,?,?,UTC_TIMESTAMP())")
	if err != nil {
		return err
	}
		
	log.Print("Preparing add_adminlog_entry statement.")
	add_adminlog_entry_stmt, err = db.Prepare("INSERT INTO `administration_logs`(`action`,`elementID`,`elementType`,`ipaddress`,`actorID`,`doneAt`) VALUES (?,?,?,?,?,UTC_TIMESTAMP())")
	if err != nil {
		return err
	}
		
	log.Print("Preparing add_forum_perms_to_group statement.")
	add_forum_perms_to_group_stmt, err = db.Prepare("REPLACE INTO `forums_permissions`(`gid`,`fid`,`preset`,`permissions`) VALUES (?,?,?,?)")
	if err != nil {
		return err
	}
		
	log.Print("Preparing add_replies_to_topic statement.")
	add_replies_to_topic_stmt, err = db.Prepare("UPDATE `topics` SET `postCount` = `postCount` + ?,`lastReplyBy` = ?,`lastReplyAt` = UTC_TIMESTAMP() WHERE `tid` = ?")
	if err != nil {
		return err
	}
		
	log.Print("Preparing remove_replies_from_topic statement.")
	remove_replies_from_topic_stmt, err = db.Prepare("UPDATE `topics` SET `postCount` = `postCount` - ? WHERE `tid` = ?")
	if err != nil {
		return err
	}
		
	log.Print("Preparing add_topics_to_forum statement.")
	add_topics_to_forum_stmt, err = db.Prepare("UPDATE `forums` SET `topicCount` = `topicCount` + ? WHERE `fid` = ?")
	if err != nil {
		return err
	}
		
	log.Print("Preparing remove_topics_from_forum statement.")
	remove_topics_from_forum_stmt, err = db.Prepare("UPDATE `forums` SET `topicCount` = `topicCount` - ? WHERE `fid` = ?")
	if err != nil {
		return err
	}
		
	log.Print("Preparing update_forum_cache statement.")
	update_forum_cache_stmt, err = db.Prepare("UPDATE `forums` SET `lastTopic` = ?,`lastTopicID` = ?,`lastReplyer` = ?,`lastReplyerID` = ?,`lastTopicTime` = UTC_TIMESTAMP() WHERE `fid` = ?")
	if err != nil {
		return err
	}
		
	log.Print("Preparing add_likes_to_topic statement.")
	add_likes_to_topic_stmt, err = db.Prepare("UPDATE `topics` SET `likeCount` = `likeCount` + ? WHERE `tid` = ?")
	if err != nil {
		return err
	}
		
	log.Print("Preparing add_likes_to_reply statement.")
	add_likes_to_reply_stmt, err = db.Prepare("UPDATE `replies` SET `likeCount` = `likeCount` + ? WHERE `rid` = ?")
	if err != nil {
		return err
	}
		
	log.Print("Preparing edit_topic statement.")
	edit_topic_stmt, err = db.Prepare("UPDATE `topics` SET `title` = ?,`content` = ?,`parsed_content` = ?,`is_closed` = ? WHERE `tid` = ?")
	if err != nil {
		return err
	}
		
	log.Print("Preparing edit_reply statement.")
	edit_reply_stmt, err = db.Prepare("UPDATE `replies` SET `content` = ?,`parsed_content` = ? WHERE `rid` = ?")
	if err != nil {
		return err
	}
		
	log.Print("Preparing stick_topic statement.")
	stick_topic_stmt, err = db.Prepare("UPDATE `topics` SET `sticky` = 1 WHERE `tid` = ?")
	if err != nil {
		return err
	}
		
	log.Print("Preparing unstick_topic statement.")
	unstick_topic_stmt, err = db.Prepare("UPDATE `topics` SET `sticky` = 0 WHERE `tid` = ?")
	if err != nil {
		return err
	}
		
	log.Print("Preparing update_last_ip statement.")
	update_last_ip_stmt, err = db.Prepare("UPDATE `users` SET `last_ip` = ? WHERE `uid` = ?")
	if err != nil {
		return err
	}
		
	log.Print("Preparing update_session statement.")
	update_session_stmt, err = db.Prepare("UPDATE `users` SET `session` = ? WHERE `uid` = ?")
	if err != nil {
		return err
	}
		
	log.Print("Preparing set_password statement.")
	set_password_stmt, err = db.Prepare("UPDATE `users` SET `password` = ?,`salt` = ? WHERE `uid` = ?")
	if err != nil {
		return err
	}
		
	log.Print("Preparing set_avatar statement.")
	set_avatar_stmt, err = db.Prepare("UPDATE `users` SET `avatar` = ? WHERE `uid` = ?")
	if err != nil {
		return err
	}
		
	log.Print("Preparing set_username statement.")
	set_username_stmt, err = db.Prepare("UPDATE `users` SET `name` = ? WHERE `uid` = ?")
	if err != nil {
		return err
	}
		
	log.Print("Preparing change_group statement.")
	change_group_stmt, err = db.Prepare("UPDATE `users` SET `group` = ? WHERE `uid` = ?")
	if err != nil {
		return err
	}
		
	log.Print("Preparing activate_user statement.")
	activate_user_stmt, err = db.Prepare("UPDATE `users` SET `active` = 1 WHERE `uid` = ?")
	if err != nil {
		return err
	}
		
	log.Print("Preparing update_user_level statement.")
	update_user_level_stmt, err = db.Prepare("UPDATE `users` SET `level` = ? WHERE `uid` = ?")
	if err != nil {
		return err
	}
		
	log.Print("Preparing increment_user_score statement.")
	increment_user_score_stmt, err = db.Prepare("UPDATE `users` SET `score` = `score` + ? WHERE `uid` = ?")
	if err != nil {
		return err
	}
		
	log.Print("Preparing increment_user_posts statement.")
	increment_user_posts_stmt, err = db.Prepare("UPDATE `users` SET `posts` = `posts` + ? WHERE `uid` = ?")
	if err != nil {
		return err
	}
		
	log.Print("Preparing increment_user_bigposts statement.")
	increment_user_bigposts_stmt, err = db.Prepare("UPDATE `users` SET `posts` = `posts` + ?,`bigposts` = `bigposts` + ? WHERE `uid` = ?")
	if err != nil {
		return err
	}
		
	log.Print("Preparing increment_user_megaposts statement.")
	increment_user_megaposts_stmt, err = db.Prepare("UPDATE `users` SET `posts` = `posts` + ?,`bigposts` = `bigposts` + ?,`megaposts` = `megaposts` + ? WHERE `uid` = ?")
	if err != nil {
		return err
	}
		
	log.Print("Preparing increment_user_topics statement.")
	increment_user_topics_stmt, err = db.Prepare("UPDATE `users` SET `topics` = `topics` + ? WHERE `uid` = ?")
	if err != nil {
		return err
	}
		
	log.Print("Preparing edit_profile_reply statement.")
	edit_profile_reply_stmt, err = db.Prepare("UPDATE `users_replies` SET `content` = ?,`parsed_content` = ? WHERE `rid` = ?")
	if err != nil {
		return err
	}
		
	log.Print("Preparing delete_forum statement.")
	delete_forum_stmt, err = db.Prepare("UPDATE `forums` SET `name` = '',`active` = 0 WHERE `fid` = ?")
	if err != nil {
		return err
	}
		
	log.Print("Preparing update_forum statement.")
	update_forum_stmt, err = db.Prepare("UPDATE `forums` SET `name` = ?,`desc` = ?,`active` = ?,`preset` = ? WHERE `fid` = ?")
	if err != nil {
		return err
	}
		
	log.Print("Preparing update_setting statement.")
	update_setting_stmt, err = db.Prepare("UPDATE `settings` SET `content` = ? WHERE `name` = ?")
	if err != nil {
		return err
	}
		
	log.Print("Preparing update_plugin statement.")
	update_plugin_stmt, err = db.Prepare("UPDATE `plugins` SET `active` = ? WHERE `uname` = ?")
	if err != nil {
		return err
	}
		
	log.Print("Preparing update_plugin_install statement.")
	update_plugin_install_stmt, err = db.Prepare("UPDATE `plugins` SET `installed` = ? WHERE `uname` = ?")
	if err != nil {
		return err
	}
		
	log.Print("Preparing update_theme statement.")
	update_theme_stmt, err = db.Prepare("UPDATE `themes` SET `default` = ? WHERE `uname` = ?")
	if err != nil {
		return err
	}
		
	log.Print("Preparing update_user statement.")
	update_user_stmt, err = db.Prepare("UPDATE `users` SET `name` = ?,`email` = ?,`group` = ? WHERE `uid` = ?")
	if err != nil {
		return err
	}
		
	log.Print("Preparing update_group_perms statement.")
	update_group_perms_stmt, err = db.Prepare("UPDATE `users_groups` SET `permissions` = ? WHERE `gid` = ?")
	if err != nil {
		return err
	}
		
	log.Print("Preparing update_group_rank statement.")
	update_group_rank_stmt, err = db.Prepare("UPDATE `users_groups` SET `is_admin` = ?,`is_mod` = ?,`is_banned` = ? WHERE `gid` = ?")
	if err != nil {
		return err
	}
		
	log.Print("Preparing update_group statement.")
	update_group_stmt, err = db.Prepare("UPDATE `users_groups` SET `name` = ?,`tag` = ? WHERE `gid` = ?")
	if err != nil {
		return err
	}
		
	log.Print("Preparing update_email statement.")
	update_email_stmt, err = db.Prepare("UPDATE `emails` SET `email` = ?,`uid` = ?,`validated` = ?,`token` = ? WHERE `email` = ?")
	if err != nil {
		return err
	}
		
	log.Print("Preparing verify_email statement.")
	verify_email_stmt, err = db.Prepare("UPDATE `emails` SET `validated` = 1,`token` = '' WHERE `email` = ?")
	if err != nil {
		return err
	}
		
	log.Print("Preparing delete_reply statement.")
	delete_reply_stmt, err = db.Prepare("DELETE FROM `replies` WHERE `rid` = ?")
	if err != nil {
		return err
	}
		
	log.Print("Preparing delete_topic statement.")
	delete_topic_stmt, err = db.Prepare("DELETE FROM `topics` WHERE `tid` = ?")
	if err != nil {
		return err
	}
		
	log.Print("Preparing delete_profile_reply statement.")
	delete_profile_reply_stmt, err = db.Prepare("DELETE FROM `users_replies` WHERE `rid` = ?")
	if err != nil {
		return err
	}
		
	log.Print("Preparing delete_forum_perms_by_forum statement.")
	delete_forum_perms_by_forum_stmt, err = db.Prepare("DELETE FROM `forums_permissions` WHERE `fid` = ?")
	if err != nil {
		return err
	}
		
	log.Print("Preparing report_exists statement.")
	report_exists_stmt, err = db.Prepare("SELECT COUNT(*) AS `count` FROM `topics` WHERE `data` = ? AND `data` != '' AND `parentID` = 1")
	if err != nil {
		return err
	}
		
	log.Print("Preparing group_count statement.")
	group_count_stmt, err = db.Prepare("SELECT COUNT(*) AS `count` FROM `users_groups`")
	if err != nil {
		return err
	}
		
	log.Print("Preparing modlog_count statement.")
	modlog_count_stmt, err = db.Prepare("SELECT COUNT(*) AS `count` FROM `moderation_logs`")
	if err != nil {
		return err
	}
		
	log.Print("Preparing add_forum_perms_to_forum_admins statement.")
	add_forum_perms_to_forum_admins_stmt, err = db.Prepare("INSERT INTO `forums_permissions`(`gid`,`fid`,`preset`,`permissions`) SELECT `gid`, ? AS `fid`, ? AS `preset`, ? AS `permissions` FROM `users_groups` WHERE `is_admin` = 1")
	if err != nil {
		return err
	}
		
	log.Print("Preparing add_forum_perms_to_forum_staff statement.")
	add_forum_perms_to_forum_staff_stmt, err = db.Prepare("INSERT INTO `forums_permissions`(`gid`,`fid`,`preset`,`permissions`) SELECT `gid`, ? AS `fid`, ? AS `preset`, ? AS `permissions` FROM `users_groups` WHERE `is_admin` = 0 AND `is_mod` = 1")
	if err != nil {
		return err
	}
		
	log.Print("Preparing add_forum_perms_to_forum_members statement.")
	add_forum_perms_to_forum_members_stmt, err = db.Prepare("INSERT INTO `forums_permissions`(`gid`,`fid`,`preset`,`permissions`) SELECT `gid`, ? AS `fid`, ? AS `preset`, ? AS `permissions` FROM `users_groups` WHERE `is_admin` = 0 AND `is_mod` = 0 AND `is_banned` = 0")
	if err != nil {
		return err
	}
		
	log.Print("Preparing notify_watchers statement.")
	notify_watchers_stmt, err = db.Prepare("INSERT INTO `activity_stream_matches`(`watcher`,`asid`) SELECT `activity_subscriptions`.`user`, `activity_stream`.`asid` FROM `activity_stream` INNER JOIN `activity_subscriptions` ON `activity_subscriptions`.`targetType` = `activity_stream`.`elementType` AND `activity_subscriptions`.`targetID` = `activity_stream`.`elementID` AND `activity_subscriptions`.`user` != `activity_stream`.`actor`  WHERE `asid` = ?")
	if err != nil {
		return err
	}
	
	return nil
}
