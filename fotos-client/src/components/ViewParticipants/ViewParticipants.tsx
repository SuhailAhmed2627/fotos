import { Card, Center, Loader, SimpleGrid, Text } from "@mantine/core";
import { useEffect, useState } from "react";
import { useParams } from "react-router-dom";
import { dataFetch, getUser } from "../../utils/helperFuntions";
import { MdAccountCircle } from "react-icons/md";

const ViewParticipants = () => {
	const user = getUser();
	const { eventId } = useParams();
	const [participants, setParticipants] = useState<
		{ username: string; name: string }[]
	>([]);
	const [isLoading, setIsLoading] = useState(true);
	const [isError, setIsError] = useState(false);

	useEffect(() => {
		const fetchParticipants = async () => {
			const response = await dataFetch(
				"/api/event/get_users",
				user,
				{
					"Content-Type": "application/json",
				},
				"POST",
				JSON.stringify({
					eventId: eventId,
				})
			);
			const responseData = await response.json();
			if (response.ok) {
				setParticipants(responseData.participants);
			} else {
				setIsError(true);
			}
			setIsLoading(false);
		};
		fetchParticipants().catch(console.error);
	}, []);

	if (isLoading) {
		return (
			<Center className="w-full h-full">
				<Loader size={"lg"} />
			</Center>
		);
	}

	if (isError || !participants) {
		return <Center className="w-full h-full">Some Error has Occurred</Center>;
	}

	if (participants.length === 0) {
		return (
			<Center className="w-full h-full">
				No Images found for this event
			</Center>
		);
	}
	return (
		<SimpleGrid
			cols={4}
			spacing="lg"
			breakpoints={[
				{ maxWidth: "md", cols: 3, spacing: "md" },
				{ maxWidth: "sm", cols: 1, spacing: "sm" },
			]}
		>
			{participants.map((participant) => (
				<Card
					className=" bg-gray-100 flex-col h-[200px] hover:bg-gray-200"
					shadow="sm"
					p="lg"
				>
					<Card.Section className="h-[70%] w-full m-0 flex items-center justify-center">
						<MdAccountCircle className="fill-gray-700" size="5rem" />
					</Card.Section>
					<Text className="text-center text-gray-700 text-h5 font-title">
						{participant.name}
					</Text>
					<Text className="text-center text-gray-700 text-h6 font-title">
						{participant.username}
					</Text>
				</Card>
			))}
		</SimpleGrid>
	);
};

export default ViewParticipants;
