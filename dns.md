# dns

域名解析系统。负责ip和域名之间的相互解析

dns中有多种记录类型：

## a记录

address，记录的是域名/主机名对应的ip

最简单的记录，输入域名，dns告诉她对应的ip

## ns记录

name server，记录的是域名服务器

做中转用的，输入域名，告诉她这个域名在哪个dns服务器进行解析

每一个域名注册时都会告诉对应的dns服务器，这个dns就可以解析这个域名，
(注册时用的是a记录，其他dns可以用ns记录，这样达到中转的目的)

## mx记录

mail exchanger，记录的是邮件交换

dns中存的是mail.qq.com对应的ip，如果发的是qq邮件，就去找对应的ip，
一般用于邮件

## cname记录

canonical name，记录别名

一般计算机有一个a记录，在这个域下提供了两个服务：www和mail，
就可以使用别名，将多个名字映射到同一台计算机，www.my.com和mail.my.com,
实际上都是指向my.com,也就是a记录中的那个ip

简单点说，cname记录表示的是域名指向另一个域名。

    www.baidu.com.		377	IN	CNAME	www.a.shifen.com.
    www.a.shifen.com.	188	IN	A	39.156.66.18
    www.a.shifen.com.	188	IN	A	39.156.66.14


## ptr记录

pointer，记录的是ip到域名/主机名的映射，和a记录是逆向的

一般用在邮件服务器上，和mx记录不一样，
mx记录主要用于发邮件时，根据邮件地址后面的域名找到对应的ip；
ptr记录主要用在收邮件时，通过发件方的ip找到对应的域名，
再和发件方邮件地址后的域名比较，判断是不是伪造邮件。

# dig

linux下的dig命令，用于从dns域名服务器查询信息，查询主机地址信息。

    查单个主机信息： dig www.google.com
    指定记录类型：   dig www.baidu.com cname
    从指定dns服务器上查询 dig @8.8.8.8 www.baidu.com
    反向查域名       dig -x 127.0.0.53     +short只显示域名
