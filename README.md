# xk6 Hello World Extension

This project is a simple example of creating a custom extension for [k6](https://k6.io/) using [xk6](https://github.com/grafana/xk6). The extension provides a basic "Hello World" functionality that can be used in k6 test scripts.

## Project Structure

```
.
├── go.mod       # Go module file
├── go.sum       # Go module dependencies
├── hello.go     # Custom k6 extension implementation
├── k6           # Built k6 binary (excluded from version control)
└── test.js      # Example k6 script using the custom extension
```

## Prerequisites

- **Go 1.17 or later**: The extension is built using Go, so you need to have Go installed on your machine.
- **xk6**: You need xk6 to build a custom k6 binary with this extension.

## Building the Custom k6 Binary

To build the custom k6 binary with the `hello` extension, follow these steps:

1. Clone the repository and navigate to the project directory.
2. Run the following command to build the k6 binary with the custom extension:

   ```bash
   xk6 build --with github.com/your-username/xk6-hello-world=.
   ```

   This will create a new `k6` binary in the project directory.

## Using the Extension

You can use the custom `hello` extension in your k6 scripts. Here's an example script (`test.js`):

```javascript
import hello from 'k6/x/hello';

export default function () {
  console.log(`${hello.helloWorld()}`);
}
```

Run the script using the custom k6 binary:

```bash
./k6 run test.js
```

You should see the following output in your console:

```
Hello World!
```
