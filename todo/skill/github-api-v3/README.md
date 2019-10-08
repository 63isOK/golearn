# rest api v3 

- overview
  - [基本术语](/todo/skill/github-api-v3/overview/hello.md)
  - [媒体类型](/todo/skill/github-api-v3/overview/media-types.md)
  - [OAuth认证api](/todo/skill/github-api-v3/overview/oauth-api.md)
  - [其他认证方式](/todo/skill/github-api-v3/overview/other-auth.md)
  - [问题解决](/todo/skill/github-api-v3/overview/trouble-shooting.md)
  - [api预览](/todo/skill/github-api-v3/overview/api-previews.md)
  - [版本](/todo/skill/github-api-v3/overview/versions.md)
- activity api，也就是对"通知/订阅/时间线"上的一些交互性提供功能
  - [事件](/todo/skill/github-api-v3/activity/events.md)，只读接口
  - [事件类型和事件payload](/todo/skill/github-api-v3/activity/event-payload.md)
  - [feed订阅](/todo/skill/github-api-v3/activity/feeds.md),给认证用户的，包含摘要信息的订阅api
  - [通知](/todo/skill/github-api-v3/activity/notifications.md)，新评论会通知到指定用户，也可以将消息标记为已读
  - [打星](/todo/skill/github-api-v3/activity/starring.md)，github打星功能，和点赞是一个意思
  - [watch](/todo/skill/github-api-v3/activity/watching.md)，watch他人项目，可接收对应的通知消息
- check api, 通过这些api来发送集成运行的结果(除了编译的成功失败，还可以包含其他信息)
  - 这些api是给github app使用的，eg：travis ci，github action也用到了这些api
