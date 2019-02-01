# MySQL to Mongo migration


### How To Run This Project
```bash
#move to directory
cd $GOPATH/src/github.com/<username>

# Clone into YOUR $GOPATH/src
git clone https://github.com/samayamnag/go-mysql-to-mongo-migration.git <project_name>

#move to project
cd <project_name>

#copy config and ensure all values are correct
cp config_example.json config/config.json

# Install Dependencies
go build .

# Run Project
go run main.go -fromId=xxx -toId=xxx

```

Or

```bash
# GET WITH GO GET
go get github.com/samayamnag/go-mysql-to-mongo-migration

# Go to directory

cd $GOPATH/src/github.com/go-mysql-to-mongo-migration

# Install Dependencies
go build .


# Run Project
go run main.go -fromId=xxx -toId=xxx
```