# VkPipelineCreationFeedback(3) Manual Page

## Name

VkPipelineCreationFeedback - Feedback about the creation of a pipeline
or pipeline stage



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkPipelineCreationFeedback` structure is defined as:

``` c
// Provided by VK_VERSION_1_3
typedef struct VkPipelineCreationFeedback {
    VkPipelineCreationFeedbackFlags    flags;
    uint64_t                           duration;
} VkPipelineCreationFeedback;
```

or the equivalent

``` c
// Provided by VK_EXT_pipeline_creation_feedback
typedef VkPipelineCreationFeedback VkPipelineCreationFeedbackEXT;
```

## <a href="#_members" class="anchor"></a>Members

- `flags` is a bitmask of
  [VkPipelineCreationFeedbackFlagBits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineCreationFeedbackFlagBits.html)
  providing feedback about the creation of a pipeline or of a pipeline
  stage.

- `duration` is the duration spent creating a pipeline or pipeline stage
  in nanoseconds.

## <a href="#_description" class="anchor"></a>Description

If the `VK_PIPELINE_CREATION_FEEDBACK_VALID_BIT` is not set in `flags`,
an implementation **must** not set any other bits in `flags`, and the
values of all other `VkPipelineCreationFeedback` data members are
undefined.

## <a href="#_see_also" class="anchor"></a>See Also

[VK_EXT_pipeline_creation_feedback](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_EXT_pipeline_creation_feedback.html),
[VK_VERSION_1_3](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_VERSION_1_3.html),
[VkPipelineCreationFeedbackCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineCreationFeedbackCreateInfo.html),
[VkPipelineCreationFeedbackFlagBits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineCreationFeedbackFlagBits.html),
[VkPipelineCreationFeedbackFlags](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineCreationFeedbackFlags.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkPipelineCreationFeedback"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
