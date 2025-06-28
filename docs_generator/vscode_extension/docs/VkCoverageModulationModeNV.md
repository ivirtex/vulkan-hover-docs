# VkCoverageModulationModeNV(3) Manual Page

## Name

VkCoverageModulationModeNV - Specify the coverage modulation mode



## [](#_c_specification)C Specification

Possible values of [VkPipelineCoverageModulationStateCreateInfoNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineCoverageModulationStateCreateInfoNV.html)::`coverageModulationMode`, specifying which color components are modulated, are:

```c++
// Provided by VK_NV_framebuffer_mixed_samples
typedef enum VkCoverageModulationModeNV {
    VK_COVERAGE_MODULATION_MODE_NONE_NV = 0,
    VK_COVERAGE_MODULATION_MODE_RGB_NV = 1,
    VK_COVERAGE_MODULATION_MODE_ALPHA_NV = 2,
    VK_COVERAGE_MODULATION_MODE_RGBA_NV = 3,
} VkCoverageModulationModeNV;
```

## [](#_description)Description

- `VK_COVERAGE_MODULATION_MODE_NONE_NV` specifies that no components are multiplied by the modulation factor.
- `VK_COVERAGE_MODULATION_MODE_RGB_NV` specifies that the red, green, and blue components are multiplied by the modulation factor.
- `VK_COVERAGE_MODULATION_MODE_ALPHA_NV` specifies that the alpha component is multiplied by the modulation factor.
- `VK_COVERAGE_MODULATION_MODE_RGBA_NV` specifies that all components are multiplied by the modulation factor.

## [](#_see_also)See Also

[VK\_NV\_framebuffer\_mixed\_samples](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NV_framebuffer_mixed_samples.html), [VkPipelineCoverageModulationStateCreateInfoNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineCoverageModulationStateCreateInfoNV.html), [vkCmdSetCoverageModulationModeNV](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdSetCoverageModulationModeNV.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkCoverageModulationModeNV)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0