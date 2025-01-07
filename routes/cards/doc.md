Carregamento dos Dados - Cards

- Separamos o carregamento por categorias (Clima, Vegetação, Parques e Praças)
- A principio, cada categoria tem 4 seções (carregadas por anos diferentes, com filtro)

Clima ex.
- v1/cards/weather/temperature?city=3549003&year=2024
- v1/cards/weather/heat?city=3549003&year=2023
- v1/cards/weather/ranking?city=3549003&year=2022
- v1/cards/weather/sec4?city=3549003

Cada seção tem dados definidos - dentro de suas categorias:

Clima:
- Temperatura e Clima
	Nivel de ilha
	Temperatura média
	Maior amplitude
	Temperatura máxima

- Calor Extremo
	Negros e indigenas
	Mulheres
	Crianças
	Idosos

- Ranking
	Brasil
		Nivel de ilha
		Temperatura média
		Temperatura máxima
	Sudeste
		Nivel de ilha
		Temperatura média
		Temperatura máxima
	Estado
		Nivel de ilha
		Temperatura média
		Temperatura máxima
	Mesorregião
		Nivel de ilha
		Temperatura média
		Temperatura máxima
	Microrregião
		Nivel de ilha
		Temperatura média
		Temperatura máxima

- Veja Mais
	Media cobertura vegetal
	Moradores prox praças
	Desigualdade amb e social


Temperatura e Clima:
[
  {
    "title":"Nível de ilha de calor",
    "subtitle":"Está na média nacional de 0",
    "value":"0"
  },
  {
    "title":"Temperatura média da superfície",
    "subtitle":"Acima da média nacional de 0",
    "value":"32°C"
  },
  {
    "title":"Maior amplitude",
    "subtitle":"É a diferença entre a temperatura mais quente e a mais fria",
    "value":"8°C"
  },
  {
    "title":"Temperatura máxima da superfície",
    "value":"41°C"
  }
]
    

Calor extremo:
[
  {
    "title":"Negros e indígenas afetados",
    "subtitle":"Porcentagem vivendo nas regiões mais quentes",
    "value":"27%"
  },
  {
    "title":"Mulheres afetadas",
    "subtitle":"Porcentagem vivendo nas regiões mais quentes",
    "value":"38%"
  },
  {
    "title":"Crianças afetadas",
    "subtitle":"Porcentagem vivendo nas regiões mais quentes",
    "value":"13%"
  },
  {
    "title":"Idosos afetados",
    "subtitle":"Porcentagem vivendo nas regiões mais quentes",
    "value":"18%"
  }
]
    
Ranking:
[
  {
    "title": "Municipios do Brasil",
    "subtitle": "Posição do seu município entre os 5568 do Brasil",
    "data": [
      {
        "title": "Nível de ilha de calor",
        "number": 2862,
        "of": 5568
      },
      {
        "title": "Temperatura média da superfície",
        "number": 298,
        "of": 5568
      },
      {
        "title": "Temperatura máxima da superfície",
        "number": 698,
        "of": 5568
      }
    ]
  },
  
  {
    "title": "Municipios da região Sudeste",
    "subtitle": "Posição do seu município entre os 1668 da região Sudeste",
    "data": [
      {
        "title": "Nível de ilha de calor",
        "number": 2862,
        "of": 5568
      },
      {
        "title": "Temperatura média da superfície",
        "number": 298,
        "of": 5568
      },
      {
        "title": "Temperatura máxima da superfície",
        "number": 698,
        "of": 5568
      }
    ]
  },
  
  {
    "title": "Municipios do Estado",
    "subtitle": "Posição do seu município entre os 5568 do Estado de São Paulo",
    "data": [
      {
        "title": "Nível de ilha de calor",
        "number": 2862,
        "of": 5568
      },
      {
        "title": "Temperatura média da superfície",
        "number": 298,
        "of": 5568
      },
      {
        "title": "Temperatura máxima da superfície",
        "number": 698,
        "of": 5568
      }
    ]
  },
  
  {
    "title": "Municipios da Mesorregião",
    "subtitle": "Posição do seu município entre os 5568 da mesorregião de Araraquara",
    "data": [
      {
        "title": "Nível de ilha de calor",
        "number": 2862,
        "of": 5568
      },
      {
        "title": "Temperatura média da superfície",
        "number": 298,
        "of": 5568
      },
      {
        "title": "Temperatura máxima da superfície",
        "number": 698,
        "of": 5568
      }
    ]
  },
  
  {
    "title": "Municipios da Microrregião",
    "subtitle": "Posição do seu município entre os 5568 da microrregião de São Carlos",
    "data": [
      {
        "title": "Nível de ilha de calor",
        "number": 2862,
        "of": 5568
      },
      {
        "title": "Temperatura média da superfície",
        "number": 298,
        "of": 5568
      },
      {
        "title": "Temperatura máxima da superfície",
        "number": 698,
        "of": 5568
      }
    ]
  }
]


Veja Mais:
[
  {
    "title": "Média da cobertura vegetal",
    "value": 41
  },
  {
    "type": "Moradores próximos a praças",
    "value": 84
  },
  {
    "type": "Desigualdade ambiental e social",
    "value": 63
  }
]

- Cada seção tem um endpoint diferente para carregamento