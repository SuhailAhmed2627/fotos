import {
	Box,
	Group,
	Tabs,
	Text,
	Image as MantineImage,
	Center,
	Button,
} from "@mantine/core";
import { DropzoneStatus } from "@mantine/dropzone";
import React from "react";
import { MdFileUpload, MdClose, MdInsertPhoto } from "react-icons/md";
import { showNotification } from "@mantine/notifications";
import { dataFetch, getUser } from "../../utils/helperFuntions";
import DropZone from "../DropZone/DropZone";
import { User } from "../../types";

function ImageUploadIcon({
	status,
	size,
}: {
	status: DropzoneStatus;
	size: number;
}) {
	if (status.accepted) {
		return <MdFileUpload size={size} />;
	}

	if (status.rejected) {
		return <MdClose size={size} />;
	}

	return <MdInsertPhoto size={size} />;
}

const eventDropzoneChildren = (status: DropzoneStatus, files: File[]) => {
	if (files.length > 0) {
		return (
			<>
				{files.map((file, index) => {
					let fileSize: string = file.size.toString();
					if (file.size > 1024) {
						fileSize = Math.round(file.size / 1024) + " KB";
					} else if (file.size > 1024 * 1024) {
						fileSize = Math.round(file.size / 1024 / 1024) + " MB";
					}
					return (
						<Box key={file.name} className="mb-4 flex flex-row gap-5">
							<div>
								<MantineImage
									height={70}
									src={URL.createObjectURL(file)}
								></MantineImage>
							</div>
							<Center className="gap-5">
								<Text className="font-body">{file.name}</Text>
								<Text className="font-body">{fileSize}</Text>
							</Center>
						</Box>
					);
				})}
				<Box className="mb-4">Drop Again to Change Images</Box>
			</>
		);
	}
	return (
		<Group
			position="center"
			spacing="xl"
			style={{ minHeight: 220, pointerEvents: "none" }}
		>
			<ImageUploadIcon status={status} size={80} />

			<div>
				<Text size="xl" inline>
					Drag images here or click to select files
				</Text>
				<Text size="sm" color="dimmed" inline mt={7}>
					Attach as many files as you like, each file should not exceed 5mb
				</Text>
			</div>
		</Group>
	);
};

const eventHandleUpload = async (
	user: User,
	files: File[],
	setIsDropZoneLoading: React.Dispatch<React.SetStateAction<boolean>>,
	setDropzoneFiles: React.Dispatch<React.SetStateAction<File[]>>,
	eventId?: number
) => {
	const formData = new FormData();
	const clickedAts: Date[] = [];
	const dimensions: { width: number; height: number }[] = [];

	for (let i = 0; i < files.length; i++) {
		formData.append("files", files[i]);
		clickedAts.push(new Date(files[i].lastModified));
		const img = new Image();
		img.src = URL.createObjectURL(files[i]);
		await img.decode();
		dimensions.push({ width: img.width, height: img.height });
	}
	formData.append("eventId", eventId ? eventId.toString() : "");
	formData.append("clickedAts", JSON.stringify({ clickedAts: clickedAts }));
	formData.append("dimensions", JSON.stringify({ dimensions: dimensions }));

	console.log(files);
	const response = await dataFetch(
		"/api/image/upload",
		user,
		{},
		"POST",
		formData
	);

	const responseData = await response.json();
	setIsDropZoneLoading(false);
	setDropzoneFiles([]);
	if (response.ok) {
		showNotification({
			title: "Upload Successful",
			message: `${responseData.message}`,
			color: "green",
		});
	} else {
		responseData.errors.forEach((errorIndex: number) => {
			showNotification({
				title: "Upload Failed",
				message: `Error Handling File: ${files[errorIndex].name}`,
				color: "red",
			});
		});
	}
};

const ImageUploader = (props: ImageUploaderProps) => {
	return (
		<>
			<Text className="font-body mb-4">
				You can Upload Images from your PC
			</Text>
			<Box>
				<Tabs variant={"pills"}>
					<Tabs.Tab label="From System Storage">
						<DropZone
							dropzoneChildren={eventDropzoneChildren}
							handleUpload={eventHandleUpload}
							eventId={props.eventId}
						/>
					</Tabs.Tab>
				</Tabs>
			</Box>
		</>
	);
};

export default ImageUploader;
