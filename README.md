# yys

#### 介绍
一个云原生项目，模拟一条完整的云原生生产线，涉及到各类云原生知识，k8s，容器化，go微服务开发，运维，服务网格istio，存储集群，日志系统，监控系统，devops等 

主要的知识点技术点：  
云原生基础环境：k8s docker linux  
微服务开发：go go-zero istio  
CI/CD：tekton gitops  
存储：rook-ceph nfs共享卷  
监控：prometheus+grafana+alertmanager  
日志：ELK Loki  
其他：mysql redis kafka ansible等  

整个项目高度容器化，k8s是基石环境，采用go语言go-zero框架编写的微服务架构的应用程序(当然这里单体和微服务之间的转换很平滑,为了更明显突出servicemesh效果我优先考虑了微服务架构)，主题上可封为两大部分，后台管理系统和服务应用，在支付和订单的逻辑处理间为了保证消息数据的正确使用，接入kafka做消息队列，部分延时功能用ansyq做延时队列，运行在istio中，做流量管理负载均衡、安全管控、链路追踪、故障注入、熔断、流量镜像等。应用提供了metrics接口，prometheus拉取metrics数据做监控grafana可视和告警通知等。使用ceph集群和nfs共享卷做存储，为了更优化本地程序的性能，部分程序采用localpv。tekton配合我本地配置的gitlab和harbor做CI/CD  


我的环境：4台各4核8g的ESC  
![image](https://user-images.githubusercontent.com/61965693/191806630-8aa15961-3ddc-4b82-81ac-d830d0f17cb2.png)

通过我的公网地址和对应的端口即可访问我对外提供的服务，上面我提到的服务大部分对外开放(后续我考虑将服务端口整理出来供网友公网访问)

个人预计后续优化：  
1.将程序做成helm包，或者使用operator来实现更贴近云环境的开发  
2.接入dapr，更好的开发微服务  
3.配置Loki，轻量且高效  
4.加入gitops,tekton+gitops理念实现更真实的CI/CD。jenking使用kubernetes插件来配置云做cicd(无他，使用用户量在，但我觉得kubernetes用jenkins做cicd有点重)  
5.缓存相应的代码逻辑进一步设计优化，缓存设计不好并发就是空谈  
6.数据库换osd，有istio接入，实现更好的高可用，更方便实现如读写分离等操作  
7.监控promenthues+grafana+alertmanager不够轻，采用promentheusOperator或者Thnos  


#### 放下部分图片介做绍  

链路追踪可视：  
![istio2](https://foruda.gitee.com/images/1663863559276064143/1e0e2f85_10984789.png "kiali11.png")  
![istio1](https://foruda.gitee.com/images/1663863513702245341/3f1a03a7_10984789.png "kiali8.png")  

CI/CD:  
![输入图片说明](https://foruda.gitee.com/images/1663863643116081908/2a924f51_10984789.png "tekton.png")  
![image](https://user-images.githubusercontent.com/61965693/191807118-ecf6aa9f-ddc1-427c-b31d-4144bfe60140.png)  
![image](https://user-images.githubusercontent.com/61965693/191807233-722bfe7c-c516-4152-a622-78f1b664af41.png)  

 
业务逻辑：  
![image](https://user-images.githubusercontent.com/61965693/191807922-ad33d60a-8924-490f-a67a-72568308db73.png)
![image](https://user-images.githubusercontent.com/61965693/191808026-726d486b-c619-4915-9373-4cf65e9acb9c.png)  

相应的service:  
![image](https://user-images.githubusercontent.com/61965693/191807377-6357f70b-331e-4319-ac59-0d4f118be3cf.png)  
![image](https://user-images.githubusercontent.com/61965693/191807432-2b1ec45d-01a3-43f6-8398-d204e8df1b53.png)  

监控：  
![image](https://user-images.githubusercontent.com/61965693/191807510-49b28b66-8a5a-4c5c-a893-857bd4e45e92.png)
![image](https://user-images.githubusercontent.com/61965693/191807594-1d5010ec-e771-461c-87c6-ac1271d533d8.png)  
![image](https://user-images.githubusercontent.com/61965693/191807685-1d6f8685-ccbd-4175-9ae9-8733af18080f.png)  
