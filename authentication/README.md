# 将token获取方法整合到dapr的go-sdk中

## 使用方式
1. 启动服务
```
dapr run --app-id serving \
         --app-protocol http \
         --app-port 8080 \
         --dapr-http-port 3500 \
         --log-level debug \
         --components-path ./config \
         go run app.go
```

当终端显示如下即成功启动服务

[![BXNP9s.png](https://s1.ax1x.com/2020/11/11/BXNP9s.png)](https://imgchr.com/i/BXNP9s)

2. 调用服务（invoke）ps：这里我使用postman进行请求

POST/GET
http://localhost:3500/v1.0/invoke/serving/method/

出现下图即成功

[![BXUPaD.png](https://s1.ax1x.com/2020/11/11/BXUPaD.png)](https://imgchr.com/i/BXUPaD)

