type DropZoneProps = {
	dropzoneChildren: (status: DropzoneStatus, files: File[]) => JSX.Element;
	handleUpload: (
		user: User,
		files: File[],
		setIsDropZoneLoading: React.Dispatch<React.SetStateAction<boolean>>,
		setDropzoneFiles: React.Dispatch<React.SetStateAction<File[]>>,
		eventId?: number
	) => Promise<void>;
	eventId?: number;
	dispatch?: Dispatch<AnyAction>;
};
