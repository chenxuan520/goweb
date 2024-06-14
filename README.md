# goweb
- 一个轻量化简单的web服务器,用于进行临时测试
# Author
- **chenxuan**
# 使用方法
- `clone`后`make`编译,或者直接下载二进制文件运行
# 配置文件
- 配置文件为`config/config.json`,配置选项为
```json
{
  "server": {
    "port": 5200, // 绑定的端口
    "host": "0.0.0.0" // 绑定的ip(直接为0.0.0.0即可)
  },
  "static_paths": [
    {
      "file_path": ".", // 静态文件的系统路径
      "web_path": "/" // 绑定到web服务器的路径
    }
  ]
}
```
