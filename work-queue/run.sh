# Run producer
go run errorHandler.go billModel.go producer.go

# Run consumer
go consumer.go billModel.go errorHandler.go