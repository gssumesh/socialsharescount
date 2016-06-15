package socialsharescount

import (
	"regexp"
	"net/http"
	"io/ioutil"
	"strconv"
)

var resultsMap = make(map[string]int)
var social = [][]string{
	{"Facebook", "http://api.facebook.com/restserver.php?method=links.getStats&urls="  , "<share_count>([0-9]+)</share_count>",  "<like_count>([0-9]+)</like_count>" , "<comment_count>([0-9]+)</comment_count>" , "<total_count>([0-9]+)</total_count>"},
	{"LinkedIn", "http://www.linkedin.com/countserv/count/share?format=json&url=", "count\":([0-9]+),"},
	{"Stumbleupon", "https://www.stumbleupon.com/services/1.01/badge.getinfo?url=", "views\":([0-9]+),"},
	{"Buffer", "https://api.bufferapp.com/1/links/shares.json?url=", "shares\":([0-9]+)"},
	{"Pintrest", "http://api.pinterest.com/v1/urls/count.json?url=", "count\":([0-9]+)"},
	{"Odnoklassniki", "https://connect.ok.ru/dk?st.cmd=extLike&uid=odklcnt0&ref=", "odklcnt0','([0-9]+)'"},
	{"Mail_ru", "https://connect.mail.ru/share_count?url_list=", "shares\":([0-9]+),"},
	{"Vkontakte", "http://vk.com/share.php?act=count&index=1&url=", "VK.Share.count.1, ([0-9]+)"},
}

var socialCount = 11 // 4 Facbook + 7 Others

func getShareCount(url string, key string, regex string, result chan map[string]int){
	totalShare := 0
	responseString := callAPI(url, "GET")
	// Get ShareCount
	re := regexp.MustCompile(regex)
	if totalShareMatch := re.FindStringSubmatch(responseString); len(totalShareMatch) > 0{
		totalShare,_ = strconv.Atoi(totalShareMatch[1])
	}
	
	totalMap := make(map[string]int)
	totalMap[key] = totalShare
	result <- totalMap
}

/* 
	GetAll function accepts the input URL and
	returns Share count for URL across social
	links : Facebook (Likes, Shares, Comments,
	Total), LinkedIn, Pinterest, Odnoklassniki,
	 Mail.ru, Vkontakte, Buffer, Stumbleupon
*/
func GetAll(url string)map[string]int{
	resultsChannel := make(chan map[string]int)
	for _, value := range social{
		if value[0] == "Facebook"{
			go	getShareCount(value[1]+url, value[0]+"Share", value[2], resultsChannel)
			go	getShareCount(value[1]+url, value[0]+"Like", value[3], resultsChannel)
			go	getShareCount(value[1]+url, value[0]+"Comment", value[4], resultsChannel)
			go	getShareCount(value[1]+url, value[0]+"Total", value[5], resultsChannel)
		}else{
			go	getShareCount(value[1]+url, value[0], value[2], resultsChannel)
		}
	}
	for i:=0;i<socialCount;i++{
		for key, value := range <-resultsChannel {
		    resultsMap[key] = value
		}
	}
	return resultsMap		
}

func callAPI(url string, verb string)string{
	resp, err := http.Get(url)
	if err != nil {
		return "FAILED"
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	return string(body)	
}
