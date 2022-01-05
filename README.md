# What is this?
I'm following the Twitch Twirp demo found at:

https://twitchtv.github.io/twirp/docs/install.html
https://twitchtv.github.io/twirp/docs/example.html

# Why?
To build the Proof of Concept.
To figure out how to install Twirp.

# Gotchas
- Need more help on installing tools correctly. It's pretty confusing.
- `protoc` options may have changed, or it doesn't like relative imports. I had to manually move the produced files to the example directory.
- Example code uses github.com as package names, so importing doesn't work as expected.
- Remember client and server are separate apps. See `twirp-haberdasher-client` for the client code.