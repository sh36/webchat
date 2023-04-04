# webchat
项目基于[openwechat](https://github.com/eatmoreapple/openwechat)开发

# 配置
### 获取chatGLM部署后端地址
配置项在 config.json ```glm_backend```

### 限制对话轮数
配置项在 config.json ```max_boxes```
每个用户单独配置上下文历史栈，超出限制轮数将清空历史。

### 配置最大并发用户数
配置项在 config.json ```user_count```

# 安装使用

### 获取项目
```git clone https://github.com/sh36/webchat.git```

### 进入项目目录
```cd webchat```

### 启动项目
1. ```nohup sh startup.sh &```进行项目后台启动

### 关闭项目
1. ```ps -aux | grep main.go```找到项目对应的进程id
2. ``` kill -9 进程id ```
