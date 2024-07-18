# VK_IMG_relaxed_line_rasterization(3) Manual Page

## Name

VK_IMG_relaxed_line_rasterization - device extension



## <a href="#_registered_extension_number" class="anchor"></a>Registered Extension Number

111

## <a href="#_revision" class="anchor"></a>Revision

1

## <a href="#_ratification_status" class="anchor"></a>Ratification Status

Not ratified

## <a href="#_extension_and_version_dependencies" class="anchor"></a>Extension and Version Dependencies

[VK_KHR_get_physical_device_properties2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_get_physical_device_properties2.html)  
or  
[Version 1.1](#versions-1.1)  

## <a href="#_special_use" class="anchor"></a>Special Use

- <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#extendingvulkan-compatibility-specialuse"
  target="_blank" rel="noopener">OpenGL / ES support</a>

## <a href="#_contact" class="anchor"></a>Contact

- James Fitzpatrick <a
  href="https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_IMG_relaxed_line_rasterization%5D%20@jamesfitzpatrick%0A*Here%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_IMG_relaxed_line_rasterization%20extension*"
  target="_blank" rel="nofollow noopener"><em></em>jamesfitzpatrick</a>

## <a href="#_other_extension_metadata" class="anchor"></a>Other Extension Metadata

**Last Modified Date**  
2023-10-22

**IP Status**  
No known IP claims.

**Contributors**  
- James Fitzpatrick, Imagination

- Andrew Garrard, Imagination

- Alex Walters, Imagination

## <a href="#_description" class="anchor"></a>Description

OpenGL specifies that implementations should rasterize lines using the
diamond exit rule (a slightly modified version of Bresenham’s
algorithm). To implement OpenGL some implementations have a device-level
compatibility mode to rasterize lines according to the OpenGL
specification.

This extension allows OpenGL emulation layers to enable the OpenGL
compatible line rasterization mode of such implementations.

## <a href="#_new_structures" class="anchor"></a>New Structures

- Extending [VkPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceFeatures2.html),
  [VkDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceCreateInfo.html):

  - [VkPhysicalDeviceRelaxedLineRasterizationFeaturesIMG](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceRelaxedLineRasterizationFeaturesIMG.html)

## <a href="#_new_enum_constants" class="anchor"></a>New Enum Constants

- `VK_IMG_RELAXED_LINE_RASTERIZATION_EXTENSION_NAME`

- `VK_IMG_RELAXED_LINE_RASTERIZATION_SPEC_VERSION`

- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html):

  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_RELAXED_LINE_RASTERIZATION_FEATURES_IMG`

## <a href="#_issues" class="anchor"></a>Issues

None.

## <a href="#_version_history" class="anchor"></a>Version History

- Revision 1, 2023-10-22 (James Fitzpatrick)

  - Initial version

## <a href="#_see_also" class="anchor"></a>See Also

No cross-references are available

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VK_IMG_relaxed_line_rasterization"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is a generated document. Fixes and changes should be made to
the generator scripts, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
