SELECT DISTINCT users.id, users.name
FROM users INNER JOIN orders
ON users.id = orders.user_id
ORDER BY name;