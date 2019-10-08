# watch

- watch 之后，就可以收到仓库相关的通知

## 显示仓库的所有订阅者

- GET /repos/:owner/:repo/subscribers

## 显示某人订阅的所有的仓库

- GET /users/:username/subscriptions
- GET /user/subscriptions 认证用户用这个

## 判断某人是否有订阅某仓库

- GET /repos/:owner/:repo/subscription (和显示仓库的所有订阅者只差几个字母)
- 200 OK 表示已订阅
- 404 Not Found 表示未订阅

## 订阅某仓库

- PUT /repos/:owner/:repo/subscription
- 参数：
  - subscribed： 表示是否接收该仓库的通知
  - ignored： 是否忽略该仓库的通知

## 取消订阅某仓库

- DELETE /repos/:owner/:repo/subscription


