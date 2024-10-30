MODULE ?=
OUTPUT ?=

proto:
	@echo "Compilando arquivos proto..."
ifeq ($(MODULE),)
	@echo "Erro: Por favor, forneça o nome do módulo.\nExemplo: make proto MODULE=<module_name> OUTPUT=<output_path>"
else ifeq ($(OUTPUT),)
    @echo "Erro: Por favor, forneça o path do output.\nExemplo: make proto MODULE=<module_name> OUTPUT=<output_path>"
else
	@protoc -I=$(MODULE) --go-grpc_out=$(OUTPUT) --go_out=$(OUTPUT) $(MODULE)/*.proto
	@echo "Arquivos do módulo '$(MODULE)' compilados com sucesso!"
endif
