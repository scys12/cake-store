# cake-store

## Instructions

1. Install Go & Docker
2. Clone this project
3. Build the project with

```
docker-compose up -d
```

4. Access through `localhost:8000/`

If you want to rebuild the project, use

```
docker-compose up -d --build
```

To destroy the project, use

```
docker-compose down -v
```

This project use `grafana` and `prometheus` for system monitoring and alerting. Read official documentation [here](https://prometheus.io/) about prometheus and [here](https://grafana.com) for grafana. There are 3 metrics that are measured, namely latency, total request and response status. Read more [here](https://github.com/scys12/cake-store/tree/master/pkg/monitoring).

## Unit Test

To check how muc coverage, run:

```
make create_cover_test
```

Explore Makefile file to know the command to visualize the coverage in HTML.

## API
### List of Cake

#### Request
`GET /cakes/`

    curl -i -H 'Accept: application/json' http://localhost:8000/cakes/

#### Response

    {
      "data": [
        {
          "id": 1,
          "title": "Lemon cheesecake",
          "description": "A cheesecake made of lemon",
          "rating": 7,
          "image": "https://img.taste.com.au/ynYrqkOs/w720-h480-cfill-q80/taste/2016/11/sunny-lemon-cheesecake-102220-1.jpeg",
          "created_at": "2020-02-01 10:56:31",
          "updated_at": "2020-02-13 09:30:23"
        }, ...
      ],
      "total_data": 6,
      "total_data_per_page":5
    }
### Detail of Cake 

#### Request
`GET /cakes/{id}`
      
      curl -i -H 'Accept: application/json' http://localhost:8000/cakes/5
    
 #### Response
      {
        "data": {
          "id": 1,
          "title": "Lemon cheesecake",
          "description": "A cheesecake made of lemon",
          "rating": 7.0,
          "image": "https://img.taste.com.au/ynYrqkOs/w720-h480-cfill-q80/taste/2016/11/sunny-lemon-cheesecake-102220-1.jpeg",
          "created_at": "2020-02-01 10:56:31",
          "updated_at": "2020-02-13 09:30:23"
        }
      }

### Add New Cake
`POST /cakes`
     
     curl -i -H 'Accept: application/json' -X POST http://localhost:8000/cakes
#### Request Body
      {
        "title": "Lemon cheesecake",
        "description": "A cheesecake made of lemon",
        "rating": 7,
        "image": "https://img.taste.com.au/ynYrqkOs/w720-h480-cfill-q80/taste/2016/11/sunny-lemon-cheesecake-102220-1.jpeg"
      }
#### Response
      {
        "data": {
          "id": 1,
          "title": "Lemon cheesecake",
          "description": "A cheesecake made of lemon",
          "rating": 7,
          "image": "https://img.taste.com.au/ynYrqkOs/w720-h480-cfill-q80/taste/2016/11/sunny-lemon-cheesecake-102220-1.jpeg"
        }
      }

### Update Cake
`PATCH /cakes/{id}`

     curl -i -H 'Accept: application/json' -X PATCH http://localhost:8000/cakes/1
#### Request Body
      {
        "title": "Lemon cheesecake",
        "description": "A cheesecake made of lemon",
        "rating": 8,
        "image": "https://img.taste.com.au/ynYrqkOs/w720-h480-cfill-q80/taste/2016/11/sunny-lemon-cheesecake-102220-1.jpeg"
      }
#### Response
      {
        "data": {
          "id": 1,
          "title": "Lemon cheesecake",
          "description": "A cheesecake made of lemon",
          "rating": 8,
          "image": "https://img.taste.com.au/ynYrqkOs/w720-h480-cfill-q80/taste/2016/11/sunny-lemon-cheesecake-102220-1.jpeg"
        }
      }
      
### Delete Cake
`DELETE /cakes/{id}`

     curl -i -H 'Accept: application/json' -X DELETE http://localhost:8000/cakes/5
     
#### Response
      {
          "data": {
              "id": 1,
          }
      }
