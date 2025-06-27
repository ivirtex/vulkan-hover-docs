# VkPipelineCoverageToColorStateCreateInfoNV(3) Manual Page

## Name

VkPipelineCoverageToColorStateCreateInfoNV - Structure specifying
whether fragment coverage replaces a color



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkPipelineCoverageToColorStateCreateInfoNV` structure is defined
as:

``` c
// Provided by VK_NV_fragment_coverage_to_color
typedef struct VkPipelineCoverageToColorStateCreateInfoNV {
    VkStructureType                                sType;
    const void*                                    pNext;
    VkPipelineCoverageToColorStateCreateFlagsNV    flags;
    VkBool32                                       coverageToColorEnable;
    uint32_t                                       coverageToColorLocation;
} VkPipelineCoverageToColorStateCreateInfoNV;
```

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- `flags` is reserved for future use.

- `coverageToColorEnable` controls whether the fragment coverage value
  replaces a fragment color output.

- `coverageToColorLocation` controls which fragment shader color output
  value is replaced.

## <a href="#_description" class="anchor"></a>Description

If the `pNext` chain of
[VkPipelineMultisampleStateCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineMultisampleStateCreateInfo.html)
includes a `VkPipelineCoverageToColorStateCreateInfoNV` structure, then
that structure controls whether the fragment coverage is substituted for
a fragment color output and, if so, which output is replaced.

If `coverageToColorEnable` is `VK_TRUE`, the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#primsrast-multisampling-coverage-mask"
target="_blank" rel="noopener">coverage mask</a> replaces the first
component of the color value corresponding to the fragment shader output
location with `Location` equal to `coverageToColorLocation` and `Index`
equal to zero. If the color attachment format has fewer bits than the
coverage mask, the low bits of the sample coverage mask are taken
without any clamping. If the color attachment format has more bits than
the coverage mask, the high bits of the sample coverage mask are filled
with zeros.

If `coverageToColorEnable` is `VK_FALSE`, these operations are skipped.
If this structure is not included in the `pNext` chain, it is as if
`coverageToColorEnable` is `VK_FALSE`.

Valid Usage

- <a
  href="#VUID-VkPipelineCoverageToColorStateCreateInfoNV-coverageToColorEnable-01404"
  id="VUID-VkPipelineCoverageToColorStateCreateInfoNV-coverageToColorEnable-01404"></a>
  VUID-VkPipelineCoverageToColorStateCreateInfoNV-coverageToColorEnable-01404  
  If `coverageToColorEnable` is `VK_TRUE`, then the render pass subpass
  indicated by
  [VkGraphicsPipelineCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkGraphicsPipelineCreateInfo.html)::`renderPass`
  and
  [VkGraphicsPipelineCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkGraphicsPipelineCreateInfo.html)::`subpass`
  **must** have a color attachment at the location selected by
  `coverageToColorLocation`, with a [VkFormat](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFormat.html) of
  `VK_FORMAT_R8_UINT`, `VK_FORMAT_R8_SINT`, `VK_FORMAT_R16_UINT`,
  `VK_FORMAT_R16_SINT`, `VK_FORMAT_R32_UINT`, or `VK_FORMAT_R32_SINT`

Valid Usage (Implicit)

- <a href="#VUID-VkPipelineCoverageToColorStateCreateInfoNV-sType-sType"
  id="VUID-VkPipelineCoverageToColorStateCreateInfoNV-sType-sType"></a>
  VUID-VkPipelineCoverageToColorStateCreateInfoNV-sType-sType  
  `sType` **must** be
  `VK_STRUCTURE_TYPE_PIPELINE_COVERAGE_TO_COLOR_STATE_CREATE_INFO_NV`

- <a
  href="#VUID-VkPipelineCoverageToColorStateCreateInfoNV-flags-zerobitmask"
  id="VUID-VkPipelineCoverageToColorStateCreateInfoNV-flags-zerobitmask"></a>
  VUID-VkPipelineCoverageToColorStateCreateInfoNV-flags-zerobitmask  
  `flags` **must** be `0`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_NV_fragment_coverage_to_color](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NV_fragment_coverage_to_color.html),
[VkBool32](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBool32.html),
[VkPipelineCoverageToColorStateCreateFlagsNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineCoverageToColorStateCreateFlagsNV.html),
[VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkPipelineCoverageToColorStateCreateInfoNV"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
