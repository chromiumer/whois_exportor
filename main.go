/*
获取域名Whois信息
更新时间：2021年7月1日
*/
package main

import (
	"bufio"
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/likexian/whois"
	whoisparser "github.com/likexian/whois-parser"
)

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func checkVarExist(v string) (rlt bool) {

	if v == "" {
		return rlt
	}
	rlt = true
	return rlt
}

func ApiMetrics(c *gin.Context) {

	//域名列表文件，每行一个域名
	//file, err := os.Open("/Users/admin/go/src/whois_exporter/domains.list")
	file, err := os.Open("/data/prometheus/prometheus-whois-exporter/domains.list")

	checkErr(err)

	defer file.Close()

	sc := bufio.NewScanner(file)
	for sc.Scan() {

		result, err := whois.Whois(sc.Text())

		if err == nil {

			result, err := whoisparser.Parse(result)
			checkErr(err)

			// 处理域名列表
			var list []string
			for i := 0; i < len(result.Domain.NameServers); i++ {
				list = append(list, result.Domain.NameServers[i])
			}

			nslist := strings.Join(list, ",")

			var status int = 0

			if checkVarExist(result.Domain.WhoisServer) {
				status = 1
			}

			var createdate time.Time
			var updatedate time.Time
			var expdate time.Time

			if result.Domain.Extension == "cn" { // cn域名时间格式单独处理

				createdate, _ = time.Parse("2006-01-02 15:04:05", result.Domain.CreatedDate)
				updatedate, _ = time.Parse("2006-01-02 15:04:05", result.Domain.UpdatedDate)
				expdate, _ = time.Parse("2006-01-02 15:04:05", result.Domain.ExpirationDate)

			} else {

				createdate, _ = time.Parse(time.RFC3339, result.Domain.CreatedDate)
				updatedate, _ = time.Parse(time.RFC3339, result.Domain.UpdatedDate)
				expdate, _ = time.Parse(time.RFC3339, result.Domain.ExpirationDate)
			}
			c.String(http.StatusOK, fmt.Sprintf("%s %v\n", "# Domain Name", sc.Text()))
			c.String(http.StatusOK, fmt.Sprintf("id{domain_name=\"%s\",id=\"%s\"} %d\n", sc.Text(), result.Domain.Id, 1))
			c.String(http.StatusOK, fmt.Sprintf("domain{domain_name=\"%s\",domain=\"%s\"} %d\n", sc.Text(), result.Domain.Domain, 1))
			c.String(http.StatusOK, fmt.Sprintf("name{domain_name=\"%s\",name=\"%s\"} %d\n", sc.Text(), result.Domain.Name, 1))
			c.String(http.StatusOK, fmt.Sprintf("extension{domain_name=\"%s\",extension=\"%s\"} %d\n", sc.Text(), result.Domain.Extension, 1))
			c.String(http.StatusOK, fmt.Sprintf("whoisserver{domain_name=\"%s\",whoisserver=\"%s\"} %d\n", sc.Text(), result.Domain.WhoisServer, status))
			c.String(http.StatusOK, fmt.Sprintf("nameservers{domain_name=\"%s\",nameservers=\"%s\"} %d\n", sc.Text(), nslist, 1))
			c.String(http.StatusOK, fmt.Sprintf("createddate{domain_name=\"%s\",createddate=\"%s\"} %d\n", sc.Text(), createdate.Format("2006-01-02 15:04:05"), createdate.Unix()))
			c.String(http.StatusOK, fmt.Sprintf("updateddate{domain_name=\"%s\",updateddate=\"%s\"} %d\n", sc.Text(), updatedate.Format("2006-01-02 15:04:05"), updatedate.Unix()))
			c.String(http.StatusOK, fmt.Sprintf("expirationdate{domain_name=\"%s\",expirationdate=\"%s\"} %d\n", sc.Text(), expdate.Format("2006-01-02 15:04:05"), expdate.Unix()))

		}
	}

}

func main() {

	route := gin.Default()
	route.GET("/", ApiMetrics)
	route.GET("/metrics", ApiMetrics)

	route.Run(":9095")

}
