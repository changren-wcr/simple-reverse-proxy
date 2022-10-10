# Simple File Reverse Proxy

## Usage

1. Clone the repo to your local machine and cd into the directory
2. Start web server on your local machine by `go run .`
3. Open link <http://localhost:8080/> in your browser
4. Append `raw.githubusercontent.com/yaml/pyyaml/master/README.md` to <http://localhost:8080/> to get url <http://localhost:8080/raw.githubusercontent.com/yaml/pyyaml/master/README.md>
5. Open <http://localhost:8080/raw.githubusercontent.com/yaml/pyyaml/master/README.md> in your browser, which indirectly fetch original file <https://raw.githubusercontent.com/yaml/pyyaml/master/README.md>

## Deploy

### Docker

```bash
docker build -t proxy-of-file .
docker run -dp 8080:8080 proxy-of-file
```
open <http://localhost:8080> in your browser

### Heroku CLI

```bash
heroku login
heroku container:login
heroku create
heroku container:push web
heroku container:release web
heroku open
```
your default browser will automatically open service url