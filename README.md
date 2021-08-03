# getCertExpiry
## Overview
getCertExpiry is a go command line tool to find the expiration dates of given server certs and check if they are expired
## Getting Started

### Installation
* ``git clone https://github.com/ryan-robinson1/getCertExpiry.git ``
### Setup
* cd into the getCertExpiry directory and build the binary with ``go build``
## Usage


 <font size="3">Gets the expiration date of the specified server cert. Use the ``--insecure`` flag to allow insecure TLS connections</font> <pre>$ ./getCertExpiry <span style="color:magenta"><i><b>ADDRESS</b></i></span>:<span style="color:magenta"><i><b>PORT</b></i></span> [--insecure] </pre>
## Exit Codes
* _Exit Code 0 : Cert is valid_
* _Exit Code 1 : Cert is expired_
* _Exit Code 3 : Certs are not supported_
* _Exit Code 4 : Cert is untrusted_
* _Exit Code 5 : Invalid Args_

## Example
    $ ./getCertExpiry example.com:5000
    $ ./getCertExpiry example.com:5000 --insecure


 

 


---
