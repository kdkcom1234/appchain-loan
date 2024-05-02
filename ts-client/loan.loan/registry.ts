import { GeneratedType } from "@cosmjs/proto-signing";
import { MsgUpdateParams } from "./types/loan/loan/tx";
import { Params } from "./types/loan/loan/params";
import { QueryGetLoanRequest } from "./types/loan/loan/query";
import { MsgUpdateParamsResponse } from "./types/loan/loan/tx";
import { GenesisState } from "./types/loan/loan/genesis";
import { QueryParamsRequest } from "./types/loan/loan/query";
import { QueryGetLoanResponse } from "./types/loan/loan/query";
import { QueryAllLoanResponse } from "./types/loan/loan/query";
import { Loan } from "./types/loan/loan/loan";
import { MsgRequestLoanResponse } from "./types/loan/loan/tx";
import { MsgRequestLoan } from "./types/loan/loan/tx";
import { QueryParamsResponse } from "./types/loan/loan/query";
import { QueryAllLoanRequest } from "./types/loan/loan/query";

const msgTypes: Array<[string, GeneratedType]>  = [
    ["/loan.loan.MsgUpdateParams", MsgUpdateParams],
    ["/loan.loan.Params", Params],
    ["/loan.loan.QueryGetLoanRequest", QueryGetLoanRequest],
    ["/loan.loan.MsgUpdateParamsResponse", MsgUpdateParamsResponse],
    ["/loan.loan.GenesisState", GenesisState],
    ["/loan.loan.QueryParamsRequest", QueryParamsRequest],
    ["/loan.loan.QueryGetLoanResponse", QueryGetLoanResponse],
    ["/loan.loan.QueryAllLoanResponse", QueryAllLoanResponse],
    ["/loan.loan.Loan", Loan],
    ["/loan.loan.MsgRequestLoanResponse", MsgRequestLoanResponse],
    ["/loan.loan.MsgRequestLoan", MsgRequestLoan],
    ["/loan.loan.QueryParamsResponse", QueryParamsResponse],
    ["/loan.loan.QueryAllLoanRequest", QueryAllLoanRequest],
    
];

export { msgTypes }