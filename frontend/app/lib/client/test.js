"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
var AccountServiceClientPb_1 = require("./AccountServiceClientPb"); // Replace with the correct path to your generated client file
var account_pb_1 = require("./account_pb"); // Replace with the correct path to your generated message file
// Define your gRPC-Web client
var client = new AccountServiceClientPb_1.AccountClient('http://localhost:50051', null, null);
// Prepare the request message
var request = new account_pb_1.CreateAccountRequest();
request.setUsername('username');
request.setPassword('password');
request.setDisplayname('Display Name');
// Make the gRPC call
client.createAccount(request, {}, function (err, response) {
    if (err) {
        console.error('Error:', err.message);
        // Handle error
    }
    else {
        console.log('Response:', response.toObject());
        // Handle success
    }
});
