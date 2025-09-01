# VK\_EXT\_image\_2d\_view\_of\_3d(3) Manual Page

## Name

VK\_EXT\_image\_2d\_view\_of\_3d - device extension



## [](#_registered_extension_number)Registered Extension Number

394

## [](#_revision)Revision

1

## [](#_ratification_status)Ratification Status

Ratified

## [](#_extension_and_version_dependencies)Extension and Version Dependencies

     [VK\_KHR\_maintenance1](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_maintenance1.html)  
     and  
     [VK\_KHR\_get\_physical\_device\_properties2](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_get_physical_device_properties2.html)  
or  
[Vulkan Version 1.1](#versions-1.1)

## [](#_special_use)Special Use

- [OpenGL / ES support](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#extendingvulkan-compatibility-specialuse)

## [](#_contact)Contact

- Mike Blumenkrantz [\[GitHub\]zmike](https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_EXT_image_2d_view_of_3d%5D%20%40zmike%0A%2AHere%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_EXT_image_2d_view_of_3d%20extension%2A)

## [](#_other_extension_metadata)Other Extension Metadata

**Last Modified Date**

2022-02-22

**IP Status**

No known IP claims.

**Contributors**

- Mike Blumenkrantz, Valve
- Piers Daniell, NVIDIA
- Spencer Fricke, Samsung
- Ricardo Garcia, Igalia
- Graeme Leese, Broadcom
- Ralph Potter, Samsung
- Stu Smith, AMD
- Shahbaz Youssefi, Google
- Alex Walters, Imagination

## [](#_description)Description

This extension allows a single slice of a 3D image to be used as a 2D view in image descriptors, matching both the functionality of glBindImageTexture in OpenGL with the `layer` parameter set to true and 2D view binding provided by the extension EGL\_KHR\_gl\_texture\_3D\_image. It is primarily intended to support GL emulation.

## [](#_new_structures)New Structures

- Extending [VkPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceFeatures2.html), [VkDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceCreateInfo.html):
  
  - [VkPhysicalDeviceImage2DViewOf3DFeaturesEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceImage2DViewOf3DFeaturesEXT.html)

## [](#_new_enum_constants)New Enum Constants

- `VK_EXT_IMAGE_2D_VIEW_OF_3D_EXTENSION_NAME`
- `VK_EXT_IMAGE_2D_VIEW_OF_3D_SPEC_VERSION`
- Extending [VkImageCreateFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageCreateFlagBits.html):
  
  - `VK_IMAGE_CREATE_2D_VIEW_COMPATIBLE_BIT_EXT`
- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html):
  
  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_IMAGE_2D_VIEW_OF_3D_FEATURES_EXT`

## [](#_version_history)Version History

- Revision 1, 2022-03-25 (Mike Blumenkrantz)
  
  - Internal revisions

## [](#_see_also)See Also

No cross-references are available

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VK_EXT_image_2d_view_of_3d).

This page is a generated document. Fixes and changes should be made to the generator scripts, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0