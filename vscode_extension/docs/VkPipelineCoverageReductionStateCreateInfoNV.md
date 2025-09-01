# VkPipelineCoverageReductionStateCreateInfoNV(3) Manual Page

## Name

VkPipelineCoverageReductionStateCreateInfoNV - Structure specifying parameters controlling coverage reduction



## [](#_c_specification)C Specification

The `VkPipelineCoverageReductionStateCreateInfoNV` structure is defined as:

```c++
// Provided by VK_NV_coverage_reduction_mode
typedef struct VkPipelineCoverageReductionStateCreateInfoNV {
    VkStructureType                                  sType;
    const void*                                      pNext;
    VkPipelineCoverageReductionStateCreateFlagsNV    flags;
    VkCoverageReductionModeNV                        coverageReductionMode;
} VkPipelineCoverageReductionStateCreateInfoNV;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `flags` is reserved for future use.
- `coverageReductionMode` is a [VkCoverageReductionModeNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCoverageReductionModeNV.html) value controlling how color sample coverage is generated from pixel coverage.

## [](#_description)Description

If this structure is not included in the `pNext` chain, or if the extension is not enabled, the default coverage reduction mode is inferred as follows:

- If the `VK_NV_framebuffer_mixed_samples` extension is enabled, then it is as if the `coverageReductionMode` is `VK_COVERAGE_REDUCTION_MODE_MERGE_NV`.
- If the `VK_AMD_mixed_attachment_samples` extension is enabled, then it is as if the `coverageReductionMode` is `VK_COVERAGE_REDUCTION_MODE_TRUNCATE_NV`.
- If both `VK_NV_framebuffer_mixed_samples` and `VK_AMD_mixed_attachment_samples` are enabled, then the default coverage reduction mode is implementation-dependent.

Valid Usage (Implicit)

- [](#VUID-VkPipelineCoverageReductionStateCreateInfoNV-sType-sType)VUID-VkPipelineCoverageReductionStateCreateInfoNV-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_PIPELINE_COVERAGE_REDUCTION_STATE_CREATE_INFO_NV`
- [](#VUID-VkPipelineCoverageReductionStateCreateInfoNV-flags-zerobitmask)VUID-VkPipelineCoverageReductionStateCreateInfoNV-flags-zerobitmask  
  `flags` **must** be `0`
- [](#VUID-VkPipelineCoverageReductionStateCreateInfoNV-coverageReductionMode-parameter)VUID-VkPipelineCoverageReductionStateCreateInfoNV-coverageReductionMode-parameter  
  `coverageReductionMode` **must** be a valid [VkCoverageReductionModeNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCoverageReductionModeNV.html) value

## [](#_see_also)See Also

[VK\_NV\_coverage\_reduction\_mode](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NV_coverage_reduction_mode.html), [VkCoverageReductionModeNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCoverageReductionModeNV.html), [VkPipelineCoverageReductionStateCreateFlagsNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineCoverageReductionStateCreateFlagsNV.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkPipelineCoverageReductionStateCreateInfoNV).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0