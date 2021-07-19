# 需求设计

todo list

- 用户通过邮箱注册，登录，找回密码

- 用户可以添加todo事项，事项包括以下信息
  - 内容
  - 日期
  - 地点
  - 标签（类型）
  - 是否置顶
- Todo分类标准
  - 日期
  - 是否置顶
  - 全部
  - 今日todo


# 数据库设计

- user

  - User id
  - name
  - password

- Todo-user

  - User id
  - User todo list -- as array

- Todo

  - Content
  - deadline
  - tag
  - location
  - Pin to top


# 接口设计
Restful





