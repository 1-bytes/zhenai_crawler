### 前言

go语言写的一个爬虫项目，欢迎大家借鉴、参考 :)

### 基础环境

- Golang (1.16.5)
- ElasticSearch (7.x)

Tips: 推荐在 docker 内一键部署 ElasticSearch

---
**docker 下载地址**：
> https://www.docker.com/products/docker-desktop


**docker 国内镜像：**

~~~
{
    "registry-mirrors": [
        "https://hub-mirror.c.163.com",
        "https://docker.mirrors.ustc.edu.cn"
    ]
}
~~~

**ElasticSearch Docker 容器镜像下载/启动：**

~~~
下载
$ docker pull docker.elastic.co/elasticsearch/elasticsearch:7.14.0

启动
$ docker run -p 9200:9200 -p 9300:9300 -e "discovery.type=single-node" docker.elastic.co/elasticsearch/elasticsearch:7.14.0
~~~
