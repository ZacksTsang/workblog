工作日记
========
记录一些开发的中所遇到的问题及解决方法，比较凌乱

## [ angular 方面](https://github.com/ZacksTsang/workblog/blob/master/angular.md)

## jquery 方面
- 有关[window.open](https://github.com/ZacksTsang/workblog/blob/master/windowopen.md)

## 一些技巧
#### 如何将 iOS 上的 `*.strings` 资源文件转成 Android 上的 `string.xml` 资源文件？
使用 Python 语言可以快速实现这一功能，具体实现请查看 [file2xml.py](https://github.com/ZacksTsang/workblog/blob/master/tools/python/iosfile2androidxml/file2xml.py)文件

使用说明：在该文件的目录下，将 iOS 的资源文件拷贝到此目录下，运行以下命令即可，<b>*此命令还可以转换资源中拼接的字符串*</b>
```
 $python file2xml.py filename tofilename
```
### Shell
#### 如何快速在控制台统一修改文件的后缀名？
此种情况一般是用在修改图片的后缀名上，在需要修改的目录下
```
for f in *
mv $f ${f%.*}.后缀名
```
注：在shell脚本中，`f="某文件"`，获取该文件名: `${f%.*}`，获取后缀名: `${f##*.}`

## Go 相关
- 如何给一组 version string 排序

	```
	[]string {"1.6beta1", "1.5rc1", "1.5beta2", "1.5beta1", "1.5.1", "1.5", "1.4rc2", "1.4rc1", "1.4beta1", "1.4.2", "1.4.1", "1.4", "1.3rc2", "1.3rc1", "1.3beta2", "1.3beta1", "1.3.3", "1.3.2", "1.3.1", "1.3", "1.2rc5", "1.2rc4", "1.2rc3", "1.2rc2", "1.2rc1", "1.2.2", "1.2.1", "1.2", "1.1.2", "1.1.1", "1.1", "1.0.3", "1.0.2", "1.5.2", "1.5alpha1"}
	```
	相关代码在：[VersionSort](https://github.com/ZacksTsang/workblog/blob/master/tools/go/VersionSort.go)