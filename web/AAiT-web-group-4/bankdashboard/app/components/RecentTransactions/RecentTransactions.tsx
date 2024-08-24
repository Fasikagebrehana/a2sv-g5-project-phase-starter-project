import React from "react";

const RecentTransactions = () => {
  return (
    <div className="flex flex-col justify-evenly tablet:w-[360px] max-tablet:w-[325px] h-3/4 bg-white rounded-3xl">
      <div className="flex px-4 justify-evenly">
        <svg
          width="40"
          height="40"
          viewBox="0 0 40 40"
          fill="none"
          xmlns="http://www.w3.org/2000/svg"
        >
          <circle cx="20" cy="20" r="20" fill="#E7EDFF" />
          <path
            d="M26.425 15.9876C26.4151 14.8945 25.9744 13.8495 25.1985 13.0795C24.4227 12.3095 23.3744 11.8767 22.2813 11.8751H16.25C16.1002 11.8731 15.9547 11.9251 15.84 12.0214C15.7253 12.1178 15.649 12.2522 15.625 12.4001L13.4875 25.8063C13.4742 25.8951 13.4801 25.9857 13.5049 26.072C13.5297 26.1582 13.5728 26.2382 13.6313 26.3063C13.6893 26.376 13.7619 26.4323 13.844 26.4712C13.926 26.51 14.0155 26.5305 14.1063 26.5313H16.6875L16.55 27.4001C16.5354 27.4901 16.5407 27.5822 16.5655 27.67C16.5903 27.7578 16.634 27.8391 16.6936 27.9082C16.7531 27.9773 16.8271 28.0325 16.9103 28.07C16.9934 28.1074 17.0838 28.1262 17.175 28.1251H20.0813C20.2302 28.1272 20.3749 28.0761 20.4895 27.981C20.6041 27.8859 20.681 27.753 20.7063 27.6063L21.3313 23.7876H23.3C24.5322 23.7826 25.7124 23.29 26.5826 22.4175C27.4527 21.545 27.9421 20.3635 27.9438 19.1313V18.9563C27.9431 18.3767 27.8052 17.8054 27.5412 17.2894C27.2772 16.7734 26.8947 16.3272 26.425 15.9876ZM16.7813 13.1251H22.2813C22.937 13.1271 23.5727 13.3512 24.0846 13.761C24.5965 14.1708 24.9544 14.7419 25.1 15.3813C24.8296 15.3179 24.5527 15.2864 24.275 15.2876H19.0625C18.9127 15.2856 18.7672 15.3376 18.6525 15.4339C18.5378 15.5303 18.4615 15.6647 18.4375 15.8126L18.0688 18.1251C18.0422 18.2908 18.0826 18.4603 18.1811 18.5963C18.2796 18.7322 18.428 18.8235 18.5938 18.8501C18.7595 18.8766 18.929 18.8362 19.065 18.7377C19.201 18.6392 19.2922 18.4908 19.3188 18.3251L19.6063 16.5251H24.2875C24.5871 16.5266 24.8838 16.5838 25.1625 16.6938C25.0437 17.6722 24.5716 18.5735 23.8351 19.2284C23.0985 19.8833 22.1481 20.2466 21.1625 20.2501H18.2625C18.1136 20.2479 17.9688 20.299 17.8543 20.3941C17.7397 20.4892 17.6628 20.6221 17.6375 20.7688L16.875 25.2813H14.8375L16.7813 13.1251ZM26.6938 19.1313C26.6921 20.032 26.3344 20.8956 25.6987 21.5337C25.063 22.1717 24.2007 22.5326 23.3 22.5376H20.8C20.6511 22.5354 20.5063 22.5865 20.3918 22.6816C20.2772 22.7767 20.2003 22.9096 20.175 23.0563L19.55 26.8751H17.9L18.0375 26.0063L18.7875 21.5126H21.15C22.3168 21.5091 23.4497 21.1198 24.3722 20.4054C25.2946 19.6909 25.9549 18.6914 26.25 17.5626C26.5401 17.9693 26.6953 18.4567 26.6938 18.9563V19.1313Z"
            fill="#396AFF"
          />
        </svg>
        <div className="flex flex-col">
          <p>Deposit from paypal</p>
          <p>date</p>
        </div>
        <p className="text-green-500">+$2500</p>
        
      </div>

      <div className="flex px-4 justify-evenly">
        <svg
          width="40"
          height="40"
          viewBox="0 0 40 40"
          fill="none"
          xmlns="http://www.w3.org/2000/svg"
        >
          <circle cx="20" cy="20" r="20" fill="#E7EDFF" />
          <path
            d="M26.425 15.9876C26.4151 14.8945 25.9744 13.8495 25.1985 13.0795C24.4227 12.3095 23.3744 11.8767 22.2813 11.8751H16.25C16.1002 11.8731 15.9547 11.9251 15.84 12.0214C15.7253 12.1178 15.649 12.2522 15.625 12.4001L13.4875 25.8063C13.4742 25.8951 13.4801 25.9857 13.5049 26.072C13.5297 26.1582 13.5728 26.2382 13.6313 26.3063C13.6893 26.376 13.7619 26.4323 13.844 26.4712C13.926 26.51 14.0155 26.5305 14.1063 26.5313H16.6875L16.55 27.4001C16.5354 27.4901 16.5407 27.5822 16.5655 27.67C16.5903 27.7578 16.634 27.8391 16.6936 27.9082C16.7531 27.9773 16.8271 28.0325 16.9103 28.07C16.9934 28.1074 17.0838 28.1262 17.175 28.1251H20.0813C20.2302 28.1272 20.3749 28.0761 20.4895 27.981C20.6041 27.8859 20.681 27.753 20.7063 27.6063L21.3313 23.7876H23.3C24.5322 23.7826 25.7124 23.29 26.5826 22.4175C27.4527 21.545 27.9421 20.3635 27.9438 19.1313V18.9563C27.9431 18.3767 27.8052 17.8054 27.5412 17.2894C27.2772 16.7734 26.8947 16.3272 26.425 15.9876ZM16.7813 13.1251H22.2813C22.937 13.1271 23.5727 13.3512 24.0846 13.761C24.5965 14.1708 24.9544 14.7419 25.1 15.3813C24.8296 15.3179 24.5527 15.2864 24.275 15.2876H19.0625C18.9127 15.2856 18.7672 15.3376 18.6525 15.4339C18.5378 15.5303 18.4615 15.6647 18.4375 15.8126L18.0688 18.1251C18.0422 18.2908 18.0826 18.4603 18.1811 18.5963C18.2796 18.7322 18.428 18.8235 18.5938 18.8501C18.7595 18.8766 18.929 18.8362 19.065 18.7377C19.201 18.6392 19.2922 18.4908 19.3188 18.3251L19.6063 16.5251H24.2875C24.5871 16.5266 24.8838 16.5838 25.1625 16.6938C25.0437 17.6722 24.5716 18.5735 23.8351 19.2284C23.0985 19.8833 22.1481 20.2466 21.1625 20.2501H18.2625C18.1136 20.2479 17.9688 20.299 17.8543 20.3941C17.7397 20.4892 17.6628 20.6221 17.6375 20.7688L16.875 25.2813H14.8375L16.7813 13.1251ZM26.6938 19.1313C26.6921 20.032 26.3344 20.8956 25.6987 21.5337C25.063 22.1717 24.2007 22.5326 23.3 22.5376H20.8C20.6511 22.5354 20.5063 22.5865 20.3918 22.6816C20.2772 22.7767 20.2003 22.9096 20.175 23.0563L19.55 26.8751H17.9L18.0375 26.0063L18.7875 21.5126H21.15C22.3168 21.5091 23.4497 21.1198 24.3722 20.4054C25.2946 19.6909 25.9549 18.6914 26.25 17.5626C26.5401 17.9693 26.6953 18.4567 26.6938 18.9563V19.1313Z"
            fill="#396AFF"
          />
        </svg>
        <div className="flex flex-col">
          <p>Deposit from paypal</p>
          <p>date</p>
        </div>
        <p className="text-green-500">+$2500</p>
        </div>

        <div className="flex px-4 justify-evenly">
        <svg
          width="40"
          height="40"
          viewBox="0 0 40 40"
          fill="none"
          xmlns="http://www.w3.org/2000/svg"
        >
          <circle cx="20" cy="20" r="20" fill="#E7EDFF" />
          <path
            d="M26.425 15.9876C26.4151 14.8945 25.9744 13.8495 25.1985 13.0795C24.4227 12.3095 23.3744 11.8767 22.2813 11.8751H16.25C16.1002 11.8731 15.9547 11.9251 15.84 12.0214C15.7253 12.1178 15.649 12.2522 15.625 12.4001L13.4875 25.8063C13.4742 25.8951 13.4801 25.9857 13.5049 26.072C13.5297 26.1582 13.5728 26.2382 13.6313 26.3063C13.6893 26.376 13.7619 26.4323 13.844 26.4712C13.926 26.51 14.0155 26.5305 14.1063 26.5313H16.6875L16.55 27.4001C16.5354 27.4901 16.5407 27.5822 16.5655 27.67C16.5903 27.7578 16.634 27.8391 16.6936 27.9082C16.7531 27.9773 16.8271 28.0325 16.9103 28.07C16.9934 28.1074 17.0838 28.1262 17.175 28.1251H20.0813C20.2302 28.1272 20.3749 28.0761 20.4895 27.981C20.6041 27.8859 20.681 27.753 20.7063 27.6063L21.3313 23.7876H23.3C24.5322 23.7826 25.7124 23.29 26.5826 22.4175C27.4527 21.545 27.9421 20.3635 27.9438 19.1313V18.9563C27.9431 18.3767 27.8052 17.8054 27.5412 17.2894C27.2772 16.7734 26.8947 16.3272 26.425 15.9876ZM16.7813 13.1251H22.2813C22.937 13.1271 23.5727 13.3512 24.0846 13.761C24.5965 14.1708 24.9544 14.7419 25.1 15.3813C24.8296 15.3179 24.5527 15.2864 24.275 15.2876H19.0625C18.9127 15.2856 18.7672 15.3376 18.6525 15.4339C18.5378 15.5303 18.4615 15.6647 18.4375 15.8126L18.0688 18.1251C18.0422 18.2908 18.0826 18.4603 18.1811 18.5963C18.2796 18.7322 18.428 18.8235 18.5938 18.8501C18.7595 18.8766 18.929 18.8362 19.065 18.7377C19.201 18.6392 19.2922 18.4908 19.3188 18.3251L19.6063 16.5251H24.2875C24.5871 16.5266 24.8838 16.5838 25.1625 16.6938C25.0437 17.6722 24.5716 18.5735 23.8351 19.2284C23.0985 19.8833 22.1481 20.2466 21.1625 20.2501H18.2625C18.1136 20.2479 17.9688 20.299 17.8543 20.3941C17.7397 20.4892 17.6628 20.6221 17.6375 20.7688L16.875 25.2813H14.8375L16.7813 13.1251ZM26.6938 19.1313C26.6921 20.032 26.3344 20.8956 25.6987 21.5337C25.063 22.1717 24.2007 22.5326 23.3 22.5376H20.8C20.6511 22.5354 20.5063 22.5865 20.3918 22.6816C20.2772 22.7767 20.2003 22.9096 20.175 23.0563L19.55 26.8751H17.9L18.0375 26.0063L18.7875 21.5126H21.15C22.3168 21.5091 23.4497 21.1198 24.3722 20.4054C25.2946 19.6909 25.9549 18.6914 26.25 17.5626C26.5401 17.9693 26.6953 18.4567 26.6938 18.9563V19.1313Z"
            fill="#396AFF"
          />
        </svg>
        <div className="flex flex-col">
          <p>Deposit from paypal</p>
          <p>date</p>
        </div>
        <p className="text-green-500">+$2500</p>
        </div>

      
    </div>
  );
};

export default RecentTransactions;
