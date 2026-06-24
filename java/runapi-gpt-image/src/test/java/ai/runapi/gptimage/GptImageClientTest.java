package ai.runapi.gptimage;

import static org.junit.jupiter.api.Assertions.assertEquals;
import static org.junit.jupiter.api.Assertions.assertNotNull;
import static org.junit.jupiter.api.Assertions.assertThrows;

import ai.runapi.core.RequestOptions;
import ai.runapi.core.errors.ValidationException;
import ai.runapi.core.http.HttpRequest;
import ai.runapi.core.http.HttpResponse;
import ai.runapi.core.http.HttpTransport;
import ai.runapi.core.http.JsonRequestBody;
import ai.runapi.core.json.Json;
import ai.runapi.gptimage.types.CompletedTextToImageResponse;
import ai.runapi.gptimage.types.TextToImageResponse;
import ai.runapi.gptimage.types.CompletedEditImageResponse;
import ai.runapi.gptimage.types.CompletedTextToImageResponse;
import ai.runapi.gptimage.types.EditImageModel;
import ai.runapi.gptimage.types.EditImageParams;
import ai.runapi.gptimage.types.EditImageResponse;
import ai.runapi.gptimage.types.TextToImageModel;
import ai.runapi.gptimage.types.TextToImageParams;
import ai.runapi.gptimage.types.TextToImageResponse;
import com.fasterxml.jackson.databind.JsonNode;
import java.io.ByteArrayOutputStream;
import java.time.Duration;
import java.util.Collections;
import org.junit.jupiter.api.Test;

class GptImageClientTest {
  @Test
  void builderCreatesClientAndUniversalResources() {
    GptImageClient client = GptImageClient.builder().apiKey("sk-test").build();

    assertNotNull(client.textToImage());
    assertNotNull(client.files());
    assertNotNull(client.account());
  }

  @Test
  void openValueClassesSerializeAsScalarStrings() throws Exception {
    String json = Json.mapper().writeValueAsString(new TextToImageModel("gpt-image-1.5"));

    assertEquals("\"gpt-image-1.5\"", json);
    assertEquals(new TextToImageModel("gpt-image-1.5"), Json.mapper().readValue(json, TextToImageModel.class));
  }

  @Test
  void createSendsExpectedRequestShape() throws Exception {
    CapturingTransport transport = new CapturingTransport("{\"id\":\"task_123\",\"status\":\"processing\"}");
    GptImageClient client = GptImageClient.builder().apiKey("sk-test").transport(transport).build();

    client.textToImage().create(
        TextToImageParams.builder()
            .model(TextToImageModel.GPT_IMAGE_1_5)
            .prompt("A small red cube on a plain white table, studio product photo")
            .aspectRatio("1:1")
            .quality("medium")
            .build()
    );

    assertEquals("POST", transport.request.getMethod().name());
    assertEquals("/api/v1/gpt_image/text_to_image", transport.request.getPath());
    JsonNode body = bodyJson(transport.request);
    assertNotNull(body);
  }

  @Test
  void getDecodesTaskResponseAndExtraFields() {
    CapturingTransport transport = new CapturingTransport("{\"id\":\"task_456\",\"status\":\"completed\",\"images\":[{\"url\":\"https://file.runapi.ai/generated\"}],\"custom\":\"kept\"}");
    GptImageClient client = GptImageClient.builder().apiKey("sk-test").transport(transport).build();

    TextToImageResponse response = client.textToImage().get("task_456");

    assertEquals("GET", transport.request.getMethod().name());
    assertEquals("/api/v1/gpt_image/text_to_image/task_456", transport.request.getPath());
    assertEquals("completed", response.getStatus().value());
    assertNotNull(response.getImages());
    assertEquals("kept", response.extraFields().get("custom").asText());
  }

  @Test
  void runPollsUntilCompletedAndKeepsExtraFields() {
    SequenceTransport transport = new SequenceTransport(
        "{\"id\":\"task_789\",\"status\":\"processing\"}",
        "{\"id\":\"task_789\",\"status\":\"completed\",\"images\":[{\"url\":\"https://file.runapi.ai/generated\"}],\"custom\":\"kept\"}");
    GptImageClient client = GptImageClient.builder().apiKey("sk-test").transport(transport).build();

    CompletedTextToImageResponse response = client.textToImage().run(
        TextToImageParams.builder()
            .model(TextToImageModel.GPT_IMAGE_1_5)
            .prompt("A small red cube on a plain white table, studio product photo")
            .aspectRatio("1:1")
            .quality("medium")
            .build(),
        RequestOptions.builder().pollingInterval(Duration.ofMillis(1)).pollingMaxWait(Duration.ofSeconds(1)).build());

    assertEquals("completed", response.getStatus().value());
    assertNotNull(response.getImages());
    assertEquals("kept", response.extraFields().get("custom").asText());
    assertEquals(2, transport.calls);
  }

  @Test
  void runRejectsCompletedResponseMissingResultField() {
    SequenceTransport transport = new SequenceTransport(
        "{\"id\":\"task_missing\",\"status\":\"processing\"}",
        "{\"id\":\"task_missing\",\"status\":\"completed\"}");
    GptImageClient client = GptImageClient.builder().apiKey("sk-test").transport(transport).build();

    assertThrows(
        ValidationException.class,
        () -> client.textToImage().run(
                TextToImageParams.builder()
                    .model(TextToImageModel.GPT_IMAGE_1_5)
                    .prompt("A small red cube on a plain white table, studio product photo")
                    .aspectRatio("1:1")
                    .quality("medium")
                    .build(),
            RequestOptions.builder().pollingInterval(Duration.ofMillis(1)).pollingMaxWait(Duration.ofSeconds(1)).build()));
  }

    @Test
    void coversEditimageResourceMethods() {
      CapturingTransport createTransport = new CapturingTransport("{\"id\":\"task_edit_image\",\"status\":\"processing\"}");
      GptImageClient createClient = GptImageClient.builder().apiKey("sk-test").transport(createTransport).build();
      assertNotNull(createClient.editImage().create(
              EditImageParams.builder()
                  .model(EditImageModel.GPT_IMAGE_1_5)
                  .prompt("A small red cube on a plain white table, studio product photo")
                  .sourceImageUrls(java.util.Arrays.asList("https://cdn.runapi.ai/public/samples/image.jpg"))
                  .aspectRatio("1:1")
                  .quality("medium")
                  .build()
      ));

      CapturingTransport createWithOptionsTransport = new CapturingTransport("{\"id\":\"task_edit_image_options\",\"status\":\"processing\"}");
      GptImageClient createWithOptionsClient = GptImageClient.builder().apiKey("sk-test").transport(createWithOptionsTransport).build();
      assertNotNull(createWithOptionsClient.editImage().create(
              EditImageParams.builder()
                  .model(EditImageModel.GPT_IMAGE_1_5)
                  .prompt("A small red cube on a plain white table, studio product photo")
                  .sourceImageUrls(java.util.Arrays.asList("https://cdn.runapi.ai/public/samples/image.jpg"))
                  .aspectRatio("1:1")
                  .quality("medium")
                  .build(),
          RequestOptions.none()));

      CapturingTransport getTransport = new CapturingTransport("{\"id\":\"task_edit_image\",\"status\":\"completed\",\"images\":[{\"url\":\"https://file.runapi.ai/generated\"}]}");
      GptImageClient getClient = GptImageClient.builder().apiKey("sk-test").transport(getTransport).build();
      assertNotNull(getClient.editImage().get("task_edit_image"));

      CapturingTransport getWithOptionsTransport = new CapturingTransport("{\"id\":\"task_edit_image_options\",\"status\":\"completed\",\"images\":[{\"url\":\"https://file.runapi.ai/generated\"}]}");
      GptImageClient getWithOptionsClient = GptImageClient.builder().apiKey("sk-test").transport(getWithOptionsTransport).build();
      assertNotNull(getWithOptionsClient.editImage().get("task_edit_image_options", RequestOptions.none()));

      SequenceTransport runTransport = new SequenceTransport(
          "{\"id\":\"task_edit_image_run\",\"status\":\"processing\"}",
          "{\"id\":\"task_edit_image_run\",\"status\":\"completed\",\"images\":[{\"url\":\"https://file.runapi.ai/generated\"}]}");
      GptImageClient runClient = GptImageClient.builder().apiKey("sk-test").transport(runTransport).build();
      CompletedEditImageResponse runResponse = runClient.editImage().run(
              EditImageParams.builder()
                  .model(EditImageModel.GPT_IMAGE_1_5)
                  .prompt("A small red cube on a plain white table, studio product photo")
                  .sourceImageUrls(java.util.Arrays.asList("https://cdn.runapi.ai/public/samples/image.jpg"))
                  .aspectRatio("1:1")
                  .quality("medium")
                  .build(),
          RequestOptions.builder().pollingInterval(Duration.ofMillis(1)).pollingMaxWait(Duration.ofSeconds(1)).build());
      assertNotNull(runResponse);

      SequenceTransport runWithOptionsTransport = new SequenceTransport(
          "{\"id\":\"task_edit_image_run_options\",\"status\":\"processing\"}",
          "{\"id\":\"task_edit_image_run_options\",\"status\":\"completed\",\"images\":[{\"url\":\"https://file.runapi.ai/generated\"}]}");
      GptImageClient runWithOptionsClient = GptImageClient.builder().apiKey("sk-test").transport(runWithOptionsTransport).build();
      assertNotNull(runWithOptionsClient.editImage().run(
              EditImageParams.builder()
                  .model(EditImageModel.GPT_IMAGE_1_5)
                  .prompt("A small red cube on a plain white table, studio product photo")
                  .sourceImageUrls(java.util.Arrays.asList("https://cdn.runapi.ai/public/samples/image.jpg"))
                  .aspectRatio("1:1")
                  .quality("medium")
                  .build(),
          RequestOptions.builder().pollingInterval(Duration.ofMillis(1)).pollingMaxWait(Duration.ofSeconds(1)).build()));
    }

    @Test
    void coversTexttoimageResourceMethods() {
      CapturingTransport createTransport = new CapturingTransport("{\"id\":\"task_text_to_image\",\"status\":\"processing\"}");
      GptImageClient createClient = GptImageClient.builder().apiKey("sk-test").transport(createTransport).build();
      assertNotNull(createClient.textToImage().create(
              TextToImageParams.builder()
                  .model(TextToImageModel.GPT_IMAGE_1_5)
                  .prompt("A small red cube on a plain white table, studio product photo")
                  .aspectRatio("1:1")
                  .quality("medium")
                  .build()
      ));

      CapturingTransport createWithOptionsTransport = new CapturingTransport("{\"id\":\"task_text_to_image_options\",\"status\":\"processing\"}");
      GptImageClient createWithOptionsClient = GptImageClient.builder().apiKey("sk-test").transport(createWithOptionsTransport).build();
      assertNotNull(createWithOptionsClient.textToImage().create(
              TextToImageParams.builder()
                  .model(TextToImageModel.GPT_IMAGE_1_5)
                  .prompt("A small red cube on a plain white table, studio product photo")
                  .aspectRatio("1:1")
                  .quality("medium")
                  .build(),
          RequestOptions.none()));

      CapturingTransport getTransport = new CapturingTransport("{\"id\":\"task_text_to_image\",\"status\":\"completed\",\"images\":[{\"url\":\"https://file.runapi.ai/generated\"}]}");
      GptImageClient getClient = GptImageClient.builder().apiKey("sk-test").transport(getTransport).build();
      assertNotNull(getClient.textToImage().get("task_text_to_image"));

      CapturingTransport getWithOptionsTransport = new CapturingTransport("{\"id\":\"task_text_to_image_options\",\"status\":\"completed\",\"images\":[{\"url\":\"https://file.runapi.ai/generated\"}]}");
      GptImageClient getWithOptionsClient = GptImageClient.builder().apiKey("sk-test").transport(getWithOptionsTransport).build();
      assertNotNull(getWithOptionsClient.textToImage().get("task_text_to_image_options", RequestOptions.none()));

      SequenceTransport runTransport = new SequenceTransport(
          "{\"id\":\"task_text_to_image_run\",\"status\":\"processing\"}",
          "{\"id\":\"task_text_to_image_run\",\"status\":\"completed\",\"images\":[{\"url\":\"https://file.runapi.ai/generated\"}]}");
      GptImageClient runClient = GptImageClient.builder().apiKey("sk-test").transport(runTransport).build();
      CompletedTextToImageResponse runResponse = runClient.textToImage().run(
              TextToImageParams.builder()
                  .model(TextToImageModel.GPT_IMAGE_1_5)
                  .prompt("A small red cube on a plain white table, studio product photo")
                  .aspectRatio("1:1")
                  .quality("medium")
                  .build(),
          RequestOptions.builder().pollingInterval(Duration.ofMillis(1)).pollingMaxWait(Duration.ofSeconds(1)).build());
      assertNotNull(runResponse);

      SequenceTransport runWithOptionsTransport = new SequenceTransport(
          "{\"id\":\"task_text_to_image_run_options\",\"status\":\"processing\"}",
          "{\"id\":\"task_text_to_image_run_options\",\"status\":\"completed\",\"images\":[{\"url\":\"https://file.runapi.ai/generated\"}]}");
      GptImageClient runWithOptionsClient = GptImageClient.builder().apiKey("sk-test").transport(runWithOptionsTransport).build();
      assertNotNull(runWithOptionsClient.textToImage().run(
              TextToImageParams.builder()
                  .model(TextToImageModel.GPT_IMAGE_1_5)
                  .prompt("A small red cube on a plain white table, studio product photo")
                  .aspectRatio("1:1")
                  .quality("medium")
                  .build(),
          RequestOptions.builder().pollingInterval(Duration.ofMillis(1)).pollingMaxWait(Duration.ofSeconds(1)).build()));
    }

  private static JsonNode bodyJson(HttpRequest request) throws Exception {
    JsonRequestBody body = (JsonRequestBody) request.getBody();
    ByteArrayOutputStream out = new ByteArrayOutputStream();
    body.writeTo(out);
    return Json.mapper().readTree(out.toByteArray());
  }

  private static final class CapturingTransport implements HttpTransport {
    private final String body;
    private HttpRequest request;

    private CapturingTransport(String body) {
      this.body = body;
    }

    public HttpResponse send(HttpRequest request) {
      this.request = request;
      return new HttpResponse(200, body, Collections.<String, java.util.List<String>>emptyMap());
    }

    public void close() {}
  }

  private static final class SequenceTransport implements HttpTransport {
    private final String[] responses;
    private int calls;

    private SequenceTransport(String... responses) {
      this.responses = responses;
    }

    public HttpResponse send(HttpRequest request) {
      String response = responses[Math.min(calls, responses.length - 1)];
      calls++;
      return new HttpResponse(200, response, Collections.<String, java.util.List<String>>emptyMap());
    }

    public void close() {}
  }
}
