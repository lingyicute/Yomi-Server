# Papercraft - Unofficial open source [mtproto](https://papercraft-official.github.io/mtproto) server written in golang
> open source mtproto server implemented in golang with compatible papercraft client.

## Introduce
Open source [mtproto](https://papercraft-official.github.io/mtproto) server implementation written in golang, support private deployment.

## Features
- MTProto 2.0
  - Abridged
  - Intermediate
  - Padded intermediate
  - Full
- API Layer: 186
- private chat
- basic group
- contacts
- web

## Architecture
![Architecture](docs/image/architecture-001.png)

## Installing Papercraft 
`Papercraft` relies on high-performance components: 

- **mysql5.7**
- [redis](https://redis.io/)
- [etcd](https://etcd.io/)
- [kafka](https://kafka.apache.org/quickstart)
- [minio](https://docs.min.io/docs/minio-quickstart-guide.html#GNU/Linux)
- [ffmpeg](https://www.johnvansickle.com/ffmpeg/)

Privatization deployment Before `Papercraft`, please make sure that the above five components have been installed. If your server does not have the above components, you must first install Missing components. 

- [Centos9 Stream Build and Install](docs/install-centos-9.md) [@A Feel]
- [CentOS7 papercraft-server环境搭建](docs/install-centos-7.md) [@saeipi]

If you have the above components, it is recommended to use them directly. If not, it is recommended to use `docker-compose-env.yaml`.


### Source code deployment
#### Install [Go environment](https://go.dev/doc/install). Make sure Go version is at least 1.17.


#### Get source code　

```
git clone https://github.com/lingyicute/papercraft-server.git
cd papercraft-server
```

#### Init data
- init database

	```
	1. create database papercraft
	2. init papercraft database
		mysql -uroot papercraft < papercraftd/sql/1_papercraft.sql
		mysql -uroot papercraft < papercraftd/sql/migrate-20220321.sql
		mysql -uroot papercraft < papercraftd/sql/migrate-20220326.sql
		mysql -uroot papercraft < papercraftd/sql/migrate-20220328.sql
		mysql -uroot papercraft < papercraftd/sql/migrate-20220401.sql
		mysql -uroot papercraft < papercraftd/sql/migrate-20220412.sql
		mysql -uroot papercraft < papercraftd/sql/migrate-20220419.sql
		mysql -uroot papercraft < papercraftd/sql/migrate-20220423.sql
		mysql -uroot papercraft < papercraftd/sql/migrate-20220504.sql
		mysql -uroot papercraft < papercraftd/sql/migrate-20220721.sql
		mysql -uroot papercraft < papercraftd/sql/migrate-20220826.sql
		mysql -uroot papercraft < papercraftd/sql/migrate-20220919.sql
		mysql -uroot papercraft < papercraftd/sql/migrate-20221008.sql
		mysql -uroot papercraft < papercraftd/sql/migrate-20221011.sql
		mysql -uroot papercraft < papercraftd/sql/migrate-20221016.sql
		mysql -uroot papercraft < papercraftd/sql/migrate-20221023.sql
		mysql -uroot papercraft < papercraftd/sql/migrate-20221101.sql
		mysql -uroot papercraft < papercraftd/sql/migrate-20221127.sql
		mysql -uroot papercraft < papercraftd/sql/migrate-20230707.sql
		mysql -uroot papercraft < papercraftd/sql/migrate-20240107.sql
		mysql -uroot papercraft < papercraftd/sql/migrate-20240108.sql
		mysql -uroot papercraft < papercraftd/sql/migrate-20240111.sql
		mysql -uroot papercraft < papercraftd/sql/migrate-20240112.sql
		mysql -uroot papercraft < papercraftd/sql/migrate-20240113.sql
		mysql -uroot papercraft < papercraftd/sql/migrate-20240114.sql
		mysql -uroot papercraft < papercraftd/sql/migrate-20240420.sql
		mysql -uroot papercraft < papercraftd/sql/migrate-20240620.sql
		mysql -uroot papercraft < papercraftd/sql/migrate-20240828.sql
		mysql -uroot papercraft < papercraftd/sql/z_init.sql
	```

- init minio buckets
	- bucket names
	  - `documents`
	  - `encryptedfiles`
	  - `photos`
	  - `videos`
	- Access `http://ip:xxxxx` and create


#### Build
	
```
make
```

#### Run

```
cd papercraftd/bin
./runall2.sh
```

### Docker deployment
#### Install [Docker](https://docs.docker.com/get-docker/)

#### Install [Docker Compose](https://docs.docker.com/compose/install/)

#### Get source code

```
git clone https://github.com/lingyicute/papercraft-server.git
cd papercraft-server
```

#### Run

```  
# run dependency
docker-compose -f ./docker-compose-env.yaml up -d

# run docker-compose
docker-compose up -d
```
	
## Compatible clients
[Android client for Papercraft](clients/papercraft-android.md)

[iOS client for Papercraft](clients/papercraft-ios.md)

[tdesktop for Papercraft](clients/papercraft-tdesktop.md)

## Feedback
Please report bugs, concerns, suggestions by issues, or join papercraft group **[Papercraft](https://papercraft-link.github.io/+TjD5LZJ5XLRlCYLF)** to discuss problems around source code.

## Notes
If need enterprise edition:

- sticker/theme/wallpaper/reactions/2fa/sms/push(apns/web/fcm)/secretchat/scheduled/...
- channel/megagroup
- audiocall/videocall/groupcall
- bots

please PM the **[author](https://papercraft-link.github.io/benqi)**

## Give a Star! ⭐

If you like or are using this project to learn or start your solution, please give it a star. Thanks!

## Visitors Count

<img align="left" src = "https://profile-counter.glitch.me/papercraft-server/count.svg" alt="Loading" />