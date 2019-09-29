# OAuth 认证

在某些OAuth认证api的响应中，不用返回token属性

下面主要介绍OAuth认证的api，
- 可以用这些api来管理访问账户的OAuth程序
- 也可以使用基础认证(用户名和密码)来访问这些api，不需要token

## 显示所有的授权程序

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

## 获取单个程序的授权

- GET /applications/grants/:grant_id
- 这里的grant_id是指授权的应用程序id，对应到上面就是 104777119

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

## 删除某个授权的应用程序

- DELETE /applications/grants/:grant_id
- 删除之后，这个应用程序就不能再访问我的账户了

    curl -X DELETE -i -u release4go:zhuwenk https://api.github.com/applications/grants/123 

    成功会返回：
    Status: 204 No Content

## 显示所有授权

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

## 显示某一个授权

- GET /authorizations/:authorization_id

    curl  -i -u release4go:zhuwenkai001 https://api.github.com/authorizations/331569766
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

## 创建一个新的授权

- 创建一个OAuth
