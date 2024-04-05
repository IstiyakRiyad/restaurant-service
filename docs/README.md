# API endpoints

Here the endpoints of the pathao technical assessment

## GET
`1. Restaurant at certain time` [/api/v1/restaurant/datetime](#) <br/>
`2. Filter with price range & no of dishes` [/api/v1/restaurant](#) <br/>
`3. Search restaurant or dishes with name` [/api/v1/search](#) <br/>

## POST
`4. Purchasing a dish` [/api/v1/purchase](#) <br/>

## Some Extra Endpoints for debug
`GET` [/api/v1/restaurant/:id](#) <br/>
`GET` [/api/v1/user/:id](#) <br/>
`GET` [/api/v1/user](#) <br/>

___

### `GET` `/api/v1/restaurant/datetime`
List all restaurants that are open at a certain datetime 

**Query Parameters**

|          Name | Required |  Type   | Description                                                                      |
| -------------:|:--------:|:-------:| -------------------------------------------------------------------------------- |
| `date_time` | optional | datetime  |  Datetime for query. Default is the current time  |
| `organization_id` | optional | string  |  organization before returning any information.    |


**Response**

```json
{
	"message": "List of restaurants",
	"data": [
		{
			"ID": 1871,
			"Name": "The Corner Office Restaurant and Martini Bar Denver",
			"CashBalance": 3234.57
		},
		{
			"ID": 2190,
			"Name": "Zodiac at Neiman Marcus – Downtown Dallas",
			"CashBalance": 1545.59
		}
	]
}
```
___

### `GET` `/api/v1/restaurant`
List top y restaurants that have more or less than x number of dishes within a price range, ranked alphabetically. More or less (than x) is a parameter that the API allows the consumer to enter.

**Query Parameters**

|          Name | Required |  Type   | Description                                                                      |
| -------------:|:--------:|:-------:| -------------------------------------------------------------------------------- |
| `limit`       | required | int    |  Number of restaurant `y`  |
| `base_count`  | required | int    |  Number of dishes `x`  |
| `base_type`   | required | enum   |  `more` or `less`  |
| `min_price`   | optional | float  |  Minimum price of dishes  |
| `max_price`   | optional | float  |  Minimum price of dishes  |


**Response**

```json
{
	"message": "List of restaurants",
	"data": [
		{
			"ID": 16,
			"Name": "24 Plates",
			"CashBalance": 4869.25
		},
		{
			"ID": 1823,
			"Name": "Texas de Brazil - Denver",
			"CashBalance": 3553.05
		}
	]
}
```
___
