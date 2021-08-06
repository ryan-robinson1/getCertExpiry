# getCertExpiry
## Overview
getCertExpiry is a go command line tool that can find the expiration dates of server certs and check if they are expired
## Getting Started

### Installation
* ``git clone https://github.com/ryan-robinson1/getCertExpiry.git ``
### Setup
* cd into the getCertExpiry directory and build the binary with ``go build``
## Usage


 <font size="3">Gets the expiration date of the specified server cert. Use the ``-c`` and ``-k`` flags to load in a client cert and key file. Use the ``-a`` flag to load in a cert authority file. Use the ``-i`` flag to allow insecure TLS connections. **NOTE:** you can  also specify the url without a flag by inputting it as the last command line argument </font> <pre>$ ./getCertExpiry -u <span style="color:magenta"><i><b>ADDRESS</b></i></span>:<span style="color:magenta"><i><b>PORT</b></i></span> [-c <span style="color:magenta"><i><b>CERT_FILE</b></i></span> -k <span style="color:magenta"><i><b>KEY_FILE</b></i></span>][-a <span style="color:magenta"><i><b>CA_FILE</b></i></span>] [-i]</pre>
 
## Exit Codes
* _Exit Code 0 : Cert is valid_
* _Exit Code 1 : Cert is expired_
* _Exit Code 3 : Certs are not supported_
* _Exit Code 4 : Cert is untrusted_
* _Exit Code 5 : Invalid inputted certs_
* _Exit Code 6 : No args_

## Examples
    $ ./getCertExpiry -u example.com:443
    $ ./getCertExpiry -u example.com:443 -i
    $ ./getCertExpiry -u example.com:443 -i -c client.crt -k client.key -a rootca.crt 
    $
    $ ./getCertExpiry example.com:443
    $ ./getCertExpiry -i example.com:443
    $ ./getCertExpiry -i -c client.crt -k client.key -a rootca.crt example.com:443
  
    


 

 


---
