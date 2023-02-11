.generate:
	python -m grpc_tools.protoc -I=api/ --python_out=. --grpc_python_out=. led-api.proto

generate: .generate

.run:
	sudo python main.py

run: .run
