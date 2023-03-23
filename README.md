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
- 通用环境变量
    - `ADDR`: 配置监听地址，默认 `:8000`
    - WebProxy 使用的代理
        - `HTTP_PROXY`
        - `http_proxy`
        - `HTTPS_PROXY`
        - `https_proxy`
    - `LEVEL`: 日志等级，默认 `info`
        - `debug`
        - `info`
        - `error`
    - `SCHEME`: 转发回源时使用的协议(默认 `https`)
- 网络代理模式环境变量
    - `SERVER_HOST`: WebProxy 地址，替换修改 gist 文件内的连接也通过 WebProxy 访问
- HTTP/TCP 转发模式环境变量
    - `RADDR`: TCP/HTTP 转发地址


## 使用样例

**场景 1** 反向代理 Github、Gist 等站点

与转发 HTTP 类似，但是会根据 Path 不同路由到不同的镜像网站，且对于 Gist 等存在特殊处理

```bash
docker run --rm -d \
    -e ADDR=0.0.0.0:8000 \
    -e SERVER_HOST=https://proxy.ohyee.cc \
    ohyee/webproxy:latest \
    -t proxy
```

**场景 2** 转发 TCP/HTTP

如存在生产环境和测试环境，请求时为了快速切换环境，设置调用方发送至 `127.0.0.1:9999`，通过切换 `RADDR` 的值来切换环境

```bash
docker run --rm -d \
    --net=host \
    -e "SCHEME=http" \
    -e "ADDR=:9999" \
    -e "RADDR=http://127.0.0.1:8888" \
    ohyee/proxy:latest \
    -t http
```
