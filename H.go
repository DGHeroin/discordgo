package discordgo

import "encoding/json"

type H map[string]interface{}

func ToJson(v interface{}, isPretty ...bool) string {
    if len(isPretty) == 1 && isPretty[0] {
        data, _ := json.MarshalIndent(v, "", "  ")
        return string(data)
    }
    data, _ := json.Marshal(v)
    return string(data)
}
