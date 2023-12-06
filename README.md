# wechatbot

> 本项目是 fork 他人的项目来进行学习和使用，本项目可以将个人微信化身GPT机器人,
> 项目基于[openwechat](https://github.com/eatmoreapple/openwechat) 开发。

### 目前实现了以下功能

* GPT机器人模型热度可配置
* 提问增加上下文
* 指令清空上下文（指令：根据配置）
* 机器人群聊@回复
* 机器人私聊回复
* 私聊回复前缀设置
* 好友添加自动通过
* 支持自定义 openai api url
* 支持 windows/linux/mac
* 支持最新 openai api


# 常见问题
* 如无法登录 login error: write storage.json: bad file descriptor 删除掉storage.json文件重新登录。
* 如无法登录 login error: wechat network error: Get "https://wx.qq.com/cgi-bin/mmwebwx-bin/webwxnewloginpage": 301 response missing Location header 一般是微信登录权限问题，先确保PC端能否正常登录。
* 其他无法登录问题，依然尝试删除掉storage.json文件，结束进程(linux一般是kill -9 进程id)之后重启程序，重新扫码登录，(如为docket部署，Supervisord进程管理工具会自动重启程序)。
* linux中二维码无法扫描，缩小命令行功能，让二维码像素尽可能清晰。（无法从代码层面解决）
* 机器人一直答非所问，可能因为上下文累积过多。切换不同问题时，发送指令：启动时配置的`session_clear_token`字段。会清空上下文

# 使用前提

> * 有openai账号，并且创建好api_key，注册事项可以参考[此文章](https://juejin.cn/post/7173447848292253704)。
> * 也可以介入local LLM + OpenAI-compatible api 推荐[轻量化openai inference engine](https://github.com/janhq/nitro)
> * 微信必须实名认证。

# 注意事项

> * 项目仅供娱乐，滥用可能有微信封禁的风险，请勿用于商业用途。
> * 请注意收发敏感信息，本项目不做信息过滤。

# Quickstart


```sh
# 运行项目
$ git clone https://github.com/kevinbin/wxbot.git
$ cd wxbot
$ make docker
# 指定环境变量
$ docker run -itd --name wechatbot -e APIKEY=openai_api_key wechatbot:latest
# 或使用配置文件
$ ./start.sh
# 查看二维码
$ docker exec -it wechatbot tail -f -n 30 /app/run.log
```

# 配置文件说明

````
base_url: 指定 openai 接口地址,可以配合本地 LLM使用
api_key：openai api_key
auto_pass:是否自动通过好友添加
session_timeout：会话超时时间，默认600秒，单位秒，在会话时间内所有发送给机器人的信息会作为上下文。
max_tokens: GPT响应字符数，默认值1024。max_tokens会影响接口响应速度，字符越大响应越慢。
model: GPT选用模型，默认gpt-3.5-turbo，具体选项参考openai官网
temperature: GPT热度，0到1，默认0.5。数字越大创造力越强，但更偏离训练事实，越低越接近训练事实
reply_prefix: 私聊回复前缀
session_clear_token: 会话清空口令，默认`会话清空`
````
