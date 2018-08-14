# Springboot2beat

Welcome to Springboot2beat.

Springboot2beat is a [Metricbeat](https://www.elastic.co/guide/en/beats/metricbeat/master/index.html) that leverages the [micrometer.io](http://micrometer.io/) application metrics facade via web, which is an instrumentation framework integrated in the [Spring Boot framework starting with version 2](https://spring.io/blog/2018/03/16/micrometer-spring-boot-2-s-new-application-metrics-collector).

## Getting Started with Springboot2beat

### Requirements
In order to allow access to Spring's metrics actuator endpoint, you need to add the following dependency:
```xml
<dependency>
    <groupId>org.springframework.boot</groupId>
    <artifactId>spring-boot-starter-actuator</artifactId>
</dependency>
```
The web app's application config `src/main/resources/application.yaml` must be adapted to expose the metrics endpoint. Please refer to the [Exposing Endpoints](https://docs.spring.io/spring-boot/docs/current/reference/html/production-ready-endpoints.html#production-ready-endpoints-exposing-endpoints) of the Spring Boot documentation to see which actuator endpoints are enabled for web access by default, in case you are depending on default functionality.
```yaml
management:
    endpoints:
        web.exposure.include: metrics
    server:
        port: 9001
        address: 127.0.0.1
```
A security based suggestion would be change the management port and accessible hosts to ensure the actuator endpoints are not accessible to the public.

### Config
In the `springboot2beat.yml` file you need to set the host name of the "to be watched" Spring Boot 2 web app. Furthermore, is this the place where you define the ingestion service, may that be logstash or elasticsearch.
The metrics section defines which metrics are collected. In order to view all available metrics:
```bash
curl -XGET http://HOST:PORT/actuator/metrics
```

### Run
To run Springboot2beat:
```
./springboot2beat -c springboot2beat.yml -e
```

## Example Document
```json
{
        "_index": "springboot2beat-7.0.0-alpha1-2018.08.14",
        "_type": "doc",
        "_id": "T9MUOmUB2ZEe1SMHweUN",
        "_score": null,
        "_source": {
          "@timestamp": "2018-08-14T20:16:02.339Z",
          "tomcat_global_received_count": 0,
          "tomcat_global_request_max_value": 0.3569999933242798,
          "tomcat_sessions_active_current_value": 0,
          "tomcat_sessions_active_max_value": 0,
          "tomcat_sessions_expired_count": 0,
          "tomcat_sessions_alive_max_value": 0,
          "jvm_memory_committed_value": 854982656,
          "jvm_gc_pause_max": 0,
          "jvm_gc_pause_count": 4,
          "jvm_memory_used_value": 429412576,
          "tomcat_global_request_count": 106,
          "host": {
            "name": "webapp003.my.network.example.com"
          },
          "jvm_threads_peak_value": 64,
          "jvm_memory_max_value": 4518313984,
          "http_server_requests_total_time": 0.7595910429954529,
          "system_cpu_usage_value": 0.10866666585206985,
          "jvm_gc_memory_promoted_count": 13190224,
          "type": "Philips-iMac.home",
          "http_server_requests_max": 0.011510051786899567,
          "tomcat_sessions_rejected_count": 0,
          "beat": {
            "name": "webapp003",
            "hostname": "webapp003.my.network.example.com",
            "version": "7.0.0-alpha1"
          },
          "tomcat_global_error_count": 84,
          "http_server_requests_count": 106,
          "jvm_buffer_count_value": 101,
          "jvm_buffer_memory_used_value": 843776,
          "system_load_average_1m_value": 2.9775390625,
          "jvm_gc_memory_allocated_count": 776123776,
          "jvm_gc_max_data_size_value": 2147483648,
          "tomcat_global_request_total_time": 0.9649999737739563,
          "jvm_buffer_total_capacity_value": 835584,
          "jvm_threads_daemon_value": 54,
          "jvm_gc_pause_total_time": 0.652999997138977,
          "jvm_threads_live_value": 61,
          "tomcat_sessions_created_count": 0
        }
      }
```

## TODO's
* add basic authentication with credentials passed via config file
* Tests
* leverage micrometer's [Percentiles, Histogram and SLA](https://micrometer.io/docs/concepts#_histograms_and_percentiles) feature, if possible
* create Kibana dashboards

## Developer Reference
Please refer to the developers guide using this link [here](https://github.com/philkra/springboot2beat/blob/master/docs/developer-guide.md).
