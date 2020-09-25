# ContinuousUpload

#### 介绍
基于golang的gin框架实现的简单断点续传模块，支持kubernetes部署

#### 软件架构
1、conf文件夹：环境变量，kubernetes配置文件

2、dockerfiles文件夹：docker镜像配置文件

3、scripts文件夹：启动脚本

4、src文件夹：代码

#### 安装教程

1.  需要golang 1.13.4版本

2.  配置环境变量，kubernetes配置文件

3.  bash scripts/compile.sh

4.  bash scripts/build_images.sh

5.  bash scripts/start.sh

#### 使用说明

如果每次upload成功，直接使用返回的文件大小作为下次上传的offset(file_size)，如果upload有错误，调用status接口获取实际的offset，再从实际的offset开始传。可以参照伪代码，大致如下逻辑：

statusRes = get(/upload/status)
if (statusRes.data.success === 0) {
    let fileName = statusRes.data.data.file_name
    let fileSize = statusRes.data.data.file_size
    while (fileSize < fileRaw.size) {
      let formData = new FormData()
      formData.append('file_size', fileSize)
      formData.append('file', fileRaw.slice(fileSize, Math.min(fileSize + 1024 * 1024, fileRaw.size)), fileName)
      let uploadRes = post(/upload, formData)
      if ((uploadRes.status !== 200) || (uploadRes.data.success !== 0)) {
        statusRes = get(/upload/status)
        fileName = statusRes.data.data.file_name
        fileSize = uploadRes.data.data.file_size
        continue
      }
      fileSize = uploadRes.data.data.file_size
    }
    return fileName
  }
  return null

#### 参与贡献

1.  Fork 本仓库
2.  新建 Feat_xxx 分支
3.  提交代码
4.  新建 Pull Request


#### 特技

1.  使用 Readme\_XXX.md 来支持不同的语言，例如 Readme\_en.md, Readme\_zh.md
2.  Gitee 官方博客 [blog.gitee.com](https://blog.gitee.com)
3.  你可以 [https://gitee.com/explore](https://gitee.com/explore) 这个地址来了解 Gitee 上的优秀开源项目
4.  [GVP](https://gitee.com/gvp) 全称是 Gitee 最有价值开源项目，是综合评定出的优秀开源项目
5.  Gitee 官方提供的使用手册 [https://gitee.com/help](https://gitee.com/help)
6.  Gitee 封面人物是一档用来展示 Gitee 会员风采的栏目 [https://gitee.com/gitee-stars/](https://gitee.com/gitee-stars/)
