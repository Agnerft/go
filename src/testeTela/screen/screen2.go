package screen

import (
	"github.com/lxn/walk"
	"testeTela/data"
)

func ScreenSecundaria() {
	walk.MsgBox(nil, "YES", data.GetInfo1(), walk.MsgBoxIconInformation)
}
