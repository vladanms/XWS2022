package com.dislinkt.agent.model;

public enum Role {
    ADMINISTRATOR("Administrator"), REGULAR_USER("Regular user"), COMPANY_OWNER("Company owner");

    public final String value;

    private Role(String value) {
        this.value = value;
    }
}
