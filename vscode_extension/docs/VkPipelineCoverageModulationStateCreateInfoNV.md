# VkPipelineCoverageModulationStateCreateInfoNV(3) Manual Page

## Name

VkPipelineCoverageModulationStateCreateInfoNV - Structure specifying parameters controlling coverage modulation



## [](#_c_specification)C Specification

As part of coverage reduction, fragment color values **can** also be modulated (multiplied) by a value that is a function of fraction of covered rasterization samples associated with that color sample.

Pipeline state controlling coverage modulation is specified through the members of the `VkPipelineCoverageModulationStateCreateInfoNV` structure.

The `VkPipelineCoverageModulationStateCreateInfoNV` structure is defined as:

```c++
// Provided by VK_NV_framebuffer_mixed_samples
typedef struct VkPipelineCoverageModulationStateCreateInfoNV {
    VkStructureType                                   sType;
    const void*                                       pNext;
    VkPipelineCoverageModulationStateCreateFlagsNV    flags;
    VkCoverageModulationModeNV                        coverageModulationMode;
    VkBool32                                          coverageModulationTableEnable;
    uint32_t                                          coverageModulationTableCount;
    const float*                                      pCoverageModulationTable;
} VkPipelineCoverageModulationStateCreateInfoNV;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `flags` is reserved for future use.
- `coverageModulationMode` is a [VkCoverageModulationModeNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCoverageModulationModeNV.html) value controlling which color components are modulated.
- `coverageModulationTableEnable` controls whether the modulation factor is looked up from a table in `pCoverageModulationTable`.
- `coverageModulationTableCount` is the number of elements in `pCoverageModulationTable`.
- `pCoverageModulationTable` is a table of modulation factors containing a value for each number of covered samples.

## [](#_description)Description

If `coverageModulationTableEnable` is `VK_FALSE`, then for each color sample the associated bits of the pixel coverage are counted and divided by the number of associated bits to produce a modulation factor R in the range (0,1] (a value of zero would have been killed due to a color coverage of 0). Specifically:

- N = value of `rasterizationSamples`
- M = value of [VkAttachmentDescription](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAttachmentDescription.html)::`samples` for any color attachments
- R = popcount(associated coverage bits) / (N / M)

If `coverageModulationTableEnable` is `VK_TRUE`, the value R is computed using a programmable lookup table. The lookup table has N / M elements, and the element of the table is selected by:

- R = `pCoverageModulationTable`\[popcount(associated coverage bits)-1]

Note that the table does not have an entry for popcount(associated coverage bits) = 0, because such samples would have been killed.

The values of `pCoverageModulationTable` **may** be rounded to an implementation-dependent precision, which is at least as fine as 1 / N, and clamped to \[0,1].

For each color attachment with a floating-point or normalized color format, each fragment output color value is replicated to M values which **can** each be modulated (multiplied) by that color sample’s associated value of R. Which components are modulated is controlled by `coverageModulationMode`.

If this structure is not included in the `pNext` chain, it is as if `coverageModulationMode` is `VK_COVERAGE_MODULATION_MODE_NONE_NV`.

If the [coverage reduction mode](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#fragops-coverage-reduction) is `VK_COVERAGE_REDUCTION_MODE_TRUNCATE_NV`, each color sample is associated with only a single coverage sample. In this case, it is as if `coverageModulationMode` is `VK_COVERAGE_MODULATION_MODE_NONE_NV`.

Valid Usage

- [](#VUID-VkPipelineCoverageModulationStateCreateInfoNV-coverageModulationTableEnable-01405)VUID-VkPipelineCoverageModulationStateCreateInfoNV-coverageModulationTableEnable-01405  
  If `coverageModulationTableEnable` is `VK_TRUE`, `coverageModulationTableCount` **must** be equal to the number of rasterization samples divided by the number of color samples in the subpass

Valid Usage (Implicit)

- [](#VUID-VkPipelineCoverageModulationStateCreateInfoNV-sType-sType)VUID-VkPipelineCoverageModulationStateCreateInfoNV-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_PIPELINE_COVERAGE_MODULATION_STATE_CREATE_INFO_NV`
- [](#VUID-VkPipelineCoverageModulationStateCreateInfoNV-flags-zerobitmask)VUID-VkPipelineCoverageModulationStateCreateInfoNV-flags-zerobitmask  
  `flags` **must** be `0`
- [](#VUID-VkPipelineCoverageModulationStateCreateInfoNV-coverageModulationMode-parameter)VUID-VkPipelineCoverageModulationStateCreateInfoNV-coverageModulationMode-parameter  
  `coverageModulationMode` **must** be a valid [VkCoverageModulationModeNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCoverageModulationModeNV.html) value

## [](#_see_also)See Also

[VK\_NV\_framebuffer\_mixed\_samples](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NV_framebuffer_mixed_samples.html), [VkBool32](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBool32.html), [VkCoverageModulationModeNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCoverageModulationModeNV.html), [VkPipelineCoverageModulationStateCreateFlagsNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineCoverageModulationStateCreateFlagsNV.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkPipelineCoverageModulationStateCreateInfoNV).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0