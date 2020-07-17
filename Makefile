TARGET=mysql-diff

clean:
	@echo "clean ..."
	@rm -rf $(TARGER)
all:
	@echo "build all"
	@go build -o mysql-diff src/main/main.go

rebuild: clean all
