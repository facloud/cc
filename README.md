# Container cloud

Container cloud (CC) is a work-in-progress. I am not sure what it is and it
will probably die after a few commits. However, hey, it's fun playing with
stuff. PS: The name 'container cloud' was chosen in case this thing becomes
a thing and it needs a name that will distinguish it from other things.

## Architecture

TODO.

## Technologies you would like to try

* DHT or Gossip or Tracker
* Flanner: For container-to-container networking
* gRPC: For communication between components
* LVM: For container volumes
* WebSockets: For transmitting the output through HTTP
* Nftables: For container firewall
* UnP: For container firewall and port forwarding

### Concepts

List of concepts mentioned bellow:

* Container: a containerized process.
* Hosts: a machine that runs containers.
* Availability zone: a physical location that contains multiple hosts.
* VPC: a local subnet in an availability zone that can be used by multiple
  containers guaranteeing high network performance.
* Volume: a data volume that is attached and used by multiple containers in
  the same host. A volume can be persistent (will survive losing the host)
  or ephemeral.

### Simplifying the prototype

To simplify the prototype following shortcuts are taken:

1. All the volumes are formatted in ext4.

## The `cc` CLI

All `cc` commands can receive the authentication details as:

* **Endpoint URL**: `--cc-endpoint-url` or `CC_ENDPOINT_URL` environment variable
* **Access token**: `--cc-access-token` or `CC_ACCESS_TOKEN` environment variable
* **Secret key**: `--cc-secret-key` or `CC_SECRET_KEY` environment variable

### Create a VPC

Create a virtual private network within a host or a local cluster:

```
cc vpc create
  --cidr <CIDR>
  --name <Name>
```


### List VPCs

```
cc vpc list
```

### Delete VPC(s)

The command will fail if one ore more VPCs are in use. The command will succeed
if non-existing or already deleted VPCs are provided.

```
cc vpc delete <ID> [...]
```

### Create a volume

```
cc volume create
  --size <Size in bytes>
  --persistent <Set flag it the volume needs to be persistent>
  --name <Name>
```

### List volumes

```
cc volume list
```

### Delete volume(s)

The command will fail if one ore more volumes are in use. The command will
succeed if non-existing or already deleted VPCs are provided.

```
cc volume delete <ID> [<ID> ...]
```

### Create container

```
cc run
  --image <Image>
  --vpc <VPC ID>
  --volume <Volume ID>:<Path> [..]
  --memory <Memory in bytes>
  --cpu <CPU in shares>
  --disk <Disk size in bytes>
  --user <User to run command as>
  --env <Environment variables, format: name=value> [..]
  --dir <Directory to run the command in>
  <Path to command> [<Command argument> ...]
```

### Get container status

```
cc status <Container ID> [...]
```

### Get container metrics

```
cc metrics <Container ID> [...]
```

### Create local-remote server

```
cc local-remote serve
  --container-id <Container ID>
  --container-port <Container port>
  --bind <Local bind address>
```
