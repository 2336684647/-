# 一键启动所有服务（构建+后台运行）
up:
	docker compose up -d --build

# 停止所有服务
stop:
	docker compose stop

# 停止并删除容器、网络
down:
	docker compose down

# 停止并删除所有数据（MySQL/Redis 数据会清空，开发环境用）
clean:
	docker compose down -v

# 重启所有服务
restart:
	docker compose restart

# 回滚到上一个稳定版本（拉取旧镜像并重启）
rollback:
	docker compose pull
	docker compose up -d