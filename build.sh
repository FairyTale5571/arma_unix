echo "Start building extension arma_unix"
echo "arma_unix_x64.dll in progress..."
GOARCH=amd64 go build -o release/arma_unix_x64.dll -buildmode=c-shared .
echo "arma_unix_x64.dll builded"

echo "arma_unix.dll in progress..."
GOARCH=386 CGO_ENABLED=1 go build -o release/arma_unix.dll -buildmode=c-shared .
echo "arma_unix.dll builded"

rm release/*.h
echo "Auto-generated headers removed"

echo "Building done, find dll's in release folder"