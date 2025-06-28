# VkCoarseSampleOrderTypeNV(3) Manual Page

## Name

VkCoarseSampleOrderTypeNV - Shading rate image sample ordering types



## [](#_c_specification)C Specification

The type [VkCoarseSampleOrderTypeNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCoarseSampleOrderTypeNV.html) specifies the technique used to order coverage samples in fragments larger than one pixel, and is defined as:

```c++
// Provided by VK_NV_shading_rate_image
typedef enum VkCoarseSampleOrderTypeNV {
    VK_COARSE_SAMPLE_ORDER_TYPE_DEFAULT_NV = 0,
    VK_COARSE_SAMPLE_ORDER_TYPE_CUSTOM_NV = 1,
    VK_COARSE_SAMPLE_ORDER_TYPE_PIXEL_MAJOR_NV = 2,
    VK_COARSE_SAMPLE_ORDER_TYPE_SAMPLE_MAJOR_NV = 3,
} VkCoarseSampleOrderTypeNV;
```

## [](#_description)Description

- `VK_COARSE_SAMPLE_ORDER_TYPE_DEFAULT_NV` specifies that coverage samples will be ordered in an implementation-dependent manner.
- `VK_COARSE_SAMPLE_ORDER_TYPE_CUSTOM_NV` specifies that coverage samples will be ordered according to the array of custom orderings provided in either the `pCustomSampleOrders` member of `VkPipelineViewportCoarseSampleOrderStateCreateInfoNV` or the `pCustomSampleOrders` member of [vkCmdSetCoarseSampleOrderNV](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdSetCoarseSampleOrderNV.html).
- `VK_COARSE_SAMPLE_ORDER_TYPE_PIXEL_MAJOR_NV` specifies that coverage samples will be ordered sequentially, sorted first by pixel coordinate (in row-major order) and then by [sample index](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#primsrast-multisampling-coverage-mask).
- `VK_COARSE_SAMPLE_ORDER_TYPE_SAMPLE_MAJOR_NV` specifies that coverage samples will be ordered sequentially, sorted first by [sample index](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#primsrast-multisampling-coverage-mask) and then by pixel coordinate (in row-major order).

## [](#_see_also)See Also

[VK\_NV\_shading\_rate\_image](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NV_shading_rate_image.html), [VkPipelineViewportCoarseSampleOrderStateCreateInfoNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineViewportCoarseSampleOrderStateCreateInfoNV.html), [vkCmdSetCoarseSampleOrderNV](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdSetCoarseSampleOrderNV.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkCoarseSampleOrderTypeNV)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0