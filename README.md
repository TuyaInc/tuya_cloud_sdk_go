# 涂鸦云云对接API SDK golang版本

## 仓库迁移说明
该仓库已经停止更新，请使用`https://github.com/tuya/tuya-cloud-sdk-go`

## 使用前需要做

1. 确定`serverHOST、AccessID、AccessKey`这些值

2. 在你运行本sdk前，显式初始化一次：

   ```
   config.SetEnv(common.URLCN, "AccessID", "AccessKey")
   ```

   

## Example

以获取设备信息接口为例，直接调用device.GetDevice()即可

```golang
    deviceID := "xxx"
    got, err := device.GetDevice(deviceID)
    if err!=nil{
        xxx
    }
    // process got
```

## 目前支持的API

|  Method                   | API                                               | 描述  |
|  ----                     | ----                                              | ----  |
| token.GetTokenAPI         | GET  /v1.0/token?grant_type=1                     | [简单模式获取access_token](https://docs.tuya.com/docDetail?code=K8uuxenajovgv) |
| token.RefreshToken     | GET  /v1.0/token/{{easy_refresh_token}}           | [刷新token](https://docs.tuya.com/docDetail?code=K8uuxfcvdsqwm) |
|  |  |  |
| device.GetDevice          | GET  /v1.0/devices/{{device_id}}                  | [获取设备信息](https://docs.tuya.com/docDetail?code=K8uuxen89a81x) |
| device.GetDeviceFunctions | GET  /v1.0/devices/{deviceId}/functions | [获取设备支持的function列表](https://docs.tuya.com/docDetail?code=K8uuxemwya69p) |
| device.GetDeviceFunctionByCategory | GET  /v1.0/functions/{category} | [根据category获取function列表](https://docs.tuya.com/docDetail?code=K8uuxemym7qkt) |
| device.GetDeviceStatus | GET  /v1.0/devices/{{device_id}}/status           | [获取设备功能点的信息](https://docs.tuya.com/docDetail?code=K8uuxen4ux749) |
| device.GetDeviceList | GET  /v1.0/devices/status?device_ids={{device_id}} | [批量获取设备状态](https://docs.tuya.com/docDetail?code=K8uuxenar6kgc) |
| device.PostDeviceCommand | POST  /v1.0/devices/{{device_id}}/commands        | [设备指令下发](https://docs.tuya.com/docDetail?code=K8uuxfcxbpwlo) |
| device.DeleteDevice | DELETE  /v1.0/devices/{device_id} | [移除设备](https://docs.tuya.com/docDetail?code=K8uuxemvwtp3z) |
|  |  |  |
| user.PostUserRegister   | POST  /v1.0/apps/{schema}/user | [云端用户注册](https://docs.tuya.com/docDetail?code=K8uuxfcuhc2ei) |
| user.GetUserList   | GET  /v1.0/apps/{schema}/users?page_no=&page_size= | [获取用户列表](https://docs.tuya.com/docDetail?code=K8uuxemwe9kwb) |
| user.GetDeviceListByUid | GET /v1.0/users/{uid}/devices | [根据用户id获取设备列表](https://docs.tuya.com/docDetail?code=K8uuxfcuesrh7) |
|  |  |  |
| device.dn.PostDeviceToken | POST /v1.0/devices/token       | [生成设备配网token](https://docs.tuya.com/docDetail?code=K8uuxfcujsk6n) |
| device.dn.GetDevicesByToken | POST  /v1.0/devices/tokens/{{pair_token}}         | [根据token获取设备列表](https://docs.tuya.com/docDetail?code=K8uuxemz174o3) |



## 常见问题

### 关于refreshToken接口

注意： refreshToken接口会返回一个新的access_token，即使旧的token还未过期。

这个逻辑在GetToken方法中已经做了，用户一般不需要调用refresh接口。

### 每次调用api之前，是否需要获取token或者刷新token？

不需要，这层逻辑已经在api方法中实现了。token信息会缓存到内存中。

### 调用某个接口时，如果token已经过期，需要手动调用refresh-token接口？

不需要，在GetToken()方法实现中，会检查token是否过期。如果过期会去重新拉取。

### 如果你的token会在多个节点中去刷新，那么你需要自己实现common.TokenLocalManage interface
涂鸦的云token，只保证面向的用户级别刷新不会有问题，但是一个用户的token在多个节点并发刷新，就会导致一个节点是成功的，
其他都是失败；
因为 `获取token`接口会返回一个access_token、refresh_token，但是 `刷新token`接口 会把当前的refresh_token 刷掉，会产生一个新的，旧的失效；

### api方法的异常信息和error需要如何处理？

接口如果返回error，一般可以为url错误或者json解析出错，可联系涂鸦相关人员帮忙修改

如果error为空，但是response的success字段为false，开发者可以根据Msg字段的详细错误信息进行排查

### 获取设备列表接口，如果多个deviceID，需要怎么拼接？

多个deviceID，以英文逗号分割拼接

### 获取用户列表接口中，schema指的是什么？

创建APP-SDK后，详情页面的渠道标识符就是schema

### v1.0/devices/tokens/{{pair_token}}接口，pair_token是指什么？如何获取？

pair_token是指app下的某个用户的配网token，可以从v1.0/devices/token获取。

### 如果SDK中的API没有及时更新，如何自己实现一个API？

有两种方法：

1. 可以通过实现common.APIRequest这个interface，如果是POST请求需要把RequestBody接口也实现了。
然后调用DoAPIRequest()即可。具体可以参考UserDevicesAPI实现
2. 提个issue，我们会及时更新 ^_^


