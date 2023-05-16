.generate:
	echo "Generating python"
	python -m grpc_tools.protoc -I=api/ --python_out=. --grpc_python_out=. led-service.proto

	echo "Generating golang"
	protoc -I=api/ --go_out=. --go-grpc_out=. led-service.proto

generate: .generate

.run:
	sudo python main.py

run: .run
