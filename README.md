<p align=center><a href="#"><img src=https://gitlab.com/rdodin/pics/-/wikis/uploads/5f77e2237d36f608d7ba810e79a324af/configtopus-L.webp?sanitize=true/></a></p>

**Configtopus** is an NDK app with a configuration tree that consists of various combinations of YANG data types. It is used to show how configuration changes made to the application tree are streamed back to the application and how to update application state for different configuration elements.

## Quickstart

Clone and enter the repository:

```bash
git clone https://github.com/srl-labs/ndk-configtopus.git && \
cd ndk-configtopus
```

One-click deployment of the lab is available via `run.sh` script. It will build the app, deploy the lab and onboard the application to SR Linux:

```bash
./run.sh deploy-all
```

Enter the SR Linux CLI:

```
ssh configtopus
```

Once entered into the SR Linux CLI, you can find `/configtopus` context available that contains the application's configuration and operational data.

Experiment with the application config and observe how the application reacts to the changes. You can see the application logs in the `./log/srl/stdout/configtopus.log` file available from the host machine.

## Configuration tree

Configuration tree can be generated by running:

```bash
./run.sh conf-tree
```

The following output should be generated:

```
module: configtopus
  +--rw configtopus!
     +--rw action-leaf-node?           enumeration
     +--rw leaf-list-node*             string
     +--rw list-node* [name]
     |  +--rw name               string
     |  +--rw child-leaf-list*   string
     |  +--ro state?             uint64
     +--rw list-with-container* [value]
     |  +--rw value             string
     |  +--ro state?            uint64
     |  +--rw container-leaf
     |     +--rw leaf-uint?   uint64
     +--rw parent-list-node* [name]
     |  +--rw name          string
     |  +--rw child-list* [name]
     |     +--rw name     string
     |     +--ro state?   uint64
     +--rw container-with-leaf!
     |  +--rw leaf-decimal?                decimal64
     |  +--rw leaf-uint?                   uint64
     |  +--rw child-container-with-leaf
     |     +--rw child-container-with-leaf-list
     |        +--rw child-leaf-list*   string
     +--rw container-with-leaf-list
     |  +--rw child-leaf-list*   string
     +--rw container-with-list
        +--rw leaf-uint?    uint64
        +--rw child-list* [name]
           +--rw name     string
           +--ro state?   uint64
```

## Shell autocompletions

To get bash autocompletions for `./run.sh` functions:

```bash
source ./run.sh
```
