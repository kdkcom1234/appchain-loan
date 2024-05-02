import { Client } from "../../ts-client"
import { AccountData } from "@cosmjs/proto-signing";


const client = new Client(
  {
    apiURL: "http://localhost:1317",
    rpcURL: "http://localhost:26657",
    prefix: "loan",
  },
);

let account: AccountData | undefined;

document.querySelector("#connect")?.addEventListener("click", async () => {
  await client.useKeplr()
  const accounts = (await client.signer?.getAccounts())
  account = accounts ? accounts[0] : undefined
  console.log(client.signer?.getAccounts)

  console.log(account);

  const balances = await client.CosmosBankV1Beta1.query.queryAllBalances(account!.address);
  console.log(balances.data);
})







