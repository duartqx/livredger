package application

var NovaUnidadeDeTrabalho fabricaDeUnidadeDeTrabalho

func SetupApplication(fabrica fabricaDeUnidadeDeTrabalho) error {
	if NovaUnidadeDeTrabalho == nil {
		NovaUnidadeDeTrabalho = fabrica
	}

	return nil
}
