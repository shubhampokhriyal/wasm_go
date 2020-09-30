# About
A go wasm example

# Steps to run:
1. Run 'GOARCH=wasm GOOS=js go build -o html/lib.wasm main.go' to create the lib.wasm file.
2. Run 'cp "$(go env GOROOT)/misc/wasm/wasm_exec.js" html' to copy the wasm_exec.js for your particular go version.
3. Serve the html file. (Can use npm serve library to serve).
