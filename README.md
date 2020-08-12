# Cufg
Collect urls from github.

从github采集相关域名的url，page默认为10，即默认采集包含相关domain最新的100个项目。

## 示例

```
./main -t <your github token> -d "kuaishou.com"

[+] Start search ......
[+] domain:  kuaishou.com  page:  10
[+] Num:  1
   Git url:  https://github.com/oldmenplus/bookmark/blob/0c3556541887065bd7ce4ae0809a9adec19bd52e/bookmork.json
   Http url:  https://cp.kuaishou.com/profile
   Response title:  <title>快手创作者服务平台</title>
[+] Num:  2
   Git url:  https://github.com/oldmenplus/bookmark/blob/0c3556541887065bd7ce4ae0809a9adec19bd52e/bookmork.json
   Http url:  https://cp.kuaishou.com/profile
   Response title:  <title>快手创作者服务平台</title>
[+] Num:  3
   Git url:  https://github.com/tzki/tzki.github.io/blob/495152bfd09dd1a7e4feffcc4d5c24d6f20f0c16/posts/bf77afb8/index.html
   Http url:  https://live.kuaishou.com
   Response title:  <title>首页-快手直播</title>
[+] Num:  4
   Git url:  https://github.com/tzki/tzki.github.io/blob/495152bfd09dd1a7e4feffcc4d5c24d6f20f0c16/posts/bf77afb8/index.html
   Http url:  https://live.kuaishou.com/live-partner&quot
   Response title:  <title>出错啦-快手直播</title>
.....
```
