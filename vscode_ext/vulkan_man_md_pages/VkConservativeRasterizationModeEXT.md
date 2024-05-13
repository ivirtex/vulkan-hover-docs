# VkConservativeRasterizationModeEXT(3) Manual Page

## Name

VkConservativeRasterizationModeEXT - Specify the conservative
rasterization mode



## <a href="#_c_specification" class="anchor"></a>C Specification

Possible values of
[VkPipelineRasterizationConservativeStateCreateInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineRasterizationConservativeStateCreateInfoEXT.html)::`conservativeRasterizationMode`,
specifying the conservative rasterization mode are:

``` c
// Provided by VK_EXT_conservative_rasterization
typedef enum VkConservativeRasterizationModeEXT {
    VK_CONSERVATIVE_RASTERIZATION_MODE_DISABLED_EXT = 0,
    VK_CONSERVATIVE_RASTERIZATION_MODE_OVERESTIMATE_EXT = 1,
    VK_CONSERVATIVE_RASTERIZATION_MODE_UNDERESTIMATE_EXT = 2,
} VkConservativeRasterizationModeEXT;
```

## <a href="#_description" class="anchor"></a>Description

- `VK_CONSERVATIVE_RASTERIZATION_MODE_DISABLED_EXT` specifies that
  conservative rasterization is disabled and rasterization proceeds as
  normal.

- `VK_CONSERVATIVE_RASTERIZATION_MODE_OVERESTIMATE_EXT` specifies that
  conservative rasterization is enabled in overestimation mode.

- `VK_CONSERVATIVE_RASTERIZATION_MODE_UNDERESTIMATE_EXT` specifies that
  conservative rasterization is enabled in underestimation mode.

## <a href="#_see_also" class="anchor"></a>See Also

[VK_EXT_conservative_rasterization](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_EXT_conservative_rasterization.html),
[VkPipelineRasterizationConservativeStateCreateInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineRasterizationConservativeStateCreateInfoEXT.html),
[vkCmdSetConservativeRasterizationModeEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetConservativeRasterizationModeEXT.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkConservativeRasterizationModeEXT"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
