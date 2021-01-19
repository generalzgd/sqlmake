# sqlmake


1. 
INSERT INTO `table_name` (`field`,`field`)VALUES(1,2,3) ON DUPLICATE KEY UPDATE `field`=VALUES(`field`);

2. 
INSERT INTO `table_name` (`a`,`b`,`c`) VALUES (1,'a',4),(2,'5h',41) ON DUPLICATE KEY UPDATE `a`=VALUES(`a`);

3. 
UPDATE `table_name` SET `abc` =CASE `id` WHEN 1 THEN '0' WHEN 34 THEN '35' END;