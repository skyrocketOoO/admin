import { createConnectTransport } from "@connectrpc/connect-web";
import { createPromiseClient } from "@connectrpc/connect";
import { AccountService } from "./account_connect.js";
import { CreateAccountRequest, ListAccountRequest } from "./account_pb.js";
import { Empty } from "./common_pb.js";

// Initialize the client
const transport = createConnectTransport({ 
  baseUrl: "http://127.0.0.1:50051", 
  useBinaryFormat: true 
});

const client = createPromiseClient(AccountService, transport);

// Function to create an account
async function createAccount() {
  try {
    const request = new CreateAccountRequest({
      // Populate request fields as needed
    });
    request.UserName="abc"
    request.Password="123"
    request.DisplayName="hi"

    const response = await client.createAccount(request);
    console.log("Account created:", response);
  } catch (error) {
    console.error("Error creating account:", error);
  }
}

// Function to list accounts
async function listAccounts() {
  try {
    const request = new ListAccountRequest({
      // Populate request fields as needed
    });

    const response = await client.listAccount(request);
    console.log("Accounts list:", response);
  } catch (error) {
    console.error("Error listing accounts:", error);
  }
}

// Call the functions
createAccount();
// listAccounts();
