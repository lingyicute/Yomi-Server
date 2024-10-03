# papercraft-android

## Install

- Get *[papercraft-android](https://github.com/papercraft/papercraft-android)* source code

- build, see [build papercraft-android](https://github.com/papercraft/papercraft-android/tree/master#compilation-guide), and google

## patch

**Default connect to Papercraft Test Server.**

If you want to connect to your own server, you can modify the following code:

[ConnectionsManager.cpp#L1691](https://github.com/papercraft/papercraft-android/blob/papercraft/TMessagesProj/jni/tgnet/ConnectionsManager.cpp#L1691)

```
https://github.com/papercraft/papercraft-android/blob/papercraft/TMessagesProj/jni/tgnet/ConnectionsManager.cpp#L1691
void ConnectionsManager::initDatacenters() {
    Datacenter *datacenter;
    if (!testBackend) {
        if (datacenters.find(1) == datacenters.end()) {
            datacenter = new Datacenter(instanceNum, 1);
            datacenter->addAddressAndPort("XXX.XXX.XXX.XXX", 10443, 0, "");
            datacenters[1] = datacenter;
        }
    } else {
        if (datacenters.find(1) == datacenters.end()) {
            datacenter = new Datacenter(instanceNum, 1);
            datacenter->addAddressAndPort("XXX.XXX.XXX.XXX", 10443, 0, "");
            datacenters[1] = datacenter;
        }
    }
}


```
