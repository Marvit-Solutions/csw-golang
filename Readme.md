## Civil Servant Warrior Backend

### How to make migrations:

###### Run:
./script/migrate_create.sh -name your_name_table

###### Example:
./script/migrate_create.sh -name alter_default_deleted_at_table

### How to up migrations:

###### Run:
./script/migrate_up.sh -u username -p password -d csw

### How to down migrations:

###### Run:
./script/migrate_down.sh -u username -p password -d csw

### To generate model from database (must be migrated first!):

###### Run: 
./modelbuilder -u username -p password -d csw -t "newtable1,newtable2" -db localhost

### To generate repository and service from models:

###### Run: 
./repogenerator
