protoa:
	@protoc -I ./payment-grpc/proto -I ./auth-grpc/proto --go_out=./auth-grpc/proto --go_opt=paths=source_relative \
	--go-grpc_out=./auth-grpc/proto --go-grpc_opt=paths=source_relative \
	auth-grpc/proto/auth.proto

mockClient:
	mockgen github.com/Edbeer/payment-proto/auth-grpc/proto AuthServiceClient > auth-grpc/client/mock/client_mock.go

mockGetAccount:
	mockgen github.com/Edbeer/payment-proto/auth-grpc/proto AuthService_GetAccountClient,AuthService_GetAccountServer > auth-grpc/client/mock/get_account_mock.go

mockGetStat:
	mockgen github.com/Edbeer/payment-proto/auth-grpc/proto AuthService_GetStatementClient,AuthService_GetStatementServer > auth-grpc/client/mock/get_statetement_mock.go

mockCreateStat:
	mockgen github.com/Edbeer/payment-proto/auth-grpc/proto AuthService_CreateStatementClient,AuthService_CreateStatementServer > auth-grpc/client/mock/create_statetemet_mock.go

mocka: mockClient mockGetAccount mockCreateStat mockGetStat

protop:
	@protoc -I ./payment-grpc/proto --go_out=./payment-grpc/proto --go_opt=paths=source_relative \
	--go-grpc_out=./payment-grpc/proto --go-grpc_opt=paths=source_relative \
	payment-grpc/proto/*.proto

mockp:
	mockgen github.com/Edbeer/payment-proto/payment-grpc/proto PaymentServiceServer > payment-grpc/client/mock/client_mock.go