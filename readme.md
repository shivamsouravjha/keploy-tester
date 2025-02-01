1. Create trigger

curl -X POST "http://localhost:4000/api/triggers" \
     -H "Content-Type: application/json" \
     -d '{
           "type": "scheduled",
           "schedule": "*/10 * * * *", 
           "endpoint": null,
           "payload": null
         }'

2. Create an API Trigger
curl -X POST "http://localhost:4000/api/triggers" \
     -H "Content-Type: application/json" \
     -d '{
           "type": "api",
           "schedule": null,
           "endpoint": "http://example.com/webhook",
           "payload": "{ \"message\": \"Hello, world!\" }"
         }'

3. Get All Triggers
curl -X GET "http://localhost:4000/api/triggers"

4. Get a Specific Trigger
curl -X GET "http://localhost:4000/api/triggers/{trigger_id}"

5. Update a Trigger
curl -X PUT "http://localhost:4000/api/triggers/{trigger_id}" \
     -H "Content-Type: application/json" \
     -d '{
           "type": "scheduled",
           "schedule": "*/15 * * * *", 
           "endpoint": null,
           "payload": null
         }'
6. Delete a Trigger
curl -X DELETE "http://localhost:4000/api/triggers/{trigger_id}"

7. Manually Execute a Trigger
curl -X POST "http://localhost:4000/api/triggers/{trigger_id}/execute"

8. Get Active Events (Last 2 Hours)
curl -X GET "http://localhost:4000/api/events"

9. Get Archived Events
curl -X GET "http://localhost:4000/api/events/archived"

10. Purge Old Events (Deletes Logs Older Than 48 Hours)
curl -X DELETE "http://localhost:4000/api/events/purge"
