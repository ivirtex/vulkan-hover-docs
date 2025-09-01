# VkRasterizationOrderAMD(3) Manual Page

## Name

VkRasterizationOrderAMD - Specify rasterization order for a graphics pipeline



## [](#_c_specification)C Specification

Possible values of [VkPipelineRasterizationStateRasterizationOrderAMD](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineRasterizationStateRasterizationOrderAMD.html)::`rasterizationOrder`, specifying the primitive rasterization order, are:

```c++
// Provided by VK_AMD_rasterization_order
typedef enum VkRasterizationOrderAMD {
    VK_RASTERIZATION_ORDER_STRICT_AMD = 0,
    VK_RASTERIZATION_ORDER_RELAXED_AMD = 1,
} VkRasterizationOrderAMD;
```

## [](#_description)Description

- `VK_RASTERIZATION_ORDER_STRICT_AMD` specifies that operations for each primitive in a subpass **must** occur in [primitive order](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#drawing-primitive-order).
- `VK_RASTERIZATION_ORDER_RELAXED_AMD` specifies that operations for each primitive in a subpass **may** not occur in [primitive order](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#drawing-primitive-order).

## [](#_see_also)See Also

[VK\_AMD\_rasterization\_order](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_AMD_rasterization_order.html), [VkPipelineRasterizationStateRasterizationOrderAMD](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineRasterizationStateRasterizationOrderAMD.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkRasterizationOrderAMD).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0