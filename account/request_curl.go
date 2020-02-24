package account

import (
	"fmt"
	"myTool/common"
	"os/exec"
	"time"
)

func Curl(url string) {

	time.Now()
	key := common.MD5String(fmt.Sprintf("%v", time.Now().UTC().UnixNano()))
	headKey := fmt.Sprintf("%v:%v","X-API-KEY", key)
	cmd := exec.Command("./curl.exe",
		"-H","Content-Type:application/json",
		"-H",headKey,
		"-X","POST",
		url,

	)

	fmt.Println(url)
	output, err := cmd.CombinedOutput()

	fmt.Println("------")
	fmt.Println(string(output), err)

}
