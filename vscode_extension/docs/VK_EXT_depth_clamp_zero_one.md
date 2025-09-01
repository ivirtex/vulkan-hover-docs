# VK\_EXT\_depth\_clamp\_zero\_one(3) Manual Page

## Name

VK\_EXT\_depth\_clamp\_zero\_one - device extension



## [](#_registered_extension_number)Registered Extension Number

422

## [](#_revision)Revision

1

## [](#_ratification_status)Ratification Status

Ratified

## [](#_extension_and_version_dependencies)Extension and Version Dependencies

[VK\_KHR\_get\_physical\_device\_properties2](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_get_physical_device_properties2.html)  
or  
[Vulkan Version 1.1](#versions-1.1)

## [](#_deprecation_state)Deprecation State

- *Promoted* to [VK\_KHR\_depth\_clamp\_zero\_one](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_depth_clamp_zero_one.html) extension

## [](#_contact)Contact

- Graeme Leese [\[GitHub\]gnl21](https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_EXT_depth_clamp_zero_one%5D%20%40gnl21%0A%2AHere%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_EXT_depth_clamp_zero_one%20extension%2A)

## [](#_other_extension_metadata)Other Extension Metadata

**Last Modified Date**

2021-07-29

**Contributors**

- Graeme Leese, Broadcom

## [](#_description)Description

This extension gives defined behavior to fragment depth values which end up outside the conventional \[0, 1] range. It can be used to ensure portability in edge cases of features like depthBias. The particular behavior is chosen to match OpenGL to aid porting or emulation.

## [](#_promotion_to_vk_khr_depth_clamp_zero_one)Promotion to `VK_KHR_depth_clamp_zero_one`

All functionality in this extension is included in `VK_KHR_depth_clamp_zero_one`, with the suffix change to KHR. The original type, enum, and command names are still available as aliases of the core functionality.

## [](#_new_structures)New Structures

- Extending [VkPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceFeatures2.html), [VkDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceCreateInfo.html):
  
  - [VkPhysicalDeviceDepthClampZeroOneFeaturesEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceDepthClampZeroOneFeaturesEXT.html)

## [](#_new_enum_constants)New Enum Constants

- `VK_EXT_DEPTH_CLAMP_ZERO_ONE_EXTENSION_NAME`
- `VK_EXT_DEPTH_CLAMP_ZERO_ONE_SPEC_VERSION`
- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html):
  
  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_DEPTH_CLAMP_ZERO_ONE_FEATURES_EXT`

## [](#_version_history)Version History

- Revision 1, 2021-07-29 (Graeme Leese)
  
  - Internal revisions

## [](#_see_also)See Also

No cross-references are available

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VK_EXT_depth_clamp_zero_one).

This page is a generated document. Fixes and changes should be made to the generator scripts, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0