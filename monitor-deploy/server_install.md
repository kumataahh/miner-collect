

### mysql

```
// Installing MySQL
sudo apt install mysql-server

// (Optional) Adjusting User Authentication and Privileges
sudo mysql

mysql> SELECT user,authentication_string,plugin,host FROM mysql.user;
mysql> ALTER USER 'root'@'localhost' IDENTIFIED WITH mysql_native_password BY 'root';
mysql> FLUSH PRIVILEGES;
mysql> SELECT user,authentication_string,plugin,host FROM mysql.user;
mysql> exit

sudo mysql
mysql -u root -p

mysql> CREATE USER 'admin'@'localhost' IDENTIFIED BY 'admin';
mysql> GRANT ALL PRIVILEGES ON *.* TO 'admin'@'localhost' WITH GRANT OPTION;
mysql> exit

// Testing MySQL
systemctl status mysql.service
mysql -u admin -p
```