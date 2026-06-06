# Managing roles 

For managing roles the general pattern is giving direct action access for user
like for an endpoint "/add-order" 
if the role is "waiter" then allow.

This won't be possible here since I want to keep this dynamic as roles are not fixed.
We can't hardcode the roles.
               ┌────────────────────────┐
               │   Is User a "Waiter"?  │
               └───────────┬────────────┘
                           │
                 ┌─────────┴─────────┐
                Yes                  No
                 │                   │
        [Access Granted]     [Access Denied]

You have tightly coupled your backend logic to a specific business model. If a café owner wants a "Kitchen Manager" role to perform the same actions, or if you adapt this software for a hotel where the role is "Room Service," you have to edit, re-test, and redeploy your Go codebase.

##  Solution: Decouple the system and make it Permission Based, i.e, Permission-Based Decoupling

To make the system universal, you introduce a layer of indirection: Permissions (sometimes called "Scopes" or "Capabilities").

Instead of checking who the user is (Role), the codebase checks what the user is allowed to do (Permission).

    USER                  ROLE                PERMISSION              API ENDPOINT
┌───────────┐         ┌───────────┐         ┌─────────────┐         ┌───────────────┐
│   Alice   ├────────►│  Waiter   ├────────►│ order:create├────────►│ POST /orders  │
└───────────┘         └───────────┘         └─────────────┘         └───────────────┘

### How it flows at runtime:

1. Authentication: The system identifies that the user is Alice and she belongs to the Waiter role.
2. Permission Check: The backend routes check: "Does Alice's role possess the  order:create  permission?"
3. Execution: If yes, the server processes the order. The endpoint itself has no idea what a "Waiter" is.

If tomorrow you add a new role like Senior Cook, you simply map  order:create  to Senior Cook in the database. The codebase remains untouched.

How to implement:

#### 1. Database driven RBAC

Create tables for each Users, Roles and Permissions and mapping table to connect them.

#### 2. Policy Engine

Think of it like a config file that will be have permission and role  mapped.
like [Role] can [Action] on [Resource] 
Query the engine to get answer. 

#### 3. Scope based tokens

When a user logs in, the authentication server looks up their role/permissions once and bakes their allowed scopes directly into a secure payload (like a session token or JWT).
That token provides scopes and based on that allow or deny.


For this project, think should go with Policy Engine. We can just write the config in .yaml file. Since the whole thing is only run on lan scope based token seems to be an overkill and a separate DB is also not needed since i am building this as native application.