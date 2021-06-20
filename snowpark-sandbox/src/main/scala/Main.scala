import com.snowflake.snowpark._
import com.snowflake.snowpark.functions._

object Main {
  def main(args: Array[String]): Unit = {
    val url = sys.env.get("SNOWFLAKE_URL").getOrElse("")
    val password = sys.env.get("SNOWFLAKE_PASSWORD").getOrElse("")
    val builder = Session.builder.configs(Map(
      "URL" -> url,
      "USER" -> "KANATA2",
      "PASSWORD" -> password,
      "WAREHOUSE" -> "COMPUTE_WH",
      "DB" -> "SANDBOX",
      "SCHEMA" -> "KANATA2_SANDBOX",
    ))
    val session = builder.create
    var dfTables = session.table("INFORMATION_SCHEMA.TABLES")
      .filter(col("TABLE_SCHEMA") === "KANATA2_SANDBOX")
    var tableCount = dfTables.count()
    var currentDB = session.getCurrentDatabase
    println(s"Number of tables in the $currentDB database: $tableCount")

    var dfPublicSchemaTables = session.table("INFORMATION_SCHEMA.TABLES")
      .filter(col("TABLE_SCHEMA") === "KANATA2_SANDBOX")
      .select(col("TABLE_NAME"))
    dfPublicSchemaTables.show()
  }
}
