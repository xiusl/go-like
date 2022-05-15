# Api

## base 
- auth

  `Authorization`: `Bearer $token`
- response
  ```javascript
  {
    "data": any,
    "code": 200,
    "message": ""
  }
  ```

## get verify code

- Path: `/verifycode`
- Method: `GET`
- Query Params:
  | Name | Type | Required |Desc|
  |---|---|---|--|
  | key    |string | true | mobile/email |
  | biz_type | int | false | type of biz |
- Response:
  ```javascript
  {
	"data": true/false,
	"code": 200,
	"message": ""
  }
  ```

## third
