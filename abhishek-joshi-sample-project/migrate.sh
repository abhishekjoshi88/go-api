DB_USERNAME: "user"
DB_PASSWORD: "user"
DB_NAME: "mydb"
DB_HOST: "localhost"
DB_PORT: "3306"


db_name=$(mysql -u $DB_USERNAME -p$DB_PASSWORD -se "SELECT SCHEMA_NAME FROM INFORMATION_SCHEMA.SCHEMATA WHERE SCHEMA_NAME = '${DB_NAME}'")
 
if [ "$db_name" != "$DB_NAME" ]; then
    mysql -u ${DB_USERNAME} -p${DB_PASSWORD} -Bse "CREATE DATABASE \`${DB_NAME}\`;" && \
	zcat ./tools/01.sql.gz | mysql -u ${DB_USERNAME} -p${DB_PASSWORD} ${DB_NAME}	
fi
 