import { Box, Button, Center, Loader, Stack, Tabs, Text } from "@mantine/core";
import { useEffect, useState } from "react";
import { useNavigate, useParams } from "react-router-dom";
import { dataFetch, getUser } from "../../utils/helperFuntions";
import { Event } from "../../types";
import { useClipboard } from "@mantine/hooks";
import { config } from "../../utils/config";
import {
	ImageGallery,
	ImageUploader,
	ViewParticipants,
} from "../../components";

const EventPage = () => {
	const clipboard = useClipboard();
	const [isCreator, setIsCreator] = useState(false);
	const [isLoading, setIsLoading] = useState(true);
	const [event, setEvent] = useState<Event | null>(null);
	const [isError, setIsError] = useState(false);
	const navigate = useNavigate();
	const { eventId } = useParams();
	const user = getUser();
	const [activeTab, setActiveTab] = useState(0);

	if (!user) {
		navigate("/login");
	}

	useEffect(() => {
		const fetchEvent = async () => {
			const response = await dataFetch(
				"/api/event/get",
				user,
				{
					"Content-Type": "application/json",
				},
				"POST",
				JSON.stringify({
					eventId,
				})
			);
			const responseData = await response.json();
			if (response.ok) {
				setEvent(responseData.event);
				setIsCreator(responseData.isCreator);
				setIsLoading(false);
			} else {
				setIsError(true);
			}

			if (response.ok) {
				const data = await response.json();
				setIsCreator(data.isCreator);
				setEvent(data.event);
			} else {
				setIsError(true);
			}
			setIsLoading(false);
		};
		fetchEvent().catch(console.error);
	}, [eventId]);

	if (isLoading) {
		return (
			<Center className="w-full h-full">
				<Loader size={"lg"} />
			</Center>
		);
	}

	if (isError || !event) {
		return <Center className="w-full h-full">Some Error has Occurred</Center>;
	}

	return (
		<Stack className="p-10 overflow-x-hidden overflow-y-scroll">
			<Text
				lineClamp={1}
				className="font-title font-bold text-h1 md:text-[4rem] text-gray-700 leading-none"
			>
				{event.name}
			</Text>
			<Text
				lineClamp={2}
				className="font-title text-h5 text-gray-600 leading-none"
			>
				{event.description}
			</Text>
			<Box className="flex gap-2">
				{isCreator && (
					<Button
						onClick={() =>
							clipboard.copy(config.baseUrl + "/event/join/" + event.url)
						}
						variant={"subtle"}
					>
						Click here copy Invite Link
					</Button>
				)}
			</Box>
			<Tabs
				className="flex-grow flex flex-col"
				classNames={{
					body: "flex-grow",
				}}
				grow
				active={activeTab}
				onTabChange={setActiveTab}
			>
				<Tabs.Tab label="All Photos">
					<ImageGallery />
				</Tabs.Tab>
				<Tabs.Tab label="Add Photos">
					<ImageUploader eventId={Number(eventId)} />
				</Tabs.Tab>
				{isCreator && (
					<Tabs.Tab label="View Participants">
						<ViewParticipants />
					</Tabs.Tab>
				)}
			</Tabs>
		</Stack>
	);
};

export default EventPage;
