package xyz.archit.demo.models;

import com.fasterxml.jackson.annotation.JsonProperty;

public class HttpResponse {

    @JsonProperty("message")
    public String message;

    @JsonProperty("code")
    public Integer code;

    public String getMessage() {
        return message;
    }

    public void setMessage(String message) {
        this.message = message;
    }
}
