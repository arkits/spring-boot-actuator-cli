package xyz.archit.demo.controllers;

import org.slf4j.Logger;
import org.slf4j.LoggerFactory;
import org.springframework.http.HttpStatus;
import org.springframework.http.ResponseEntity;
import org.springframework.stereotype.Controller;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.RequestParam;
import xyz.archit.demo.models.HttpResponse;

@Controller
public class LogController {

    private static final Logger logger = LoggerFactory.getLogger(LogController.class);

    @GetMapping(
            value = "/log",
            produces = { "application/json" }
    )
    public ResponseEntity<HttpResponse> makeSomeLog(@RequestParam(name = "n", required = false, defaultValue ="1") Integer numberOfLogs) {

        for(int i = 0; i < numberOfLogs; i++){
            logger.info("Some random log here... i={}", i);
        }

        HttpResponse response = new HttpResponse();
        response.message = "OK";
        response.code = 200;

        return new ResponseEntity<>(response, HttpStatus.valueOf(response.code));
    }
    
}
