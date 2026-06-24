package ai.runapi.gptimage.types;

import com.fasterxml.jackson.annotation.JsonCreator;

/** Model slug for edit image operations. */
public final class EditImageModel extends GptimageValue {
  /** gpt-image-1.5 model slug. */
  public static final EditImageModel GPT_IMAGE_1_5 = new EditImageModel("gpt-image-1.5");

  /** Creates a model value from a literal model slug. */
  @JsonCreator
  public EditImageModel(String value) {
    super(value);
  }
}
