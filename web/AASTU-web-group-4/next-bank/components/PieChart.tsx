"use client"
import { Pie, PieChart } from "recharts"
import {
  Card,
  CardContent,
  CardHeader,
  CardTitle,
} from "@/components/ui/card"
import {
  ChartConfig,
  ChartContainer,
  ChartTooltip,
  ChartTooltipContent,
} from "@/components/ui/chart"
import { useEffect, useState } from "react"

const chartData = [
  { category: "Entertainment", amount: 275, fill: "hsl(210, 100%, 40%)" },  // Dim Blue
  { category: "Shopping", amount: 200, fill: "hsl(120, 100%, 30%)" },       // Dim Green
  { category: "Groceries", amount: 187, fill: "hsl(30, 100%, 40%)" },       // Dim Orange
  { category: "Bills", amount: 173, fill: "hsl(0, 80%, 40%)" },             // Dim Red
  { category: "Other", amount: 90, fill: "hsl(60, 100%, 40%)" },            // Dim Yellow
]



export default function Component() {
  const [pierad , setpierad] = useState(130)
  useEffect(()=>{

    const fun= () =>{
      if(window.innerWidth < 1024){
        setpierad(20)
        alert('less than 1024')
    
      }
      else{
        setpierad(150)
      }
    fun()
    window.addEventListener('resize' , fun)
    return()=>{
      window.removeEventListener('resize' , fun)
    }
    }
    },[window.innerWidth])
  return (
    <Card className=" md:py-10  ">
      <CardHeader className="items-center pb-0">
        
      </CardHeader>
      <CardContent className="flex-1 pb-0">
        <ChartContainer
          className="mx-auto  max-h-[90%] "
          config={{}}
        >
          <PieChart>
            <ChartTooltip
              content={<ChartTooltipContent nameKey="category" />}
            />
            <Pie
              data={chartData}
              dataKey="amount"
              nameKey="category"
              cx="50%"
              cy="50%"
              outerRadius={pierad}
              // className="md:"
              labelLine={false}
              label={({ cx, cy, midAngle, innerRadius, outerRadius, percent, index }) => {
                const RADIAN = Math.PI / 180;
                const radius = innerRadius + (outerRadius - innerRadius) * 0.5;
                const x = cx + radius * Math.cos(-midAngle * RADIAN);
                const y = cy + radius * Math.sin(-midAngle * RADIAN);
                return (
                  <text
                    x={x}
                    y={y}
                    fill="white"
                    textAnchor="middle"
                    dominantBaseline="central"
                    fontSize={12}
                  >
                    {`${(percent * 100).toFixed(0)}%`}
                  </text>
                );
              }}
            />
          </PieChart>
        </ChartContainer>
      </CardContent>
    </Card>
  )
}
