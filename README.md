### whois_exportor
#### prometheus whois exportor

#### 域名到期时间监控

##### 1️⃣1️⃣1️⃣ 依赖模块及使用
```
go get github.com/gin-gonic/gin
go get github.com/likexian/whois
go get github.com/likexian/whois-parser

git clone https://github.com/chromiumer/whois_exportor.git

cd whois_exportor

根据需要修改域名列表路径

file, err := os.Open("/data/prometheus/prometheus-whois-exporter/domains.list")

go build main.go
```

##### 2️⃣2️⃣2️⃣ 启动配置

```
[program:prometheus-whois-exporter]
command = /data/prometheus/prometheus-whois-exporter/whois_exporter
directory = /data/prometheus/prometheus-whois-exporter
user = root
autostart = true
autorestart=true
redirect_stderr=true
stdout_logfile = /data/prometheus/prometheus-whois-exporter/stdout.log

```
##### 3️⃣3️⃣3️⃣ prometheus添加exporter

```
  - job_name: 'domain_name_whois_info'
    static_configs:
    - targets: ["172.17.7.7:9095"]
```

##### 4️⃣4️⃣4️⃣告警规则
##### alert rules
```
- name: Domain_Name_Alert
  rules:
  - alert: domain_name_exp_alert
    expr: expirationdate - time() < 86400 * 90
    for: 10s
    labels:
      severity: critical
    annotations:
      description: "Domain Name {{ $labels.domain_name}} expiration date less than 90 days"
```

```
# Domain Name mysql.com
id{domain_name="mysql.com",id="3383022_DOMAIN_COM-VRSN"} 1
domain{domain_name="mysql.com",domain="mysql.com"} 1
name{domain_name="mysql.com",name="mysql"} 1
extension{domain_name="mysql.com",extension="com"} 1
whoisserver{domain_name="mysql.com",whoisserver="whois.markmonitor.com"} 1
nameservers{domain_name="mysql.com",nameservers="ns1.p04.dynect.net,ns2.p04.dynect.net,ns3.p04.dynect.net,ns4.p04.dynect.net,orcldns1.ultradns.com,orcldns2.ultradns.net"} 1
createddate{domain_name="mysql.com",createddate="1999-02-03 05:00:00"} 918018000
updateddate{domain_name="mysql.com",updateddate="2021-01-27 10:22:00"} 1611742920
expirationdate{domain_name="mysql.com",expirationdate="2022-02-28 02:45:40"} 1646016340
# Domain Name github.com
id{domain_name="github.com",id="1264983250_DOMAIN_COM-VRSN"} 1
domain{domain_name="github.com",domain="github.com"} 1
name{domain_name="github.com",name="github"} 1
extension{domain_name="github.com",extension="com"} 1
whoisserver{domain_name="github.com",whoisserver="whois.markmonitor.com"} 1
nameservers{domain_name="github.com",nameservers="dns1.p08.nsone.net,dns2.p08.nsone.net,dns3.p08.nsone.net,dns4.p08.nsone.net,ns-1283.awsdns-32.org,ns-1707.awsdns-21.co.uk,ns-421.awsdns-52.com,ns-520.awsdns-01.net"} 1
createddate{domain_name="github.com",createddate="2007-10-09 18:20:50"} 1191954050
updateddate{domain_name="github.com",updateddate="2020-09-08 09:18:27"} 1599556707
expirationdate{domain_name="github.com",expirationdate="2022-10-09 18:20:50"} 1665339650
```

💢💢💢💢💢💢💢完结🌻🌻🌻撒花💢💢💢💢💢💢💢
