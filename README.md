
##### 背景
- 学习了golang语言需要在工程中去实践
- 爬虫项目想要做的很丰富将会是一个很大的项目，如果只想爬爬数据，也可以做的很轻松。计划如下:
    1. [x] V1.0 ✅简单的获取到数据->解析数据->存储数据->展示数据；采用的技术栈是: Golang, ES, GolangTemplate。
           问题:1. 爬的比较慢 2. 单节点 
    2. [x] V2.0 ❌重构V1.0，增加并发，提高爬取效率，完善数据价值
    3. [x] V3.0 ❌重构V2.0，服务拆分，增加RPC调用，分布式节点
    4. [x] V4.0 ❌重构V3.0，Docker + Kubernetes 部署，至此分布式爬虫完成
    5. [x] V5.0 ❌增加数据分析,丰富数据价值，采用 ELK Stack(Elasticsearch, Logstash, Kibana) 
- 目标: 掌握Golang, Docker, Kubernetes, ELK Stack 在工程中的关键点
- V2.0 架构设计如下:

![img](https://raw.githubusercontent.com/MichaelYgZhang/michaelygzhang.github.io/master/images/crawler.jpg)

##### 记录V1.0
- Golang-Spider-XXX网站
    - 拉取城市列表解析
    - 解析每个城市列表 + 每个人的信息(别名,性别,年龄,收入,婚姻状况....等)
    - 数据入ES存储
    - Golang, Template前端展示
##### TODO
- 重构,增加队列,抽象出解析器,存储接口
- 城市列表分页,可以增加数据量
- 调用Face++颜值分数服务对每个人打颜值分 ✅
- 结合颜值/身高/收入/年龄/婚姻情况，设计算法产出一个综合排名分值
##### 终极目标
- 分布式
- 反爬机制，爬动态数据
- 分布式去重 (redis?)
- 大数据分析ELK技术
- 脚本部署，Docker + Kubernetes 部署
- 集成服务注册与发现 consul
- 使用Logstash 汇总和分析日志
