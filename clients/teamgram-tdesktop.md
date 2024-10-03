# [papercraft-tdesktop](https://github.com/papercraft/papercraft-tdesktop)

## Install

- Get *[papercraft-tdesktop](https://github.com/papercraft/papercraft-tdesktop)* source code

- build, see [Build papercraft-tdesktop](https://github.com/papercraft/papercraft-tdesktop#build-instructions)

## Patch

**Default connect to Papercraft Test Server.**

If you want to connect to your own server, you can modify the following code:

[mtproto_dc_options.cpp#L31](https://github.com/papercraft/papercraft-tdesktop/blob/papercraft2/papercraft/SourceFiles/mtproto/mtproto_dc_options.cpp#L31)

```
https://github.com/papercraft/papercraft-tdesktop/blob/papercraft2/papercraft/SourceFiles/mtproto/mtproto_dc_options.cpp#L31

const BuiltInDc kBuiltInDcs[] = {
    { 1, "XXX.XXX.XXX.XXX" , 10443 },
};

const BuiltInDc kBuiltInDcsTest[] = {
    { 1, "XXX.XXX.XXX.XXX" , 10443 },
};

```
