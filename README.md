## Setup Instructions

1. **Install dependencies**:
    ```bash
    npm install
    ```

2. **Build the project**:
    ```bash
    go build .
    ```

3. **Run the CLI tool**:
    ```bash
    ./tailwindcss-parser-cli "w-24"
    ```

## `--minify` Flag

The `--minify` flag is used to reduce the size of the output by removing unnecessary characters, such as whitespace and comments. This is especially useful for optimizing files for production, where smaller file sizes can lead to faster load times and better performance.

### Usage

To use the `--minify` flag with the `tailwindcss-parser-cli`, include it in your command like this:

```bash
./tailwindcss-parser-cli --minify "w-24"
```
