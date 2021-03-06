package resultcheckforward

import (
	"github.com/PharbersDeveloper/max-go/phmodel/config"
	"github.com/PharbersDeveloper/max-go/phmodel/resultcheck"
	"github.com/alfredyang1986/blackmirror/bmerror"
	"github.com/alfredyang1986/blackmirror/bmpipe"
	"github.com/alfredyang1986/blackmirror/bmrouter"
	"github.com/alfredyang1986/blackmirror/jsonapi"
	"io"
	"net/http"
)

type PHResultCheckForwardBrick struct {
	bk *bmpipe.BMBrick
}

/*------------------------------------------------
 * brick interface
 *------------------------------------------------*/

var maxForward maxconfig.PhForwardConfig

func (b *PHResultCheckForwardBrick) Exec() error {
	return nil
}

func (b *PHResultCheckForwardBrick) Prepare(pr interface{}) error {
	req := pr.(resultcheck.ResultCheck)
	b.BrickInstance().Pr = req
	return nil
}

func (b *PHResultCheckForwardBrick) Done(pkg string, idx int64, e error) error {
	maxForward.GenerateConfig()
	host := maxForward.HostA
	port := maxForward.PortA
	bmrouter.ForWardNextBrick(host, port, pkg, idx, b)
	return nil
}

func (b *PHResultCheckForwardBrick) BrickInstance() *bmpipe.BMBrick {
	if b.bk == nil {
		b.bk = &bmpipe.BMBrick{}
	}
	return b.bk
}

func (b *PHResultCheckForwardBrick) ResultTo(w io.Writer) error {
	pr := b.BrickInstance().Pr
	tmp := pr.(resultcheck.ResultCheck)
	err := jsonapi.ToJsonAPI(&tmp, w)
	return err
}

func (b *PHResultCheckForwardBrick) Return(w http.ResponseWriter) {
	ec := b.BrickInstance().Err
	if ec != 0 {
		bmerror.ErrInstance().ErrorReval(ec, w)
	} else {
		var reval resultcheck.ResultCheck = b.BrickInstance().Pr.(resultcheck.ResultCheck)
		jsonapi.ToJsonAPI(&reval, w)
	}
}
