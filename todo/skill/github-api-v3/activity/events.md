# 事件api

- 只读
- 支持github的多种活动流
- 为了轮询，使用Etag头进行了优化
- 使用api时，如果没有新的事件被触发，会响应一个304 Not Modified，限流也不会被触发
- 在响应头中还指定了 X-Poll-Interval，表明轮询的间隔，单位秒(尽量遵循github服务器给的间隔)
- 事件是支持分页的
  - 此时不支持per_page选项，每页固定为30项
  - 最多支持10页，也就是说最多可查询300个事件
  - 超过90天的事件是取不到的(即使是没有超过300个事件)，时间线也是一样

    所有的事件响应格式都是一样的
    payload/repo/actor/org 4个对象信息
    还有就是和事件相关的几个标签属性

    Status: 200 OK
    Link: <https://api.github.com/resource?page=2>; rel="next",
          <https://api.github.com/resource?page=5>; rel="last"
    [
      {
        "type": "Event",
        "public": true,
        "payload": {
        },
        "repo": {
          "id": 3,
          "name": "octocat/Hello-World",
          "url": "https://api.github.com/repos/octocat/Hello-World"
        },
        "actor": {
          "id": 1,
          "login": "octocat",
          "gravatar_id": "",
          "avatar_url": "https://github.com/images/error/octocat_happy.gif",
          "url": "https://api.github.com/users/octocat"
        },
        "org": {
          "id": 1,
          "login": "github",
          "gravatar_id": "",
          "url": "https://api.github.com/orgs/github",
          "avatar_url": "https://github.com/images/error/octocat_happy.gif"
        },
        "created_at": "2011-09-06T17:26:27Z",
        "id": "12345"
      }
    ]

## 显示公共事件

- 公共活动事件延时了5分钟
  - 大多数最近的事件，用公共事件api来查询，只会查询5分钟之前的事件
- GET /events

    curl -i -u release4go:zhuwenk -H 'If-None-Match:"45d757cd54d75ea3d6b7917c133e7951"' https://api.github.com/events 
    会查询最近30个事件
    响应的头信息会包含以下信息：
    X-Poll-Interval: 60
    Link: <https://api.github.com/events?page=2>; rel="next", <https://api.github.com/events?page=10>; rel="last"

## 显示仓库事件

- GET /repos/:owner/:repo/events
- curl -i -H 'If-None-Match: "bede2a7eec0349bc9d218dc136cf6675"' https://api.github.com/repos/63isOK/golearn/events
- 显示的是指定仓库的事件

## 显示仓库的issue事件

- GET /repos/:owner/:repo/issues/events
- issue事件的格式其他事件的格式有些不同
- issue中有很多事件，包括了分配/合并/关闭/引用 等等

    curl -i https://api.github.com/repos/fight100year/golearn/issues/events

## 显示一个仓库的，和网络相关的公共事件

- GET /networks/:owner/:repo/events
- 显示的都是和网络相关的push 创建分支等

    curl -i https://api.github.com/networks/fight100year/golearn/events

## 显示一个组织的公共事件

- GET /orgs/:org/events

    curl -i https://api.github.com/orgs/fight100year/events

## 显示一个用户接收到的事件

- watch某个仓库，或是follow某个用户，收到的事件
- 请求中如果添加了认证信息，也会显示私有事件;否则只显示公共事件
- GET /users/:username/received_events

    curl -i https://api.github.com/users/63isOK/received_events

## 只显示用户接收到的公共事件

- GET /users/:username/received_events/public

## 显示用户执行的事件

- 显示的某用户执行的事件 
- 请求中如果添加了认证信息，也会显示私有事件;否则只显示公共事件

    curl -i https://api.github.com/users/63isOK/events

## 只显示功胡执行的公共事件

- GET /users/:username/events/public

## 显示一个组织的事件

- 就是组织的看板(dashboard)，需要认证
- GET /users/:username/events/orgs/:org

