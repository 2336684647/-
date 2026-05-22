# NekoCafé 智慧餐饮预约平台
> 实验三 DevOps 容器化部署 PoC

## 项目简介
本项目为 NekoCafé 猫咪主题餐饮预约平台的 DevOps 部署原型，包含两个核心微服务：
- 会员服务（member-service）：用户注册、登录、积分管理
- 预约服务（reservation-service）：门店查询、桌位预约、排队管理

采用 Docker + Docker Compose 实现容器化编排，支持一键启动、健康检查、服务隔离与自动回滚。

---

## 前置依赖
- Docker 24.0+
- Docker Compose v2+
- Make（可选，用于快捷命令）

---

## 一键启动服务
```bash
# 方式1：使用Make快捷命令（推荐）
make up

# 方式2：直接使用Docker Compose
docker compose up -d --build