---
title: hospital v1.0.0
language_tabs:
  - shell: Shell
  - http: HTTP
  - javascript: JavaScript
  - ruby: Ruby
  - python: Python
  - php: PHP
  - java: Java
  - go: Go
toc_footers: []
includes: []
search: true
code_clipboard: true
highlight_theme: darkula
headingLevel: 2
generator: "@tarslib/widdershins v4.0.17"

---

# hospital

> v1.0.0

Base URLs:

* <a href="http://127.0.0.1:8888">开发环境: http://127.0.0.1:8888</a>

# UserController

## POST Login

POST /user/login

> Body 请求参数

```json
{
  "id": 0,
  "phone": "string",
  "password": "string"
}
```

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|body|body|[User](#schemauser)| 否 |none|

> 返回示例

> 成功

```json
{
  "success": false,
  "code": 0,
  "message": "",
  "data": {
    "": {}
  }
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|[Result](#schemaresult)|

## GET Info

GET /user/info

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|token|query|string| 否 |none|

> 返回示例

> 成功

```json
{
  "success": false,
  "code": 0,
  "message": "",
  "data": {
    "": {}
  }
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|[Result](#schemaresult)|

## POST 注册

POST /user/register

> Body 请求参数

```json
{
  "name": "string",
  "sex": 0,
  "birthday": "string",
  "idNum": "string",
  "phone": "string",
  "fee": 0,
  "introduce": "string",
  "departId": 0,
  "password": "string"
}
```

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|body|body|[RegisterDTO](#schemaregisterdto)| 否 |none|

> 返回示例

> 成功

```json
{
  "success": false,
  "code": 0,
  "message": "",
  "data": {
    "": {}
  }
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|[Result](#schemaresult)|

## POST 更换头像

POST /user/updateAvatar

> Body 请求参数

```yaml
file: string

```

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|body|body|object| 否 |none|
|» file|body|string(binary)| 是 |none|

> 返回示例

> 成功

```json
{
  "success": false,
  "code": 0,
  "message": "",
  "data": {
    "": {}
  }
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|[Result](#schemaresult)|

## GET Chart

GET /user/list

> 返回示例

> 成功

```json
{
  "success": false,
  "code": 0,
  "message": "",
  "data": {
    "": {}
  }
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|[Result](#schemaresult)|

# <p>

## GET All

GET /room/list

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|searchName|query|string| 否 |none|
|pageNum|query|integer| 否 |none|
|pageSize|query|integer| 否 |none|
|roomId|query|integer| 否 |none|
|departId|query|integer| 否 |none|

> 返回示例

> 成功

```json
{
  "success": false,
  "code": 0,
  "message": "",
  "data": {
    "": {}
  }
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|[Result](#schemaresult)|

## POST Del

POST /room/delete

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|id|query|integer| 是 |none|

> 返回示例

> 成功

```json
{
  "success": false,
  "code": 0,
  "message": "",
  "data": {
    "": {}
  }
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|[Result](#schemaresult)|

## POST Add

POST /room/add

> Body 请求参数

```json
{
  "roomId": 0,
  "roomName": "string",
  "roomIntroduce": "string",
  "departId": 0
}
```

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|body|body|[Room](#schemaroom)| 否 |none|

> 返回示例

> 成功

```json
{
  "success": false,
  "code": 0,
  "message": "",
  "data": {
    "": {}
  }
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|[Result](#schemaresult)|

## POST Update

POST /room/update

> Body 请求参数

```json
{
  "roomId": 0,
  "roomName": "string",
  "roomIntroduce": "string",
  "departId": 0
}
```

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|body|body|[Room](#schemaroom)| 否 |none|

> 返回示例

> 成功

```json
{
  "success": false,
  "code": 0,
  "message": "",
  "data": {
    "": {}
  }
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|[Result](#schemaresult)|

## GET All

GET /depart/list

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|searchName|query|string| 否 |none|
|pageNum|query|integer| 否 |none|

> 返回示例

> 成功

```json
{
  "success": false,
  "code": 0,
  "message": "",
  "data": {
    "": {}
  }
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|[Result](#schemaresult)|

## POST Del

POST /depart/delete

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|id|query|integer| 是 |none|

> 返回示例

> 成功

```json
{
  "success": false,
  "code": 0,
  "message": "",
  "data": {
    "": {}
  }
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|[Result](#schemaresult)|

## POST Add

POST /depart/add

> Body 请求参数

```json
{
  "departId": 0,
  "departName": "string",
  "departIntroduce": "string"
}
```

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|body|body|[Depart](#schemadepart)| 否 |none|

> 返回示例

> 成功

```json
{
  "success": false,
  "code": 0,
  "message": "",
  "data": {
    "": {}
  }
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|[Result](#schemaresult)|

## POST Update

POST /depart/update

> Body 请求参数

```json
{
  "departId": 0,
  "departName": "string",
  "departIntroduce": "string"
}
```

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|body|body|[Depart](#schemadepart)| 否 |none|

> 返回示例

> 成功

```json
{
  "success": false,
  "code": 0,
  "message": "",
  "data": {
    "": {}
  }
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|[Result](#schemaresult)|

# DoctorController

## GET All

GET /doctor/list

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|name|query|string| 否 |none|
|pageNum|query|string| 否 |none|
|pageSize|query|string| 否 |none|
|id|query|string| 否 |none|
|idNum|query|string| 否 |none|
|phone|query|string| 否 |none|
|fee|query|string| 否 |none|
|departId|query|string| 否 |none|
|role|query|string| 否 |none|

> 返回示例

> 成功

```json
{
  "success": false,
  "code": 0,
  "message": "",
  "data": {
    "": {}
  }
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|[Result](#schemaresult)|

## POST All

POST /doctor/list

> Body 请求参数

```json
{
  "name": "string",
  "pageNum": 0,
  "pageSize": 0,
  "id": 0,
  "idNum": "string",
  "phone": "string",
  "fee": 0,
  "departId": 0,
  "role": 0
}
```

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|body|body|[DoctorSearchDto](#schemadoctorsearchdto)| 否 |none|

> 返回示例

> 成功

```json
{
  "success": false,
  "code": 0,
  "message": "",
  "data": {
    "": {}
  }
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|[Result](#schemaresult)|

## POST add

POST /doctor/add

> Body 请求参数

```json
{
  "name": "string",
  "sex": 0,
  "birthday": "string",
  "idNum": "string",
  "phone": "string",
  "fee": 0,
  "introduce": "string",
  "departId": 0,
  "password": "string"
}
```

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|body|body|[RegisterDTO](#schemaregisterdto)| 否 |none|

> 返回示例

> 成功

```json
{
  "success": false,
  "code": 0,
  "message": "",
  "data": {
    "": {}
  }
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|[Result](#schemaresult)|

## POST delete

POST /doctor/delete

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|id|query|integer| 是 |none|

> 返回示例

> 成功

```json
{
  "success": false,
  "code": 0,
  "message": "",
  "data": {
    "": {}
  }
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|[Result](#schemaresult)|

## POST update

POST /doctor/update

> Body 请求参数

```json
{
  "id": 0,
  "name": "string",
  "sex": 0,
  "birthday": "string",
  "idNum": "string",
  "phone": "string",
  "fee": 0,
  "introduce": "string",
  "departId": 0,
  "role": 0
}
```

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|body|body|[DoctorDTO](#schemadoctordto)| 否 |none|

> 返回示例

> 成功

```json
{
  "success": false,
  "code": 0,
  "message": "",
  "data": {
    "": {}
  }
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|[Result](#schemaresult)|

# 数据模型

<h2 id="tocS_DoctorDTO">DoctorDTO</h2>

<a id="schemadoctordto"></a>
<a id="schema_DoctorDTO"></a>
<a id="tocSdoctordto"></a>
<a id="tocsdoctordto"></a>

```json
{
  "id": 0,
  "name": "string",
  "sex": 0,
  "birthday": "string",
  "idNum": "string",
  "phone": "string",
  "fee": 0,
  "introduce": "string",
  "departId": 0,
  "role": 0
}

```

### 属性

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|id|integer|false|none||none|
|name|string|false|none||none|
|sex|integer|false|none||none|
|birthday|string|false|none||none|
|idNum|string|false|none||none|
|phone|string|false|none||none|
|fee|number|false|none||none|
|introduce|string|false|none||none|
|departId|integer|false|none||none|
|role|integer|false|none||none|

<h2 id="tocS_DoctorSearchDto">DoctorSearchDto</h2>

<a id="schemadoctorsearchdto"></a>
<a id="schema_DoctorSearchDto"></a>
<a id="tocSdoctorsearchdto"></a>
<a id="tocsdoctorsearchdto"></a>

```json
{
  "name": "string",
  "pageNum": 0,
  "pageSize": 0,
  "id": 0,
  "idNum": "string",
  "phone": "string",
  "fee": 0,
  "departId": 0,
  "role": 0
}

```

### 属性

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|name|string|false|none||none|
|pageNum|integer|false|none||none|
|pageSize|integer|false|none||none|
|id|integer|false|none||none|
|idNum|string|false|none||none|
|phone|string|false|none||none|
|fee|number|false|none||none|
|departId|integer|false|none||none|
|role|integer|false|none||none|

<h2 id="tocS_DoctorDto">DoctorDto</h2>

<a id="schemadoctordto"></a>
<a id="schema_DoctorDto"></a>
<a id="tocSdoctordto"></a>
<a id="tocsdoctordto"></a>

```json
{
  "name": "string",
  "pageNum": 0,
  "pageSize": 0,
  "id": 0,
  "idNum": "string",
  "phone": "string",
  "fee": 0,
  "departId": 0,
  "role": 0
}

```

### 属性

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|name|string|false|none||none|
|pageNum|integer|false|none||none|
|pageSize|integer|false|none||none|
|id|integer|false|none||none|
|idNum|string|false|none||none|
|phone|string|false|none||none|
|fee|number|false|none||none|
|departId|integer|false|none||none|
|role|integer|false|none||none|

<h2 id="tocS_RegisterDTO">RegisterDTO</h2>

<a id="schemaregisterdto"></a>
<a id="schema_RegisterDTO"></a>
<a id="tocSregisterdto"></a>
<a id="tocsregisterdto"></a>

```json
{
  "name": "string",
  "sex": 0,
  "birthday": "string",
  "idNum": "string",
  "phone": "string",
  "fee": 0,
  "introduce": "string",
  "departId": 0,
  "password": "string"
}

```

### 属性

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|name|string|false|none||医生姓名|
|sex|integer|false|none||性别 0:女，1:男，2：其他|
|birthday|string|false|none||出生日期|
|idNum|string|false|none||身份证号|
|phone|string|false|none||手机号码|
|fee|number|false|none||挂号费|
|introduce|string|false|none||基本介绍|
|departId|integer|false|none||科室ID|
|password|string|false|none||登陆密码|

<h2 id="tocS_Depart">Depart</h2>

<a id="schemadepart"></a>
<a id="schema_Depart"></a>
<a id="tocSdepart"></a>
<a id="tocsdepart"></a>

```json
{
  "departId": 0,
  "departName": "string",
  "departIntroduce": "string"
}

```

### 属性

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|departId|integer|false|none||科室ID|
|departName|string|false|none||科室名称|
|departIntroduce|string|false|none||科室介绍|

<h2 id="tocS_Room">Room</h2>

<a id="schemaroom"></a>
<a id="schema_Room"></a>
<a id="tocSroom"></a>
<a id="tocsroom"></a>

```json
{
  "roomId": 0,
  "roomName": "string",
  "roomIntroduce": "string",
  "departId": 0
}

```

### 属性

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|roomId|integer|false|none||诊室ID|
|roomName|string|false|none||诊室名称|
|roomIntroduce|string|false|none||诊室介绍|
|departId|integer|false|none||隶属科室ID|

<h2 id="tocS_Result">Result</h2>

<a id="schemaresult"></a>
<a id="schema_Result"></a>
<a id="tocSresult"></a>
<a id="tocsresult"></a>

```json
{
  "success": true,
  "code": 0,
  "message": "string",
  "data": {
    "key": {}
  }
}

```

### 属性

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|success|boolean|false|none||none|
|code|integer|false|none||none|
|message|string|false|none||none|
|data|[Map«Object»](#schemamap%c2%abobject%c2%bb)|false|none||java.util.Map<java.lang.String,java.lang.Object>|

<h2 id="tocS_Map«Object»">Map«Object»</h2>

<a id="schemamap«object»"></a>
<a id="schema_Map«Object»"></a>
<a id="tocSmap«object»"></a>
<a id="tocsmap«object»"></a>

```json
{
  "key": {}
}

```

### 属性

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|key|[key](#schemakey)|false|none||data.key|

<h2 id="tocS_key">key</h2>

<a id="schemakey"></a>
<a id="schema_key"></a>
<a id="tocSkey"></a>
<a id="tocskey"></a>

```json
{}

```

### 属性

*None*

<h2 id="tocS_User">User</h2>

<a id="schemauser"></a>
<a id="schema_User"></a>
<a id="tocSuser"></a>
<a id="tocsuser"></a>

```json
{
  "id": 0,
  "phone": "string",
  "password": "string"
}

```

### 属性

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|id|integer|false|none||none|
|phone|string|false|none||none|
|password|string|false|none||none|

<h2 id="tocS_SaResult">SaResult</h2>

<a id="schemasaresult"></a>
<a id="schema_SaResult"></a>
<a id="tocSsaresult"></a>
<a id="tocssaresult"></a>

```json
{
  "key": {}
}

```

### 属性

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|key|object|false|none||none|

