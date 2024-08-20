import React from 'react';
import Table from './Table'; 

interface Transaction {
  description: string;
  transactionId: string;
  type: string;
  card?: string; // Mark this as optional
  date: string;
  amount: string;
  receipt?: string; // Mark this as optional
}


interface TransactionsListProps {
  transactions: Transaction[];
}

const TransactionsList: React.FC<TransactionsListProps> = ({ transactions }) => {
  const columns: Array<{ Header: string; accessor: keyof Transaction; Cell?: (row: any) => JSX.Element }> = [
    {
      Header: '',
      accessor: 'amount',
      Cell: ({ value }: { value: string | undefined }) => {
        const amountValue = value ? parseFloat(value.replace(/[^0-9.-]/g, '')) : 0;
        const isPositive = amountValue > 0;
        return (
          <div className="flex items-center">
            <div className="border border-solid border-[#718EBF] rounded-full flex justify-center items-center h-6 w-6 -mr-4">
              <img
                src={isPositive ? '/downArrow.svg' : '/upArrow.svg'}
                alt={isPositive ? 'down arrow' : 'up arrow'}
                className="h-3 w-3"
              />
            </div>
          </div>
        );
      },
    },
    {
      Header: 'Description',
      accessor: 'description',
      Cell: ({ value }: { value: string }) => (
        <div className="flex items-center">
          <span className="ml-2">{value}</span>
        </div>
      ),
    },
    {
      Header: 'Transaction ID',
      accessor: 'transactionId',
    },
    {
      Header: 'Type',
      accessor: 'type',
    },
    {
      Header: 'Card',
      accessor: 'card',
    },
    {
      Header: 'Date',
      accessor: 'date',
    },
    {
      Header: 'Amount',
      accessor: 'amount',
      Cell: ({ value }: { value: string | undefined }) => {
        const amountValue = value ? parseFloat(value.replace(/[^0-9.-]/g, '')) : 0;
        const isPositive = amountValue > 0;
        return (
          <span className={isPositive ? 'text-green-500' : 'text-red-500'}>
            {value || 'N/A'}
          </span>
        );
      },
    },
    {
      Header: 'Receipt',
      accessor: 'receipt',
      Cell: () => (
        <button className="text-[#123288] border border-[#123288] rounded-full px-2 py-1 hover:text-blue-600 hover:border-blue-600 transition-colors duration-300">
          Download
        </button>
      ),
    },
  ];

  return (
    <div className="p-4">
      {/* Desktop View */}
      <div className="hidden md:block">
        <Table columns={columns} data={transactions} />
      </div>
      
      {/* Mobile View */}
      <div className="block md:hidden">
        {transactions.map((transaction, index) => (
          <div key={index} className="border-b border-gray-200 py-2">
            <div className="flex justify-between items-center">
              <div className="flex items-center">
                <div className="border border-solid border-[#718EBF] rounded-full flex justify-center items-center h-10 w-10">
                  <img
                    src={
                      parseFloat(transaction.amount.replace(/[^0-9.-]/g, '')) > 0
                        ? '/downArrow.svg'
                        : '/upArrow.svg'
                    }
                    alt={
                      parseFloat(transaction.amount.replace(/[^0-9.-]/g, '')) > 0
                        ? 'down arrow'
                        : 'up arrow'
                    }
                    className="h-5 w-5"
                  />
                </div>
                <div className="ml-2">
                  <span>{transaction.description}</span>
                  <div className="text-sm text-gray-500">{transaction.date}</div>
                </div>
              </div>
              <span
                className={
                  parseFloat(transaction.amount.replace(/[^0-9.-]/g, '')) > 0
                    ? 'text-green-500'
                    : 'text-red-500'
                }
              >
                {transaction.amount}
              </span>
            </div>
          </div>
        ))}
      </div>

      <div className="flex items-center mt-4 space-x-2">
        <div className="ml-auto flex items-center space-x-1">
          <img src="/back.svg" alt="back" className="h-4 w-4 text-blue-600" />
          <button className="text-blue-600 rounded-full px-2 py-1">
            Previous
          </button>
        </div>
        <div className="flex space-x-1">
          <button className="px-2 py-1 bg-[#1814F3] text-white rounded">1</button>
          <button className="px-3 py-1 text-[#1814F3] rounded">2</button>
          <button className="px-3 py-1 text-[#1814F3] rounded">3</button>
          <button className="px-3 py-1 text-[#1814F3] rounded">4</button>
        </div>
        <div className="flex items-center space-x-1">
          <button className="text-blue-600 rounded-full px-1 py-1">
            Next
          </button>
          <img src="/forward.svg" alt="forward" className="h-4 w-4 text-blue-600" />
        </div>
      </div>
    </div>
  );
};

export default TransactionsList;
