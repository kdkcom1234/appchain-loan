import { GeneratedType } from "@cosmjs/proto-signing";
import { MsgRepayLoan } from "./types/loan/loan/tx";
import { MsgRequestLoan } from "./types/loan/loan/tx";
import { MsgUpdateParamsResponse } from "./types/loan/loan/tx";
import { QueryParamsRequest } from "./types/loan/loan/query";
import { QueryGetLoanResponse } from "./types/loan/loan/query";
import { QueryAllLoanRequest } from "./types/loan/loan/query";
import { MsgApproveLoan } from "./types/loan/loan/tx";
import { MsgUpdateParams } from "./types/loan/loan/tx";
import { MsgRequestLoanResponse } from "./types/loan/loan/tx";
import { QueryAllLoanResponse } from "./types/loan/loan/query";
import { MsgRepayLoanResponse } from "./types/loan/loan/tx";
import { GenesisState } from "./types/loan/loan/genesis";
import { MsgApproveLoanResponse } from "./types/loan/loan/tx";
import { Params } from "./types/loan/loan/params";
import { QueryParamsResponse } from "./types/loan/loan/query";
import { QueryGetLoanRequest } from "./types/loan/loan/query";
import { Loan } from "./types/loan/loan/loan";

const msgTypes: Array<[string, GeneratedType]>  = [
    ["/loan.loan.MsgRepayLoan", MsgRepayLoan],
    ["/loan.loan.MsgRequestLoan", MsgRequestLoan],
    ["/loan.loan.MsgUpdateParamsResponse", MsgUpdateParamsResponse],
    ["/loan.loan.QueryParamsRequest", QueryParamsRequest],
    ["/loan.loan.QueryGetLoanResponse", QueryGetLoanResponse],
    ["/loan.loan.QueryAllLoanRequest", QueryAllLoanRequest],
    ["/loan.loan.MsgApproveLoan", MsgApproveLoan],
    ["/loan.loan.MsgUpdateParams", MsgUpdateParams],
    ["/loan.loan.MsgRequestLoanResponse", MsgRequestLoanResponse],
    ["/loan.loan.QueryAllLoanResponse", QueryAllLoanResponse],
    ["/loan.loan.MsgRepayLoanResponse", MsgRepayLoanResponse],
    ["/loan.loan.GenesisState", GenesisState],
    ["/loan.loan.MsgApproveLoanResponse", MsgApproveLoanResponse],
    ["/loan.loan.Params", Params],
    ["/loan.loan.QueryParamsResponse", QueryParamsResponse],
    ["/loan.loan.QueryGetLoanRequest", QueryGetLoanRequest],
    ["/loan.loan.Loan", Loan],
    
];

export { msgTypes }