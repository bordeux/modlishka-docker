# ..Modlishka..
[It is unofficial docker image for Modlishka project](https://github.com/drk1wi/Modlishka)

More about project

Modlishka is a flexible and powerful reverse proxy, that will take your phishing campaigns to the next level (with minimal effort required from your side). 
 
Enjoy :-)


Features
--------

Some of the most important 'Modlishka' features :

-   Support for majority of 2FA authentication schemes (by design).
-   No website templates (just point Modlishka to the target domain - in most cases, it will be handled automatically).
-   Full control of "cross" origin TLS traffic flow from your victims browsers (through custom new techniques).
-   Flexible  and easily configurable phishing scenarios through configuration options.
-   Pattern based JavaScript payload injection.
-   Striping website from all encryption and security headers (back to 90's MITM style). 
-   User credential harvesting (with context based on URL parameter passed identifiers).
-   Can be extended with your ideas through plugins.
-   Stateless design. Can be scaled up easily for an arbitrary number of users - ex. through a DNS load balancer.
-   Web panel with a summary of collected credentials and user session impersonation (beta).
-   Written in Go.


Action
------
_"A picture is worth a thousand words":_

 Modlishka in action against an example 2FA (SMS) enabled authentication scheme:

[![Watch the video](https://i.vimeocdn.com/video/749353683.jpg)](https://vimeo.com/308709275)

[https://vimeo.com/308709275](https://vimeo.com/308709275)

Note: google.com was chosen here just as a POC.
    
![alt text](https://raw.githubusercontent.com/drk1wi/assets/master/7d0426a133a85a46a76a424574bf5a2acf99815e.png)


  
    
    Usage of docker file - Environment variables
          
      -ML_CERT string
        	base64 encoded TLS certificate
      
      -ML_CERT_KEY string
        	base64 encoded TLS certificate key
      
      -ML_CERT_POOL string
        	base64 encoded Certification Authority certificate
      
      -ML_CONFIG string
        	JSON configuration file. Convenient instead of using command line switches.
      
      -ML_CRED_PARAMS string
          	Credential regexp collector with matching groups. Example: base64(username_regex),base64(password_regex)

      -ML_DEBUG
        	Print debug information
      
      -ML_DISABLE_SECURITY
        	Disable security features like anti-SSRF. Disable at your own risk.
      
      -ML_JS_RULES string
        	Comma separated list of URL patterns and JS base64 encoded payloads that will be injected. 
      
      -ML_LISTENING_ADDRESS string
        	Listening address (default "127.0.0.1")
      
      -ML_LISTENING_PORT string
        	Listening port (default "443")
      
      -ML_LOG string
        	Local file to which fetched requests will be written (appended)
      
      -ML_PHISHING string
        	Phishing domain to create - Ex.: target.co
      
      -ML_PLUGINS string
        	Comma seperated list of enabled plugin names (default "all")
      
      -ML_POST_ONLY
        	Log only HTTP POST requests
      
      -ML_RULES string
        	Comma separated list of 'string' patterns and their replacements. 
      
      -ML_TARGET string
        	Main target to proxy - Ex.: https://target.com
      
      -ML_TARGET_RES string
        	Comma separated list of target subdomains that need to pass through the  proxy 
      
      -ML_TERMINATE_TRIGGERS string
        	Comma separated list of URLs from target's origin which will trigger session termination
      
      -ML_TERMINATE_URL string
        	URL to redirect the client after session termination triggers
      
      -ML_TLS
        	Enable TLS (default false)
      
      -ML_TRACKING_COOKIE string
        	Name of the HTTP cookie used to track the victim (default "id")
      
      -ML_TRACKING_PARAM string
        	Name of the HTTP parameter used to track the victim (default "id")




Usage
-----

 * Check out the [wiki](https://github.com/drk1wi/Modlishka/wiki) page for a more detailed overview of the tool usage.
 * [FAQ](https://github.com/drk1wi/Modlishka/wiki/FAQ) (Frequently Asked Questions)
 * [Blog post](https://blog.duszynski.eu/phishing-ng-bypassing-2fa-with-modlishka/)


License
-------
Modlishka was made by Piotr Duszyński ([@drk1wi](https://twitter.com/drk1wi)). You can find the license [here](https://github.com/drk1wi/Modlishka/blob/master/LICENSE).

Credits
-------
Thanks for helping with the code go to Giuseppe Trotta ([@Giutro](https://twitter.com/giutro)) 


Disclaimer
----------
This tool is made only for educational purposes and can be only used in legitimate penetration tests. Author does not take any responsibility for any actions taken by its users.

-------

[![Twitter](https://img.shields.io/badge/twitter-drk1wi-blue.svg)](https://twitter.com/drk1wi)

