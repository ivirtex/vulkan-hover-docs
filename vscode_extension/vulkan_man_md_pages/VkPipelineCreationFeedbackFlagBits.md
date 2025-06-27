# VkPipelineCreationFeedbackFlagBits(3) Manual Page

## Name

VkPipelineCreationFeedbackFlagBits - Bitmask specifying pipeline or
pipeline stage creation feedback



## <a href="#_c_specification" class="anchor"></a>C Specification

Possible values of the `flags` member of
[VkPipelineCreationFeedback](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineCreationFeedback.html) are:

``` c
// Provided by VK_VERSION_1_3
typedef enum VkPipelineCreationFeedbackFlagBits {
    VK_PIPELINE_CREATION_FEEDBACK_VALID_BIT = 0x00000001,
    VK_PIPELINE_CREATION_FEEDBACK_APPLICATION_PIPELINE_CACHE_HIT_BIT = 0x00000002,
    VK_PIPELINE_CREATION_FEEDBACK_BASE_PIPELINE_ACCELERATION_BIT = 0x00000004,
    VK_PIPELINE_CREATION_FEEDBACK_VALID_BIT_EXT = VK_PIPELINE_CREATION_FEEDBACK_VALID_BIT,
    VK_PIPELINE_CREATION_FEEDBACK_APPLICATION_PIPELINE_CACHE_HIT_BIT_EXT = VK_PIPELINE_CREATION_FEEDBACK_APPLICATION_PIPELINE_CACHE_HIT_BIT,
    VK_PIPELINE_CREATION_FEEDBACK_BASE_PIPELINE_ACCELERATION_BIT_EXT = VK_PIPELINE_CREATION_FEEDBACK_BASE_PIPELINE_ACCELERATION_BIT,
} VkPipelineCreationFeedbackFlagBits;
```

or the equivalent

``` c
// Provided by VK_EXT_pipeline_creation_feedback
typedef VkPipelineCreationFeedbackFlagBits VkPipelineCreationFeedbackFlagBitsEXT;
```

## <a href="#_description" class="anchor"></a>Description

- `VK_PIPELINE_CREATION_FEEDBACK_VALID_BIT` indicates that the feedback
  information is valid.

- `VK_PIPELINE_CREATION_FEEDBACK_APPLICATION_PIPELINE_CACHE_HIT_BIT`
  indicates that a readily usable pipeline or pipeline stage was found
  in the `pipelineCache` specified by the application in the pipeline
  creation command.

  An implementation **should** set the
  `VK_PIPELINE_CREATION_FEEDBACK_APPLICATION_PIPELINE_CACHE_HIT_BIT` bit
  if it was able to avoid the large majority of pipeline or pipeline
  stage creation work by using the `pipelineCache` parameter of
  [vkCreateGraphicsPipelines](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCreateGraphicsPipelines.html),
  [vkCreateRayTracingPipelinesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCreateRayTracingPipelinesKHR.html),
  [vkCreateRayTracingPipelinesNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCreateRayTracingPipelinesNV.html),
  or [vkCreateComputePipelines](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCreateComputePipelines.html). When an
  implementation sets this bit for the entire pipeline, it **may** leave
  it unset for any stage.

  <table>
  <colgroup>
  <col style="width: 50%" />
  <col style="width: 50%" />
  </colgroup>
  <tbody>
  <tr>
  <td class="icon"><em></em></td>
  <td class="content">Note
  <p>Implementations are encouraged to provide a meaningful signal to
  applications using this bit. The intention is to communicate to the
  application that the pipeline or pipeline stage was created “as fast as
  it gets” using the pipeline cache provided by the application. If an
  implementation uses an internal cache, it is discouraged from setting
  this bit as the feedback would be unactionable.</p></td>
  </tr>
  </tbody>
  </table>

- `VK_PIPELINE_CREATION_FEEDBACK_BASE_PIPELINE_ACCELERATION_BIT`
  indicates that the base pipeline specified by the `basePipelineHandle`
  or `basePipelineIndex` member of the `Vk*PipelineCreateInfo` structure
  was used to accelerate the creation of the pipeline.

  An implementation **should** set the
  `VK_PIPELINE_CREATION_FEEDBACK_BASE_PIPELINE_ACCELERATION_BIT` bit if
  it was able to avoid a significant amount of work by using the base
  pipeline.

  <table>
  <colgroup>
  <col style="width: 50%" />
  <col style="width: 50%" />
  </colgroup>
  <tbody>
  <tr>
  <td class="icon"><em></em></td>
  <td class="content">Note
  <p>While “significant amount of work” is subjective, implementations are
  encouraged to provide a meaningful signal to applications using this
  bit. For example, a 1% reduction in duration may not warrant setting
  this bit, while a 50% reduction would.</p></td>
  </tr>
  </tbody>
  </table>

## <a href="#_see_also" class="anchor"></a>See Also

[VK_EXT_pipeline_creation_feedback](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_EXT_pipeline_creation_feedback.html),
[VK_VERSION_1_3](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_VERSION_1_3.html),
[VkPipelineCreationFeedback](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineCreationFeedback.html),
[VkPipelineCreationFeedbackCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineCreationFeedbackCreateInfo.html),
[VkPipelineCreationFeedbackFlags](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineCreationFeedbackFlags.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkPipelineCreationFeedbackFlagBits"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
