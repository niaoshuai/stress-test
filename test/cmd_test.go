package test

import (
	"bytes"
	"os/exec"
	"testing"

	logger "stress-test/pkg/log"
)

func TestHeyCmdExecute(t *testing.T) {
	svnLookLogCmd := exec.Command("hey", "-m", "POST", "-H", "ontent-Type:application/x-www-form-urlencoded; charset=UTF-8", "-d", "params=%7B%22pageIndex%22%3A0%2C%22pageSize%22%3A10%2C%22pageType%22%3A%22P%22%2C%22pageId%22%3A%22%22%2C%22pageName%22%3A%22%22%2C%22pageCreator%22%3A%22%22%2C%22pageEseeid%22%3A%22%22%2C%22pageChanneltype%22%3A%22%22%2C%22channelId%22%3A%22%22%2C%22pageWay%22%3A%22%22%2C%22projectId%22%3A%22%22%2C%22pagePosturl%22%3A%22%22%2C%22pageStatus%22%3A%22Y%22%2C%22pageCreateddate_start%22%3A%22%22%2C%22pageCreateddate_end%22%3A%22%22%2C%22menuCode%22%3A%22es%2Fpage%2Flist%22%7D&TOKEN=83746A1D63764609B1B8AFAA1BFC25B2", "-z", perApiTime, "http://umsauat.niceloo.com/api/es/page/list")
	var out1 bytes.Buffer
	svnLookLogCmd.Stdout = &out1
	err := svnLookLogCmd.Start()
	if err != nil {
		logger.Fatal(err)
	}
	err = svnLookLogCmd.Wait()
	if err != nil {
		logger.Fatal(err)
	}
	logger.Info(out1.String())
}
