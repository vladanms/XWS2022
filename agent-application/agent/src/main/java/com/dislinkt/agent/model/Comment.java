package com.dislinkt.agent.model;

import org.springframework.data.annotation.Id;
import org.springframework.data.mongodb.core.mapping.Document;

public class Comment {
    
    private User author;
    private String content;

    public Comment(User user, Comment comment)
    {
        this.author = user;
        this.content = content;
    }

    public User getAuthor()
    {
        return author;
    }

    public String cetContent()
    {
        return content;
    }

    public void setAuthor(User author)
    {
        this.author = author;
    }

    public void setContent(String content)
    {
        this.content = content;
    }
}
