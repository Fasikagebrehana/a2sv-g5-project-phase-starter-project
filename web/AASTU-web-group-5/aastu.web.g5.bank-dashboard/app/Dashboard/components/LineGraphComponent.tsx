"use client";

import React, { useEffect, useState } from "react";
import {
	Area,
	AreaChart,
	CartesianGrid,
	XAxis,
	YAxis,
	ResponsiveContainer,
} from "recharts";
import { Card, CardContent } from "@/components/ui/card";
import {
	ChartConfig,
	ChartContainer,
	ChartTooltip,
	ChartTooltipContent,
} from "@/components/ui/chart";
import { useSession } from "next-auth/react";
import axios from "axios";

const chartConfig = {
	desktop: {
		label: "Desktop",
		color: "hsl(var(--chart-1))",
	},
} satisfies ChartConfig;

const formatMonth = (dateString: string) => {
	const date = new Date(`${dateString}-01`);
	return date.toLocaleString("default", { month: "short" }); // Use 'short' for abbreviated month names
};

export function LineGraphComponent() {
	const [chartData, setChartData] = useState<{ time: string; value: number }[]>(
		[]
	);
	const [loading, setLoading] = useState<boolean>(true);
	const [error, setError] = useState<string | null>(null);
	interface ExtendedUser {
		name?: string;
		email?: string;
		image?: string;
		accessToken?: string;
	}
	const { data: session, status } = useSession();
	const user = session?.user as ExtendedUser;

	const accessToken = user?.accessToken;
	useEffect(() => {
		const fetchChartData = async () => {
			try {
				const response = await axios.get(
					"https://bank-dashboard-1tst.onrender.com/transactions/balance-history",
					{
						headers: {
							Authorization: `Bearer ${accessToken}`,
						},
					}
				);
				if (response.data.success) {
					setChartData(response.data.data);
				} else {
					setError("Failed to fetch data");
				}
			} catch (err) {
				console.error("Error fetching chart data:", err);
				setError("Failed to fetch data");
			} finally {
				setLoading(false);
			}
		};

		fetchChartData();
	}, []);

	if (loading) return <p>Loading...</p>;
	if (error) return <p>Error: {error}</p>;

	return (
		<Card>
			<CardContent>
				<ChartContainer config={chartConfig}>
					<ResponsiveContainer width="100%" height={300}>
						<AreaChart
							data={chartData}
							margin={{
								top: 16,
								left: 4,
								right: 8,
								bottom: 4,
							}}
						>
							<CartesianGrid vertical={false} strokeDasharray="3 3" />
							<XAxis
								dataKey="time"
								tickLine={false}
								axisLine={false}
								tickMargin={8}
								tickFormatter={(value) => formatMonth(value)}
							/>
							<YAxis
								domain={[0, "auto"]}
								tickCount={5}
								tickFormatter={(value) => `${value}`}
							/>
							<ChartTooltip
								cursor={false}
								content={<ChartTooltipContent indicator="line" />}
							/>
							<Area
								dataKey="value"
								type="monotone"
								fill="#1814F3"
								fillOpacity={0.2}
								stroke="#1814F3"
							/>
						</AreaChart>
					</ResponsiveContainer>
				</ChartContainer>
			</CardContent>
		</Card>
	);
}

export default LineGraphComponent;
