# MongoDB Connection Builder

Esta librería proporciona una forma sencilla de construir conexiones a MongoDB utilizando un patrón de constructor (builder pattern) en Go.

## Estructura del Proyecto

- `mongodb.go`: Contiene la implementación del constructor para la conexión a MongoDB.
- `go.mod` y `go.sum`: Archivos de gestión de dependencias de Go.
- `LICENSE`: Licencia del proyecto.

## Uso

### Crear un Constructor de Conexión

Para crear un constructor de conexión a MongoDB, utiliza la función `NewBuilder` y configura las credenciales y el endpoint:

```go
builder := mongodb.NewBuilder("host", "port", "username", "password")
```

### Configurar TLS

Si necesitas configurar TLS para la conexión, utiliza el método `WithTLS`:

```go
builder = builder.WithTLS("base64EncodedCACertificate")
```

### Configurar Retry Writes

Para configurar el parámetro `retryWrites`, utiliza el método `WithRetryWrites`:

```go
builder = builder.WithRetryWrites(true)
```

### Obtener el Cliente de MongoDB

Para obtener un cliente de MongoDB, utiliza el método `GetClient`:

```go
ctx := context.TODO()
client, err := builder.GetClient(ctx)
if err != nil {
    log.Fatalf("Error creando el cliente: %v", err)
}
defer client.Disconnect(ctx)
```

## Licencia

Este proyecto está licenciado bajo la **Licencia Pública General Affero de GNU, versión 3 (AGPLv3)**.

Esto significa que:
1. Puedes usar, copiar, modificar y distribuir este software libremente.
2. Si realizas modificaciones y utilizas este software como parte de un servicio público (por ejemplo, una API, aplicación web o SaaS), estás obligado a:
   - Liberar el código fuente de las modificaciones realizadas.
   - Asegurarte de que las modificaciones estén disponibles bajo los mismos términos de la AGPLv3.

Puedes encontrar el texto completo de la licencia en el archivo [LICENSE](./LICENSE) o consultarlo en el sitio oficial de GNU:  
[https://www.gnu.org/licenses/agpl-3.0.html](https://www.gnu.org/licenses/agpl-3.0.html).
