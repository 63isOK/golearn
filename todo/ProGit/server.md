# 服务器上的git

使用git协作,要么搭建自己的服务器,要么使用托管服务.

现在来了解一下服务器上git的部分概念:
- git使用4种协议来传输资料
    - 本地协议
        - 远程仓库就是硬盘上的一个目录
        - 适用于协作者可以访问一个共享文件系统,远程仓库就放在这个这个文件系统
        - 优点是简单
        - 缺点是共享文件系统比较难配置,速度不占优势,权限也无法很好配置
    - http协议
        - 主流协议
        - 相比其他协议,唯一的缺点的https的部署复杂很多
        - 其他方面,其他协议都不如http协议
    - ssh协议
        - 优点:很多服务器上都有ssh
        - 缺点:不能匿名访问(即使仅仅是读)
    - git协议
        - 优点:项目很大时,git协议的传输速度是最快的
        - 缺点:缺乏授权机制,要么谁都可以clone,要么谁都不可以

自己搭建git服务器,按不同协议,一步步搭建即可,
自己不想搭建git服务器,就是用第三方托管服务,github是一个很好的选择


