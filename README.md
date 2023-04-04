# Dongtai_USB

# 转换器配置文件(放在当前目录下)：config-tutorial.ini.example 
```
ip: 访问地址白名单
iast_url：iast地址
dast_token： iast对应dast_token
xray_url：商业版xray地址
xray_token： 商业版xray-token
```

# 启动
```
mv config-tutorial.ini.example config-tutorial.ini
docker-compose up -d
添加代理：IP:10802
```
