import React from 'react'
interface props{
    title:string;
    amount: string;
    icon:string;
    color:string;
    date:string;
}
const InvoiceCard = ({
    title,amount,icon,color,date
}:props) => {
  return (
    <div className='flex justify-between items-center mb-8 overflow-x-visible'>
        <div className='flex items-center gap-3'>
        <div className="border border-solid rounded-2xl  w-[60px] h-[60px] flex justify-center items-center" style={{borderColor: color , backgroundColor:color}}>
<img src={icon}/>
</div>
<div className='flex flex-col'>
    <p className='font-inter font-medium text-base text-[#B1B1B1]'>{title}</p>
    <p className='font-inter font-normal text-[15px] text-[#718EBF]'>{date}</p>
</div>
        </div>
        <p className='font-inter font-normal text-base text-[#718EBF]'>{amount}</p>
      
    </div>
  )
}

export default InvoiceCard
