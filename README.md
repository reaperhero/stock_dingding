# 股市分析

<div align="center">
  <a href="https://alist.nn.ci"><img height="100px" alt="logo" src="https://cdn.jsdelivr.net/gh/alist-org/logo@main/logo.svg"/></a>
  <p><em>🗂️A file list program that supports multiple storage, powered by Gin and React.</em></p>
  <a href="https://github.com/Xhofe/alist/releases"><img src="https://img.shields.io/github/release/Xhofe/alist?style=flat-square" alt="latest version"></a>
  <a href="https://github.com/Xhofe/alist/discussions"><img src="https://img.shields.io/github/discussions/Xhofe/alist?color=%23ED8936&style=flat-square" alt="discussions"></a>
  <a href="https://github.com/Xhofe/alist/actions?query=workflow%3ABuild"><img src="https://img.shields.io/github/workflow/status/Xhofe/alist/build?style=flat-square" alt="Build status"></a>
  <a href="https://github.com/Xhofe/alist/releases"><img src="https://img.shields.io/github/downloads/Xhofe/alist/total?style=flat-square&color=%239F7AEA" alt="Downloads"></a>
  <a href="https://github.com/Xhofe/alist/blob/v2/LICENSE"><img src="https://img.shields.io/github/license/Xhofe/alist?style=flat-square" alt="License"></a>
  <a href="https://pay.xhofe.top">
    <img src="https://img.shields.io/badge/%24-donate-ff69b4.svg?style=flat-square" alt="donate">
  </a>
</div>

---

## 同步数据

### - step 1：下载数据数据

<div align=center>
<img src="static/image/1.png" width=800 height=400 />
</div>

<div align=center>
<img src="static/image/Snipaste_2022-07-26_19-19-42.png" width=800 height=400 />
</div>

### - step 2：配置数据库

```
log_level = "DEBUG"

[mysql]
ip = "127.0.0.1"
port = 3306
user = "root"
password = "=nXCwk8l<@Dx5bL%"
database = "dongfang_stock"


[web]
port = 8080
```

### - step 3：开始同步

- 把剪切板的数据粘贴到`service/clipbroad/clipbroad.txt`文件
- 执行命令同步
```
go run cmd/cmd.go sync
```

## 每日分析报告

```
go run .\cmd\cmd.go report tofile
```

<div align=center>
<img src="static/image/3.png"  />
</div>


## 最近涨停票走势

```
go run .\cmd\cmd.go report stion
```

<div align=center>
<img src="static/image/Snipaste_2022-07-26_19-19-42.png" width=400 height=200 />
</div>