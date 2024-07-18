# VK_EXT_line_rasterization(3) Manual Page

## Name

VK_EXT_line_rasterization - device extension



## <a href="#_registered_extension_number" class="anchor"></a>Registered Extension Number

260

## <a href="#_revision" class="anchor"></a>Revision

1

## <a href="#_ratification_status" class="anchor"></a>Ratification Status

Not ratified

## <a href="#_extension_and_version_dependencies" class="anchor"></a>Extension and Version Dependencies

[VK_KHR_get_physical_device_properties2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_get_physical_device_properties2.html)  
or  
[Version 1.1](#versions-1.1)  

## <a href="#_deprecation_state" class="anchor"></a>Deprecation State

- *Promoted* to
  [VK_KHR_line_rasterization](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_line_rasterization.html) extension

## <a href="#_special_use" class="anchor"></a>Special Use

- <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#extendingvulkan-compatibility-specialuse"
  target="_blank" rel="noopener">CAD support</a>

## <a href="#_contact" class="anchor"></a>Contact

- Jeff Bolz <a
  href="https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_EXT_line_rasterization%5D%20@jeffbolznv%0A*Here%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_EXT_line_rasterization%20extension*"
  target="_blank" rel="nofollow noopener"><em></em>jeffbolznv</a>

## <a href="#_other_extension_metadata" class="anchor"></a>Other Extension Metadata

**Last Modified Date**  
2019-05-09

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

## <a href="#_promotion_to_vk_khr_line_rasterization" class="anchor"></a>Promotion to `VK_KHR_line_rasterization`

All functionality in this extension is included in
[`VK_KHR_line_rasterization`](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_line_rasterization.html), with the
suffix changed to KHR. The original enum names are still available as
aliases of the KHR functionality.

## <a href="#_new_commands" class="anchor"></a>New Commands

- [vkCmdSetLineStippleEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetLineStippleEXT.html)

## <a href="#_new_structures" class="anchor"></a>New Structures

- Extending [VkPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceFeatures2.html),
  [VkDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceCreateInfo.html):

  - [VkPhysicalDeviceLineRasterizationFeaturesEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceLineRasterizationFeaturesEXT.html)

- Extending
  [VkPhysicalDeviceProperties2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceProperties2.html):

  - [VkPhysicalDeviceLineRasterizationPropertiesEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceLineRasterizationPropertiesEXT.html)

- Extending
  [VkPipelineRasterizationStateCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineRasterizationStateCreateInfo.html):

  - [VkPipelineRasterizationLineStateCreateInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineRasterizationLineStateCreateInfoEXT.html)

## <a href="#_new_enums" class="anchor"></a>New Enums

- [VkLineRasterizationModeEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkLineRasterizationModeEXT.html)

## <a href="#_new_enum_constants" class="anchor"></a>New Enum Constants

- `VK_EXT_LINE_RASTERIZATION_EXTENSION_NAME`

- `VK_EXT_LINE_RASTERIZATION_SPEC_VERSION`

- Extending [VkDynamicState](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDynamicState.html):

  - `VK_DYNAMIC_STATE_LINE_STIPPLE_EXT`

- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html):

  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_LINE_RASTERIZATION_FEATURES_EXT`

  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_LINE_RASTERIZATION_PROPERTIES_EXT`

  - `VK_STRUCTURE_TYPE_PIPELINE_RASTERIZATION_LINE_STATE_CREATE_INFO_EXT`

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
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VK_EXT_line_rasterization"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is a generated document. Fixes and changes should be made to
the generator scripts, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
