# Go modules

Install :

- (OPTIONAL) To generate swagger, `make installSwag` (api)
- (OPTIONAL) To generate protos, run `make`
- `go build` in each directory

See <https://github.com/alextanhongpin/go-protobuf-example> and <https://github.com/simplesteph/protobuf-example-go> for examples on how to use protobuf in go projects

See <https://github.com/gogo/grpc-example> for complete example using go, grpc and swagger auto generated

See <https://github.com/gogo/protobuf> for protobuf gen

See <https://jbrandhorst.com/post/go-protobuf-tips/> for protobuf tips

Messages published in Pulsar have this structure:

```json
{
    "profile": {
        "gender": "male",
        "name": {
            "first": "William",
            "last": "Brown",
            "title": "Mr"
        },
        "location": {
            "street": "32 Franklin St",
            "city": "Brandwell",
            "state": "Florida",
            "postcode": 30006
        },
        "email": "william.brown@example.com",
        "login": {
            "username": "Handsurf",
            "password": "Stingerpalm",
            "salt": "x2bA1E96GsShnUp9",
            "md5": "fe38204ff77ca7657821f3f5a5d99acc",
            "sha1": "4QN0GAD1dAT_CxVsJ0qliqwYhLA=",
            "sha256": "ASmpMZN3C5eZPDKTK1GBNO-Q7GZMsG9o8P923948yRo="
        },
        "dob": "Monday 30 Mar 2020",
        "registered": "Friday 14 Feb 2020",
        "phone": "+370 66 4617 36 76 2",
        "cell": "+44 77 2 252280 66 5",
        "id": {
            "name": "SSN"
        },
        "picture": {
            "large": "https://randomuser.me/api/portraits/men/16.jpg",
            "medium": "https://randomuser.me/api/portraits/med/men/16.jpg",
            "thumbnail": "https://randomuser.me/api/portraits/thumb/men/16.jpg"
        },
        "nat": "US"
    }
}
```
