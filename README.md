# nicehash-exporter
Nicehash API exporter with prometheus


## Docker usage

```yml
nicehashexporter:
    image: thomaslacaze/nicehash-exporter
    container_name: nicehash_exporter
    restart: unless-stopped
    environment:
    - APIUrl=https://api2.nicehash.com
    - APIKey=****
    - APISecret=****
    - XOrganizationId=****
    - ENVIRONMENT=prod
    ports: 
    - 9159:9159
```

## Exporter metrics 

```
# HELP nicehash_miningSpeed Mining speed in MH
# TYPE nicehash_miningSpeed gauge
nicehash_miningSpeed{algo="****",device="****",localisation="****"} 0
nicehash_miningSpeed{algo="****",device="****",localisation="****"} 0
# HELP nicehash_nextpayouttimestamp next payout timestamp
# TYPE nicehash_nextpayouttimestamp gauge
nicehash_nextpayouttimestamp 0
# HELP nicehash_temperatureDevice Temperature in °C
# TYPE nicehash_temperatureDevice gauge
nicehash_temperatureDevice{device="****",localisation="****"} 0
nicehash_temperatureDevice{device="****",localisation="****"} 0
# HELP nicehash_temperatureVRAM Temperature in °C
# TYPE nicehash_temperatureVRAM gauge
nicehash_temperatureVRAM{device="****",localisation="****"} 0
# HELP nicehash_unpaidAmount unpaid in BTC
# TYPE nicehash_unpaidAmount gauge
nicehash_unpaidAmount 0
# HELP nicehash_walletbalance Balance in BTC
# TYPE nicehash_walletbalance gauge
nicehash_walletbalance 0
```

## API call to interaction with your rig or device 

If you want to specify a device use 'deviceId' instead of 'rigId'

Start mining :
```bash
$ curl --location --request POST 'http://****:9159/api/mining?rigId=****&action=START'
```

Stop mining :
```bash
$ curl --location --request POST 'http://****:9159/api/mining?rigId=****&action=STOP'
```

Status mining (Mining=1, Not mininig = 0)
```bash
$ curl --location --request POST 'http://****:9159/api/status?rigId=****'
```

Change GPU optimisation :
```bash
curl -X POST -G 'http://****:9159/api/mining' -d rigId=**** -d action=NHQM_SET -d 'options=V=1%3BOP=2%3B'
```

`V=1%3BOP=2%3B` equal to `V=1;OP=2;`

[See more for actions options](https://github.com/nicehash/rest-clients-demo/issues/38)
