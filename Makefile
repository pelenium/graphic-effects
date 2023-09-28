spiralBuild: ./spiral
	go build -C spiral -o spiral.exe -ldflags '-extldflags "-static" -H=windowsgui'

plasmBuild: ./plasm
	go build -C plasm -o plasm.exe -ldflags '-extldflags "-static" -H=windowsgui'

ray-castingBuild: ./ray-casting
	go build -C ray-casting -o ray-casting.exe -ldflags '-extldflags "-static" -H=windowsgui'