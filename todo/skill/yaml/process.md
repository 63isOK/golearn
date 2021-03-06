# 处理yaml信息

本节主语介绍yaml的组成部分

- yaml既是一个文本格式，也是一个方法(将原生数据格式转换成文本格式的方法)
- 所以yaml spec中对yaml定义了两个概念
  - 数据对象的类，用于yaml的表现(tagged nodes有根/连接/有向图)
  - 一种语法(将yaml的表现作为一系列字符串呈现出来)，叫yaml流
- yaml处理器就是一套工具，将yaml表现和yaml流之间做转换
- 使用yaml处理器的其他模块叫程序

本节主要围绕yaml信息结构，这些信息由yaml处理器或程序提供

yaml信息要么给机器用，要么给人用
- 所以yaml信息有3种状态：
  - representation 表现，是一个由root节点/连接线/有向图组成
    - 用于解决 原生数据结构和编程环境之间的移植转换
  - serialization 序列化树，
    - 关注于将representation转换成一个串行格式(这个格式需要顺序访问)
  - presentation 流，人类可读
    - 将serialization处理为一系列人类可读的字符

## 处理

- 原生数据结构和字符流之间的转换要经过几个逻辑阶段
- 每个阶段都定义好了输入输出的数据模型

处理过程描述：
- yaml信息主要有4个不同的状态：
  - 原生数据结构 native data structure， 对应编程环境中的数据结构
  - 节点图 node graph，对应representation
  - 事件树 event tree，对应serialization
  - 字符流 character stream，对应presentation
- 4个状态中有3种转换
  - 数据结构 --represent--> 节点图 --serialize--> 事件树 --present--> 字符流
  - 数据结构 <--construct-- 节点图 <--compose--   事件树 <--parse--   字符流
- yaml处理只需要对外暴露字符流阶段即可
  - 从数据结构到字符流，叫dump
  - 反过来，从字符流到数据结构，叫load

dump的三个阶段：
- 第一阶段，转成节点图
  - 利用3种节点(node)来映射数据结构
    - sequence，序列，可理解为数组
    - map，kv对
    - 字面量
  - 这些原语组合成有向图结构
    - sequence对应很多编程语言中的数据和列表
    - map对应hash表，字典
    - 字面量对应 string int date 以及其他数据类型
  - 节点node，包含(类型/内容/用于指明数据类型的tag)
- 第二阶段，序列化成事件树
  - 对于要顺序访问的媒介(元素)，需要序列化成一个有序树
  - map是无序的，一个node可能被引用多次
    - 所以序列化操作对key要强加一个排序
    - 对于第node的第二次及后面的引用，都用一个别名占位符来代替
  - yaml sepc对具体的序列化细节并未规定
  - 序列化完之后就是序列化的树了，这个树叫事件树
- 第三阶段，转字符流
  - yaml提供了多种可选风格，来让字符流更具有可读性
  - 所以在这个阶段，yaml处理器就需要更多信息：
    - 选择哪种node风格
    - 字面量格式
    - 缩进
    - 使用哪些tag
    - 未指定的保留tag
    - 提供哪些指令集
    - 添加注释时可能的事件
  - 有时，上面部分信息由程序决定，有时，需要让用户按自己的喜好来选择

load是dump的逆向，将可读性的字符流转换成程序的额数据结构：
- 第一阶段，解析
  - 将字符流解析为一系列事件
  - 解析时不用关心那么多可选风格，只需要将当前风格的数据解析为序列化事件即可
  - 解析可能会失败(eg：错误的格式会导致失败)
- 第二阶段，组合成图
  - 组合成节点图也可能会失败，具体失败原因会在后续展开
- 第三阶段，构造数据结构
  - 构造时会丢弃很多附加信息：
    - 注释
    - 指令
    - map的有序key
    - node风格
    - 字面量内容格式
    - 缩进等级
  - 构造阶段也可能会失败(eg：找不到语言对应的数据结构)

## 信息模型

- 要最大化地将编程语言中的数据结构移植到yaml中，需要分清楚yaml处理的第一第二阶段的区别
- yaml处理第一阶段，更多考虑的是yaml和不同编程语言数据转换的一致性和可移植性
- 为了将有可读性的字符流呈现给人类阅读，有序是比不可少的，所以yaml第二阶段更多关注于序列化(成一个有序图)

### 节点图的模型

- 编程语言的数据结构，用yaml的3种原语来表示
- yaml的节点图,就是由 根节点，连接节点，带tag节点的有向图组成
- 其中有向图是：有一系列node和有向带箭头的连线，edge指(node到node之间的连线)
- 在节点图中，所有的node都是要有连接的，最后会连在root节点上，连接的方式是edge
- 所以在节点图中，会存在循环，一个node可能会有多个edge(连线)
- 下面用nodes术语表示有关联的node集合，nodes并不包含字面量node
- nodes要么是map要么是数组，对应3种原语

从[节点图的信息模型](https://yaml.org/spec/1.2/spec.html#id2763754)中可以看出：
- 对应3种语言，有3类节点：
  - 数组node，里面是有序内容
  - map node，里面是无序内容
  - 字面量node，里面是基本内容(字符串 整数 时间 日期等)
- 数组node，map node，字面量node都是继承于Node,和Node也有使用的依赖关系：
  - 数组内容也是Node，多对多的关系
  - map内容是包含多个kv对，一个kv对包含一个key和一个value
    - key和value都可以包含多个Node
  - 字面量node就是由string node/ int node/ 时间 node/ 日期node 组合而成
  - 所以说，字面量node里是基础内容，map和数组都是由Node组成，map可以包含字面量，数组可包含map
- Node是一个父类，派生出上面3种node
  - Node里是可以包含Tag的，tag是原生数据结构的一个简称
  - tag里面包含简称name和数据结构
  - 字面量tag就是对基础类型的一个别名

上面提到了很多形象化的概念，下面来解释一下：

node概念:
- 一个yaml node用于表示一个原生的数据结构
- 按内容分可分3种：
  - 包含基础内容的字面量node
    - 字面量node的内容是一个不透明的数据，可以用0或者更多unicode字符来表示
  - 包含有序内容的数组node
    - 数组node的内容是0个或多个nodes，带顺序的
    - 实际使用中，数组node更多的是包含同一数据类型的node
  - 包含无序内容的map node
    - map node里的内容是无序的kv对，key是唯一的
    - 实际使用中，key可以使任意nodes，value也是一样
- 每个node都包含一个tag，tag用来限制内容中可能出现的值

tag概念:
- 原生数据结构的一个简称叫tag
- 推荐使用tag uri方案来表示yaml tags
- 全局tag是uris，可以被所有应用使用(应用指使用yaml的程序)
- 本地tag是给某一个应用使用的，本地tag用!开头，不是一个uri，也不预期是全局唯一
- yaml提供TAG指令来减少tag符号冗余，这些指令可以将本地tag和全局tag进行相互转换
- tag不管理相同字符串开头的tag之间的关系
  - 也不要求tag结尾一定要有uri fragments(#后面的部分)
  - tag共享基础uri，并用后面的fragments来区别不同
  - fragments也是一种另类的标签，可以使用/ 来表示命名空间继承，这是惯例
    - 也可以使用其他：perl使用::来表示命名空间，java使用.来表示命名空间继承
- tag用于关联每个node的元信息
  - 每个tag都要指明node的kide，node的类型就3种

node比较：
- map的key是需要唯一的，所以node需要有比较的概念
- 普通格式
  - 也就是字面量的比较
  - 需要将所有需要比较的字面量转成一种普通格式
  - 然后通过普通格式的比较来判断是否一致
  - 当然，其中会使用字面量tag
- 相等
  - 两个node的tag和内容都相同，就是相等
  - tag表示的类型，内容表示的数据，两者相同就表示node相等
  - 字面量/数组/map的比较都是类似，嵌套比较每一个node的tag和内容
- 恒等
  - 意味着她们表示相同的原生数据结构
  - 一般来说，她们都有相同的内存地址

### 事件树的模型

- 针对节点图，事件树添加了：
  - map的有序key
  - 别名机制(一个类型第二次出现时，使用占位符表示)
- 这个过程的输出称为序列化树，也叫事件树

从[事件树信息模型](https://yaml.org/spec/1.2/spec.html#id2765410)中可以看出：
- 新增了一个别名node;map node添加了有序key;Node中添加了占位符

概念解释：

有序key：
- 前面也提到过，第一阶段map是无序的，为了序列化，就给key加了一个排序，就是有序key
- 这个有序key只是用于辅助序列化的，对第一阶段和第三阶段的数据并无副作用
- 在yaml中，有序的是用数组表示，map的key变成有序key之后，yaml是这样表示的：
  - 以前是map，里面是kv对
  - 现在是数组，里面方的是map，map里面只有一个kv对
  - yaml对这种有特殊处理，更加紧凑

别名和占位符：
- 在第一阶段里(节点图中)，一个node可能出现在多个集合(数组/map)中
- 在第二阶段序列化中，node第一次出现会被标识为anchor，后面出现都用占位符来表示那个anchor
- anchor可以不唯一，占位符只会找最近一个anchor
- 有anchor，并不意味着一定有相应的占位符来引用

### 字节流的模型

- unicode 字符流，里面包含了：
  - 风格
  - 字面量格式
  - 注释
  - 指令
  - 其他增加可读性的细节

相对于事件树来说，字节流的模型新增来了：
- 注释
- 指令
- 非特指的tag
- 字面量node的内容包含了格式化
- Node新增了风格/空格/行相关的缩进和表示
 
概念解释：

node风格：
- 在字节流中，node在呈现时，会依据其kind(类型)，可选择呈现的风格
- 这里的呈现风格，和事件树/节点图没有任何关系
- 总共有两种风格组
  - block 块式 用缩进来表示结构体
  - flow  流式 更依赖明确的指标(约定的指示符)
- yaml提供了很多字面量风格
  - 块式字面量风格包含了literal风格(文字风格)和folded风格(折叠风格)
  - 流式字面量风格包含了plain风格(简洁风格)和two quoted风格(两种引号风格)
  - 这些风格是在表现力和可读性上做了一定的权衡
- yaml针对数组和map
  - 块式，一般是从第二行开始，称为next-line风格，也叫下一行风格
    - 同时也提供了更加紧凑的，从第一行开始的，称为in-line风格，也称行内风格
  - 流式，提供了explicit风格，叫明确风格
    - map还提供了single-pair，单kv对风格，也是实现有序key的基础

|#| 字面量 | 数组 | map |
|---|---|---|---|
|flow|quoted(引号风格：包括单引号和双引号风格)<br /> plain(简洁风格)|explicit(明确风格)|explicit(明确风格) <br />single-pair(单kv对风格)|
|block|literal(文字风格) <br />floded(折叠风格)|next-line(下行风格) <br />in-line(行内风格)|next-line <br />in-line|

- 在yaml中，字面量有多种表示方法
  - 整数有各种进制
  - node比较时还需要一种机制来保证各种不同格式的转换，需要tag来指定
  - 但是呈现风格，和事件树/节点图没有任何关系，这里只是用不同的风格去呈现给人类，可读性是重点设计的
- 注释
  - 注释不影响事件树和节点图，甚至和node也每什么关系
  - 注释是用于维护者之间的交流
  - 一般既是配置文件中才会出现注释
  - 字面量中是不会出现注释的
- 指令
  - 指令不影响事件树和节点图
  - 指令是针对文档的，一个指令会有一个名字和一些可选参数
  - 这些指令是告诉yaml的处理器的
  - 目前只有YAML和TAG两个指令

## 加载失败点

从yaml字节流加载，到生成原生数据结构，会有多个失败点，这些叫加载失败点

这些失败的地方可能有很多：
- 字符流格式错误
- 别名未定义
- 未指定的tag可能无法解析
- tag无法识别
- 内容无效
- 原生类型不可用
- etc

分析：
- yaml信息分3个阶段来处理，每个阶段只处理一个目的
- 坏处是每个阶段对其他阶段的错误是发现不了的，好处是非常弱的关系导致新增类型是非常容易的
  - eg：地铁安检，包包检测用专门的仪器，罪犯检测用人脸识别，没有必要将人脸检测装到包包仪器上
  - eg：对于后续的包包的细菌检查，再加一台检测机器就行，这是基于当前最省的方式，不然包包一起加细菌检测的花费要高10倍了

格式ok的字符流和正确的别名
- 格式ok的字符流要符合bnf规范
- 使用别名，需要前面出现过anchor的

tag解析
- 字符流中很少需要显示出现tag的
- 在解析阶段(字符流解析成事件树)，没有显示指定tag的node，会指定一个非指定tag
- !可表示非plain字面量node，?可表示其他node
- 在组合阶段(事件树解析成节点图)，需要将未指定tag解析为指定tag(也就是说要解析为全局tag或本地tag)
- tag解析有3个参数：
  - node的未指定tag
  - root到node的路径
  - node的内容(基于kind)
  - 其中还需要考虑到别名，遇到别名，就要找到第一次出现node的地方
- 解析不考虑注释/缩进/风格
- 也不考虑其他node的内容，除非key node的内容和root有某些关系
- 也不考虑集合中的其他兄弟node
- 如果key node已经解析，就不考虑value node的内容

