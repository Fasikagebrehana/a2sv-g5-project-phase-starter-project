"use client";
import { useRef, useEffect } from "react";
import { Chart, ChartData, ChartOptions } from "chart.js/auto";
import { useGetInvestmentHistoryQuery } from "@/lib/service/TransactionService";
import { useSession } from "next-auth/react";

interface CustomCanvasElement extends HTMLCanvasElement {
  chart?: Chart;
}

function YearlyTotalInvestment() {
  const chartRef = useRef<CustomCanvasElement | null>(null);
  const { data: session } = useSession();
  const accessToken = session?.user.accessToken!;
  const { data, isError, isLoading } = useGetInvestmentHistoryQuery(accessToken);

  useEffect(() => {
    if (chartRef.current && data && data.success) {
      // Destroy the chart instance if it already exists
      if (chartRef.current.chart) {
        chartRef.current.chart.destroy();
      }

      const context = chartRef.current.getContext("2d");

      if (context) {
        const chartItem = context.canvas;

        // Extract the labels and values from the API response
        const labels = data.data.yearlyTotalInvestment.map((item: { time: any; }) => item.time);
        const values = data.data.yearlyTotalInvestment.map((item: { value: any; }) => item.value);

        const chartData: ChartData<"line"> = {
          labels,
          datasets: [
            {
              label: "Balance",
              data: values,
              fill: true,
              backgroundColor: "rgba(252, 170, 11,0)", // Color for the filled area
              borderColor: "rgba(252, 170, 11,1)", // Color of the line
              borderWidth: 4,
              pointRadius: 5,
              pointBackgroundColor: "rgba(255, 255, 255, 1)",
              pointBorderColor: "rgba(252, 170, 11,1)",
            },
          ],
        };

        const chartOptions: ChartOptions<"line"> = {
          responsive: true,
          plugins: {
            legend: {
              display: false,
              labels: {
                color: "rgb(113, 142, 191)", // Change legend label color
              },
            },
          },
          scales: {
            x: {
              stacked: false,
              ticks: {
                align: "end",
                autoSkip: true,
                color: "rgb(113, 142, 191)", // Change x-axis tick label color
              },
              grid: {
                display: false,
                tickBorderDash: [1, 1],
              },
              border: {
                display: false,
              },
            },
            y: {
              beginAtZero: true,
              ticks: {
                callback: function (value) {
                  return `$${value}`;
                },
                color: "rgb(113, 142, 191)", // Change y-axis tick label color
              },
              border: {
                display: false,
              },
            },
          },
        };

        const newChart = new Chart(chartItem, {
          type: "line",
          data: chartData,
          options: chartOptions,
        });

        chartRef.current.chart = newChart;
      }
    }
  }, [data]);

  return (
    <div className="text-gray-500 border rounded-[22px] bg-white p-2  lg:w-[540px] lg:h-[282px] md:w-[359px] md:h-[226px] w-[325px] h-[225px]">
      <div>
        <div className="mt-8 expense-chart lg:mx-[20px] lg:w-[481px] lg:h-[228px] md:w-[321px] md:h-[190px] w-[283px] h-[157px]">
          <canvas ref={chartRef} />
        </div>
      </div>
    </div>
  );
}

export default YearlyTotalInvestment;
