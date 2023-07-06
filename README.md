## Introduce
框架：wails2
前端实现： Vue3 + Arco
后端实现： Golang

基于 Wails2 实现读取后缀为 qc 的文件增删改的
前端 vue3+Arco 页面显示如下：
1、左边栏显示文件列表，文件来源默认本地文件目录，如 /.qc，文件后缀为 .qc，类似导航栏。
2、文件列表每个文件增加一个开关，可以切换开启/关闭状态
3、右边栏显示文件内容，默认首页。双击左边栏文件列表的文件，显示对应的文件内容，支持修改保存
4、右击文件显示 编辑文件信息以及移动到回收站 的选项
5、上方栏显示两个按钮，一个点击可以关闭文件导航栏，只显示文件内容编辑框；一个用于可以添加文件

# 后端功能接口
- gorm 数据库存储☑️
- 数据初始化☑️
- 获取文件列表及文件内容☑️
- 保存文件修改☑️
- 新增文件☑️
- 删除文件☑️
- 数据库保存文件内容☑️
- 启动初始化创建 sqlite 文件☑️
- 初始化数据不重复问题
- 文件内容规范检查
- source 文件不存在则删除

# 前端后端接口对接
- 文件列表☑️
- 新增文件☑️
- 保存内容☑️
- 删除文件☑️
- 文件使用开关☑️
- 回收站功能

## 安装 waild
```
go install github.com/wailsapp/wails/v2/cmd/wails@latest
```

## 运行项目测试
```
wails dev
```

## 编译项目
```
wails build
```