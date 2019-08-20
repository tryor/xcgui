// 炫彩界面库，用的免费版dll。
//
// 有钱的朋友可以去官方网站购买正式版，支持正版将有更多功能。
//
// 官方网站：http://www.xcgui.com
package xc

/*
enum   edit_type_ { edit_type_none = 0, edit_type_richedit, edit_type_chat}
*/
type Edit_type_ int32

const (
	Edit_type_none = iota + 0
	Edit_type_richedit
	Edit_type_chat
)
