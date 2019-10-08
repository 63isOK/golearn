# 星星

- 星星表示了一个仓库的收欢迎程度
- 星星和通知/时间线没什么影响

## 查看了哪些观星者

- GET /repos/:owner/:repo/stargazers

## 响应中带打星的时间戳

- 请求头上指定资源类型： Accept: application/vnd.github.v3.star+json

## 显示当前用户打星的仓库

- GET /users/:username/starred
- 显示认证用户的打星仓库 GET /user/starred
- 参数：
  - sort： created 按打星时间排序;updated按仓库更新时间排序，默认updated
  - direction: asc 升序; desc 降序

## 检查用户是否给指定仓库打过星

- GET /user/starred/:owner/:repo
- 需认证
- 204 No Content 表示已打星
- 404 Not Found 表示未打星

## 给指定仓库打星

- PUT /user/starred/:owner/:repo
- 需认证
- 正确响应204 No Content

## 取消打星

- DELETE /user/starred/:owner/:repo
