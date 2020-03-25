package tcb

import (
	"github.com/machao520/wechat"
)

//Tcb Tencent Cloud Base
type Tcb struct{
	*wechat.Wechat
}

//NewTcb new Tencent Cloud Base
func NewTcb(wechat *wechat.Wechat)*Tcb{
	return &Tcb{
		wechat,
	}
}

