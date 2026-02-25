module github.com/poymanov/codemania-task-board/notification

go 1.25.5

replace github.com/poymanov/codemania-task-board/shared => ../shared

replace github.com/poymanov/codemania-task-board/platform => ../platform

require (
	github.com/caarlos0/env/v11 v11.4.0
	github.com/joho/godotenv v1.5.1
	github.com/poymanov/codemania-task-board/platform v0.0.0-00010101000000-000000000000
	github.com/rs/zerolog v1.34.0
	github.com/segmentio/kafka-go v0.4.50
)

require (
	github.com/klauspost/compress v1.18.3 // indirect
	github.com/mattn/go-colorable v0.1.14 // indirect
	github.com/mattn/go-isatty v0.0.20 // indirect
	github.com/pierrec/lz4/v4 v4.1.22 // indirect
	github.com/stretchr/testify v1.11.1 // indirect
	golang.org/x/sys v0.40.0 // indirect
)
