// Cynhyrchwyd y ffeil hon yn awtomatig. PEIDIWCH Â MODIWL
// This file is automatically generated. DO NOT EDIT
import {main} from '../models';
import {tables} from '../models';

export function AddFile(arg1:main.Form):Promise<string>;

export function CleanBin():Promise<string>;

export function EditFile(arg1:tables.Qc):Promise<string>;

export function GetFiles():Promise<Array<tables.Qc>>;

export function GetRecycleList():Promise<Array<tables.Qc>>;

export function Greet(arg1:string):Promise<string>;

export function RemoveFile(arg1:tables.Qc):Promise<string>;
