Civil Servant Warrior Backend

How to make migrations:

Install sql-migrate:
https://github.com/rubenv/sql-migrate

Run:
sql-migrate new -env=yourenv create-yourtables-table
sql-migrate up -env=yourenv

To generate model from database (must be migrated first!):

Install:
go install gorm.io/gen/tools/gentool@latest

To generate models from table (details see pkg/modelbuilder/readme.md):
./modelbuilder -u username -p password -d csw -t newtable1,newtable2 -db 127.0.0.1

To generate repository and service from models (details see pkg/repogenerator/readme.md):
./repogenerator