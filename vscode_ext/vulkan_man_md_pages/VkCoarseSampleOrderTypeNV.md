# VkCoarseSampleOrderTypeNV(3) Manual Page

## Name

VkCoarseSampleOrderTypeNV - Shading rate image sample ordering types



## <a href="#_c_specification" class="anchor"></a>C Specification

The type [VkCoarseSampleOrderTypeNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCoarseSampleOrderTypeNV.html)
specifies the technique used to order coverage samples in fragments
larger than one pixel, and is defined as:

``` c
// Provided by VK_NV_shading_rate_image
typedef enum VkCoarseSampleOrderTypeNV {
    VK_COARSE_SAMPLE_ORDER_TYPE_DEFAULT_NV = 0,
    VK_COARSE_SAMPLE_ORDER_TYPE_CUSTOM_NV = 1,
    VK_COARSE_SAMPLE_ORDER_TYPE_PIXEL_MAJOR_NV = 2,
    VK_COARSE_SAMPLE_ORDER_TYPE_SAMPLE_MAJOR_NV = 3,
} VkCoarseSampleOrderTypeNV;
```

## <a href="#_description" class="anchor"></a>Description

- `VK_COARSE_SAMPLE_ORDER_TYPE_DEFAULT_NV` specifies that coverage
  samples will be ordered in an implementation-dependent manner.

- `VK_COARSE_SAMPLE_ORDER_TYPE_CUSTOM_NV` specifies that coverage
  samples will be ordered according to the array of custom orderings
  provided in either the `pCustomSampleOrders` member of
  `VkPipelineViewportCoarseSampleOrderStateCreateInfoNV` or the
  `pCustomSampleOrders` member of
  [vkCmdSetCoarseSampleOrderNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetCoarseSampleOrderNV.html).

- `VK_COARSE_SAMPLE_ORDER_TYPE_PIXEL_MAJOR_NV` specifies that coverage
  samples will be ordered sequentially, sorted first by pixel coordinate
  (in row-major order) and then by <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#primsrast-multisampling-coverage-mask"
  target="_blank" rel="noopener">sample index</a>.

- `VK_COARSE_SAMPLE_ORDER_TYPE_SAMPLE_MAJOR_NV` specifies that coverage
  samples will be ordered sequentially, sorted first by <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#primsrast-multisampling-coverage-mask"
  target="_blank" rel="noopener">sample index</a> and then by pixel
  coordinate (in row-major order).

## <a href="#_see_also" class="anchor"></a>See Also

[VK_NV_shading_rate_image](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NV_shading_rate_image.html),
[VkPipelineViewportCoarseSampleOrderStateCreateInfoNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineViewportCoarseSampleOrderStateCreateInfoNV.html),
[vkCmdSetCoarseSampleOrderNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetCoarseSampleOrderNV.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkCoarseSampleOrderTypeNV"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
