# VkCoverageModulationModeNV(3) Manual Page

## Name

VkCoverageModulationModeNV - Specify the coverage modulation mode



## <a href="#_c_specification" class="anchor"></a>C Specification

Possible values of
[VkPipelineCoverageModulationStateCreateInfoNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineCoverageModulationStateCreateInfoNV.html)::`coverageModulationMode`,
specifying which color components are modulated, are:

``` c
// Provided by VK_NV_framebuffer_mixed_samples
typedef enum VkCoverageModulationModeNV {
    VK_COVERAGE_MODULATION_MODE_NONE_NV = 0,
    VK_COVERAGE_MODULATION_MODE_RGB_NV = 1,
    VK_COVERAGE_MODULATION_MODE_ALPHA_NV = 2,
    VK_COVERAGE_MODULATION_MODE_RGBA_NV = 3,
} VkCoverageModulationModeNV;
```

## <a href="#_description" class="anchor"></a>Description

- `VK_COVERAGE_MODULATION_MODE_NONE_NV` specifies that no components are
  multiplied by the modulation factor.

- `VK_COVERAGE_MODULATION_MODE_RGB_NV` specifies that the red, green,
  and blue components are multiplied by the modulation factor.

- `VK_COVERAGE_MODULATION_MODE_ALPHA_NV` specifies that the alpha
  component is multiplied by the modulation factor.

- `VK_COVERAGE_MODULATION_MODE_RGBA_NV` specifies that all components
  are multiplied by the modulation factor.

## <a href="#_see_also" class="anchor"></a>See Also

[VK_NV_framebuffer_mixed_samples](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NV_framebuffer_mixed_samples.html),
[VkPipelineCoverageModulationStateCreateInfoNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineCoverageModulationStateCreateInfoNV.html),
[vkCmdSetCoverageModulationModeNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetCoverageModulationModeNV.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkCoverageModulationModeNV"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
