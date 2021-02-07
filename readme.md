---
   统一告警接口
---

# 配置
```lua
    local alert = rock.alert{
        host = "http://xxxx.com/api",
        dns  = "114.114.114.114",
        origin = "security",
        notifer = '[{"name": "张三" , "mobile": "15900010002" , "email":"zhangsan@qq.com" , "serial_number": "000000" , "notify_methods": "weixin,sms"},{},{}]' 
    }
```
