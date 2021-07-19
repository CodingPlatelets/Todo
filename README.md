# Yiban Todo

# 简介: 创建一个 管理TODO待办事项的 web管理后台

完成效果参考:  ( 仅考虑后端接口实现 )

![](https://img2020.cnblogs.com/blog/1609091/202107/1609091-20210715201417230-1650707211.png)

接口定义包括以下部分:

## 用户相关 /user

```
{{server}}/user/register
desc: 用户注册
method: POST
body:
    {
        username: str,
        password: str
    }
returns:
    {
        code: 200,
        msg: "OK"
    }


{{server}}/user/login
desc: 用户登录
method: POST
body:
    {
        username: str,
        password: str
    }
returns:
    {
        code: 200,
        msg: "OK"
    }


{{server}}/user/logout
desc: 用户登出
method: POST
body:
    user_id: int
returns:
    {
        code: 200,
        msg: "OK"
    }

```



## TODO事项分组相关 /todo_group

```
{{server}}/todo_group/add
desc: 创建待办事项的分组
method: POST
body:
    {
        user_id: int,
        todo_group_name: str
    }
returns:
    {
        code: 200,
        msg: "OK"
    }

{{server}}/todo_group/list
desc: 获取待办事项的所有分组基本信息
method: GET
returns:
    {
        code: 200,
        msg: "OK",
        todo_group_list:{
            {
                todo_group_id: int,
                todo_group_name: str,
                item_count: int  # 该 todo_group 中 todo_item的个数
            },
            {
                todo_group_id: int,
                todo_group_name: str,
                item_count: int
            }
        }
    }


{{server}}/todo_group/<id>
desc: 修改待办事项组的信息
method: PUT
body:
    {
        todo_group_id: int,
        todo_group_name: str
    }
returns:
    {
        code: 200,
        msg: "OK"
    }
    
{{server}}/todo_group/<id>
desc: 删除待办事项组的信息  # 需要考虑删除相关数据
method: DELETE
body:  
    todo_group_id: int
returns:
    {
        code: 200,
        msg: "OK"
    }


```



## TODO事项相关 /todo

```
{{server}}/todo/add
desc: 创建 TODO_item
method: POST
body:
    {
        user_id: int
        todo_group_id: int
        todo_title: str
        todo_content: str
    }
returns:
    {
        code: 200,
        msg: "OK"
    }


{{server}}/todo/list
desc: 查看 TODO 列表
method: GET
params:  # 筛选条件
    create_at: 'datetime'  # 根据创建时间筛选
    keyword: str  # 根据关键词筛选
    todo_group_id: int  # 根据分组筛选
    is_finished: bool  # 根据是否已完成筛选
returns:
    {
        code: 200,
        msg: "OK",
        todo_list:[
            {
                todo_id: int
                user_id: int
                todo_group_id: int
                todo_title: str
                todo_content: str
                is_finished: bool
                create_at: "datatime"
            },
            {
                todo_id: int
                user_id: int
                todo_group_id: int
                todo_title: str
                todo_content: str
                is_finished: bool
                create_at: "datatime"
            }
        ]
    }


{{server}}/todo/<id>
desc: 修改 TODO_item
method: PUT
body:
    {
        todo_id: int  # 根据该id 即todo_item_id 确定要修改的具体对象
        todo_group_id: int  # 要修改到的分组 不传默认不修改
        todo_title: str
        todo_content: str
        is_finished: bool  # 只传id 与 此字段 即修改todo_item的完成状态
    }
returns:
    {
        code: 200,
        msg: "OK"
    }


{{server}}/todo/<id>
desc: 删除 TODO_item
method: DELETE
body:
    todo_id: int
returns:
    {
        code: 200,
        msg: "OK"
    }
```
