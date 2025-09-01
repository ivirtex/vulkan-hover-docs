# VK\_EXT\_shader\_image\_atomic\_int64(3) Manual Page

## Name

VK\_EXT\_shader\_image\_atomic\_int64 - device extension



## [](#_registered_extension_number)Registered Extension Number

235

## [](#_revision)Revision

1

## [](#_ratification_status)Ratification Status

Ratified

## [](#_extension_and_version_dependencies)Extension and Version Dependencies

[VK\_KHR\_get\_physical\_device\_properties2](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_get_physical_device_properties2.html)  
or  
[Vulkan Version 1.1](#versions-1.1)

## [](#_spir_v_dependencies)SPIR-V Dependencies

- [SPV\_EXT\_shader\_image\_int64](https://github.khronos.org/SPIRV-Registry/extensions/EXT/SPV_EXT_shader_image_int64.html)

## [](#_contact)Contact

- Tobias Hector [\[GitHub\]tobski](https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_EXT_shader_image_atomic_int64%5D%20%40tobski%0A%2AHere%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_EXT_shader_image_atomic_int64%20extension%2A)

## [](#_other_extension_metadata)Other Extension Metadata

**Last Modified Date**

2020-07-14

**IP Status**

No known IP claims.

**Interactions and External Dependencies**

- This extension provides API support for [`GLSL_EXT_shader_image_int64`](https://github.com/KhronosGroup/GLSL/blob/main/extensions/ext/GLSL_EXT_shader_image_int64.txt)

**Contributors**

- Matthaeus Chajdas, AMD
- Graham Wihlidal, Epic Games
- Tobias Hector, AMD
- Jeff Bolz, Nvidia
- Faith Ekstrand, Intel

## [](#_description)Description

This extension extends existing 64-bit integer atomic support to enable these operations on images as well.

When working with large 2- or 3-dimensional data sets (e.g. rasterization or screen-space effects), image accesses are generally more efficient than equivalent buffer accesses. This extension allows applications relying on 64-bit integer atomics in this manner to quickly improve performance with only relatively minor code changes.

64-bit integer atomic support is guaranteed for optimally tiled images with the `VK_FORMAT_R64_UINT` and `VK_FORMAT_R64_SINT` formats.

## [](#_new_structures)New Structures

- Extending [VkPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceFeatures2.html), [VkDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceCreateInfo.html):
  
  - [VkPhysicalDeviceShaderImageAtomicInt64FeaturesEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceShaderImageAtomicInt64FeaturesEXT.html)

## [](#_new_enum_constants)New Enum Constants

- `VK_EXT_SHADER_IMAGE_ATOMIC_INT64_EXTENSION_NAME`
- `VK_EXT_SHADER_IMAGE_ATOMIC_INT64_SPEC_VERSION`
- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html):
  
  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_SHADER_IMAGE_ATOMIC_INT64_FEATURES_EXT`

## [](#_version_history)Version History

- Revision 1, 2020-07-14 (Tobias Hector)
  
  - Initial draft

## [](#_see_also)See Also

No cross-references are available

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VK_EXT_shader_image_atomic_int64).

This page is a generated document. Fixes and changes should be made to the generator scripts, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0