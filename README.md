urlutils
========

[![Build Status](https://travis-ci.org/ernestas-poskus/urlutils.svg?branch=master)](https://travis-ci.org/ernestas-poskus/urlutils)

Golang standard URL wrapper, adds syntactic sugar and few new methods.


- ResolveURL: resolves relative URL to absolute URL
- IsAsset: matches asset URL's like: .css, .js, etc.
- IsRelative: checks whether URL is relative, e.g.: /news/article/13.html
- IsAbsolute: checks whether URL has absolute (full) path
- SameDomain: compares URL's checks if they have same domain
