# VkPipelineCoverageReductionStateCreateInfoNV(3) Manual Page

## Name

VkPipelineCoverageReductionStateCreateInfoNV - Structure specifying
parameters controlling coverage reduction



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkPipelineCoverageReductionStateCreateInfoNV` structure is defined
as:

``` c
// Provided by VK_NV_coverage_reduction_mode
typedef struct VkPipelineCoverageReductionStateCreateInfoNV {
    VkStructureType                                  sType;
    const void*                                      pNext;
    VkPipelineCoverageReductionStateCreateFlagsNV    flags;
    VkCoverageReductionModeNV                        coverageReductionMode;
} VkPipelineCoverageReductionStateCreateInfoNV;
```

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- `flags` is reserved for future use.

- `coverageReductionMode` is a
  [VkCoverageReductionModeNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCoverageReductionModeNV.html) value
  controlling how color sample coverage is generated from pixel
  coverage.

## <a href="#_description" class="anchor"></a>Description

If this structure is not included in the `pNext` chain, or if the
extension is not enabled, the default coverage reduction mode is
inferred as follows:

- If the
  [`VK_NV_framebuffer_mixed_samples`](VK_NV_framebuffer_mixed_samples.html)
  extension is enabled, then it is as if the `coverageReductionMode` is
  `VK_COVERAGE_REDUCTION_MODE_MERGE_NV`.

- If the
  [`VK_AMD_mixed_attachment_samples`](VK_AMD_mixed_attachment_samples.html)
  extension is enabled, then it is as if the `coverageReductionMode` is
  `VK_COVERAGE_REDUCTION_MODE_TRUNCATE_NV`.

- If both
  [`VK_NV_framebuffer_mixed_samples`](VK_NV_framebuffer_mixed_samples.html)
  and
  [`VK_AMD_mixed_attachment_samples`](VK_AMD_mixed_attachment_samples.html)
  are enabled, then the default coverage reduction mode is
  implementation-dependent.

Valid Usage (Implicit)

- <a href="#VUID-VkPipelineCoverageReductionStateCreateInfoNV-sType-sType"
  id="VUID-VkPipelineCoverageReductionStateCreateInfoNV-sType-sType"></a>
  VUID-VkPipelineCoverageReductionStateCreateInfoNV-sType-sType  
  `sType` **must** be
  `VK_STRUCTURE_TYPE_PIPELINE_COVERAGE_REDUCTION_STATE_CREATE_INFO_NV`

- <a
  href="#VUID-VkPipelineCoverageReductionStateCreateInfoNV-flags-zerobitmask"
  id="VUID-VkPipelineCoverageReductionStateCreateInfoNV-flags-zerobitmask"></a>
  VUID-VkPipelineCoverageReductionStateCreateInfoNV-flags-zerobitmask  
  `flags` **must** be `0`

- <a
  href="#VUID-VkPipelineCoverageReductionStateCreateInfoNV-coverageReductionMode-parameter"
  id="VUID-VkPipelineCoverageReductionStateCreateInfoNV-coverageReductionMode-parameter"></a>
  VUID-VkPipelineCoverageReductionStateCreateInfoNV-coverageReductionMode-parameter  
  `coverageReductionMode` **must** be a valid
  [VkCoverageReductionModeNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCoverageReductionModeNV.html) value

## <a href="#_see_also" class="anchor"></a>See Also

[VK_NV_coverage_reduction_mode](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NV_coverage_reduction_mode.html),
[VkCoverageReductionModeNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCoverageReductionModeNV.html),
[VkPipelineCoverageReductionStateCreateFlagsNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineCoverageReductionStateCreateFlagsNV.html),
[VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkPipelineCoverageReductionStateCreateInfoNV"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
