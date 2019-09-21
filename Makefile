run:
	export `cat .env` && go run ./cmd/server/main.go

server:
	protoc --twirp_out=. --go_out=. ./rpc/psu/service.proto

iface:
	ifacemaker -f pkg/store -s Store -i StoreInterface -p store -o pkg/store/store_interface.go
	ifacemaker -f pkg/service/student.go -s StudentService -i StudentService -p serviceiface -o pkg/service/serviceiface/studentservice.go

migrations:
	mysql -u root < migrations/dump.sql