go.gravatar
===========

Library for handing gravatar requests server-side.


General information
===========

When usin gravatar, requests to gravatar.com are 
handled client-side, thus giving gravatar insight
into who visited what and when. From a privacy standpoint
that is unacceptable. This library fetches and stores
avatars from gravatar 100% server side.

To protect your user's privacy it is essential that
you use this library correctly. If you fetch the
gravatar every time a request is made, it's possible
to deduce the identity of your visitors through
traffic analysis.


