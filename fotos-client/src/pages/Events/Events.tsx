import { useEffect, useState } from "react";
import { useNavigate, useLocation } from "react-router-dom";
import { dataFetch, getUser } from "../../utils/helperFuntions";
import { Event } from "../../types";
import { Box, Center, Loader, Text } from "@mantine/core";

const Events = () => {
	const navigate = useNavigate();
	const user = getUser();

	const location = useLocation();

	const [isLoading, setIsLoading] = useState(true);
	const [isError, setIsError] = useState(false);
	const [events, setEvents] = useState<Event[]>([]);

	if (!user) {
		navigate("/login");
	}

	useEffect(() => {
		const fetchEvents = async () => {
			const response = await dataFetch(
				"/api/event/get_all",
				user,
				{},
				"GET"
			);
			const responseData = await response.json();
			if (response.ok) {
				setEvents(responseData.events);
			} else {
				setIsError(true);
			}
			setIsLoading(false);
		};
		fetchEvents().catch(console.error);
	}, []);

	if (isLoading) {
		return (
			<Center className="w-full h-full">
				<Loader size={"lg"} />
			</Center>
		);
	}

	if (isError && !events) {
		return <Center className="w-full h-full">Some Error has Occurred</Center>;
	}

	return (
		<>
			{events.map((event) => (
				<Box
					onClick={() => navigate(`/event/${event.eventId}`)}
					className=" transition-all w-full p-2 mb-3 bg-slate-200 rounded-md shadow-sm cursor-pointer hover:bg-slate-100"
					key={event.eventId}
				>
					<Text className="font-title text-h5 font-semibold">
						{event.name}
						<span className="text-h6 font-light">
							{" "}
							by {event.createdBy}
						</span>
					</Text>

					<Text className="font-body text-sm">{event.description}</Text>
				</Box>
			))}
		</>
	);
};

export default Events;
