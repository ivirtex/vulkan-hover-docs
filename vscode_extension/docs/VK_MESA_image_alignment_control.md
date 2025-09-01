# VK\_MESA\_image\_alignment\_control(3) Manual Page

## Name

VK\_MESA\_image\_alignment\_control - device extension



## [](#_registered_extension_number)Registered Extension Number

576

## [](#_revision)Revision

1

## [](#_ratification_status)Ratification Status

Not ratified

## [](#_extension_and_version_dependencies)Extension and Version Dependencies

[VK\_KHR\_get\_physical\_device\_properties2](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_get_physical_device_properties2.html)  
or  
[Vulkan Version 1.1](#versions-1.1)

## [](#_special_use)Special Use

- [D3D support](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#extendingvulkan-compatibility-specialuse)

## [](#_contact)Contact

- Hans-Kristian Arntzen [\[GitHub\]HansKristian-Work](https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_MESA_image_alignment_control%5D%20%40HansKristian-Work%0A%2AHere%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_MESA_image_alignment_control%20extension%2A)

## [](#_other_extension_metadata)Other Extension Metadata

**Last Modified Date**

2024-05-03

**IP Status**

No known IP claims.

**Contributors**

- Hans-Kristian Arntzen, Valve

## [](#_description)Description

This extension allows applications to request a narrower alignment for images than an implementation would otherwise require. Some implementations internally support multiple image layouts in `VK_IMAGE_TILING_OPTIMAL`, each with different alignment requirements and performance trade-offs. In some API layering use cases such as D3D12, it is beneficial to be able to control the alignment, since certain alignments for placed resources are guaranteed to be supported, and emulating that expectation requires unnecessary padding of allocations.

[VkImageAlignmentControlCreateInfoMESA](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageAlignmentControlCreateInfoMESA.html) **can** be chained to [VkImageCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageCreateInfo.html), requesting that the alignment is no more than the provided alignment. If the requested alignment is not supported for a given [VkImageCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageCreateInfo.html), a larger alignment **may** be returned.

While something similar could be achieved with `VK_EXT_image_drm_format_modifier` in theory, this is not the intended way to use that extension. Format modifiers are generally used for externally shareable images, and would not be platform portable. It is also a cumbersome API to use just to lower the alignment.

## [](#_new_structures)New Structures

- Extending [VkImageCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageCreateInfo.html):
  
  - [VkImageAlignmentControlCreateInfoMESA](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageAlignmentControlCreateInfoMESA.html)
- Extending [VkPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceFeatures2.html), [VkDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceCreateInfo.html):
  
  - [VkPhysicalDeviceImageAlignmentControlFeaturesMESA](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceImageAlignmentControlFeaturesMESA.html)
- Extending [VkPhysicalDeviceProperties2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceProperties2.html):
  
  - [VkPhysicalDeviceImageAlignmentControlPropertiesMESA](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceImageAlignmentControlPropertiesMESA.html)

## [](#_new_enum_constants)New Enum Constants

- `VK_MESA_IMAGE_ALIGNMENT_CONTROL_EXTENSION_NAME`
- `VK_MESA_IMAGE_ALIGNMENT_CONTROL_SPEC_VERSION`
- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html):
  
  - `VK_STRUCTURE_TYPE_IMAGE_ALIGNMENT_CONTROL_CREATE_INFO_MESA`
  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_IMAGE_ALIGNMENT_CONTROL_FEATURES_MESA`
  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_IMAGE_ALIGNMENT_CONTROL_PROPERTIES_MESA`

## [](#_version_history)Version History

- Revision 1, 2024-04-05 (Hans-Kristian Arntzen)
  
  - Initial specification

## [](#_see_also)See Also

No cross-references are available

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VK_MESA_image_alignment_control).

This page is a generated document. Fixes and changes should be made to the generator scripts, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0