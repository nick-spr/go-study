## 7. Инверсия булевой карты

**Название:** Инвертировать значения флагов

**Описание:**  
Есть карта `map[string]bool`, которая хранит флаги, включена ли какая-то настройка. Нужно вернуть новую карту, где все значения будут инвертированы (`true` → `false`, `false` → `true`).

**Ожидаемые входные данные:**
```go
map[string]bool{
    "featureA": true,
    "featureB": false,
    "featureC": true,
}
```

**Ожидаемые выходные данные:**
```go
map[string]bool{
    "featureA": false,
    "featureB": true,
    "featureC": false,
}
```