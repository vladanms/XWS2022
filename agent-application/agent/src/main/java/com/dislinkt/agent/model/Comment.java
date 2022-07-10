package com.dislinkt.agent.model;

import org.springframework.data.annotation.Id;
import org.springframework.data.mongodb.core.mapping.Document;

public class Comment {
    
    private String id;
    private User author;
    private String content;

    public Comment(String id, User user, Comment comment)
    {
        this.id = id;
        this.author = user;
        this.content = content;
    }

    public User getAuthor()
    {
        return author;
    }

    public String getContent()
    {
        return content;
    }

    public String getId()
    {
        return id;
    }

    public void setId(String id)
    {
        this.id = id;
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
