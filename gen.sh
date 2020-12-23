(
flatc --go --go-namespace a -o internal/fb ./schema/thingA.fbs;
flatc --go --go-namespace b -o internal/fb ./schema/thingB.fbs;
echo gopath:$GOPATH;
)