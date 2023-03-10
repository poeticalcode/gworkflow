## gworkflow








### process 节点模型




startModel : 开始节点

startModel : 结束节点

workModel : 工作节点

taskModel : 任务节点

joinModel : 会签节点

subProcessModel : 子流程节点



### process 定义

```xml
<?xml version="1.0" encoding="UTF-8" standalone="no"?>
<process displayName="测试流程创建模型" name="notallow">
  <!-- 开始节点 必须存在-->
  <start displayName="start1" layout="64,83,-1,-1" name="start1">
      <!-- 节点流转，to 是流转到那个节点-->
      <transition g="" name="transition1" offset="0,0" to="task1"/>
  </start>
  <!-- 定义具体工作流 -->
  <task assignee="task1.operator" displayName="task1" layout="213,80,-1,-1" name="task1" performType="ANY">
      <!-- 节点流转，to 是流转到那个节点-->
      <transition g="" name="transition2" offset="0,0" to="end1"/>
  </task>
  <!-- 结束节点 -->
  <end displayName="end1" layout="454,82,-1,-1" name="end1"/>
</process>
```


