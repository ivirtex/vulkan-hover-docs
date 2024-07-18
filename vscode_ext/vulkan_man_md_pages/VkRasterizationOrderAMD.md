# VkRasterizationOrderAMD(3) Manual Page

## Name

VkRasterizationOrderAMD - Specify rasterization order for a graphics
pipeline



## <a href="#_c_specification" class="anchor"></a>C Specification

Possible values of
[VkPipelineRasterizationStateRasterizationOrderAMD](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineRasterizationStateRasterizationOrderAMD.html)::`rasterizationOrder`,
specifying the primitive rasterization order, are:

``` c
// Provided by VK_AMD_rasterization_order
typedef enum VkRasterizationOrderAMD {
    VK_RASTERIZATION_ORDER_STRICT_AMD = 0,
    VK_RASTERIZATION_ORDER_RELAXED_AMD = 1,
} VkRasterizationOrderAMD;
```

## <a href="#_description" class="anchor"></a>Description

- `VK_RASTERIZATION_ORDER_STRICT_AMD` specifies that operations for each
  primitive in a subpass **must** occur in <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#drawing-primitive-order"
  target="_blank" rel="noopener">primitive order</a>.

- `VK_RASTERIZATION_ORDER_RELAXED_AMD` specifies that operations for
  each primitive in a subpass **may** not occur in <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#drawing-primitive-order"
  target="_blank" rel="noopener">primitive order</a>.

## <a href="#_see_also" class="anchor"></a>See Also

[VK_AMD_rasterization_order](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_AMD_rasterization_order.html),
[VkPipelineRasterizationStateRasterizationOrderAMD](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineRasterizationStateRasterizationOrderAMD.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkRasterizationOrderAMD"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
