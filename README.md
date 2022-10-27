# smart-tool
一个简单的快捷方式管理工具，可以将一些常用文件、文件夹、网址、脚本加到里面进行管理，目前只适用于Windows系统。

如果加进去的是某些文件路径比如`c:\myDoc.docx`文件将会调用系统默认打开程序进行打开，如果是目录就会直接打开文件管理器并进入某个目录，如果是网址将会使用默认浏览器打开，如果需要选择某个浏览器，可以将浏览器安装路径添加到PATH，然后用类似`chrome https://baidu.com`的格式填入即可选择浏览器，当然也可以编写一些脚本比如js脚本，通过node快速调用：`node c:\test.js`。

<p>
  <a href="https://github.com/Wyatex/smart-tool/releases/latest">
    <img src="https://img.shields.io/github/v/release/Wyatex/smart-tool" />
  </a>
  <a href="https://github.com/Wyatex/smart-tool/actions">
    <img src="https://github.com/Wyatex/smart-tool/actions/workflows/build.yml/badge.svg" />
  </a>
</p>

## 技术栈

- 语言：Go、TypeScript
- 打包工具：wails
- 前端框架：Vue
- UI 库：Navie-UI

## 使用方式

1. 安装 go1.17+、node1.16+
2. 安装 wails 工具：`go install github.com/wailsapp/wails/v2/cmd/wails@latest`
3. 检查环境依赖:`wails doctor`，按提示安装必须的依赖
3. 下载项目，命令行打开到项目目录
4. 启动 dev：`wails dev`， 打包：`wails build`，建议安装 upx，打包时加上`-upx`进行压缩能大幅减少体积

> To Linux用户：编译可能需要先cd到frontend，然后执行pnpm i和pnpm build，再执行wails dev或者wails build
