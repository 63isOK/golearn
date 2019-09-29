# github rest api v3

本文的主要关注点是 [github rest api v3](https://developer.github.com/v3/https://developer.github.com/v3/)

## api 版本

- 本次聊到的是v3版本，
- 下面所有的请求，默认都是v3版本的，因为https://api.github.com接收的默认是v3版本
- 现在v4版本也出来了，也推荐在请求中显示指明版本

    Accept: application/vnd.github.v3+json

## schema 架构

- 所有的api都是通过https来访问https://api.github.com的
- 所有的数据收发都是序列化成json格式的
- 空白字段，表示是null，而不是忽略
- 时间戳格式遵循iso 8601格式(这也是rest api的风格)

    curl -i https://api.github.com
    // -i 表示显示响应头

    HTTP/1.1 200 OK
    Date: Sun, 29 Sep 2019 01:03:17 GMT
    Content-Type: application/json; charset=utf-8
    Content-Length: 2165
    Server: GitHub.com
    Status: 200 OK
    X-RateLimit-Limit: 60
    X-RateLimit-Remaining: 59
    X-RateLimit-Reset: 1569722597
    Cache-Control: public, max-age=60, s-maxage=60
    Vary: Accept
    ETag: "7dc470913f1fe9bb6c7355b50a0737bc"
    X-GitHub-Media-Type: github.v3; format=json
    Access-Control-Expose-Headers: ETag, Link, Location, Retry-After, X-GitHub-OTP, X-RateLimit-Limit, X-RateLimit-Remaining, X-RateLimit-Reset, X-OAuth-Scopes, X-Accepted-OAuth-Scopes, X-Poll-Interval, X-GitHub-Media-Type
    Access-Control-Allow-Origin: *
    Strict-Transport-Security: max-age=31536000; includeSubdomains; preload
    X-Frame-Options: deny
    X-Content-Type-Options: nosniff
    X-XSS-Protection: 1; mode=block
    Referrer-Policy: origin-when-cross-origin, strict-origin-when-cross-origin
    Content-Security-Policy: default-src 'none'
    Vary: Accept-Encoding
    X-GitHub-Request-Id: 8286:71B1:29441:39D39:5D9002D5

    {
      "current_user_url": "https://api.github.com/user",
      "current_user_authorizations_html_url": "https://github.com/settings/connections/applications{/client_id}",
      "authorizations_url": "https://api.github.com/authorizations",
      "code_search_url": "https://api.github.com/search/code?q={query}{&page,per_page,sort,order}",
      "commit_search_url": "https://api.github.com/search/commits?q={query}{&page,per_page,sort,order}",
      "emails_url": "https://api.github.com/user/emails",
      "emojis_url": "https://api.github.com/emojis",
      "events_url": "https://api.github.com/events",
      "feeds_url": "https://api.github.com/feeds",
      "followers_url": "https://api.github.com/user/followers",
      "following_url": "https://api.github.com/user/following{/target}",
      "gists_url": "https://api.github.com/gists{/gist_id}",
      "hub_url": "https://api.github.com/hub",
      "issue_search_url": "https://api.github.com/search/issues?q={query}{&page,per_page,sort,order}",
      "issues_url": "https://api.github.com/issues",
      "keys_url": "https://api.github.com/user/keys",
      "notifications_url": "https://api.github.com/notifications",
      "organization_repositories_url": "https://api.github.com/orgs/{org}/repos{?type,page,per_page,sort}",
      "organization_url": "https://api.github.com/orgs/{org}",
      "public_gists_url": "https://api.github.com/gists/public",
      "rate_limit_url": "https://api.github.com/rate_limit",
      "repository_url": "https://api.github.com/repos/{owner}/{repo}",
      "repository_search_url": "https://api.github.com/search/repositories?q={query}{&page,per_page,sort,order}",
      "current_user_repositories_url": "https://api.github.com/user/repos{?type,page,per_page,sort}",
      "starred_url": "https://api.github.com/user/starred{/owner}{/repo}",
      "starred_gists_url": "https://api.github.com/gists/starred",
      "team_url": "https://api.github.com/teams",
      "user_url": "https://api.github.com/users/{user}",
      "user_organizations_url": "https://api.github.com/user/orgs",
      "user_repositories_url": "https://api.github.com/users/{user}/repos{?type,page,per_page,sort}",
      "user_search_url": "https://api.github.com/search/users?q={query}{&page,per_page,sort,order}"
    }

### 显示摘要

- 当获取一个资源列表时，响应中一般会包含资源的一个属性集
- 而有时候，这个属性集不会将详细的信息带上，所以这个就叫摘要

    curl https://api.github.com/orgs/fight100year/repos
    获取github100year组织的所有仓库信息，响应中会包含每个仓库的摘要信息

### 详细信息

- 一般访问单个资源，响应会将资源的所有属性全部显示出来
- 这里一般叫详细信息(对应列表的摘要)

    curl https://api.github.com/repos/fight100year/golearn
    获取当资源的信息

## 认证

- v3 总共有2大类认证方式，共3种认证方法
- 部分请求需要认证信息，如果请求中没有，响应会返回404 Not Found
- 错误的认证信息，会返回401 未认证成功

### 基础认证

    手动输入密码
    curl -u "63isOK" https://api.github.com
    用户名和密码一起输入
    curl -u "63isOK:zhuwenka" https://api.github.com

### OAuth2 token认证

- 通过http头进行发送
- 这里利用的是token认证方式，基础认证也是通过http头进行发送，只是认证方式不一样而已
- 这种方式是github推荐使用的方式
- 这种方式有一点不好，因为请求路径中包含了url，所以可以被日志记录，少了点安全性

    curl -H "Authorization: token OAUTH-TOKEN" https://api.github.com 
    通过http头进行发送

### OAuth2 key/secret认证

    curl 'https://api.github.com/users/whatever?client_id=xxxx&client_secret=yyyy'

- 这种方案并不是为了作为一个用户角色来认证
- 而是为了解决网速限制
- 换句话说：这种认证方式是给app去做认证的，而不是给用户做认证用的

### 认证响应合集

- 404 需要认证的时候没有认证信息
- 401 错误的认证信息
- 403 短时间内受到很多无效认证后，github会返回403表示短时间内拒绝用户的所有认证请求 

## 参数

- 很多api都是带有参数的
- 对于GET请求，在路径中，任何未指定为segment的都被认为是参数

    curl -i "https://api.github.com/repos/fight100year/golearn/issues?state=closed"
    fight100year 被认为是所有者
    golearn被认为是仓库
    closed被显示指定为issues的状态

- 对于POST/PATHC/PUT/DELETE请求，不再url中的参数，需要丢到json中
- 此时http头中的Content-Type要指定为'application/json'

    curl -i -u 63isOK:zhuwenkai001 -d '{"scopes":["public_repo"]}' https://api.github.com/authorizations

## 根节点

- 可以利用curl https://api.github.com获取所有节点分类
- 说白了就是你能获取哪些资源，都可以用这个命令来查看

## graphql 全局node的id

- 可以在rest api v3中获取node的id，然后在graphql中使用
- GraphQL也是一种api风格标准，和rpc/rest一起称为3种api风格

## 客户端错误

- 在http中：
  - 1xx： 信息提示，client会收到多个1xx
  - 2xx： 成功
  - 3xx： 重定向
  - 4xx： 客户端错误
  - 5xx： 服务端错误
- 在rest中，4xx表示客户端错误
  - 400 错误请求，一般是请求参数不是一个有效的json格式
  - 400 ，也可能是json格式中的数据类型是错误的
  - 422 ，field错误，eg：json格式中缺少了某个field

```json
// HTTP/1.1 422 Unprocessable Entity
// Content-Length: 149

{
  "message": "Validation Failed",
  "errors": [
    {
      "resource": "Issue",
      "field": "title",
      "code": "missing_field"
    }
  ]
}
```

- 客户端422错误，会有以下几种错误码：
  - missing 资源不存在
  - missing_field 资源不存在这个field
  - invalid 资源的field的格式是无效的
  - already_exists 已经存在相同的field通向的value(一般出在需要唯一的场景)
  - 自定义错误，需要带一个message来描述这个错误码

## 重定向

- 接收一个http重定向并不是错误
- 一般http头信息中会有一个Location field会指明新的rui
- 重新请求新的uri即可
- 支持的返回码：
  - 301，永久重定向
  - 302/307, 临时重定向
  - 其他重定向的返回码可参考http1.1 spec

## http支持的动词 verbs

- HEAD，只想获取头信息
- GET，获取资源
- POST，创建资源
- PATCH，用部分json数据去更新资源，局部更新用这个。这个动作较新，所以不常用
- PUT，替换资源或集合，也就是更新资源
- DELETE，删除资源

## 超媒体 hypermedia

- 每个资源都可能有多个\*\_url属性，用于链接其他资源
- 这些链接符合RFC6570, 也可以自己扩展模版

## 分页 pagination

- 如果相应包含太多元素，会按30进行分页
- 可通过?page参数指定后面的页
- 30是默认的，也可以用?per_page指定每页100
- 不是所有的api都支持分页，也有部分不支持
- 分页都是从1开始，忽略这个参数，默认取第一页

    curl 'https://api.github.com/user/repos?page=2&per_page=100'

### 链接头

    Link: <https://api.github.com/user/repos?page=3&per_page=100>; rel="next",
    <https://api.github.com/user/repos?page=50&per_page=100>; rel="last"

- 链接头里也可以指定分页参数
- Link一般出现在响应头中，一个相应头会包含多个Link
- rel可能出现的值：
  - next：和当前页的关系是：link指向的是下一页
  - last：link指向的是最后一页
  - first：link指向的是第一页
  - prev： link指向的是当前页的前一页

## rate limiting 速率限制

- 如果api请求是基于认证的，不管是基础认证还是OAuth认证
- 一小时最多有5000次请求
- 认证请求是基于认证用户的，也就是同一个用户的认证，最多5000次/小时
- 对于非认证请求，60次/小时
- 非认证请求是基于原始ip地址的，不是基于用户的

    curl -i https://api.github.com/users/release4go
    curl -i https://api.github.com/users/63isOK
    结果显示的X-RateLimit-Remaining就是调用一次少一次
    非认证，基于ip

- 在响应中，和限流有关的有3个属性：
  - X-RateLimit-Limit 每小时最大请求数
  - X-RateLimit-Remaining 当前环境 还能请求多少次
  - X-RateLimit-Reset 重置的时间，UTC秒数(在浏览器控制台输入new Date(1000\*秒数)可查看具体时间)
- 如果超过了限制，会返回403

### OAuth 程序未认证限流增加上限

- 60次/小时如果不够用，可带上id和secret

    curl -i 'https://api.github.com/users/whatever?client_id=xxxx&client_secret=yyyy'
    HTTP/1.1 200 OK
    Date: Mon, 01 Jul 2013 17:27:06 GMT
    Status: 200 OK
    X-RateLimit-Limit: 5000
    X-RateLimit-Remaining: 4966
    X-RateLimit-Reset: 1372700873

### 保持不要超过限流的限制

- 如果确定会超过限流，可以缓存api响应，并使用有条件请求
- 不要滥用资源，不然超过限流会报403

## 用户代理

- 所有的api请求都需要包含一个有效的User-Agent头
- 没有这个User-Agent头的请求会被拒绝
- http中User-Agent用于表示用户使用的什么系统/什么浏览器/什么自定义消息
- github建议，这个User-Agent要么使用github用户名，要么使用程序名
- curl默认会添加一个User-Agent,github会识别这个，并认为是一个有效的
- 如果通过curl并设置了一个无效的User-Agent，响应会返回403

## 条件请求

- 超过限流，可使用条件请求
- 很多响应会返回一个ETag头
- 同时还会返回一个Last-Modified头
- 可以通过这些头信息，创建一个子请求(会用到If-None-Match和If-Modified-Since) 
- 如果这些资源并未改变，会返回304 Not Modified
- 使用条件请求，如果返回的是304,就不会触发限流计数
- 所以推荐任何时候都使用条件请求

    // 先来一个请求
    curl -i https://api.github.com/users/63isOK
    HTTP/1.1 200 OK
    Date: Sun, 29 Sep 2019 06:36:48 GMT
    Content-Type: application/json; charset=utf-8
    Content-Length: 1294
    Server: GitHub.com
    Status: 200 OK
    X-RateLimit-Limit: 60
    X-RateLimit-Remaining: 59
    X-RateLimit-Reset: 1569742607
    Cache-Control: public, max-age=60, s-maxage=60
    Vary: Accept
    ETag: "963c0790703410579a9b4ff7db92f822"
    Last-Modified: Thu, 26 Sep 2019 09:25:55 GMT
    X-GitHub-Media-Type: github.v3; format=json
    Access-Control-Expose-Headers: ETag, Link, Location, Retry-After, X-GitHub-OTP, X-RateLimit-Limit, X-RateLimit-Remaining, X-RateLimit-Reset, X-OAuth-Scopes, X-Accepted-OAuth-Scopes, X-Poll-Interval, X-GitHub-Media-Type
    Access-Control-Allow-Origin: *
    Strict-Transport-Security: max-age=31536000; includeSubdomains; preload
    X-Frame-Options: deny
    X-Content-Type-Options: nosniff
    X-XSS-Protection: 1; mode=block
    Referrer-Policy: origin-when-cross-origin, strict-origin-when-cross-origin
    Content-Security-Policy: default-src 'none'
    Vary: Accept-Encoding
    X-GitHub-Request-Id: C576:5C94:2A8D9B:37D276:5D9050FF

    // 条件请求，参数使用ETag
    // 注意写法，Etag是一个字符串，需要用单引号包裹
    curl -i https://api.github.com/users/63isOK -H 'If-None-Match: "963c0790703410579a9b4ff7db92f822"'
    HTTP/1.1 304 Not Modified
    Date: Sun, 29 Sep 2019 06:38:28 GMT
    Server: GitHub.com
    Status: 304 Not Modified
    X-RateLimit-Limit: 60
    X-RateLimit-Remaining: 59
    X-RateLimit-Reset: 1569742708
    Cache-Control: public, max-age=60, s-maxage=60
    Vary: Accept
    ETag: "963c0790703410579a9b4ff7db92f822"
    Last-Modified: Thu, 26 Sep 2019 09:25:55 GMT
    Access-Control-Expose-Headers: ETag, Link, Location, Retry-After, X-GitHub-OTP, X-RateLimit-Limit, X-RateLimit-Remaining, X-RateLimit-Reset, X-OAuth-Scopes, X-Accepted-OAuth-Scopes, X-Poll-Interval, X-GitHub-Media-Type
    Access-Control-Allow-Origin: *
    Strict-Transport-Security: max-age=31536000; includeSubdomains; preload
    X-Frame-Options: deny
    X-Content-Type-Options: nosniff
    X-XSS-Protection: 1; mode=block
    Referrer-Policy: origin-when-cross-origin, strict-origin-when-cross-origin
    Content-Security-Policy: default-src 'none'
    Vary: Accept-Encoding
    X-GitHub-Request-Id: B947:0BE3:642785:7F23EA:5D905164

    // 条件请求，参数使用Last-Modified
    // 注意写法，Last-Modified 不是字符串，不需要单独用引号包裹
    curl -i https://api.github.com/users/63isOK -H "If-Modified-Since: Thu, 26 Sep 2019 09:25:55 GMT"  
    HTTP/1.1 304 Not Modified
    Date: Sun, 29 Sep 2019 06:41:51 GMT
    Server: GitHub.com
    Status: 304 Not Modified
    X-RateLimit-Limit: 60
    X-RateLimit-Remaining: 59
    X-RateLimit-Reset: 1569742910
    Cache-Control: public, max-age=60, s-maxage=60
    Vary: Accept
    ETag: "963c0790703410579a9b4ff7db92f822"
    Last-Modified: Thu, 26 Sep 2019 09:25:55 GMT
    Access-Control-Expose-Headers: ETag, Link, Location, Retry-After, X-GitHub-OTP, X-RateLimit-Limit, X-RateLimit-Remaining, X-RateLimit-Reset, X-OAuth-Scopes, X-Accepted-OAuth-Scopes, X-Poll-Interval, X-GitHub-Media-Type
    Access-Control-Allow-Origin: *
    Strict-Transport-Security: max-age=31536000; includeSubdomains; preload
    X-Frame-Options: deny
    X-Content-Type-Options: nosniff
    X-XSS-Protection: 1; mode=block
    Referrer-Policy: origin-when-cross-origin, strict-origin-when-cross-origin
    Content-Security-Policy: default-src 'none'
    Vary: Accept-Encoding
    X-GitHub-Request-Id: CEF6:64FB:5E7A24:7B005B:5D90522E

- 使用条件请求，后面两次请求并没有触发限流计数

## CORS

- CORS (Cross origin resource sharing), 跨域资源共享，是一种机制
- 利用这种机制，ajax请求可以请求任何源的资源

## json-p 回调

- 使用?callback参数，可以在GET指定json函数
- 一般用在将gtihub内容嵌入一个web页面，用于解决跨域问题

## 时区

- 部分api请求创建新数据时，需要指明时区，用于生成时间戳
- 下面按优先级排了几个规则：
  - 显式使用ISO 8601时间戳
  - 使用Time-Zone头
  - 使用最后一个已知的时区
  - 默认使用UTC，这时不需要时区信息

