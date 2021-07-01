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
