# How to run the project
1. Clone the project  
2. Being at the root of the project, just run `docker-compose up`
***

# Making request
## Create an announce

````
curl -X POST  http://127.0.0.1:8000/api/annonce/add \
-H 'Content-Type: application/json' \
-d '{"titre" : "Voiture mise en vente", "contenu" : "Vente de voiture BMW", "categorie" : "Automobile", "marque" : "BMW", "modele" : "M4" }' 
````
### Reponse  Create annonce
````
{
    "id": 1,
    "message": "Annonce créé avec success"
}
```` 
***

## Update an announce

````
curl -X POST   http://127.0.0.1:8000/api/annonce/update/1 \
-H 'Content-Type: application/json' \
-d '{"titre" : "Une voiture de luxe", "contenu" : "Vente de voiture BMW", "categorie" : "Automobile", "marque" : "BMW", "modele" : "M4" }'
````
### Reponse  Create annonce
````
{
    "message:": "Annonce modifié"
}
```` 
***

## Retreive all announce

````
curl  http://127.0.0.1:8000/api/annonce/all \
-H "Accept: application/json"
````
### Reponse  list annonunce
````
{
    "annonces:": [
        {
            "id": 1,
            "titre": "Vente d'une voiture",
            "contenu": "Vente de voiture BMW",
            "categorie": "automobile",
            "marque": "bmw",
            "modele": "M4"
        },
        {
            "id": 2,
            "titre": "Vente d'une voiture",
            "contenu": "Vente de voiture BMW",
            "categorie": "automobile",
            "marque": "bmw",
            "modele": "M4"
        }
    ]
}
```` 
***

## Get announce detail

````
curl   http://127.0.0.1:8000/api/annonce/detail/1 \
-H "Accept: application/json"
````
### Reponse  annonce details
````
    "annonce:":  {
            "id": 1,
            "titre": "Voiture mise en vente",
            "contenu": "Vente de voiture BMW",
            "categorie": "automobile",
            "marque": "bmw",
            "modele": "M4"
    }
```` 
***


## Delete Announce

````
curl http://127.0.0.1:8000/api/annonce/delete/1 \
-H "Accept: application/json"
````
### Reponse  annonce details
````
{
    "message:": "Annonce supprimé"
}
```` 
***


## Search annonce by model libelle

````
curl    http://127.0.0.1:8000/api/annonce/search \
-H "Accept: application/json" \
-d '{ "modele" : "avant M4" }'
````
### Reponse  annonce details
````
{
    "annonce": {
        "id": 1,
        "titre": "Vente d'une voiture",
        "contenu": "Vente de voiture BMW",
        "categorie": "automobile",
        "marque": "bmw",
        "modele": "M4"
    }
}
```` 
***