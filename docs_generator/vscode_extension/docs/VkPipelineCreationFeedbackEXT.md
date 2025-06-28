# VkPipelineCreationFeedback(3) Manual Page

## Name

VkPipelineCreationFeedback - Feedback about the creation of a pipeline or pipeline stage



## [](#_c_specification)C Specification

The `VkPipelineCreationFeedback` structure is defined as:

```c++
// Provided by VK_VERSION_1_3
typedef struct VkPipelineCreationFeedback {
    VkPipelineCreationFeedbackFlags    flags;
    uint64_t                           duration;
} VkPipelineCreationFeedback;
```

or the equivalent

```c++
// Provided by VK_EXT_pipeline_creation_feedback
typedef VkPipelineCreationFeedback VkPipelineCreationFeedbackEXT;
```

## [](#_members)Members

- `flags` is a bitmask of [VkPipelineCreationFeedbackFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineCreationFeedbackFlagBits.html) providing feedback about the creation of a pipeline or of a pipeline stage.
- `duration` is the duration spent creating a pipeline or pipeline stage in nanoseconds.

## [](#_description)Description

If the `VK_PIPELINE_CREATION_FEEDBACK_VALID_BIT` is not set in `flags`, an implementation **must** not set any other bits in `flags`, and the values of all other `VkPipelineCreationFeedback` data members are undefined.

## [](#_see_also)See Also

[VK\_EXT\_pipeline\_creation\_feedback](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_pipeline_creation_feedback.html), [VK\_VERSION\_1\_3](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_VERSION_1_3.html), [VkPipelineCreationFeedbackCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineCreationFeedbackCreateInfo.html), [VkPipelineCreationFeedbackFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineCreationFeedbackFlagBits.html), [VkPipelineCreationFeedbackFlags](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineCreationFeedbackFlags.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkPipelineCreationFeedback)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0