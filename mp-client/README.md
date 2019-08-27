## 小程序客户端

> [mpvue quickstart](http://mpvue.com/mpvue/quickstart.html)

### 1. 先检查下 Node.js 是否安装成功

```
$ node -v
v8.9.0

$ npm -v
5.6.0
```

### 2. 由于众所周知的原因，可以考虑切换源为 taobao 源

```
$ npm set registry https://registry.npm.taobao.org/
```

### 3. 全局安装 vue-cli

一般是要 sudo 权限的

```
$ npm install --global vue-cli@2.9
```

### 4. 创建一个基于 mpvue-quickstart 模板的新项目

新手一路回车选择默认就可以了

```
$ vue init mpvue/mpvue-quickstart my-project
```

### 5. 安装依赖，走你

```
$ cd my-project
$ npm install
$ npm run dev
```
