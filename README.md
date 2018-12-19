# vpn-timer
[![FOSSA Status](https://app.fossa.io/api/projects/git%2Bgithub.com%2Fpeteretelej%2Fvpn-timer.svg?type=shield)](https://app.fossa.io/projects/git%2Bgithub.com%2Fpeteretelej%2Fvpn-timer?ref=badge_shield)


Alerts you when you're on VPN for too long. 


## Install
Grab an binary executable for your OS from [Releases](https://github.com/peteretelej/vpn-timer/releases)

Or, if you have Go:
```
git clone https://github.com/peteretelej/vpn-timer.git

cd vpn-timer

go install
```



## Usage
Run `vpn-timer` while specifying the VPN IP (see [ifconfig.co](https://ifconfig.co), and **_optionally_** you can specify the time _limit_ you want eg `10m`,`1h`,`2h30m` (default `30m`)
```
vpn-timer -ip 123.45.67.89 
# specify the VPN public IP so that it can be monitored
# by default uses a time limit of 30min

vpn-timer -ip 123.45.67.89 -limit 1h
# send notifications after an hour
```


## License
[![FOSSA Status](https://app.fossa.io/api/projects/git%2Bgithub.com%2Fpeteretelej%2Fvpn-timer.svg?type=large)](https://app.fossa.io/projects/git%2Bgithub.com%2Fpeteretelej%2Fvpn-timer?ref=badge_large)