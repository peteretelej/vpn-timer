# vpn-timer

Alerts you when you're on VPN for too long. 


## Install
```
go get github.com/peteretelej/vpn-timer
```

## Use
```
vpn-timer -vpn 123.45.67.89 
# specify the VPN static IP so that it can be monitored
# by default uses a time limit of 30min

vpn-timer -vpn 123.45.67.89 -limit 1h
# send notifications after an hour
```