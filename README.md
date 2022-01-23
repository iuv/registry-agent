# registry-agent

实现k8s环境下spring cloud 项目不改代码去掉注册中心

将服务发现转为k8s service方式

使用方法：

1. 为服务添加与FeignClient调用名相同的Service
2. 将本项目编译为docker镜像
3. 部署服务时使用Sidecar形式（或多容器）运行
4. 指定服务注册中心地址为127.0.0.1:80

完成以上步骤即可

> 注： 目前仅支持nacos、consul的替换

> 注2： 暂时不支持使用注册中心获取配置功能
