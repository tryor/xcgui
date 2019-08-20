package main

import (
	"fmt"
)

import (
	"log"

	xcgui "github.com/tryor/xcgui/xc"
)

var template = `<?xml version="1.0" encoding="utf8" ?>
<xc_template_comboBox>
	<layoutObject type="item" layout.horizon="true" layout.alignH="left"  layout.alignV="top" layout.width="auto" layout.height="fill" layout.space="2" layout.padding="2,2,2,2" >
		<shapeText itemID="1" layout.width="auto" layout.height="fill" content="{binding field=name source=data bSync=false}" />
	</layoutObject>
</xc_template_comboBox>`

func main() {
	hWindow := xcgui.XWnd_Create(0, 0, 300, 200, "炫彩界面库窗口", 0, xcgui.XC_WINDOW_STYLE_DEFAULT)
	xcgui.CloseBtn(hWindow)

	hComboBox := xcgui.XComboBox_Create(20, 40, 120, 28, xcgui.HXCGUI(hWindow))
	log.Println("hComboBox:", hComboBox)
	xcgui.XComboBox_SetItemTemplateXML(hComboBox, "../xml-template/ComboBox_ListBox_Item.xml")
	//xcgui.XComboBox_SetItemTemplateXMLFromString(hComboBox, template)

	//xcgui.XRichEdit_SetText(hComboBox, "123")
	xcgui.XComboBox_EnableEdit(hComboBox, false)

	//	xcgui.XComboBox_AddItemText(hComboBox, "111111")
	//xcgui.XComboBox_AddItemText(hComboBox, "222222")

	//xcgui.XComboBox_InsertItemText(hComboBox, 0, "222222")

	hAdapter := xcgui.XAdTable_Create()
	xcgui.XComboBox_BindAdapter(hComboBox, hAdapter)
	xcgui.XComboBox_SetBindName(hComboBox, "name")

	xcgui.XAdTable_AddColumn(hAdapter, "name")
	for i := 0; i < 20; i++ {
		xcgui.XAdTable_AddItemText(hAdapter, "name-"+fmt.Sprint(i)+"-0")
		//		xcgui.XAdTable_AddItemTextEx(hAdapter, "name", "name-"+fmt.Sprint(i)+"-0")
	}

	//	hAdapter := xcgui.XAdListView_Create()
	//	xcgui.XComboBox_BindAdapter(hComboBox, hAdapter)
	//	xcgui.XAdListView_GroupAddColumn(hAdapter, "name")

	//	for i := 0; i < 20; i++ {
	//		xcgui.XAdListView_GroupAddItemText(hAdapter, "name-"+fmt.Sprint(i)+"-0")
	//	}

	xcgui.XWnd_ShowWindow(hWindow, xcgui.SW_SHOW)
	xcgui.XRunXCGUI()
	xcgui.XExitXCGUI()
}
