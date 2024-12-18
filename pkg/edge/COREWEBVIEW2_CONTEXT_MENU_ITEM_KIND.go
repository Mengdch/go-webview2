//go:build windows

package edge

type COREWEBVIEW2_CONTEXT_MENU_ITEM_KIND uint32

const (
	COREWEBVIEW2_CONTEXT_MENU_ITEM_KIND_COMMAND   = 0
	COREWEBVIEW2_CONTEXT_MENU_ITEM_KIND_CHECK_BOX = 1
	COREWEBVIEW2_CONTEXT_MENU_ITEM_KIND_RADIO     = 2
	COREWEBVIEW2_CONTEXT_MENU_ITEM_KIND_SEPARATOR = 3
	COREWEBVIEW2_CONTEXT_MENU_ITEM_KIND_SUBMENU   = 4
)
