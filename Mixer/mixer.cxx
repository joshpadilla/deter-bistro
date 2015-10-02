#include "crow_all.h"

int main()
{
    crow::SimpleApp app;

    CROW_ROUTE(app, "/mix")([](){
        return "dough";
    });

    app.port(8085).multithreaded().run();
}
