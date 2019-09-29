# 媒体类型

- 可以理解为资源类型
- 调用api请求时，在Accept头中可以添加可接收资源的类型
- github的资源类型写法如下：
  - application/vnd.github[.version].param[+json]
- github中支持的基础类型大部分如下：
  - application/json
  - application/vnd.github+json
  - 这里是没有指定版本的，默认取当前json指定的资源
- 目前使用的是v3版本的api，未来会升级，所以写法最好指定版本
  - application/vnd.github.v3+json
- 版本后买能是可以指定一些属性的
  - 属性包含full/raw/etc
  - application/vnd.github.v3.raw+json

    curl -I https://api.github.com/users/63isOK                                                      
    // -I 表示只显示响应头
    // -i 表表示显示响应头和内容

    HTTP/1.1 200 OK
    Date: Sun, 29 Sep 2019 07:34:38 GMT
    Content-Type: application/json; charset=utf-8
    Content-Length: 1294
    Server: GitHub.com
    Status: 200 OK
    X-RateLimit-Limit: 60
    X-RateLimit-Remaining: 59
    X-RateLimit-Reset: 1569746078
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
    X-GitHub-Request-Id: A98E:71B4:5EB36A:7BFB04:5D905E8D

- 看上面的响应头，X-GitHub-Media-Type: github.v3; format=json 可以看到当前版本是v3

## comment 的属性

- comment可以用在很多地方，一般使用md写法：
  - issues
  - issues的comment
  - pr comment
  - gist comment
- application/vnd.github.VERSION.raw+json 返回原始的md
  - 这也是默认的资源类型
  - 原始的md在响应的body字段下
- application/vnd.github.VERSION.text+json 返回文本
  - 文本里放的就是md
  - 在响应的body_text字段下
- application/vnd.github.VERSION.html+json 返回md生成的html
  - 在响应的body_html字段下
- application/vnd.github.VERSION.full+json 返回raw / text /html
  - 响应中会包含 body/body_text/body_html

## git blob 的属性

- application/vnd.github.VERSION+json
- application/json
  - 这两个都是json格式
  - 就是说blob的内容进行base64编码之后，放在json里
  - 这是默认的
- application/vnd.github.VERSION.raw
  - 响应中会返回原始的blob数据

## commit,commit比较，pr

- application/vnd.github.VERSION.diff 
- application/vnd.github.VERSION.patch
- application/vnd.github.VERSION.sha

## 仓库内容

- application/vnd.github.VERSION.raw 
  - 这个是默认的
  - 返回一个文件的原始内容
- application/vnd.github.VERSION.html
  - 对于标记语言md或asciidoc，可以返回html

## gist

- application/vnd.github.VERSION.raw
  - 这是默认的
  - 返回gist文件的原始内容
- application/vnd.github.VERSION.base64
  - 返回经过base64编码之后的内容
  - 如果gist里包含了非utf-8的字符，用base64就非常有用了

## 总结

- 这节主要介绍了不同资源有不同的返回类型
- 可在合适的时候选择合适的类型
