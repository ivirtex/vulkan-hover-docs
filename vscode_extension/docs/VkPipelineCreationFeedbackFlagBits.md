# VkPipelineCreationFeedbackFlagBits(3) Manual Page

## Name

VkPipelineCreationFeedbackFlagBits - Bitmask specifying pipeline or pipeline stage creation feedback



## [](#_c_specification)C Specification

Possible values of the `flags` member of [VkPipelineCreationFeedback](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineCreationFeedback.html) are:

```c++
// Provided by VK_VERSION_1_3
typedef enum VkPipelineCreationFeedbackFlagBits {
    VK_PIPELINE_CREATION_FEEDBACK_VALID_BIT = 0x00000001,
    VK_PIPELINE_CREATION_FEEDBACK_APPLICATION_PIPELINE_CACHE_HIT_BIT = 0x00000002,
    VK_PIPELINE_CREATION_FEEDBACK_BASE_PIPELINE_ACCELERATION_BIT = 0x00000004,
  // Provided by VK_EXT_pipeline_creation_feedback
    VK_PIPELINE_CREATION_FEEDBACK_VALID_BIT_EXT = VK_PIPELINE_CREATION_FEEDBACK_VALID_BIT,
  // Provided by VK_EXT_pipeline_creation_feedback
    VK_PIPELINE_CREATION_FEEDBACK_APPLICATION_PIPELINE_CACHE_HIT_BIT_EXT = VK_PIPELINE_CREATION_FEEDBACK_APPLICATION_PIPELINE_CACHE_HIT_BIT,
  // Provided by VK_EXT_pipeline_creation_feedback
    VK_PIPELINE_CREATION_FEEDBACK_BASE_PIPELINE_ACCELERATION_BIT_EXT = VK_PIPELINE_CREATION_FEEDBACK_BASE_PIPELINE_ACCELERATION_BIT,
} VkPipelineCreationFeedbackFlagBits;
```

or the equivalent

```c++
// Provided by VK_EXT_pipeline_creation_feedback
typedef VkPipelineCreationFeedbackFlagBits VkPipelineCreationFeedbackFlagBitsEXT;
```

## [](#_description)Description

- `VK_PIPELINE_CREATION_FEEDBACK_VALID_BIT` specifies that the feedback information is valid.
- `VK_PIPELINE_CREATION_FEEDBACK_APPLICATION_PIPELINE_CACHE_HIT_BIT` specifies that a readily usable pipeline or pipeline stage was found in the `pipelineCache` specified by the application in the pipeline creation command.
  
  An implementation **should** set the `VK_PIPELINE_CREATION_FEEDBACK_APPLICATION_PIPELINE_CACHE_HIT_BIT` bit if it was able to avoid the large majority of pipeline or pipeline stage creation work by using the `pipelineCache` parameter of [vkCreateGraphicsPipelines](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCreateGraphicsPipelines.html), [vkCreateRayTracingPipelinesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCreateRayTracingPipelinesKHR.html), [vkCreateRayTracingPipelinesNV](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCreateRayTracingPipelinesNV.html), [vkCreateDataGraphPipelinesARM](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCreateDataGraphPipelinesARM.html), or [vkCreateComputePipelines](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCreateComputePipelines.html). When an implementation sets this bit for the entire pipeline, it **may** leave it unset for any stage.
  
  Note
  
  Implementations are encouraged to provide a meaningful signal to applications using this bit. The intention is to communicate to the application that the pipeline or pipeline stage was created “as fast as it gets” using the pipeline cache provided by the application. If an implementation uses an internal cache, it is discouraged from setting this bit as the feedback would be unactionable.
- `VK_PIPELINE_CREATION_FEEDBACK_BASE_PIPELINE_ACCELERATION_BIT` specifies that the base pipeline specified by the `basePipelineHandle` or `basePipelineIndex` member of the `Vk*PipelineCreateInfo` structure was used to accelerate the creation of the pipeline.
  
  An implementation **should** set the `VK_PIPELINE_CREATION_FEEDBACK_BASE_PIPELINE_ACCELERATION_BIT` bit if it was able to avoid a significant amount of work by using the base pipeline.
  
  Note
  
  While “significant amount of work” is subjective, implementations are encouraged to provide a meaningful signal to applications using this bit. For example, a 1% reduction in duration may not warrant setting this bit, while a 50% reduction would.

## [](#_see_also)See Also

[VK\_EXT\_pipeline\_creation\_feedback](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_pipeline_creation_feedback.html), [VK\_VERSION\_1\_3](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_VERSION_1_3.html), [VkPipelineCreationFeedback](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineCreationFeedback.html), [VkPipelineCreationFeedbackCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineCreationFeedbackCreateInfo.html), [VkPipelineCreationFeedbackFlags](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineCreationFeedbackFlags.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkPipelineCreationFeedbackFlagBits).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0