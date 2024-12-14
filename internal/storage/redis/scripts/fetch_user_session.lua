local access_token = ARGV[1]
local pattern = "session:" .. access_token .. ":*"
local keys = redis.call("KEYS", pattern)

for _, key in ipairs(keys) do
    local session_data = redis.call("GET", key)
    if session_data then
        local status, session = pcall(cjson.decode, session_data)
        if status and session.access_token == access_token then
            return session_data
        end
    end
end

return nil