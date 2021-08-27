# Features
Fetch clickup tasks by ids

# Install
```
go get github.com/quangvo09/clickup-fetch
```

# How to use

### Execute

```
clickup-fetch xxx
```

### Response

```json
[
    {
        "id": "xxx",
        "custom_id": null,
        "name": "Ticket 1",
        "text_content": "\n",
        "description": "\n",
        "status": {
            "status": "to do",
            "color": "#d3d3d3",
            "type": "open",
            "orderindex": 0
        },
        "orderindex": "2.87104567388116900000000000000000",
        "date_created": "1630051869229",
        "date_updated": "1630051869229",
        "date_closed": null,
        "archived": false,
        "creator": {
            "id": 25541457,
            "username": "XXX",
            "color": "#795548",
            "email": "xxx@gmail.com",
            "profilePicture": null
        },
        "assignees": [
            {
                "id": "xxx",
                "username": "xxx",
                "color": "#795548",
                "initials": "Q",
                "email": "xxx@gmail.com",
                "profilePicture": null
            }
        ],
        "watchers": [],
        "checklists": [],
        "tags": [],
        "parent": null,
        "priority": null,
        "due_date": null,
        "start_date": null,
        "points": null,
        "time_estimate": null,
        "custom_fields": [],
        "dependencies": [],
        "linked_tasks": [],
        "team_id": "xxx",
        "url": "https://app.clickup.com/t/xxx",
        "permission_level": "create",
        "list": {
            "id": "xxx",
            "name": "List",
            "access": true
        },
        "project": {
            "id": "xxx",
            "name": "hidden",
            "hidden": true,
            "access": true
        },
        "folder": {
            "id": "xxx",
            "name": "hidden",
            "hidden": true,
            "access": true
        },
        "space": {
            "id": "xxx"
        }
    }
]

```
