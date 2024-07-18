# VkPipelineCreationFeedbackCreateInfo(3) Manual Page

## Name

VkPipelineCreationFeedbackCreateInfo - Request for feedback about the
creation of a pipeline



## <a href="#_c_specification" class="anchor"></a>C Specification

Feedback about the creation of a particular pipeline object **can** be
obtained by adding a `VkPipelineCreationFeedbackCreateInfo` structure to
the `pNext` chain of
[VkGraphicsPipelineCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkGraphicsPipelineCreateInfo.html),
[VkRayTracingPipelineCreateInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkRayTracingPipelineCreateInfoKHR.html),
[VkRayTracingPipelineCreateInfoNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkRayTracingPipelineCreateInfoNV.html),
or [VkComputePipelineCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkComputePipelineCreateInfo.html). The
`VkPipelineCreationFeedbackCreateInfo` structure is defined as:

``` c
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

``` c
// Provided by VK_EXT_pipeline_creation_feedback
typedef VkPipelineCreationFeedbackCreateInfo VkPipelineCreationFeedbackCreateInfoEXT;
```

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- `pPipelineCreationFeedback` is a pointer to a
  [VkPipelineCreationFeedback](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineCreationFeedback.html)
  structure.

- `pipelineStageCreationFeedbackCount` is the number of elements in
  `pPipelineStageCreationFeedbacks`.

- `pPipelineStageCreationFeedbacks` is a pointer to an array of
  `pipelineStageCreationFeedbackCount`
  [VkPipelineCreationFeedback](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineCreationFeedback.html)
  structures.

## <a href="#_description" class="anchor"></a>Description

An implementation **should** write pipeline creation feedback to
`pPipelineCreationFeedback` and **may** write pipeline stage creation
feedback to `pPipelineStageCreationFeedbacks`. An implementation
**must** set or clear the `VK_PIPELINE_CREATION_FEEDBACK_VALID_BIT` in
[VkPipelineCreationFeedback](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineCreationFeedback.html)::`flags`
for `pPipelineCreationFeedback` and every element of
`pPipelineStageCreationFeedbacks`.

<table>
<colgroup>
<col style="width: 50%" />
<col style="width: 50%" />
</colgroup>
<tbody>
<tr>
<td class="icon"><em></em></td>
<td class="content">Note
<p>One common scenario for an implementation to skip per-stage feedback
is when
<code>VK_PIPELINE_CREATION_FEEDBACK_APPLICATION_PIPELINE_CACHE_HIT_BIT</code>
is set in <code>pPipelineCreationFeedback</code>.</p></td>
</tr>
</tbody>
</table>

When chained to
[VkRayTracingPipelineCreateInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkRayTracingPipelineCreateInfoKHR.html),
[VkRayTracingPipelineCreateInfoNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkRayTracingPipelineCreateInfoNV.html),
or [VkGraphicsPipelineCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkGraphicsPipelineCreateInfo.html),
the `i` element of `pPipelineStageCreationFeedbacks` corresponds to the
`i` element of
[VkRayTracingPipelineCreateInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkRayTracingPipelineCreateInfoKHR.html)::`pStages`,
[VkRayTracingPipelineCreateInfoNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkRayTracingPipelineCreateInfoNV.html)::`pStages`,
or
[VkGraphicsPipelineCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkGraphicsPipelineCreateInfo.html)::`pStages`.
When chained to
[VkComputePipelineCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkComputePipelineCreateInfo.html), the
first element of `pPipelineStageCreationFeedbacks` corresponds to
[VkComputePipelineCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkComputePipelineCreateInfo.html)::`stage`.

Valid Usage (Implicit)

- <a href="#VUID-VkPipelineCreationFeedbackCreateInfo-sType-sType"
  id="VUID-VkPipelineCreationFeedbackCreateInfo-sType-sType"></a>
  VUID-VkPipelineCreationFeedbackCreateInfo-sType-sType  
  `sType` **must** be
  `VK_STRUCTURE_TYPE_PIPELINE_CREATION_FEEDBACK_CREATE_INFO`

- <a
  href="#VUID-VkPipelineCreationFeedbackCreateInfo-pPipelineCreationFeedback-parameter"
  id="VUID-VkPipelineCreationFeedbackCreateInfo-pPipelineCreationFeedback-parameter"></a>
  VUID-VkPipelineCreationFeedbackCreateInfo-pPipelineCreationFeedback-parameter  
  `pPipelineCreationFeedback` **must** be a valid pointer to a
  [VkPipelineCreationFeedback](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineCreationFeedback.html)
  structure

- <a
  href="#VUID-VkPipelineCreationFeedbackCreateInfo-pPipelineStageCreationFeedbacks-parameter"
  id="VUID-VkPipelineCreationFeedbackCreateInfo-pPipelineStageCreationFeedbacks-parameter"></a>
  VUID-VkPipelineCreationFeedbackCreateInfo-pPipelineStageCreationFeedbacks-parameter  
  If `pipelineStageCreationFeedbackCount` is not `0`,
  `pPipelineStageCreationFeedbacks` **must** be a valid pointer to an
  array of `pipelineStageCreationFeedbackCount`
  [VkPipelineCreationFeedback](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineCreationFeedback.html)
  structures

## <a href="#_see_also" class="anchor"></a>See Also

[VK_EXT_pipeline_creation_feedback](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_EXT_pipeline_creation_feedback.html),
[VK_VERSION_1_3](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_VERSION_1_3.html),
[VkComputePipelineCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkComputePipelineCreateInfo.html),
[VkGraphicsPipelineCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkGraphicsPipelineCreateInfo.html),
[VkPipelineCreationFeedback](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineCreationFeedback.html),
[VkRayTracingPipelineCreateInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkRayTracingPipelineCreateInfoKHR.html),
[VkRayTracingPipelineCreateInfoNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkRayTracingPipelineCreateInfoNV.html),
[VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkPipelineCreationFeedbackCreateInfo"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
