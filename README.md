# WebProxy

## 拉取镜像

```bash
docker pull ohyee/proxy:latest
```

## 启动代理

```bash
docker run -it -p 8000:8000 --rm ohyee/proxy:latest
```

## 配置选项

启动参数
- `--type`: 支持 `http`(转发 HTTP)、`tcp`(转发 TCP)、`proxy`(网络代理模式)

环境变量
- `ADDR`: 配置监听地址，默认 `:8000`
- WebProxy 使用的代理
    - `HTTP_PROXY`
    - `http_proxy`
    - `HTTPS_PROXY`
    - `https_proxy`
- `SERVER_HOST`: WebProxy 地址，替换修改 gist 文件内的连接也通过 WebProxy 访问
- `LEVEL`: 日志等级，默认 `info`
    - `debug`
    - `info`
    - `error`
- `RADDR`: TCP/HTTP 转发地址
- `SCHEME`: 转发协议(默认 `https`)