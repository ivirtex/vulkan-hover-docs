# VkPipelineCoverageToColorStateCreateInfoNV(3) Manual Page

## Name

VkPipelineCoverageToColorStateCreateInfoNV - Structure specifying whether fragment coverage replaces a color



## [](#_c_specification)C Specification

The `VkPipelineCoverageToColorStateCreateInfoNV` structure is defined as:

```c++
// Provided by VK_NV_fragment_coverage_to_color
typedef struct VkPipelineCoverageToColorStateCreateInfoNV {
    VkStructureType                                sType;
    const void*                                    pNext;
    VkPipelineCoverageToColorStateCreateFlagsNV    flags;
    VkBool32                                       coverageToColorEnable;
    uint32_t                                       coverageToColorLocation;
} VkPipelineCoverageToColorStateCreateInfoNV;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `flags` is reserved for future use.
- `coverageToColorEnable` controls whether the fragment coverage value replaces a fragment color output.
- `coverageToColorLocation` controls which fragment shader color output value is replaced.

## [](#_description)Description

If the `pNext` chain of [VkPipelineMultisampleStateCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineMultisampleStateCreateInfo.html) includes a `VkPipelineCoverageToColorStateCreateInfoNV` structure, then that structure controls whether the fragment coverage is substituted for a fragment color output and, if so, which output is replaced.

If `coverageToColorEnable` is `VK_TRUE`, the [coverage mask](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#primsrast-multisampling-coverage-mask) replaces the first component of the color value corresponding to the fragment shader output location with `Location` equal to `coverageToColorLocation` and `Index` equal to zero. If the color attachment format has fewer bits than the coverage mask, the low bits of the sample coverage mask are taken without any clamping. If the color attachment format has more bits than the coverage mask, the high bits of the sample coverage mask are filled with zeros.

If `coverageToColorEnable` is `VK_FALSE`, these operations are skipped. If this structure is not included in the `pNext` chain, it is as if `coverageToColorEnable` is `VK_FALSE`.

Valid Usage

- [](#VUID-VkPipelineCoverageToColorStateCreateInfoNV-coverageToColorEnable-01404)VUID-VkPipelineCoverageToColorStateCreateInfoNV-coverageToColorEnable-01404  
  If `coverageToColorEnable` is `VK_TRUE`, then the render pass subpass indicated by [VkGraphicsPipelineCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkGraphicsPipelineCreateInfo.html)::`renderPass` and [VkGraphicsPipelineCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkGraphicsPipelineCreateInfo.html)::`subpass` **must** have a color attachment at the location selected by `coverageToColorLocation`, with a [VkFormat](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFormat.html) of `VK_FORMAT_R8_UINT`, `VK_FORMAT_R8_SINT`, `VK_FORMAT_R16_UINT`, `VK_FORMAT_R16_SINT`, `VK_FORMAT_R32_UINT`, or `VK_FORMAT_R32_SINT`

Valid Usage (Implicit)

- [](#VUID-VkPipelineCoverageToColorStateCreateInfoNV-sType-sType)VUID-VkPipelineCoverageToColorStateCreateInfoNV-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_PIPELINE_COVERAGE_TO_COLOR_STATE_CREATE_INFO_NV`
- [](#VUID-VkPipelineCoverageToColorStateCreateInfoNV-flags-zerobitmask)VUID-VkPipelineCoverageToColorStateCreateInfoNV-flags-zerobitmask  
  `flags` **must** be `0`

## [](#_see_also)See Also

[VK\_NV\_fragment\_coverage\_to\_color](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NV_fragment_coverage_to_color.html), [VkBool32](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBool32.html), [VkPipelineCoverageToColorStateCreateFlagsNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineCoverageToColorStateCreateFlagsNV.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkPipelineCoverageToColorStateCreateInfoNV)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0