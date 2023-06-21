## Dongtai_USB
<br />

### 转换器配置文件(放在当前目录下)：config-tutorial.ini.example 
```
ip: xray访问地址白名单,默认不需要修改
iast_url：iast地址
dast_token： iast对应dast_token
type: 扫描器类型（xray）
xray_url：商业版xray地址
xray_token： 商业版xray-token
```

### 启动
```
mv config-tutorial.ini.example config-tutorial.ini
docker-compose up -d
添加代理：IP:10802
```
<br />
<br />

## 开发配置
### 数据上报流程: 
```
用户 ==> 浏览器代理 ==> mitmproxy ==> xray ==> dongtai_usb ==> 洞态IAST
```
![Alt text](image-1.png)


1. 需要添加两个 header 响应头
```
Dt-Request-Id
dt-mark-header
```
2. 通过 mitmproxy 自动添加 dt-mark-header 响应头
```
flow.request.headers["dt-mark-header"] = uuid.uuid4().hex
```
3. 通过 洞态agent 自动添加 Dt-Request-Id  响应头
```
dt-request-id : <agent_id>.<uuid>
```


4. 发送给IAST的数据格式，可参考如下结构体Response，可直接调用
service.Client(Response) 发送数据给洞态iast,
```
type Response struct {
	VulName         string            `json:"vul_name"`
	Detail          string            `json:"detail"`
	VulLevel        string            `json:"vul_level"`
	Urls            []string          `json:"urls"`
	Payload         string            `json:"payload"`
	CreateTime      int64             `json:"create_time"`
	VulType         string            `json:"vul_type"`
	RequestMessages []RequestMessages `json:"request_messages"`
	Target          string            `json:"target"`
	DtUUIDID        []string          `json:"dt_uuid_id"`
	AgentID         []string          `json:"agent_id"`
	DongtaiVulType  []string          `json:"dongtai_vul_type"`
	Dtmark          []string          `json:"dt_mark"`
	DastTag         string            `json:"dast_tag"`
}
type RequestMessages struct {
	Request  string `json:"request"`
	Response string `json:"response"`
}

type Target struct {
	Params []struct {
		Path     []string `json:"path"`
		Position string   `json:"position"`
	} `json:"params"`
	URL string `json:"url"`
}
```

5. Response 结构体详解
```
{
    "vul_name": "",#漏洞名 格式为 target+漏洞类型
    "detail":"", #漏洞详情
    "vul_level": "HIGH", #HIGH,MEDIUM,LOW,NOTE 漏洞等级，对应现在洞态的4个等级
    "urls":[""],# 黑盒扫描发送的多个 url 地址
    "payload":"", #  黑盒扫描触发漏洞的 payload, 可为空
    "create_time":1679020853, # 时间戳(秒)
    "vul_type":"",#黑盒扫描的漏洞类型
    "request_messages":[{ # 一组扫描对应的所有请求和响应信息
        "request":"",
        "response":""
      }
    ],
    #以下为dongtai对接相关信息。
    "dt_mark": [""], # dt-mark-header 的值
    "target":"", # 原始请求地址
    "dt_uuid_id":[""], # 需要在 dt-request-id 响应头拆分出来
    "agent_id":[""], # 需要在 dt-request-id 响应头拆分出来
    "dongtai_vul_type":[""],# 洞态的漏洞类型, 多个类型，为空数组即对应所有调用链漏洞
    "dast_tag":"", # 所集成的黑盒扫描器标识
}
```
6. 上报消息示例
```
{
	"vul_name": "http://192.168.1.37:8001/xxe/Digester/vuln xxe",
	"detail": "在http://192.168.1.37:8001/xxe/Digester/vuln发现了xxe",
	"vul_level": "MEDIUM",
	"urls": ["/xxe/Digester/vuln"],
	"payload": "\u003c?xml version=\"1.0\"?\u003e\u003c!DOCTYPE ANY [\u003c!ENTITY content SYSTEM \"XRAY_RURL\"\u003e]\u003e\u003ca\u003e\u0026content;\u003c/a\u003e",
	"create_time": 1687334492,
	"vul_type": "xxe",
	"request_messages": [{
		"request": "POST /xxe/Digester/vuln HTTP/1.1\r\nHost: 192.168.1.37:8001\r\nUser-Agent: Xray_Test\r\nContent-Length: 123\r\nAccept: */*\r\nContent-Type: application/xml\r\nCookie: XSRF-TOKEN=b8442e83-8a47-4703-9b4c-b0e74dc214f8; JSESSIONID=60EB9A6A04AAF9416210F3951788EA81\r\nDt-Dast: Xray\r\nDt-Mark-Header: b9c9ddd59de04c479577851db7dd0242\r\nXray: x\r\nAccept-Encoding: gzip\r\n\r\n\u003c?xml version=\"1.0\"?\u003e\u003c!DOCTYPE ANY [\u003c!ENTITY content SYSTEM \"http://177.7.0.11:8091/i/4d6816/1pqo/b9t8/\"\u003e]\u003e\u003ca\u003e\u0026content;\u003c/a\u003e",
		"response": "HTTP/1.1 200 \r\nCache-Control: no-cache, no-store, max-age=0, must-revalidate\r\nContent-Length: 22\r\nContent-Type: text/plain;charset=UTF-8\r\nDate: Wed, 21 Jun 2023 08:01:31 GMT\r\nDongtai: v1.10.0\r\nDt-Request-Id: 17.0ce4ae9dca114fd084a2070cf96e99cd\r\nExpires: 0\r\nPragma: no-cache\r\nX-Application-Context: application\r\nX-Content-Type-Options: nosniff\r\nX-Frame-Options: DENY\r\nX-Xss-Protection: 1; mode=block\r\n\r\nDigester xxe vuln code"
	}],
	"target": "http://192.168.1.37:8001/xxe/Digester/vuln",
	"dt_uuid_id": ["0ce4ae9dca114fd084a2070cf96e99cd"],
	"agent_id": ["17"],
	"dongtai_vul_type": ["xxe"],
	"dt_mark": ["b9c9ddd59de04c479577851db7dd0242"],
	"dast_tag": "Xray"
}
```


## 如何开发一个新的黑盒扫描器
```
1. 在dongtai_usb/目录下创建一个新的文件夹，文件夹名字为黑盒扫描器的名字
2. 在新建的文件夹下创建三个子目录,可参考xray目录
	1. dongtai_usb/xxx/engine/  # 数据处理转换 代码
	2. dongtai_usb/xxx/model/ # 请求结构体 代码
	3. dongtai_usb/xxx/request/ # 接收或拉取请求实现 代码

3. 漏洞类型可使用map对应，参考: dongtai_usb/xray/model/vultype.go
	1. Vultype 为漏洞类型命名
	2. VulLevel 为漏洞等级命名
```
### 漏洞类型等级对应关系
```
func Vultype() map[string]string {
	return map[string]string{
		//xray漏洞类型     //洞态漏洞类型
		"xss":            "reflected-xss",
		"sqldet":         "sql-injection",
		"cmd-injection":  "cmd-injection",
		"path-traversal": "path-traversal",
		"xxe":            "xxe",
		"ssrf":           "ssrf",
		"brute-force":    "crypto-bad-ciphers",
		"redirect":       "unvalidated-redirect",
	}
}
func VulLevel() map[string]string {
	return map[string]string{
		//xray漏洞名字     //洞态漏洞等级
		"xss":            "MEDIUM",
		"sqldet":         "HIGH",
		"cmd-injection":  "HIGH",
		"path-traversal": "HIGH",
		"xxe":            "MEDIUM",
		"ssrf":           "ssrf",
		"brute-force":    "LOW",
		"redirect":       "LOW",
	}
}
```
### 运行方式为两个场景
```
推流模式(webhook)：即黑盒扫描器主动推送数据给dongtai_usb
拉流模式(cron): 即dongtai_usb主动从黑盒扫描器定时拉取数据
```

![Alt text](image-2.png)
### 开发完成后在main方法添加启用即可，如xray
```
推流模式(webhook):
router.POST("/v1/xray", USB_Xray.Xray)

拉流模式(cron):
usb.Cron("xray", USB_Xray.Xray_cron)
```