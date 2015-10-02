import static spark.Spark.*;

public class WoodbrickOven {
    public static void main(String[] args) {
        port(8082);
        get("/cook", (req, res) -> {
          Thread.sleep(3);
          return "done";
        });
    }
}
