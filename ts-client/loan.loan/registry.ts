import { GeneratedType } from "@cosmjs/proto-signing";
import { GenesisState } from "./types/loan/loan/genesis";
import { MsgUpdateParams } from "./types/loan/loan/tx";
import { Params } from "./types/loan/loan/params";
import { QueryParamsRequest } from "./types/loan/loan/query";
import { QueryParamsResponse } from "./types/loan/loan/query";
import { MsgUpdateParamsResponse } from "./types/loan/loan/tx";

const msgTypes: Array<[string, GeneratedType]>  = [
    ["/loan.loan.GenesisState", GenesisState],
    ["/loan.loan.MsgUpdateParams", MsgUpdateParams],
    ["/loan.loan.Params", Params],
    ["/loan.loan.QueryParamsRequest", QueryParamsRequest],
    ["/loan.loan.QueryParamsResponse", QueryParamsResponse],
    ["/loan.loan.MsgUpdateParamsResponse", MsgUpdateParamsResponse],
    
];

export { msgTypes }