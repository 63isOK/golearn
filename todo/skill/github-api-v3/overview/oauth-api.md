# OAuth 认证

基于[OAuth2.0](https://tools.ietf.org/html/rfc6749)认证框架

在某些OAuth认证api的响应中，不用返回token属性

下面主要介绍OAuth认证的api，
- 可以用这些api来管理访问账户的OAuth程序
- 也可以使用基础认证(用户名和密码)来访问这些api，不需要token

下面主要介绍两类api
- grants api
  - 主要针对应用程序的授权
  - 这个授权之后，是不需要管理独立的token的
  - 按照OAuth2.0的标准看，github oauth用的是授权码模式
  - 这个grants api主要是用于处理授权码的
- authorizations api
  - 这个主要处理授权码模式中的token
  - 一个授权码可以对应多个token

## 显示所有有授权码的程序 grants

- GET /applications/grants
- 获取所有通过OAuth认证访问账户的程序

    curl -i -u release4go:zhuwenk https://api.github.com/applications/grants
    // 下面只有一个符合的程序： travis-ci

    HTTP/1.1 200 OK
    Date: Sun, 29 Sep 2019 08:44:32 GMT
    Content-Type: application/json; charset=utf-8
    Content-Length: 457
    Server: GitHub.com
    Status: 200 OK
    X-RateLimit-Limit: 5000
    X-RateLimit-Remaining: 4999
    X-RateLimit-Reset: 1569750272
    Cache-Control: private, max-age=60, s-maxage=60
    Vary: Accept, Authorization, Cookie, X-GitHub-OTP
    ETag: "78a5dbe5539d0c19a810fb2d9e30adf9"
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
    X-GitHub-Request-Id: FDE5:71B4:62D36D:8160A6:5D906EEF

    [
      {
        "id": 104777119,
        "url": "https://api.github.com/applications/grants/104777119",
        "app": {
          "name": "Travis CI for Open Source",
          "url": "https://travis-ci.org",
          "client_id": "f244293c729d5066cf27"
        },
        "created_at": "2019-08-20T03:27:46Z",
        "updated_at": "2019-08-20T03:27:46Z",
        "scopes": [
          "read:org",
          "repo:status",
          "repo_deployment",
          "user:email",
          "write:repo_hook"
        ]
      }
    ]

- 响应中的scopes是授权的范围

## 通过授权码查对应的程序

- GET /applications/grants/:grant_id
- 这里的grant_id是指授权的应用程序对应的授权程序id(授权码)，对应到上面就是 104777119

    curl -i -u release4go:zhuwenk https://api.github.com/applications/grants/104777119

    HTTP/1.1 200 OK
    Date: Sun, 29 Sep 2019 09:04:58 GMT
    Content-Type: application/json; charset=utf-8
    Content-Length: 417
    Server: GitHub.com
    Status: 200 OK
    X-RateLimit-Limit: 5000
    X-RateLimit-Remaining: 4999
    X-RateLimit-Reset: 1569751498
    Cache-Control: private, max-age=60, s-maxage=60
    Vary: Accept, Authorization, Cookie, X-GitHub-OTP
    ETag: "80136b9f88c236ab3b3051655fdeff07"
    Last-Modified: Tue, 20 Aug 2019 03:27:46 GMT
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
    X-GitHub-Request-Id: EAC2:666C:66B906:869809:5D9073B9

    {
      "id": 104777119,
      "url": "https://api.github.com/applications/grants/104777119",
      "app": {
        "name": "Travis CI for Open Source",
        "url": "https://travis-ci.org",
        "client_id": "f244293c729d5066cf27"
      },
      "created_at": "2019-08-20T03:27:46Z",
      "updated_at": "2019-08-20T03:27:46Z",
      "scopes": [
        "read:org",
        "repo:status",
        "repo_deployment",
        "user:email",
        "write:repo_hook"
      ]
    }

## 通过授权码删除相关的程序

- DELETE /applications/grants/:grant_id
- 删除之后，这个应用程序就不能再访问我的账户了

    curl -X DELETE -i -u release4go:zhuwenk https://api.github.com/applications/grants/123 

    成功会返回：
    Status: 204 No Content

## 查找所有授权,这个授权会管理token

- 显示所有授权, 此时更多的是为了拿到token
- GET /authorizations 

    curl  -i -u release4go:zhuwenk https://api.github.com/authorizations

    HTTP/1.1 200 OK
    Date: Sun, 29 Sep 2019 09:11:33 GMT
    Content-Type: application/json; charset=utf-8
    Content-Length: 1759
    Server: GitHub.com
    Status: 200 OK
    X-RateLimit-Limit: 5000
    X-RateLimit-Remaining: 4999
    X-RateLimit-Reset: 1569751893
    Cache-Control: private, max-age=60, s-maxage=60
    Vary: Accept, Authorization, Cookie, X-GitHub-OTP
    ETag: "0daec668ba6a9a853723744ff0141295"
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
    X-GitHub-Request-Id: 89C6:71B4:645D42:8363E8:5D907545

    [
      {
        "id": 319734067,
        "url": "https://api.github.com/authorizations/319734067",
        "app": {
          "name": "Travis CI for Open Source",
          "url": "https://travis-ci.org",
          "client_id": "f244293c729d5066cf27"
        },
        "token": "",
        "hashed_token": "386ef32db57434730937be344bb5bc1b6ba64c7a3ccd6fe97aa0a66a29b6fc4d",
        "token_last_eight": "186cf8d4",
        "note": null,
        "note_url": null,
        "created_at": "2019-08-20T03:27:46Z",
        "updated_at": "2019-08-20T03:27:46Z",
        "scopes": [
          "read:org",
          "repo:status",
          "repo_deployment",
          "user:email",
          "write:repo_hook"
        ],
        "fingerprint": null
      },
      {
        "id": 333752504,
        "url": "https://api.github.com/authorizations/333752504",
        "app": {
          "name": "GitHub Pages",
          "url": "http://pages.github.com",
          "client_id": "24df84fbaf82e89de25c"
        },
        "token": "",
        "hashed_token": "fc25bf7ce6f535d975a89a7af67acda21b12e63cf002c397a124a50c96fc1ee2",
        "token_last_eight": "58cd9f4b",
        "note": null,
        "note_url": null,
        "created_at": "2019-09-28T13:23:07Z",
        "updated_at": "2019-09-28T13:23:07Z",
        "scopes": [
          "repo"
        ],
        "fingerprint": null
      },
      {
        "id": 331569766,
        "url": "https://api.github.com/authorizations/331569766",
        "app": {
          "name": "GitHub Launch",
          "url": "https://github.com/github",
          "client_id": "9308c497974c96a06e71"
        },
        "token": "",
        "hashed_token": "e0e09e5f7611bb84f6b7519cfc8830b6b24a6bce6b2a05fe6fd8b8e0f7674f9d",
        "token_last_eight": "2dc53943",
        "note": null,
        "note_url": null,
        "created_at": "2019-09-23T08:39:38Z",
        "updated_at": "2019-09-23T08:39:38Z",
        "scopes": [
          "repo"
        ],
        "fingerprint": null
      }
    ]

- 可以看出除了travis-ci，还有github page和github launch都有授权
- 授权id和应用程序id是不一样的

## 通过一个授权id显示授权信息

- GET /authorizations/:authorization_id

    curl  -i -u release4go:zhuwenk https://api.github.com/authorizations/331569766
    更具授权id来查看具体的授权信息

    HTTP/1.1 200 OK
    Date: Sun, 29 Sep 2019 09:19:39 GMT
    Content-Type: application/json; charset=utf-8
    Content-Length: 511
    Server: GitHub.com
    Status: 200 OK
    X-RateLimit-Limit: 5000
    X-RateLimit-Remaining: 4999
    X-RateLimit-Reset: 1569752379
    Cache-Control: private, max-age=60, s-maxage=60
    Vary: Accept, Authorization, Cookie, X-GitHub-OTP
    ETag: "df5857b3f2b05fe885f6f36f7f062d9f"
    Last-Modified: Mon, 23 Sep 2019 08:39:38 GMT
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
    X-GitHub-Request-Id: B2D0:7AAF:2F59C9:3E8123:5D90772B

    {
      "id": 331569766,
      "url": "https://api.github.com/authorizations/331569766",
      "app": {
        "name": "GitHub Launch",
        "url": "https://github.com/github",
        "client_id": "9308c497974c96a06e71"
      },
      "token": "",
      "hashed_token": "e0e09e5f7611bb84f6b7519cfc8830b6b24a6bce6b2a05fe6fd8b8e0f7674f9d",
      "token_last_eight": "2dc53943",
      "note": null,
      "note_url": null,
      "created_at": "2019-09-23T08:39:38Z",
      "updated_at": "2019-09-23T08:39:38Z",
      "scopes": [
        "repo"
      ],
      "fingerprint": null
    }

## 创建一个新的授权(手工创建一个token/或者通过ouath创建token)

- 创建一个OAuth token,可以利用基础认证的方式获取
- 如果使用2次认证的方式，可以利用otp(一次性密码) + 用户密码 来代替token(这种方式不需要token)
- 给一个具体的OAuth程序创建一个token，使用endpoint即可
  - 此时你要创建授权，就需要对这个创建授权操作的人进行验证
  - 并提供程序的client-id和client-secret(这两个参数在OAuth程序的设置界面获取)
- 如果一个OAuth程序要给一个用户创建多个token，需要使用fingerprint参数来区别这些token
- 当然，如果不走github的OAuth2.0认证去创建token，也可以自己手动在github上创建一个token
  - 此时需要手动将一些权限和token绑定在一起
  - 在github 个人设置/开发设置/个人访问token 创建token
  - 创建授权有两种方式：手动创建和通过OAuth程序走OAuth2.0的授权码模式去创建
- POST /authorizations

参数如下：

参数名|类型|描述
--|--|--
scopes|数组|包含了授权范围
note|字符串|必选，可以理解为授权名，如果是手动创建的授权，名字还要唯一
note_url|字符串|和授权名关联的一个url，和note的作用一样，都是为了提示
client_id|字符串|20个字符，用于生成token
client_secret|字符串|40个字符，用于生成token
fingerprint|字符串|用于区别同一个用户或同一个client_id创建的不同授权

    curl -X POST -i -u release4go:zhuwenk https://api.github.com/authorizations -d '{"scopes":["public_repo"],"note":"admin script"}' -H "Content-Type: application/json; charset=utf-8"
    此处使用到了post，使用了json传参
    功能只是手工创建了一个授权
    返回201表示创建成功

    HTTP/1.1 201 Created
    Date: Mon, 30 Sep 2019 03:11:54 GMT
    Content-Type: application/json; charset=utf-8
    Content-Length: 595
    Server: GitHub.com
    Status: 201 Created
    X-RateLimit-Limit: 5000
    X-RateLimit-Remaining: 4998
    X-RateLimit-Reset: 1569816685
    Cache-Control: private, max-age=60, s-maxage=60
    Vary: Accept, Authorization, Cookie, X-GitHub-OTP
    ETag: "ea3a183c2e968ef691724871830fd592"
    Location: https://api.github.com/authorizations/334105717
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
    X-GitHub-Request-Id: 8442:2F91:486418:5E448D:5D91727A

    {
      "id": 334105717,
      "url": "https://api.github.com/authorizations/334105717",
      "app": {
        "name": "admin script",
        "url": "https://developer.github.com/v3/oauth_authorizations/",
        "client_id": "00000000000000000000"
      },
      "token": "55b2bc7d744dc01067c389c65fbe4bbca5fb9357",
      "hashed_token": "51c4d57c261d311e9aed0255173d8f2de29930afc3fcf487be2e1f101760340a",
      "token_last_eight": "a5fb9357",
      "note": "admin script",
      "note_url": null,
      "created_at": "2019-09-30T03:11:54Z",
      "updated_at": "2019-09-30T03:11:54Z",
      "scopes": [
        "public_repo"
      ],
      "fingerprint": null
    }

## 对于某个app，获取或创建一个授权

- 本小节主题，对于用户来说，某个app没有某些授权，那么就为app创建一个新的授权
- url里面需要包含client-id，github用这个client-id来区分应用程序
- PUT /authorizations/clients/:client_id
  - client-id放在url中
  - json参数还会包含client-secret
  - 参数还包含了scopes/note/note_url/fingerprint
- 这个api，相当于：获取OAuth程序获取授权，如果授权不存在就新建授权
- 如果授权已存在，就返回200,如果新建，就返回201


## 对于某个app的某个授权，获取或创建一个授权

- 一个app可能有多个授权，授权用fingerprint来区分
- 相对于上一个api来说，这个api在app限制下，还添加了fingerprint限制
- PUT /authorizations/clients/:client_id/:fingerprint
- 相对于上一个api来说，api调用方面，fingerprint从json参数，放到了url中，表示的api也不一样


## 更新一个已存在的授权

- PATCH /authorizations/:authorization_id
- json参数如下：
  - scopes 数组，新的授权指定的范围
  - add_scopes, 数组，新增的授权范围
  - remove-scopes,数组，要删除的授权范围
  - note, 手动创建的授权，需要带上这个
  - note_url
  - fingerprint
- 参数中的scopes相关的3个参数，不必全写，可以省略部分
- 返回200 表示成功

## 删除一个授权

- DELETE /authorizations/:authorization_id
- 成功，返回204 No Content

## 授权的检查

- OAuth程序可以通过指定api来检查授权，而不影响限流
- 此时检查的是检查授权对应的token
- 授权一般配合其他endpoint一起工作
- 我们可以通过基础认证取访问授权，也可以通过OAUth程序(通过client-id/client-secret访问)
- GET /applications/:client_id/tokens/:access_token
- 这个api就是给OAuth程序来使用的
- 成功会返回200

## 重置一个授权

- 也是给OAuth程序使用的api
- 执行完这个api之后，需要立马更新token(在响应中可以获取这个token)，因为是立马生效的
- POST /applications/:client_id/tokens/:access_token
- 可以把这个api理解为：重新获取token

## 撤销某个程序的授权(删除程序中指定的某一项授权)

- OAuth2.0协议中支持的一种场景
- 使用的是基础认证默认 curl -u user:pwd url
  - 不过不是用户的用户和密码，而是程序的client-id:client-secret
- DELETE /applications/:client_id/tokens/:access_token
- 成功返回204 No Content

## 撤销某个应用的授权码(删除程序的整个授权)

- OAuth2.0协议中支持的一种场景
- 使用的是基础认证默认 curl -u user:pwd url
  - 不过不是用户的用户和密码，而是程序的client-id:client-secret
- DELETE /applications/:client_id/grants/:access_token
- 删除之后，会删除所有有关联的token，且在github页面也不会再显示这个应用程序
