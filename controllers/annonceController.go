package controllers

import (
	"errors"
	"fmt"
	dbCon "github.com/franck-djacoto/announce-service/db-connection"
	. "github.com/franck-djacoto/announce-service/models"
	. "github.com/franck-djacoto/announce-service/repository"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"regexp"
	"strconv"
)

type AnnonceController struct {
	catRepo CategoryRepository
	mapRepo MarqueRepository
	modRepo ModelRepository
	model   ModelRepository
	anRepo  AnnonceRepository
}

type SearchForm struct {
	Modele string `json:"modele" binding:"required"`
}

func (AnCon *AnnonceController) New(dbConnet *dbCon.DbConnection) *AnnonceController {
	return &AnnonceController{
		catRepo: CategoryRepository{DbConnect: dbConnet},
		mapRepo: MarqueRepository{DbConnect: dbConnet},
		anRepo:  AnnonceRepository{DbConnect: dbConnet},
		model:   ModelRepository{DbConnect: dbConnet},
	}
}

func (AnCon *AnnonceController) ChecIfServiceRespond(c *gin.Context) {
	c.JSON(
		200,
		gin.H{
			"data": "Hello world",
		},
	)
}
func (AnCon *AnnonceController) validateId(id string) (int, error) {
	regexId := regexp.MustCompile("[1-9]+")

	if !regexId.MatchString(id) {
		return 0, errors.New("Invalid Id")
	}

	idToInt, err := strconv.Atoi(id)

	if err != nil {
		return 0, err
	}

	return idToInt, nil
}

func (AnCon *AnnonceController) getAnnonceAddionalInfo(annonce Annonce, c *gin.Context) (allIds map[string]int, errorMap map[string]string, err error) {
	var errorsMap map[string]string
	catId, err := AnCon.catRepo.GetCatIdByLibelle(annonce.Categorie)
	if err != nil {
		return nil, nil, err
	}
	if catId == 0 {
		errorsMap["Categorie"] = "La catégorie" + annonce.Categorie + "est invalide"
	}

	maqId, err := AnCon.mapRepo.GetMarqueIdByLibelle(annonce.Marque)
	if err != nil {
		return nil, nil, err
	}
	if maqId == 0 {
		errorsMap["Marque"] = "La marque " + annonce.Marque + "est invalide"
	}

	mdId, err := AnCon.model.GetModelIdByLibelle(annonce.Modele)
	if err != nil {
		return nil, nil, err
	}
	if mdId == 0 {
		errorsMap["Model"] = "Le model " + annonce.Modele + "est invalide"
	}

	if len(errorsMap) > 0 {
		return nil, errorsMap, nil
	}

	allIds = map[string]int{
		"categorie": catId,
		"marque":    maqId,
		"model":     mdId,
	}
	return allIds, nil, nil
}

func (AnCon *AnnonceController) Add(c *gin.Context) {
	var annonce Annonce
	var errorsMap map[string]string
	err := c.ShouldBindJSON(&annonce)

	if err != nil {
		log.Printf("Error while binding ===> %v", err)
		c.AbortWithStatusJSON(http.StatusBadRequest,
			gin.H{"error": fmt.Sprintf("Value sent are incorrect : %v", err)},
		)
		return
	}

	if annonce.Categorie == "Automobile" && (annonce.Modele == "" || annonce.Marque == "") {
		c.AbortWithStatusJSON(
			http.StatusBadRequest,
			gin.H{
				"errors": "Vous devez renseigner la marque et le modele de l'automobile",
			},
		)
		return
	}

	allIds, errorsMap, err := AnCon.getAnnonceAddionalInfo(annonce, c)

	if err != nil {
		log.Printf("Error while getting Annonce additional Infor ===>  %v", err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("%v", err) })
		return
	}

	if len(errorsMap) > 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"errors": errorsMap})
		return
	}

	annonceModel := AnnonceModel{
		Titre:       annonce.Titre,
		Contenu:     annonce.Contenu,
		CategorieId: allIds["categorie"],
		MarqueId:    allIds["marque"],
		ModeleId:    allIds["model"],
	}

	idAnnonce, err := AnCon.anRepo.Save(annonceModel)

	if err != nil {
		log.Printf("Error while saving annonce to db ===>   %v", err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("%v", err)})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Annonce créé avec success",
		"id":      idAnnonce,
	})
}

func (AnCon *AnnonceController) All(c *gin.Context) {
	annonces, err := AnCon.anRepo.GetAll()

	if err != nil {
		log.Printf("Error while getting annonce from db ===>   %v", err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("%v", err)})
		return
	}

	if len(annonces) > 0 {
		c.JSON(http.StatusOK, gin.H{"annonces:": annonces})
	} else {
		c.JSON(http.StatusOK, gin.H{"annonces:": "Aunce annonce trouvée"})
	}
}

func (AnCon *AnnonceController) Detail(c *gin.Context) {
	id := c.Param("id")
	idToInt, err := AnCon.validateId(id)

	if err != nil {
		log.Printf("Error while validation id ===>   %v", err)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("%v", err)})
		return
	}
	annonce, err := AnCon.anRepo.GetById(idToInt)

	if err != nil {
		log.Printf("Error while getting annonce for id %d ===>   %v",idToInt, err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("%v", err)})
		return
	}

	if annonce.Id > 0 {
		c.JSON(http.StatusOK, gin.H{"annonce:": annonce})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"message:": "Annonce inexistante"})
	}
}

func (AnCon *AnnonceController) Update(c *gin.Context) {
	id := c.Param("id")
	idToInt, err := AnCon.validateId(id)

	if err != nil {
		log.Printf("Error while validation id ===>   %v", err)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("%v", err)})
		return
	}

	var annonce Annonce

	err = c.ShouldBindJSON(&annonce)

	if err != nil {
		log.Printf("Error while binding ===> %v", err)
		c.AbortWithStatusJSON(http.StatusBadRequest,
			gin.H{"error": fmt.Sprintf("Value sent are incorrect : %v", err)},
		)
		return
	}

	if annonce.Categorie == "Automobile" && (annonce.Modele == "" || annonce.Marque == "") {
		c.AbortWithStatusJSON(
			http.StatusBadRequest,
			gin.H{
				"errors": "Vous devez renseigner la marque et le modele de l'automobile",
			},
		)
		return
	}

	annonceToModifiy, err := AnCon.anRepo.GetById(idToInt)
	if err != nil {
		log.Printf("Error while getting annonce for id %d ===>   %v",idToInt, err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("%v", err)})
		return
	}

	if annonceToModifiy.Id > 0 {

		allIds, errorsMap, err := AnCon.getAnnonceAddionalInfo(annonce, c)

		if err != nil {
			log.Printf("Error while getting Annonce additional Infor ===>  %v", err)
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("%v", err)})
			return
		}

		if len(errorsMap) > 0 {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"errors":  errorsMap})
			return
		}

		annonceModel := AnnonceModel{
			Id:          int64(idToInt),
			Titre:       annonce.Titre,
			Contenu:     annonce.Contenu,
			CategorieId: allIds["categorie"],
			MarqueId:    allIds["marque"],
			ModeleId:    allIds["model"],
		}

		isUpdated, err := AnCon.anRepo.Update(annonceModel)

		if err != nil {
			log.Printf("Error while updating annonce ===>  %v", err)
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("%v", err)})
			return
		}

		if isUpdated {
			c.JSON(http.StatusOK, gin.H{"message:": "Annonce modifié"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"message:": "Une erreur est survenue lors de la modification de l'annonce. Veuillez eessayer"})
		}

	} else {
		c.JSON(http.StatusBadRequest, gin.H{"message:": "Annonce inexistante"})
	}
}

func (AnCon *AnnonceController) Delete(c *gin.Context) {
	id := c.Param("id")
	idToInt, err := AnCon.validateId(id)

	if err != nil {
		log.Printf("Error while validation id ===>   %v", err)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("%v", err)})
		return
	}

	annonce, err := AnCon.anRepo.GetById(idToInt)

	if err != nil {
		log.Printf("Error while getting annonce for id %d ===>   %v",idToInt, err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("%v", err)})
		return
	}

	if annonce.Id > 0 {
		isDeleted, err := AnCon.anRepo.Delete(idToInt)

		if err != nil {
			log.Printf("Error while deleting annonce ===>  %v", err)
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("%v", err)})
			return
		}

		if isDeleted {
			c.JSON(http.StatusOK, gin.H{"message:": "Annonce supprimé"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"message:": "Une erreur est survenue lors de la suppression de l'annonce. Veuillez eessayer"})
		}

	} else {
		c.JSON(http.StatusBadRequest, gin.H{"message:": "Annonce inexistante"})
	}
}

func (AnCon *AnnonceController) Search(c *gin.Context) {
	modelRegex := regexp.MustCompile("^[a-zA-Z0-9\\s]*$")
	var sarchForm SearchForm

	err := c.ShouldBindJSON(&sarchForm)

	if err != nil {
		log.Printf("Error occured while binding searchForm  %v",  err)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("%v", err)})
		return
	}

	if !modelRegex.MatchString(sarchForm.Modele) {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Modèle invalide"})
		return
	}
	annonces, err := AnCon.anRepo.GetByModelName(sarchForm.Modele)

	if err != nil {
		log.Printf("Error while retreiving annonce for model %s ===>   %v", sarchForm.Modele, err)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("%v", err)})
		return
	}

	if len(annonces) > 0 {
		c.JSON(http.StatusOK, gin.H{"message":fmt.Sprintf("Annonce correspondant au  modele %s", sarchForm.Modele),"annonces": annonces})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"message:": fmt.Sprintf("Aunce annonce trouvé pour le modèle %s", sarchForm.Modele)})
	}
}
