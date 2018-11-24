# vpn-timer

Alerts you when you're on VPN for too long. 


## Install
Grab an binary executable for your OS from [Releases](https://github.com/peteretelej/vpn-timer/releases)

Or, if you have Go:
```
go get github.com/peteretelej/vpn-timer
```



## Use
```
vpn-timer -ip 123.45.67.89 
# specify the VPN public IP so that it can be monitored
# by default uses a time limit of 30min

vpn-timer -ip 123.45.67.89 -limit 1h
# send notifications after an hour
```
