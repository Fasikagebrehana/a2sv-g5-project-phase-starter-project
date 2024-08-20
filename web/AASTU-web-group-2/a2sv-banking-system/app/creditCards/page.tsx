import React from "react";
import PieChart from "./PieChart";
import CreditCard from "./CreditCard";
import CardSettingList from "./CardSettingList";
import AddCardForm from "./AddCardForm";
import MainCreditCard from "./MainCreditCard";
import Card from "../components/Page2/Card";

const HeadingTitle = ({ title }: { title: string }) => {
  return (
    <h1 className="text-[#343C6A] font-semibold lg:text-xl md:text-lg">
      {title}
    </h1>
  );
};

const CreditCards = () => {
  return (
    <div className="bg-[#f5f7fb] w-full p-5 gap-5 flex flex-col">
      <div className="flex-col gap-5">
        <HeadingTitle title="My Cards" />

        <div className="flex overflow-scroll justify-between">
          <Card
            balance="$5,756"
            cardHolder="Eddy Cusuma"
            validThru="12/22"
            cardNumber="3778 **** **** 1234"
            filterClass=""
            bgColor="from-[#0A06F4] to-[#0A06F4]"
            textColor="text-white"
            iconBgColor="bg-opacity-10"
            showIcon={true}
          />
          <Card
            balance="$5,756"
            cardHolder="Eddy Cusuma"
            validThru="12/22"
            cardNumber="3778 **** **** 1234"
            filterClass=""
            bgColor="from-[#4C49ED] to-[#4C49ED]"
            textColor="text-white"
            iconBgColor="bg-opacity-10"
            showIcon={true}
          />
          <Card
            balance="$5,756"
            cardHolder="Eddy Cusuma"
            validThru="12/22"
            cardNumber="3778 **** **** 1234"
            filterClass=""
            bgColor="from-[#FFF] to-[#FFF]"
            textColor="text-black"
            iconBgColor="bg-opacity-10"
            showIcon={true}
          />
        </div>
      </div>
      <div className="flex flex-col gap-6 md:flex-row">
        <div className="flex flex-col gap-5 basis-5/12 ">
          <HeadingTitle title="Card Expense Statistics" />
          <PieChart />
        </div>
        <div className="flex flex-col gap-3 md:justify-between w-full h-full">
          <HeadingTitle title="Card List" />
          <CreditCard
            icon={<img src="card1.svg" />}
            linkUrl=""
            data={[
              ["Card Type", "Secondary"],
              ["Card Type", "Secondary"],
              ["Card Type", "Secondary"],
              ["Card Type", "Secondary"],
            ]}
          />

          <CreditCard
            icon={<img src="card1.svg" />}
            linkUrl=""
            data={[
              ["Card Type", "Secondary"],
              ["Card Type", "Secondary"],
              ["Card Type", "Secondary"],
              ["Card Type", "Secondary"],
            ]}
          />
          <CreditCard
            icon={<img src="card1.svg" />}
            linkUrl=""
            data={[
              ["Card Type", "Secondary"],
              ["Card Type", "Secondary"],
              ["Card Type", "Secondary"],
              ["Card Type", "Secondary"],
            ]}
          />
        </div>
      </div>

      <div className="flex flex-col gap-6 md:flex-row">
        <div className="flex flex-col gap-5">
          <HeadingTitle title="Add New Card" />
          <AddCardForm />
        </div>
        <div className="flex flex-col gap-5 min-w-64 h-full">
          <HeadingTitle title="Card Setting" />
          <CardSettingList />
        </div>
      </div>
    </div>
  );
};

export default CreditCards;
