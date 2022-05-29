import { Center, Loader } from "@mantine/core";
import { useState, useEffect } from "react";
import { useNavigate, useParams } from "react-router-dom";
import user from "../../reducers/user";
import { dataFetch, getUser } from "../../utils/helperFuntions";

const JoinEvent = () => {
	const { joinLink } = useParams();
	const [isLoading, setIsLoading] = useState(true);
	const [isError, setIsError] = useState(false);
	const [events, setEvents] = useState<Event[]>([]);
	const navigate = useNavigate();
	const user = getUser();

	if (!user) {
		navigate("/login");
	}

	useEffect(() => {
		const fetchEvents = async () => {
			const response = await dataFetch(
				"/api/event/join",
				user,
				{},
				"POST",
				JSON.stringify({
					url: joinLink,
				})
			);
			const responseData = await response.json();
			if (response.ok) {
				navigate("/event/" + responseData.id);
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

	return <Center className="w-full h-full">Loading</Center>;
};

export default JoinEvent;
