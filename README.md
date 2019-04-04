# Synopsis


```sh
docker-compose up
```
A running http server is listening on localhost:9001

Then

```sh
curl --header "Content-Type: application/json"   --request POST   --data '{"taskId":"004DFGCdNuyccjCHQ"}' localhost:9001
```

Retrieves

```
{"taskId":"004DFGCdNuyccjCHQ","applicants":{"siderId":"aLfphBDYrkYbDFcAT","score":100,"firstName":"Cinda","lastName":"Meyerson"},"description":"\u003cp\u003eFrichti a besoin de toi pour l'aider dans ses préparations de commandes.\u003c/p\u003e\u003ch4 class=\"tpl-task-description-title\"\u003eTon rôle\u003c/h4\u003e\u003cp\u003eTri/Rangement\u003cbr /\u003eTrier, répartir et ranger les produits selon des règles pré-déterminées.\u003c/p\u003e\u003cp\u003ePicking\u003cbr /\u003ePrélever et rassembler les produits de manière ordonnée.\u003c/p\u003e\u003cp\u003ePacking\u003cbr /\u003eAgencer les produits selon une charte, mettre sous blister, conditionner dans un carton/une enveloppe.\u003c/p\u003e\u003cp\u003eTu n\u0026#x27;auras pas à porter de charges.\u003cbr /\u003eTu seras debout à un poste fixe.\u003c/p\u003e\u003ch4 class=\"tpl-task-description-title\"\u003eTes objectifs\u003c/h4\u003e\u003cp\u003eRanger les sacs dans l\u0026#x27;etagere, restocking, etc. \u003c/p\u003e\u003ch4 class=\"tpl-task-description-title\"\u003eTu devras apporter\u003c/h4\u003e\u003cp\u003eDes chaussures confortables\u003c/p\u003e\u003cp\u003eDes vêtements confortables\u003c/p\u003e\u003cp\u003eUne veste chaude\u003c/p\u003e\u003ch4 class=\"tpl-task-description-title\"\u003eTu devras arriver\u003c/h4\u003e\u003cp\u003e5 minutes à l\u0026#x27;avance pour être pro.\u003c/p\u003e\u003ch4 class=\"tpl-task-description-title\"\u003eLe profil recherché\u003c/h4\u003e\u003cp\u003eUn(e) Sider : \u003cbr /\u003e- sachant faire preuve de concentration, persévérance, minutie et rigueur ;\u003cbr /\u003e- disponible sur le maximum de créneaux ;\u003cbr /\u003e- avec de préférence une ou plusieurs expériences similaires.\u003c/p\u003e","country":"FR","tags":["Log | Packer"]}
```
