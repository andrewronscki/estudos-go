- Testar os arquivos de teste naquela pasta
go test
- Testar os arquivos de teste em todas as pastas
go test ./...
- Cobertura de teste para todos os arquivos
go test ./... --cover 
- Gerar arquivo txt de cobertura de teste
go test ./... --coverprofile cover.txt
- Gerar html de cobertura de teste a partir do txt
go tool cover --html=cover.txt