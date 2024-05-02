import { GeneratedType } from "@cosmjs/proto-signing";
import { MsgRequestLoanResponse } from "./types/loan/loan/tx";
import { QueryParamsRequest } from "./types/loan/loan/query";
import { Loan } from "./types/loan/loan/loan";
import { MsgUpdateParamsResponse } from "./types/loan/loan/tx";
import { MsgRequestLoan } from "./types/loan/loan/tx";
import { MsgApproveLoan } from "./types/loan/loan/tx";
import { QueryParamsResponse } from "./types/loan/loan/query";
import { QueryGetLoanRequest } from "./types/loan/loan/query";
import { GenesisState } from "./types/loan/loan/genesis";
import { Params } from "./types/loan/loan/params";
import { MsgApproveLoanResponse } from "./types/loan/loan/tx";
import { QueryGetLoanResponse } from "./types/loan/loan/query";
import { QueryAllLoanRequest } from "./types/loan/loan/query";
import { QueryAllLoanResponse } from "./types/loan/loan/query";
import { MsgUpdateParams } from "./types/loan/loan/tx";

const msgTypes: Array<[string, GeneratedType]>  = [
    ["/loan.loan.MsgRequestLoanResponse", MsgRequestLoanResponse],
    ["/loan.loan.QueryParamsRequest", QueryParamsRequest],
    ["/loan.loan.Loan", Loan],
    ["/loan.loan.MsgUpdateParamsResponse", MsgUpdateParamsResponse],
    ["/loan.loan.MsgRequestLoan", MsgRequestLoan],
    ["/loan.loan.MsgApproveLoan", MsgApproveLoan],
    ["/loan.loan.QueryParamsResponse", QueryParamsResponse],
    ["/loan.loan.QueryGetLoanRequest", QueryGetLoanRequest],
    ["/loan.loan.GenesisState", GenesisState],
    ["/loan.loan.Params", Params],
    ["/loan.loan.MsgApproveLoanResponse", MsgApproveLoanResponse],
    ["/loan.loan.QueryGetLoanResponse", QueryGetLoanResponse],
    ["/loan.loan.QueryAllLoanRequest", QueryAllLoanRequest],
    ["/loan.loan.QueryAllLoanResponse", QueryAllLoanResponse],
    ["/loan.loan.MsgUpdateParams", MsgUpdateParams],
    
];

export { msgTypes }