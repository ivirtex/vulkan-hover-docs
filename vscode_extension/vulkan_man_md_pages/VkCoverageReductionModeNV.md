# VkCoverageReductionModeNV(3) Manual Page

## Name

VkCoverageReductionModeNV - Specify the coverage reduction mode



## <a href="#_c_specification" class="anchor"></a>C Specification

Possible values of
[VkPipelineCoverageReductionStateCreateInfoNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineCoverageReductionStateCreateInfoNV.html)::`coverageReductionMode`,
specifying how color sample coverage is generated from pixel coverage,
are:

``` c
// Provided by VK_NV_coverage_reduction_mode
typedef enum VkCoverageReductionModeNV {
    VK_COVERAGE_REDUCTION_MODE_MERGE_NV = 0,
    VK_COVERAGE_REDUCTION_MODE_TRUNCATE_NV = 1,
} VkCoverageReductionModeNV;
```

## <a href="#_description" class="anchor"></a>Description

- `VK_COVERAGE_REDUCTION_MODE_MERGE_NV` specifies that each color sample
  will be associated with an implementation-dependent subset of samples
  in the pixel coverage. If any of those associated samples are covered,
  the color sample is covered.

- `VK_COVERAGE_REDUCTION_MODE_TRUNCATE_NV` specifies that for color
  samples present in the color attachments, a color sample is covered if
  the pixel coverage sample with the same <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#primsrast-multisampling-coverage-mask"
  target="_blank" rel="noopener">sample index</a> i is covered; other
  pixel coverage samples are discarded.

## <a href="#_see_also" class="anchor"></a>See Also

[VK_NV_coverage_reduction_mode](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NV_coverage_reduction_mode.html),
[VkFramebufferMixedSamplesCombinationNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFramebufferMixedSamplesCombinationNV.html),
[VkPipelineCoverageReductionStateCreateInfoNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineCoverageReductionStateCreateInfoNV.html),
[vkCmdSetCoverageReductionModeNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetCoverageReductionModeNV.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkCoverageReductionModeNV"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
