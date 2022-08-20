## 抖音项目服务端简单示例

具体功能内容参考飞书说明文档

工程无其他依赖，直接编译运行即可

```shell
go build && ./simple-demo
```

### 功能说明

接口功能已全部实现，但仅能在局域网运行

* 用户登录数据保存在数据库中
* 视频上传后会保存到本地 public 目录中，访问时用本地的ip地址和端口8080访问