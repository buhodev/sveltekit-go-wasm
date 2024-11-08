# sv

Everything you need to build a Svelte project, powered by [`sv`](https://github.com/sveltejs/cli).

## Creating a project

If you're seeing this, you've probably already done this step. Congrats!

```bash
# create a new project in the current directory
npx sv create

# create a new project in my-app
npx sv create my-app
```

## Developing

Once you've created a project and installed dependencies with `npm install` (or `pnpm install` or `yarn`), start a development server:

```bash
npm run dev

# or start the server and open the app in a new browser tab
npm run dev -- --open
```

Now, you'll need to do the following steps locally on your machine:

Create a `static/wasm` directory in your project and copy the `wasm_exec.js` file from your Go installation:

```sh
mkdir -p static/wasm
cp "$(go env GOROOT)/misc/wasm/wasm_exec.js" static/wasm/
```

Compile the Go code to WebAssembly (from the `go` directory):

```sh
GOOS=js GOARCH=wasm go build -o ../static/wasm/main.wasm main.go
```

The application expects a CSV file with prices in the first column. For example:

```csv
10.50
25.75
30.00
```

Key features of this setup:

- The page has a file input that accepts CSV files
- The Go WASM module processes the CSV file locally (no server upload needed)
- The total is displayed in the UI with 2 decimal places
- Error handling is implemented both in the UI and WASM module
- The WASM module is loaded only once when the page mounts
- After compiling the WASM module locally and placing it in the correct directory, the application will be ready to use. Upload a CSV file and the total will be calculated and displayed instantly.

## Building

To create a production version of your app:

```bash
npm run build
```

You can preview the production build with `npm run preview`.

> To deploy your app, you may need to install an [adapter](https://svelte.dev/docs/kit/adapters) for your target environment.
