package helpers

/* import (
	"EOP/database"

	"github.com/gofiber/fiber/v2"
	"github.com/wpcodevo/golang-fiber-jwt/models"
)

func Authorize(c *fiber.Ctx, subject string, action string) error {
	if subject == "" && action == "" {
		return nil
	}
	var sql = `
	Select count(*) from users u
	LEFT JOIN core.user_actions cas ON u.id = cas.user_id
	LEFT JOIN core.module_actions cma on cas.system_module_action_id = cma.id
	LEFT JOIN core.module m on m.id = cma.module_id
	WHERE u.id = ? AND m.path = ? AND cma.key = ?;`
	var result int
	database.DB.Db.Raw(sql, c.Locals("user").(models.User).ID, subject, action).Scan(&result)
	if result == 0 {
		if c.Locals("user").(models.User).IsAdmin {
			return nil
		}
		return fiber.NewError(fiber.StatusForbidden, "Bu eylemi gerçekleştirme yetkiniz yok.")
	}
	return nil
}
*/
