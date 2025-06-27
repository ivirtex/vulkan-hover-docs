# VkPipelineCoverageModulationStateCreateInfoNV(3) Manual Page

## Name

VkPipelineCoverageModulationStateCreateInfoNV - Structure specifying
parameters controlling coverage modulation



## <a href="#_c_specification" class="anchor"></a>C Specification

As part of coverage reduction, fragment color values **can** also be
modulated (multiplied) by a value that is a function of fraction of
covered rasterization samples associated with that color sample.

Pipeline state controlling coverage modulation is specified through the
members of the `VkPipelineCoverageModulationStateCreateInfoNV`
structure.

The `VkPipelineCoverageModulationStateCreateInfoNV` structure is defined
as:

``` c
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

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- `flags` is reserved for future use.

- `coverageModulationMode` is a
  [VkCoverageModulationModeNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCoverageModulationModeNV.html) value
  controlling which color components are modulated.

- `coverageModulationTableEnable` controls whether the modulation factor
  is looked up from a table in `pCoverageModulationTable`.

- `coverageModulationTableCount` is the number of elements in
  `pCoverageModulationTable`.

- `pCoverageModulationTable` is a table of modulation factors containing
  a value for each number of covered samples.

## <a href="#_description" class="anchor"></a>Description

If `coverageModulationTableEnable` is `VK_FALSE`, then for each color
sample the associated bits of the pixel coverage are counted and divided
by the number of associated bits to produce a modulation factor R in the
range (0,1\] (a value of zero would have been killed due to a color
coverage of 0). Specifically:

- N = value of `rasterizationSamples`

- M = value of
  [VkAttachmentDescription](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAttachmentDescription.html)::`samples` for
  any color attachments

- R = popcount(associated coverage bits) / (N / M)

If `coverageModulationTableEnable` is `VK_TRUE`, the value R is computed
using a programmable lookup table. The lookup table has N / M elements,
and the element of the table is selected by:

- R = `pCoverageModulationTable`\[popcount(associated coverage bits)-1\]

Note that the table does not have an entry for popcount(associated
coverage bits) = 0, because such samples would have been killed.

The values of `pCoverageModulationTable` **may** be rounded to an
implementation-dependent precision, which is at least as fine as 1 / N,
and clamped to \[0,1\].

For each color attachment with a floating-point or normalized color
format, each fragment output color value is replicated to M values which
**can** each be modulated (multiplied) by that color sample’s associated
value of R. Which components are modulated is controlled by
`coverageModulationMode`.

If this structure is not included in the `pNext` chain, it is as if
`coverageModulationMode` is `VK_COVERAGE_MODULATION_MODE_NONE_NV`.

If the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#fragops-coverage-reduction"
target="_blank" rel="noopener">coverage reduction mode</a> is
`VK_COVERAGE_REDUCTION_MODE_TRUNCATE_NV`, each color sample is
associated with only a single coverage sample. In this case, it is as if
`coverageModulationMode` is `VK_COVERAGE_MODULATION_MODE_NONE_NV`.

Valid Usage

- <a
  href="#VUID-VkPipelineCoverageModulationStateCreateInfoNV-coverageModulationTableEnable-01405"
  id="VUID-VkPipelineCoverageModulationStateCreateInfoNV-coverageModulationTableEnable-01405"></a>
  VUID-VkPipelineCoverageModulationStateCreateInfoNV-coverageModulationTableEnable-01405  
  If `coverageModulationTableEnable` is `VK_TRUE`,
  `coverageModulationTableCount` **must** be equal to the number of
  rasterization samples divided by the number of color samples in the
  subpass

Valid Usage (Implicit)

- <a
  href="#VUID-VkPipelineCoverageModulationStateCreateInfoNV-sType-sType"
  id="VUID-VkPipelineCoverageModulationStateCreateInfoNV-sType-sType"></a>
  VUID-VkPipelineCoverageModulationStateCreateInfoNV-sType-sType  
  `sType` **must** be
  `VK_STRUCTURE_TYPE_PIPELINE_COVERAGE_MODULATION_STATE_CREATE_INFO_NV`

- <a
  href="#VUID-VkPipelineCoverageModulationStateCreateInfoNV-flags-zerobitmask"
  id="VUID-VkPipelineCoverageModulationStateCreateInfoNV-flags-zerobitmask"></a>
  VUID-VkPipelineCoverageModulationStateCreateInfoNV-flags-zerobitmask  
  `flags` **must** be `0`

- <a
  href="#VUID-VkPipelineCoverageModulationStateCreateInfoNV-coverageModulationMode-parameter"
  id="VUID-VkPipelineCoverageModulationStateCreateInfoNV-coverageModulationMode-parameter"></a>
  VUID-VkPipelineCoverageModulationStateCreateInfoNV-coverageModulationMode-parameter  
  `coverageModulationMode` **must** be a valid
  [VkCoverageModulationModeNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCoverageModulationModeNV.html) value

## <a href="#_see_also" class="anchor"></a>See Also

[VK_NV_framebuffer_mixed_samples](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NV_framebuffer_mixed_samples.html),
[VkBool32](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBool32.html),
[VkCoverageModulationModeNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCoverageModulationModeNV.html),
[VkPipelineCoverageModulationStateCreateFlagsNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineCoverageModulationStateCreateFlagsNV.html),
[VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkPipelineCoverageModulationStateCreateInfoNV"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
