# IPRO-VAULT-SERVER

## Installing

Requirements:
Go >=1.4


*NOTE: Make sure your $GOPATH is set*

```
go get -u github.com/revel/cmd/revel
go get -u github.com/skrzepto/IPRO-VAULT-SERVER
```

## Running the program

```
revel run github.com/skrzepto/IPRO-VAULT-SERVER
```

now go to localhost:9000


## Running unittests

### Console mode

```
revel test github.com/skrzepto/IPRO-VAULT-SERVER dev
```

### GUI mode
```
revel run github.com/skrzepto/IPRO-VAULT-SERVER
```
go to "http://localhost:9000/@tests" and click on run all tests


