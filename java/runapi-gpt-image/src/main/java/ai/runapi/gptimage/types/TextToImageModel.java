package ai.runapi.gptimage.types;

import com.fasterxml.jackson.annotation.JsonCreator;

/** Model slug for text to image operations. */
public final class TextToImageModel extends GptimageValue {
  /** gpt-image-1.5 model slug. */
  public static final TextToImageModel GPT_IMAGE_1_5 = new TextToImageModel("gpt-image-1.5");

  /** Creates a model value from a literal model slug. */
  @JsonCreator
  public TextToImageModel(String value) {
    super(value);
  }
}
