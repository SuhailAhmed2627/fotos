import { Button } from "@mantine/core";
import { Dropzone, DropzoneStatus } from "@mantine/dropzone";
import React, { useState } from "react";
import { MdFileUpload, MdClose, MdInsertPhoto } from "react-icons/md";
import { FileRejection } from "react-dropzone";
import { showNotification } from "@mantine/notifications";
import { getUser } from "../../utils/helperFuntions";
import { useDispatch } from "react-redux";

const handleRejection = (files: FileRejection[]): void => {
	files.forEach((file) => {
		showNotification({
			title: "File Rejected",
			message: `Error Uploading ${file.file.name}`,
			color: "red",
		});
	});
};

const DropZone = (props: DropZoneProps) => {
	const [dropzoneStatus, setDropzoneStatus] = useState<DropzoneStatus>({
		accepted: false,
		rejected: false,
	});
	const [dropzoneFiles, setDropzoneFiles] = useState<File[]>([]);
	const [isDropZoneLoading, setIsDropZoneLoading] =
		React.useState<boolean>(false);

	const user = getUser();
	const dispatch = useDispatch();
	return (
		<>
			<Dropzone
				loading={isDropZoneLoading}
				onDrop={(files) => {
					// remove files that are already in dropzonefiles
					const newFiles = files.filter(
						(file) =>
							!dropzoneFiles.some(
								(dropzoneFile) => dropzoneFile.name === file.name
							)
					);
					setDropzoneFiles([...dropzoneFiles, ...newFiles]);
					setDropzoneStatus({ accepted: true, rejected: false });
				}}
				onReject={(files) => handleRejection(files)}
				maxSize={5 * 1024 * 1024}
			>
				{(status) => props.dropzoneChildren(status, dropzoneFiles)}
			</Dropzone>
			<Button
				onClick={() => {
					if (dropzoneFiles.length > 0) {
						setIsDropZoneLoading(true);
						props.handleUpload(
							user,
							dropzoneFiles,
							setIsDropZoneLoading,
							setDropzoneFiles,
							props.eventId as number
						);
					}
				}}
				variant={"filled"}
				classNames={{
					label: "font-medium font-title group-hover:font-semibold",
				}}
				className={
					" transition-all group text-white bg-secondary mt-5 hover:bg-secondary-700"
				}
			>
				Upload
			</Button>
		</>
	);
};

export default DropZone;
