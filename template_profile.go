// +build !no_templategen

// Code generated by Gosora. More below:
/* This file was automatically generated by the software. Please don't edit it as your changes may be overwritten at any moment. */
package main
import "net/http"
import "./common"
import "strconv"

var profile_tmpl_phrase_id int

// nolint
func init() {
	common.Template_profile_handle = Template_profile
	common.Ctemplates = append(common.Ctemplates,"profile")
	common.TmplPtrMap["profile"] = &common.Template_profile_handle
	common.TmplPtrMap["o_profile"] = Template_profile
	profile_tmpl_phrase_id = common.RegisterTmplPhraseNames([]string{
		"menu_forums_aria",
		"menu_forums_tooltip",
		"menu_topics_aria",
		"menu_topics_tooltip",
		"menu_alert_counter_aria",
		"menu_alert_list_aria",
		"menu_account_aria",
		"menu_account_tooltip",
		"menu_profile_aria",
		"menu_profile_tooltip",
		"menu_panel_aria",
		"menu_panel_tooltip",
		"menu_logout_aria",
		"menu_logout_tooltip",
		"menu_register_aria",
		"menu_register_tooltip",
		"menu_login_aria",
		"menu_login_tooltip",
		"menu_hamburger_tooltip",
		"profile_login_for_options",
		"profile_add_friend",
		"profile_unban",
		"profile_ban",
		"profile_report_user_aria",
		"profile_report_user_tooltip",
		"profile_ban_user_head",
		"profile_ban_user_notice",
		"profile_ban_user_days",
		"profile_ban_user_weeks",
		"profile_ban_user_months",
		"profile_ban_user_reason",
		"profile_ban_user_button",
		"profile_comments_head",
		"profile_comments_edit_tooltip",
		"profile_comments_edit_aria",
		"profile_comments_delete_tooltip",
		"profile_comments_delete_aria",
		"profile_comments_report_tooltip",
		"profile_comments_report_aria",
		"profile_comments_edit_tooltip",
		"profile_comments_edit_aria",
		"profile_comments_delete_tooltip",
		"profile_comments_delete_aria",
		"profile_comments_report_tooltip",
		"profile_comments_report_aria",
		"profile_comments_form_content",
		"profile_comments_form_button",
		"footer_powered_by",
		"footer_made_with_love",
		"footer_theme_selector_aria",
	})
}

// nolint
func Template_profile(tmpl_profile_vars common.ProfilePage, w http.ResponseWriter) error {
var phrases = common.GetTmplPhrasesBytes(profile_tmpl_phrase_id)
w.Write(header_frags[0])
w.Write([]byte(tmpl_profile_vars.Title))
w.Write(header_frags[1])
w.Write([]byte(tmpl_profile_vars.Header.Site.Name))
w.Write(header_frags[2])
w.Write([]byte(tmpl_profile_vars.Header.Theme.Name))
w.Write(header_frags[3])
if len(tmpl_profile_vars.Header.Stylesheets) != 0 {
for _, item := range tmpl_profile_vars.Header.Stylesheets {
w.Write(header_frags[4])
w.Write([]byte(item))
w.Write(header_frags[5])
}
}
w.Write(header_frags[6])
if len(tmpl_profile_vars.Header.Scripts) != 0 {
for _, item := range tmpl_profile_vars.Header.Scripts {
w.Write(header_frags[7])
w.Write([]byte(item))
w.Write(header_frags[8])
}
}
w.Write(header_frags[9])
w.Write([]byte(tmpl_profile_vars.CurrentUser.Session))
w.Write(header_frags[10])
w.Write([]byte(tmpl_profile_vars.Header.Site.URL))
w.Write(header_frags[11])
if tmpl_profile_vars.Header.MetaDesc != "" {
w.Write(header_frags[12])
w.Write([]byte(tmpl_profile_vars.Header.MetaDesc))
w.Write(header_frags[13])
}
w.Write(header_frags[14])
if !tmpl_profile_vars.CurrentUser.IsSuperMod {
w.Write(header_frags[15])
}
w.Write(header_frags[16])
w.Write(menu_frags[0])
w.Write([]byte(common.BuildWidget("leftOfNav",tmpl_profile_vars.Header)))
w.Write(menu_frags[1])
w.Write(menu_frags[2])
w.Write([]byte(tmpl_profile_vars.Header.Site.ShortName))
w.Write(menu_frags[3])
w.Write(phrases[0])
w.Write(menu_frags[4])
w.Write(phrases[1])
w.Write(menu_frags[5])
w.Write(phrases[2])
w.Write(menu_frags[6])
w.Write(phrases[3])
w.Write(menu_frags[7])
w.Write(phrases[4])
w.Write(menu_frags[8])
w.Write(phrases[5])
w.Write(menu_frags[9])
if tmpl_profile_vars.CurrentUser.Loggedin {
w.Write(menu_frags[10])
w.Write(phrases[6])
w.Write(menu_frags[11])
w.Write(phrases[7])
w.Write(menu_frags[12])
w.Write([]byte(tmpl_profile_vars.CurrentUser.Link))
w.Write(menu_frags[13])
w.Write(phrases[8])
w.Write(menu_frags[14])
w.Write(phrases[9])
w.Write(menu_frags[15])
w.Write(phrases[10])
w.Write(menu_frags[16])
w.Write(phrases[11])
w.Write(menu_frags[17])
w.Write([]byte(tmpl_profile_vars.CurrentUser.Session))
w.Write(menu_frags[18])
w.Write(phrases[12])
w.Write(menu_frags[19])
w.Write(phrases[13])
w.Write(menu_frags[20])
} else {
w.Write(menu_frags[21])
w.Write(phrases[14])
w.Write(menu_frags[22])
w.Write(phrases[15])
w.Write(menu_frags[23])
w.Write(phrases[16])
w.Write(menu_frags[24])
w.Write(phrases[17])
w.Write(menu_frags[25])
}
w.Write(menu_frags[26])
w.Write(phrases[18])
w.Write(menu_frags[27])
w.Write([]byte(common.BuildWidget("rightOfNav",tmpl_profile_vars.Header)))
w.Write(menu_frags[28])
w.Write(header_frags[17])
if tmpl_profile_vars.Header.Widgets.RightSidebar != "" {
w.Write(header_frags[18])
}
w.Write(header_frags[19])
if len(tmpl_profile_vars.Header.NoticeList) != 0 {
for _, item := range tmpl_profile_vars.Header.NoticeList {
w.Write(header_frags[20])
w.Write([]byte(item))
w.Write(header_frags[21])
}
}
w.Write(header_frags[22])
w.Write(profile_frags[0])
w.Write([]byte(tmpl_profile_vars.ProfileOwner.Avatar))
w.Write(profile_frags[1])
w.Write([]byte(tmpl_profile_vars.ProfileOwner.Name))
w.Write(profile_frags[2])
w.Write([]byte(tmpl_profile_vars.ProfileOwner.Name))
w.Write(profile_frags[3])
w.Write([]byte(tmpl_profile_vars.ProfileOwner.Name))
w.Write(profile_frags[4])
w.Write([]byte(tmpl_profile_vars.ProfileOwner.Name))
w.Write(profile_frags[5])
if tmpl_profile_vars.ProfileOwner.Tag != "" {
w.Write(profile_frags[6])
w.Write([]byte(tmpl_profile_vars.ProfileOwner.Tag))
w.Write(profile_frags[7])
w.Write([]byte(tmpl_profile_vars.ProfileOwner.Tag))
w.Write(profile_frags[8])
}
w.Write(profile_frags[9])
if !tmpl_profile_vars.CurrentUser.Loggedin {
w.Write(profile_frags[10])
w.Write(phrases[19])
w.Write(profile_frags[11])
} else {
w.Write(profile_frags[12])
w.Write(phrases[20])
w.Write(profile_frags[13])
if tmpl_profile_vars.CurrentUser.IsSuperMod && !tmpl_profile_vars.ProfileOwner.IsSuperMod {
w.Write(profile_frags[14])
if tmpl_profile_vars.ProfileOwner.IsBanned {
w.Write(profile_frags[15])
w.Write([]byte(strconv.Itoa(tmpl_profile_vars.ProfileOwner.ID)))
w.Write(profile_frags[16])
w.Write([]byte(tmpl_profile_vars.CurrentUser.Session))
w.Write(profile_frags[17])
w.Write(phrases[21])
w.Write(profile_frags[18])
} else {
w.Write(profile_frags[19])
w.Write(phrases[22])
w.Write(profile_frags[20])
}
w.Write(profile_frags[21])
}
w.Write(profile_frags[22])
w.Write([]byte(strconv.Itoa(tmpl_profile_vars.ProfileOwner.ID)))
w.Write(profile_frags[23])
w.Write([]byte(tmpl_profile_vars.CurrentUser.Session))
w.Write(profile_frags[24])
w.Write(phrases[23])
w.Write(profile_frags[25])
w.Write(phrases[24])
w.Write(profile_frags[26])
}
w.Write(profile_frags[27])
if tmpl_profile_vars.CurrentUser.Perms.BanUsers {
w.Write(profile_frags[28])
w.Write(phrases[25])
w.Write(profile_frags[29])
w.Write([]byte(strconv.Itoa(tmpl_profile_vars.ProfileOwner.ID)))
w.Write(profile_frags[30])
w.Write([]byte(tmpl_profile_vars.CurrentUser.Session))
w.Write(profile_frags[31])
w.Write(profile_frags[32])
w.Write(phrases[26])
w.Write(profile_frags[33])
w.Write(phrases[27])
w.Write(profile_frags[34])
w.Write(phrases[28])
w.Write(profile_frags[35])
w.Write(phrases[29])
w.Write(profile_frags[36])
w.Write(phrases[30])
w.Write(profile_frags[37])
w.Write(phrases[31])
w.Write(profile_frags[38])
}
w.Write(profile_frags[39])
w.Write(phrases[32])
w.Write(profile_frags[40])
if tmpl_profile_vars.Header.Theme.BgAvatars {
if len(tmpl_profile_vars.ItemList) != 0 {
for _, item := range tmpl_profile_vars.ItemList {
w.Write(profile_comments_row_frags[0])
w.Write([]byte(item.ClassName))
w.Write(profile_comments_row_frags[1])
w.Write([]byte(item.Avatar))
w.Write(profile_comments_row_frags[2])
if item.ContentLines <= 5 {
w.Write(profile_comments_row_frags[3])
}
w.Write(profile_comments_row_frags[4])
w.Write([]byte(item.ContentHtml))
w.Write(profile_comments_row_frags[5])
w.Write([]byte(item.UserLink))
w.Write(profile_comments_row_frags[6])
w.Write([]byte(item.CreatedByName))
w.Write(profile_comments_row_frags[7])
if tmpl_profile_vars.CurrentUser.IsMod {
w.Write(profile_comments_row_frags[8])
w.Write([]byte(strconv.Itoa(item.ID)))
w.Write(profile_comments_row_frags[9])
w.Write([]byte(tmpl_profile_vars.CurrentUser.Session))
w.Write(profile_comments_row_frags[10])
w.Write(phrases[33])
w.Write(profile_comments_row_frags[11])
w.Write(phrases[34])
w.Write(profile_comments_row_frags[12])
w.Write([]byte(strconv.Itoa(item.ID)))
w.Write(profile_comments_row_frags[13])
w.Write([]byte(tmpl_profile_vars.CurrentUser.Session))
w.Write(profile_comments_row_frags[14])
w.Write(phrases[35])
w.Write(profile_comments_row_frags[15])
w.Write(phrases[36])
w.Write(profile_comments_row_frags[16])
}
w.Write(profile_comments_row_frags[17])
w.Write([]byte(strconv.Itoa(item.ID)))
w.Write(profile_comments_row_frags[18])
w.Write([]byte(tmpl_profile_vars.CurrentUser.Session))
w.Write(profile_comments_row_frags[19])
w.Write(phrases[37])
w.Write(profile_comments_row_frags[20])
w.Write(phrases[38])
w.Write(profile_comments_row_frags[21])
if item.Tag != "" {
w.Write(profile_comments_row_frags[22])
w.Write([]byte(item.Tag))
w.Write(profile_comments_row_frags[23])
}
w.Write(profile_comments_row_frags[24])
}
}
} else {
if len(tmpl_profile_vars.ItemList) != 0 {
for _, item := range tmpl_profile_vars.ItemList {
w.Write(profile_comments_row_frags[25])
w.Write([]byte(item.ClassName))
w.Write(profile_comments_row_frags[26])
w.Write([]byte(item.Avatar))
w.Write(profile_comments_row_frags[27])
w.Write([]byte(item.CreatedByName))
w.Write(profile_comments_row_frags[28])
w.Write([]byte(item.CreatedByName))
w.Write(profile_comments_row_frags[29])
w.Write([]byte(item.UserLink))
w.Write(profile_comments_row_frags[30])
w.Write([]byte(item.CreatedByName))
w.Write(profile_comments_row_frags[31])
if item.Tag != "" {
w.Write(profile_comments_row_frags[32])
w.Write([]byte(item.Tag))
w.Write(profile_comments_row_frags[33])
}
w.Write(profile_comments_row_frags[34])
if tmpl_profile_vars.CurrentUser.IsMod {
w.Write(profile_comments_row_frags[35])
w.Write([]byte(strconv.Itoa(item.ID)))
w.Write(profile_comments_row_frags[36])
w.Write([]byte(tmpl_profile_vars.CurrentUser.Session))
w.Write(profile_comments_row_frags[37])
w.Write(phrases[39])
w.Write(profile_comments_row_frags[38])
w.Write(phrases[40])
w.Write(profile_comments_row_frags[39])
w.Write([]byte(strconv.Itoa(item.ID)))
w.Write(profile_comments_row_frags[40])
w.Write([]byte(tmpl_profile_vars.CurrentUser.Session))
w.Write(profile_comments_row_frags[41])
w.Write(phrases[41])
w.Write(profile_comments_row_frags[42])
w.Write(phrases[42])
w.Write(profile_comments_row_frags[43])
}
w.Write(profile_comments_row_frags[44])
w.Write([]byte(strconv.Itoa(item.ID)))
w.Write(profile_comments_row_frags[45])
w.Write([]byte(tmpl_profile_vars.CurrentUser.Session))
w.Write(profile_comments_row_frags[46])
w.Write(phrases[43])
w.Write(profile_comments_row_frags[47])
w.Write(phrases[44])
w.Write(profile_comments_row_frags[48])
w.Write([]byte(item.ContentHtml))
w.Write(profile_comments_row_frags[49])
}
}
}
w.Write(profile_frags[41])
if !tmpl_profile_vars.CurrentUser.IsBanned {
w.Write(profile_frags[42])
w.Write([]byte(tmpl_profile_vars.CurrentUser.Session))
w.Write(profile_frags[43])
w.Write([]byte(strconv.Itoa(tmpl_profile_vars.ProfileOwner.ID)))
w.Write(profile_frags[44])
w.Write(phrases[45])
w.Write(profile_frags[45])
w.Write(phrases[46])
w.Write(profile_frags[46])
}
w.Write(profile_frags[47])
w.Write(profile_frags[48])
w.Write(footer_frags[0])
w.Write([]byte(common.BuildWidget("footer",tmpl_profile_vars.Header)))
w.Write(footer_frags[1])
w.Write(phrases[47])
w.Write(footer_frags[2])
w.Write(phrases[48])
w.Write(footer_frags[3])
w.Write(phrases[49])
w.Write(footer_frags[4])
if len(tmpl_profile_vars.Header.Themes) != 0 {
for _, item := range tmpl_profile_vars.Header.Themes {
if !item.HideFromThemes {
w.Write(footer_frags[5])
w.Write([]byte(item.Name))
w.Write(footer_frags[6])
if tmpl_profile_vars.Header.Theme.Name == item.Name {
w.Write(footer_frags[7])
}
w.Write(footer_frags[8])
w.Write([]byte(item.FriendlyName))
w.Write(footer_frags[9])
}
}
}
w.Write(footer_frags[10])
w.Write([]byte(common.BuildWidget("rightSidebar",tmpl_profile_vars.Header)))
w.Write(footer_frags[11])
	return nil
}
