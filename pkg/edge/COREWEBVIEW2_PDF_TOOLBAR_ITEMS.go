//go:build windows

package edge

type COREWEBVIEW2_PDF_TOOLBAR_ITEMS uint32

const (
	COREWEBVIEW2_PDF_TOOLBAR_ITEMS_NONE          = 0x0
	COREWEBVIEW2_PDF_TOOLBAR_ITEMS_SAVE          = 0x0001
	COREWEBVIEW2_PDF_TOOLBAR_ITEMS_PRINT         = 0x0002
	COREWEBVIEW2_PDF_TOOLBAR_ITEMS_SAVE_AS       = 0x0004
	COREWEBVIEW2_PDF_TOOLBAR_ITEMS_ZOOM_IN       = 0x0008
	COREWEBVIEW2_PDF_TOOLBAR_ITEMS_ZOOM_OUT      = 0x0010
	COREWEBVIEW2_PDF_TOOLBAR_ITEMS_ROTATE        = 0x0020
	COREWEBVIEW2_PDF_TOOLBAR_ITEMS_FIT_PAGE      = 0x0040
	COREWEBVIEW2_PDF_TOOLBAR_ITEMS_PAGE_LAYOUT   = 0x0080
	COREWEBVIEW2_PDF_TOOLBAR_ITEMS_BOOKMARKS     = 0x0100
	COREWEBVIEW2_PDF_TOOLBAR_ITEMS_PAGE_SELECTOR = 0x0200
	COREWEBVIEW2_PDF_TOOLBAR_ITEMS_SEARCH        = 0x0400
	COREWEBVIEW2_PDF_TOOLBAR_ITEMS_FULL_SCREEN   = 0x0800
	COREWEBVIEW2_PDF_TOOLBAR_ITEMS_MORE_SETTINGS = 0x1000
)
