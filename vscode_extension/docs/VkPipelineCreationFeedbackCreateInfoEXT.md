# VkPipelineCreationFeedbackCreateInfo(3) Manual Page

## Name

VkPipelineCreationFeedbackCreateInfo - Request for feedback about the creation of a pipeline



## [](#_c_specification)C Specification

Feedback about the creation of a particular pipeline object **can** be obtained by adding a `VkPipelineCreationFeedbackCreateInfo` structure to the `pNext` chain of [VkGraphicsPipelineCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkGraphicsPipelineCreateInfo.html), [VkRayTracingPipelineCreateInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkRayTracingPipelineCreateInfoKHR.html), [VkRayTracingPipelineCreateInfoNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkRayTracingPipelineCreateInfoNV.html), [VkDataGraphPipelineCreateInfoARM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDataGraphPipelineCreateInfoARM.html), or [VkComputePipelineCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkComputePipelineCreateInfo.html). The `VkPipelineCreationFeedbackCreateInfo` structure is defined as:

```c++
// Provided by VK_VERSION_1_3
typedef struct VkPipelineCreationFeedbackCreateInfo {
    VkStructureType                sType;
    const void*                    pNext;
    VkPipelineCreationFeedback*    pPipelineCreationFeedback;
    uint32_t                       pipelineStageCreationFeedbackCount;
    VkPipelineCreationFeedback*    pPipelineStageCreationFeedbacks;
} VkPipelineCreationFeedbackCreateInfo;
```

or the equivalent

```c++
// Provided by VK_EXT_pipeline_creation_feedback
typedef VkPipelineCreationFeedbackCreateInfo VkPipelineCreationFeedbackCreateInfoEXT;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `pPipelineCreationFeedback` is a pointer to a [VkPipelineCreationFeedback](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineCreationFeedback.html) structure.
- `pipelineStageCreationFeedbackCount` is the number of elements in `pPipelineStageCreationFeedbacks`.
- `pPipelineStageCreationFeedbacks` is a pointer to an array of `pipelineStageCreationFeedbackCount` [VkPipelineCreationFeedback](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineCreationFeedback.html) structures.

## [](#_description)Description

An implementation **should** write pipeline creation feedback to `pPipelineCreationFeedback` and **may** write pipeline stage creation feedback to `pPipelineStageCreationFeedbacks`. An implementation **must** set or clear the `VK_PIPELINE_CREATION_FEEDBACK_VALID_BIT` in [VkPipelineCreationFeedback](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineCreationFeedback.html)::`flags` for `pPipelineCreationFeedback` and every element of `pPipelineStageCreationFeedbacks`.

Note

One common scenario for an implementation to skip per-stage feedback is when `VK_PIPELINE_CREATION_FEEDBACK_APPLICATION_PIPELINE_CACHE_HIT_BIT` is set in `pPipelineCreationFeedback`.

When chained to [VkRayTracingPipelineCreateInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkRayTracingPipelineCreateInfoKHR.html), [VkRayTracingPipelineCreateInfoNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkRayTracingPipelineCreateInfoNV.html), or [VkGraphicsPipelineCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkGraphicsPipelineCreateInfo.html), the `i` element of `pPipelineStageCreationFeedbacks` corresponds to the `i` element of [VkRayTracingPipelineCreateInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkRayTracingPipelineCreateInfoKHR.html)::`pStages`, [VkRayTracingPipelineCreateInfoNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkRayTracingPipelineCreateInfoNV.html)::`pStages`, or [VkGraphicsPipelineCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkGraphicsPipelineCreateInfo.html)::`pStages`. When chained to [VkComputePipelineCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkComputePipelineCreateInfo.html), the first element of `pPipelineStageCreationFeedbacks` corresponds to [VkComputePipelineCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkComputePipelineCreateInfo.html)::`stage`.

Valid Usage (Implicit)

- [](#VUID-VkPipelineCreationFeedbackCreateInfo-sType-sType)VUID-VkPipelineCreationFeedbackCreateInfo-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_PIPELINE_CREATION_FEEDBACK_CREATE_INFO`
- [](#VUID-VkPipelineCreationFeedbackCreateInfo-pPipelineCreationFeedback-parameter)VUID-VkPipelineCreationFeedbackCreateInfo-pPipelineCreationFeedback-parameter  
  `pPipelineCreationFeedback` **must** be a valid pointer to a [VkPipelineCreationFeedback](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineCreationFeedback.html) structure
- [](#VUID-VkPipelineCreationFeedbackCreateInfo-pPipelineStageCreationFeedbacks-parameter)VUID-VkPipelineCreationFeedbackCreateInfo-pPipelineStageCreationFeedbacks-parameter  
  If `pipelineStageCreationFeedbackCount` is not `0`, `pPipelineStageCreationFeedbacks` **must** be a valid pointer to an array of `pipelineStageCreationFeedbackCount` [VkPipelineCreationFeedback](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineCreationFeedback.html) structures

## [](#_see_also)See Also

[VK\_EXT\_pipeline\_creation\_feedback](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_pipeline_creation_feedback.html), [VK\_VERSION\_1\_3](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_VERSION_1_3.html), [VkComputePipelineCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkComputePipelineCreateInfo.html), [VkGraphicsPipelineCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkGraphicsPipelineCreateInfo.html), [VkPipelineCreationFeedback](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineCreationFeedback.html), [VkRayTracingPipelineCreateInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkRayTracingPipelineCreateInfoKHR.html), [VkRayTracingPipelineCreateInfoNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkRayTracingPipelineCreateInfoNV.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkPipelineCreationFeedbackCreateInfo)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0