# <center>Message增删改查测试</center>

## <center>用户部分</center>

## 基本信息

BaseURL `http://127.0.0.1:8080`

验证方式

## 接口详情

### POST `/user` 创建用户

form-data

```json
{
	"username": "wuhan",
	"phone": "123666",  //手机号，不可重复
	"age": 19
}
```

Response

```json
{
	"data": {
		"uid": 11   //新建用户的ID
	},
	"error": 0,
	"msg": "success"
}
```

### GET `/users` 获取所有用户

Response

```json
{
	"data": [{
		"ID": 6,
		"CreatedAt": "2019-03-17T18:50:39+08:00", //创建时间
		"UpdatedAt": "2019-03-17T18:50:51+08:00", //更新时间
		"DeletedAt": null,
		"Username": "testuser",
        "Password":"password1",	//密码
		"Phone": "1066",
		"Age": 18,
        "Rank":1	//等级，默认为1，普通用户，2为管理以上，以后还可以改
	}, {
		"ID": 9,
		"CreatedAt": "2019-03-17T19:45:56+08:00",
		"UpdatedAt": "2019-03-17T19:45:56+08:00",
		"DeletedAt": null,
		"Username": "wuhan",
        "Password":"password2",
		"Phone": "12366",
		"Age": 19,
        "Rank":2
	}],
	"error": 0,
	"msg": "success"
}
```

### GET `/user/id/{id}` 通过id查找用户信息

Param

- `id`:(int)用户ID

Response

```json
{
	"data": {
		"ID": 10,
		"CreatedAt": "2019-03-17T19:54:42+08:00",
		"UpdatedAt": "2019-03-17T19:54:42+08:00",
		"DeletedAt": null,
		"Username": "wuhan",
        "Password": "password2",
		"Phone": "123666",
		"Age": 19,
        "Rank":0
	},
	"error": 0,
	"msg": "success"
}
```

### GET `/user/phone/{phone}` 通过电话查找用户信息

Param

- `phone`:(phone)用户电话

Response

```json
{
	"data": {
		"ID": 10,
		"CreatedAt": "2019-03-17T19:54:42+08:00",
		"UpdatedAt": "2019-03-17T19:54:42+08:00",
		"DeletedAt": null,
		"Username": "wuhan",
        "Password": "password2",
		"Phone": "123666",
		"Age": 19,
        "Rank":0
	},
	"error": 0,
	"msg": "success"
}
```

### PUT `/user` 修改信息



### DELETE `/user/id/{id}` 通过id删除用户

Param

- `id`:(int)用户ID

Success

HttpStatusCode:`200`



### DELETE `/user/phone/{phone}` 通过电话删除用户

Param

- `phone`:(string)用户电话

Success

HttpStatusCode:`200`

## 错误对照

| HttpStatusCode | Error | Msg                | Meaning       |
| -------------- | ----- | ------------------ | ------------- |
| 400            | 40000 | bad param          | url入参错误   |
| 400            | 40001 | bad payload        | payload不完整 |
| 403            | 40300 | phone already used | 手机号被占用  |
| 404            | 40400 | not found          | 用户不存在    |



---

## <center>留言部分</center>

## 基本信息

BaseURL `http://127.0.0.1:8080`

验证方式

## 接口详情



### POST `/message` 创建留言

form-data

```json
{
    "content": "留言内容",
    "uid":1 	//创建留言的用户uid
}
```

Response

```json
{
	"data": {
		"mid": 7
	},
	"error": 0,
	"msg": "success"
}
```

### GET `/users` 获取所有留言

Response

```json
{
	"data": [{
		"ID": 1,
		"CreatedAt": "2019-04-08T22:36:41+08:00",
		"UpdatedAt": "2019-04-08T22:48:19+08:00",
		"DeletedAt": null,
		"Content": "messageEdit",
		"UID": 2
	}, {
		"ID": 3,
		"CreatedAt": "2019-04-08T22:43:13+08:00",
		"UpdatedAt": "2019-04-08T22:59:24+08:00",
		"DeletedAt": null,
		"Content": "messageEdit",
		"UID": 2
	}, {
		"ID": 4,
		"CreatedAt": "2019-04-08T22:45:00+08:00",
		"UpdatedAt": "2019-04-08T22:45:00+08:00",
		"DeletedAt": null,
		"Content": "messageEdited",
		"UID": 2
	}, {
		"ID": 5,
		"CreatedAt": "2019-04-08T22:45:08+08:00",
		"UpdatedAt": "2019-04-08T22:45:08+08:00",
		"DeletedAt": null,
		"Content": "messageEdited",
		"UID": 2
	}, {
		"ID": 6,
		"CreatedAt": "2019-04-09T14:59:05+08:00",
		"UpdatedAt": "2019-04-09T14:59:05+08:00",
		"DeletedAt": null,
		"Content": "测试一下根据留言内容查找",
		"UID": 2
	}],
	"error": 0,
	"msg": "success"
}
```

### GET `/message/id/{id}` 通过id查找留言信息

Param

- `id`:(int)留言id(mid)

Response

```json
{
	"data": {
		"ID": 7,
		"CreatedAt": "2019-04-09T16:36:34+08:00",
		"UpdatedAt": "2019-04-09T16:36:34+08:00",
		"DeletedAt": null,
		"Content": "留言内容",
		"UID": 2
	},
	"error": 0,
	"msg": "success"
}
```

### GET `/message/uid/{uid}` 通过uid查找某用户所有留言

Param

- `uid`:(int)用户id

Response

```json
{
	"data": [{
		"ID": 7,
		"CreatedAt": "2019-04-09T16:36:34+08:00",
		"UpdatedAt": "2019-04-09T16:36:34+08:00",
		"DeletedAt": null,
		"Content": "留言内容",
		"UID": 2
	}, {
		"ID": 8,
		"CreatedAt": "2019-04-09T16:43:10+08:00",
		"UpdatedAt": "2019-04-09T16:43:10+08:00",
		"DeletedAt": null,
		"Content": "我又留言了",
		"UID": 2
	}, {
		"ID": 9,
		"CreatedAt": "2019-04-09T16:43:17+08:00",
		"UpdatedAt": "2019-04-09T16:43:17+08:00",
		"DeletedAt": null,
		"Content": "哈哈哈还是我",
		"UID": 2
	}],
	"error": 0,
	"msg": "success"
}
```

### GET `/message/content/{content} `根据留言内容查询

Param

* `content`:(string)留言部`分内容

Get `http://127.0.0.1:8080/message/content/测试`
Response

```json
{
	"data": [{
		"ID": 6,
		"CreatedAt": "2019-04-09T14:59:05+08:00",
		"UpdatedAt": "2019-04-09T14:59:05+08:00",
		"DeletedAt": null,
		"Content": "测试一下根据留言内容查找",
		"UID": 2
	}],
	"error": 0,
	"msg": "success"
}
```

### PUT `/message` 修改留言信息



### DELETE `/message/id/{id}` 通过留言id删除留言

Param

- `id`:(int)留言id

Success

HttpStatusCode:`200`



### DELETE `/user/uid/{uid}` 删除某人所有留言

Param

- `uid`:(int)用户id

Success

HttpStatusCode:`200`



## 错误对照

| HttpStatusCode | Error | Msg                | Meaning       |
| -------------- | ----- | ------------------ | ------------- |
| 400            | 40000 | bad param          | url入参错误   |
| 400            | 40001 | bad payload        | payload不完整 |
| 404            | 40400 | not found          | 留言不存在    |

