# VkPipelineViewportCoarseSampleOrderStateCreateInfoNV(3) Manual Page

## Name

VkPipelineViewportCoarseSampleOrderStateCreateInfoNV - Structure specifying parameters controlling sample order in coarse fragments



## [](#_c_specification)C Specification

If the `pNext` chain of [VkPipelineViewportStateCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineViewportStateCreateInfo.html) includes a `VkPipelineViewportCoarseSampleOrderStateCreateInfoNV` structure, then that structure includes parameters controlling the order of coverage samples in fragments larger than one pixel.

The `VkPipelineViewportCoarseSampleOrderStateCreateInfoNV` structure is defined as:

```c++
// Provided by VK_NV_shading_rate_image
typedef struct VkPipelineViewportCoarseSampleOrderStateCreateInfoNV {
    VkStructureType                       sType;
    const void*                           pNext;
    VkCoarseSampleOrderTypeNV             sampleOrderType;
    uint32_t                              customSampleOrderCount;
    const VkCoarseSampleOrderCustomNV*    pCustomSampleOrders;
} VkPipelineViewportCoarseSampleOrderStateCreateInfoNV;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `sampleOrderType` specifies the mechanism used to order coverage samples in fragments larger than one pixel.
- `customSampleOrderCount` specifies the number of custom sample orderings to use when ordering coverage samples.
- `pCustomSampleOrders` is a pointer to an array of `customSampleOrderCount` [VkCoarseSampleOrderCustomNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCoarseSampleOrderCustomNV.html) structures, each structure specifying the coverage sample order for a single combination of fragment area and coverage sample count.

## [](#_description)Description

If this structure is not present, `sampleOrderType` is considered to be `VK_COARSE_SAMPLE_ORDER_TYPE_DEFAULT_NV`.

If `sampleOrderType` is `VK_COARSE_SAMPLE_ORDER_TYPE_CUSTOM_NV`, the coverage sample order used for any combination of fragment area and coverage sample count not enumerated in `pCustomSampleOrders` will be identical to that used for `VK_COARSE_SAMPLE_ORDER_TYPE_DEFAULT_NV`.

If the pipeline was created with `VK_DYNAMIC_STATE_VIEWPORT_COARSE_SAMPLE_ORDER_NV`, the contents of this structure (if present) are ignored, and the coverage sample order is instead specified by [vkCmdSetCoarseSampleOrderNV](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdSetCoarseSampleOrderNV.html).

Valid Usage

- [](#VUID-VkPipelineViewportCoarseSampleOrderStateCreateInfoNV-sampleOrderType-02072)VUID-VkPipelineViewportCoarseSampleOrderStateCreateInfoNV-sampleOrderType-02072  
  If `sampleOrderType` is not `VK_COARSE_SAMPLE_ORDER_TYPE_CUSTOM_NV`, `customSamplerOrderCount` **must** be `0`
- [](#VUID-VkPipelineViewportCoarseSampleOrderStateCreateInfoNV-pCustomSampleOrders-02234)VUID-VkPipelineViewportCoarseSampleOrderStateCreateInfoNV-pCustomSampleOrders-02234  
  The array `pCustomSampleOrders` **must** not contain two structures with matching values for both the `shadingRate` and `sampleCount` members

Valid Usage (Implicit)

- [](#VUID-VkPipelineViewportCoarseSampleOrderStateCreateInfoNV-sType-sType)VUID-VkPipelineViewportCoarseSampleOrderStateCreateInfoNV-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_PIPELINE_VIEWPORT_COARSE_SAMPLE_ORDER_STATE_CREATE_INFO_NV`
- [](#VUID-VkPipelineViewportCoarseSampleOrderStateCreateInfoNV-sampleOrderType-parameter)VUID-VkPipelineViewportCoarseSampleOrderStateCreateInfoNV-sampleOrderType-parameter  
  `sampleOrderType` **must** be a valid [VkCoarseSampleOrderTypeNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCoarseSampleOrderTypeNV.html) value
- [](#VUID-VkPipelineViewportCoarseSampleOrderStateCreateInfoNV-pCustomSampleOrders-parameter)VUID-VkPipelineViewportCoarseSampleOrderStateCreateInfoNV-pCustomSampleOrders-parameter  
  If `customSampleOrderCount` is not `0`, `pCustomSampleOrders` **must** be a valid pointer to an array of `customSampleOrderCount` valid [VkCoarseSampleOrderCustomNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCoarseSampleOrderCustomNV.html) structures

## [](#_see_also)See Also

[VK\_NV\_shading\_rate\_image](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NV_shading_rate_image.html), [VkCoarseSampleOrderCustomNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCoarseSampleOrderCustomNV.html), [VkCoarseSampleOrderTypeNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCoarseSampleOrderTypeNV.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkPipelineViewportCoarseSampleOrderStateCreateInfoNV)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0