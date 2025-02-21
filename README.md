# Render-MD: A Client-Side Markdown Editor
**Render-MD** is a lightweight client-side Markdown editor built with Go + WebAssembly. It provides real-time Markdown preview with syntax highlighting, making it a powerful tool for writing and formatting Markdown directly in the browser.

This is a learning project to explore Go WebAssembly and how it can be used for frontend applications.

## Installation & Running Locally
1️⃣ Clone the Repository
```sh
git clone https://github.com/your-username/render-md.git
cd render-md
```
2️⃣ Build WebAssembly (WASM)
Run the following command to compile the Go code into a .wasm file:
```sh
GOOS=js GOARCH=wasm go build -o main.wasm main.go
```
3️⃣ Copy WebAssembly Runtime File
Copy wasm_exec.js (required to run WASM in the browser):
```sh
cp "$(go env GOROOT)/misc/wasm/wasm_exec.js" .
```
4️⃣ Start a Local Server
You cannot open the HTML file directly due to WebAssembly security restrictions. Use a local web server:
```sh
python3 -m http.server 8080
#or
goexec 'http.ListenAndServe(":8080", http.FileServer(http.Dir(".")))'
```
Then open: http://localhost:8080/

## Technologies Used
* Go (for WebAssembly)
* Goldmark (Markdown parser)
* highlight.js (Syntax highlighting)
* HTML/CSS/JavaScript (Frontend)
* GitHub Pages (Hosting)

## License
This project is open-source and available under the MIT License.