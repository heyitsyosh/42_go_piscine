NAME	= bsq
GMOD	= go mod init
GBUILD 	= go build .

all: $(NAME)
$(NAME):
	$(GMOD) $(NAME)
	$(GBUILD)

re: clean all
clean:
	rm -f $(NAME) 
	rm -f go.mod
.PHONY: all clean re