import {
	Stack,
	TextInput,
	Text,
	Box,
	Button,
	Center,
	Loader,
} from "@mantine/core";
import { showNotification } from "@mantine/notifications";
import { AnyAction, Dispatch } from "@reduxjs/toolkit";
import { useState } from "react";
import { MdDangerous } from "react-icons/md";
import { useDispatch } from "react-redux";
import { NavigateFunction, useNavigate } from "react-router-dom";
import { createEvent } from "../../actions/user";
import { User } from "../../types";
import { dataFetch, getUser } from "../../utils/helperFuntions";

const validate = (eventname: string, eventdescription: string) => {
	const errors: string[] = [];

	if (eventname.length < 3) {
		errors.push("Event Name should be Atleast 3 characters");
	}
	if (eventdescription.length < 10) {
		errors.push("Event Sescription should be Atleast 10 characters");
	}

	return errors;
};

const handleCreateEvent = async (
	eventName: string,
	eventDescription: string,
	user: User,
	navigate: NavigateFunction,
	setIsLoading: React.Dispatch<React.SetStateAction<boolean>>
) => {
	const errors = validate(eventName, eventDescription);
	if (errors.length > 0) {
		errors.forEach((error) => {
			showNotification({
				message: error,
				title: "Error",
				color: "red",
				icon: <MdDangerous />,
				autoClose: 2000,
			});
		});
		return;
	}

	const response = await dataFetch(
		"/api/event/create",
		user,
		{
			"Content-Type": "application/json",
		},
		"POST",
		JSON.stringify({
			name: eventName,
			description: eventDescription,
		})
	);

	const data = await response.json();
	if (response.ok) {
		navigate(`/event/${data.id}`);
	} else {
		showNotification({
			message: data.message,
			title: "Error",
			color: "red",
			autoClose: 2000,
		});
	}
	setIsLoading(false);
};

const CreateEvent = () => {
	const [eventName, setEventName] = useState("");
	const [eventDescription, setEventDescription] = useState("");
	const [isLoading, setIsLoading] = useState(false);
	const user = getUser();
	const navigate = useNavigate();

	if (!user) {
		return null;
	}

	if (isLoading) {
		return (
			<Center className="w-full h-full">
				<Loader size={"lg"} />
			</Center>
		);
	}
	return (
		<Stack className="w-full h-full items-center justify-center bg-gray-100 p-5 rounded-md shadow-sm text-gray-800">
			<Text className="font-title text-h3 font-semibold ">
				Create a New Event
			</Text>
			<Box className="w-full md:w-[30%] flex flex-col gap-5">
				<TextInput
					onChange={(e) => setEventName(e.target.value)}
					value={eventName}
					label="Event name"
					placeholder="Enter Event name"
					required
				></TextInput>
				<TextInput
					onChange={(e) => setEventDescription(e.target.value)}
					value={eventDescription}
					label="Event Description"
					placeholder="Enter Event Description"
					required
				></TextInput>
				<Button
					onClick={() => {
						setIsLoading(true);
						handleCreateEvent(
							eventName,
							eventDescription,
							user,
							navigate,
							setIsLoading
						);
					}}
					variant={"filled"}
					classNames={{
						label: "font-medium font-title group-hover:font-semibold",
					}}
					className={
						" transition-all group text-white bg-secondary mt-5 hover:bg-secondary-700"
					}
				>
					Create New Event
				</Button>
			</Box>
		</Stack>
	);
};

export default CreateEvent;
