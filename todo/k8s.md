# k8s社区流程

这个流程(交流规范)的目的是：提高文档质量 提高代码质量 提供交流机会

## 项目治理

主要包括了项目的结构和项目中的组织,记录的是目前项目运行中的工作方式

### 原则

* 开源，公开
* 欢迎贡献者，尊重社区规范等行为守则
* 透明，可访问
* 发挥长处，想法和贡献符合技术长处，并和项目目标、规模、设计原则保持一致

### 行为守则

符合k8s的行为守则，说白了，就是遵循cncf行为守则。cncf(云原生计算基金会)。
简单点就是体面一点，不管是贡献者还是维护者，不要做出情感色彩明显的举动

### 社区价值

k8s社区文化常被认为是k8s迅速崛起过程中的一个重要贡献者，
下面是k8s社区多年来总结出来的精华：
* 分布式比集中式好。项目规模的伸缩是基于高信任高可见来分配工作，
包括授权、决策、技术设计、代码所有权和文档，
分散式非同步的所有权和协作、沟通、决策是全社区的基石
* 社区先于产品和公司。作为社区一员，主动管理项目，
以达到为"项目成员和项目的使用者"造福的目的。
个人通过工作(贡献或维护项目)来获得社区地位；公司通过支持和提供资源获得社区地位
 
### 社区关系

区分贡献者在项目中的角色和职责，一个大项目会细分成多个子项目，
这里的贡献者角色都是指子项目中的。

| 角色 |责任 | 要求 | 任命 |
| ---- |---- | ---- | ---- |
| 成员 | 项目的积极贡献者 | 有两位审查员负责，对项目要有多次贡献 | 有意愿参与的 |
| 审查员 | 审查其他成员的的贡献 | 有审查经验，是子项目的作者 | OWNERS 中定义 |
| 批准员 | 批准接受贡献 | 经验丰富，积极的审查者，子项目的贡献者 | OWNERS 中定义 |
| 子项目所有者 | 制定子项目的方向和优先级 | 负责子项目，对子项目有优秀的技术判断力 | 子项目中的OWNERS中定义 |

有项目新成员时：
* 欢迎
* 熟悉RP(request pull)流程
* 项目相关文档和沟通渠道

新进入的成员要认同社区文化，包括认同 项目组织 角色 政策 手续 惯例等，
要有一定的技术能力或是写作能力。  
针对上面的角色或其他定义，下面有详细讲解。

#### 成员

sig，Special Interest Groups 特殊兴趣小组

积极的贡献者，可以提出issues，也可推送和问题相关的rp。
作为sig的一员，在预提交rp之前，要提供rp相关的自动化测试

成员要求：
* 开启github的两步认证(固定密码+随机的6位数字)
* 多次贡献，包括并不局限于:
    * 以作者或审查者的角色推送rp
    * 提交或评论issues
    * 为子项目，sig，或社区讨论 做贡献 (meeting, 邮件，stack overflow)等
* 订阅项目开发讨论组，如果有类似google group的话
* 阅读贡献规则
* 积极为1个或多个子项目做贡献
* 由两位审查者赞助，审查者要满足以下条件
    * 这个目前还不理解 后面再补 TODO
* 打开一个issues
    * 确保赞助商被提及到
    * checklist上所有的项都已完成
    * 列表中包含的贡献表示相应项目中的的工作
* 让赞助审查者确认赞助关系
* 一旦审查者响应了，社区(项目中的某个team)会进行审查，依据slo规则。如果遗漏了信息，会被要求补全

生态：k8s github组织成员，自动是相关组织的成员；反之不是，需要申请k8s github组织成员

责任和特权：
* 对issues负责，并提供响应的rp
* 如果sig的成员提及到自己，需要响应，(可以理解为微信上的@)
* 现有代码的所有者(如果所有权没有明确转让)，所有者要是积极的
    * 代码要是易于测试的
    * 测试要能通过的
    * 在接收代码后，如果发现bug和issuse，需要定位
* 在打开的rp上，有权利标记 /lgtm(look good to me),一般在审查后觉得没问题打的标识
* 可以被分配到指定的issues和rp，需要某人来审查代码，可以使用 /cc @用户名
* rp可以自动化测试， 不需要指定 /ok-to-test
* 可以为一个带 needs-ok-to-test 的rp 打一个 /ok-to-test 标记；或是打一个 /close 去关闭一个rp

成员需要主动推进审查工作，进一步，要成为子项目的主要审查员。

### 审查员





