---
   统一告警接口
---

# 配置
```lua
    local alert = rock.alert{
        host = "http://xxxx.com/api",
        dns  = "114.114.114.114",
        origin = "security",
        notifer = '[{"name": "张三" , "mobile": "15900010002" , "notify_methods": "weixin,sms"},{},{}]' 
    }
  
    --测试模块功能
    alert.send{
      severity = "high",    
      attribute = "demo",
      subject = "this is security notice",    
      body = "very good xxxxxxxxxxxxxxxxxxxxxxxxxxxxx",    
      tags = "security",
    }
```

# 代码
```golang
    var ud  *alert.Alert
    ud.Do(
      //severity = "high",    
      //attribute = "demo",
      //subject = "this is security notice",    
      //body = "very good xxxxxxxxxxxxxxxxxxxxxxxxxxxxx",    
      //tags = "security",
    )
```
