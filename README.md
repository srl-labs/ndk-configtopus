<p align=center><a href="#"><img src=https://gitlab.com/rdodin/pics/-/wikis/uploads/5f77e2237d36f608d7ba810e79a324af/configtopus-L.webp?sanitize=true/></a></p>

Configtopus is an NDK app with a configuration tree that consists of all possible combination of YANG data types. It is used to show how configuration changes made to the application tree are streamed back to the application and how to update application state for different configuration elements.

## Quickstart

Clone and enter the repository:

```bash
git clone https://github.com/srl-labs/ndk-configtopus.git && \
cd ndk-configtopus
```

Build the application and deploy it to the lab:

```
./run.sh deploy-all
```

Once the lab is deployed, the application is automatically onboarded to SR Linux.

Enter the SR Linux CLI:

```
ssh configtopus
```

Once entered into the SR Linux CLI, you can finde `/configtopus` context available that contains the application's configuration and operational data.

Configure the desired name:

```
--{ + running }--[  ]--
A:greeter# enter candidate

--{ + candidate shared default }--[  ]--
A:greeter# greeter name srlinux-user
```

Commit the configuration:

```
--{ +* candidate shared default }--[  ]--
A:greeter# commit now
All changes have been committed. Leaving candidate mode.
```

The application will now greet you when you list its operational state:

```
--{ + running }--[  ]--
A:greeter# info from state greeter
    greeter {
        name srlinux-user
        greeting "👋 Hello srlinux-user, I was last booted at 2023-11-26T10:24:27.374Z"
    }
```

## Shell autocompletions

To get bash autocompletions for `./run.sh` functions:

```bash
source ./run.sh
```
