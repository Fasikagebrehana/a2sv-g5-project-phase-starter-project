
import React from 'react'
import { Color } from 'chart.js'
import { colors  , logo} from '@/constants'
import DesktopCreditCart from '@/components/DesktopCreditCard'
import RecentTransaction from '@/components/Recent Transaction'
import ExpensesChart from '@/components/ExpensesCart'
import { icons } from 'lucide-react'
import { text } from 'stream/consumers'
import BarChart from '@/components/BarChart'
import  PieChart  from '@/components/PieChart'  

const page = () => {
  return (
    <div className={`${colors.graybg} p-6`}>
      <div className='flex flex-col justify-between w-screen  md:flex-row max-w-screen-2xl gap-10 '>
        <div className=' md:w-3/5 py-4 '>
          <div className={`${colors.navbartext} flex justify-between `}>
            <h1 className='font-bold text-2xl'>My Cards</h1>
            <p className='my-2'> See All</p>
          </div>
          <div className='flex justify-between gap-3 overflow-x-auto md: max-w-4xl'>
            <div className='flex py-3 '> <DesktopCreditCart bgColor={colors.blue} textColor={colors.textwhite}/></div>
            <div className='flex py-3'> <DesktopCreditCart bgColor={colors.white} textColor={colors.textblack}/></div>
          </div>
        </div>
        <div className='w-2/5'>
          <div className={`${colors.navbartext} flex justify-between  py-4`}>
            <h1 className='font-bold text-2xl'>Recent Transaction</h1>
          </div>
          <div className=''>
            <RecentTransaction icon={logo.RT1} color={colors.lightorange} colortext={colors.textred} name="deposit from my" date = "23 january 2023" amount='-850'/>
            <RecentTransaction icon={logo.RT2} color={colors.lightblue} colortext={colors.textgreen} name="deposit from my" date = "23 january 2023" amount='+2500'/>
            <RecentTransaction icon={logo.RT3} color={colors.lightgreen} colortext={colors.textgreen} name="deposit from my" date = "23 january 2023" amount='+5400'/>
          </div>
        </div>
      </div>
      <div className='flex flex-col  justify-between gap-10 md:flex-row max-w-screen-2xl' >
        <div className=' md:w-3/5 py-5'>
          <div className={`${colors.navbartext} flex justify-between py-4`}>
            <h1 className='font-bold text-2xl'>Weekly Activity</h1>
          </div>
          <div><BarChart/></div>
        </div>
        <div className='  md:w-2/5 py-5'>
          <div className={`${colors.navbartext} flex justify-between py-4`}>
            <h1 className='font-bold text-2xl'>Expense Statstics</h1>
          </div>
          <div><PieChart/></div>
        </div>
            </div>
      </div>
    
  )
}

export default page