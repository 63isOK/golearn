# github 其他认证方法

上一节谈到了OAuth2.0的认证方式，包括了和认证码相关/和token相关的，
这里主要讨论一下其他认证的方法

虽然github提供了除开OAuth的其他认证方式，github还是强烈推荐使用OAuth认证，
其他认证方式的提供，更多的是为了脚本化/测试

## 基础认证

- 基础认证是基于rfc2617
- github实现和rfc也稍有不同：
  - 遇到未认证请求时，rfc要求返回401 Unauthorized responses
  - github返回404 Not Found，为的是不泄漏用户信息
  - 遇到这种情况，如果使用http库，可能会出错，可以手动添加头信息来解决

### 用户密码来访问账户

- 最简单的，也是不安全的，

    curl -u release4go https://api.github.com/user

### 通过OAuth token来访问账户

- 可以使用OAuth token，或是私人的访问token来访问账户
- 适合：所用的工具只支持基础认证，但又想具备OAuth安全访问的特性，就适合用这种

    curl -u release4go:token https://api.github.com/user

### SAML SSO 单点认证

- 对于组织org来说，可以用api进行SAML SSO单点认证
- 此时需要创建一个PAT(personal access token 私人访问toekn)和token白名单
- 也就是说通过api访问组织时，请求头里的X-GitHub-SSO对应的url里放的是token白名单
- 没有token白名单，一个用户是无法直接取访问这些信息的
- 如果api涉及多个组织，那X-GitHub-SSO里放的就是组织号了(github会根据组织号自动去找白名单)
  - 多个组织用逗号分割

## 二次认证

- 启用二次认证后，大多数endpoint的基础认证(v3)需要的是私人访问token或OAuth token，不需要用户名和密码
- 私人访问token可以在github页面创建
- OAuth token用api生成
- 用户名和密码只在OAuth 认证api中会使用到
- 开启二次认证后，基础认证需要用户名和密码，也需要一个otp(一次性密码)

开启二次认证后，会需要一个一次性的密码，类似短信验证码，认证也会变得复杂

```shell
curl --request POST \
  --url https://api.github.com/authorizations \
  --header 'authorization: Basic PASSWORD' \
  --header 'content-type: application/json' \
  --header 'x-github-otp: OTP' \
  --data '{"scopes": ["public_repo"], "note": "test"}'
```
