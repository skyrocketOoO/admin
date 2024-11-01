import { createClient, Client } from "@connectrpc/connect";
import { createConnectTransport } from "@connectrpc/connect-web";
import { ListAccountReq, ListAccountResp } from "@/proto/main_pb";
import { MainService } from "@/proto/main_connect";

const serverSideTransport = createConnectTransport({
  baseUrl: "http://localhost:50051",
});

const clientSideTransport = createConnectTransport({
  baseUrl: "https://demo.connectrpc.com",
});

export const serverSideClient = createClient(MainService, serverSideTransport);
export const clientSideClient = createClient(MainService, clientSideTransport);