# VkFramebufferMixedSamplesCombinationNV(3) Manual Page

## Name

VkFramebufferMixedSamplesCombinationNV - Structure specifying a
supported sample count combination



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkFramebufferMixedSamplesCombinationNV` structure is defined as:

``` c
// Provided by VK_NV_coverage_reduction_mode
typedef struct VkFramebufferMixedSamplesCombinationNV {
    VkStructureType              sType;
    void*                        pNext;
    VkCoverageReductionModeNV    coverageReductionMode;
    VkSampleCountFlagBits        rasterizationSamples;
    VkSampleCountFlags           depthStencilSamples;
    VkSampleCountFlags           colorSamples;
} VkFramebufferMixedSamplesCombinationNV;
```

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- `coverageReductionMode` is a
  [VkCoverageReductionModeNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCoverageReductionModeNV.html) value
  specifying the coverage reduction mode.

- `rasterizationSamples` is a
  [VkSampleCountFlagBits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSampleCountFlagBits.html) specifying the
  number of rasterization samples in the supported combination.

- `depthStencilSamples` specifies the number of samples in the depth
  stencil attachment in the supported combination. A value of 0
  indicates the combination does not have a depth stencil attachment.

- `colorSamples` specifies the number of color samples in a color
  attachment in the supported combination. A value of 0 indicates the
  combination does not have a color attachment.

## <a href="#_description" class="anchor"></a>Description

Valid Usage (Implicit)

- <a href="#VUID-VkFramebufferMixedSamplesCombinationNV-sType-sType"
  id="VUID-VkFramebufferMixedSamplesCombinationNV-sType-sType"></a>
  VUID-VkFramebufferMixedSamplesCombinationNV-sType-sType  
  `sType` **must** be
  `VK_STRUCTURE_TYPE_FRAMEBUFFER_MIXED_SAMPLES_COMBINATION_NV`

- <a href="#VUID-VkFramebufferMixedSamplesCombinationNV-pNext-pNext"
  id="VUID-VkFramebufferMixedSamplesCombinationNV-pNext-pNext"></a>
  VUID-VkFramebufferMixedSamplesCombinationNV-pNext-pNext  
  `pNext` **must** be `NULL`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_NV_coverage_reduction_mode](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NV_coverage_reduction_mode.html),
[VkCoverageReductionModeNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCoverageReductionModeNV.html),
[VkSampleCountFlagBits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSampleCountFlagBits.html),
[VkSampleCountFlags](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSampleCountFlags.html),
[VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html),
[vkGetPhysicalDeviceSupportedFramebufferMixedSamplesCombinationsNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceSupportedFramebufferMixedSamplesCombinationsNV.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkFramebufferMixedSamplesCombinationNV"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
