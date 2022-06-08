package domain

// SurveyQuestionnaire is the root element of the question XML file
// When we get a survey using the ODK API, it will be in XML format:
/*
	#head => not required when creating forms for normal ODK
	<head>
		#model => Required when creating forms for normal ODK
		<model>
			# instance => Required when creating forms for normal ODK
			<instance></instance>  => required primary instance
			<instance></instance>  => optional secondary instance
			<instance></instance><instance id="countries" src="jr://file/country-data.xml"/>  => optional external secondary instance

			#bind => Required when creating forms for normal ODK
			<bind></bind>
		</model>
	</head>
	#body => Required required when creating forms for normal ODK
	<body></body>
*/
// XMl ODK at a high level looks like this:
/*
	<model>
		<instance></instance>
		<bind></bind>
	</model>
	<body></body>
*/
type SurveyQuestionnaire struct {
	Head Head `xml:"head"`
	Body Body `xml:"body"`
}
