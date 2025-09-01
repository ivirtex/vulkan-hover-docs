# VkPipelineViewportStateCreateInfo(3) Manual Page

## Name

VkPipelineViewportStateCreateInfo - Structure specifying parameters of a newly created pipeline viewport state



## [](#_c_specification)C Specification

The `VkPipelineViewportStateCreateInfo` structure is defined as:

```c++
// Provided by VK_VERSION_1_0
typedef struct VkPipelineViewportStateCreateInfo {
    VkStructureType                       sType;
    const void*                           pNext;
    VkPipelineViewportStateCreateFlags    flags;
    uint32_t                              viewportCount;
    const VkViewport*                     pViewports;
    uint32_t                              scissorCount;
    const VkRect2D*                       pScissors;
} VkPipelineViewportStateCreateInfo;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `flags` is reserved for future use.
- `viewportCount` is the number of viewports used by the pipeline.
- `pViewports` is a pointer to an array of [VkViewport](https://registry.khronos.org/vulkan/specs/latest/man/html/VkViewport.html) structures, defining the viewport transforms. If the viewport state is dynamic, this member is ignored.
- `scissorCount` is the number of [scissors](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#fragops-scissor) and **must** match the number of viewports.
- `pScissors` is a pointer to an array of [VkRect2D](https://registry.khronos.org/vulkan/specs/latest/man/html/VkRect2D.html) structures defining the rectangular bounds of the scissor for the corresponding viewport. If the scissor state is dynamic, this member is ignored.

## [](#_description)Description

Valid Usage

- [](#VUID-VkPipelineViewportStateCreateInfo-viewportCount-01216)VUID-VkPipelineViewportStateCreateInfo-viewportCount-01216  
  If the [`multiViewport`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-multiViewport) feature is not enabled, `viewportCount` **must** not be greater than `1`
- [](#VUID-VkPipelineViewportStateCreateInfo-scissorCount-01217)VUID-VkPipelineViewportStateCreateInfo-scissorCount-01217  
  If the [`multiViewport`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-multiViewport) feature is not enabled, `scissorCount` **must** not be greater than `1`
- [](#VUID-VkPipelineViewportStateCreateInfo-viewportCount-01218)VUID-VkPipelineViewportStateCreateInfo-viewportCount-01218  
  `viewportCount` **must** be less than or equal to `VkPhysicalDeviceLimits`::`maxViewports`
- [](#VUID-VkPipelineViewportStateCreateInfo-scissorCount-01219)VUID-VkPipelineViewportStateCreateInfo-scissorCount-01219  
  `scissorCount` **must** be less than or equal to `VkPhysicalDeviceLimits`::`maxViewports`
- [](#VUID-VkPipelineViewportStateCreateInfo-x-02821)VUID-VkPipelineViewportStateCreateInfo-x-02821  
  The `x` and `y` members of `offset` member of any element of `pScissors` **must** be greater than or equal to `0`
- [](#VUID-VkPipelineViewportStateCreateInfo-offset-02822)VUID-VkPipelineViewportStateCreateInfo-offset-02822  
  Evaluation of (`offset.x` + `extent.width`) **must** not cause a signed integer addition overflow for any element of `pScissors`
- [](#VUID-VkPipelineViewportStateCreateInfo-offset-02823)VUID-VkPipelineViewportStateCreateInfo-offset-02823  
  Evaluation of (`offset.y` + `extent.height`) **must** not cause a signed integer addition overflow for any element of `pScissors`
- [](#VUID-VkPipelineViewportStateCreateInfo-scissorCount-04134)VUID-VkPipelineViewportStateCreateInfo-scissorCount-04134  
  If `scissorCount` and `viewportCount` are both not dynamic, then `scissorCount` and `viewportCount` **must** be identical
- [](#VUID-VkPipelineViewportStateCreateInfo-viewportCount-04135)VUID-VkPipelineViewportStateCreateInfo-viewportCount-04135  
  If the graphics pipeline is being created with `VK_DYNAMIC_STATE_VIEWPORT_WITH_COUNT` set then `viewportCount` **must** be `0`, otherwise `viewportCount` **must** be greater than `0`
- [](#VUID-VkPipelineViewportStateCreateInfo-scissorCount-04136)VUID-VkPipelineViewportStateCreateInfo-scissorCount-04136  
  If the graphics pipeline is being created with `VK_DYNAMIC_STATE_SCISSOR_WITH_COUNT` set then `scissorCount` **must** be `0`, otherwise `scissorCount` **must** be greater than `0`
- [](#VUID-VkPipelineViewportStateCreateInfo-viewportWScalingEnable-01726)VUID-VkPipelineViewportStateCreateInfo-viewportWScalingEnable-01726  
  If the `viewportWScalingEnable` member of a [VkPipelineViewportWScalingStateCreateInfoNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineViewportWScalingStateCreateInfoNV.html) structure included in the `pNext` chain is `VK_TRUE`, the `viewportCount` member of the [VkPipelineViewportWScalingStateCreateInfoNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineViewportWScalingStateCreateInfoNV.html) structure **must** be greater than or equal to [VkPipelineViewportStateCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineViewportStateCreateInfo.html)::`viewportCount`

Valid Usage (Implicit)

- [](#VUID-VkPipelineViewportStateCreateInfo-sType-sType)VUID-VkPipelineViewportStateCreateInfo-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_PIPELINE_VIEWPORT_STATE_CREATE_INFO`
- [](#VUID-VkPipelineViewportStateCreateInfo-pNext-pNext)VUID-VkPipelineViewportStateCreateInfo-pNext-pNext  
  Each `pNext` member of any structure (including this one) in the `pNext` chain **must** be either `NULL` or a pointer to a valid instance of [VkPipelineViewportCoarseSampleOrderStateCreateInfoNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineViewportCoarseSampleOrderStateCreateInfoNV.html), [VkPipelineViewportDepthClampControlCreateInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineViewportDepthClampControlCreateInfoEXT.html), [VkPipelineViewportDepthClipControlCreateInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineViewportDepthClipControlCreateInfoEXT.html), [VkPipelineViewportExclusiveScissorStateCreateInfoNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineViewportExclusiveScissorStateCreateInfoNV.html), [VkPipelineViewportShadingRateImageStateCreateInfoNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineViewportShadingRateImageStateCreateInfoNV.html), [VkPipelineViewportSwizzleStateCreateInfoNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineViewportSwizzleStateCreateInfoNV.html), or [VkPipelineViewportWScalingStateCreateInfoNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineViewportWScalingStateCreateInfoNV.html)
- [](#VUID-VkPipelineViewportStateCreateInfo-sType-unique)VUID-VkPipelineViewportStateCreateInfo-sType-unique  
  The `sType` value of each structure in the `pNext` chain **must** be unique
- [](#VUID-VkPipelineViewportStateCreateInfo-flags-zerobitmask)VUID-VkPipelineViewportStateCreateInfo-flags-zerobitmask  
  `flags` **must** be `0`

## [](#_see_also)See Also

[VK\_VERSION\_1\_0](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_VERSION_1_0.html), [VkGraphicsPipelineCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkGraphicsPipelineCreateInfo.html), [VkPipelineViewportStateCreateFlags](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineViewportStateCreateFlags.html), [VkRect2D](https://registry.khronos.org/vulkan/specs/latest/man/html/VkRect2D.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html), [VkViewport](https://registry.khronos.org/vulkan/specs/latest/man/html/VkViewport.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkPipelineViewportStateCreateInfo).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0