# sqlTool

This repository has the code for **sqlTool** which works as a [postgresql](https://www.postgresql.org/) database migrating tool.
It's written using [Golang](https://go.dev/dl/).


## Getting Started

### Prerequisites
- [postgres](https://www.postgresql.org/) DB drive
- [Golang](https://go.dev/dl/) programming language

### Installation

```
git clone https://github.com/praveenmahasena/sqlTool.git
cd sqlTool
make app
```

### How to use it?
Make sure u create a yaml file with the named *psql.yaml* in the working or your project directory
The yaml file should contain these many fields in it

```
host: #psql host
port: # port ID your postgresql is running on
user: #DB user
password: #DB password
dbname: #DB name
sslmode:
```

Then by giving the proper cli args **sqlTool up** or **sqlTool down** you could make the sql migration happen
