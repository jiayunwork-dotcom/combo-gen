# combo-gen

`combo-gen` 是一个纯 Go 标准库实现的组合数学生成器：给定一组元素，可枚举全排列、k-排列、k-组合、可重复组合与笛卡尔积。

## 构建 / 运行 / 测试

```text
go build ./...     # 编译
go run .
go test ./...      # 测试
```

## 评测镜像

本目录评测专用文件（勿覆盖项目自带 Dockerfile/README）：

- `benzhi.Dockerfile`
- `build_benzhi_docker.sh`
- `BENZHI_README.md`（本文件）

两种架构都要构建并进容器验证：

```bash
chmod +x build_benzhi_docker.sh
./build_benzhi_docker.sh <image-name> linux/arm64
./build_benzhi_docker.sh <image-name> linux/amd64
docker run -it <image-name>:latest
```
