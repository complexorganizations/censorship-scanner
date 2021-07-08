### Censorship-Scanner

Censorship-Scanner will scan your local network and determine if there is any network-based censorship present.

---
### Features
- Check for censorship on the local network.
- Check TOR, HTTPS, DNS

---
### Installation
Let's begin by obtaining the most recent version of the Censorship-Scanner binary.
```
go get -v -u github.com/complexorganizations/censorship-scanner
```
Let's start the binary.
```
censorship-scanner -advanced
```
You can usually disregard `Error`, but if something `Failed`, you most likely have a problem.

---
### Q&A

What is the purpose of censorship-scanner?
- censorship-scanner is a high-level network scanner that will scan your local network and detect if network-based censorship is present.

Is it possible to get around browser-based censorship using this?
- No, This is just a scanner.

Are there any logs on this?
- No

Is it possible to use this to see if your DNS provider is limiting access to a particular platform?
- Yes

What is the total number of scans?
- 250 Domain(s)
- 250 Tor IP(s)

---
### Author
* Name: Prajwal Koirala
* Website: [prajwalkoirala.com](https://www.prajwalkoirala.com)

---	
### Credits
Open Source Community

---
### License
Copyright © [Prajwal](https://github.com/prajwal-koirala)

This project is unlicensed.
