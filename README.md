<p align="center">
	<img src="https://user-images.githubusercontent.com/30529572/72455010-fb38d400-37e7-11ea-9c1e-8cdeb5f5906e.png" />
	<h2 align="center"> Generic Events Registration Forms </h2>
	<h4 align="center"> Stores the registration details of participants in a provided DB and exposes APIs for viewing their count and other details. <h4>
</p>

---
[![DOCS](https://img.shields.io/badge/Documentation-see%20docs-green?style=flat-square&logo=appveyor)](https://documenter.getpostman.com/view/3896915/SWTD9HqJ?version=latest) 
  [![UI ](https://img.shields.io/badge/User%20Interface-Link%20to%20UI-orange?style=flat-square&logo=appveyor)](https://dsc-eventsreg.herokuapp.com/)


## Functionalities
- [X]  Simple registration form data capture
- [X]  Count event registrations by event name
- [X]  Count all live event registrations

<br>

## Description

Note that this project is capable of storing the following data points:

* Name
* Registration Number
* Phone Number
* Email Address
* DeviceID (For notifications)

Not all of them are required though. One additional field is required in the payload, which signifies which event is the entire payload for. It only gets captured if the event is whitelisted in the backend. For whitelisting your event, let the [maintainer](https://github.com/L04DB4L4NC3R) of this repository know.

<br>

## Instructions to run

* Pre-requisites:
	-  go v1.11+
	- The following `.env` file

```
HOST=localhost

# Not needed for heroku
PORT=3000

DATABASE_URI=mongodb://<username>:<password>@<host>:<port>/<database>
```

* How to build and execute
```bash
$ make build
$ make begin
```


<br>

## Contributors

* [Angad Sharma](https://github.com/L04DB4L4NC3R)



<br>
<br>

<p align="center">
	Made with :heart: by DSC VIT
</p>


