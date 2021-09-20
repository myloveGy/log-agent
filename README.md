# LogAgent 

监听日志写入，然后输出到redis、获取mongodb中

## 配置说明

使用ini配置、配置文件在目录下的.ini 中

| 配置项 | 说明 |
|-------|-----|
| tail | 需要监听的日志，多个使用,分隔 |

```ini
app_name=日志收集处理
tail=log.laravel,log.wechat
[redis]
host=127.0.0.1
port=6379
password=
db=0
[mongodb]
host=localhost
port=27017
db_name=log
[log.laravel]
path=/www/admin/storage/logs/
filename=laravel-2006-01-02.log
rule=true
type=json
driver=redis
name=laravel
[log.wechat]
path=/www/admin/storage/logs/
filename=wechat-2006-01-02.log
rule=true
type=json
driver=mongo
name=wechat
```

### 日志配置项说明
| 配置项 | 说明 | 
|-------|-----|
| path | 监听文件路径 |
| filename | 监听文件名称 |
| rule | 文件是否格式化(默认false) |
| type | 读取到每一行的数据格式(取决于你日志写入格式) 目前: `json`, `text`|
| driver | 写入的类型; 目前: `redis`, `mongo` |
| name | 写入的字段名称(对应`mongo`的表名称, `redis` key 名称)

### 监听文件名称： path + 格式化后的 `filename` 

例如： 

目前处理 laravel 的日志 laravel-2006-01-02.log

rule=true 会被格式化: laravel-年-月-日.log

