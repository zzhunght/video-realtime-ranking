# Realtime Video Ranking

### Description
- Hệ thống mô phỏng xếp hạng video realtime theo tương tác của người dùng

### High Level Design
![alt text](image.png)


- Dùng Redis (ZSet) để lưu ranking, (HSET) để lưu metadata (low latency)
- Postgres làm db chính nơi lưu trữ metadata của video
- Kafka làm Message Queue để xử lí event theo bất đồng bộ (dễ dàng scale khi có lượng lớn user)
- Sử dụng EDA parttern

### Note

### Code Struct

```
project-name/
├── cmd/
│   └──  main.go                 

├── internal/
│   ├── application/                # Application layer
│   |   |── video_service.go        # Xử lí bussiness logic
│   |   └── ranking_service.go
|   |
|   ├── docs/                      #API doc
│   |   ├── api     
│   │   │   ├── docs.go             
│   │   │   ├── swagger.json             
│   │   │   ├── swagger.yaml        
|   |
|   ├── config/
│   |   ├── config.go                   # Config
|   |
│   ├── domain/                     # Domain layer
│   │   ├── models/                 # Entity liên quan đến ranking
│   │   │   ├── category.go
│   │   │   ├── event.go
│   │   │   └── score.go            
│   │   │   └── user.go            
│   │   │   └── video.go            
│   │   ├── repository/             # Interface cho repositories
│   │   │   ├── cached_repository.go
│   │   │   ├── ranking_repository.go
│   │   │   └── video_repository.go
│   ├── infrastructure/             # Infrastructure layer
│   │   ├── messaging/              # Kafka
│   │   │   ├── kafka_producer.go        
│   │   │   ├── score_consumer.go  
│   │   ├── persistence/
│   │   │   ├── postgres/         
│   │   │   │   ├── connection.go
│   │   │   │   ├── postgres_video_repository.go 
│   │   │   └── redis/              
│   │   │       ├── connection.go
│   │   │       ├── ranking_repository.go  # Chỗ này dùng để get và add score video
│   │   │       └── cached_repository.go   # Dùng để get video metadata trong cached
|   |   
│   └── interfaces/                 # Interface layer
│       ├── api/
│       │   ├── dto/              
│       │   │   ├── event.go
│       │   ├── handler/            # API handler  
│       │   │   ├── ranking_handler.go
│       │   │   ├── handler_helper.go
│       │   ├── responses/              
│       │   │   ├── response.go
│       │   ├── router/             # API router
│       │   │   ├── router.go
│
├── pkg/                           # Mã dùng chung, tiện ích
│   └── errors/
│       └── errors.go
│
├── scripts/              
│   ├── mock.go  #Mock scripts tự động call api add event
│
├── config.yaml                    # Env
├── go.mod
├── go.sum
├── Dockerfile                     # Docker file cho app
├── Dockerfile.mock                # Docker file cho scripts
├── docker-compose.yml             # Setup Postgres & Redis
└── readme.md
```


### Cách chạy ứng dụng
Sử dụng docker compose
```
docker compose up
```

Khi chạy docker xong sẽ tự động chạy 1 mock scripts mô phỏng lại user tương tác với video
và cập nhật score liên tục

[Truy cập vào và test API tại đây](http://localhost:8080/swagger/index.html)
