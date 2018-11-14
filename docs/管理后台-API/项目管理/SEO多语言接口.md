    
**简要描述：** 

- SEO多语言接口

**请求URL：** 
- ` http://domain/sites/seo `
  
**请求方式：**
- POST 

**参数：** 

|参数名|必选|类型|说明|
|:----    |:---|:----- |-----   |
|lang |是  |array | 支持语言 |

 **请求示例**
``` 
{
    "siteId": 2,
    "lang": [
        {
            "lang_id": 1,
            "title": "League of Angels III",
            "keyword": "League of Angels III",
            "desc": "League of Angels III"
        },
        {
            "lang_id": 2,
            "title": "女神聯盟III",
            "keyword": "女神聯盟III",
            "desc": "女神聯盟III"
        },
        {
            "lang_id": 3,
            "title": "女神联盟III",
            "keyword": "女神联盟III",
            "desc": "女神联盟III"
        }
    ]
}
```
 **返回示例**

``` 
{
    "code": 0,
    "msg": "Sucess",
    "data": []
}
```

 **返回参数说明** 

|参数名|类型|说明|
|:-----  |:-----|-----                           |
|code |int   |状态  |
|msg  |strint   |消息  |
|data |array   |数据  |

 **备注** 

- 更多返回错误代码请看首页的错误代码描述











