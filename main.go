package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"myProject/videoWater/account"
	"myProject/videoWater/common"
	"myProject/videoWater/deal"
	"net/http"
	"strings"
	"time"
)
var conFile = flag.String("f", "", "config file")
var videoPath = flag.String("v", "", "config file")
var line = "********************************************************************************"
func main() {
	Run()
}

func Run()  {

	flag.Parse()
	con := common.ReadConfig(*conFile)
	if con == nil {
		log.Println("配置文件有误")
		time.Sleep(time.Second * 5)
		return
	}

	if len(con.AppId) == 0 {
		fmt.Println("    请购买密钥   ")
		time.Sleep(time.Second * 500)
		return
	}

	acc := account.GetAccount(con.AppId)

	fmt.Println(line)
	fmt.Println(line)
	fmt.Println(line)
	fmt.Println()
	fmt.Println()
	fmt.Println(formatline(fmt.Sprintf("伪原创视频批量剪辑器 %v", common.Version)))
	fmt.Println()
	if acc.AccType < 0 {
		fmt.Println(formatline("账户密钥:"+acc.AppId))
		fmt.Println(formatline("密钥无效，请购买密钥"))
		fmt.Println(formatline("密钥购买地址："+"https://www.kuaifaka.com/purchasing?link=3ZUpQ"))
	} else {
		fmt.Printf("账户 密钥：%v \n",acc.AppId)
		fmt.Printf("账户类型：%v \n",acc.TYPE())
		fmt.Printf("%v \n",acc.Time)
		fmt.Println(acc.Msg)
	}

	fmt.Println()
	fmt.Println()
	fmt.Println(line)
	fmt.Println(line)
	fmt.Println(line)

	if acc.AccType < 0 {
		time.Sleep(time.Second * 500)
		return
	}
	time.Sleep(time.Second * 5)
	if len(*videoPath) > 0 {
		con.VideoPath = *videoPath
		fmt.Println(*videoPath)
	}

	deal.DoFactory(con)
	time.Sleep(time.Second * 5)

}

func formatline(text string)string  {

	r := strings.Repeat(" ", 10)
	return r + text + r

}

type Data []struct {
	URL           string `json:"url"`
	RepositoryURL string `json:"repository_url"`
	LabelsURL     string `json:"labels_url"`
	CommentsURL   string `json:"comments_url"`
	EventsURL     string `json:"events_url"`
	HTMLURL       string `json:"html_url"`
	ID            int    `json:"id"`
	NodeID        string `json:"node_id"`
	Number        int    `json:"number"`
	Title         string `json:"title"`
	User          struct {
		Login             string `json:"login"`
		ID                int    `json:"id"`
		NodeID            string `json:"node_id"`
		AvatarURL         string `json:"avatar_url"`
		GravatarID        string `json:"gravatar_id"`
		URL               string `json:"url"`
		HTMLURL           string `json:"html_url"`
		FollowersURL      string `json:"followers_url"`
		FollowingURL      string `json:"following_url"`
		GistsURL          string `json:"gists_url"`
		StarredURL        string `json:"starred_url"`
		SubscriptionsURL  string `json:"subscriptions_url"`
		OrganizationsURL  string `json:"organizations_url"`
		ReposURL          string `json:"repos_url"`
		EventsURL         string `json:"events_url"`
		ReceivedEventsURL string `json:"received_events_url"`
		Type              string `json:"type"`
		SiteAdmin         bool   `json:"site_admin"`
	} `json:"user"`
	Labels            []interface{} `json:"labels"`
	State             string        `json:"state"`
	Locked            bool          `json:"locked"`
	Assignee          interface{}   `json:"assignee"`
	Assignees         []interface{} `json:"assignees"`
	Milestone         interface{}   `json:"milestone"`
	Comments          int           `json:"comments"`
	CreatedAt         time.Time     `json:"created_at"`
	UpdatedAt         time.Time     `json:"updated_at"`
	ClosedAt          interface{}   `json:"closed_at"`
	AuthorAssociation string        `json:"author_association"`
	Body              string        `json:"body"`
}

type Message struct {
	Code int	`json:"code"`
	Msg string  `json:"msg"`
}

func check() bool  {

	url := "https://api.github.com/repos/suifengqjn/videoWater/issues"
	client := http.Client{Timeout:time.Second*20}
	resp, err := client.Get(url)
	if err != nil {
		fmt.Println("请检查网络")
		return false
	}
	defer resp.Body.Close()

	buf, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("请检查网络")
		return false
	}

	var res Data
	err = json.Unmarshal(buf, &res)
	if err != nil {
		fmt.Println("请检查网络")
		return false
	}
	var msg Message
	for _, d := range res {
		if d.Title == "2.0" {
			err = json.Unmarshal([]byte(d.Body), &msg)
			break
		}
	}


	if len(msg.Msg) > 0 {
		fmt.Println("===========================")
		fmt.Println(msg.Msg)
		fmt.Println("===========================")
	}

	if msg.Code == 1 {
		return true
	}

	return false
}

