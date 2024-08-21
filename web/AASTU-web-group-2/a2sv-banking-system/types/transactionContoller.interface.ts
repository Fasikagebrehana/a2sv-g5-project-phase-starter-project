export interface TransactionResponse {
  success: boolean;
  message: string;
  data: TransactionData;
}

export interface TransactionData {
  transactionId: string;
  type: string;
  senderUserName: string;
  description: string;
  date: string;
  amount: number;
  receiverUserName: string;
}

export interface GetTransactionsResponse {
  transactions: TransactionResponse[];
}

export interface PostDepositTransactionRequest {
  description: string;
  amount: number;
}

export interface BalanceHistoryData {
  time: string;
  value: number;
}

export interface BalanceHistoryResponse {
  success: boolean;
  message: string;
  data: BalanceHistoryData[];
}

export interface QuickTransferData {
  id: string;
  name: string;
  username: string;
  city: string;
  country: string;
  profilePicture: string;
}

export interface PostTransactionRequest {
  type: string;
  description: string;
  amount: number;
  receiverUserName: string;
}

export interface QuickTransferResponse {
  success: boolean;
  message: string;
  data: QuickTransferData[];
}

export interface PaginatedTransactionsResponse {
  success: boolean;
  message: string;
  data: TransactionData[];
}

export interface PostTransactionResponse extends TransactionResponse {}
export interface GetTransactionById extends TransactionResponse {}
