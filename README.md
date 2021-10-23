# About this project 
This project is just a simple REST Api created in Go language that allow you to create, read, update, delete 
an announcement. 

***

## Created with MVC architecture and using docker
Although it's a simple REST Api, I decide to created on an MVC architecture from scratch (yeah it's true the 'V' part is missing here ^ ^). <br/>
I used docker to make it easier to be launched for anyone cloning to project and having docker on is machine.

***

# How to run the project
1. Clone the project  
2. Being at the root of the project, just run `docker-compose up`
***

# Making request
Once your have started the program you can open a `terminal` and make request using `curl` 
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
    "message": "Annonce créée avec success"
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
    "message:": "Annonce modifiée"
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
            "titre": "Voiture de luxe",
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
    "message:": "Annonce supprimée"
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
    "message": "Annonces correspondantes au  modele avant M4", 
    "annonces": [
        {
            "id": 2,
            "titre": "Une voiture de luxe",
            "contenu": "Vente de voiture BMW",
            "categorie": "automobile",
            "marque": "bmw",
            "modele": "M4"
        },
        {
            "id": 3,
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