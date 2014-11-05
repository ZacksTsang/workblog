# angular学习记
记录在使用 angular 所遇到的问题及解决方案，和一些 angular 常用的技巧

## 问题

- 如何给 directive 中的自定义属性赋值
比如有一个自定义的属性 `<... date-from="o.date" .../>`，如何在 directive 中改变 o.date 的值？

因为可以通过 `$parse($attrs.dateFrom)($scope)` 获取此属性的值，那么改变此属性就 

```
    $parse($attrs.dateFrom + "='需要改变的值'")($scope);
    $scope.$apply();
```
或
```
    $scope.$apply(function(){
        $parse($attrs.dateFrom + "='需要改变的值'")($scope);
    });
```
