/* WIP Under Construction */
package main

import (
	"fmt"
	"log"
	"os"
	"runtime/debug"

	"./lib"
)

// TODO: Make sure all the errors in this file propagate upwards properly
func main() {
	// Capture panics instead of closing the window at a superhuman speed before the user can read the message on Windows
	defer func() {
		r := recover()
		if r != nil {
			fmt.Println(r)
			debug.PrintStack()
			return
		}
	}()

	log.Println("Running the query generator")
	for _, adapter := range qgen.Registry {
		log.Printf("Building the queries for the %s adapter", adapter.GetName())
		qgen.Install.SetAdapterInstance(adapter)
		qgen.Install.RegisterPlugin(NewPrimaryKeySpitter()) // TODO: Do we really need to fill the spitter for every adapter?

		err := writeStatements(adapter)
		if err != nil {
			log.Print(err)
		}
		err = qgen.Install.Write()
		if err != nil {
			log.Print(err)
		}
		err = adapter.Write()
		if err != nil {
			log.Print(err)
		}
	}
}

// nolint
func writeStatements(adapter qgen.Adapter) error {
	err := createTables(adapter)
	if err != nil {
		return err
	}

	err = seedTables(adapter)
	if err != nil {
		return err
	}

	err = writeSelects(adapter)
	if err != nil {
		return err
	}

	err = writeLeftJoins(adapter)
	if err != nil {
		return err
	}

	err = writeInnerJoins(adapter)
	if err != nil {
		return err
	}

	err = writeInserts(adapter)
	if err != nil {
		return err
	}

	/*err = writeReplaces(adapter)
	if err != nil {
		return err
	}

	err = writeUpserts(adapter)
	if err != nil {
		return err
	}*/

	err = writeUpdates(adapter)
	if err != nil {
		return err
	}

	err = writeDeletes(adapter)
	if err != nil {
		return err
	}

	err = writeSimpleCounts(adapter)
	if err != nil {
		return err
	}

	err = writeInsertSelects(adapter)
	if err != nil {
		return err
	}

	err = writeInsertLeftJoins(adapter)
	if err != nil {
		return err
	}

	err = writeInsertInnerJoins(adapter)
	if err != nil {
		return err
	}

	return nil
}

func seedTables(adapter qgen.Adapter) error {
	qgen.Install.SimpleInsert("sync", "last_update", "UTC_TIMESTAMP()")

	qgen.Install.SimpleInsert("settings", "name, content, type", "'url_tags','1','bool'")
	qgen.Install.SimpleInsert("settings", "name, content, type, constraints", "'activation_type','1','list','1-3'")
	qgen.Install.SimpleInsert("settings", "name, content, type", "'bigpost_min_words','250','int'")
	qgen.Install.SimpleInsert("settings", "name, content, type", "'megapost_min_words','1000','int'")
	qgen.Install.SimpleInsert("themes", "uname, default", "'tempra-simple',1")
	qgen.Install.SimpleInsert("emails", "email, uid, validated", "'admin@localhost',1,1") // ? - Use a different default email or let the admin input it during installation?

	/*
		The Permissions:

		Global Permissions:
		BanUsers
		ActivateUsers
		EditUser
		EditUserEmail
		EditUserPassword
		EditUserGroup
		EditUserGroupSuperMod
		EditUserGroupAdmin
		EditGroup
		EditGroupLocalPerms
		EditGroupGlobalPerms
		EditGroupSuperMod
		EditGroupAdmin
		ManageForums
		EditSettings
		ManageThemes
		ManagePlugins
		ViewAdminLogs
		ViewIPs

		Non-staff Global Permissions:
		UploadFiles

		Forum Permissions:
		ViewTopic
		LikeItem
		CreateTopic
		EditTopic
		DeleteTopic
		CreateReply
		EditReply
		DeleteReply
		PinTopic
		CloseTopic
	*/

	qgen.Install.SimpleInsert("users_groups", "name, permissions, plugin_perms, is_mod, is_admin, tag", `'Administrator','{"BanUsers":true,"ActivateUsers":true,"EditUser":true,"EditUserEmail":true,"EditUserPassword":true,"EditUserGroup":true,"EditUserGroupSuperMod":true,"EditUserGroupAdmin":false,"EditGroup":true,"EditGroupLocalPerms":true,"EditGroupGlobalPerms":true,"EditGroupSuperMod":true,"EditGroupAdmin":false,"ManageForums":true,"EditSettings":true,"ManageThemes":true,"ManagePlugins":true,"ViewAdminLogs":true,"ViewIPs":true,"UploadFiles":true,"ViewTopic":true,"LikeItem":true,"CreateTopic":true,"EditTopic":true,"DeleteTopic":true,"CreateReply":true,"EditReply":true,"DeleteReply":true,"PinTopic":true,"CloseTopic":true}','{}',1,1,"Admin"`)

	qgen.Install.SimpleInsert("users_groups", "name, permissions, plugin_perms, is_mod, tag", `'Moderator','{"BanUsers":true,"ActivateUsers":false,"EditUser":true,"EditUserEmail":false,"EditUserGroup":true,"ViewIPs":true,"UploadFiles":true,"ViewTopic":true,"LikeItem":true,"CreateTopic":true,"EditTopic":true,"DeleteTopic":true,"CreateReply":true,"EditReply":true,"DeleteReply":true,"PinTopic":true,"CloseTopic":true}','{}',1,"Mod"`)

	qgen.Install.SimpleInsert("users_groups", "name, permissions, plugin_perms", `'Member','{"UploadFiles":true,"ViewTopic":true,"LikeItem":true,"CreateTopic":true,"CreateReply":true}','{}'`)

	qgen.Install.SimpleInsert("users_groups", "name, permissions, plugin_perms, is_banned", `'Banned','{"ViewTopic":true}','{}',1`)

	qgen.Install.SimpleInsert("users_groups", "name, permissions, plugin_perms", `'Awaiting Activation','{"ViewTopic":true}','{}'`)

	qgen.Install.SimpleInsert("users_groups", "name, permissions, plugin_perms, tag", `'Not Loggedin','{"ViewTopic":true}','{}','Guest'`)

	//
	// TODO: Stop processFields() from stripping the spaces in the descriptions in the next commit

	qgen.Install.SimpleInsert("forums", "name, active, desc", "'Reports',0,'All the reports go here'")

	qgen.Install.SimpleInsert("forums", "name, lastTopicID, lastReplyerID, desc", "'General',1,1,'A place for general discussions which don't fit elsewhere'")

	//

	qgen.Install.SimpleInsert("forums_permissions", "gid, fid, permissions", `1,1,'{"ViewTopic":true,"CreateReply":true,"CreateTopic":true,"PinTopic":true,"CloseTopic":true}'`)
	qgen.Install.SimpleInsert("forums_permissions", "gid, fid, permissions", `2,1,'{"ViewTopic":true,"CreateReply":true,"CloseTopic":true}'`)
	qgen.Install.SimpleInsert("forums_permissions", "gid, fid, permissions", "3,1,'{}'")
	qgen.Install.SimpleInsert("forums_permissions", "gid, fid, permissions", "4,1,'{}'")
	qgen.Install.SimpleInsert("forums_permissions", "gid, fid, permissions", "5,1,'{}'")
	qgen.Install.SimpleInsert("forums_permissions", "gid, fid, permissions", "6,1,'{}'")

	//

	qgen.Install.SimpleInsert("forums_permissions", "gid, fid, permissions", `1,2,'{"ViewTopic":true,"CreateReply":true,"CreateTopic":true,"LikeItem":true,"EditTopic":true,"DeleteTopic":true,"EditReply":true,"DeleteReply":true,"PinTopic":true,"CloseTopic":true}'`)

	qgen.Install.SimpleInsert("forums_permissions", "gid, fid, permissions", `2,2,'{"ViewTopic":true,"CreateReply":true,"CreateTopic":true,"LikeItem":true,"EditTopic":true,"DeleteTopic":true,"EditReply":true,"DeleteReply":true,"PinTopic":true,"CloseTopic":true}'`)

	qgen.Install.SimpleInsert("forums_permissions", "gid, fid, permissions", `3,2,'{"ViewTopic":true,"CreateReply":true,"CreateTopic":true,"LikeItem":true}'`)

	qgen.Install.SimpleInsert("forums_permissions", "gid, fid, permissions", `4,2,'{"ViewTopic":true}'`)

	qgen.Install.SimpleInsert("forums_permissions", "gid, fid, permissions", `5,2,'{"ViewTopic":true}'`)

	qgen.Install.SimpleInsert("forums_permissions", "gid, fid, permissions", `6,2,'{"ViewTopic":true}'`)

	//

	qgen.Install.SimpleInsert("topics", "title, content, parsed_content, createdAt, lastReplyAt, lastReplyBy, createdBy, parentID, ipaddress", "'Test Topic','A topic automatically generated by the software.','A topic automatically generated by the software.',UTC_TIMESTAMP(),UTC_TIMESTAMP(),1,1,2,'::1'")

	qgen.Install.SimpleInsert("replies", "tid, content, parsed_content, createdAt, createdBy, lastUpdated, lastEdit, lastEditBy, ipaddress", "1,'A reply!','A reply!',UTC_TIMESTAMP(),1,UTC_TIMESTAMP(),0,0,'::1'")

	return nil
}

func writeSelects(adapter qgen.Adapter) error {
	// Looking for getTopic? Your statement is in another castle

	adapter.Select("getPassword").Table("users").Columns("password, salt").Where("uid = ?").Parse()

	adapter.Select("getSettings").Table("settings").Columns("name, content, type").Parse()

	adapter.Select("getSetting").Table("settings").Columns("content, type").Where("name = ?").Parse()

	adapter.Select("getFullSetting").Table("settings").Columns("name, type, constraints").Where("name = ?").Parse()

	adapter.Select("isPluginActive").Table("plugins").Columns("active").Where("uname = ?").Parse()

	//adapter.SimpleSelect("isPluginInstalled","plugins","installed","uname = ?","","")

	adapter.Select("getUsersOffset").Table("users").Columns("uid, name, group, active, is_super_admin, avatar").Orderby("uid ASC").Limit("?,?").Parse()

	adapter.SimpleSelect("isThemeDefault", "themes", "default", "uname = ?", "", "")

	adapter.SimpleSelect("getModlogs", "moderation_logs", "action, elementID, elementType, ipaddress, actorID, doneAt", "", "", "")

	adapter.SimpleSelect("getModlogsOffset", "moderation_logs", "action, elementID, elementType, ipaddress, actorID, doneAt", "", "doneAt DESC", "?,?")

	adapter.SimpleSelect("getReplyTID", "replies", "tid", "rid = ?", "", "")

	adapter.SimpleSelect("getTopicFID", "topics", "parentID", "tid = ?", "", "")

	adapter.SimpleSelect("getUserReplyUID", "users_replies", "uid", "rid = ?", "", "")

	adapter.SimpleSelect("getUserName", "users", "name", "uid = ?", "", "")

	adapter.SimpleSelect("getEmailsByUser", "emails", "email, validated, token", "uid = ?", "", "")

	adapter.SimpleSelect("getTopicBasic", "topics", "title, content", "tid = ?", "", "")

	adapter.SimpleSelect("getActivityEntry", "activity_stream", "actor, targetUser, event, elementType, elementID", "asid = ?", "", "")

	adapter.SimpleSelect("forumEntryExists", "forums", "fid", "name = ''", "fid ASC", "0,1")

	adapter.SimpleSelect("groupEntryExists", "users_groups", "gid", "name = ''", "gid ASC", "0,1")

	adapter.SimpleSelect("getForumTopicsOffset", "topics", "tid, title, content, createdBy, is_closed, sticky, createdAt, lastReplyAt, lastReplyBy, parentID, postCount, likeCount", "parentID = ?", "sticky DESC, lastReplyAt DESC, createdBy DESC", "?,?")

	adapter.SimpleSelect("getAttachment", "attachments", "sectionID, sectionTable, originID, originTable, uploadedBy, path", "path = ? AND sectionID = ? AND sectionTable = ?", "", "")

	return nil
}

func writeLeftJoins(adapter qgen.Adapter) error {
	adapter.SimpleLeftJoin("getTopicRepliesOffset", "replies", "users", "replies.rid, replies.content, replies.createdBy, replies.createdAt, replies.lastEdit, replies.lastEditBy, users.avatar, users.name, users.group, users.url_prefix, users.url_name, users.level, replies.ipaddress, replies.likeCount, replies.actionType", "replies.createdBy = users.uid", "replies.tid = ?", "replies.rid ASC", "?,?")

	adapter.SimpleLeftJoin("getTopicList", "topics", "users", "topics.tid, topics.title, topics.content, topics.createdBy, topics.is_closed, topics.sticky, topics.createdAt, topics.parentID, users.name, users.avatar", "topics.createdBy = users.uid", "", "topics.sticky DESC, topics.lastReplyAt DESC, topics.createdBy DESC", "")

	adapter.SimpleLeftJoin("getTopicReplies", "replies", "users", "replies.rid, replies.content, replies.createdBy, replies.createdAt, replies.lastEdit, replies.lastEditBy, users.avatar, users.name, users.group, users.url_prefix, users.url_name, users.level, replies.ipaddress", "replies.createdBy = users.uid", "tid = ?", "", "")

	adapter.SimpleLeftJoin("getForumTopics", "topics", "users", "topics.tid, topics.title, topics.content, topics.createdBy, topics.is_closed, topics.sticky, topics.createdAt, topics.lastReplyAt, topics.parentID, users.name, users.avatar", "topics.createdBy = users.uid", "topics.parentID = ?", "topics.sticky DESC, topics.lastReplyAt DESC, topics.createdBy desc", "")

	adapter.SimpleLeftJoin("getProfileReplies", "users_replies", "users", "users_replies.rid, users_replies.content, users_replies.createdBy, users_replies.createdAt, users_replies.lastEdit, users_replies.lastEditBy, users.avatar, users.name, users.group", "users_replies.createdBy = users.uid", "users_replies.uid = ?", "", "")

	return nil
}

func writeInnerJoins(adapter qgen.Adapter) (err error) {
	_, err = adapter.SimpleInnerJoin("getWatchers", "activity_stream", "activity_subscriptions", "activity_subscriptions.user", "activity_subscriptions.targetType = activity_stream.elementType AND activity_subscriptions.targetID = activity_stream.elementID AND activity_subscriptions.user != activity_stream.actor", "asid = ?", "", "")
	if err != nil {
		return err
	}

	return nil
}

func writeInserts(adapter qgen.Adapter) error {
	adapter.SimpleInsert("createReport", "topics", "title, content, parsed_content, createdAt, lastReplyAt, createdBy, lastReplyBy, data, parentID, css_class", "?,?,?,UTC_TIMESTAMP(),UTC_TIMESTAMP(),?,?,?,1,'report'")

	adapter.SimpleInsert("addActivity", "activity_stream", "actor, targetUser, event, elementType, elementID", "?,?,?,?,?")

	adapter.SimpleInsert("notifyOne", "activity_stream_matches", "watcher, asid", "?,?")

	adapter.SimpleInsert("addEmail", "emails", "email, uid, validated, token", "?,?,?,?")

	adapter.SimpleInsert("addSubscription", "activity_subscriptions", "user, targetID, targetType, level", "?,?,?,2")

	adapter.SimpleInsert("addForumPermsToForum", "forums_permissions", "gid,fid,preset,permissions", "?,?,?,?")

	adapter.SimpleInsert("addPlugin", "plugins", "uname, active, installed", "?,?,?")

	adapter.SimpleInsert("addTheme", "themes", "uname, default", "?,?")

	adapter.SimpleInsert("addAttachment", "attachments", "sectionID, sectionTable, originID, originTable, uploadedBy, path", "?,?,?,?,?,?")

	adapter.SimpleInsert("createWordFilter", "word_filters", "find, replacement", "?,?")

	return nil
}

func writeReplaces(adapter qgen.Adapter) (err error) {
	return nil
}

// ! Upserts are broken atm
/*func writeUpserts(adapter qgen.Adapter) (err error) {
	_, err = adapter.SimpleUpsert("addForumPermsToGroup", "forums_permissions", "gid, fid, preset, permissions", "?,?,?,?", "gid = ? AND fid = ?")
	if err != nil {
		return err
	}

	_, err = adapter.SimpleUpsert("replaceScheduleGroup", "users_groups_scheduler", "uid, set_group, issued_by, issued_at, revert_at, temporary", "?,?,?,UTC_TIMESTAMP(),?,?", "uid = ?")
	if err != nil {
		return err
	}

	return nil
}*/

func writeUpdates(adapter qgen.Adapter) error {
	adapter.SimpleUpdate("editReply", "replies", "content = ?, parsed_content = ?", "rid = ?")

	adapter.SimpleUpdate("editProfileReply", "users_replies", "content = ?, parsed_content = ?", "rid = ?")

	adapter.SimpleUpdate("updateSetting", "settings", "content = ?", "name = ?")

	adapter.SimpleUpdate("updatePlugin", "plugins", "active = ?", "uname = ?")

	adapter.SimpleUpdate("updatePluginInstall", "plugins", "installed = ?", "uname = ?")

	adapter.SimpleUpdate("updateTheme", "themes", "default = ?", "uname = ?")

	adapter.SimpleUpdate("updateForum", "forums", "name = ?, desc = ?, active = ?, preset = ?", "fid = ?")

	adapter.SimpleUpdate("updateUser", "users", "name = ?, email = ?, group = ?", "uid = ?")

	adapter.SimpleUpdate("updateGroupPerms", "users_groups", "permissions = ?", "gid = ?")

	adapter.SimpleUpdate("updateGroup", "users_groups", "name = ?, tag = ?", "gid = ?")

	adapter.SimpleUpdate("updateEmail", "emails", "email = ?, uid = ?, validated = ?, token = ?", "email = ?")

	adapter.SimpleUpdate("verifyEmail", "emails", "validated = 1, token = ''", "email = ?") // Need to fix this: Empty string isn't working, it gets set to 1 instead x.x -- Has this been fixed?

	adapter.SimpleUpdate("setTempGroup", "users", "temp_group = ?", "uid = ?")

	adapter.SimpleUpdate("updateWordFilter", "word_filters", "find = ?, replacement = ?", "wfid = ?")

	adapter.SimpleUpdate("bumpSync", "sync", "last_update = UTC_TIMESTAMP()", "")

	return nil
}

func writeDeletes(adapter qgen.Adapter) error {
	adapter.SimpleDelete("deleteProfileReply", "users_replies", "rid = ?")

	//adapter.SimpleDelete("deleteForumPermsByForum", "forums_permissions", "fid = ?")

	adapter.SimpleDelete("deleteActivityStreamMatch", "activity_stream_matches", "watcher = ? AND asid = ?")
	//adapter.SimpleDelete("deleteActivityStreamMatchesByWatcher","activity_stream_matches","watcher = ?")

	adapter.SimpleDelete("deleteWordFilter", "word_filters", "wfid = ?")

	return nil
}

func writeSimpleCounts(adapter qgen.Adapter) error {
	adapter.SimpleCount("reportExists", "topics", "data = ? AND data != '' AND parentID = 1", "")

	adapter.SimpleCount("modlogCount", "moderation_logs", "", "")

	return nil
}

func writeInsertSelects(adapter qgen.Adapter) error {
	/*adapter.SimpleInsertSelect("addForumPermsToForumAdmins",
		qgen.DB_Insert{"forums_permissions", "gid, fid, preset, permissions", ""},
		qgen.DB_Select{"users_groups", "gid, ? AS fid, ? AS preset, ? AS permissions", "is_admin = 1", "", ""},
	)*/

	/*adapter.SimpleInsertSelect("addForumPermsToForumStaff",
		qgen.DB_Insert{"forums_permissions", "gid, fid, preset, permissions", ""},
		qgen.DB_Select{"users_groups", "gid, ? AS fid, ? AS preset, ? AS permissions", "is_admin = 0 AND is_mod = 1", "", ""},
	)*/

	/*adapter.SimpleInsertSelect("addForumPermsToForumMembers",
		qgen.DB_Insert{"forums_permissions", "gid, fid, preset, permissions", ""},
		qgen.DB_Select{"users_groups", "gid, ? AS fid, ? AS preset, ? AS permissions", "is_admin = 0 AND is_mod = 0 AND is_banned = 0", "", ""},
	)*/

	return nil
}

// nolint
func writeInsertLeftJoins(adapter qgen.Adapter) error {
	return nil
}

func writeInsertInnerJoins(adapter qgen.Adapter) error {
	adapter.SimpleInsertInnerJoin("notifyWatchers",
		qgen.DBInsert{"activity_stream_matches", "watcher, asid", ""},
		qgen.DBJoin{"activity_stream", "activity_subscriptions", "activity_subscriptions.user, activity_stream.asid", "activity_subscriptions.targetType = activity_stream.elementType AND activity_subscriptions.targetID = activity_stream.elementID AND activity_subscriptions.user != activity_stream.actor", "asid = ?", "", ""},
	)

	return nil
}

func writeFile(name string, content string) (err error) {
	f, err := os.Create(name)
	if err != nil {
		return err
	}
	_, err = f.WriteString(content)
	if err != nil {
		return err
	}
	err = f.Sync()
	if err != nil {
		return err
	}
	return f.Close()
}
