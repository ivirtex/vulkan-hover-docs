# VK\_EXT\_line\_rasterization(3) Manual Page

## Name

VK\_EXT\_line\_rasterization - device extension



## [](#_registered_extension_number)Registered Extension Number

260

## [](#_revision)Revision

1

## [](#_ratification_status)Ratification Status

Not ratified

## [](#_extension_and_version_dependencies)Extension and Version Dependencies

[VK\_KHR\_get\_physical\_device\_properties2](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_get_physical_device_properties2.html)  
or  
[Vulkan Version 1.1](#versions-1.1)

## [](#_deprecation_state)Deprecation State

- *Promoted* to [VK\_KHR\_line\_rasterization](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_line_rasterization.html) extension
  
  - Which in turn was *promoted* to [Vulkan 1.4](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#versions-1.4-promotions)

## [](#_special_use)Special Use

- [CAD support](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#extendingvulkan-compatibility-specialuse)

## [](#_contact)Contact

- Jeff Bolz [\[GitHub\]jeffbolznv](https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_EXT_line_rasterization%5D%20%40jeffbolznv%0A%2AHere%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_EXT_line_rasterization%20extension%2A)

## [](#_other_extension_metadata)Other Extension Metadata

**Last Modified Date**

2019-05-09

**IP Status**

No known IP claims.

**Contributors**

- Jeff Bolz, NVIDIA
- Allen Jensen, NVIDIA
- Faith Ekstrand, Intel

## [](#_description)Description

This extension adds some line rasterization features that are commonly used in CAD applications and supported in other APIs like OpenGL. Bresenham-style line rasterization is supported, smooth rectangular lines (coverage to alpha) are supported, and stippled lines are supported for all three line rasterization modes.

## [](#_promotion_to_vk_khr_line_rasterization)Promotion to `VK_KHR_line_rasterization`

All functionality in this extension is included in `VK_KHR_line_rasterization`, with the suffix changed to KHR. The original enum names are still available as aliases of the KHR functionality.

## [](#_new_commands)New Commands

- [vkCmdSetLineStippleEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdSetLineStippleEXT.html)

## [](#_new_structures)New Structures

- Extending [VkPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceFeatures2.html), [VkDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceCreateInfo.html):
  
  - [VkPhysicalDeviceLineRasterizationFeaturesEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceLineRasterizationFeaturesEXT.html)
- Extending [VkPhysicalDeviceProperties2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceProperties2.html):
  
  - [VkPhysicalDeviceLineRasterizationPropertiesEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceLineRasterizationPropertiesEXT.html)
- Extending [VkPipelineRasterizationStateCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineRasterizationStateCreateInfo.html):
  
  - [VkPipelineRasterizationLineStateCreateInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineRasterizationLineStateCreateInfoEXT.html)

## [](#_new_enums)New Enums

- [VkLineRasterizationModeEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkLineRasterizationModeEXT.html)

## [](#_new_enum_constants)New Enum Constants

- `VK_EXT_LINE_RASTERIZATION_EXTENSION_NAME`
- `VK_EXT_LINE_RASTERIZATION_SPEC_VERSION`
- Extending [VkDynamicState](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDynamicState.html):
  
  - `VK_DYNAMIC_STATE_LINE_STIPPLE_EXT`
- Extending [VkLineRasterizationMode](https://registry.khronos.org/vulkan/specs/latest/man/html/VkLineRasterizationMode.html):
  
  - `VK_LINE_RASTERIZATION_MODE_BRESENHAM_EXT`
  - `VK_LINE_RASTERIZATION_MODE_DEFAULT_EXT`
  - `VK_LINE_RASTERIZATION_MODE_RECTANGULAR_EXT`
  - `VK_LINE_RASTERIZATION_MODE_RECTANGULAR_SMOOTH_EXT`
- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html):
  
  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_LINE_RASTERIZATION_FEATURES_EXT`
  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_LINE_RASTERIZATION_PROPERTIES_EXT`
  - `VK_STRUCTURE_TYPE_PIPELINE_RASTERIZATION_LINE_STATE_CREATE_INFO_EXT`

## [](#_issues)Issues

1\) Do we need to support Bresenham-style and smooth lines with more than one rasterization sample? i.e. the equivalent of glDisable(GL\_MULTISAMPLE) in OpenGL when the framebuffer has more than one sample?

**RESOLVED**: Yes. For simplicity, Bresenham line rasterization carries forward a few restrictions from OpenGL, such as not supporting per-sample shading, alpha to coverage, or alpha to one.

## [](#_version_history)Version History

- Revision 1, 2019-05-09 (Jeff Bolz)
  
  - Initial draft

## [](#_see_also)See Also

No cross-references are available

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VK_EXT_line_rasterization)

This page is a generated document. Fixes and changes should be made to the generator scripts, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0