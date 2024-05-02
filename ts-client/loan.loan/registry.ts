import { GeneratedType } from "@cosmjs/proto-signing";
import { MsgApproveLoan } from "./types/loan/loan/tx";
import { MsgApproveLoanResponse } from "./types/loan/loan/tx";
import { MsgRequestLoan } from "./types/loan/loan/tx";
import { MsgLiquidateLoan } from "./types/loan/loan/tx";
import { Params } from "./types/loan/loan/params";
import { QueryGetLoanRequest } from "./types/loan/loan/query";
import { QueryAllLoanRequest } from "./types/loan/loan/query";
import { MsgRepayLoan } from "./types/loan/loan/tx";
import { QueryParamsRequest } from "./types/loan/loan/query";
import { MsgRepayLoanResponse } from "./types/loan/loan/tx";
import { QueryAllLoanResponse } from "./types/loan/loan/query";
import { Loan } from "./types/loan/loan/loan";
import { MsgRequestLoanResponse } from "./types/loan/loan/tx";
import { MsgLiquidateLoanResponse } from "./types/loan/loan/tx";
import { GenesisState } from "./types/loan/loan/genesis";
import { QueryParamsResponse } from "./types/loan/loan/query";
import { QueryGetLoanResponse } from "./types/loan/loan/query";
import { MsgUpdateParams } from "./types/loan/loan/tx";
import { MsgUpdateParamsResponse } from "./types/loan/loan/tx";

const msgTypes: Array<[string, GeneratedType]>  = [
    ["/loan.loan.MsgApproveLoan", MsgApproveLoan],
    ["/loan.loan.MsgApproveLoanResponse", MsgApproveLoanResponse],
    ["/loan.loan.MsgRequestLoan", MsgRequestLoan],
    ["/loan.loan.MsgLiquidateLoan", MsgLiquidateLoan],
    ["/loan.loan.Params", Params],
    ["/loan.loan.QueryGetLoanRequest", QueryGetLoanRequest],
    ["/loan.loan.QueryAllLoanRequest", QueryAllLoanRequest],
    ["/loan.loan.MsgRepayLoan", MsgRepayLoan],
    ["/loan.loan.QueryParamsRequest", QueryParamsRequest],
    ["/loan.loan.MsgRepayLoanResponse", MsgRepayLoanResponse],
    ["/loan.loan.QueryAllLoanResponse", QueryAllLoanResponse],
    ["/loan.loan.Loan", Loan],
    ["/loan.loan.MsgRequestLoanResponse", MsgRequestLoanResponse],
    ["/loan.loan.MsgLiquidateLoanResponse", MsgLiquidateLoanResponse],
    ["/loan.loan.GenesisState", GenesisState],
    ["/loan.loan.QueryParamsResponse", QueryParamsResponse],
    ["/loan.loan.QueryGetLoanResponse", QueryGetLoanResponse],
    ["/loan.loan.MsgUpdateParams", MsgUpdateParams],
    ["/loan.loan.MsgUpdateParamsResponse", MsgUpdateParamsResponse],
    
];

export { msgTypes }