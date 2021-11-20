/*
 * @lc app=leetcode.cn id=811 lang=golang
 *
 * [811] 子域名访问计数
 */

// @lc code=start
import "strings"
import "strconv"
func subdomainVisits(cpdomains []string) []string {
    m := make(map[string]int)
    for _,v := range cpdomains {
        tmp := strings.Split(v, " ")  // 截取 访问次数
        if num,err := strconv.Atoi(tmp[0]);err==nil{  // string 转化成 int
            for ;len(tmp)==2; tmp = strings.SplitAfterN(tmp[1], ".", 2){
                m[tmp[1]] +=num
            }
        }
    }
    rs := []string{}
    for domains,cnt := range m {
        rs = append(rs, strconv.Itoa(cnt)+" "+ domains)  // map 转化成 []string
    }
    return rs
}

/**
	hashmap 操作， 和 golang string strconv 两个 API的使用
	时间复杂度： 2n
	空间复杂度:  n 
*/

// @lc code=end

