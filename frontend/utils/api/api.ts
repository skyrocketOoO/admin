import { createClient, Client } from "@connectrpc/connect";
import { createConnectTransport } from "@connectrpc/connect-web";
import { ListAccountReq, ListAccountResp } from "@/proto/main_pb";
import { Main } from "@/proto/main_connect";

// The transport defines what type of endpoint we're hitting.
// In our example we'll be communicating with a Connect endpoint.
// If your endpoint only supports gRPC-web, make sure to use
// `createGrpcWebTransport` instead.
const serverSideTransport = createConnectTransport({
  baseUrl: "http://localhost:50051",
});

const clientSideTransport = createConnectTransport({
  baseUrl: "https://demo.connectrpc.com",
});

// Here we make the client itself, combining the service
// definition with the transport.
export const serverSideClient = createClient(Main, serverSideTransport);
export const clientSideClient = createClient(Main, clientSideTransport);

export async function listAccount(
  client: Client<typeof Main>,
  req: ListAccountReq
): Promise<ListAccountResp> {
  return await client.listAccount(req) as ListAccountResp;
}
