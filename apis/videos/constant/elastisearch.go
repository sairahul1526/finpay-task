package constant

// elasti search indexes
const (
	VideosIndex = "video"
)

const VideosIndexMapping = `
{
	"mappings":{
		"properties": {
			"video":{
				"properties":{
					"title":{
						"type":"text"
					},
					"description":{
						"type":"text"
					},
					"created":{
						"type":"date"
					}
				}
			}
		}
	}
}`
