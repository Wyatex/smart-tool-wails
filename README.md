# smart-tool
一个简单的快捷方式管理工具

## 技术栈

- 语言：Go、TypeScript
- 打包工具：wails
- 前端框架：Vue
- UI 库：Navie-UI

## 使用方式

1. 安装 go1.17+、node1.16+
2. 安装 wails 工具：`go install github.com/wailsapp/wails/v2/cmd/wails@latest`
3. 下载项目，命令行打开到项目目录
4. 启动 dev：`wails dev`， 打包：`wails build`，建议安装 upx，打包时加上`-upx`进行压缩能大幅减少体积
