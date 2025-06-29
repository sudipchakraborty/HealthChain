import { GeneratedType } from "@cosmjs/proto-signing";
import { MsgCreateRecord } from "./types/healthchain/healthchain/tx";
import { MsgUpdateRecord } from "./types/healthchain/healthchain/tx";
import { MsgDeleteRecord } from "./types/healthchain/healthchain/tx";

const msgTypes: Array<[string, GeneratedType]>  = [
    ["/healthchain.healthchain.MsgCreateRecord", MsgCreateRecord],
    ["/healthchain.healthchain.MsgUpdateRecord", MsgUpdateRecord],
    ["/healthchain.healthchain.MsgDeleteRecord", MsgDeleteRecord],
    
];

export { msgTypes }