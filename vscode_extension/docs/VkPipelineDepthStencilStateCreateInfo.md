# VkPipelineDepthStencilStateCreateInfo(3) Manual Page

## Name

VkPipelineDepthStencilStateCreateInfo - Structure specifying parameters of a newly created pipeline depth stencil state



## [](#_c_specification)C Specification

The `VkPipelineDepthStencilStateCreateInfo` structure is defined as:

```c++
// Provided by VK_VERSION_1_0
typedef struct VkPipelineDepthStencilStateCreateInfo {
    VkStructureType                           sType;
    const void*                               pNext;
    VkPipelineDepthStencilStateCreateFlags    flags;
    VkBool32                                  depthTestEnable;
    VkBool32                                  depthWriteEnable;
    VkCompareOp                               depthCompareOp;
    VkBool32                                  depthBoundsTestEnable;
    VkBool32                                  stencilTestEnable;
    VkStencilOpState                          front;
    VkStencilOpState                          back;
    float                                     minDepthBounds;
    float                                     maxDepthBounds;
} VkPipelineDepthStencilStateCreateInfo;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `flags` is a bitmask of [VkPipelineDepthStencilStateCreateFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineDepthStencilStateCreateFlagBits.html) specifying additional depth/stencil state information.
- `depthTestEnable` controls whether [depth testing](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#fragops-depth) is enabled.
- `depthWriteEnable` controls whether [depth writes](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#fragops-depth-write) are enabled when `depthTestEnable` is `VK_TRUE`. Depth writes are always disabled when `depthTestEnable` is `VK_FALSE`.
- `depthCompareOp` is a [VkCompareOp](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCompareOp.html) value specifying the comparison operator to use in the [Depth Comparison](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#fragops-depth-comparison) step of the [depth test](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#fragops-depth).
- `depthBoundsTestEnable` controls whether [depth bounds testing](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#fragops-dbt) is enabled.
- `stencilTestEnable` controls whether [stencil testing](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#fragops-stencil) is enabled.
- `front` and `back` are [VkStencilOpState](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStencilOpState.html) values controlling the corresponding parameters of the [stencil test](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#fragops-stencil).
- `minDepthBounds` is the minimum depth bound used in the [depth bounds test](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#fragops-dbt).
- `maxDepthBounds` is the maximum depth bound used in the [depth bounds test](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#fragops-dbt).

## [](#_description)Description

Valid Usage

- [](#VUID-VkPipelineDepthStencilStateCreateInfo-depthBoundsTestEnable-00598)VUID-VkPipelineDepthStencilStateCreateInfo-depthBoundsTestEnable-00598  
  If the [`depthBounds`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-depthBounds) feature is not enabled, `depthBoundsTestEnable` **must** be `VK_FALSE`
- [](#VUID-VkPipelineDepthStencilStateCreateInfo-separateStencilMaskRef-04453)VUID-VkPipelineDepthStencilStateCreateInfo-separateStencilMaskRef-04453  
  If the `VK_KHR_portability_subset` extension is enabled, and [VkPhysicalDevicePortabilitySubsetFeaturesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDevicePortabilitySubsetFeaturesKHR.html)::`separateStencilMaskRef` is `VK_FALSE`, and the value of [VkPipelineDepthStencilStateCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineDepthStencilStateCreateInfo.html)::`stencilTestEnable` is `VK_TRUE`, and the value of [VkPipelineRasterizationStateCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineRasterizationStateCreateInfo.html)::`cullMode` is `VK_CULL_MODE_NONE`, the value of `reference` in each of the [VkStencilOpState](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStencilOpState.html) structs in `front` and `back` **must** be the same
- [](#VUID-VkPipelineDepthStencilStateCreateInfo-rasterizationOrderDepthAttachmentAccess-06463)VUID-VkPipelineDepthStencilStateCreateInfo-rasterizationOrderDepthAttachmentAccess-06463  
  If the [`rasterizationOrderDepthAttachmentAccess`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-rasterizationOrderDepthAttachmentAccess) feature is not enabled, `flags` **must** not include `VK_PIPELINE_DEPTH_STENCIL_STATE_CREATE_RASTERIZATION_ORDER_ATTACHMENT_DEPTH_ACCESS_BIT_EXT`
- [](#VUID-VkPipelineDepthStencilStateCreateInfo-rasterizationOrderStencilAttachmentAccess-06464)VUID-VkPipelineDepthStencilStateCreateInfo-rasterizationOrderStencilAttachmentAccess-06464  
  If the [`rasterizationOrderStencilAttachmentAccess`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-rasterizationOrderStencilAttachmentAccess) feature is not enabled, `flags` **must** not include `VK_PIPELINE_DEPTH_STENCIL_STATE_CREATE_RASTERIZATION_ORDER_ATTACHMENT_STENCIL_ACCESS_BIT_EXT`

Valid Usage (Implicit)

- [](#VUID-VkPipelineDepthStencilStateCreateInfo-sType-sType)VUID-VkPipelineDepthStencilStateCreateInfo-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_PIPELINE_DEPTH_STENCIL_STATE_CREATE_INFO`
- [](#VUID-VkPipelineDepthStencilStateCreateInfo-pNext-pNext)VUID-VkPipelineDepthStencilStateCreateInfo-pNext-pNext  
  `pNext` **must** be `NULL`
- [](#VUID-VkPipelineDepthStencilStateCreateInfo-flags-parameter)VUID-VkPipelineDepthStencilStateCreateInfo-flags-parameter  
  `flags` **must** be a valid combination of [VkPipelineDepthStencilStateCreateFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineDepthStencilStateCreateFlagBits.html) values
- [](#VUID-VkPipelineDepthStencilStateCreateInfo-depthCompareOp-parameter)VUID-VkPipelineDepthStencilStateCreateInfo-depthCompareOp-parameter  
  `depthCompareOp` **must** be a valid [VkCompareOp](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCompareOp.html) value
- [](#VUID-VkPipelineDepthStencilStateCreateInfo-front-parameter)VUID-VkPipelineDepthStencilStateCreateInfo-front-parameter  
  `front` **must** be a valid [VkStencilOpState](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStencilOpState.html) structure
- [](#VUID-VkPipelineDepthStencilStateCreateInfo-back-parameter)VUID-VkPipelineDepthStencilStateCreateInfo-back-parameter  
  `back` **must** be a valid [VkStencilOpState](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStencilOpState.html) structure

## [](#_see_also)See Also

[VK\_VERSION\_1\_0](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_VERSION_1_0.html), [VkBool32](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBool32.html), [VkCompareOp](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCompareOp.html), [VkGraphicsPipelineCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkGraphicsPipelineCreateInfo.html), [VkPipelineDepthStencilStateCreateFlags](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineDepthStencilStateCreateFlags.html), [VkStencilOpState](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStencilOpState.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkPipelineDepthStencilStateCreateInfo)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0