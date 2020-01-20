
go编写http/websocket代理：
  支持后端多地址轮询转发，支持校验http header中的字段检查登录状态并拒绝未登录用户
  支持websocket代理，websocket消息通过消息队列与后端服务通信

示例：
    根据示例配置可以直接运行如下命令：
    go build proxy.go
    ./proxy
    之后在浏览器访问 http://127.0.0.1，反复刷新可以看到是 www.baidu.com 
    或 www.qq.com 两个页面，且页面展示和真实官网一模一样
    