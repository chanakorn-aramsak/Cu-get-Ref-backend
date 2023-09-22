const PROTO_PATH = "./restaurant.proto";
const dotenv = require("dotenv");
const db = require("../config/db");

var grpc = require("@grpc/grpc-js");
var protoLoader = require("@grpc/proto-loader");

//Load env vars
dotenv.config({ path: "./config/config.env" });

//Connect to database
db.connectDB();

var packageDefinition = protoLoader.loadSync(PROTO_PATH, {
  keepCase: true,
  longs: String,
  enums: String,
  arrays: true,
});

var restaurantProto = grpc.loadPackageDefinition(packageDefinition);

const { v4: uuidv4 } = require("uuid");

const server = new grpc.Server();
const menu = [
  {
    id: "a68b823c-7ca6-44bc-b721-fb4d5312cafc",
    name: "Tomyam Gung",
    price: 500,
  },
  {
    id: "34415c7c-f82d-4e44-88ca-ae2a1aaa92b7",
    name: "Somtam",
    price: 60,
  },
  {
    id: "8551887c-f82d-4e44-88ca-ae2a1ccc92b7",
    name: "Pad-Thai",
    price: 120,
  },
];

server.addService(restaurantProto.RestaurantService.service, {
  getAllMenu: (_, callback) => {
    db.getAllMenus()
      .then((result) => {
        if (!result) {
          callback(null, { menu: [] });
        }
        callback(null, { menu: result });
      })
      .catch((e) => {
        callback(e);
      });
  },
  get: (call, callback) => {
    const { id } = call.request;
    db.getMenu(id)
      .then((result) => {
        if (!result) {
          callback({
            code: grpc.status.NOT_FOUND,
            details: "Not found",
          });
        }
        callback(null, result);
      })
      .catch((e) => {
        callback(e);
      });
  },
  insert: (call, callback) => {
    let menuItem = call.request;
    menuItem.id = uuidv4();
    menu.push(menuItem);

    db.insertMenu(menuItem)
      .then(() => {
        callback(null, menuItem);
      })
      .catch((e) => {
        callback(e);
      });
  },
  update: (call, callback) => {
    const menu = call.request;
    const { id } = menu;
    db.updateMenu(id, menu)
      .then(() => {
        callback(null, menu);
      })
      .catch((e) => {
        callback({
          code: grpc.status.NOT_FOUND,
          details: "Not Found",
        });
      });
  },
  remove: (call, callback) => {
    const { id } = call.request;
    db.removeMenu(id)
      .then(() => {
        callback(null, {});
      })
      .catch((e) => {
        callback({
          code: grpc.status.NOT_FOUND,
          details: "NOT Found",
        });
      });
  },
});

server.bindAsync(
  "127.0.0.1:30043",
  grpc.ServerCredentials.createInsecure(),
  () => {
    server.start();
  }
);
console.log("Server running at http://127.0.0.1:30043");
