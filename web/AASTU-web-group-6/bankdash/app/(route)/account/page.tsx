import Card from '../../components/Accounts/account'
import Last_trans from "../../components/Accounts/Last_trans";
import { DebitCreditOver } from "../../components/Accounts/DebitCreditOver";
import InvoiceCard from "../../components/Accounts/InvoiceCard";
import VisaCard from '@/app/components/Card/VisaCard';
export default function Home() {
  return (
    <>
      <div className="flex flex-col lg:flex-row gap-3 lg:gap-7">
      <div className='flex w-full lg:w-[45%] gap-3 lg:gap-12'>
        <Card
          title="My Balance"
          amount="$12,750"
          color="#FFF5D9"
          icon="/assets/money-tag 1.svg"
          width='w-[45%]'
        />
        <Card
          title="Income"
          amount="$5,600"
          color="#E7EDFF"
          icon="/assets/expense.svg"
          width='w-[45%]'
        />
        </div>
        <div className='flex  w-full lg:w-[45%] gap-3 lg:gap-12'>
        <Card
          title="Expense"
          amount="$3,460"
          color="#FFE0EB"
          icon="/assets/income.svg"
          width='w-[45%]'
        />
        <Card
          title="Total Saving"
          amount="$7,920"
          color="#DCFAF8"
          icon="/assets/saving.svg"
          width='w-[45%]'
        />
        </div>
      </div>

      <div className="flex  flex-col lg:flex-row my-5 justify-between">
        <div className="lg:w-[65%]">
          <p className="font-inter font-semibold text-[22px] text-[#333B69] mb-5">
            Last Transaction
          </p>
          <div className=" bg-white border rounded-3xl p-3 shadow-lg border-gray-300">
            <Last_trans
              title="Spotify Subscription"
              date="25 Jan 2021"
              type="Shopping"
              account_no="1234 ****"
              status="Pending"
              amount="-$150"
              color="#DCFAF8"
              icon="/assets/renew.svg"
            />
            <Last_trans
              title="Mobile Service"
              date="25 Jan 2021"
              type="Service"
              account_no="1234 ****"
              status="Completed"
              amount="-$340"
              color="#E7EDFF"
              icon="/assets/settings.svg"
            />
            <Last_trans
              title="Emilly Wilson"
              date="25 Jan 2021"
              type="Transfer"
              account_no="1234 ****"
              status="Completed"
              amount="+$780"
              color="#FFE0EB"
              icon="/assets/userr.svg"
            />
          </div>
        </div>
        <div className='lg:w-[30%] lg:h-[250px]'>
          <div className='flex justify-between'>
          <p className="font-inter font-semibold text-[22px] text-[#333B69] mb-5">
           My Card
          </p>
          <p className="font-inter font-semibold text-[22px] text-[#333B69] mb-5">See All</p>
          </div>
          <VisaCard isBlack={false} isFade={true}/>
       
        </div>
        
      </div>
      <div className="flex flex-col lg:flex-row justify-between my-5">
        <div className="lg:w-[65%] lb:h-[364px]">
          <p className="font-inter font-semibold text-[22px] text-[#333B69] mb-5">
            Debit & Credit Overview
          </p>
          <DebitCreditOver />
        </div>
        <div className="lg:w-[30%]">
          <p className="font-inter font-semibold text-[22px] text-[#333B69] mb-5 ">
            Invoices Sent
          </p>
          <div className="border border-solid rounded-3xl p-9  bg-white shadow-lg border-gray-300">
            <InvoiceCard
              title="Apple Store"
              date="5h ago"
              amount="$450"
              icon="/assets/apple.svg"
              color="#DCFAF8"
            />
            <InvoiceCard
              title="Michael"
              date="2 days ago"
              amount="$160"
              icon="/assets/userr.svg"
              color="#FFF5D9"
            />
            <InvoiceCard
              title="Playstation"
              date="5 days ago"
              amount="$1085"
              icon="/assets/Group.svg"
              color="#E7EDFF"
            />
            <InvoiceCard
              title="William"
              date="10 days ago"
              amount="$90"
              icon="/assets/userr.svg"
              color="#FFE0EB"
            />
          </div>
        </div>
      </div>
    </>
  );
}
