# 通知

用户会接收watch仓库的消息：
- issues和评论
- pr和评论
- 提供的的评论

对于未watch的仓库，以下情况也会收到消息：
- @mentions系统
- issue分配
- 是提交的作者或其他人基于自己的提交
- 用户参加的讨论

对于任何通知：
- 都需要notifications或repo 域，说白了就是要以这两个开头才能获取通知

通过请求有优化：
- 通过Last-Modified头
  - 没有新通知，返回304 Not Modified，不会触发限流
- X-Poll-Interval 响应中包含的头，用户提示轮询间隔

通知原因：
- 每个通知都有一个字段reason，表明是哪种事件触发了通知

## 获取通知

- GET /notifications
- 显示当前用户的所有通知，按最近更新时间排序
- 参数：
  - all： 是否显示已读消息
  - participating： 是否过滤出直接参与的活动
  - since： 过滤掉指定时间之前的通知
  - before： 过滤掉指定时间之后的通知

    curl -u username:token https://api.github.com/notifications?since=2014-11-07T08:00:00Z
    curl -i -u release4go:zhuwenk https://api.github.com/notifications\?since\=2019-09-07T08:00:00Z

## 获取某个仓库的通知

- GET /repos/:owner/:repo/notifications
- 同样的认证，同样的参数

    curl -i -u release4go:zhuwenk https://api.github.com/repos/spf13/pflag/notifications\?since\=2019-09-07T08:00:00Z

## 将通知标记为已读

- PUT /notifications
- 如果未读消息太多，会收到一个202Accepted
- 检查是否还有未读消息，可用上面的获取通知，带all=true参数
- 参数：
  - last_read_at： 最后的检查点(这个时间点之后的通知不做处理)

## 将某个仓库的通知标记为已读

- PUT /repos/:owner/:repo/notifications
- 参数和条件都是一样的

## 查看单个thread(主题)

- GET /notifications/threads/:thread_id

## 将thread标记为已读

- PATCH /notifications/threads/:thread_id

## 获取thread的订阅

- 用户判断当前用户是否订阅了thread
- GET /notifications/threads/:thread_id/subscription

## 订阅一个thread

- PUT /notifications/threads/:thread_id/subscription
- 参数：
  - ignored： 是否忽略thread的所有通知，默认false

## 删除一个thread订阅

- DELETE /notifications/threads/:thread_id/subscription
- 屏蔽thread的所有通知
- 除非我们去添加了评论或被人用@mention提及到
