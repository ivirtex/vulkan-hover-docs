# VkCoverageReductionModeNV(3) Manual Page

## Name

VkCoverageReductionModeNV - Specify the coverage reduction mode



## [](#_c_specification)C Specification

Possible values of [VkPipelineCoverageReductionStateCreateInfoNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineCoverageReductionStateCreateInfoNV.html)::`coverageReductionMode`, specifying how color sample coverage is generated from pixel coverage, are:

```c++
// Provided by VK_NV_coverage_reduction_mode
typedef enum VkCoverageReductionModeNV {
    VK_COVERAGE_REDUCTION_MODE_MERGE_NV = 0,
    VK_COVERAGE_REDUCTION_MODE_TRUNCATE_NV = 1,
} VkCoverageReductionModeNV;
```

## [](#_description)Description

- `VK_COVERAGE_REDUCTION_MODE_MERGE_NV` specifies that each color sample will be associated with an implementation-dependent subset of samples in the pixel coverage. If any of those associated samples are covered, the color sample is covered.
- `VK_COVERAGE_REDUCTION_MODE_TRUNCATE_NV` specifies that for color samples present in the color attachments, a color sample is covered if the pixel coverage sample with the same [sample index](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#primsrast-multisampling-coverage-mask) i is covered; other pixel coverage samples are discarded.

## [](#_see_also)See Also

[VK\_NV\_coverage\_reduction\_mode](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NV_coverage_reduction_mode.html), [VkFramebufferMixedSamplesCombinationNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFramebufferMixedSamplesCombinationNV.html), [VkPipelineCoverageReductionStateCreateInfoNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineCoverageReductionStateCreateInfoNV.html), [vkCmdSetCoverageReductionModeNV](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdSetCoverageReductionModeNV.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkCoverageReductionModeNV)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0