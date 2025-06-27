# VK_KHR_line_rasterization(3) Manual Page

## Name

VK_KHR_line_rasterization - device extension



## <a href="#_registered_extension_number" class="anchor"></a>Registered Extension Number

535

## <a href="#_revision" class="anchor"></a>Revision

1

## <a href="#_ratification_status" class="anchor"></a>Ratification Status

Ratified

## <a href="#_extension_and_version_dependencies" class="anchor"></a>Extension and Version Dependencies

[VK_KHR_get_physical_device_properties2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_get_physical_device_properties2.html)  
or  
[Version 1.1](#versions-1.1)  

## <a href="#_contact" class="anchor"></a>Contact

- Piers Daniell <a
  href="https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_KHR_line_rasterization%5D%20@pdaniell-nv%0A*Here%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_KHR_line_rasterization%20extension*"
  target="_blank" rel="nofollow noopener"><em></em>pdaniell-nv</a>

## <a href="#_other_extension_metadata" class="anchor"></a>Other Extension Metadata

**Last Modified Date**  
2023-06-08

**IP Status**  
No known IP claims.

**Contributors**  
- Jeff Bolz, NVIDIA

- Allen Jensen, NVIDIA

- Faith Ekstrand, Intel

## <a href="#_description" class="anchor"></a>Description

This extension adds some line rasterization features that are commonly
used in CAD applications and supported in other APIs like OpenGL.
Bresenham-style line rasterization is supported, smooth rectangular
lines (coverage to alpha) are supported, and stippled lines are
supported for all three line rasterization modes.

## <a href="#_new_commands" class="anchor"></a>New Commands

- [vkCmdSetLineStippleKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetLineStippleKHR.html)

## <a href="#_new_structures" class="anchor"></a>New Structures

- Extending [VkPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceFeatures2.html),
  [VkDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceCreateInfo.html):

  - [VkPhysicalDeviceLineRasterizationFeaturesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceLineRasterizationFeaturesKHR.html)

- Extending
  [VkPhysicalDeviceProperties2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceProperties2.html):

  - [VkPhysicalDeviceLineRasterizationPropertiesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceLineRasterizationPropertiesKHR.html)

- Extending
  [VkPipelineRasterizationStateCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineRasterizationStateCreateInfo.html):

  - [VkPipelineRasterizationLineStateCreateInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineRasterizationLineStateCreateInfoKHR.html)

## <a href="#_new_enums" class="anchor"></a>New Enums

- [VkLineRasterizationModeKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkLineRasterizationModeKHR.html)

## <a href="#_new_enum_constants" class="anchor"></a>New Enum Constants

- `VK_KHR_LINE_RASTERIZATION_EXTENSION_NAME`

- `VK_KHR_LINE_RASTERIZATION_SPEC_VERSION`

- Extending [VkDynamicState](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDynamicState.html):

  - `VK_DYNAMIC_STATE_LINE_STIPPLE_KHR`

- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html):

  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_LINE_RASTERIZATION_FEATURES_KHR`

  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_LINE_RASTERIZATION_PROPERTIES_KHR`

  - `VK_STRUCTURE_TYPE_PIPELINE_RASTERIZATION_LINE_STATE_CREATE_INFO_KHR`

## <a href="#_issues" class="anchor"></a>Issues

1\) Do we need to support Bresenham-style and smooth lines with more
than one rasterization sample? i.e. the equivalent of
glDisable(GL_MULTISAMPLE) in OpenGL when the framebuffer has more than
one sample?

**RESOLVED**: Yes. For simplicity, Bresenham line rasterization carries
forward a few restrictions from OpenGL, such as not supporting
per-sample shading, alpha to coverage, or alpha to one.

## <a href="#_version_history" class="anchor"></a>Version History

- Revision 1, 2019-05-09 (Jeff Bolz)

  - Initial draft

## <a href="#_see_also" class="anchor"></a>See Also

No cross-references are available

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VK_KHR_line_rasterization"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is a generated document. Fixes and changes should be made to
the generator scripts, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
