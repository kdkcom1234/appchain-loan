import { GeneratedType } from "@cosmjs/proto-signing";
import { QueryParamsRequest } from "./types/loan/loan/query";
import { QueryGetLoanResponse } from "./types/loan/loan/query";
import { QueryAllLoanResponse } from "./types/loan/loan/query";
import { Loan } from "./types/loan/loan/loan";
import { QueryGetLoanRequest } from "./types/loan/loan/query";
import { MsgUpdateParams } from "./types/loan/loan/tx";
import { MsgUpdateParamsResponse } from "./types/loan/loan/tx";
import { QueryAllLoanRequest } from "./types/loan/loan/query";
import { GenesisState } from "./types/loan/loan/genesis";
import { Params } from "./types/loan/loan/params";
import { QueryParamsResponse } from "./types/loan/loan/query";

const msgTypes: Array<[string, GeneratedType]>  = [
    ["/loan.loan.QueryParamsRequest", QueryParamsRequest],
    ["/loan.loan.QueryGetLoanResponse", QueryGetLoanResponse],
    ["/loan.loan.QueryAllLoanResponse", QueryAllLoanResponse],
    ["/loan.loan.Loan", Loan],
    ["/loan.loan.QueryGetLoanRequest", QueryGetLoanRequest],
    ["/loan.loan.MsgUpdateParams", MsgUpdateParams],
    ["/loan.loan.MsgUpdateParamsResponse", MsgUpdateParamsResponse],
    ["/loan.loan.QueryAllLoanRequest", QueryAllLoanRequest],
    ["/loan.loan.GenesisState", GenesisState],
    ["/loan.loan.Params", Params],
    ["/loan.loan.QueryParamsResponse", QueryParamsResponse],
    
];

export { msgTypes }