Bugfix: Config file value not being read

There was a bug in which phoenix config is always set to the default values and the contents of the config file were actually ignored.

https://github.com/owncloud/mono/phoenix/pull/45
https://github.com/owncloud/mono/phoenix/issues/46
https://github.com/owncloud/mono/phoenix/issues/47
