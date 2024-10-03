# [papercraft-ios](https://github.com/papercraft/papercraft-ios)

## Install

- Get *[papercraft-ios](https://github.com/papercraft/papercraft-ios)* source code
```
mkdir ~/Papercraft
cd ~/Papercraft
git clone --recursive https://github.com/papercraft/papercraft-ios.git
```

- build, see [build papercraft-ios](https://github.com/papercraft/papercraft-ios#compilation-guide), and google
```
cd ~/Papercraft/papercraft-ios
sh r.sh
```

## Patch

**Default connect to Papercraft Test Server.**

If you want to connect to your own server, you can modify the following code:

[Network.swift#L473](https://github.com/papercraft/papercraft-ios/blob/papercraft/submodules/papercraftCore/Sources/Network/Network.swift#L473)

```
https://github.com/papercraft/papercraft-ios/blob/papercraft/submodules/papercraftCore/Sources/Network/Network.swift#L473
if testingEnvironment {
    seedAddressList = [
        1: ["XXX.XXX.XXX.XXX"]
    ]
} else {
    seedAddressList = [
        1: ["XXX.XXX.XXX.XXX"]
    ]
}
```
