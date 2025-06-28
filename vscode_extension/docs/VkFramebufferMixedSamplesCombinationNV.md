# VkFramebufferMixedSamplesCombinationNV(3) Manual Page

## Name

VkFramebufferMixedSamplesCombinationNV - Structure specifying a supported sample count combination



## [](#_c_specification)C Specification

The `VkFramebufferMixedSamplesCombinationNV` structure is defined as:

```c++
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

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `coverageReductionMode` is a [VkCoverageReductionModeNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCoverageReductionModeNV.html) value specifying the coverage reduction mode.
- `rasterizationSamples` is a [VkSampleCountFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSampleCountFlagBits.html) specifying the number of rasterization samples in the supported combination.
- `depthStencilSamples` specifies the number of samples in the depth stencil attachment in the supported combination. A value of 0 indicates the combination does not have a depth stencil attachment.
- `colorSamples` specifies the number of color samples in a color attachment in the supported combination. A value of 0 indicates the combination does not have a color attachment.

## [](#_description)Description

Valid Usage (Implicit)

- [](#VUID-VkFramebufferMixedSamplesCombinationNV-sType-sType)VUID-VkFramebufferMixedSamplesCombinationNV-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_FRAMEBUFFER_MIXED_SAMPLES_COMBINATION_NV`
- [](#VUID-VkFramebufferMixedSamplesCombinationNV-pNext-pNext)VUID-VkFramebufferMixedSamplesCombinationNV-pNext-pNext  
  `pNext` **must** be `NULL`

## [](#_see_also)See Also

[VK\_NV\_coverage\_reduction\_mode](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NV_coverage_reduction_mode.html), [VkCoverageReductionModeNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCoverageReductionModeNV.html), [VkSampleCountFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSampleCountFlagBits.html), [VkSampleCountFlags](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSampleCountFlags.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html), [vkGetPhysicalDeviceSupportedFramebufferMixedSamplesCombinationsNV](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetPhysicalDeviceSupportedFramebufferMixedSamplesCombinationsNV.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkFramebufferMixedSamplesCombinationNV)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0