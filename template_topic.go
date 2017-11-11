// +build !no_templategen

// Code generated by Gosora. More below:
/* This file was automatically generated by the software. Please don't edit it as your changes may be overwritten at any moment. */
package main
import "net/http"
import "./common"
import "strconv"

// nolint
func init() {
	common.Template_topic_handle = Template_topic
	common.Ctemplates = append(common.Ctemplates,"topic")
	common.TmplPtrMap["topic"] = &common.Template_topic_handle
	common.TmplPtrMap["o_topic"] = Template_topic
}

// nolint
func Template_topic(tmpl_topic_vars common.TopicPage, w http.ResponseWriter) error {
w.Write(header_0)
w.Write([]byte(tmpl_topic_vars.Title))
w.Write(header_1)
w.Write([]byte(tmpl_topic_vars.Header.Site.Name))
w.Write(header_2)
w.Write([]byte(tmpl_topic_vars.Header.ThemeName))
w.Write(header_3)
if len(tmpl_topic_vars.Header.Stylesheets) != 0 {
for _, item := range tmpl_topic_vars.Header.Stylesheets {
w.Write(header_4)
w.Write([]byte(item))
w.Write(header_5)
}
}
w.Write(header_6)
if len(tmpl_topic_vars.Header.Scripts) != 0 {
for _, item := range tmpl_topic_vars.Header.Scripts {
w.Write(header_7)
w.Write([]byte(item))
w.Write(header_8)
}
}
w.Write(header_9)
w.Write([]byte(tmpl_topic_vars.CurrentUser.Session))
w.Write(header_10)
w.Write([]byte(tmpl_topic_vars.Header.Site.URL))
w.Write(header_11)
if !tmpl_topic_vars.CurrentUser.IsSuperMod {
w.Write(header_12)
}
w.Write(header_13)
w.Write(menu_0)
w.Write(menu_1)
w.Write([]byte(tmpl_topic_vars.Header.Site.ShortName))
w.Write(menu_2)
if tmpl_topic_vars.CurrentUser.Loggedin {
w.Write(menu_3)
w.Write([]byte(tmpl_topic_vars.CurrentUser.Link))
w.Write(menu_4)
w.Write([]byte(tmpl_topic_vars.CurrentUser.Session))
w.Write(menu_5)
} else {
w.Write(menu_6)
}
w.Write(menu_7)
w.Write(header_14)
if tmpl_topic_vars.Header.Widgets.RightSidebar != "" {
w.Write(header_15)
}
w.Write(header_16)
if len(tmpl_topic_vars.Header.NoticeList) != 0 {
for _, item := range tmpl_topic_vars.Header.NoticeList {
w.Write(header_17)
w.Write([]byte(item))
w.Write(header_18)
}
}
w.Write(topic_0)
w.Write([]byte(strconv.Itoa(tmpl_topic_vars.Topic.ID)))
w.Write(topic_1)
if tmpl_topic_vars.Page > 1 {
w.Write(topic_2)
w.Write([]byte(strconv.Itoa(tmpl_topic_vars.Topic.ID)))
w.Write(topic_3)
w.Write([]byte(strconv.Itoa(tmpl_topic_vars.Page - 1)))
w.Write(topic_4)
w.Write([]byte(strconv.Itoa(tmpl_topic_vars.Topic.ID)))
w.Write(topic_5)
w.Write([]byte(strconv.Itoa(tmpl_topic_vars.Page - 1)))
w.Write(topic_6)
}
if tmpl_topic_vars.LastPage != tmpl_topic_vars.Page {
w.Write(topic_7)
w.Write([]byte(strconv.Itoa(tmpl_topic_vars.Topic.ID)))
w.Write(topic_8)
w.Write([]byte(strconv.Itoa(tmpl_topic_vars.Page + 1)))
w.Write(topic_9)
w.Write([]byte(strconv.Itoa(tmpl_topic_vars.Topic.ID)))
w.Write(topic_10)
w.Write([]byte(strconv.Itoa(tmpl_topic_vars.Page + 1)))
w.Write(topic_11)
}
w.Write(topic_12)
if tmpl_topic_vars.Topic.Sticky {
w.Write(topic_13)
} else {
if tmpl_topic_vars.Topic.IsClosed {
w.Write(topic_14)
}
}
w.Write(topic_15)
w.Write([]byte(tmpl_topic_vars.Topic.Title))
w.Write(topic_16)
if tmpl_topic_vars.Topic.IsClosed {
w.Write(topic_17)
}
if tmpl_topic_vars.CurrentUser.Perms.EditTopic {
w.Write(topic_18)
w.Write([]byte(tmpl_topic_vars.Topic.Title))
w.Write(topic_19)
}
w.Write(topic_20)
w.Write([]byte(tmpl_topic_vars.Topic.ClassName))
w.Write(topic_21)
if tmpl_topic_vars.Topic.Avatar != "" {
w.Write(topic_22)
w.Write([]byte(tmpl_topic_vars.Topic.Avatar))
w.Write(topic_23)
w.Write([]byte(tmpl_topic_vars.Header.ThemeName))
w.Write(topic_24)
if tmpl_topic_vars.Topic.ContentLines <= 5 {
w.Write(topic_25)
}
w.Write(topic_26)
}
w.Write(topic_27)
w.Write([]byte(tmpl_topic_vars.Topic.ContentHTML))
w.Write(topic_28)
w.Write([]byte(tmpl_topic_vars.Topic.Content))
w.Write(topic_29)
w.Write([]byte(tmpl_topic_vars.Topic.UserLink))
w.Write(topic_30)
w.Write([]byte(tmpl_topic_vars.Topic.CreatedByName))
w.Write(topic_31)
if tmpl_topic_vars.CurrentUser.Perms.LikeItem {
w.Write(topic_32)
w.Write([]byte(strconv.Itoa(tmpl_topic_vars.Topic.ID)))
w.Write(topic_33)
if tmpl_topic_vars.Topic.Liked {
w.Write(topic_34)
}
w.Write(topic_35)
}
if tmpl_topic_vars.CurrentUser.Perms.EditTopic {
w.Write(topic_36)
w.Write([]byte(strconv.Itoa(tmpl_topic_vars.Topic.ID)))
w.Write(topic_37)
}
if tmpl_topic_vars.CurrentUser.Perms.DeleteTopic {
w.Write(topic_38)
w.Write([]byte(strconv.Itoa(tmpl_topic_vars.Topic.ID)))
w.Write(topic_39)
}
if tmpl_topic_vars.CurrentUser.Perms.CloseTopic {
if tmpl_topic_vars.Topic.IsClosed {
w.Write(topic_40)
w.Write([]byte(strconv.Itoa(tmpl_topic_vars.Topic.ID)))
w.Write(topic_41)
} else {
w.Write(topic_42)
w.Write([]byte(strconv.Itoa(tmpl_topic_vars.Topic.ID)))
w.Write(topic_43)
}
}
if tmpl_topic_vars.CurrentUser.Perms.PinTopic {
if tmpl_topic_vars.Topic.Sticky {
w.Write(topic_44)
w.Write([]byte(strconv.Itoa(tmpl_topic_vars.Topic.ID)))
w.Write(topic_45)
} else {
w.Write(topic_46)
w.Write([]byte(strconv.Itoa(tmpl_topic_vars.Topic.ID)))
w.Write(topic_47)
}
}
if tmpl_topic_vars.CurrentUser.Perms.ViewIPs {
w.Write(topic_48)
w.Write([]byte(tmpl_topic_vars.Topic.IPAddress))
w.Write(topic_49)
}
w.Write(topic_50)
w.Write([]byte(strconv.Itoa(tmpl_topic_vars.Topic.ID)))
w.Write(topic_51)
w.Write([]byte(tmpl_topic_vars.CurrentUser.Session))
w.Write(topic_52)
if tmpl_topic_vars.Topic.LikeCount > 0 {
w.Write(topic_53)
w.Write([]byte(strconv.Itoa(tmpl_topic_vars.Topic.LikeCount)))
w.Write(topic_54)
}
if tmpl_topic_vars.Topic.Tag != "" {
w.Write(topic_55)
w.Write([]byte(tmpl_topic_vars.Topic.Tag))
w.Write(topic_56)
} else {
w.Write(topic_57)
w.Write([]byte(strconv.Itoa(tmpl_topic_vars.Topic.Level)))
w.Write(topic_58)
}
w.Write(topic_59)
if len(tmpl_topic_vars.ItemList) != 0 {
for _, item := range tmpl_topic_vars.ItemList {
if item.ActionType != "" {
w.Write(topic_60)
w.Write([]byte(item.ActionIcon))
w.Write(topic_61)
w.Write([]byte(item.ActionType))
w.Write(topic_62)
} else {
w.Write(topic_63)
w.Write([]byte(item.ClassName))
w.Write(topic_64)
if item.Avatar != "" {
w.Write(topic_65)
w.Write([]byte(item.Avatar))
w.Write(topic_66)
w.Write([]byte(tmpl_topic_vars.Header.ThemeName))
w.Write(topic_67)
if item.ContentLines <= 5 {
w.Write(topic_68)
}
w.Write(topic_69)
}
w.Write(topic_70)
w.Write(topic_71)
w.Write([]byte(item.ContentHtml))
w.Write(topic_72)
w.Write([]byte(item.UserLink))
w.Write(topic_73)
w.Write([]byte(item.CreatedByName))
w.Write(topic_74)
if tmpl_topic_vars.CurrentUser.Perms.LikeItem {
w.Write(topic_75)
w.Write([]byte(strconv.Itoa(item.ID)))
w.Write(topic_76)
if item.Liked {
w.Write(topic_77)
}
w.Write(topic_78)
}
if tmpl_topic_vars.CurrentUser.Perms.EditReply {
w.Write(topic_79)
w.Write([]byte(strconv.Itoa(item.ID)))
w.Write(topic_80)
}
if tmpl_topic_vars.CurrentUser.Perms.DeleteReply {
w.Write(topic_81)
w.Write([]byte(strconv.Itoa(item.ID)))
w.Write(topic_82)
}
if tmpl_topic_vars.CurrentUser.Perms.ViewIPs {
w.Write(topic_83)
w.Write([]byte(item.IPAddress))
w.Write(topic_84)
}
w.Write(topic_85)
w.Write([]byte(strconv.Itoa(item.ID)))
w.Write(topic_86)
w.Write([]byte(tmpl_topic_vars.CurrentUser.Session))
w.Write(topic_87)
if item.LikeCount > 0 {
w.Write(topic_88)
w.Write([]byte(strconv.Itoa(item.LikeCount)))
w.Write(topic_89)
}
if item.Tag != "" {
w.Write(topic_90)
w.Write([]byte(item.Tag))
w.Write(topic_91)
} else {
w.Write(topic_92)
w.Write([]byte(strconv.Itoa(item.Level)))
w.Write(topic_93)
}
w.Write(topic_94)
}
}
}
w.Write(topic_95)
if tmpl_topic_vars.CurrentUser.Perms.CreateReply {
w.Write(topic_96)
w.Write([]byte(strconv.Itoa(tmpl_topic_vars.Topic.ID)))
w.Write(topic_97)
if tmpl_topic_vars.CurrentUser.Perms.UploadFiles {
w.Write(topic_98)
}
w.Write(topic_99)
}
w.Write(topic_100)
w.Write(footer_0)
if len(tmpl_topic_vars.Header.Themes) != 0 {
for _, item := range tmpl_topic_vars.Header.Themes {
if !item.HideFromThemes {
w.Write(footer_1)
w.Write([]byte(item.Name))
w.Write(footer_2)
if tmpl_topic_vars.Header.ThemeName == item.Name {
w.Write(footer_3)
}
w.Write(footer_4)
w.Write([]byte(item.FriendlyName))
w.Write(footer_5)
}
}
}
w.Write(footer_6)
if tmpl_topic_vars.Header.Widgets.RightSidebar != "" {
w.Write(footer_7)
w.Write([]byte(string(tmpl_topic_vars.Header.Widgets.RightSidebar)))
w.Write(footer_8)
}
w.Write(footer_9)
	return nil
}
