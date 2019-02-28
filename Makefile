goGen:
	@protoc -I proto/ proto/game.proto --go_out=plugins=grpc:game

dartGen: 
	@protoc -I=proto --dart_out=grpc:lib/src/generated proto/game.proto
