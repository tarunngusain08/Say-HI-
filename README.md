# Say-Hi

## Minimum Viable Product (MVP) for Chat Application

### User Authentication
- User registration
- User login/logout
- Password hashing for security

### Messaging System
- Enable users to send text messages
- Real-time updates for new messages
- Message timestamp

### 1-1 Chat
- Create individual chat rooms for each pair of users
- Display sender and receiver information

### Security
- Implement end-to-end encryption for messages (optional but recommended)

### Error Handling
- Implement proper error handling for user inputs and server-side operations
- Provide meaningful error messages to users


## Additional Functional Requirements

### Message Formatting
- Support for basic text formatting (bold, italic, etc.)

### Search Functionality
- Implement a search feature to find and retrieve past messages


## Database Design

For our chat application, we'll employ a hybrid approach, utilizing both traditional relational databases and blob storage for efficient data management.

1. **Relational Database:**
   - **Entities:**
     - **User:** Stores user information, including user details and profile settings.
     - **Message:** Manages message metadata such as sender, receiver, timestamp, and references to message content.
     - **Notification:** Holds user-specific notifications about various events.

2. **Blob Storage:**
   - Stores large message bodies, file attachments, and other binary data.
   - Efficiently handles scalability and performance for non-relational data.


## Entity-Relationship (ER) Relationships:**

1. **User to Message (1:N):**
   - One user can send multiple messages.

2. **User to Chat History (N:N):**
   - Many users can have many chat histories, each associated with a specific chat partner.

3. **User to Notification (1:N):**
   - One user can have multiple notifications.

4. **Message to Notification (1:1 or 1:N, depending on use case):**
   - One message may be associated with one or more notifications, depending on the notification requirements.
