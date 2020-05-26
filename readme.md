# Day 4: Build a Cash Register App

## Instructions

Build a web api that fullfills the following 3 endpoints.

`POST request localhost:8080/add`
```json
{
    "item": "Toilet Paper",
    "price_per_unit": 20.00,
    "units": 2.0
}
```

`POST response localhost:8080/add`
```json
{
    "items": [
        {
            "item": "Toiler Paper", "cost_per_unit": 20.00,
            "units": 2.0,
            "total_cost": 40.00
        }
    ],
    "totals": {
        "items_count": 1,
        "cost": 40.00
    }
}
```

Items should be held in the order that they were recieved. Units must be floats to accomodate to items that are sold be weight or other non-integer quantities.

`GET response localhost:8080/list`
```json
{
    "items": [
        {
            "item": "Toiler Paper", "cost_per_unit": 20.00,
            "units": 2.0,
            "total_cost": 40.00
        }
    ],
    "totals": {
        "items_count": 1,
        "cost": 40.00
    }
}
```

Same output as above, except no form values required.

`POST response localhost:8080/clear`

Clears the cart, returns a 204.
