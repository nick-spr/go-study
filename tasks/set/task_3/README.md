## 3. Объединение множеств

**Название:** Объединить два множества  
**Описание:**  
Даны множества строк `A` и `B` в виде `map[string]bool{}`. Нужно создать новое множество, содержащее все уникальные элементы из обоих множеств.

**Ожидаемые входные данные:**
```go
A = map[string]bool{"go":true, "java":true, "rust":true}
B = map[string]bool{"rust":true, "python":true}
```

**Ожидаемые выходные данные:**
```go
map[string]bool{"go":true, "java":true, "rust":true, "python":true}
```