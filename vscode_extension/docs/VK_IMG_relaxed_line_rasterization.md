# VK\_IMG\_relaxed\_line\_rasterization(3) Manual Page

## Name

VK\_IMG\_relaxed\_line\_rasterization - device extension



## [](#_registered_extension_number)Registered Extension Number

111

## [](#_revision)Revision

1

## [](#_ratification_status)Ratification Status

Not ratified

## [](#_extension_and_version_dependencies)Extension and Version Dependencies

[VK\_KHR\_get\_physical\_device\_properties2](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_get_physical_device_properties2.html)  
or  
[Vulkan Version 1.1](#versions-1.1)

## [](#_special_use)Special Use

- [OpenGL / ES support](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#extendingvulkan-compatibility-specialuse)

## [](#_contact)Contact

- James Fitzpatrick [\[GitHub\]jamesfitzpatrick](https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_IMG_relaxed_line_rasterization%5D%20%40jamesfitzpatrick%0A%2AHere%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_IMG_relaxed_line_rasterization%20extension%2A)

## [](#_other_extension_metadata)Other Extension Metadata

**Last Modified Date**

2023-10-22

**IP Status**

No known IP claims.

**Contributors**

- James Fitzpatrick, Imagination
- Andrew Garrard, Imagination
- Alex Walters, Imagination

## [](#_description)Description

OpenGL specifies that implementations should rasterize lines using the diamond exit rule (a slightly modified version of Bresenham’s algorithm). To implement OpenGL some implementations have a device-level compatibility mode to rasterize lines according to the OpenGL specification.

This extension allows OpenGL emulation layers to enable the OpenGL compatible line rasterization mode of such implementations.

## [](#_new_structures)New Structures

- Extending [VkPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceFeatures2.html), [VkDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceCreateInfo.html):
  
  - [VkPhysicalDeviceRelaxedLineRasterizationFeaturesIMG](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceRelaxedLineRasterizationFeaturesIMG.html)

## [](#_new_enum_constants)New Enum Constants

- `VK_IMG_RELAXED_LINE_RASTERIZATION_EXTENSION_NAME`
- `VK_IMG_RELAXED_LINE_RASTERIZATION_SPEC_VERSION`
- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html):
  
  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_RELAXED_LINE_RASTERIZATION_FEATURES_IMG`

## [](#_issues)Issues

None.

## [](#_version_history)Version History

- Revision 1, 2023-10-22 (James Fitzpatrick)
  
  - Initial version

## [](#_see_also)See Also

No cross-references are available

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VK_IMG_relaxed_line_rasterization).

This page is a generated document. Fixes and changes should be made to the generator scripts, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0