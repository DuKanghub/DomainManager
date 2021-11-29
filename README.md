# 1. 基本介绍

本项目是基于框架[gin-vue-admin](https://www.gin-vue-admin.com)编写，目前已实现功能：

- 域名解析管理
  - 阿里云域名管理
  - 腾讯云域名管理
  - 多账号多平台
- 通过域名检查证书过期

# 2. 使用说明

```
- node版本 > v12.18.3
- golang版本 >= v1.16
- IDE推荐：Goland
- 初始化项目：不要手动修改config.yaml的数据库信息
- 替换掉项目中的七牛云公钥，私钥，仓名和默认url地址，以免发生测试文件数据错乱
```

## server项目

使用 `Goland` 等编辑工具，打开server目录，不可以打开`DomainManager`根目录：

```sh

# 克隆项目
git clone https://github.com/Dukanghub/DomainManager.git
# 进入server文件夹
cd server

# 使用 go mod 并安装go依赖包
go generate

# 编译 
go build -o server main.go # (windows编译命令为go build -o server.exe main.go )

# 生成API文档(可选)，
go get -u github.com/swaggo/swag/cmd/swag
swag init	#每次写了新的api时，都可以使用此命令生成api文档
# api文档地址：http://127.0.0.1:8888/swagger/index.html
# 运行二进制
./server (windows运行命令为 server.exe)

```

如果需要修改后端端口，修改`server/config.yaml`里的`system.addr`，这里用的是`8888`，如果修改了，前端一样要修改，否则会连接不上。

## 前端项目

前提：安装node

进入web目录或使用IDE打开web目录：

- 打开工具的终端,输入`npm i` 或者 `cnpm i` 进行安装web项目的环境
- 安装完成之后使用`npm run serve`或者`cnpm run serve`即可启动项目

前端端口：8080，默认使用的是`.env.development`配置，配置的后端地址是：127.0.0.1:8888。

也就是说打开`localhost:8080`即可看到后台界面了，用户名：admin，密码：123456

## 初始化

打开后台后，选择初始化，配置数据库连接信息，注意账号一定要有读写权限。

数据库和账号需要自己提前创建。表会自己创建。

## 域名管理使用

### 1. 添加RAM用户

我们的域名管理都是使用RAM用户对接云平台的。这步的RAM用户只是添加一个名字，一个易识别的名字，方便在网页展示。

所以我们将其放在`超级管理员`-->`字典管理`-->`RAM账号`中，点击`详情`，`新增字典项`，表单说明如下：

- 展示值：可以是中文，方便区分即可，可以一眼看出是哪个账号的
- 字典值：设置为非0即可，
- 启用状态：不用改
- 排序标记：建议从1开始，从小到大设置。

### 2. 添加RAM用户信息

这里是真正添加RAM用户token信息的地方。

- dns子用户：输入你在云平台创建这个RAM用户时的名字即可。
- 平台：按下拉框选择，目前可选项有：阿里云，腾讯云
- AccessKey：RAM用户的access_key
- AccessSecret：RAM用户的
- 所属账号：下拉选择`添加RAM用户`创建的用户。**如果下拉找不到之前在字典管理那里创建的用户，可以退出重新登陆重试。**

### 3. 从云平台拉取账号上的域名

转到 域名列表，点击`刷新域名缓存`，获取所有账号的域名到数据库里。如果域名较多，可能用时较长。

当然，也可以自己手动添加的方式，添加域名，这个页面还有个字段，就是公司。前面没有提到，如果你的账号属于不同公司，可以在字典里定义好公司名

这里可以批量修改公司名。

支持批量修改的字段有：公司、DNS提供商(即平台)、所属账号。

### 4. 获取所有域名的已有解析结果。

转到解析列表，域名列表，点击任一域名的详情进入解析列表。

点击 刷新解析缓存，和3的作用差不多，只不过这个是拉各域名的解析结果到数据库。

### 5. 接下来就是自行体验

## 部署

- 前端

  在web目录下执行 

  ```sh
  npm run build 
  ```

  得到 dist文件夹，将dist文件夹上传到服务器。

  建议使用nginx进行代理，并且设置 proxy 把请求代理到后端。

- 后端：

  在 server目录下 

  ```sh
  go build 
  ```

  得到一个可执行文件, 然后将`可执行文件和config.ymal 以及 resource 文件夹`上传至服务器 三者最好放在同一路径下 最终服务器目录结构可能如下

  ```sh
  
      ├── breakpointDir  # 后续断点续传自动生成
      ├── chunk   # 后续断点续传自动生成
      ├── fileDir   # 后续断点续传自动生成
      ├── finish   # 后续断点续传自动生成
      ├── resource
      │   └── 子目录文件                   
      ├── dist	#这个是前端打包的
      │   └── 子目录文件
      ├── gin-vue-admin	#后端二进制文件
      ├── config.ymal	#配置文件
      
  ```

- 前后端可以分开机器部署，需要修改前端配置。

- nginx反向代理配置

  ```sh
  server {
  	...
  	location / {
          root /data/www/domainManager/wwwroot/dist;
          index index.html index.htm;
      }
      location /api {
          proxy_set_header Host $host;
          root /data/www/domainManager/wwwroot;
          proxy_http_version 1.1;
          proxy_set_header Connection "upgrade";
          proxy_set_header Upgrade $http_upgrade;
          proxy_set_header X-Real-IP $remote_addr;
          proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
          proxy_set_header X-Forwarded-Proto $scheme;
          rewrite ^/api/(.*)$ /$1 break;
          proxy_pass http://localhost:8888;
      }
  }
  	
  ```

  