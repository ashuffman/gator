# Gator: A Silly RSS Aggregator CLI Project

This is my implementation of the [bootdev](https://www.boot.dev/) project [Build a Blog Aggregator in Go](https://www.boot.dev/courses/build-blog-aggregator-golang). It's surely not the most convenient or fullest featured RSS aggregator out there. But it *is* a fun little toy to play with. 

All code in this project was written by me, by hand. The idea and project direction came from bootdev.

## Table of Contents
- Installation
	- 1. Install Go
	- 2. Install Postgres
	- 3. Install Gator
- Config
	- 1. Config File
	- 2. Database Schema
- Running the Program
	- 1. Overview
	- 2. Commands
  

## Installation:

### 1. Install Go
To use the Gator CLI, you need an up-to-date Golang toolchain installed on your system.

**Option 1 (Linux/WSL/macOS):** The Webi installer is the simplest way for most people. Just run this in your terminal:

```shell
curl -sS https://webi.sh/golang | sh
```

*Read the output of the command and follow any instructions.*

**Option 2 (any platform, including Windows/PowerShell):** Use the [official Golang installation instructions](https://go.dev/doc/install). On Windows, this means downloading and running a `.msi` installer package; the rest should be taken care of automatically.

After installing Golang, _open a new shell session_ and run `go version` to make sure everything works. If it does, _move on to step 2_.

**Optional troubleshooting:**

- If you already had Go installed with Webi, you should be able to run the same Webi command to update it.
    
- If you already had a version of Go installed a different way, on Linux/macOS you can run `which go` to find out where it's installed, and (if needed) remove the old version manually. (In PowerShell on Windows, the equivalent is `Get-Command go`.)
    
- If you're getting a "command not found" error after installation, it's most likely because the directory containing the `go` program isn't in your [`PATH`](https://opensource.com/article/17/6/set-path-linux). You need to add the directory to your `PATH` by modifying your shell's configuration file. First, you need to know _where_ the `go` command was installed. It might be in:
    
    - `~/.local/opt/go/bin` (Webi)
    - `/usr/local/go/bin` (official installation)
    - Somewhere else?

You can ensure that the program exists by attempting to run `go` using its full filepath. For example, if you think it's in `~/.local/opt/go/bin`, you can run `~/.local/opt/go/bin/go version`. If that works, then you just need to add `~/.local/opt/go/bin` to your `PATH` and reload your shell:
```shell
# For Linux/WSL
echo 'export PATH=$PATH:$HOME/.local/opt/go/bin' >> ~/.bashrc
# Next, reload your shell configuration
source ~/.bashrc
```

```shell
# For macOS
echo 'export PATH=$PATH:$HOME/.local/opt/go/bin' >> ~/.zshrc
# Next, reload your shell configuration
source ~/.zshrc
```

```shell
# For fish
fish_add_path $HOME/.local/opt/go/bin
```
### 2. Install Postgres
Gator uses a locally hosted Postgres database to store data. Since you'll be hosting it yourself, you'll need to install Postgres v15 or newer.
#### macOS with [brew](https://brew.sh/):
```shell
brew install postgresql@15
```

Ensure the installation worked. The `psql` command line utility is the default client for Postgres. Use it to make sure you're on version 15+ of Postgres
```shell
psql --version
```

Start the Postgres server in the background:
```shell
brew services start postgresql@15
```

Connect to the server. The `psql` client is the "default" for Postgres.
- Enter the `psql` shell
```shell
psql postgres
```
- Your new prompt should look like:
```shell
postgres=#
```

Create a new database called `gator`:
```shell
CREATE DATABASE gator;
```

Connect to the new database:
```shell
\c gator
```
You should see a prompt that looks like:
```shell
gator=#
```

Type `exit` to leave the `psql` shell.

#### Linux / WSL (Debian)
Here are the [docs from Microsoft](https://learn.microsoft.com/en-us/windows/wsl/tutorials/wsl-database#install-postgresql), but simply:
```shell
sudo apt update
sudo apt install postgresql postgresql-contrib
```

Update Postgres password:
```shell
sudo passwd postgres
```
Enter a password and be sure that you won't forget it. You can just use something easy like `postgres`

Start the Postgres server in the background:
```shell
sudo service postgresql start
```

Connect to the server. The `psql` client is the "default" for Postgres.
- Enter the `psql` shell
```shell
sudo -u postgres psql
```
- Your new prompt should look like:
```shell
postgres=#
```

Create a new database called `gator`:
```shell
CREATE DATABASE gator;
```

Connect to the new database:
```shell
\c gator
```
You should see a prompt that looks like:
```shell
gator=#
```

Set the user password:
```shell
ALTER USER postgres PASSWORD `postgres`;
```
> [! info] `postgres` is used here for simplicity. Before we altered the user's password, now we're altering the *database* user's password.

Type `exit` to leave the `psql` shell.

### 3. Install The Gator CLI
The following command will download, build, and install the `gator` CLI tool into your Go toolchain's `bin` directory:
```shell
go install github.com/ashuffman/gator@latest
```

Run `gator --version` on your command line to make sure the installation worked. If it did, move on the the Config section.

If you're getting a "command not found" error for `gator`, it's likely because the directory containing the `gator` program isn't in your [`PATH`](https://opensource.com/article/17/6/set-path-linux). You need to add the directory to `PATH` by modifying your shell's configuration file. In most cases, this means adding `$HOME/go/bin` (the default `GOBIN` directory where `go` installs programs):

```shell
# For Linux/WSL
echo 'export PATH=$PATH:$HOME/go/bin' >> ~/.bashrc
# Next, reload your shell configuration
source ~/.bashrc
```

```shell
# For macOS
echo 'export PATH=$PATH:$HOME/go/bin' >> ~/.zshrc
# Next, reload your shell configuration
source ~/.zshrc
```

```shell
# For fish
fish_add_path $HOME/go/bin
```

## Config
Gator is a multi-user CLI application. There's no _server_ (other than the database), so it's only intended for local use but, that doesn't mean we can't have multi-user functionality.

Gator uses a single JSON file to keep track of two things:
1. Who is currently logged in
2. The connection credentials for the PostgreSQL database

### 1. Create the Config File
Manually create a config file in your home directory `~/.gatorconfig.json`. 

Add the following content to the file:
```json
{
  "db_url": "postgres://<username>:<password>@localhost:5432/<dbname>?sslmode=disable"
}
```

### 2. Build the Database Schema
Use goose to run the the database migrations. If you don't already have goose installed you can run the following command to install it:
```shell
go install github.com/pressly/goose/v3/cmd/goose@latest
```

Once goose is installed `cd` into the `gator/sql/schema` directory and run an up migration:
```shell
goose postgres '<connection_string>' up

# example:
# goose postgres 'postgres://ashuffman:@localhost:5432/gator' up
```

## Running the Program
### 1. Overview
Once some basic set up has been done, Gator's `agg` function is intended to run in its own terminal window while you interact with the rest of the program from another terminal window. Before you run `gator agg` for the first time, we need to give it something to work with. 
- [ ] `register`:
```shell
gator register <username>
```
- `username` is a string with no spaces
- `register` automatically logs you in. To switch users once you've already registered, see the `login` command below.
- [ ] `addfeed`
```shell
gator addfeed <feed_name> <feed_url>
```
- `feed_name` and `url` are both string values
- `feed_name` can have spaces
- Make sure to wrap the `feed_name` and `feed_url` in quotes so that `gator` correctly interprets special characters and treats the arguments as intended.
- `addfeed` automatically creates a follow record for the currently logged in user
- add as many feeds as you like, then start aggregating!
- [ ] `agg`
```shell
gator agg <duration>

# examples: 
# gator agg 30s
# gator agg 10m
# gator agg 1h
```
- `duration` is the interval after which gator will scrape the feeds for the user 
- [ ] `browse`
```shell
gator browse <limit>

# example:
# gator browse 5
```
- `limit` is an integer value that controls how many articles are displayed in the terminal
	- the default `limit` is 2
- `browse` displays the most recent articles from your feeds
### 2. Other Commands
- `login` changes the current active user
```shell
gator login <username>
```
- `users` lists registered users
```shell
gator users
```
- `feeds` lists active feeds
```shell
gator feeds
```
- `follow` adds a specified feed to the current user's `following` list
```shell
gator follow <feed_url>
```
- `following` lists feeds the current user is following
```shell
gator following
```
- `unfollow` removes a specified feed from the current user's `following` list
```shell
gator unfollow <feed_url>
```
