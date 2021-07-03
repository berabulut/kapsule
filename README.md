
<h1>kapsule</h1>

<div align="center">
  <img src="https://github.com/berabulut/kapsule/blob/main/web/static/logo.png?raw=true" alt="kapsule" />
  <h3>
    <a href="https://kapsule.click/">
      Live
    </a>
    <span> | </span>
    <a href="https://github.com/berabulut/kapsule-ui">
      kapsule-ui
     </a>
    <span> | </span>
    <a href="https://github.com/berabulut/kapsule-server">
      kapsule-server
    </a>
  </h3><br>
  A URL shortener that collects simple user data when someone clicks a shortened link.
</div>


## What does it collect ?

- Operating system name
- Browser name
- Browser language
- Country name
- Device type (Mobile / Desktop)

## How does it collect ?

All of the data is collected through HTTP headers.

- It parses User-Agent for operating system name, browser name, device type. 
- It parses Accept-Language for browser language.
- It uses [geojs](https://www.geojs.io/) API for country name. It calls geojs with IP address it takes from X-FORWARDED-FOR header. 

## Apps

- ### api

	It shortens URLs and return URL records to client.

	- **/shorten** (POST)

		```
		{
			"url" : "https://github.com/berabulut",
			"options_enabled": true, // activate waiting page
			"duration": 6, // duration of waiting page
			"note": "Letter to a stranger." // small note to be shown in waiting page
		}
		```

	- **/:key** (GET)

		Returns URL data. Key is unique id of shortened link. Usage:

		`localhost:8080/ab123`

	- **/details** (GET) 

		Does same thing with `/:key` but returns multiple records (up to 3). Usage :

		`localhost:8080/details?keys[0]=ab123&keys[1]=ab321&keys[2]=ab231`

- ### redirect

	Redirects client to target URL and collects user data.

	- **/:key** (GET)

		Redirects to target URL. Key is unique id of shortened link. Usage:

		`localhost:8080/ab123`

## Build and Run

`sh build.sh`

Exposes **api** to localhost:8080 and **redirect** to localhost:8081. 