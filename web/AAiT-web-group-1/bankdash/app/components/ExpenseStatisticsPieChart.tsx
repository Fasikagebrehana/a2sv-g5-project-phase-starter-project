"use client";

import React, { useRef, useEffect } from 'react';
import { Chart, CategoryScale, LinearScale, PieController, ArcElement, Tooltip, Legend } from 'chart.js';
import ChartDataLabels from 'chartjs-plugin-datalabels';

Chart.register(CategoryScale, LinearScale, PieController, ArcElement, Tooltip, Legend, ChartDataLabels);

// const sectors = ['Service', 'Others','Shopping', 'Transfer']
// const bgColors = ['#FC7900', '#1814F3', '#FA00FF', '#343C6A']

interface ExpenseStatisticsPieChartProps {
    sectors: string[],
    bgColors: string[],
}

const ExpenseStatisticsPieChart = ({sectors, bgColors} : ExpenseStatisticsPieChartProps) => {
    const chartRef = useRef<HTMLCanvasElement>(null);
    const chartInstanceRef = useRef<Chart<"pie", number[], string> | null>(null);

    useEffect(() => {
        if (chartRef.current && !chartInstanceRef.current) {
            const context = chartRef.current.getContext('2d');

            if (context) {
                chartInstanceRef.current = new Chart(context, {
                    type: 'pie',
                    data: {
                        labels: sectors, // Categories
                        datasets: [{
                            data: [15, 35, 20, 30], // Percentages for each category
                            backgroundColor: bgColors,
                            borderColor: '#ffffff',
                            borderWidth: 2,
                            offset: [50, 30, 0, 40]
                        }]
                    },
                    options: {
                        layout: {
                            padding: {
                                left: 20,
                                right: 20,
                                top: 10,
                                bottom: 10
                            }
                        },
                        plugins: {
                            legend: {
                                display: false,
                            },
                            tooltip: {
                                enabled: false,
                            },
                            datalabels: {
                                color: '#fff', // Color of the text
                                font: {
                                    size: 14, // Font size
                                    weight: 'bold' // Font weight
                                },
                                formatter: (value, context) => {
                                    const label = context.chart.data.labels![context.dataIndex];
                                    return `${value}% \n ${label} `;
                                },
                                anchor: 'end', 
                                align: 'start', 
                                offset: 10, 
                            }
                        },
                        maintainAspectRatio: false,
                    }
                });
            }
        }

        // Cleanup: Destroy the chart instance when the component unmounts
        return () => {
            if (chartInstanceRef.current) {
                chartInstanceRef.current.destroy();
                chartInstanceRef.current = null;
            }
        };
    }, []);

    return (
        <div className='relative ml-4 pb-4 h-80 text-sm'>
            <canvas ref={chartRef} />
        </div>
    );
};

export default ExpenseStatisticsPieChart;

