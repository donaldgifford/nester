# nester

Go tool for pulling Nest Data to InfluxDB

**NOTE** This will only work if ran locally on a machine. The OAuth2 flow requires the local machine browser to open as its configured currently for localhost only.

## what is this?

nester is a tool for polling a Nest Thermostat and pushing data to InfluxDB. It uses OAuth2 to manage tokens for requests made to the Nest API.

## Setup and usage

either copy the `.nester.yaml.example` config file to where you will run the binary from or you can run the `nester init config` command to generate the config file.

After that you can enter in all the relevant information into the config file and try running it using `nester run`. Add the `-p` flag if you want to output the metrics received to stdout.

## Local setup and development

pulling the repo and running the nester commands from go run `go run main.go commands`. Also running `docker-compose up -d` to run a local influxdb environment.

For that be sure to create a `.env` file with the following content:

```bash
#influxdb
INFLUXDB_HTTP_AUTH_ENABLED=true
INFLUXDB_ADMIN_USER=admin
INFLUXDB_ADMIN_PASSWORD=admin
INFLUXDB_DB=mydb
INFLUXDB_ORG=myorg
INFLUXDB_BUCKET=mybucket
INFLUXDB_USERNAME=user
INFLUXDB_PASSWORD=password
INFLUXDB_ADMIN_TOKEN=useradminsecret
```

Change the values to whatever you want to use. Then you can access it locally at `http://localhost:8086`

### Example Caddy File

```caddyfile
nester.poop.software {
        reverse_proxy nester:8080

        tls {
                dns namecheap {
                        api_key APIKEY
                        user USER
                        api_endpoint https://api.namecheap.com/xml.response
                        client_ip MYEXTERNALIP
                }
        }
}
```

## Todos

- [ ] Setup GitHub Actions for builds and packages/releases
- [ ] Setup CI/tests? coverage etc
- [ ] Docs for deploying into a cronjob
- [ ] Maybe add http endpoints to run as a service with a `/metrics` and `/status` for Prom style metrics. Also this might make it easier to deal with the OAuth2 flow stuff instead of spitting up a endpoint just for the ApprovalFlow.
- [ ] add configuration so that instead of spawning initial browser code to localhost, add an IP in so that it can be configured remotely.
