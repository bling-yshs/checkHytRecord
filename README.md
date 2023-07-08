# checkHytRecord
花雨庭战绩查询，一键鉴定纪狗(没战绩的就是新玩家，纪狗几率+80%)

### 编译

```
go build -o 花雨庭查成分.exe
```

### 修改图标

```
windres -o app.syso -i app.rc
```

### 请求 UAC 并添加图标

```
rsrc -manifest manifest.xml -ico app.ico -o app.syso
```

注意 `syso` 文件只能存在一个