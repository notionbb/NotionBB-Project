// +build mssql

// This file was generated by Gosora's Query Generator. Please try to avoid modifying this file, as it might change at any time.
package main

import "log"
import "database/sql"
import "./common"

// nolint
type Stmts struct {
	isPluginActive *sql.Stmt
	isThemeDefault *sql.Stmt
	forumEntryExists *sql.Stmt
	groupEntryExists *sql.Stmt
	getForumTopics *sql.Stmt
	addForumPermsToForum *sql.Stmt
	addPlugin *sql.Stmt
	addTheme *sql.Stmt
	createWordFilter *sql.Stmt
	updatePlugin *sql.Stmt
	updatePluginInstall *sql.Stmt
	updateTheme *sql.Stmt
	updateUser *sql.Stmt
	updateGroupPerms *sql.Stmt
	updateGroup *sql.Stmt
	updateEmail *sql.Stmt
	setTempGroup *sql.Stmt
	updateWordFilter *sql.Stmt
	bumpSync *sql.Stmt
	deleteActivityStreamMatch *sql.Stmt
	deleteWordFilter *sql.Stmt

	getActivityFeedByWatcher *sql.Stmt
	getActivityCountByWatcher *sql.Stmt
	todaysPostCount *sql.Stmt
	todaysTopicCount *sql.Stmt
	todaysReportCount *sql.Stmt
	todaysNewUserCount *sql.Stmt

	Mocks bool
}

// nolint
func _gen_mssql() (err error) {
	common.DebugLog("Building the generated statements")
	
	common.DebugLog("Preparing isPluginActive statement.")
	stmts.isPluginActive, err = db.Prepare("SELECT [active] FROM [plugins] WHERE [uname] = ?1")
	if err != nil {
		log.Print("Error in isPluginActive statement.")
		log.Print("Bad Query: ","SELECT [active] FROM [plugins] WHERE [uname] = ?1")
		return err
	}
		
	common.DebugLog("Preparing isThemeDefault statement.")
	stmts.isThemeDefault, err = db.Prepare("SELECT [default] FROM [themes] WHERE [uname] = ?1")
	if err != nil {
		log.Print("Error in isThemeDefault statement.")
		log.Print("Bad Query: ","SELECT [default] FROM [themes] WHERE [uname] = ?1")
		return err
	}
		
	common.DebugLog("Preparing forumEntryExists statement.")
	stmts.forumEntryExists, err = db.Prepare("SELECT [fid] FROM [forums] WHERE [name] = '' ORDER BY fid ASC OFFSET 0 ROWS FETCH NEXT 1 ROWS ONLY")
	if err != nil {
		log.Print("Error in forumEntryExists statement.")
		log.Print("Bad Query: ","SELECT [fid] FROM [forums] WHERE [name] = '' ORDER BY fid ASC OFFSET 0 ROWS FETCH NEXT 1 ROWS ONLY")
		return err
	}
		
	common.DebugLog("Preparing groupEntryExists statement.")
	stmts.groupEntryExists, err = db.Prepare("SELECT [gid] FROM [users_groups] WHERE [name] = '' ORDER BY gid ASC OFFSET 0 ROWS FETCH NEXT 1 ROWS ONLY")
	if err != nil {
		log.Print("Error in groupEntryExists statement.")
		log.Print("Bad Query: ","SELECT [gid] FROM [users_groups] WHERE [name] = '' ORDER BY gid ASC OFFSET 0 ROWS FETCH NEXT 1 ROWS ONLY")
		return err
	}
		
	common.DebugLog("Preparing getForumTopics statement.")
	stmts.getForumTopics, err = db.Prepare("SELECT [topics].[tid],[topics].[title],[topics].[content],[topics].[createdBy],[topics].[is_closed],[topics].[sticky],[topics].[createdAt],[topics].[lastReplyAt],[topics].[parentID],[users].[name],[users].[avatar] FROM [topics] LEFT JOIN [users] ON [topics].[createdBy] = [users].[uid]  WHERE [topics].[parentID] = ?1 ORDER BY topics.sticky DESC,topics.lastReplyAt DESC,topics.createdBy DESC")
	if err != nil {
		log.Print("Error in getForumTopics statement.")
		log.Print("Bad Query: ","SELECT [topics].[tid],[topics].[title],[topics].[content],[topics].[createdBy],[topics].[is_closed],[topics].[sticky],[topics].[createdAt],[topics].[lastReplyAt],[topics].[parentID],[users].[name],[users].[avatar] FROM [topics] LEFT JOIN [users] ON [topics].[createdBy] = [users].[uid]  WHERE [topics].[parentID] = ?1 ORDER BY topics.sticky DESC,topics.lastReplyAt DESC,topics.createdBy DESC")
		return err
	}
		
	common.DebugLog("Preparing addForumPermsToForum statement.")
	stmts.addForumPermsToForum, err = db.Prepare("INSERT INTO [forums_permissions] ([gid],[fid],[preset],[permissions]) VALUES (?,?,?,?)")
	if err != nil {
		log.Print("Error in addForumPermsToForum statement.")
		log.Print("Bad Query: ","INSERT INTO [forums_permissions] ([gid],[fid],[preset],[permissions]) VALUES (?,?,?,?)")
		return err
	}
		
	common.DebugLog("Preparing addPlugin statement.")
	stmts.addPlugin, err = db.Prepare("INSERT INTO [plugins] ([uname],[active],[installed]) VALUES (?,?,?)")
	if err != nil {
		log.Print("Error in addPlugin statement.")
		log.Print("Bad Query: ","INSERT INTO [plugins] ([uname],[active],[installed]) VALUES (?,?,?)")
		return err
	}
		
	common.DebugLog("Preparing addTheme statement.")
	stmts.addTheme, err = db.Prepare("INSERT INTO [themes] ([uname],[default]) VALUES (?,?)")
	if err != nil {
		log.Print("Error in addTheme statement.")
		log.Print("Bad Query: ","INSERT INTO [themes] ([uname],[default]) VALUES (?,?)")
		return err
	}
		
	common.DebugLog("Preparing createWordFilter statement.")
	stmts.createWordFilter, err = db.Prepare("INSERT INTO [word_filters] ([find],[replacement]) VALUES (?,?)")
	if err != nil {
		log.Print("Error in createWordFilter statement.")
		log.Print("Bad Query: ","INSERT INTO [word_filters] ([find],[replacement]) VALUES (?,?)")
		return err
	}
		
	common.DebugLog("Preparing updatePlugin statement.")
	stmts.updatePlugin, err = db.Prepare("UPDATE [plugins] SET [active] = ? WHERE [uname] = ?")
	if err != nil {
		log.Print("Error in updatePlugin statement.")
		log.Print("Bad Query: ","UPDATE [plugins] SET [active] = ? WHERE [uname] = ?")
		return err
	}
		
	common.DebugLog("Preparing updatePluginInstall statement.")
	stmts.updatePluginInstall, err = db.Prepare("UPDATE [plugins] SET [installed] = ? WHERE [uname] = ?")
	if err != nil {
		log.Print("Error in updatePluginInstall statement.")
		log.Print("Bad Query: ","UPDATE [plugins] SET [installed] = ? WHERE [uname] = ?")
		return err
	}
		
	common.DebugLog("Preparing updateTheme statement.")
	stmts.updateTheme, err = db.Prepare("UPDATE [themes] SET [default] = ? WHERE [uname] = ?")
	if err != nil {
		log.Print("Error in updateTheme statement.")
		log.Print("Bad Query: ","UPDATE [themes] SET [default] = ? WHERE [uname] = ?")
		return err
	}
		
	common.DebugLog("Preparing updateUser statement.")
	stmts.updateUser, err = db.Prepare("UPDATE [users] SET [name] = ?,[email] = ?,[group] = ? WHERE [uid] = ?")
	if err != nil {
		log.Print("Error in updateUser statement.")
		log.Print("Bad Query: ","UPDATE [users] SET [name] = ?,[email] = ?,[group] = ? WHERE [uid] = ?")
		return err
	}
		
	common.DebugLog("Preparing updateGroupPerms statement.")
	stmts.updateGroupPerms, err = db.Prepare("UPDATE [users_groups] SET [permissions] = ? WHERE [gid] = ?")
	if err != nil {
		log.Print("Error in updateGroupPerms statement.")
		log.Print("Bad Query: ","UPDATE [users_groups] SET [permissions] = ? WHERE [gid] = ?")
		return err
	}
		
	common.DebugLog("Preparing updateGroup statement.")
	stmts.updateGroup, err = db.Prepare("UPDATE [users_groups] SET [name] = ?,[tag] = ? WHERE [gid] = ?")
	if err != nil {
		log.Print("Error in updateGroup statement.")
		log.Print("Bad Query: ","UPDATE [users_groups] SET [name] = ?,[tag] = ? WHERE [gid] = ?")
		return err
	}
		
	common.DebugLog("Preparing updateEmail statement.")
	stmts.updateEmail, err = db.Prepare("UPDATE [emails] SET [email] = ?,[uid] = ?,[validated] = ?,[token] = ? WHERE [email] = ?")
	if err != nil {
		log.Print("Error in updateEmail statement.")
		log.Print("Bad Query: ","UPDATE [emails] SET [email] = ?,[uid] = ?,[validated] = ?,[token] = ? WHERE [email] = ?")
		return err
	}
		
	common.DebugLog("Preparing setTempGroup statement.")
	stmts.setTempGroup, err = db.Prepare("UPDATE [users] SET [temp_group] = ? WHERE [uid] = ?")
	if err != nil {
		log.Print("Error in setTempGroup statement.")
		log.Print("Bad Query: ","UPDATE [users] SET [temp_group] = ? WHERE [uid] = ?")
		return err
	}
		
	common.DebugLog("Preparing updateWordFilter statement.")
	stmts.updateWordFilter, err = db.Prepare("UPDATE [word_filters] SET [find] = ?,[replacement] = ? WHERE [wfid] = ?")
	if err != nil {
		log.Print("Error in updateWordFilter statement.")
		log.Print("Bad Query: ","UPDATE [word_filters] SET [find] = ?,[replacement] = ? WHERE [wfid] = ?")
		return err
	}
		
	common.DebugLog("Preparing bumpSync statement.")
	stmts.bumpSync, err = db.Prepare("UPDATE [sync] SET [last_update] = GETUTCDATE()")
	if err != nil {
		log.Print("Error in bumpSync statement.")
		log.Print("Bad Query: ","UPDATE [sync] SET [last_update] = GETUTCDATE()")
		return err
	}
		
	common.DebugLog("Preparing deleteActivityStreamMatch statement.")
	stmts.deleteActivityStreamMatch, err = db.Prepare("DELETE FROM [activity_stream_matches] WHERE [watcher] = ? AND [asid] = ?")
	if err != nil {
		log.Print("Error in deleteActivityStreamMatch statement.")
		log.Print("Bad Query: ","DELETE FROM [activity_stream_matches] WHERE [watcher] = ? AND [asid] = ?")
		return err
	}
		
	common.DebugLog("Preparing deleteWordFilter statement.")
	stmts.deleteWordFilter, err = db.Prepare("DELETE FROM [word_filters] WHERE [wfid] = ?")
	if err != nil {
		log.Print("Error in deleteWordFilter statement.")
		log.Print("Bad Query: ","DELETE FROM [word_filters] WHERE [wfid] = ?")
		return err
	}
	
	return nil
}
