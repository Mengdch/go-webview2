//go:build windows

package edge

type COREWEBVIEW2_CONTEXT_MENU_TARGET_KIND uint32

const (
	COREWEBVIEW2_CONTEXT_MENU_TARGET_KIND_PAGE          = 0 //含超链接处右键 && 页面内其他区域右键 && 页面内输入框右键
	COREWEBVIEW2_CONTEXT_MENU_TARGET_KIND_IMAGE         = 1 //页面内图片处右键
	COREWEBVIEW2_CONTEXT_MENU_TARGET_KIND_SELECTED_TEXT = 2 //页面内文字选中右键 && 页面内链接选中右键
	COREWEBVIEW2_CONTEXT_MENU_TARGET_KIND_AUDIO         = 3
	COREWEBVIEW2_CONTEXT_MENU_TARGET_KIND_VIDEO         = 4
)
const (
	MENU_SAVEIMAGEAS         = "saveImageAs"         //图像另存为
	MENU_MAGNIFYIMAGE        = "magnifyImage"        //放大图像
	MENU_SHARE               = "share"               //共享
	MENU_COPYLINKTOHIGHLIGHT = "copyLinkToHighlight" //复制指向突出显示的链接
	MENU_SAVEAS              = "saveAs"              //另存为
	MENU_WEBCAPTURE          = "webCapture"          //截屏
	MENU_OTHER               = "other"               //分隔线
	MENU_RELOAD              = "reload"              //刷新
	MENU_BACK                = "back"                //返回
	MENU_FORWARD             = "forward"             //前进
	MENU_PRINT               = "print"               //打印
)
