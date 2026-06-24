package ai.runapi.gptimage.types;

import java.util.LinkedHashMap;
import java.util.List;
import java.util.Map;

/** Parameters for edit image operations. */
public final class EditImageParams {
  private final String model;
  private final String prompt;
  private final List<String> sourceImageUrls;
  private final String aspectRatio;
  private final String quality;
  private final String callbackUrl;

  private EditImageParams(Builder builder) {
    this.model = builder.model;
    this.prompt = GptimageParamUtils.requireNonBlank(builder.prompt, "prompt");
    this.sourceImageUrls = GptimageParamUtils.requiredStrings(builder.sourceImageUrls, "sourceImageUrls");
    this.aspectRatio = GptimageParamUtils.requireNonBlank(builder.aspectRatio, "aspectRatio");
    this.quality = GptimageParamUtils.requireNonBlank(builder.quality, "quality");
    this.callbackUrl = builder.callbackUrl;
  }

  /** Creates a new EditImageParams builder. */
  public static Builder builder() {
    return new Builder();
  }

  /** Returns the RunAPI action key for this request. */
  public String action() {
    return "gpt-image/edit-image";
  }

  /** Converts these parameters to the JSON request body shape. */
  public Map<String, Object> toMap() {
    Map<String, Object> raw = new LinkedHashMap<String, Object>();
    raw.put("model", GptimageParamUtils.wireValue(model));
    raw.put("prompt", GptimageParamUtils.wireValue(prompt));
    raw.put("source_image_urls", GptimageParamUtils.wireValue(sourceImageUrls));
    raw.put("aspect_ratio", GptimageParamUtils.wireValue(aspectRatio));
    raw.put("quality", GptimageParamUtils.wireValue(quality));
    raw.put("callback_url", GptimageParamUtils.wireValue(callbackUrl));
    return GptimageParamUtils.compact(raw);
  }



  /** Builder for {@link EditImageParams}. */
  public static final class Builder {
    private String model;
    private String prompt;
    private List<String> sourceImageUrls;
    private String aspectRatio;
    private String quality;
    private String callbackUrl;

    private Builder() {}

    /** Sets the model slug using a typed model value. */
    public Builder model(EditImageModel value) {
      this.model = java.util.Objects.requireNonNull(value, "model").value();
      return this;
    }

    /** Sets the model slug using a string value. */
    public Builder model(String value) {
      this.model = GptimageParamUtils.requireNonBlankTrim(value, "model");
      return this;
    }


    /** Sets the text prompt. */
    public Builder prompt(String value) {
      this.prompt = GptimageParamUtils.requireNonBlank(value, "prompt");
      return this;
    }

    /** Sets the source image URLs. */
    public Builder sourceImageUrls(List<String> value) {
      this.sourceImageUrls = value;
      return this;
    }

    /** Sets the output aspect ratio. */
    public Builder aspectRatio(String value) {
      this.aspectRatio = GptimageParamUtils.requireNonBlank(value, "aspectRatio");
      return this;
    }

    /** Sets the quality. */
    public Builder quality(String value) {
      this.quality = GptimageParamUtils.requireNonBlank(value, "quality");
      return this;
    }

    /** Sets the webhook URL for task completion notifications. */
    public Builder callbackUrl(String value) {
      this.callbackUrl = GptimageParamUtils.requireNonBlank(value, "callbackUrl");
      return this;
    }

    /** Builds immutable edit image parameters. */
    public EditImageParams build() {
      return new EditImageParams(this);
    }
  }
}
