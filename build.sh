"""
A build script for golang projects
"""
APP_NAME=findreplace

if [[ $(uname -m) == "x86_64" ]]; then
    APP_ARCH=amd64
elif [[ $(uname -m) == "i686" ]]; then
    APP_ARCH=386
elif [[ $(uname -m) == "arm64" ]]; then
    APP_ARCH=arm64
else
    echo "Unsupported architecture: $(uname -m)"
    exit 1
fi

INSTALL=false
for (( i=0; i <= "$#"; i++ )) ; do
    if [[ ${!i} == "--install" ]] ; then
        INSTALL=true
    fi

    if [[ ${!i} == "--version" ]]; then
        VERSION=${@:$i+1:1}
        echo "Building version: $APP_NAME $VERSION"
    fi

    if [[ ${!i} == "--os-type" ]] ; then
        OS=${@:$i+1:1}
        echo "Building for OS: $OS"
    fi
done

if [[ -z $OS ]]; then
    if [[ "$OSTYPE" == "linux-gnu"* ]]; then
        OS=linux
    elif [[ "$OSTYPE" == "darwin"* ]]; then
        OS=darwin
    elif [[ "$OSTYPE" == "cygwin" ]]; then
        OS=cygwin
    elif [[ "$OSTYPE" == "msys" ]]; then
        # Lightweight shell and GNU utilities compiled for Windows (part of MinGW)
        OS=mysys
    elif [[ "$OSTYPE" == "win32" ]]; then
        OS=windows
    elif [[ "$OSTYPE" == "freebsd"* ]]; then
        OS=freebsd
    else
        OS=unknown
    fi
    echo "Auomatically detected OS: $OS"
fi


echo Building for arch: ${OS}_${APP_ARCH}
BIN_FILENAME=bin/${APP_NAME}_${OS}_${APP_ARCH}
echo FILENAME = $BIN_FILENAME
echo "Executing tests with: go test ./..."
GOOS=$OS GOARCH=$APP_ARCH go test ./...
TEST_STATUS=$?
if [[ $TEST_STATUS -ne 0 ]]; then
    echo "!!!!!!!!!WARNING: TESTS FAILED!!!!!!!!!"
    echo "!!!!!!!!!WARNING: TESTS FAILED!!!!!!!!!"
    echo "!!!!!!!!!WARNING: TESTS FAILED!!!!!!!!!"
    echo "!!!!!!!!!WARNING: TESTS FAILED!!!!!!!!!"
fi
GOOS=$OS GOARCH=$APP_ARCH go build -o $BIN_FILENAME -ldflags "-X main.version=$VERSION" main.go
chmod +x $BIN_FILENAME
echo "Build finished: $BIN_FILENAME"
if [ $INSTALL == "true" ]; then
    echo "Installing..."
    INSTALL_FILEPATH=/usr/local/bin/$APP_NAME
    echo "Installing to $INSTALL_FILEPATH"
    cp $BIN_FILENAME $INSTALL_FILEPATH
    echo "Installation complete!"
else
    echo
    echo "Next time, run with the --install flag to install to /usr/local/bin"
fi